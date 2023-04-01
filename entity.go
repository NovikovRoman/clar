package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type entity struct {
	structName string
	name       string
	snakeName  string
	table      string
	nameRunes  []rune
	symb       string
}

func newEntity(name string) (ent *entity) {
	ent = &entity{
		nameRunes: []rune(name),
		snakeName: toSnake(name),
	}
	ent.symb = strings.ToLower(string(ent.nameRunes[0]))
	ent.name = strings.ToLower(string(ent.nameRunes[0])) + string(ent.nameRunes[1:])
	ent.structName = cases.Title(language.English, cases.NoLower).String(ent.name)

	ent.table = ent.snakeName
	switch []rune(ent.table)[len(ent.table)-1] {
	case 's':
		ent.table += "es"

	case 'y':
		ent.table = string([]rune(ent.table)[:len(ent.table)-1]) + "ies"

	default:
		ent.table += "s"
	}
	return
}

func createEntity(name string, dbType *DBType, empty, simple, internal bool) (err error) {
	ent := newEntity(name)
	if err = initBasicEntityFiles(ent, dbType, empty, simple, internal); err != nil {
		return
	}

	switch dbType.code {
	case MysqlCode:
		return createMysqlEntityFiles(ent, dbType, empty, internal)
	}

	err = errors.New("Database type not supported.")
	return
}

func initBasicEntityFiles(ent *entity, dbType *DBType, empty, simple, internal bool) (err error) {
	mPath := modulePath
	if internal {
		mPath = filepath.Join(modulePath, "internal")
	}
	data := struct {
		Module     string
		Backtick   string
		Entity     string
		EntitySymb string
		EntityName string
	}{
		Module:     mPath,
		Backtick:   backtick,
		Entity:     cases.Title(language.English, cases.NoLower).String(ent.name),
		EntitySymb: ent.symb,
		EntityName: ent.name,
	}

	tmplEntity := "entity.normal"
	tmplRepository := "repository.interface"
	if empty {
		tmplEntity = "entity.empty"
		tmplRepository = "repository.interface.empty"

	} else if simple {
		tmplEntity = "entity.simple"
	}

	entitityFilename := ent.snakeName + ".go"

	dirE := getPathLocation(dirEntity, internal)
	if err = createDir(dirE); err != nil {
		return
	}
	filePath := filepath.Join(dirE, entitityFilename)
	if err = saveTemplate(filePath, getTemplateByDBType(dbType, tmplEntity), data); err != nil {
		return
	}

	dirR := getPathLocation(dirRepository, internal)
	if err = createDir(dirR); err != nil {
		return
	}
	filePath = filepath.Join(dirR, entitityFilename)
	err = saveTemplate(filePath, getTemplate(tmplRepository), data)
	return
}

func createMysqlEntityFiles(ent *entity, dbType *DBType, empty, internal bool) (err error) {
	dirDbType := filepath.Join(getPathLocation(dirRepository, internal), dbType.name)
	if _, err = os.Stat(dirDbType); os.IsNotExist(err) {
		if err = createDir(dirDbType); err != nil {
			return
		}

		if err = initClar(dbType, internal); err != nil {
			return
		}
	}

	mPath := modulePath
	if internal {
		mPath = filepath.Join(modulePath, "internal")
	}
	data := struct {
		Module      string
		Backtick    string
		Table       string
		Entity      string
		EntityName  string
		EntityTable string
	}{
		Module:      mPath,
		Backtick:    backtick,
		Table:       ent.table,
		Entity:      ent.structName,
		EntityName:  ent.name,
		EntityTable: ent.table,
	}

	tmpl := "repository.entity"
	if empty {
		tmpl = "repository.empty.entity"
	}

	filePath := filepath.Join(dirDbType, ent.snakeName+".go")
	err = saveTemplate(filePath, getTemplateByDBType(dbType, tmpl), data)
	return
}
