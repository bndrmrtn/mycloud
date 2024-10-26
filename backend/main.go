package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database"
	"github.com/bndrmrtn/my-cloud/implementations"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm/logger"
)

var listenAddr = flag.String("listenAddr", ":3000", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	db, err := database.New(logger.Info)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v\n", err)
	}

	sizeLimit, fileLimit := config.Containers()
	svc, err := services.NewStorageServiceV1(os.Getenv("DATADIR"), db, sizeLimit, fileLimit)
	if err != nil {
		log.Fatalf("failed to create storage service: %v", err)
	}

	store := implementations.NewRedisSessionStore(context.Background(), database.NewRedisClient())
	api := NewApiServer(db, store, svc)

	log.Fatal(api.Serve(*listenAddr))
}
