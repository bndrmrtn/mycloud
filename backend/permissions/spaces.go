package permissions

import (
	"fmt"

	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CanUserAccessSpace(rdb *redis.Client, db *gorm.DB, user *models.User, space *models.FileSpace) bool {
	var key = fmt.Sprintf("permission:space:%s-%s", user.ID, space.ID)

	// Check if the cache has the value
	ok, err := redisBoolReturn(rdb, key)
	if err == nil {
		return ok
	}

	var count int64
	result := db.Model(&models.FileSpace{}).
		Where("file_spaces.id = ? and (file_spaces.user_id = ? or exists (select 1 from space_user where space_user.file_space_id = file_spaces.id and space_user.user_id = ?))", space.ID, user.ID, user.ID).
		Count(&count)

	return redisBoolSave(rdb, result.Error == nil && count > 0, key)
}

func CanUserUploadFile(rdb *redis.Client, db *gorm.DB, user *models.User, space *models.FileSpace) bool {
	var key = fmt.Sprintf("permission:space.create:%s-%s", user.ID, space.ID)

	// Check if the cache has the value
	ok, err := redisBoolReturn(rdb, key)
	if err == nil {
		return ok
	}

	if user.ID == space.UserID {
		return true
	}

	spaceUser, err := spaceUserPerm(db, user.ID, space.ID)
	if err != nil {
		return false
	}

	return redisBoolSave(rdb, spaceUser.Permission.UploadFile, key)
}

func CanUserReadFile(rdb *redis.Client, db *gorm.DB, user *models.User, space *models.FileSpace) bool {
	var key = fmt.Sprintf("permission:space.read:%s-%s", user.ID, space.ID)

	// Check if the cache has the value
	ok, err := redisBoolReturn(rdb, key)
	if err == nil {
		return ok
	}

	if user.ID == space.UserID {
		return true
	}

	spaceUser, err := spaceUserPerm(db, user.ID, space.ID)
	if err != nil {
		return false
	}

	return redisBoolSave(rdb, spaceUser.Permission.ReadFile, key)
}

func spaceUserPerm(db *gorm.DB, userID, spaceID string) (*models.SpaceUser, error) {
	var spaceUser models.SpaceUser

	result := db.Model(&models.SpaceUser{}).
		Where("space_user.user_id = ? and space_user.file_space_id = ?", userID, spaceID).
		First(&spaceUser)

	return &spaceUser, result.Error
}
