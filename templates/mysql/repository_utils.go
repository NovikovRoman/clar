package mysql

const RepositoryUtils = `package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"{{.Module}}/domain/entity"
	"github.com/jmoiron/sqlx"
)

// save saves the record to the database.
//
// New record - creates in a DB, existing - updates in a DB. If the entry:
//
// - entity.SimpleEntity regular update of a record in the database,
//
// - entity.Entity regular update of the record in the database
// 		and the auto-update of the date in the UpdatedAt field.
func save(ctx context.Context, db *sqlx.DB, table string, ent entity.SimpleBaseEntity) (err error) {
	if ent.GetID() == 0 {
		return create(ctx, db, table, ent)
	}

	if v, ok := ent.(entity.BaseEntity); ok {
		setUpdatedAt(v, time.Now())
	}

	var set string
	if set, err = fieldsForUpdate(ent); err != nil {
		return
	}

	query := "UPDATE {{.Backtick}}" + table + "{{.Backtick}} SET " + set + " WHERE id=:id"
	if ctx == nil {
		_, err = db.NamedExec(query, ent)
	} else {
		_, err = db.NamedExecContext(ctx, query, ent)
	}
	return
}

// create creates a record in the database.
//
// If the entry:
//
// - entity.SimpleBaseEntity the usual creation of a record in the database,
//
// - entity.BaseEntity the usual creation of a record in the database,
//		setting CreatedAt to the current time, UpdatedAt if not set - to the current time,
//  	DeletedAt if not set - nil.
func create(ctx context.Context, db *sqlx.DB, table string, ent entity.SimpleBaseEntity) (err error) {
	if ent.GetID() > 0 {
		return save(ctx, db, table, ent)
	}

	if vEnt, ok := ent.(entity.BaseEntity); ok {
		setCreatedAt(vEnt, time.Now())
		if vEnt.GetUpdatedAt().IsZero() {
			setUpdatedAt(vEnt, time.Now())
		}
		if vEnt.GetDeletedAt() != nil && vEnt.GetDeletedAt().IsZero() {
			setDeletedAt(vEnt, nil)
		}
	}

	var (
		set    string
		values string
	)
	if set, values, err = fieldsForInsert(ent); err != nil {
		return
	}

	var res sql.Result
	query := "INSERT INTO {{.Backtick}}" + table + "{{.Backtick}} (" + set + ") VALUES (" + values + ")"
	if ctx == nil {
		res, err = db.NamedExec(query, ent)
	} else {
		res, err = db.NamedExecContext(ctx, query, ent)
	}

	if err == nil {
		var id int64
		id, err = res.LastInsertId()
		setID(ent, id)
	}
	return
}

// update updates a record in the database. Alias save.
func update(ctx context.Context, db *sqlx.DB, table string, ent entity.SimpleBaseEntity) (err error) {
	if ent.GetID() == 0 {
		err = errors.New("This is a new entry. ")
		return
	}

	return save(ctx, db, table, ent)
}

// remove deleting an entry from the database or marking it as a deleted entry in cases where entity:
//
// - entity.SimpleBaseEntity - removes a record from the database,
//
// - entity.BaseEntity - marks the entry in the database as deleted.
func remove(ctx context.Context, db *sqlx.DB, table string, ent entity.SimpleBaseEntity) (err error) {
	if ent.GetID() == 0 {
		return
	}

	var sql string
	if _, ok := ent.(entity.BaseEntity); ok {
		sql = "UPDATE {{.Backtick}}" + table + "{{.Backtick}} SET {{.Backtick}}deleted_at{{.Backtick}}=null WHERE {{.Backtick}}id{{.Backtick}}=?"
	} else {
		sql = "DELETE FROM {{.Backtick}}" + table + "{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}}=?"
	}

	if ctx == nil {
		_, err = db.Exec(sql, ent.GetID())
	} else {
		_, err = db.ExecContext(ctx, sql, ent.GetID())
	}
	return
}

func setID(ent entity.SimpleBaseEntity, id int64) {
	v := reflect.ValueOf(ent)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("ID").SetInt(id)
}

func setCreatedAt(ent entity.BaseEntity, t time.Time) {
	v := reflect.ValueOf(ent)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("CreatedAt").Set(reflect.ValueOf(t))
}

func setUpdatedAt(ent entity.BaseEntity, t time.Time) {
	v := reflect.ValueOf(ent)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(t))
}

func setDeletedAt(ent entity.BaseEntity, t *time.Time) {
	v := reflect.ValueOf(ent)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if t != nil {
		v.FieldByName("DeletedAt").SetPointer(unsafe.Pointer(t))
		return
	}
	v.FieldByName("DeletedAt").SetPointer(nil)
}

func fieldsForInsert(ent interface{}) (set string, values string, err error) {
	var fields []string
	if fields, err = tableFields(ent); err != nil {
		return
	}

	for i, name := range fields {
		if i > 0 {
			set += ","
			values += ","
		}
		set += "{{.Backtick}}" + name + "{{.Backtick}}"
		values += ":" + name
	}
	return
}

func fieldsForUpdate(ent interface{}) (set string, err error) {
	var fields []string
	if fields, err = tableFields(ent); err != nil {
		return
	}

	for i, name := range fields {
		if i > 0 {
			set += ","
		}
		set += "{{.Backtick}}" + name + "{{.Backtick}}=:" + name
	}
	return
}

func tableFields(ent interface{}) (fields []string, err error) {
	v := reflect.ValueOf(ent)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fields = []string{}
	switch {
	case v.Kind() == reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field == "-" {
				continue

			} else if field == "" {
				fields = append(fields, strings.ToLower(v.Type().Field(i).Name))
				continue
			}

			fields = append(fields, field)
		}
		return

	case v.Kind() == reflect.Map:
		fields = make([]string, len(v.MapKeys()))
		for i, k := range v.MapKeys() {
			fields[i] = k.String()
		}
		return
	}

	err = fmt.Errorf("dbTableFields requires a struct or a map, found: %s", v.Kind().String())
	return
}
`
