package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/bndrmrtn/go-bolt"
)

func validateFileName(name string) error {
	if len(name) > 50 {
		return bolt.NewError(http.StatusBadRequest, "File name is too long")
	}

	if strings.Contains(name, "/") {
		return bolt.NewError(http.StatusBadRequest, "File name cannot contain slashes")
	}

	return nil
}

func validateDir(s string) string {
	dir := filepath.Clean(s)
	if dir == "." {
		dir = "/"
	}

	if !strings.HasPrefix(dir, "/") {
		dir = "/" + dir
	}

	if dir != "/" {
		dir = strings.TrimSuffix(dir, "/")
	}

	return dir
}