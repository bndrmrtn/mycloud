package handlers

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"gorm.io/gorm"
)

type CreateSpacesRequest struct {
	Name string
}

func HandleCreateSpace(db *gorm.DB) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		var data CreateSpacesRequest

		if err := c.Body().ParseJSON(&data); err != nil {
			return err
		}

		if len(data.Name) < 1 {
			return bolt.NewError(http.StatusBadRequest, "Name cannot be empty")
		}

		if len(data.Name) > 50 {
			return bolt.NewError(http.StatusBadRequest, "Name cannot be longer than 50 characters")
		}

		userID, err := ctxUserID(c)
		if err != nil {
			return err
		}

		space := models.FileSpace{
			HasUser: models.HasUserID(userID),
			Name:    data.Name,
		}
		db.Create(&space)

		return c.JSON(space)
	}
}

func HandleGetSpaces(db *gorm.DB) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		user, err := ctxUser(c)
		if err != nil {
			return err
		}

		spaces, err := repository.GetAllSpacesForUser(db, user.ID)
		if err != nil {
			return err
		}

		return c.JSON(spaces)
	}
}
