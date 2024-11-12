package handlers

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	"gorm.io/gorm"
)

func HandleGetFiles(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
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

func HandleGetFS(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
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

func HandleUploadFile(db *gorm.DB, svc services.StorageService, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
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
		}

		if err := validateFileName(name); err != nil {
			return err
		}

		dir := validateDir(c.Request().FormValue("directory"))

		osFile, err := svc.StoreMultipartFile(header)
		if err != nil {
			return err
		}

		exists, err := repository.IsFileExists(db, space.ID, dir, name)
		if err != nil {
			return err
		}

		if exists {
			return gale.NewError(http.StatusFound, "File already exists in this directory")
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

		wsWriter(ws, userID, gale.Map{
			"type":          "space_file_upload_succeed",
			"file_space_id": space.ID,
		})

		return c.JSON(file)
	}
}

func HandleGetCodeFileContent(db *gorm.DB, svc services.StorageService) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		if file.OSFile.FileSize > 5*utils.MB {
			return gale.NewError(http.StatusBadRequest, "File is too big")
		}

		content, err := svc.ReadFile(file.OSFile)
		if err != nil {
			return err
		}

		return c.ContentType(gale.ContentTypeText).Send(content)
	}
}

func HandleDeleteFile(db *gorm.DB, svc services.StorageService, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		userID, err := ctxUserID(c)
		if err != nil {
			return err
		}

		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		can, err := repository.CanDeleteOSFile(db, file.OSFileID)
		if err != nil {
			return err
		}

		if can {
			svc.Delete(file.OSFile)
		}

		if err := db.Delete(&file).Error; err != nil {
			return err
		}

		wsWriter(ws, userID, gale.Map{
			"type":          "space_file_delete_succeed",
			"file_space_id": file.FileSpaceID,
		})

		return c.JSON(gale.Map{"message": "File deleted"})
	}
}

func HandleUpdateFileInfo(db *gorm.DB) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		var data struct {
			Name      string `json:"name"`
			Directory string `json:"directory"`
		}

		if err := c.Body().ParseJSON(&data); err != nil {
			return err
		}

		if err := validateFileName(data.Name); err != nil {
			return err
		}

		data.Directory = validateDir(data.Directory)

		ok, err := repository.IsFileExists(db, file.FileSpaceID, data.Directory, data.Name)
		if err != nil {
			return err
		}

		if !ok {
			return gale.NewError(http.StatusFound, "File already exists in this directory")
		}

		file.FileName = data.Name
		file.Directory = data.Directory

		if err := db.Save(&file).Error; err != nil {
			return gale.NewError(http.StatusInternalServerError, "Failed to update file info")
		}

		return c.JSON(file)
	}
}
