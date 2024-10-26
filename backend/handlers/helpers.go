package handlers

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/utils"
)

func ctxUser(c bolt.Ctx) (*models.User, error) {
	user := c.Get(utils.RequestAuthUserKey)
	if user != nil {
		return user.(*models.User), nil
	}

	return nil, bolt.NewError(http.StatusUnauthorized, "Failed to get user")
}

func ctxUserID(c bolt.Ctx) (string, error) {
	user, err := ctxUser(c)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func ctxSpace(c bolt.Ctx) (*models.FileSpace, error) {
	space := c.Get(utils.RequestSpaceKey)
	if space != nil {
		return space.(*models.FileSpace), nil
	}

	return nil, bolt.NewError(http.StatusNotFound, "Space not found")
}
