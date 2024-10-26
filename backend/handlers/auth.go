package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bndrmrtn/go-bolt"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func HandleCreateAuthURL(c bolt.Ctx) error {
	state := utils.NewRandom().String(10)
	c.Session().Set("auth_state", []byte(state))

	conf := config.GoogleOAuth()
	url := conf.AuthCodeURL(state)

	return c.JSON(bolt.Map{
		"redirect_url": url,
	})
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func HandleAuthUser(db *gorm.DB, svc services.StorageService) bolt.HandlerFunc {
	return func(c bolt.Ctx) error {
		state, err := c.Session().Get("auth_state")

		if err != nil {
			return err
		}

		if string(state) != c.URL().Query().Get("state") {
			return bolt.NewError(http.StatusNotAcceptable, "Invalid state parameter")
		}

		_ = c.Session().Delete("auth_state")

		conf := config.GoogleOAuth()

		token, err := conf.Exchange(context.Background(), c.URL().Query().Get("code"))
		if err != nil {
			return bolt.NewError(http.StatusNotAcceptable, "Failed to exchange token")
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return bolt.NewError(http.StatusInternalServerError, "User Data Fetch Failed")
		}

		userDataB, err := io.ReadAll(resp.Body)
		if err != nil {
			return bolt.NewError(http.StatusInternalServerError, "JSON Reading Failed")
		}

		var userData GoogleUser
		err = json.Unmarshal(userDataB, &userData)
		if err != nil {
			return bolt.NewError(http.StatusInternalServerError, "JSON Unmarshal Failed")
		}

		if !userData.VerifiedEmail {
			return bolt.NewError(http.StatusNotAcceptable, "Email is not verified")
		}

		user, err := repository.FindUserByEmail(db, userData.Email)
		if err != nil {
			return bolt.NewError(http.StatusNotFound, "User not found")
		}

		session, err := repository.NewSession(db, user.ID, c.IP(), c.Request().UserAgent())
		if err != nil {
			return bolt.NewError(http.StatusInternalServerError, "Session creation failed")
		}

		if err = c.Session().Set(utils.AuthSessionKey, []byte(session.ID)); err != nil {
			return err
		}

		go func() {
			err = updateUserImage(db, svc, userData.Picture, &user)
			if err != nil {
				log.Warn("Failed to update user image")
			}
		}()

		return c.JSON(bolt.Map{
			"user":    user,
			"session": session,
		})
	}
}

func HandleGetAuthUser(c bolt.Ctx) error {
	user, err := ctxUser(c)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func updateUserImage(db *gorm.DB, svc services.StorageService, gImageURL string, user *models.User) error {
	if user.Image != "" {
		return nil
	}

	res, err := http.Get(gImageURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := svc.Store(res.Body, res.ContentLength, ".png")
	if err != nil {
		return err
	}

	image := models.ImageURL{
		HasOSFile: models.HasOSFileID(file.ID),
	}

	if err := db.Create(&image).Error; err != nil {
		return err
	}

	user.Image = image.ID + ".png"

	return db.Save(&user).Error
}
