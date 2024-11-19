package models

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Base
	// GID is the google id of the user
	GID string `json:"gid"`
	// Name is the name of the user
	Name string `json:"name"`
	// Email is the email of the user
	Email string `json:"email"`
	// Image is the image url of the user
	Image string `json:"image_url"`

	// SQLite does not support enum, so we need to use string
	Role Role `json:"role" gorm:"default:'user'"`
}
