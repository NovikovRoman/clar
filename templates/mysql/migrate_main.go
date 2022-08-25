package mysql

const MigrateMain = `package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var opts struct {
	Dir   string {{.Backtick}}long:"dir" description:"Путь к директории миграции." required:"true"{{.Backtick}}
	Env   string {{.Backtick}}short:"e" long:"env" description:"Конфигурация приложения. Должен быть параметр DB=…." default:".env"{{.Backtick}}
	Up    bool   {{.Backtick}}short:"u" long:"up" description:"Миграция назад."{{.Backtick}}
	Down  bool   {{.Backtick}}short:"d" long:"down" description:"Миграция назад."{{.Backtick}}
	Step  int    {{.Backtick}}short:"s" long:"step" description:"Количество шагов." default:"1"{{.Backtick}}
	Force bool   {{.Backtick}}short:"f" long:"force" description:"Принудительно выполнить текущую version."{{.Backtick}}
}

func main() {
	var (
		dirFullPath string
		err         error
		db          *sql.DB
		driver      database.Driver
		m           *migrate.Migrate
		version     uint
	)

	if _, err = flags.ParseArgs(&opts, os.Args); err != nil {
		log.Fatalln(err)
	}

	if err = godotenv.Load(opts.Env); err != nil {
		log.Fatalln("Error loading env file")
	}

	if dirFullPath, err = filepath.Abs(opts.Dir); err != nil {
		log.Fatalln(err)
	}

	if db, err = sql.Open("mysql", os.Getenv("DB")); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if derr := db.Close(); derr != nil {
			log.Fatal(derr)
		}
	}()

	driver, err = mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "schema_migrations",
		DatabaseName:    "migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}

	m, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", dirFullPath), "migrations", driver)
	if err != nil {
		log.Fatalln(err)
	}

	if opts.Force {
		version, _, err = m.Version()
		if err != nil {
			log.Fatalln(err)
		}
		if err = m.Force(int(version)); err != nil {
			log.Fatalln(err)
		}
	}

	if opts.Up {
		err = m.Steps(opts.Step)

	} else if opts.Down {
		err = m.Steps(-opts.Step)

	} else {
		err = m.Up()
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln(err)

	} else if err == migrate.ErrNoChange {
		log.Println(err)
	}
}
`
