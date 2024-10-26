package services

import (
	"io"
	"mime/multipart"

	"github.com/bndrmrtn/my-cloud/database/models"
)

type StorageService interface {
	Store(r io.Reader, size int64, ext string) (*models.OSFile, error)
	StoreMultipartFile(f *multipart.FileHeader) (*models.OSFile, error)
	Delete(f *models.OSFile) error
	ReadFile(f *models.OSFile) ([]byte, error)
	GetTmpDir() string
	GetRealPath(f *models.OSFile) string
}
