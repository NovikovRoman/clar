package main

import (
	"path/filepath"
)

const (
	dirEntity     = "internal/domain/entity"
	dirRepository = "internal/domain/repository"
)

func initClar(module string, dbType *DBType) (err error) {
	if err = createDir(dirEntity); err != nil {
		return
	}
	if err = createDir(filepath.Join(dirRepository, dbType.name)); err != nil {
		return
	}

	filename := filepath.Join(dirEntity, "entity_interface.go")
	if err = saveTemplate(filename, getTemplate("interface"), nil); err != nil {
		return
	}

	data := struct {
		Backtick string
		Module   string
	}{
		Backtick: backtick,
		Module:   module,
	}
	filename = filepath.Join(dirRepository, dbType.name, "utils.go")
	err = saveTemplate(filename, getTemplateByDBType(dbType, "repository.utils"), data)
	return
}
