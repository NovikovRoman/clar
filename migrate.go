package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func migrateCmd(dbType string) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Aliases: []string{"m"},
		Short: "Create migrate",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := createMigrate(dbType); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func createMigrate(dbType string) (err error) {
	path := filepath.Join("internal", dirMigrate, dbType)
	pathSql := filepath.Join(path, "migrations")
	if err = createDir(pathSql); err != nil {
		return
	}

	data := struct {
		Backtick string
	}{
		Backtick: backtick,
	}

	if err = saveTemplate(path+"/migrate.go", getTemplateByDBType(dbType, "migrate"), data); err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202205041600_begin.up.sql"), getTemplateByDBType(dbType, "migrate.up"), data)
	if err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202205041600_begin.down.sql"), getTemplateByDBType(dbType, "migrate.down"), data)
	return
}
