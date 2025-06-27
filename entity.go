package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func entityCmd() *cobra.Command {
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

			ent := newEntity(args[0], empty, simple)
			if err = ent.create(); err != nil {
				fmt.Println(err)
			}
			if err = ent.createModel(); err != nil {
				fmt.Println(err)
			}
			if err = save("internal/db/utils.go", "templates/db.utils.tmpl", nil); err != nil {
				fmt.Println(err)
			}
			if err = createConnection(); err != nil {
				fmt.Println(err)
			}
			if err = createRepo(ent); err != nil {
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

	empty  bool
	simple bool
}

func newEntity(name string, empty, simple bool) (ent *entity) {
	ent = &entity{
		nameRunes: []rune(name),
		snakeName: toSnake(name),
		empty:     empty,
		simple:    simple,
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

func (ent *entity) create() error {
	tmplEntity := "entity"
	if ent.empty {
		tmplEntity = "entity.empty"

	} else if ent.simple {
		tmplEntity = "entity.simple"
	}

	entityFilename := ent.snakeName + ".go"

	dir := filepath.Join("internal/db", db, "entity")
	if err := createDir(dir); err != nil {
		return err
	}

	data := struct {
		ImportModels string
		Entity       string
		EntitySymb   string
		EntityName   string
	}{
		ImportModels: filepath.Join(modulePath, "internal/domain/models"),
		Entity:       ent.structName,
		EntitySymb:   ent.symb,
		EntityName:   ent.name,
	}
	return save(filepath.Join(dir, entityFilename), "templates/"+tmplEntity+".tmpl", data)
}

func (ent *entity) createModel() error {
	if ent.empty {
		return nil
	}
	tmplModel := "model"
	if ent.simple {
		tmplModel = "model.simple"
	}

	dir := "internal/domain/models"
	if err := createDir(dir); err != nil {
		return err
	}

	fileErr := filepath.Join(dir, "errors.go")
	if fileNotExists(fileErr) {
		if err := save(fileErr, "templates/models.errors.tmpl", nil); err != nil {
			return err
		}
	}

	data := struct {
		Entity     string
		EntitySymb string
	}{
		Entity:     ent.structName,
		EntitySymb: ent.symb,
	}
	return save(filepath.Join(dir, ent.snakeName+".go"), "templates/"+tmplModel+".tmpl", data)
}
