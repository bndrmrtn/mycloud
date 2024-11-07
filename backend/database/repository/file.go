package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func IsFileExists(db *gorm.DB, spaceID string, dir string, name string) (bool, error) {
	var count int64
	result := db.Model(&models.File{}).Where("file_space_id = ? and directory = ? and file_name = ?", spaceID, dir, name).Count(&count)
	return count > 0, result.Error
}

func FindFileByID(db *gorm.DB, id string) (models.File, error) {
	var file models.File
	result := db.Where("id = ?", id).Preload("OSFile").First(&file)
	return file, result.Error
}
