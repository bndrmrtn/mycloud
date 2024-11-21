package dto

import (
	"time"

	"github.com/bndrmrtn/my-cloud/database/models"
)

type WSEvent struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

type WSEventDownloadPreparedData struct {
	RequestID      string    `json:"request_id"`
	DownloadID     string    `json:"download_id"`
	DownloadExpiry time.Time `json:"download_expiry"`
}

type WSEventFileUploaded struct {
	File *models.File `json:"file"`
}

type WSEventFileDeleted struct {
	FileID string `json:"file_id"`
}
