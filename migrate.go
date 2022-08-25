package main

import (
	"path/filepath"
)

func createMigrate(path string, dbType *DBType) (err error) {
	path = filepath.Join(path, dbType.name)
	pathSql := filepath.Join(path, "/migrations")

	if err = createDir(path); err != nil {
		return
	}
	if err = createDir(pathSql); err != nil {
		return
	}

	data := struct {
		Backtick string
	}{
		Backtick: backtick,
	}

	if err = saveTemplate(path+"/README.md", getTemplateByDBType(dbType, "migrate.readme"), data); err != nil {
		return
	}

	if err = saveTemplate(path+"/main.go", getTemplateByDBType(dbType, "migrate.main"), data); err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202205041600_begin.up.sql"), getTemplateByDBType(dbType, "migrate.up"), data)
	if err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202205041600_begin.down.sql"), getTemplateByDBType(dbType, "migrate.down"), data)
	return
}
