package templates

const BaseEntity = `package entity

import "time"

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
