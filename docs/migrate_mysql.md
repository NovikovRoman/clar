# migrate.go

```go
package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func newMigrate(db *sql.DB, path string, force bool) (m *migrate.Migrate, err error) {
	var driver database.Driver

	driver, err = mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "schema_migrations",
		DatabaseName:    "migrations",
	})
	if err != nil {
		return
	}

	m, err = migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), "migrations", driver)
	if err != nil {
		return
	}

	if !force {
		return
	}

	var version uint
	if version, _, err = m.Version(); err != nil {
		return
	}

	err = m.Force(int(version))
	return
}

func Migration(db *sql.DB, path string, force bool) (err error) {
	err = MigrationUp(db, path, 0, force)
	return
}

func MigrationUp(db *sql.DB, path string, step int, force bool) (err error) {
	var m *migrate.Migrate
	if m, err = newMigrate(db, path, force); err != nil {
		if err == migrate.ErrNoChange {
			err = nil
		}
		return
	}

	if step > 0 {
		if err = m.Steps(step); err == migrate.ErrNoChange {
			err = nil
		}
		return
	}

	if err = m.Up(); err == migrate.ErrNoChange {
		err = nil
	}
	return
}

func MigrationDown(db *sql.DB, path string, step int, force bool) (err error) {
	var m *migrate.Migrate
	if m, err = newMigrate(db, path, force); err != nil {
		if err == migrate.ErrNoChange {
			err = nil
		}
		return
	}

	if step > 0 {
		err = errors.New("Steps not specified.")
		return
	}

	if err = m.Steps(-step); err == migrate.ErrNoChange {
		err = nil
	}
	return
}
```

## Usage example
```go
…
import (
	…
	migrate "[your_module]/domain/migrate/mysql"
	…
)
…
step, _ := strconv.Atoi(strings.TrimSpace(os.Getenv("MIGRATE_STEP")))
force := strings.TrimSpace(os.Getenv("MIGRATE_FORCE")) != ""
if os.Getenv("MIGRATE_STEP") == "down" {
	err = migrate.MigrationDown(db.DB, migrations, step, force)
} else {
	err = migrate.MigrationUp(db.DB, migrations, step, force)
}

if err != nil {
	log.Fatalf("migrate %s\n", err)
}
```