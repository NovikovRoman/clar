package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func entityCmd(dbType, modulePath string) *cobra.Command {
	var (
		empty  bool
		simple bool
	)
	entityCmd := &cobra.Command{
		Use:     "entity [name]",
		Aliases: []string{"e"},
		Short:   "Create entity",
		Args:    cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if empty && simple {
				fmt.Println("Error: Flags -s and -e cannot be used together")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			empty, err := cmd.Flags().GetBool("empty")
			if err != nil {
				fmt.Println(err)
				return
			}
			simple, err := cmd.Flags().GetBool("simple")
			if err != nil {
				fmt.Println(err)
				return
			}
			if err = createEntity(modulePath, dbType, args[0], empty, simple); err != nil {
				fmt.Println(err)
			}
		},
	}

	entityCmd.Flags().BoolVarP(&empty, "empty", "e", false, "Create empty entity")
	entityCmd.Flags().BoolVarP(&simple, "simple", "s", false, "Create simple entity")
	return entityCmd
}

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

func createEntity(modulePath, dbType, name string, empty, simple bool) error {
	ent := newEntity(name)
	if err := createEntityFile(dbType, ent, empty, simple); err != nil {
		return err
	}

	switch dbType {
	case dbMysql:
		return createMysqlRepositoryFile(modulePath, dbType, ent, empty)
	}

	return errors.New("Database type not supported.")
}

func createEntityFile(dbType string, ent *entity, empty, simple bool) error {
	tmplEntity := "entity.normal"
	if empty {
		tmplEntity = "entity.empty"

	} else if simple {
		tmplEntity = "entity.simple"
	}

	entityFilename := ent.snakeName + ".go"

	dirE := getPathLocation(dbType, dirEntity)
	if err := createDir(dirE); err != nil {
		return err
	}

	data := struct {
		Module     string
		Backtick   string
		Entity     string
		EntitySymb string
		EntityName string
	}{
		Backtick:   backtick,
		Entity:     cases.Title(language.English, cases.NoLower).String(ent.name),
		EntitySymb: ent.symb,
		EntityName: ent.name,
	}
	filePath := filepath.Join(dirE, entityFilename)
	return saveTemplate(filePath, getTemplateByDBType(dbType, tmplEntity), data)
}

func createMysqlRepositoryFile(modulePath, dbType string, ent *entity, empty bool) (err error) {
	dirDbType := getPathLocation(dbType, dirRepository)
	if _, err = os.Stat(filepath.Join(dirDbType, "utils.go")); os.IsNotExist(err) {
		if err = createDir(dirDbType); err != nil {
			return
		}

		if err = initClar(modulePath, dbType); err != nil {
			return
		}
	}

	data := struct {
		Module      string
		DBType      string
		Backtick    string
		Table       string
		Entity      string
		EntityName  string
		EntityTable string
	}{
		Module:      modulePath,
		DBType:      dbType,
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
