> ## Datainfra YugabyteDB
>
> This is a library to enable generation of data infrastructure for YugabyteDB

# Datainfra YugabyteDB

YugabyteDB (YSQL) support for repository code generation, built on the SQL/gorm
implementation in [datainfra](https://github.com/activatedio/datainfra).

Because YSQL is PostgreSQL-compatible, this package is a thin adapter: code
generation, the runtime templates, migrations (goose), and setup all come from
`datainfra`'s gorm stack. This module adds:

* YugabyteDB-flavored connection configuration (multi-host, YSQL port 5433)
* Setup and migration config bridges
* An example generation target and test suite running against YugabyteDB

## Structure

* `pkg/data/yb` - runtime configuration mapping to the gorm data stack
* `pkg/setup/yb` - setup (database/user creation) config bridge
* `pkg/migrate/yb` - migrator config bridge
* `pkg/data/yb/testing` - test fixture helpers
* `examples` - generation example and integration tests

## Development

```
make dev_containers   # start YugabyteDB via docker compose
go generate ./...     # re-run code generation (examples/data/gen)
go test ./...
```
