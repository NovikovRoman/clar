package templates

const EntityInterface = `package entity

import "time"

type SimpleEntityInterface interface {
	GetID() int64
}

type EntityInterface interface {
	SimpleEntityInterface

	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
}
`
