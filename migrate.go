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
	switch db {
	case dbPostgres:
		err := save(filepath.Join(pathSql, "202505200000_types.up.sql"), "templates/migrate.pg.types.up.tmpl", nil)
		if err != nil {
			return err
		}
		err = save(filepath.Join(pathSql, "202505200000_types.down.sql"), "templates/migrate.pg.types.down.tmpl", nil)
		if err != nil {
			return err
		}
		err = save(filepath.Join(pathSql, "202505201000_users.up.sql"), "templates/migrate.pg.users.up.tmpl", nil)
		if err != nil {
			return err
		}
		err = save(filepath.Join(pathSql, "202505201000_users.down.sql"), "templates/migrate.users.down.tmpl", nil)
		return err

	default:
		err := save(filepath.Join(pathSql, "202505200000_users.up.sql"), "templates/migrate."+db+".users.up.tmpl", nil)
		if err != nil {
			return err
		}
		err = save(filepath.Join(pathSql, "202505200000_users.down.sql"), "templates/migrate.users.down.tmpl", nil)
		return err
	}
}
