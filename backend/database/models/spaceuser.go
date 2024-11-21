package models

import (
	"gorm.io/gorm"
)

type SpaceUser struct {
	Base
	HasUser
	HasFileSpace
	PermissionInt int64                `json:"-"`
	Permission    *SpaceUserPermission `json:"permission,omitempty" gorm:"-:all"`
}

type SpaceUserPermission struct {
	ReadFile   bool `json:"read_file"`
	UpdateFile bool `json:"update_file"`
	DeleteFile bool `json:"delete_file"`
	UploadFile bool `json:"upload_file"`
}

const (
	ReadFileBit   = 1 << 0 // 0001
	UpdateFileBit = 1 << 1 // 0010
	DeleteFileBit = 1 << 2 // 0100
	UploadFileBit = 1 << 3 // 1000
)

func (s *SpaceUser) TableName() string {
	return "space_user"
}

func (s *SpaceUser) BeforeSave(tx *gorm.DB) error {
	s.EncodePermissions()
	return nil
}

func (s *SpaceUser) AfterFind(tx *gorm.DB) error {
	s.DecodePermissions()
	return nil
}

// Update `PermissionInt` based on `Permission` struct
func (s *SpaceUser) EncodePermissions() {
	if s.Permission == nil {
		return
	}

	s.PermissionInt = 0
	if s.Permission.ReadFile {
		s.PermissionInt |= ReadFileBit
	}
	if s.Permission.UpdateFile {
		s.PermissionInt |= UpdateFileBit
	}
	if s.Permission.DeleteFile {
		s.PermissionInt |= DeleteFileBit
	}
	if s.Permission.UploadFile {
		s.PermissionInt |= UploadFileBit
	}
}

// Decode `PermissionInt` into `Permission` struct
func (s *SpaceUser) DecodePermissions() {
	s.Permission = &SpaceUserPermission{}

	s.Permission.ReadFile = s.PermissionInt&ReadFileBit != 0
	s.Permission.UpdateFile = s.PermissionInt&UpdateFileBit != 0
	s.Permission.DeleteFile = s.PermissionInt&DeleteFileBit != 0
	s.Permission.UploadFile = s.PermissionInt&UploadFileBit != 0
}
