package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/config"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/database/repository"
	"github.com/bndrmrtn/my-cloud/handlers/dto"
	"github.com/bndrmrtn/my-cloud/services"
	"github.com/bndrmrtn/my-cloud/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func HandleCreateAuthURL(c gale.Ctx) error {
	state := utils.NewRandom().String(10)
	if err := c.Session().Set("auth_state", []byte(state)); err != nil {
		return err
	}

	conf := config.GoogleOAuth()
	url := conf.AuthCodeURL(state)

	return c.JSON(gale.Map{
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

func HandleAuthUser(db *gorm.DB, svc services.StorageService, authConf *config.AuthorizationConfig) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		userData, err := getGoogleUser(c)
		if err != nil {
			return err
		}

		user, err := checkUser(db, userData, authConf)
		if err != nil {
			return gale.NewError(http.StatusBadRequest, err.Error())
		}

		session, err := repository.NewSession(db, user.ID, c.IP(), c.Request().UserAgent())
		if err != nil {
			return gale.NewError(http.StatusInternalServerError, "Session creation failed")
		}

		if err = c.Session().Set(utils.AuthSessionKey, []byte(session.ID)); err != nil {
			return err
		}

		go func() {
			err = updateUserImage(db, svc, userData.Picture, user)
			if err != nil {
				log.Warn("Failed to update user image:", err)
			}
		}()

		return c.JSON(dto.UserSession{
			User:    user,
			Session: &session,
		})
	}
}

func HandleGetAuthUser(c gale.Ctx) error {
	user, err := ctxUser(c)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func HandleLogout(c gale.Ctx) error {
	return c.Session().Destroy()
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

func getGoogleUser(c gale.Ctx) (*GoogleUser, error) {
	state, err := c.Session().Get("auth_state")

	if err != nil {
		return nil, err
	}

	if string(state) != c.URL().Query().Get("state") {
		return nil, gale.NewError(http.StatusNotAcceptable, "Invalid state parameter")
	}

	_ = c.Session().Delete("auth_state")

	conf := config.GoogleOAuth()

	token, err := conf.Exchange(context.Background(), c.URL().Query().Get("code"))
	if err != nil {
		return nil, gale.NewError(http.StatusNotAcceptable, "Failed to exchange token")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, gale.NewError(http.StatusInternalServerError, "User Data Fetch Failed")
	}

	userDataB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, gale.NewError(http.StatusInternalServerError, "JSON Reading Failed")
	}

	var userData GoogleUser
	err = json.Unmarshal(userDataB, &userData)
	if err != nil {
		return nil, gale.NewError(http.StatusInternalServerError, "JSON Unmarshal Failed")
	}

	if !userData.VerifiedEmail {
		return nil, gale.NewError(http.StatusNotAcceptable, "Email is not verified")
	}

	return &userData, nil
}

// checkUser checks if the user is in the whitelist or blacklist and returns the user if the user is authorized to access the application
func checkUser(db *gorm.DB, guser *GoogleUser, conf *config.AuthorizationConfig) (*models.User, error) {
	user, err := repository.FindUserByEmail(db, guser.Email)

	// If the user found, check if the user can access the application
	if err == nil {
		// Return if user is admin
		if conf.Admin.PrimaryAdminEmail == guser.Email {
			return &user, nil
		}

		// Return if user is in whitelist
		if conf.UseWhitelist {
			ok, err := repository.CheckEmailInWhitelist(db, guser.Email)
			if err != nil {
				return nil, err
			}

			if ok {
				return &user, nil
			}

			return nil, errors.New("User is not in whitelist")
		}

		// Return if user is in blacklist
		if conf.UseBlacklist {
			ok, err := repository.CheckEmailPassBlacklist(db, guser.Email)
			if err != nil {
				return nil, err
			}

			if ok {
				return &user, nil
			}

			return nil, errors.New("User is in blacklist")
		}

		// Return if the user is not in the blacklist
		return &user, nil
	}

	// Otherwise, create a new user if it can access the application

	// Return if the user is the primary admin
	if conf.Admin.PrimaryAdminEmail == guser.Email {
		user, err := createUser(db, guser)
		if err == nil {
			user.Role = models.RoleAdmin
		}
		return user, err
	}

	// Return if the user is in whitelist
	if conf.UseWhitelist {
		ok, err := repository.CheckEmailInWhitelist(db, guser.Email)
		if err != nil {
			return nil, err
		}

		if ok {
			return createUser(db, guser)
		}

		return nil, errors.New("User is not in whitelist")
	}

	// Return if the user isn't in blacklist
	if conf.UseBlacklist {
		ok, err := repository.CheckEmailPassBlacklist(db, guser.Email)
		if err != nil {
			return nil, err
		}

		if ok {
			return createUser(db, guser)
		}

		return nil, errors.New("User is in blacklist")
	}

	// Otherwise, create a new user
	return createUser(db, guser)
}

// createUser creates a new user in the database from the Google user data
func createUser(db *gorm.DB, guser *GoogleUser) (*models.User, error) {
	var user models.User

	user.GID = guser.ID
	user.Name = guser.Name
	user.Email = guser.Email

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
