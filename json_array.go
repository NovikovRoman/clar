package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func jsonArrayCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "array [name]",
		Aliases: []string{"a"},
		Short:   "Create json array",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := createJsonArray(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func createJsonArray(name string) (err error) {
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
	return save(filepath.Join(dir, toSnake(name)+".go"), "templates/json_array.tmpl", data)
}
