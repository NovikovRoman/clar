package main

import (
	"os"
	"path/filepath"
)

const (
	dirInternal   = "internal"
	dirEntity     = "domain/entity"
	dirRepository = "domain/repository"
	dirMigrate    = "domain/migrate"
)

// createDir creates a directory if it does not exist.
func createDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
	}
	return
}

func getPathLocation(path string, internal bool) string {
	if internal {
		return filepath.Join(dirInternal, path)
	}
	return path
}
