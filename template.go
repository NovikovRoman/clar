package main

import (
	"errors"
	"os"
	"text/template"

	"github.com/NovikovRoman/clar/templates"
	"github.com/NovikovRoman/clar/templates/mysql"
)

func getTemplate(name string) string {
	switch name {
	case "interface":
		return templates.EntityInterface

	case "json_struct":
		return templates.JsonStruct
	case "json_array":
		return templates.StringArray

	case "repository.interface":
		return templates.RepositoryInterface

	default:
		return ""
	}
}

func getTemplateByDBType(dbType *DBType, name string) string {
	switch dbType.code {
	case MysqlCode:
		return getTemplateMysql(name)

	default:
		return ""
	}
}

func getTemplateMysql(name string) string {
	switch name {
	case "migrate.readme":
		return mysql.MigrateReadme
	case "migrate.main":
		return mysql.MigrateMain
	case "migrate.up":
		return mysql.MigrateUp
	case "migrate.down":
		return mysql.MigrateDown

	case "repository.utils":
		return mysql.RepositoryUtils
	case "repository.entity":
		return mysql.EntityRepository

	case "entity.normal":
		return mysql.Entity
	case "entity.simple":
		return mysql.SimpleEntity

	default:
		return ""
	}
}

// saveTemplate creates a file from a template.
func saveTemplate(filename string, tmpl string, data interface{}) (err error) {
	var f *os.File

	if _, err = os.Stat(filename); err == nil {
		return errors.New(filename + " file exists.")
	}

	if f, err = os.Create(filename); err != nil {
		return
	}

	defer func() {
		_ = f.Close()
	}()

	t := template.Must(template.New("").Parse(tmpl))
	err = t.Execute(f, data)
	return
}
