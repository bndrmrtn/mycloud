package config

import (
	"os"
	"strings"

	"github.com/bndrmrtn/go-bolt"
	"github.com/sirupsen/logrus"
)

func Mode() bolt.Mode {
	mode := strings.ToLower(os.Getenv("MODE"))
	switch mode {
	case "production":
		return bolt.Production
	case "development":
		return bolt.Development
	default:
		logrus.Warn("No MODE environment variable found, running in development mode")
		return bolt.Development
	}
}
