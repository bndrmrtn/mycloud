package models

type OSFile struct {
	Base

	Container   string `json:"-"`
	Path        string `json:"-"`
	Extension   string `json:"extension"`
	Type        string `json:"type"`
	FileSize    int64  `json:"size"`
	ContentHash string `json:"hash"`
}
