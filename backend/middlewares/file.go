package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/permissions"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func FileMiddleware(rdb *redis.Client, db *gorm.DB, param string) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		defer logrus.Info("Middleware: FileMiddleware")

		unauthorized := gale.NewError(http.StatusNotFound, "File not found")

		file, err := repository.FindFileByID(db, c.Param(param))
		if err != nil {
			return unauthorized
		}

		user := c.Get(utils.RequestAuthUserKey).(*models.User)

		if !permissions.CanUserAccessFile(rdb, db, user, &file) {
			return gale.NewError(http.StatusForbidden, "Forbidden")
		}

		c.Set(utils.RequestSpaceFileKey, &file)

		return nil
	}
}
