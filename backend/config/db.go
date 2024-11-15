package config

import (
	"os"

	"gorm.io/gorm/logger"
)

func DBLogLevel() logger.LogLevel {
	lvl := os.Getenv("DB_LOG_LEVEL")
	switch lvl {
	case "SILENT":
		return logger.Silent
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	default:
		return logger.Error
	}
}
