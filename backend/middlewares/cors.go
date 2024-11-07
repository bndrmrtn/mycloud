package middlewares

import (
	"net/http"

	"github.com/bndrmrtn/go-bolt"
)

func CORSMiddleware(c bolt.Ctx) (bool, error) {
	c.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header().Add("Access-Control-Allow-Credentials", "true")

	if c.Method() == "OPTIONS" {
		c.Status(204) // No Content
		return false, bolt.NewError(http.StatusOK, "")
	}

	return true, nil
}
