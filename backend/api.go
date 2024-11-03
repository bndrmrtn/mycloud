package main

import (
	"errors"
	"strings"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/handlers"
	"github.com/bndrmrtn/my-cloud/middlewares"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm"
)

func NewApiServer(db *gorm.DB, store bolt.SessionStore, svc services.StorageService) *bolt.Bolt {
	conf := config.Api(store)
	conf.Mode = config.Mode()

	app := bolt.New(&conf)
	ws := NewWSServer(app, db)

	registerValidators(app.CompleteRouter)
	app.Hook(bolt.PreRequestHook, func(c bolt.Ctx) {
		middlewares.CORSMiddleware(c)
	})
	registerRoutes(app, db, store, svc, ws)

	return app
}

func registerRoutes(r bolt.Router, db *gorm.DB, store bolt.SessionStore, svc services.StorageService, ws bolt.WSServer) {
	r.Get("/auth-redirect", handlers.HandleCreateAuthURL).Name("auth.redirect")
	r.Get("/gauth", handlers.HandleAuthUser(db, svc)).Name("auth.google")
	r.Get("/profileimage/{id@png}", handlers.HandleGetProfileImage(db, svc, store))

	auth := r.Group("/", middlewares.AuthMiddleware(db))

	auth.Get("/me", handlers.HandleGetAuthUser)
	auth.Get("/logout", handlers.HandleLogout)

	// Manage spaces
	{
		auth.Get("/spaces", handlers.HandleGetSpaces(db)).Name("spaces.all")
		auth.Post("/spaces", handlers.HandleCreateSpace(db)).Name("spaces.create")
	}

	// Manage files in a space
	// TODO: Go Bolt does not chain middlewares with groups
	spaces := r.Group("/spaces/{space_id@uuid}", middlewares.AuthMiddleware(db), middlewares.SpaceMiddleware(db, "space_id"))
	{
		spaces.Get("/", handlers.HandleGetSpace(db)).Name("spaces.get")
		spaces.Get("/files", handlers.HandleGetFiles(db)).Name("spaces.files")
		spaces.Get("/fs", handlers.HandleGetFS(db))
		spaces.Get("/download", handlers.HandleDownloadDir(db, svc, ws))

		spaces.Post("/upload", handlers.HandleUploadFile(db, svc)).Name("spaces.upload")
	}
}

func registerValidators(r bolt.CompleteRouter) {
	r.RegisterRouteParamValidator("png", func(value string) (string, error) {
		strs := strings.SplitN(value, ".", 2)
		if strs[len(strs)-1] != "png" {
			return "", errors.New("invalid file extension")
		}

		return strs[0], nil
	})
}
