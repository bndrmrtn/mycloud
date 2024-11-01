package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func GetAllSpacesForUser(db *gorm.DB, userID string) ([]models.FileSpace, error) {
	var spaces []models.FileSpace
	result := db.Where("user_id = ?", userID).Find(&spaces)
	return spaces, result.Error
}

func FindSpaceByID(db *gorm.DB, id string) (models.FileSpace, error) {
	var space models.FileSpace
	result := db.Where("id = ?", id).First(&space)
	return space, result.Error
}

func GetSpaceFiles(db *gorm.DB, spaceID string, dir string) ([]models.File, error) {
	var files []models.File
	result := db.Model(&models.File{}).Where("file_space_id = ? and directory = ?", spaceID, dir).Preload("OSFile").Find(&files)
	return files, result.Error
}

func GetAllSpaceFiles(db *gorm.DB, spaceID string, dir string) ([]models.File, error) {
	var files []models.File
	result := db.Model(&models.File{}).Where("file_space_id = ? and directory like ?", spaceID, dir+"%").Preload("OSFile").Find(&files)
	return files, result.Error
}

func GetSpaceFS(db *gorm.DB, spaceID string, dir string) ([]string, error) {
	var files []string
	result := db.Raw(`
		SELECT DISTINCT
            SUBSTRING_INDEX(SUBSTRING_INDEX(directory, '/', LENGTH(TRIM(TRAILING '/' FROM ?)) - LENGTH(REPLACE(TRIM(TRAILING '/' FROM ?), '/', '')) + 2), '/', -1) AS next_directory
        FROM files
        WHERE directory LIKE CONCAT(TRIM(TRAILING '/' FROM ?), '/%')
            AND directory != '/'
            AND file_space_id = ?
            AND LENGTH(TRIM(TRAILING '/' FROM ?)) < LENGTH(REPLACE(TRIM(TRAILING '/' FROM directory), '/', ''));
		`, dir, dir, dir, spaceID, dir).Find(&files)
	return files, result.Error
}
