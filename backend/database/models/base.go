package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	// ID is the primary key of the model
	ID string `json:"id" gorm:"primaryKey;unique;type:varchar(191)"`
	// CreatedAt is the time when the model is created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time when the model is updated
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) Exists() bool {
	return b.ID != ""
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}
