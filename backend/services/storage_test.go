package services

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/stretchr/testify/assert"
)

func Test_NewStorageService(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a new storage service
	storageService, err := NewStorageServiceV1(tempDir, nil, 1024*utils.MB, 100)
	// Check if the storage service is created without error
	assert.Nil(t, err, "storage service should be created without error")
	assert.NotNil(t, storageService, "storage service should not be nil")

	// Check if the storage service has the correct values
	assert.Equal(t, filepath.Join(tempDir, "tmp"), storageService.GetTmpDir(), "storage service should have a tmp directory")

	// Check if the storage service has 0 containers
	assert.Equal(t, 0, storageService.Containers(), "storage service should have 0 containers at the beginning")
}

func Test_StorageFileOperations(t *testing.T) {
	// Create a temporary database and directory for testing
	mockDB, err := setupTestDB()
	if err != nil {
		t.Fatal("error setting up test db", err)
	}
	tempDir := t.TempDir()

	// Create a new storage service
	storageService, err := NewStorageServiceV1(tempDir, mockDB, 1024*utils.MB, 100)
	// Check if the storage service is created without error
	assert.Nil(t, err, "storage service should be created without error")
	assert.NotNil(t, storageService, "storage service should not be nil")

	// Create a new file with text content
	textContent := "Hello world!"
	osFile, err := storageService.Store(strings.NewReader(textContent), int64(len(textContent)), ".txt")

	// Check if the file is stored without error
	assert.Nil(t, err, "file should be stored without error")
	assert.NotNil(t, osFile, "file should not be nil")
	// Check if the storage created a container for the file
	assert.Equal(t, 1, storageService.Containers(), "storage service should have 1 container")

	// Read the file content
	content, err := storageService.ReadFile(osFile)
	// Check if the file content is read without error
	assert.Nil(t, err, "file content should be read without error")
	assert.Equal(t, textContent, string(content), "file content should be the same as the text content")

	// Delete the file
	err = storageService.Delete(osFile)
	assert.Nil(t, err, "file should be deleted without error")

	content, err = storageService.ReadFile(osFile)
	// Check if the file content is read without error
	assert.NotNil(t, err, "file content should not be read after deletion")
	assert.Nil(t, content, "file content should be nil after deletion")
}

func Test_GetFileRealPath(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a new storage service
	storageService, err := NewStorageServiceV1(tempDir, nil, 1024*utils.MB, 100)
	// Check if the storage service is created without error
	assert.Nil(t, err, "storage service should be created without error")
	assert.NotNil(t, storageService, "storage service should not be nil")

	var (
		fileID    = "test-file-id"
		container = "container0"
	)

	osFile := &models.OSFile{
		Base: models.Base{
			ID: fileID,
		},
		Container: container,
	}

	realPath := storageService.GetRealPath(osFile)
	assert.Equal(t, filepath.Join(tempDir, container, fileID), realPath, "real path should be the same as the expected path")
}
