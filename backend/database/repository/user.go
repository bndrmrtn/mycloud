package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository/paginator"
	"gorm.io/gorm"
)

func FindUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func CheckEmailInWhitelist(db *gorm.DB, email string) (bool, error) {
	var whitelist int64
	result := db.Model(&models.UserWhitelist{}).Where("email = ?", email).Count(&whitelist)
	return whitelist > 0, result.Error
}

func CheckEmailPassBlacklist(db *gorm.DB, email string) (bool, error) {
	var blacklist int64
	result := db.Model(&models.UserBlacklist{}).Where("email = ?", email).Count(&blacklist)
	return blacklist < 1, result.Error
}

func PaginateUsers(db *gorm.DB, cursor string) (*paginator.Pagination[models.User], error) {
	data, err := paginator.Paginate[models.User](db.Model(&models.User{}), &paginator.Config{
		Cursor:     cursor,
		Order:      "desc",
		PointsNext: true,
		Limit:      15,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func FindUserByID(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	result := db.Where("id = ?", id).First(&user)
	return user, result.Error
}
