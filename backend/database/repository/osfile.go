package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func FindOSFileByHash(db *gorm.DB, h string) (*models.OSFile, error) {
	var f models.OSFile
	err := db.Where("content_hash = ?", h).First(&f).Error
	return &f, err
}

func CanDeleteOSFile(db *gorm.DB, osFileID string) (bool, error) {
	var files int64
	result := db.Model(&models.File{}).InnerJoins("OSFile").Where("os_files.id = ?", osFileID).Limit(2).Count(&files)
	return files < 2, result.Error
}
