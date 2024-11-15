package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gale.MiddlewareFunc {
	return func(c gale.Ctx) error {
		unauthorized := gale.NewError(http.StatusUnauthorized, "Unauthorized")
		defer logrus.Info("Middleware: AuthMiddleware")

		token, err := c.Session().Get(utils.AuthSessionKey)
		if err != nil {
			return unauthorized
		}

		user, err := repository.FindUserBySessionID(db, string(token))
		if err != nil {
			return unauthorized
		}

		// Set user to context as a pointer
		// This state can be accessed by any upcoming handler
		// The data will be dropped after the request is done
		c.Set(utils.RequestAuthUserKey, &user)

		return nil
	}
}
