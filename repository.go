package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func createConnection() error {
	dir := filepath.Join("internal/domain/db", db)
	if fileNotExists(dir) {
		if err := createDir(dir); err != nil {
			return err
		}
	}

	fileUtils := "internal/domain/db/utils.go"
	if fileNotExists(fileUtils) {
		if err := save(fileUtils, "templates/db.utils.tmpl", nil); err != nil {
			return err
		}
	}

	connFile := "internal/domain/db/" + db + "/connecton.go"
	if !fileNotExists(connFile) {
		fmt.Println("File already exists:", connFile)
		return nil
	}

	data := struct {
		ModulePath string
	}{
		ModulePath: modulePath,
	}
	err := save(connFile, "templates/connection."+db+".tmpl", data)
	if err != nil {
		return err
	}
	if fileNotExists("migrations") {
		err = createMigrate()
	}
	return err
}

func createRepo(ent entity) error {
	if ent.empty {
		return nil
	}

	dir := filepath.Join("internal/domain/db", db)
	if err := createDir(dir); err != nil {
		return err
	}

	fileHelper := filepath.Join(dir, "helpers.go")
	if fileNotExists(fileHelper) {
		if err := save(fileHelper, "templates/helpers."+db+".tmpl", nil); err != nil {
			return err
		}
	}

	data := struct {
		ModulePath     string
		DbType         string
		FirstUpperName string
		EntitySymb     string
		FirstLowerName string
		EntityTable    string
		SnakeName      string
		Alias          string
	}{
		ModulePath:     modulePath,
		DbType:         db,
		FirstUpperName: ent.firstUpperName,
		EntitySymb:     ent.symb,
		FirstLowerName: ent.firstLowerName,
		EntityTable:    ent.tableName,
		SnakeName:      ent.snakeName,
		Alias:          ent.snakeName,
	}
	if strings.Contains(ent.snakeName, "_") {
		data.Alias = ent.packageName
	}

	tmpl := "templates/repository.db.tmpl"
	if ent.simple {
		tmpl = "templates/repository.db.simple.tmpl"
	}

	if err := save(filepath.Join(dir, ent.snakeName+".go"), tmpl, data); err != nil {
		return err
	}

	tmpl = "templates/repository.tmpl"
	if ent.simple {
		tmpl = "templates/repository.simple.tmpl"
	}
	data2 := struct {
		PackageName string
		EntityName  string
	}{
		PackageName: ent.packageName,
		EntityName:  ent.firstLowerName,
	}
	return save(filepath.Join(ent.getDir(), "repository.go"), tmpl, data2)
}
