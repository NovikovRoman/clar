package templates

const RepositoryInterface = `package repository

import (
	"context"

	"{{.Module}}/domain/entity"
)

type {{.Entity}}RepositoryInterface interface {
	GetTable() string
	GetByID(ctx context.Context, id int64) ({{.EntityName}} *entity.{{.Entity}}, err error)
	Save(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
	Update(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
	Remove(ctx context.Context, {{.EntityName}} *entity.{{.Entity}}) (err error)
}
`

const SimpleRepositoryInterface = `package repository

import (
	"context"

	"{{.Module}}/domain/entity"
)

type {{.Entity}}RepositoryInterface interface {
	GetTable() string
	GetByID(ctx context.Context, id int64) ({{.EntityName}} *entity.{{.Entity}}, err error)
}
`
