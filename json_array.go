package main

import (
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func createJsonArray(name string) (err error) {
	if err = createDir(dirEntity); err != nil {
		return
	}

	data := struct {
		Struct     string
		StructSymb string
	}{
		Struct:     cases.Title(language.English, cases.NoLower).String(name),
		StructSymb: strings.ToLower(string([]rune(name)[0])),
	}

	filename := filepath.Join(dirEntity, toSnake(name)+".go")
	err = saveTemplate(filename, getTemplate("json_array"), data)
	return
}
