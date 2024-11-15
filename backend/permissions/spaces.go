package permissions

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func CanUserAccessSpace(_ *gorm.DB, user *models.User, space *models.FileSpace) bool {
	// This logic is simple for now, but it can be extended to check for more complex permissions
	return user.ID == space.UserID
}
