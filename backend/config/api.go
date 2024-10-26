package config

import (
	"net/http"
	"time"

	"github.com/bndrmrtn/go-bolt"
	"github.com/google/uuid"
)

func Api(store bolt.SessionStore) bolt.Config {
	return bolt.Config{
		NotFoundHandler: func(c bolt.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(bolt.Map{
				"error": "Not Found",
			})
		},
		Session: &bolt.SessionConfig{
			Enabled:     true,
			TokenExpire: time.Hour * 12,
			TokenFunc: func(c bolt.Ctx) (string, error) {
				cookie, err := c.Cookie().Get("session")
				if err != nil {
					token := uuid.New().String()
					c.Cookie().Set(&http.Cookie{
						Name:    "session",
						Value:   token,
						Expires: time.Now().Add(time.Hour * 12),
					})
					return token, nil
				}
				return cookie.Value, nil
			},
			Store: store,
		},
	}
}
