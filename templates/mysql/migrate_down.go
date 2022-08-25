package mysql

const MigrateDown = `DROP TABLE IF EXISTS {{.Backtick}}table{{.Backtick}};`
