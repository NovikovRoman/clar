package mysql

const RepositoryUtils = `package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"{{.Module}}/{{.DBType}}/entity"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

const (
	tagName = "db"
	tagPrimaryKey = "pkey"
)

// save saves the record to the database.
//
// New record - creates in a DB, existing - updates in a DB. If the entry:
//
// - entity.SimpleBaseEntity regular update of a record in the database,
//
// - entity.BaseEntity regular update of the record in the database
//     and the auto-update of the date in the UpdatedAt field.
func save(ctx context.Context, db *sqlx.DB, table string, ent entity.SimpleBaseEntity) (err error) {
	if ent.GetID() == 0 {
		return create(ctx, db, table, ent)
	}

	if v, ok := ent.(entity.BaseEntity); ok {
		setUpdatedAt(v, time.Now())
	}

	query := "UPDATE {{.Backtick}}" + table + "{{.Backtick}} SET " + fieldsForUpdate(ent) + " WHERE id=:id"
	_, err = db.NamedExecContext(ctx, query, ent)
	return
}

// saveMultiple saves multiple entries to the database. Adds new, updates existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func saveMultiple(ctx context.Context, db *sqlx.DB, table string, ents ...any) (err error) {
	return saveMultipleBase(ctx, db, table, false, ents...)
}

// saveMultipleIgnoreDuplicates saves multiple entries to the database. Adds new, ignore existing entities.
// Entities must be of the same type.
// [!] Use with caution.
// - for new entries, does not return an ID.
// - be sure to specify primaryKey (pkey) if present.
// Example: ID int64 {{.Backtick}}db:"id" pkey:"true"{{.Backtick}}
func saveMultipleIgnoreDuplicates(ctx context.Context, db *sqlx.DB, table string, ents ...any) (err error) {
	return saveMultipleBase(ctx, db, table, true, ents...)
}

func saveMultipleBase(ctx context.Context, db *sqlx.DB, table string, ignore bool, ents ...any) (err error) {
	if len(ents) == 0 {
		return
	}

	query, fields, args := partQueryMultiInsert(table, ignore, ents...)
	if !ignore {
		primaryKey := getPrimaryKey(ents[0])
		comma := false
		for _, field := range fields {
			if field == primaryKey {
				continue
			}

			if comma {
				query += ","
			} else {
				comma = true
			}
			query += fmt.Sprintf("{{.Backtick}}%s{{.Backtick}}=VALUES({{.Backtick}}%s{{.Backtick}})", field, field)
		}
	}

	_, err = db.ExecContext(ctx, query, args...)
	return
}

// partQueryMultiInsert inserts multiple records into the database.
// There is a space at the end of the line.
//
// Example:
// INSERT [IGNORE] INTO table({{.Backtick}}id{{.Backtick}},{{.Backtick}}field1{{.Backtick}},{{.Backtick}}field2{{.Backtick}}) VALUES (0,'str','str2'),(10,'str1','str3'),(2,'str4','str5') as t
// [ON DUPLICATE KEY UPDATE]
func partQueryMultiInsert(table string, ignore bool, ents ...any) (query string, fields []string, args []any) {
	fields = tableFields(ents[0])

	for _, ent := range ents {
		if v, ok := ent.(entity.BaseEntity); ok {
			setUpdatedAt(v, time.Now())

			if v.GetID() == 0 {
				setCreatedAt(v, time.Now())
				if v.GetUpdatedAt().IsZero() {
					setUpdatedAt(v, time.Now())
				}
				if v.GetDeletedAt() != nil && v.GetDeletedAt().IsZero() {
					setDeletedAt(v, nil)
				}
			}
		}
	}

	/*
	   INSERT INTO table({{.Backtick}}id{{.Backtick}},{{.Backtick}}field1{{.Backtick}},{{.Backtick}}field2{{.Backtick}})
	   VALUES (0,'str','str2'),(10,'str1','str3'),(2,'str4','str5') as t
	   ON DUPLICATE KEY UPDATE id=t.id,t.field1=t.field1,field2=t.field2;
	*/
	args = []any{}
	values := ""
	m := reflectx.NewMapper(tagName)
	for _, ent := range ents {
		for _, field := range fields {
			v := m.FieldByName(reflect.ValueOf(ent), field)
			if v.Type().String() == "*time.Time" && v.Interface().(*time.Time).IsZero() {
				args = append(args, nil)

			} else {
				args = append(args, v.Interface())
			}
		}

		if values != "" {
			values += ","
		}
		values += "(?" + strings.Repeat(",?", len(fields)-1) + ")"
	}

	query = "INSERT "
	if ignore {
		query += "IGNORE "
	}
	query += "INTO {{.Backtick}}" + table + "{{.Backtick}}" + "({{.Backtick}}" + strings.Join(fields, "{{.Backtick}},{{.Backtick}}") + "{{.Backtick}})" + " VALUES " +
		values + " "
	if !ignore {
		query += "ON DUPLICATE KEY UPDATE "
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
//     setting CreatedAt to the current time, UpdatedAt if not set - to the current time,
//     DeletedAt if not set - nil.
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

	set, values := fieldsForInsert(ent)

	var res sql.Result
	query := "INSERT INTO {{.Backtick}}" + table + "{{.Backtick}} (" + set + ") VALUES (" + values + ")"
	res, err = db.NamedExecContext(ctx, query, ent)

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

    if _, ok := ent.(entity.BaseEntity); ok {
        query := "UPDATE {{.Backtick}}" + table + "{{.Backtick}} SET {{.Backtick}}deleted_at{{.Backtick}}=? WHERE {{.Backtick}}id{{.Backtick}}=?"
        _, err = db.ExecContext(ctx, query, time.Now(), ent.GetID())
        return
    }

    query := "DELETE FROM {{.Backtick}}" + table + "{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}}=?"
    _, err = db.ExecContext(ctx, query, ent.GetID())
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

func fieldsForInsert(ent any) (set string, values string) {
	for i, name := range tableFields(ent) {
		if i > 0 {
			set += ","
			values += ","
		}
		set += "{{.Backtick}}" + name + "{{.Backtick}}"
		values += ":" + name
	}
	return
}

func fieldsForUpdate(ent any) (set string) {
	for i, name := range tableFields(ent) {
		if i > 0 {
			set += ","
		}
		set += "{{.Backtick}}" + name + "{{.Backtick}}=:" + name
	}
	return
}

func tableFields(ent any) (fields []string) {
	m := reflectx.NewMapperFunc(tagName, func(s string) string { return s })
	for _, n := range m.TypeMap(reflect.TypeOf(ent)).Names {
		if n.Field.Tag.Get("readonly") == "true" {
			continue
		}
		if len(n.Index) == 1 {
			fields = append(fields, n.Path)
		}
	}
	return
}

func getPrimaryKey(ent any) string {
	m := reflectx.NewMapperFunc(tagName, func(s string) string { return s })
	for k, n := range m.TypeMap(reflect.TypeOf(ent)).Names {
		if n.Field.Tag.Get(tagPrimaryKey) != "" {
			return k
		}
	}
	return ""
}
`
