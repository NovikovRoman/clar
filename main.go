package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

const (
	permDir  = 0755
	backtick = "`"
)

const (
	dbMysql = "mysql"
)

func main() {
	var err error
	if _, err = os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("go.mod not found.")
		os.Exit(1)
	}

	var b []byte
	if b, err = os.ReadFile("go.mod"); err != nil {
		fmt.Println("go.mod cannot be read.")
		os.Exit(1)
	}

	m := regexp.MustCompile(`(?si)module\s+(.+?)\s`).FindSubmatch(b)
	if len(m) == 0 {
		fmt.Println("Module not found in go.mod.")
		os.Exit(1)
	}
	modulePath := filepath.Join(string(m[1]), "internal", "db")

	rootCmd := &cobra.Command{
		Use:   "clar [entity|struct|array|migrate]",
		Short: "",
	}

	var dbType string
	rootCmd.PersistentFlags().StringVarP(&dbType, "db", "d", dbMysql, "Database type")

	migrateCmd := migrateCmd(dbType)
	jsonStructCmd := jsonStructCmd(dbType)
	jsonArrayCmd := jsonArrayCmd(dbType)
	entityCmd := entityCmd(dbType, modulePath)

	rootCmd.AddCommand(migrateCmd, jsonStructCmd, jsonArrayCmd, entityCmd)

	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
	if err = dbIsSupported(dbType); err != nil {
		fmt.Println(err)
	}
}

func dbIsSupported(dbType string) error {
	if dbType == dbMysql {
		return nil
	}
	return fmt.Errorf("Database type %s not supported", dbType)
}
