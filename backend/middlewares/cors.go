package middlewares

import "github.com/bndrmrtn/go-gale"

func CORSMiddleware(c gale.Ctx) error {
	c.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header().Add("Access-Control-Allow-Credentials", "true")

	if c.Method() == "OPTIONS" {
		c.Status(204).Break() // No Content
		return nil
	}

	return nil
}
