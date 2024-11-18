package handlers

import (
	"github.com/bndrmrtn/go-gale"
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
