package services

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.OSFile{}, &models.File{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
