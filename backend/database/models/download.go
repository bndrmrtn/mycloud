package models

import "time"

type Download struct {
	Base
	HasUser
	Path   string    `json:"path"`
	Expiry time.Time `json:"expiry"`
}
