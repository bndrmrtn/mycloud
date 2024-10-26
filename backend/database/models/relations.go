package models

// HasUser is a relation to the User model
type HasUser struct {
	UserID string `json:"-" gorm:"type:varchar(191)"`
	User   *User  `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
}

// HasUserID is a helper function to create a HasRole struct
func HasUserID(id string) HasUser {
	return HasUser{UserID: id}
}

// HasFileSpace is a relation to the FileSpace model
type HasFileSpace struct {
	FileSpaceID string     `json:"-" gorm:"type:varchar(191)"`
	FileSpace   *FileSpace `json:"filespace,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:FileSpaceID"`
}

// HasFileSpaceID is a helper function to create a HasFileSpace struct
func HasFileSpaceID(id string) HasFileSpace {
	return HasFileSpace{FileSpaceID: id}
}

// HasFile is a relation to the File model
type HasFile struct {
	FileID string `json:"-" gorm:"type:varchar(191)"`
	File   *File  `json:"file,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:FileID"`
}

// HasFileID is a helper function to create a HasFile struct
func HasFileID(id string) HasFile {
	return HasFile{FileID: id}
}

// HasOSFile is a relation to the OSFile model
type HasOSFile struct {
	OSFileID string  `json:"-" gorm:"type:varchar(191)"`
	OSFile   *OSFile `json:"info" gorm:"constraint:OnDelete:CASCADE;foreignKey:OSFileID"`
}

// HasFileID is a helper function to create a HasRole struct
func HasOSFileID(id string) HasOSFile {
	return HasOSFile{OSFileID: id}
}
