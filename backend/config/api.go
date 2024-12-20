package config

import (
	"net/http"
	"time"

	"github.com/bndrmrtn/go-gale"
	"github.com/google/uuid"
)

func Api(store gale.SessionStore) gale.Config {
	return gale.Config{
		Session: &gale.SessionConfig{
			Enabled:     true,
			TokenExpire: time.Hour * 12,
			TokenFunc: func(c gale.Ctx) (string, error) {
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
