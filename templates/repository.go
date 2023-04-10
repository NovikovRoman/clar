package templates

const Repository = `package repository

import (
	"context"

	"{{.Module}}/domain/entity"
)

type {{.Entity}}Repository interface {
	Table() string
	GetByID(ctx context.Context, id int64) ({{.EntityName}} *entity.{{.Entity}}, err error)
	Save(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
	SaveMultiple(ctx context.Context, {{.EntityName}} ...*entity.{{.Entity}}) error
	Update(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
	Remove(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
}
`

const EmptyRepository = `package repository

type {{.Entity}}Repository interface {
	Table() string
}
`
