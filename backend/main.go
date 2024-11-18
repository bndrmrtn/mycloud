package main

import (
	"context"
	"flag"
	"log"

	_ "net/http/pprof"

	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database"
	"github.com/bndrmrtn/my-cloud/implementations"
	"github.com/bndrmrtn/my-cloud/services"
)

var listenAddr = flag.String("listenAddr", ":3000", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	// Load the .env file into the environment
	config.LoadDotEnv()

	// Read the app config
	conf, err := config.ReadAppConfig()
	if err != nil {
		log.Fatal("Failed to read app config: ", err)
	}

	// Connect to the database
	db, err := database.New(config.DBLogLevel())
	if err != nil {
		log.Fatalf("failed to connect to the database: %v\n", err)
	}

	// Create the redis session store
	store := implementations.NewRedisSessionStore(context.Background(), database.NewRedisClient())

	// Create the storage service
	sizeLimit, fileLimit := config.Containers()
	svc, err := services.NewStorageService(conf.Service.Version, conf.Service.AppdataDir, db, sizeLimit, fileLimit)
	if err != nil {
		log.Fatalf("failed to create storage service: %v", err)
	}

	// Create the API server
	api := NewApiServer(conf, db, store, svc)

	// Start the API server
	log.Fatal(api.Serve(*listenAddr))
}
