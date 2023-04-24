# user.go

```go
package entity

import "time"

// User entity structure.
type User struct {
	ID   int64  `db:"id" pkey:"true"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) GetDeletedAt() *time.Time {
	return u.DeletedAt
}
```