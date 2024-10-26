package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if d := os.Getenv("DATADIR"); d == "" {
		os.Setenv("DATADIR", "mycloud-appdata")
	} else {
		os.Setenv("DATADIR", filepath.Clean(d))
	}

	if err := os.MkdirAll(os.Getenv("DATADIR"), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(os.Getenv("DATADIR"), "tmp"), 0755); err != nil {
		return err
	}

	return nil
}
