package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PermissionChecker func(rdb *redis.Client, db *gorm.DB, user *models.User, fs *models.FileSpace) bool

func SpaceMiddleware(rdb *redis.Client, db *gorm.DB, param string, permFn PermissionChecker) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		defer logrus.Info("Middleware: SpaceMiddleware")

		notfound := gale.NewError(http.StatusNotFound, "Space not found")

		space, err := repository.FindSpaceByID(db, c.Param(param))
		if err != nil {
			return notfound
		}

		// Check if user has access to the space
		user := c.Get(utils.RequestAuthUserKey).(*models.User)
		if !permFn(rdb, db, user, &space) {
			return gale.NewError(http.StatusForbidden, "Forbidden")
		}

		// Set space to context as a pointer
		c.Set(utils.RequestSpaceKey, &space)

		return nil
	}
}
