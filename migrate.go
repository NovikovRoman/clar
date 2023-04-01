package main

import (
	"path/filepath"
)

func createMigrate(dbType *DBType, internal bool) (err error) {
	path := ""
	if internal {
		path = dirInternal
	}

	path = filepath.Join(path, dirMigrate, dbType.name)
	pathSql := filepath.Join(path, "migrations")
	if err = createDir(pathSql); err != nil {
		return
	}

	data := struct {
		Backtick string
	}{
		Backtick: backtick,
	}

	if err = saveTemplate(path+"/migrate.go", getTemplateByDBType(dbType, "migrate"), data); err != nil {
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
