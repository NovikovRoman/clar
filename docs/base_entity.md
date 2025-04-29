# base_entity.go

```go
package entity

import (
    "errors"
    "time"
)

type SimpleBaseEntity interface {
    GetID() int64
}

type BaseEntity interface {
    SimpleEntity

    GetCreatedAt() time.Time
    GetUpdatedAt() time.Time
    GetDeletedAt() *time.Time
}
```
