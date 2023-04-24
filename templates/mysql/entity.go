package mysql

const Entity = `package entity

import "time"

// {{.Entity}} entity structure.
type {{.Entity}} struct {
	ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}

	CreatedAt time.Time  {{.Backtick}}db:"created_at"{{.Backtick}}
	UpdatedAt time.Time  {{.Backtick}}db:"updated_at"{{.Backtick}}
	DeletedAt *time.Time {{.Backtick}}db:"deleted_at"{{.Backtick}}
}

func ({{.EntitySymb}} *{{.Entity}}) GetID() int64 {
	return {{.EntitySymb}}.ID
}

func ({{.EntitySymb}} *{{.Entity}}) GetCreatedAt() time.Time {
	return {{.EntitySymb}}.CreatedAt
}

func ({{.EntitySymb}} *{{.Entity}}) GetUpdatedAt() time.Time {
	return {{.EntitySymb}}.UpdatedAt
}

func ({{.EntitySymb}} *{{.Entity}}) GetDeletedAt() *time.Time {
	return {{.EntitySymb}}.DeletedAt
}
`

const SimpleEntity = `package entity

// {{.Entity}} entity structure.
type {{.Entity}} struct {
	ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
}

func ({{.EntitySymb}} *{{.Entity}}) GetID() int64 {
	return {{.EntitySymb}}.ID
}
`

const EmptyEntity = `package entity

// {{.Entity}} entity structure.
type {{.Entity}} struct {
}
`
