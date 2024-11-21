package dto

import "github.com/bndrmrtn/my-cloud/database/models"

type UserSession struct {
	User    *models.User    `json:"user"`
	Session *models.Session `json:"session"`
}
