package config

import (
	"errors"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// AppConfig is the struct for the config file
type AppConfig struct {
	Service     ServiceConfig     `json:"service" yaml:"service"`
	Application ApplicationConfig `json:"application" yaml:"application"`
}

// ServiceConfig is the configuration for StorageService
type ServiceConfig struct {
	Version    string `json:"version" yaml:"version"`
	AppdataDir string `json:"-" yaml:"appdata_dir"`
}

// ApplicationConfig is the main application's configuration
type ApplicationConfig struct {
	Authorization AuthorizationConfig `json:"authorization" yaml:"authorization"`
}

// AuthorizationConfig configures the authorization of the application
type AuthorizationConfig struct {
	UseWhitelist bool       `json:"use_whitelist" yaml:"use_whitelist"`
	UseBlacklist bool       `json:"use_blacklist" yaml:"use_blacklist"`
	Admin        AdminCofig `json:"admin" yaml:"admin"`
}

// AdminCofig configures the admin(s) of the application
type AdminCofig struct {
	PrimaryAdminEmail string `json:"-" yaml:"primary_admin_email"`
	EnableMultiAdmin  bool   `json:"enable_multi_admin" yaml:"enable_multi_admin"`
}

// ReadAppConfig reads the config file and returns the AppConfig struct
func ReadAppConfig() (*AppConfig, error) {
	// Get and check config file
	var (
		conf     AppConfig
		filename = "config"
		exts     = []string{".yaml", ".yml"}
	)

	var (
		data []byte
		err  error
	)
	for _, ext := range exts {
		data, err = os.ReadFile(filename + ext)
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	// Check config file content

	if conf.Service.Version == "" || conf.Service.Version != "1" {
		return nil, errors.New("service.version should be 1")
	}

	if conf.Service.AppdataDir == "" {
		conf.Service.AppdataDir = "mycloud-appdata"
	} else {
		conf.Service.AppdataDir = filepath.Clean(conf.Service.AppdataDir)
	}

	if err := os.MkdirAll(conf.Service.AppdataDir, 0755); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Join(conf.Service.AppdataDir, "tmp"), 0755); err != nil {
		return nil, err
	}

	adminEmail, err := mail.ParseAddress(conf.Application.Authorization.Admin.PrimaryAdminEmail)
	if err != nil || adminEmail.Address != conf.Application.Authorization.Admin.PrimaryAdminEmail {
		return nil, errors.New("app.admin.primary_admin_email is not a valid email address")
	}

	if conf.Application.Authorization.UseWhitelist && conf.Application.Authorization.UseBlacklist {
		return nil, errors.New("app.authorization.use_whitelist and app.authorization.use_blacklist cannot be both true")
	}

	if !conf.Application.Authorization.UseWhitelist && !conf.Application.Authorization.UseBlacklist {
		logrus.Warn("Both app.authorization.use_whitelist and app.authorization.use_blacklist are false, every email will be allowed")
	}

	return &conf, nil
}
