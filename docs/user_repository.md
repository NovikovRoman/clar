# user_repository.go

```go
package mysql

import (
	"context"

	"github.com/NovikovRoman/clar/domain/entity"
	"github.com/NovikovRoman/clar/domain/repository"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	table string
	db    *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepositoryInterface {
	return &userRepository{
		table: "users",
		db:    db,
	}
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (user *entity.User, err error) {
	// our code …
	return
}

func (r *userRepository) Save(ctx context.Context, user *entity.User) error {
	return save(ctx, r.db, r.table, user)
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return update(ctx, r.db, r.table, user)
}

func (r *userRepository) Remove(ctx context.Context, user *entity.User) error {
	return remove(ctx, r.db, r.table, user)
}
```