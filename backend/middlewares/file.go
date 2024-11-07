package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func FileMiddleware(db *gorm.DB, param string) bolt.MiddlewareFunc {
	return func(c bolt.Ctx) (bool, error) {
		defer logrus.Info("Middleware: FileMiddleware")

		unauthorized := bolt.NewError(http.StatusNotFound, "File not found")

		file, err := repository.FindFileByID(db, c.Param(param))
		if err != nil {
			return false, unauthorized
		}

		c.Set(utils.RequestSpaceFileKey, &file)

		return true, nil
	}
}
