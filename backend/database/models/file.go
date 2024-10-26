package models

type File struct {
	Base
	// DATABASE DRIVEN

	// HasFileSpace is the file space that the file belongs to
	HasFileSpace
	// HasUser is the user who created the file
	HasUser

	Directory string `json:"directory"`
	FileName  string `json:"file_name"`

	// OS DRIVEN

	HasOSFile `json:"fileinfo"`
}
