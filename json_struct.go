package main

import (
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func createJsonStruct(name string, internal bool) (err error) {
	dirE := getPathLocation(dirEntity, internal)
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
	err = saveTemplate(filename, getTemplate("json_struct"), data)
	return
}
