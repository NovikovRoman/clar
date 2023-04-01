package main

import (
	"path/filepath"
)

func initClar(dbType *DBType, internal bool) (err error) {
	dirE := getPathLocation(dirEntity, internal)
	if err = createDir(dirE); err != nil {
		return
	}
	dirR := getPathLocation(dirRepository, internal)
	if err = createDir(filepath.Join(dirR, dbType.name)); err != nil {
		return
	}

	filename := filepath.Join(dirE, "base_entity.go")
	if err = saveTemplate(filename, getTemplate("interface"), nil); err != nil {
		return
	}

	mPath := modulePath
	if internal {
		mPath = filepath.Join(mPath, "internal")
	}

	data := struct {
		Backtick string
		Module   string
	}{
		Backtick: backtick,
		Module:   mPath,
	}
	filename = filepath.Join(dirR, dbType.name, "utils.go")
	err = saveTemplate(filename, getTemplateByDBType(dbType, "repository.utils"), data)
	return
}
