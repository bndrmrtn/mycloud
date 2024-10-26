package handlers

import (
	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/services"
	"gorm.io/gorm"
)

func HandleGetProfileImage(db *gorm.DB, svc services.StorageService, store bolt.SessionStore) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		id := c.Param("id")

		if b, err := store.Get("image:" + id); err == nil {
			return c.SendFile(string(b))
		}

		var img models.ImageURL

		if err := db.Where("id = ?", id).Preload("OSFile").First(&img).Error; err != nil {
			return err
		}

		path := svc.GetRealPath(img.OSFile)
		store.Set("image:"+id, []byte(path))
		return c.SendFile(path)
	}
}
