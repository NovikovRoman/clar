# CLean ARchitecture

> This is a Go application for creating a clean architecture of the application being developed.

**⚠ Currently only MySQL is supported**

## Table of Contents

- [CLean ARchitecture](#clean-architecture)
- [Table of Contents](#table-of-contents)
  - [Build](#build)
  - [Usage](#usage)
    - [Initialization](#initialization)
    - [Entity And Repository](#entity-and-repository)
    - [Array Structure For JSON Columns](#array-structure-for-json-columns)
    - [Structure For JSON Columns](#structure-for-json-columns)
    - [Migration Tools](#migration-tools)
  - [Help](#help)
  - [License](#license)

### Build

```shell script
make && sudo mv bin/clar /usr/local/bin/
```

### Usage

#### Initialization

```shell script
clar i [-d dbtype|--db=dbtype]
```

or

```shell script
clar init [-d dbtype|--db=dbtype]
```

- `dbtype` - for what type of database (default: `mysql`).

Will be created:

- [internal/db/mysql/entity/base_entity.go](docs/base_entity.md)

- For MySQL: [internal/db/mysql/repository/mysql/utils.go](docs/utils.md)

#### Entity And Repository

Creating a entity with repository.

```shell script
clar e name
```

or

```shell script
clar entity name [-d dbtype|--db=dbtype] [-s|--simple] or [-e|--empty]
```

- `name` - entity name (required).
- `dbtype` - for what type of database (default: `mysql`).
- `-s`, `--simple` - simple entity.
- `-e`, `--empty` - empty entity.

Will be created (example `clar e user`):

- [`internal/db/mysql/entity/user.go`](docs/user.md) or if the `-s` flag is specified [`user.go`](docs/simple_user.md) or if the `-e` flag is specified [`user.go`](docs/empty_user.md)
- For MySQL: [internal/db/mysql/repository/user.go](docs/user_repository.md) or if the `-e` flag is specified [user.go](docs/empty_user_repository.md)

if there was no initialization, there will be automatic initialization for the selected database.

#### Array Structure For JSON Columns

Creating an array structure for a json column.

```shell script
clar a name
```

or

```shell script
clar array name
```

- `name` - struct name (required).

Will be created (example `clar i myArr`) [`internal/db/mysql/entity/my_arr.go`](docs/my_arr.md).

#### Structure For JSON Columns

Creating a structure for a json column.

```shell script
clar s name
```

or

```shell script
clar struct name
```

- `name` - struct name (required).

Will be created (example `clar s myStruct`) [`internal/db/mysql/entity/my_struct.go`](docs/my_struct.md).

#### Migration Tools

Creating migration tools code.

```shell script
clar m [-d dbtype|--db=dbtype]
```

or

```shell script
clar migrate [-d dbtype|--db=dbtype]
```

- `dbtype` - for what type of database (default: `mysql`).

Will be created (for MySQL):

- [migrations/mysql/migrate.go](docs/migrate_mysql.md)

- [migrations/mysql/migrations/….[up|down].sql](docs/migrations_mysql.md)

### Help

Any of the options:

```shell script
clar -h|--help
```

### License

[MIT License](LICENSE) © Roman Novikov
