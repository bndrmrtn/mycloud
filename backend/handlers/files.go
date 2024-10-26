package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm"
)

func HandleGetFiles(db *gorm.DB) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		path := queryPath(c)
		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		files, err := repository.GetSpaceFiles(db, space.ID, path)
		if err != nil {
			return err
		}

		return c.JSON(files)
	}
}

func HandleGetFS(db *gorm.DB) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		path := c.URL().Query().Get("path")
		if path == "" || path == "." {
			path = "/"
		}

		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		files, err := repository.GetSpaceFS(db, space.ID, path)
		if err != nil {
			return err
		}

		return c.JSON(files)
	}
}

func HandleUploadFile(db *gorm.DB, svc services.StorageService) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		userID, err := ctxUserID(c)
		if err != nil {
			return err
		}

		_, header, err := c.Body().File("file")
		if err != nil {
			return err
		}

		name := c.Request().FormValue("filename")
		if len(name) < 1 {
			name = header.Filename
		} else if len(name) > 50 {
			return bolt.NewError(http.StatusBadRequest, "File name is too long")
		}

		dir := c.Request().FormValue("directory")
		dir = filepath.Clean(dir)
		if dir == "." {
			dir = "/"
		}

		osFile, err := svc.StoreMultipartFile(header)
		if err != nil {
			return err
		}

		file := models.File{
			HasFileSpace: models.HasFileSpaceID(space.ID),
			HasUser:      models.HasUserID(userID),
			HasOSFile:    models.HasOSFileID(osFile.ID),
			FileName:     name,
			Directory:    dir,
		}

		if err := db.Create(&file).Error; err != nil {
			return err
		}

		return c.JSON(file)
	}
}

func queryPath(c bolt.Ctx) string {
	path := c.URL().Query().Get("path")
	if path == "" || path == "." {
		path = "/"
	}

	return path
}
