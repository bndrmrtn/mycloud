package services

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StorageServiceV1 struct {
	datadir    string
	db         *gorm.DB
	containers []container
	sizeLimit  int64
	fileLimit  int64

	mu sync.Mutex
}

type container struct {
	name  string
	files int64
	size  int64
}

func NewStorageServiceV1(datadir string, db *gorm.DB, sizeLimit int64, fileLimit int64) (StorageService, error) {
	var (
		svc StorageServiceV1
		err error
	)

	svc = StorageServiceV1{
		datadir:   datadir,
		db:        db,
		sizeLimit: sizeLimit,
		fileLimit: fileLimit,
	}

	logrus.Infof("initializing file storage service: sizeLimit: %d, fileLimit: %d", sizeLimit, fileLimit)
	// Walk through the data directory and count the number of files in each container
	err = filepath.Walk(datadir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() || !strings.Contains(info.Name(), "container") {
			return nil
		}

		svc.containers = append(svc.containers, container{
			name: info.Name(),
			size: info.Size(),
		})

		return filepath.Walk(filepath.Join(datadir, info.Name()), func(path string, i fs.FileInfo, err error) error {
			if i.IsDir() {
				return nil
			}
			svc.containers[len(svc.containers)-1].files++
			return nil
		})
	})

	return &svc, err
}

func (s *StorageServiceV1) StoreMultipartFile(f *multipart.FileHeader) (*models.OSFile, error) {
	multiFile, err := f.Open()
	if err != nil {
		return nil, err
	}

	defer multiFile.Close()

	return s.Store(multiFile, f.Size, filepath.Ext(f.Filename))
}

func (s *StorageServiceV1) Store(r io.Reader, size int64, ext string) (*models.OSFile, error) {
	// Read file data
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Compute hash to avoid duplicates
	sum := md5.Sum(data)
	hash := hex.EncodeToString(sum[:])

	// Check if the file already exists
	hFile, err := repository.FindOSFileByHash(s.db, hash)
	if err == nil {
		return hFile, nil
	}

	// Create a new file
	var file models.OSFile

	file.Extension = filepath.Ext(ext)
	file.Type = mime.TypeByExtension(file.Extension)
	file.FileSize = size
	file.Container, err = s.findContainer(size)
	file.ContentHash = hash
	// TODO: Detech charset for text files

	if err != nil {
		return nil, err
	}

	// Save the file to create an ID for it
	if err = s.db.Create(&file).Error; err != nil {
		return nil, err
	}

	// Write the file to the disk
	if err := os.WriteFile(filepath.Join(s.datadir, file.Container, file.ID), data, os.ModePerm); err != nil {
		// Rollback the file creation
		_ = s.Delete(&file)
		return nil, err
	}

	return &file, nil
}

func (s *StorageServiceV1) Delete(f *models.OSFile) error {
	can, err := repository.CanDeleteOSFile(s.db, f.ID)
	if err != nil || !can {
		return err
	}

	if err := os.Remove(filepath.Join(s.datadir, f.Container, f.ID)); err != nil {
		return err
	}

	go func(cont string, size int64) {
		s.mu.Lock()
		defer s.mu.Unlock()

		containerID, err := strconv.Atoi(strings.TrimPrefix(cont, "container"))
		if err != nil {
			logrus.Errorln("Failed to remove container from storage service:", cont)
			return
		}

		s.containers[containerID].files--
		s.containers[containerID].size -= size
	}(f.Container, f.FileSize)

	return nil
}

func (s *StorageServiceV1) GetTmpDir() string {
	return filepath.Join(s.datadir, "tmp")
}

func (s *StorageServiceV1) ReadFile(f *models.OSFile) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.datadir, f.Container, f.ID))
}

func (s *StorageServiceV1) GetRealPath(f *models.OSFile) (string, error) {
	path := filepath.Join(s.datadir, f.Container, f.ID)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, nil
}

func (s *StorageServiceV1) Containers() int {
	return len(s.containers)
}

func (s *StorageServiceV1) findContainer(size int64) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.containers {
		if s.containers[i].size+size <= s.sizeLimit && s.containers[i].files < s.fileLimit {
			s.containers[i].size += size
			s.containers[i].files++
			return "container" + strconv.Itoa(i), nil
		}
	}

	newDir := "container" + strconv.Itoa(len(s.containers))
	if err := os.MkdirAll(filepath.Join(s.datadir, newDir), 0755); err != nil {
		return "", err
	}

	s.containers = append(s.containers, container{
		files: 1,
		size:  size,
	})

	return newDir, nil
}
