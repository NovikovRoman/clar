# user.go

```go
package entity

// User entity structure.
type User struct {
	ID   int64  `db:"id" pkey:"true"`
}

func (u *User) GetID() int64 {
	return u.ID
}
```