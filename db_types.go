package main

import (
	"errors"
	"strings"
)

const (
	MysqlCode = iota
)

type DBType struct {
	code int
	name string
}

func newDBType(s string) (t *DBType, err error) {
	s = strings.ToLower(s)
	switch s {
	case "mysql":
		t = &DBType{
			code: MysqlCode,
			name: s,
		}

	default:
		err = errors.New("Database type not supported.")
	}
	return
}
