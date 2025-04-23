package templates

const BaseEntity = `package entity

import (
	"errors"
	"time"
)

var ErrNotFound error = errors.New("Entity not found")

type SimpleBaseEntity interface {
	GetID() int64
}

type BaseEntity interface {
	SimpleBaseEntity

	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
}
`
