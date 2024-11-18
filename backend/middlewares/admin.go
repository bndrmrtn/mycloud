package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AdminMiddleware(db *gorm.DB) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		defer logrus.Info("Middleware: AdminMiddleware")

		user := c.Get(utils.RequestAuthUserKey).(*models.User)

		if user.Role != models.RoleAdmin {
			return gale.NewError(http.StatusForbidden, "User is not an admin")
		}

		return nil
	}
}
