package middlewares

import "github.com/bndrmrtn/go-bolt"

func CORSMiddleware(c bolt.Ctx) (bool, error) {
	c.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000") // Portot is add meg, ha szükséges
	c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header().Add("Access-Control-Allow-Credentials", "true")
	return true, nil
}
