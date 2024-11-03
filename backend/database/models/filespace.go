package models

type FileSpace struct {
	Base
	// HasUser is the user who created the file space
	HasUser
	// Name is the name of the file space
	Name string `json:"name" gorm:"not null"`
	// Size is the total size of the file space
	Size int64 `json:"size" gorm:"-:migration"`
}
