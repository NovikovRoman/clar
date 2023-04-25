# user_repository_interface.go

```go
package repository

import (
	"context"

	"github.com/NovikovRoman/clar/domain/entity"
)

type UserRepository interface {
	Table() string
	GetByID(ctx context.Context, id int64) (user *entity.User, err error)
	Save(ctx context.Context, user *entity.User) (err error)
	SaveMultiple(ctx context.Context, user ...*entity.User) error
	SaveMultipleIgnoreDuplicates(ctx context.Context, user ...*entity.User) error
	Update(ctx context.Context, user *entity.User) (err error)
	Remove(ctx context.Context, user *entity.User) (err error)
}
```