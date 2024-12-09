package handlers

import (
	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"gorm.io/gorm"
)

func HandleAdminGetUsers(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		data, err := repository.PaginateUsers(db, c.URL().Query().Get("cursor"))
		if err != nil {
			return err
		}

		return c.JSON(data)
	}
}

func HandleAdminDeleteUser(db *gorm.DB, conf *config.AdminCofig) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		// Get id from URL
		id := c.Param("userID")

		// Check if the user is trying to delete themselves
		uid, err := ctxUserID(c)
		if err != nil {
			return err
		}

		if uid == id {
			return gale.NewError(400, "You can't delete yourself")
		}

		// Check if the user is trying to delete the primary admin
		if err := checkPrimaryAdmin(db, id, conf); err != nil {
			return err
		}

		// Delete the user
		if err := db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
			return err
		}

		return c.JSON("User deleted")
	}
}

func HandleAdminGetWhitelist(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		data, err := repository.PaginateList[models.UserWhitelist](db, c.URL().Query().Get("cursor"))
		if err != nil {
			return err
		}

		return c.JSON(data)
	}
}

func HandleAdminGetBlacklist(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		data, err := repository.PaginateList[models.UserBlacklist](db, c.URL().Query().Get("cursor"))
		if err != nil {
			return err
		}

		return c.JSON(data)
	}
}

func HandleAdminGetAnalytics(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		osFilesContainer, _ := repository.GetOSFilesSeparatedByContainers(db)
		osFileSizeDiff, _ := repository.GetSizeDiff(db)

		return c.JSON(gale.Map{
			"os_file_container": osFilesContainer,
			"file_difference":   osFileSizeDiff,
		})
	}
}
