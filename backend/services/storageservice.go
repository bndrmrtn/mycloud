package services

import (
	"errors"
	"io"
	"mime/multipart"

	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

type StorageService interface {
	Store(r io.Reader, size int64, ext string) (*models.OSFile, error)
	StoreMultipartFile(f *multipart.FileHeader) (*models.OSFile, error)

	Delete(f *models.OSFile) error
	ReadFile(f *models.OSFile) ([]byte, error)

	GetTmpDir() string
	GetRealPath(f *models.OSFile) (string, error)

	Containers() int
}

func NewStorageService(version, datadir string, db *gorm.DB, sizeLimit int64, fileLimit int64) (StorageService, error) {
	switch version {
	case "1":
		return NewStorageServiceV1(datadir, db, sizeLimit, fileLimit)
	default:
		return nil, errors.New("unknown storage service version")
	}
}
