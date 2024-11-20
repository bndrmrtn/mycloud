package utils

const (
	MB = 1 << 20

	AuthSessionKey = "auth.session_id"

	MyCloudSpaceName     = "mycloud"
	MyCloudImageSpaceDir = "images"

	RequestAuthUserKey  = "auth.user_pointer"
	RequestSpaceKey     = "spaces.space_pointer"
	RequestSpaceFileKey = "spaces.space_file_pointer"

	WSUserID = "ws.user_id"

	WSFileUploadEvent      = "space_file_upload"
	WSFileDeleteEvent      = "space_file_delete"
	WSFileUpdateEvent      = "space_file_update"
	WSDownloadRequestEvent = "download_request"
)
