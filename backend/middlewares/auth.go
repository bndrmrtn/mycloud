package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) bolt.MiddlewareFunc {
	return func(c bolt.Ctx) (bool, error) {
		unauthorized := bolt.NewError(http.StatusUnauthorized, "Unauthorized")
		defer logrus.Info("Middleware: AuthMiddleware")

		token, err := c.Session().Get(utils.AuthSessionKey)
		if err != nil {
			return false, unauthorized
		}

		user, err := repository.FindUserBySessionID(db, string(token))
		if err != nil {
			return false, unauthorized
		}

		// Set user to context as a pointer
		// This state can be accessed by any upcoming handler
		// The data will be dropped after the request is done
		c.Set(utils.RequestAuthUserKey, &user)

		return true, nil
	}
}

func WSAuthMiddleware(db *gorm.DB) bolt.MiddlewareFunc {
	return func(c bolt.Ctx) (bool, error) {
		unauthorized := bolt.NewError(http.StatusUnauthorized, "Unauthorized")
		defer logrus.Info("Middleware: AuthMiddleware")

		token, err := c.Session().From(c.URL().Query().Get("auth")).Get(utils.AuthSessionKey)
		if err != nil {
			return false, unauthorized
		}

		user, err := repository.FindUserBySessionID(db, string(token))
		if err != nil {
			return false, unauthorized
		}

		// Set user to context as a pointer
		// This state can be accessed by any upcoming handler
		// The data will be dropped after the request is done
		c.Set(utils.RequestAuthUserKey, &user)

		return true, nil
	}
}
