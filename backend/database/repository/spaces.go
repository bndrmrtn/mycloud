package repository

import (
	"github.com/bndrmrtn/my-cloud/database/models"
	"gorm.io/gorm"
)

func GetAllSpacesForUser(db *gorm.DB, userID string) ([]models.FileSpaceWithSize, error) {
	var spaces []models.FileSpaceWithSize
	result := db.Raw(`
		select
			file_spaces.id as id,
			file_spaces.created_at as created_at,
			file_spaces.updated_at as updated_at,
			file_spaces.name as name,
			sum(os_files.file_size) as size,
			space_user.user_id as space_user_id
		from file_spaces
			left join files on files.file_space_id = file_spaces.id
			left join os_files on os_files.id = files.os_file_id
			left join space_user on space_user.file_space_id = file_spaces.id
			where file_spaces.user_id = ? or space_user.user_id = ?
			group by file_spaces.id, space_user.user_id
	`, userID, userID).Find(&spaces)
	return spaces, result.Error
}

func FindSpaceByID(db *gorm.DB, id string) (models.FileSpace, error) {
	var space models.FileSpace
	result := db.Where("id = ?", id).First(&space)
	return space, result.Error
}

func GetSpaceFiles(db *gorm.DB, spaceID string, dir string) ([]models.File, error) {
	var files []models.File
	result := db.Model(&models.File{}).Where("file_space_id = ? and directory = ?", spaceID, dir).Preload("OSFile").Find(&files)
	return files, result.Error
}

func GetAllSpaceFiles(db *gorm.DB, spaceID string, dir string) ([]models.File, error) {
	var (
		files []models.File
	)

	dirPrefix := dir
	if dir != "/" {
		dirPrefix += "/"
	}

	result := db.Model(&models.File{}).Where("file_space_id = ? and (directory = ? or directory like ?)", spaceID, dir, dirPrefix+"%").Preload("OSFile").Find(&files)
	return files, result.Error
}

// GetSpaceFS returns the next directories in the given directory
//
// ChatGPT's help: https://chatgpt.com/share/673784f7-5224-8010-a870-025fb7fdd0db
func GetSpaceFS(db *gorm.DB, spaceID string, dir string) ([]string, error) {
	var directories []string
	result := db.Raw(`
		SELECT DISTINCT
			SUBSTRING_INDEX(
				SUBSTRING_INDEX(directory, '/',
					CASE
						WHEN ? = '/' THEN 2
						ELSE LENGTH(?) - LENGTH(REPLACE(?, '/', '')) + 2
					END
				),
				'/',
				-1
			) AS next_directory
		FROM files
		WHERE
			(directory LIKE CONCAT(TRIM(TRAILING '/' FROM ?), '/%') OR ? = '/')
			AND directory != ?
			AND file_space_id = ?
			AND LENGTH(directory) > LENGTH(TRIM(TRAILING '/' FROM ?))
			AND SUBSTRING_INDEX(directory, '/',
				CASE
					WHEN ? = '/' THEN 2
					ELSE LENGTH(?) - LENGTH(REPLACE(?, '/', '')) + 2
				END
			) != '';
	`, dir, dir, dir, dir, dir, dir, spaceID, dir, dir, dir, dir).Find(&directories)

	return directories, result.Error
}
