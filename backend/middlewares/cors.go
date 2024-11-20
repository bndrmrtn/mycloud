package middlewares

import (
	"slices"

	"github.com/bndrmrtn/go-gale"
	"github.com/sirupsen/logrus"
)

var allowedOrigins = []string{"http://localhost:3000"}

func CORSMiddleware(c gale.Ctx) error {
	defer logrus.Info("Middleware: CORSMiddleware")

	origin := c.Request().Header.Get("Origin")

	if slices.Contains(allowedOrigins, origin) {
		c.Header().Add("Access-Control-Allow-Origin", origin)
		c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header().Add("Access-Control-Allow-Credentials", "true")
	}

	if c.Method() == "OPTIONS" {
		c.Status(204).Break() // No Content
		return c.Send(nil)
	}

	return nil
}
