package config

import (
	"os"
	"strings"

	"github.com/bndrmrtn/go-gale"
	"github.com/sirupsen/logrus"
)

func Mode() gale.Mode {
	mode := strings.ToLower(os.Getenv("MODE"))
	switch mode {
	case "production":
		return gale.Production
	case "development":
		return gale.Development
	default:
		logrus.Warn("No MODE environment variable found, running in development mode")
		return gale.Development
	}
}
