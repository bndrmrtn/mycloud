package handlers

import (
	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
)

func HandleIndexRoute(conf *config.AppConfig) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		return c.JSON(conf)
	}
}
