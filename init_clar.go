package main

import (
	"path/filepath"
)

func initClar(modulePath, dbType string) (err error) {
	dirE := getPathLocation(dbType, dirEntity)
	if err = createDir(dirE); err != nil {
		return
	}
	dirR := getPathLocation(dbType, dirRepository)
	if err = createDir(dirR); err != nil {
		return
	}

	filename := filepath.Join(dirE, "base_entity.go")
	if err = saveTemplate(filename, getTemplate("interface"), nil); err != nil {
		return
	}

	data := struct {
		Backtick string
		Module   string
		DBType   string
	}{
		Backtick: backtick,
		Module:   modulePath,
		DBType:   dbType,
	}
	err = saveTemplate(filepath.Join(dirR, "utils.go"),
		getTemplateByDBType(dbType, "repository.utils"), data)
	return
}
