package main

import (
	"path/filepath"
)

const dirMigrate = "internal/domain/migrate"

func createMigrate(dbType *DBType) (err error) {
	path := filepath.Join(dirMigrate, dbType.name)
	/* if err = createDir(path); err != nil {
		return
	} */
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
