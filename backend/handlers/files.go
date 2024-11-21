package handlers

import (
	"fmt"
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/handlers/dao"
	"github.com/bndrmrtn/my-cloud/handlers/dto"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/sirupsen/logrus"
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

		exists, err := repository.IsFileExists(db, space.ID, "", dir, name)
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

		go func() {
			file, err = repository.FindFileByID(db, file.ID)
			if err != nil {
				logrus.Warn("failed to find file", err)
				return
			}

			_ = wsWriter(ws, userID, dto.WSEvent{
				Event: utils.WSFileUploadEvent,
				Data: dto.WSEventFileUploaded{
					File: &file,
				},
			})
		}()

		return c.JSON(file)
	}
}

// HandleGetFile returns the file content
func HandleGetFile(db *gorm.DB, svc services.StorageService) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		if file.OSFile.FileSize > 5*utils.MB {
			return gale.NewError(http.StatusBadRequest, "file size exceeded the 5mb limit")
		}

		path, err := svc.GetRealPath(file.OSFile)
		if err != nil {
			return err
		}

		return c.ContentType(file.OSFile.Type).SendFile(path)
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
			err = svc.Delete(file.OSFile)
			if err != nil {
				logrus.Warn("failed to delete file", err)
			}
		}

		if err := db.Delete(&file).Error; err != nil {
			return err
		}

		_ = wsWriter(ws, userID, dto.WSEvent{
			Event: utils.WSFileDeleteEvent,
			Data: dto.WSEventFileDeleted{
				FileID: file.ID,
			},
		})

		return c.JSON(dto.Message{
			Message: "File deleted successfully",
		})
	}
}

func HandleUpdateFileInfo(db *gorm.DB, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		var data dao.FileUploadInformation

		if err := c.Body().ParseJSON(&data); err != nil {
			return err
		}

		if err := validateFileName(data.Name); err != nil {
			return err
		}

		data.Directory = validateDir(data.Directory)

		ok, err := repository.IsFileExists(db, file.FileSpaceID, file.ID, data.Directory, data.Name)
		if err != nil {
			return err
		}

		if ok {
			return gale.NewError(http.StatusFound, "File already exists in this directory")
		}

		file.FileName = data.Name
		file.Directory = data.Directory

		if err := db.Save(&file).Error; err != nil {
			return gale.NewError(http.StatusInternalServerError, "Failed to update file info")
		}

		go func() {
			_ = wsWriter(ws, file.UserID, dto.WSEvent{
				Event: utils.WSFileUpdateEvent,
				Data: dto.WSEventFileUploaded{
					File: file,
				},
			})
		}()

		return c.JSON(file)
	}
}

func HandleDownloadFile(db *gorm.DB, svc services.StorageService) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		file, err := ctxSpaceFile(c)
		if err != nil {
			return err
		}

		path, err := svc.GetRealPath(file.OSFile)
		if err != nil {
			return err
		}

		c.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.FileName))
		c.Header().Add("Content-Type", "application/octet-stream")

		return c.ContentType(file.OSFile.Type).SendFile(path)
	}
}
