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

func createEntity(module string, name string, dbType *DBType, simple bool) (err error) {
	if err = createDir(dirEntity); err != nil {
		return
	}
	if err = createDir(dirRepository); err != nil {
		return
	}

	ent := newEntity(name)
	if err = initBasicEntityFiles(ent, dbType, simple); err != nil {
		return
	}

	switch dbType.code {
	case MysqlCode:
		return createMysqlEntityFiles(module, ent, dbType)
	}

	err = errors.New("Database type not supported.")
	return
}

func initBasicEntityFiles(ent *entity, dbType *DBType, simple bool) (err error) {
	data := struct {
		Module     string
		Backtick   string
		Entity     string
		EntitySymb string
		EntityName string
	}{
		Module:     modulePath,
		Backtick:   backtick,
		Entity:     cases.Title(language.English, cases.NoLower).String(ent.name),
		EntitySymb: ent.symb,
		EntityName: ent.name,
	}

	tmplName := "entity.normal"
	if simple {
		tmplName = "entity.simple"
	}
	filename := filepath.Join(dirEntity, ent.snakeName+".go")
	if err = saveTemplate(filename, getTemplateByDBType(dbType, tmplName), data); err != nil {
		return
	}

	tmplName = "repository.interface"
	if simple {
		tmplName = "repository.simple_interface"
	}
	filename = filepath.Join(dirRepository, ent.snakeName+"_repository_interface.go")
	err = saveTemplate(filename, getTemplate(tmplName), data)
	return
}

func createMysqlEntityFiles(module string, ent *entity, dbType *DBType) (err error) {
	if _, err = os.Stat(filepath.Join(dirRepository, dbType.name)); os.IsNotExist(err) {
		if err = createDir(filepath.Join(dirRepository, dbType.name)); err != nil {
			return
		}

		if err = initClar(module, dbType); err != nil {
			return
		}
	}

	data := struct {
		Module      string
		Backtick    string
		Table       string
		Entity      string
		EntityName  string
		EntityTable string
	}{
		Module:      modulePath,
		Backtick:    backtick,
		Table:       ent.table,
		Entity:      ent.structName,
		EntityName:  ent.name,
		EntityTable: ent.table,
	}

	filename := filepath.Join(dirRepository, dbType.name, ent.snakeName+"_repository.go")
	err = saveTemplate(filename, getTemplateByDBType(dbType, "repository.entity"), data)
	return
}
