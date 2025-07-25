package main

import "path/filepath"

func createFactory() error {
	dir := filepath.Join("internal/domain/db", db, "factory")
	if fileNotExists(dir) {
		if err := createDir(dir); err != nil {
			return err
		}
	}

	fileRepos := filepath.Join(dir, "repositories.go")
	if fileNotExists(fileRepos) {
		if err := save(fileRepos, "templates/factory.repositories.tmpl", nil); err != nil {
			return err
		}
	}

	fileServices := filepath.Join(dir, "services.go")
	if fileNotExists(fileServices) {
		if err := save(fileServices, "templates/factory.services.tmpl", nil); err != nil {
			return err
		}
	}
	return nil
}
