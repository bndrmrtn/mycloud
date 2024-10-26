package repository

import (
	"net"

	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func NewSession(db *gorm.DB, userID string, ipAddr net.IP, agent string) (models.Session, error) {
	var session = models.Session{
		HasUser:   models.HasUserID(userID),
		IP:        ipAddr.String(),
		UserAgent: agent,
	}
	result := db.Create(&session)
	return session, result.Error
}

func FindUserBySessionID(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	result := db.Raw(`select users.* from users inner join sessions on sessions.user_id = users.id where sessions.id = ?`, id).First(&user)
	return user, result.Error
}
