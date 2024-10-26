package handlers

import (
	"archive/zip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func HandleDownloadDir(db *gorm.DB, svc services.StorageService, ws bolt.WSServer) bolt.HandlerFunc {
	var wsWriter = func(ws bolt.WSServer, userID string, data any) error {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}

		ws.BroadcastFunc(b, func(conn bolt.WSConn) bool {
			return conn.Ctx().Get(utils.WSUserID) == userID
		})

		return nil
	}

	return func(c bolt.Ctx) error {
		userID, err := ctxUserID(c)
		if err != nil {
			return err
		}

		space, err := ctxSpace(c)
		if err != nil {
			return err
		}

		// Prepare the download and notify the user
		go func(db *gorm.DB, ws bolt.WSServer, tmpFile, spaceID, userID, path, reqID string) (err error) {
			defer func() {
				if err != nil {
					logrus.Error("Failed to download directory", err)
					wsWriter(ws, userID, bolt.Map{
						"error": err.Error(),
					})
				}
			}()

			files, err := repository.GetAllSpaceFiles(db, spaceID, path)
			if err != nil {
				logrus.Fatal("Failed to get space files")
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
				return err
			}

			if err = z.Close(); err != nil {
				return err
			}

			download := models.Download{
				HasUser: models.HasUserID(userID),
				Path:    tmpFile,
				Expiry:  time.Now().Add(time.Hour * 24),
			}
			db.Create(&download)

			return wsWriter(ws, userID, bolt.Map{
				"download_request_finished": true,
				"request_id":                reqID,
				"download_id":               download.ID,
				"download_expiry":           download.Expiry,
			})
		}(db, ws, filepath.Join(svc.GetTmpDir(), uuid.New().String()+".zip"), space.ID, userID, queryPath(c), c.ID())

		return c.Status(http.StatusAccepted).JSON(bolt.Map{
			"accapted":   true,
			"request_id": c.ID(),
		})
	}
}
