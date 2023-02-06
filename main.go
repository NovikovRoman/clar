package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jessevdk/go-flags"
)

const (
	permDir  = 0755
	backtick = "`"
)

type initCommand struct {
	DBType string `long:"db" short:"d" description:"DB type." default:"mysql"`
}

func (c initCommand) Execute(_ []string) (err error) {
	var dbType *DBType
	if dbType, err = newDBType(opts.Init.DBType); err != nil {
		return
	}
	return initClar(modulePath, dbType)
}

type arrayCommand struct {
	Name string `long:"name" short:"n" description:"Structure name." required:"true"`
}

func (c arrayCommand) Execute(_ []string) error {
	return createJsonArray(opts.Array.Name)
}

type structCommand struct {
	Name string `long:"name" short:"n" description:"Structure name." required:"true"`
}

func (c structCommand) Execute(_ []string) error {
	return createJsonStruct(opts.Struct.Name)
}

type entityCommand struct {
	Name   string `long:"name" short:"n" description:"Entity name." required:"true"`
	DBType string `long:"db" short:"d" description:"DB type." default:"mysql"`
	Empty  bool   `long:"empty" short:"e" description:"Empty entity."`
	Simple bool   `long:"simple" short:"s" description:"Simple entity."`
}

func (c entityCommand) Execute(_ []string) (err error) {
	var dbType *DBType
	if dbType, err = newDBType(opts.Entity.DBType); err != nil {
		return
	}
	return createEntity(modulePath, opts.Entity.Name, dbType, opts.Entity.Empty, opts.Entity.Simple)
}

type migrateCommand struct {
	DBType string `long:"db" short:"d" description:"DB type." default:"mysql"`
}

func (c migrateCommand) Execute(_ []string) (err error) {
	var dbType *DBType
	if dbType, err = newDBType(opts.Migrate.DBType); err != nil {
		return
	}
	return createMigrate(dbType)
}

var opts struct {
	Init    initCommand    `command:"init" alias:"i" description:"Creates a model interface."`
	Array   arrayCommand   `command:"array" alias:"a" description:"Creates a array structure template for json columns."`
	Struct  structCommand  `command:"struct" alias:"s" description:"Creates a structure template for json columns."`
	Entity  entityCommand  `command:"entity" alias:"e" description:"Creates a entity with a repository."`
	Migrate migrateCommand `command:"migrate" alias:"m" description:"Generates code for migration."`
}

var modulePath string

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
	modulePath = filepath.Join(string(m[1]), "internal")

	_, _ = flags.Parse(&opts)
}
