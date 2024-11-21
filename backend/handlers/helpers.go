package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"gorm.io/gorm"
)

func ctxUser(c gale.Ctx) (*models.User, error) {
	user := c.Get(utils.RequestAuthUserKey)
	if user != nil {
		return user.(*models.User), nil
	}

	return nil, gale.NewError(http.StatusUnauthorized, "Failed to get user")
}

func ctxUserID(c gale.Ctx) (string, error) {
	user, err := ctxUser(c)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func ctxSpace(c gale.Ctx) (*models.FileSpace, error) {
	space := c.Get(utils.RequestSpaceKey)
	if space != nil {
		return space.(*models.FileSpace), nil
	}

	return nil, gale.NewError(http.StatusNotFound, "Space not found")
}

func ctxSpaceFile(c gale.Ctx) (*models.File, error) {
	space := c.Get(utils.RequestSpaceFileKey)
	if space != nil {
		return space.(*models.File), nil
	}

	return nil, gale.NewError(http.StatusNotFound, "File not found")
}

// TODO: Send realtime events to all users in the same space (even participants)
func wsWriter(ws gale.WSServer, userID string, data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ws.BroadcastFunc(b, func(conn gale.WSConn) bool {
		return conn.Ctx().Get(utils.WSUserID) == userID
	})

	return nil
}

func queryPath(c gale.Ctx) string {
	path := c.URL().Query().Get("path")
	if path == "" || path == "." {
		path = "/"
	}

	return filepath.Clean(path)
}

func checkPrimaryAdmin(db *gorm.DB, id string, conf *config.AdminCofig) error {
	user, err := repository.FindUserByID(db, id)
	if err != nil {
		return err
	}

	if user.Email == conf.PrimaryAdminEmail {
		return gale.NewError(400, "You can't delete the primary admin")
	}

	return nil
}
