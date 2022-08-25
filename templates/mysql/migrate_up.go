package mysql

const MigrateUp = `CREATE TABLE IF NOT EXISTS {{.Backtick}}table{{.Backtick}}
(
    {{.Backtick}}id{{.Backtick}}         bigint(20)                              NOT NULL AUTO_INCREMENT PRIMARY KEY,
    {{.Backtick}}host{{.Backtick}}       varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL UNIQUE,
    {{.Backtick}}success{{.Backtick}}    boolean                                 NOT NULL,

    {{.Backtick}}created_at{{.Backtick}} datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    {{.Backtick}}updated_at{{.Backtick}} datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    {{.Backtick}}deleted_at{{.Backtick}} datetime                                         DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;`
