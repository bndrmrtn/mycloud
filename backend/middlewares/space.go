package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SpaceMiddleware(db *gorm.DB, param string) bolt.MiddlewareFunc {
	return func(c bolt.Ctx) (bool, error) {
		defer logrus.Info("Middleware: SpaceMiddleware")

		unauthorized := bolt.NewError(http.StatusNotFound, "Space not found")

		space, err := repository.FindSpaceByID(db, c.Param(param))
		if err != nil {
			return false, unauthorized
		}

		c.Set(utils.RequestSpaceKey, &space)

		return true, nil
	}
}
