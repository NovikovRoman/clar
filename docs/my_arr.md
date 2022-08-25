# my_arr.go

```go
package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MyArr []string // can change the type int, int32, etc.

func (m MyArr) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MyArr) Scan(val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &m)

	case string:
		return json.Unmarshal([]byte(v), &m)

	default:
		return fmt.Errorf("Unsupported type: %T. ", v)
	}
}

func (m MyArr) Value() (driver.Value, error) {
	return json.Marshal(m)
}
```