package main

import (
	"errors"
	"strings"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/handlers"
	"github.com/bndrmrtn/my-cloud/middlewares"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm"
)

func NewApiServer(db *gorm.DB, store gale.SessionStore, svc services.StorageService) *gale.Gale {
	conf := config.Api(store)
	conf.Mode = config.Mode()

	app := gale.New(&conf)
	ws := NewWSServer(app, db)

	registerValidators(app.CompleteRouter)
	registerRoutes(app.Group("/", middlewares.CORSMiddleware), db, store, svc, ws)

	if conf.Mode == gale.Development {
		// Add devtools in development mode
		app.Use(gale.NewUIDevtools())
		// Register pprof routes
		middlewares.RegisterPprof(app.Router())
		// Dump the routes in development mode
		app.Dump()
	}

	return app
}

func registerRoutes(r gale.Router, db *gorm.DB, store gale.SessionStore, svc services.StorageService, ws gale.WSServer) {
	r.Get("/auth-redirect", handlers.HandleCreateAuthURL).Name("auth.redirect")
	r.Get("/gauth", handlers.HandleAuthUser(db, svc)).Name("auth.google")
	r.Get("/profileimage/{id@png}", handlers.HandleGetProfileImage(db, svc, store)).Name("cdn.profileimage")

	auth := r.Group("/", middlewares.AuthMiddleware(db))

	auth.Get("/me", handlers.HandleGetAuthUser).Name("auth.me")
	auth.Get("/logout", handlers.HandleLogout).Name("auth.logout")

	// Manage spaces
	{
		auth.Get("/spaces", handlers.HandleGetSpaces(db)).Name("spaces.all")
		auth.Post("/spaces", handlers.HandleCreateSpace(db)).Name("spaces.create")
	}

	// Manage file spaces
	spaces := auth.Group("/spaces/{space_id@uuid}", middlewares.SpaceMiddleware(db, "space_id"))
	{
		spaces.Get("/", handlers.HandleGetSpace(db)).Name("spaces.get")
		spaces.Get("/fs", handlers.HandleGetFS(db))

		spaces.Get("/files", handlers.HandleGetFiles(db)).Name("spaces.files")

		spaces.Post("/upload", handlers.HandleUploadFile(db, svc, ws)).Name("spaces.upload")
		spaces.Get("/download", handlers.HandleDownloadDir(db, svc, ws)).Name("spaces.download")
	}

	// Manage files in a space
	files := auth.Group("/files/{file_id@uuid}", middlewares.FileMiddleware(db, "file_id"))
	{
		files.Get("/", handlers.HandleGetCodeFileContent(db, svc)).Name("files.get")
		files.Delete("/", handlers.HandleDeleteFile(db, svc, ws)).Name("files.delete")
		files.Put("/", handlers.HandleUpdateFileInfo(db)).Name("files.update")
	}
}

func registerValidators(r gale.RouterParamValidator) {
	r.RegisterRouteParamValidator("png", func(value string) (string, error) {
		strs := strings.SplitN(value, ".", 2)
		if strs[len(strs)-1] != "png" {
			return "", errors.New("invalid file extension")
		}

		return strs[0], nil
	})
}
