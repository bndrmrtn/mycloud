package models

type Session struct {
	Base
	HasUser
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
