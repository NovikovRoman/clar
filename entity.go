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
			if err = createConnection(); err != nil {
				fmt.Println(err)
			}
			if err = createRepo(ent); err != nil {
				fmt.Println(err)
			}
			if err = createService(ent); err != nil {
				fmt.Println(err)
			}
			if err = createFactory(); err != nil {
				fmt.Println(err)
			}
		},
	}

	entityCmd.Flags().BoolVarP(&empty, "empty", "e", false, "Create empty entity")
	entityCmd.Flags().BoolVarP(&simple, "simple", "s", false, "Create simple entity")
	return entityCmd
}

type entity struct {
	firstLowerName string
	firstUpperName string
	packageName    string
	snakeName      string
	tableName      string
	nameRunes      []rune
	symb           string

	empty  bool
	simple bool
}

func newEntity(name string, empty, simple bool) entity {
	ent := entity{
		nameRunes: []rune(name),
		snakeName: toSnake(name),
		empty:     empty,
		simple:    simple,
	}
	ent.symb = strings.ToLower(string(ent.nameRunes[0]))
	ent.firstLowerName = strings.ToLower(string(ent.nameRunes[0])) + string(ent.nameRunes[1:])
	ent.firstUpperName = cases.Title(language.English, cases.NoLower).String(ent.firstLowerName)
	ent.packageName = strings.ReplaceAll(ent.snakeName, "_", "")

	ent.tableName = ent.snakeName
	switch []rune(ent.tableName)[len(ent.tableName)-1] {
	case 's', 'x', 'h':
		ent.tableName += "es"

	case 'y':
		ent.tableName = string([]rune(ent.tableName)[:len(ent.tableName)-1]) + "ies"

	default:
		ent.tableName += "s"
	}
	return ent
}

func (ent *entity) create() error {
	tmplEntity := "entity"
	if ent.empty {
		tmplEntity = "entity.empty"

	} else if ent.simple {
		tmplEntity = "entity.simple"
	}

	dir := ent.getDir()
	if err := createDir(dir); err != nil {
		return err
	}
	data := struct {
		PackageName string
	}{
		PackageName: ent.packageName,
	}
	return save(filepath.Join(dir, "entity.go"), "templates/"+tmplEntity+".tmpl", data)
}

func (ent *entity) getDir() string {
	return filepath.Join("internal/domain/", ent.snakeName)
}
