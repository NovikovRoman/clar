package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func jsonArrayCmd(dbType string) *cobra.Command {
	return &cobra.Command{
		Use:     "array [name]",
		Aliases: []string{"a"},
		Short:   "Create json array",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := createJsonArray(dbType, args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func createJsonArray(dbType, name string) (err error) {
	dirE := getPathLocation(dbType, dirEntity)
	if err = createDir(dirE); err != nil {
		return
	}

	data := struct {
		Struct     string
		StructSymb string
	}{
		Struct:     cases.Title(language.English, cases.NoLower).String(name),
		StructSymb: strings.ToLower(string([]rune(name)[0])),
	}

	filename := filepath.Join(dirE, toSnake(name)+".go")
	err = saveTemplate(filename, getTemplate("json_array"), data)
	return
}
