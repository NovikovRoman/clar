package main

import (
	"fmt"
	"path/filepath"
)

func createConnection() error {
	filepath := "internal/db/" + db + "/connecton.go"
	if !fileNotExists(filepath) {
		fmt.Println("File already exists:", filepath)
		return nil
	}

	data := struct {
		ModulePath string
	}{
		ModulePath: modulePath,
	}
	err := save(filepath, "templates/connection."+db+".tmpl", data)
	if err != nil {
		return err
	}
	if fileNotExists("migrations") {
		err = createMigrate()
	}
	return err
}

func createRepo(ent *entity) error {
	if ent.empty {
		return nil
	}

	dirR := "internal/repository"
	if err := createDir(dirR); err != nil {
		return err
	}
	dirDbR := filepath.Join("internal/db", db, "repositories")
	if err := createDir(dirDbR); err != nil {
		return err
	}

	fileHelper := filepath.Join(dirDbR, "helpers.go")
	if fileNotExists(fileHelper) {
		if err := save(fileHelper, "templates/helpers."+db+".tmpl", nil); err != nil {
			return err
		}
	}

	data := struct {
		ModulePath  string
		DbType      string
		Entity      string
		EntitySymb  string
		EntityName  string
		EntityTable string
	}{
		ModulePath:  modulePath,
		DbType:      db,
		Entity:      ent.structName,
		EntitySymb:  ent.symb,
		EntityName:  ent.name,
		EntityTable: ent.table,
	}

	tmpl := "templates/repository.db.tmpl"
	if ent.simple {
		tmpl = "templates/repository.db.simple.tmpl"
	}

	if err := save(filepath.Join(dirDbR, ent.name+".go"), tmpl, data); err != nil {
		return err
	}

	if !ent.simple {
		err := save(filepath.Join(dirR, data.EntityName+".go"), "templates/repository.tmpl", data)
		if err != nil {
			return err
		}
		fileRepositories := filepath.Join(dirR, "repositories.go")
		if fileNotExists(fileRepositories) {
			if err = save(fileRepositories, "templates/repositories.tmpl", data); err != nil {
				return err
			}
		}
		fileDbRepositories := filepath.Join(dirDbR, "repositories.go")
		if fileNotExists(fileDbRepositories) {
			if err = save(fileDbRepositories, "templates/repositories.db.tmpl", data); err != nil {
				return err
			}
		}
	}
	return nil
}
