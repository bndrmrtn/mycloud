package middlewares

import "github.com/bndrmrtn/go-bolt"

func CORSMiddleware(c bolt.Ctx) (bool, error) {
	c.Header().Add("Access-Control-Allow-Origin", "*")
	c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	return true, nil
}
