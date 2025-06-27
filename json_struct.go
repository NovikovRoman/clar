package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func jsonStructCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "struct [name]",
		Aliases: []string{"s"},
		Short:   "Create json struct",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := createJsonStruct(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func createJsonStruct(name string) (err error) {
	dir := "internal/domain/models"
	if err = createDir(dir); err != nil {
		return
	}

	data := struct {
		Struct     string
		StructSymb string
	}{
		Struct:     cases.Title(language.English, cases.NoLower).String(name),
		StructSymb: strings.ToLower(string([]rune(name)[0])),
	}
	return save(filepath.Join(dir, toSnake(name)+".go"), "templates/json_struct.tmpl", data)
}
