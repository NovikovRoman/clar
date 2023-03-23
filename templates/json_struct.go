package templates

const JsonStruct = `package entity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type {{.Struct}} struct {
	Field1 int
	Field2 string
	Field3 bool
}

func ({{.StructSymb}} *{{.Struct}}) String() string {
	b, _ := json.Marshal({{.StructSymb}})
	return string(b)
}

func ({{.StructSymb}} *{{.Struct}}) Scan(val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		if bytes.Equal(v, []byte("[]")) {
			return
		}
		err = json.Unmarshal(v, {{.StructSymb}})
		return

	case string:
		if v == "[]" {
			return
		}
		err = json.Unmarshal([]byte(v), {{.StructSymb}})
		return

	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
}

func ({{.StructSymb}} *{{.Struct}}) Value() (driver.Value, error) {
	return json.Marshal({{.StructSymb}})
}

func ({{.StructSymb}} *{{.Struct}}) ConvertValue() (string, error) {
	b, err := json.Marshal({{.StructSymb}})
	if err != nil {
		return "[]", err
	}
	return string(b), nil
}
`
