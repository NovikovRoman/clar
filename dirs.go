package main

import (
	"os"
)

// createDir creates a directory if it does not exist.
func createDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, permDir)
	}
	return
}
