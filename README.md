# CLean ARchitecture

> This is a Go application for creating a clean architecture of the application being developed.

**⚠ Currently only MySQL/MariaDB and PostgreSQL is supported**

## Table of Contents

- [CLean ARchitecture](#clean-architecture)
- [Table of Contents](#table-of-contents)
  - [Build](#build)
  - [Usage](#usage)
    - [Entity And Repository](#entity-and-repository)
    - [Array Structure For JSON Columns](#array-structure-for-json-columns)
    - [Structure For JSON Columns](#structure-for-json-columns)
  - [Help](#help)
  - [License](#license)

### Build

```shell script
make && sudo mv bin/clar /usr/local/bin/
```

### Usage

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
- `dbtype` - for what type of database (default: `pg` postgreSQL).
- `-s`, `--simple` - simple entity (ID only).
- `-e`, `--empty` - empty entity (Without fields).

if migration tools were missing, they will be created.

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

### Help

Any of the options:

```shell script
clar -h|--help
```

### License

[MIT License](LICENSE) © Roman Novikov
