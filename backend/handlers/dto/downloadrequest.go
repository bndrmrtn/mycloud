package dto

type DownloadRequestAccepted struct {
	Accapted  bool   `json:"accepted"`
	RequestID string `json:"request_id"`
}
