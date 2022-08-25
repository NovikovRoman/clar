# CLean ARchitecture
> This is a Go application for creating a clean architecture of the application being developed.

**⚠ Currently only MySQL is supported**

# Table of Contents
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

## Build

```shell script
make && sudo mv bin/clar /usr/local/bin/
```

## Usage

### Initialization

```shell script
clar i [-d dbtype|--db=dbtype]
```
or
```shell script
clar init [-d dbtype|--db=dbtype]
```
- `dbtype` - for what type of database (default: `mysql`).

Will be created:

- [domain/entity/entity_interface.go](docs/entity_interface.md)

- For MySQL: [domain/repository/mysql/utils.go](docs/utils.md)

### Entity And Repository

Creating a entity with repository.

```shell script
clar e [-n name|--name=name]
```
or
```shell script
clar entity -n name|--name=name [-d dbtype|--db=dbtype] [-s|--simple]
```
- `name` - entity name (required).
- `dbtype` - for what type of database (default: `mysql`).
- `-s`, `--simple` - simple entity.

Will be created (example `clar i -nuser`):

- [`domain/repository/user_repository_interface.go`](docs/user_repository_interface.md)
- [`domain/repository/entity/user.go`](docs/user.md) or if the `-s` flag is specified [`domain/repository/entity/user.go`](docs/simple_user.md)
- For MySQL: [domain/repository/mysql/user_repository.go](docs/user_repository.md)

if there was no initialization, there will be automatic initialization for the selected database.

### Array Structure For JSON Columns

Creating an array structure for a json column.

```shell script
clar a -n name|--name=name
```
or
```shell script
clar array -n name|--name=name
```
- `name` - struct name (required).

Will be created (example `clar i -nmyArr`) [`domain/entity/my_arr.go`](docs/my_arr.md).

### Structure For JSON Columns

Creating a structure for a json column.

```shell script
clar s -n name|--name=name
```
or
```shell script
clar struct -n name|--name=name
```
- `name` - struct name (required).

Will be created (example `clar i -nmyStruct`) [`domain/entity/my_struct.go`](docs/my_struct.md).

### Migration Tools

Creating migration tools code.

```shell script
clar m [-p dirpath|--path=dirpath] [-d dbtype|--db=dbtype]
```
or
```shell script
clar migrate [-p dirpath|--path=dirpath] [-d dbtype|--db=dbtype]
```
- `dirpath` - for what type of database (default: `cmd/migrate`).
- `dbtype` - for what type of database (default: `mysql`).

This will create the directory `dirpath/dbtype`, with a migration tools and [documentation](docs/migrate_readme.md).

## Help

Any of the options:

```shell script
clar -h|--help
```

## License

[MIT License](LICENSE) © Roman Novikov