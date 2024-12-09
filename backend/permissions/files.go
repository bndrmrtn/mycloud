package permissions

import (
	"fmt"

	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CanUserAccessFile(rdb *redis.Client, db *gorm.DB, user *models.User, file *models.File) bool {
	var key = fmt.Sprintf(FileReadFormat, user.ID, file.ID)
	return checkFilePerm(db, rdb, key, user.ID, file, models.ReadFileBit)
}

func CanUserDeleteFile(rdb *redis.Client, db *gorm.DB, user *models.User, file *models.File) bool {
	var key = fmt.Sprintf(FileDeleteFormat, user.ID, file.ID)
	return checkFilePerm(db, rdb, key, user.ID, file, models.DeleteFileBit)
}

func CanUserUpdateFile(rdb *redis.Client, db *gorm.DB, user *models.User, file *models.File) bool {
	var key = fmt.Sprintf(FileUpdateFormat, user.ID, file.ID)
	return checkFilePerm(db, rdb, key, user.ID, file, models.UpdateFileBit)
}

func checkFilePerm(db *gorm.DB, rdb *redis.Client, key string, userID string, file *models.File, bit int) bool {
	if userID == file.UserID {
		// No need to cache this, the user is the owner of the file
		return true
	}

	// Check if the cache has the value
	ok, err := redisBoolReturn(rdb, key)
	if err == nil {
		return ok
	}

	var can bool
	result := db.Raw(`
		SELECT EXISTS (
			SELECT 1
			FROM space_user
			JOIN file_spaces ON file_spaces.id = space_user.file_space_id
			JOIN files ON files.file_space_id = file_spaces.id
			WHERE files.id = ?
			AND (file_spaces.user_id = ?
				OR (space_user.user_id = ? AND (space_user.permission_int & ? != 0))
			)
		)`,
		file.ID, userID, userID, bit).Scan(&can)

	if result.Error != nil {
		return false
	}

	return redisBoolSave(rdb, can, key)
}
