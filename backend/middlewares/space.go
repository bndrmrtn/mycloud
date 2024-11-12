package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SpaceMiddleware(db *gorm.DB, param string) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		defer logrus.Info("Middleware: SpaceMiddleware")

		unauthorized := gale.NewError(http.StatusNotFound, "Space not found")

		space, err := repository.FindSpaceByID(db, c.Param(param))
		if err != nil {
			return unauthorized
		}

		c.Set(utils.RequestSpaceKey, &space)

		return nil
	}
}
