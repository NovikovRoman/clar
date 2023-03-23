# my_struct.go

```go
package entity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 bool
}

func (m MyStruct) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MyStruct) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		if bytes.Equal(v, []byte("[]")) || bytes.Equal(v, []byte("{}")) {
			return nil
		}
		return json.Unmarshal(v, m)

	case string:
		if v == "[]" || v == "{}" {
			return nil
		}
		return json.Unmarshal([]byte(v), m)

	default:
		returnfmt.Errorf("Unsupported type: %T", v)
	}
}

func (m MyStruct) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m MyStruct) ConvertValue() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "[]", err
	}
	return string(b), nil
}
```