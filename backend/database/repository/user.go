package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func FindUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}
