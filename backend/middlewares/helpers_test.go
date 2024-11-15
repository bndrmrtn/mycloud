package middlewares

import (
	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Session{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupGaleServer() (*gale.Gale, func(ctx gale.Ctx, err error) error) {
	sessionStore := gale.NewMemStorage()
	conf := config.Api(sessionStore)
	app := gale.New(&conf)

	return app, app.Config().ErrorHandler
}

func createTestUserSession(db *gorm.DB) (*models.Session, error) {
	user := &models.User{
		GID:   "1111111",
		Name:  "Jane",
		Email: "jane@mail.com",
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	sess := &models.Session{
		HasUser:   models.HasUserID(user.ID),
		IP:        "127.0.0.1",
		UserAgent: "Mozilla/5.0",
	}

	if err := db.Create(sess).Error; err != nil {
		return nil, err
	}

	return sess, nil
}
