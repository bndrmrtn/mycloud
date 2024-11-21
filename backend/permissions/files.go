package permissions

import (
	"fmt"

	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CanUserAccessFile(rdb *redis.Client, db *gorm.DB, user *models.User, file *models.File) bool {
	var key = fmt.Sprintf("permission:file.read:%s-%s", user.ID, file.ID)

	// Check if the cache has the value
	ok, err := redisBoolReturn(rdb, key)
	if err == nil {
		return ok
	}

	if user.ID == file.UserID {
		// No need to cache this, the user is the owner of the file
		return true
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
		file.ID, user.ID, user.ID, models.ReadFileBit).Scan(&can)

	if result.Error != nil {
		return false
	}

	return redisBoolSave(rdb, can, key)
}
