package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func IsFileExists(db *gorm.DB, spaceID, fileID, dir, name string) (bool, error) {
	var count int64
	result := db.Model(&models.File{}).Where("file_space_id = ? and id != ? and directory = ? and file_name = ?", spaceID, fileID, dir, name).Count(&count)
	return count > 0, result.Error
}

func FindFileByID(db *gorm.DB, id string) (models.File, error) {
	var file models.File
	result := db.Where("id = ?", id).Preload("OSFile").First(&file)
	return file, result.Error
}

func CanUserAccessFile(db *gorm.DB, userID, fileID string) (bool, error) {
	var count int64
	result := db.Model(&models.File{}).Joins("JOIN file_spaces ON file_spaces.id = files.file_space_id").Where("files.id = ? and file_spaces.user_id = ?", fileID, userID).Count(&count)
	return count > 0, result.Error
}
