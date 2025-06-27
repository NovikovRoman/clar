package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func migrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"m"},
		Short:   "Create migrate",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := createMigrate(); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func createMigrate() error {
	path := filepath.Join("migrations", string(db))
	pathSql := filepath.Join(path, "migrations")
	if err := createDir(pathSql); err != nil {
		return err
	}
	if err := save(path+"/migrate.go", "templates/migrate."+db+".tmpl", nil); err != nil {
		return err
	}
	err := save(filepath.Join(pathSql, "202205041600_begin.up.sql"), "templates/migrate.up."+db+".tmpl", nil)
	if err != nil {
		return err
	}
	err = save(filepath.Join(pathSql, "202205041600_begin.down.sql"), "templates/migrate.down.tmpl", nil)
	return err
}
