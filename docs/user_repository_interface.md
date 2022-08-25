# user_repository_interface.go

```go
package repository

import (
	"context"

	"github.com/NovikovRoman/clar/domain/entity"
)

type UserRepositoryInterface interface {
	GetTable() string
	GetByID(ctx context.Context, id int64) (user *entity.User, err error)
	Save(ctx context.Context, user *entity.User) (err error)
	Update(ctx context.Context, user *entity.User) (err error)
	Remove(ctx context.Context, user *entity.User) (err error)
}
```