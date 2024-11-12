package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func FileMiddleware(db *gorm.DB, param string) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		defer logrus.Info("Middleware: FileMiddleware")

		unauthorized := gale.NewError(http.StatusNotFound, "File not found")

		file, err := repository.FindFileByID(db, c.Param(param))
		if err != nil {
			return unauthorized
		}

		c.Set(utils.RequestSpaceFileKey, &file)

		return nil
	}
}
