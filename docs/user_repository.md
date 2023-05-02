# user_repository.go

```go
package mysql

import (
	"context"
	"database/sql"

	"github.com/NovikovRoman/clar/domain/entity"
	"github.com/NovikovRoman/clar/domain/repository"
	"github.com/jmoiron/sqlx"
)

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

func (r *userRepository) GetByID(ctx context.Context, id int64) (user *entity.User, err error) {
	user = &entity.User{}
	err = r.db.GetContext(ctx, &user, "SELECT * FROM `"+r.table+"` WHERE `id` = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		user = nil
	}
	return
}

// SaveMultiple saves multiple entries to the database. Adds new, updates existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 `db:"id" pkey:"true"`
func (r *userRepository) SaveMultiple(ctx context.Context, user ...*entity.User) error {
	items := make([]interface{}, len(user))
	for i, item := range user {
		items[i] = item
	}
	return saveMultiple(ctx, r.db, r.table, items...)
}

// SaveMultipleIgnoreDuplicates saves multiple entries to the database. Adds new, ignore existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
func (r *userRepository) SaveMultipleIgnoreDuplicates(ctx context.Context, user ...*entity.User) error {
	items := make([]interface{}, len(user))
	for i, item := range user {
		items[i] = item
	}
	return saveMultipleIgnoreDuplicates(ctx, r.db, r.table, items...)
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