package mysql

const MigrateReadme = `# Migration Tools

Build and run:
{{.Backtick}}{{.Backtick}}{{.Backtick}}
cd migrate && go build -o "migrate"
migrate /path/migrations/dir/ /path/to/.env
{{.Backtick}}{{.Backtick}}{{.Backtick}}

**Optional keys:**
- {{.Backtick}}-u{{.Backtick}} migration up. Default one step.

- {{.Backtick}}-d{{.Backtick}} rollback migration. Default one step.

- {{.Backtick}}-sn{{.Backtick}} - {{.Backtick}}n{{.Backtick}} migration steps.

- {{.Backtick}}-f{{.Backtick}} force the current version of the migration to be installed.

Running without optional keys will execute all required versions of the migration.

You cannot specify the keys {{.Backtick}}-u{{.Backtick}} and {{.Backtick}}-d{{.Backtick}} at the same time. Only {{.Backtick}}-u{{.Backtick}} will be executed.

{{.Backtick}}/path/to/.env{{.Backtick}} is optional. Default {{.Backtick}}.env{{.Backtick}}.

{{.Backtick}}.env{{.Backtick}} must contain the {{.Backtick}}db{{.Backtick}} parameter. Example:
{{.Backtick}}{{.Backtick}}{{.Backtick}}
…
db = user:pass@tcp(localhost:3306)/dbname?parseTime=true&multiStatements=true
…
{{.Backtick}}{{.Backtick}}{{.Backtick}}

## Naming Migration

Migration up:

{{.Backtick}}{{.Backtick}}{{.Backtick}}
[version]_[name].up.sql
{{.Backtick}}{{.Backtick}}{{.Backtick}}

Rollback migration:

{{.Backtick}}{{.Backtick}}{{.Backtick}}
[version]_[name].down.sql
{{.Backtick}}{{.Backtick}}{{.Backtick}}

{{.Backtick}}version{{.Backtick}} - unsigned integer. Usually as a date and time in the format {{.Backtick}}YYYYmmddHHii{{.Backtick}}.
{{.Backtick}}name{{.Backtick}} - the name of the migration for convenience.`

