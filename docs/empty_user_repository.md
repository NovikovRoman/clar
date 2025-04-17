# user_repository.go

```go
package mysql

import (
    "github.com/NovikovRoman/clar/domain/repository"
    "github.com/jmoiron/sqlx"
)

type UserRepository interface {
    Table() string
}

type userRepository struct {
    table string
    db    *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
    return &userRepository{
        table: "users",
        db:    db,
    }
}

func (r *userRepository) Table() string {
    return r.table
}
