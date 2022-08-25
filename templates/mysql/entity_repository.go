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

func New{{.Entity}}Repository(db *sqlx.DB) repository.{{.Entity}}RepositoryInterface {
	return &{{.EntityName}}Repository{
		table: "{{.EntityTable}}",
		db:    db,
	}
}

func (r *{{.EntityName}}Repository) GetTable() string {
	return r.table
}

func (r *{{.EntityName}}Repository) GetByID(ctx context.Context, id int64) ({{.EntityName}} *entity.{{.Entity}}, err error) {
	{{.EntityName}} = &entity.{{.Entity}}{}
	err = r.db.Get(&{{.EntityName}}, "SELECT * FROM {{.Backtick}}"+r.table+"{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}} = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		{{.EntityName}} = nil
	}
	return
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
