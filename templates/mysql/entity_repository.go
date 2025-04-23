package mysql

const EntityRepository = `package repository

import (
	"context"
	"database/sql"

	"{{.Module}}/{{.DBType}}/entity"
	"github.com/jmoiron/sqlx"
)

type {{.Entity}}Repository interface {
	Table() string
	ByID(ctx context.Context, id int64) (entity.{{.Entity}}, error)
	Save(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error
	SaveMultiple(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error
	SaveMultipleIgnoreDuplicates(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error
	Update(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error
	Remove(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) error
}

type {{.EntityName}}Repository struct {
	table string
	db    *sqlx.DB
}

func New{{.Entity}}Repository(db *sqlx.DB) {{.Entity}}Repository {
	return &{{.EntityName}}Repository{
		table: "{{.EntityTable}}",
		db:    db,
	}
}

func (r *{{.EntityName}}Repository) Table() string {
	return r.table
}

// start CLAR generation ---------------------------------------------------------------------------

func (r *{{.EntityName}}Repository) ByID(ctx context.Context, id int64) (entity.{{.Entity}}, error) {
	var {{.EntityName}} entity.{{.Entity}}
	err := r.db.GetContext(ctx, &{{.EntityName}}, "SELECT * FROM {{.Backtick}}"+r.table+"{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}} = ?", id)
	if err == sql.ErrNoRows {
		return {{.EntityName}}, ErrNotFound
	}
	return {{.EntityName}}, err
}

// SaveMultiple saves multiple entries to the database. Adds new, updates existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func (r *{{.EntityName}}Repository) SaveMultiple(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error {
	items := make([]any, len({{.EntityName}}))
	for i, item := range {{.EntityName}} {
		items[i] = item
	}
	return saveMultiple(ctx, r.db, r.table, items...)
}

// SaveMultipleIgnoreDuplicates saves multiple entries to the database. Adds new, ignore existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func (r *{{.EntityName}}Repository) SaveMultipleIgnoreDuplicates(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error {
	items := make([]any, len({{.EntityName}}))
	for i, item := range {{.EntityName}} {
		items[i] = item
	}
	return saveMultipleIgnoreDuplicates(ctx, r.db, r.table, items...)
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

// end CLAR generation -----------------------------------------------------------------------------
`

const EmptyEntityRepository = `package repository

import (
	"github.com/jmoiron/sqlx"
)

type {{.Entity}}Repository interface {
	Table() string
}

type {{.EntityName}}Repository struct {
	table string
	db    *sqlx.DB
}

func New{{.Entity}}Repository(db *sqlx.DB) {{.Entity}}Repository {
	return &{{.EntityName}}Repository{
		table: "{{.EntityTable}}",
		db:    db,
	}
}

func (r *{{.EntityName}}Repository) Table() string {
	return r.table
}
`
