package handlers

import (
	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm"
)

func HandleGetProfileImage(db *gorm.DB, svc services.StorageService, store gale.SessionStore) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		id := c.Param("id")

		if b, err := store.Get("image:" + id); err == nil {
			return c.SendFile(string(b))
		}

		var img models.ImageURL

		if err := db.Where("id = ?", id).Preload("OSFile").First(&img).Error; err != nil {
			return err
		}

		path, err := svc.GetRealPath(img.OSFile)
		if err != nil {
			return err
		}

		_ = store.Set("image:"+id, []byte(path))
		return c.SendFile(path)
	}
}
