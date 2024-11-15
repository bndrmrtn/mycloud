package permissions

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"gorm.io/gorm"
)

func CanUserAccessFile(db *gorm.DB, user *models.User, file *models.File) bool {
	// This logic is simple for now, but it can be extended to check for more complex permissions
	can, err := repository.CanUserAccessFile(db, user.ID, file.ID)
	if err != nil {
		return false
	}
	return can
}
