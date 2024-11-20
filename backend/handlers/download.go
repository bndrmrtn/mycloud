package handlers

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func HandleDownloadDir(db *gorm.DB, svc services.StorageService, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		userID, err := ctxUserID(c)
		if err != nil {
			return err
		}

		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		// Prepare the download and notify the user
		go func(db *gorm.DB, ws gale.WSServer, tmpFile, spaceID, userID, path, reqID string) {
			var err error

			defer func() {
				if err != nil {
					logrus.Error("Failed to download directory", err)
					_ = wsWriter(ws, userID, gale.Map{
						"error": err.Error(),
					})
				}
			}()

			files, err := repository.GetAllSpaceFiles(db, spaceID, path)
			if err != nil {
				logrus.Error("Failed to get space files")
				return
			}

			z, err := os.Create(tmpFile)
			if err != nil {
				return
			}

			zw := zip.NewWriter(z)

			for _, file := range files {
				var f []byte
				f, err = svc.ReadFile(file.OSFile)
				if err != nil {
					logrus.Fatal("Failed to read file", err)
					return
				}

				var fw io.Writer
				fw, err = zw.Create(filepath.Join(file.Directory, file.FileName))
				if err != nil {
					logrus.Fatal("Failed to create zip file", err)
					return
				}

				_, err = fw.Write(f)
				if err != nil {
					logrus.Fatal("Failed to write to zip file", err)
					return
				}
			}

			if err = zw.Close(); err != nil {
				return
			}

			if err = z.Close(); err != nil {
				return
			}

			download := models.Download{
				HasUser: models.HasUserID(userID),
				Path:    tmpFile,
				Expiry:  time.Now().Add(time.Hour * 24),
			}

			if err = db.Create(&download).Error; err != nil {
				return
			}

			err = wsWriter(ws, userID, gale.Map{
				"type":            utils.WSDownloadRequestEvent,
				"request_id":      reqID,
				"download_id":     download.ID,
				"download_expiry": download.Expiry,
			})
		}(db, ws, filepath.Join(svc.GetTmpDir(), uuid.New().String()+".zip"), space.ID, userID, queryPath(c), c.ID())

		return c.Status(http.StatusAccepted).JSON(gale.Map{
			"accapted":   true,
			"request_id": c.ID(),
		})
	}
}
