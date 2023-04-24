package mysql

const EntityRepository = `package mysql

import (
	"context"
	"database/sql"

	"{{.Module}}/domain/entity"
	"{{.Module}}/domain/repository"
	"github.com/jmoiron/sqlx"
)

type {{.EntityName}}Repository struct {
	table string
	db    *sqlx.DB
}

func New{{.Entity}}Repository(db *sqlx.DB) repository.{{.Entity}}Repository {
	return &{{.EntityName}}Repository{
		table: "{{.EntityTable}}",
		db:    db,
	}
}

func (r *{{.EntityName}}Repository) Table() string {
	return r.table
}

func (r *{{.EntityName}}Repository) GetByID(ctx context.Context, id int64) ({{.EntityName}} *entity.{{.Entity}}, err error) {
	{{.EntityName}} = &entity.{{.Entity}}{}
	err = r.db.GetContext(ctx, {{.EntityName}}, "SELECT * FROM {{.Backtick}}"+r.table+"{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}} = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		{{.EntityName}} = nil
	}
	return
}

// SaveMultiple saves multiple entries to the database. Adds new, updates existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func (r *{{.EntityName}}Repository) SaveMultiple(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error {
	items := make([]entity.SimpleBaseEntity, len({{.EntityName}}))
	for i, item := range {{.EntityName}} {
		items[i] = item
	}
	return saveMultiple(ctx, r.db, r.table, items...)
}

// InsertIgnoreDuplicates inserts multiple records into the database.
// [!] be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func (r *{{.EntityName}}Repository) InsertIgnoreDuplicates(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error {
	items := make([]entity.SimpleBaseEntity, len({{.EntityName}}))
	for i, item := range {{.EntityName}} {
		items[i] = item
	}
	return insertIgnoreDuplicates(ctx, r.db, r.table, items...)
}

func (r *{{.EntityName}}Repository) Save(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error {
	return save(ctx, r.db, r.table, {{.EntityName}})
}

func (r *{{.EntityName}}Repository) Update(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error {
	return update(ctx, r.db, r.table, {{.EntityName}})
}

func (r *{{.EntityName}}Repository) Remove(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error {
	return remove(ctx, r.db, r.table, {{.EntityName}})
}
`

const EmptyEntityRepository = `package mysql

import (
	"{{.Module}}/domain/repository"
	"github.com/jmoiron/sqlx"
)

type {{.EntityName}}Repository struct {
	table string
	db    *sqlx.DB
}

func New{{.Entity}}Repository(db *sqlx.DB) repository.{{.Entity}}Repository {
	return &{{.EntityName}}Repository{
		table: "{{.EntityTable}}",
		db:    db,
	}
}

func (r *{{.EntityName}}Repository) Table() string {
	return r.table
}
`
