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

func ({{.StructSymb}} {{.Struct}}) String() string {
	b, _ := json.Marshal({{.StructSymb}})
	return string(b)
}

func ({{.StructSymb}} *{{.Struct}}) Scan(val interface{}) error {
	var value {{.Struct}}
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
		*{{.StructSymb}} = value
	}
	return
}

func ({{.StructSymb}} {{.Struct}}) Value() (driver.Value, error) {
	return json.Marshal({{.StructSymb}})
}

func ({{.StructSymb}} {{.Struct}}) ConvertValue() (string, error) {
	b, err := json.Marshal({{.StructSymb}})
	if err != nil {
		return "{}", err
	}
	return string(b), nil
}
`
