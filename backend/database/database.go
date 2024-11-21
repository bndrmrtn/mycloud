package database

import (
	"fmt"
	"log"
	"os"

	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(loglevel logger.LogLevel) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %v\n", err.Error())
	}

	log.Println("successfully connected to the database!")
	db.Logger = logger.Default.LogMode(loglevel)

	log.Println("Running migrations")
	err = db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.FileSpace{},
		&models.File{},
		&models.OSFile{},
		&models.Download{},
		&models.ImageURL{},
		&models.UserWhitelist{},
		&models.UserBlacklist{},
		&models.SpaceUser{},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to migrate: " + err.Error())
	}

	return db, nil
}
