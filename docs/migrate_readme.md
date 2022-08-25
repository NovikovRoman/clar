# Migration Tools

Build and run:
```
cd migrate && go build -o "migrate"
migrate /path/migrations/dir/ /path/to/.env
```

**Optional keys:**
- `-u` migration up. Default one step.

- `-d` rollback migration. Default one step.

- `-sn` - `n` migration steps.

- `-f` force the current version of the migration to be installed.

Running without optional keys will execute all required versions of the migration.

You cannot specify the keys `-u` and `-d` at the same time. Only `-u` will be executed.

`/path/to/.env` is optional. Default `.env`.

`.env` must contain the `db` parameter. Example:
```
…
db = user:pass@tcp(localhost:3306)/dbname?parseTime=true&multiStatements=true
…
```

## Naming Migration

Migration up:

```
[version]_[name].up.sql
```

Rollback migration:

```
[version]_[name].down.sql
```

`version` - unsigned integer. Usually as a date and time in the format `YYYYmmddHHii`.
`name` - the name of the migration for convenience.