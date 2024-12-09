package repository

import "gorm.io/gorm"

type ContainerWithSize struct {
	Container string `json:"container"`
	Size      int    `json:"size"`
}

func GetOSFilesSeparatedByContainers(db *gorm.DB) ([]*ContainerWithSize, error) {
	var files []*ContainerWithSize

	result := db.Raw(`select container, file_size as size from os_files group by container`).Scan(&files)
	return files, result.Error
}

func GetSizeDiff(db *gorm.DB) (map[string]int64, error) {
	var files struct {
		TotalFileSize  int64 `json:"total_file_size"`
		UniqueFileSize int64 `json:"unique_file_size"`
	}

	err := db.Raw(`
		select sum(os_files.file_size) AS total_file_size, sum(DISTINCT os_files.file_size) AS unique_file_size
		from files join os_files ON os_files.id = files.os_file_id
	`).Scan(&files).Error

	if err != nil {
		return nil, err
	}

	return map[string]int64{
		"total_file_size":  files.TotalFileSize,
		"unique_file_size": files.UniqueFileSize,
	}, nil
}
