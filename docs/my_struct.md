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

func (m *MyStruct) Scan(val interface{}) (err error) {
    var value MyStruct
    switch v := val.(type) {
    case []byte:
        if bytes.Equal(v, []byte("[]")) || bytes.Equal(v, []byte("{}")) {
            return nil
        }
        err = json.Unmarshal(v, &value)

    case string:
        if v == "[]" || v == "{}" {
            return nil
        }
        err = json.Unmarshal([]byte(v), &value)

    default:
        err = fmt.Errorf("Unsupported type: %T", v)
    }

    if err == nil {
        *m = value
    }
    return
}

func (m MyStruct) Value() (driver.Value, error) {
    return json.Marshal(m)
}

func (m MyStruct) ConvertValue() (string, error) {
    b, err := json.Marshal(m)
    if err != nil {
        return "{}", err
    }
    return string(b), nil
}
```
