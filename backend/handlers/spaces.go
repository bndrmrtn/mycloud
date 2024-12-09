package handlers

import (
	"errors"
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/handlers/dao"
	"github.com/bndrmrtn/my-cloud/permissions"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CreateSpacesRequest struct {
	Name string `json:"name"`
}

func HandleGetSpace(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		space, err := ctxSpace(c)
		if err != nil {
			return err
		}
		return c.JSON(space)
	}
}

func HandleCreateSpace(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		var data CreateSpacesRequest

		if err := c.Body().ParseJSON(&data); err != nil {
			return err
		}

		if len(data.Name) < 1 {
			return gale.NewError(http.StatusBadRequest, "Name cannot be empty")
		}

		if len(data.Name) > 50 {
			return gale.NewError(http.StatusBadRequest, "Name cannot be longer than 50 characters")
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

		return c.JSON(space.WithSize(0))
	}
}

func HandleGetSpaces(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
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

func HandleGetCollaborators(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		collaborators, err := repository.GetSpaceCollaborators(db, space.ID)
		if err != nil {
			return err
		}

		return c.JSON(collaborators)
	}
}

func HandleUpdateCollaborator(db *gorm.DB, rdb *redis.Client) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		user, err := ctxUser(c)
		if err != nil {
			return err
		}

		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		if user.ID != space.UserID {
			return gale.NewError(http.StatusForbidden, "Only the owner can manage collaborators")
		}

		var data dao.Collaborator
		if err := c.Body().ParseJSON(&data); err != nil {
			return err
		}

		if data.Email == user.Email {
			return gale.NewError(http.StatusBadRequest, "You cannot manage yourself")
		}

		collaboratorUser, err := repository.FindUserByEmail(db, data.Email)
		if err != nil {
			return err
		}

		if data.DoRemove() {
			err = db.Delete(&models.SpaceUser{}, "file_space_id = ? and user_id = ?", space.ID, collaboratorUser.ID).Error
			if err != nil {
				return err
			}
			permissions.CleanUp(rdb, collaboratorUser.ID, space.ID)
			return c.JSON(gale.Map{
				"deleted": true,
			})
		}

		collaborator, err := repository.FindSpaceCollaborator(db, space.ID, collaboratorUser.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return createCollaborator(c, db, space, &collaboratorUser, data)
			}
			return err
		}

		collaborator.Permission.UploadFile = data.Permission.Create
		collaborator.Permission.ReadFile = data.Permission.Read
		collaborator.Permission.UpdateFile = data.Permission.Update
		collaborator.Permission.DeleteFile = data.Permission.Delete

		if err := db.Save(&collaborator).Error; err != nil {
			return err
		}

		permissions.CleanUp(rdb, collaboratorUser.ID, space.ID)

		return c.JSON(gale.Map{
			"updated": true,
		})
	}
}

func createCollaborator(c gale.Ctx, db *gorm.DB, space *models.FileSpace, user *models.User, data dao.Collaborator) error {
	collaborator := models.SpaceUser{
		HasFileSpace: models.HasFileSpaceID(space.ID),
		HasUser:      models.HasUserID(user.ID),
		Permission: &models.SpaceUserPermission{
			UploadFile: data.Permission.Create,
			ReadFile:   data.Permission.Read,
			UpdateFile: data.Permission.Update,
			DeleteFile: data.Permission.Delete,
		},
	}

	if err := db.Create(&collaborator).Error; err != nil {
		return err
	}

	return c.JSON(collaborator)
}
