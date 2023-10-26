# Migration

## Naming Migration

Migration up:

```plain
[version]_[name].up.sql
```

Rollback migration:

```plain
[version]_[name].down.sql
```

`version` - unsigned integer. Usually as a date and time in the format `YYYYmmddHHii`.
`name` - the name of the migration for convenience.
