# my_struct.go

```go
package entity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 bool
}

func (m *MyStruct) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MyStruct) Scan(val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		if bytes.Compare(v, []byte("[]")) == 0 {
			return
		}
		err = json.Unmarshal(v, m)
		return

	case string:
		if v == "[]" {
			return
		}
		err = json.Unmarshal([]byte(v), m)
		return

	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (m *MyStruct) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MyStruct) ConvertValue() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "[]", err
	}
	return string(b), nil
}
```