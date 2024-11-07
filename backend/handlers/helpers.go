package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

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

func ctxSpaceFile(c bolt.Ctx) (*models.File, error) {
	space := c.Get(utils.RequestSpaceFileKey)
	if space != nil {
		return space.(*models.File), nil
	}

	return nil, bolt.NewError(http.StatusNotFound, "File not found")
}

func wsWriter(ws bolt.WSServer, userID string, data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ws.BroadcastFunc(b, func(conn bolt.WSConn) bool {
		return conn.Ctx().Get(utils.WSUserID) == userID
	})

	return nil
}

func queryPath(c bolt.Ctx) string {
	path := c.URL().Query().Get("path")
	if path == "" || path == "." {
		path = "/"
	}

	return filepath.Clean(path)
}
