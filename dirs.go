package main

import (
	"os"
	"path/filepath"
)

const (
	dirDB         = "db"
	dirEntity     = "entity"
	dirRepository = "repository"
	dirMigrate    = "migrations"
)

// createDir creates a directory if it does not exist.
func createDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
	}
	return
}

func getPathLocation(dbType, path string) string {
	return filepath.Join("internal", dirDB, dbType, path)
}
