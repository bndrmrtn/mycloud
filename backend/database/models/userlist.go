package models

// UserWhitelist allows users to authenticate with the system
type UserWhitelist struct {
	Base
	Email string `json:"email" gorm:"unique;not null"`
}

// UserBlacklist prevents users from authenticating with the system
type UserBlacklist struct {
	Base
	Email string `json:"email" gorm:"unique;not null"`
}
