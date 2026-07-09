# AI Boilerplate Backend

Go backend template for AI Boilerplate. It uses Kratos, protobuf-first API
definitions, generated HTTP/gRPC bindings, PostgreSQL/GORM, Redis, and generated
Swagger output for frontend clients.

## Setup

Install Go 1.24 or newer, then run commands from this directory.

```bash
go mod download
make migrate-install
make migrate-up
go test ./...
```

The development config lives in `configs/config.development.yaml`. Copy
`configs/config.example.yaml` when creating environment-specific config and set
the database, Redis, JWT, WeChat, and push-service values for that environment.
The initial seed creates a local `admin` / `123456` administrator for bootstrap
only; change it before sharing any non-local environment.

## Local Development Path

1. Run `go mod download`.
2. Configure local PostgreSQL and Redis in `configs/config.development.yaml`.
3. Run `make migrate-install` once, then `make migrate-up` against the local
   PostgreSQL database.
4. Start the service with `make run` when integration testing is needed.
5. Run `go test ./...` before committing backend changes.
6. If protobuf, Swagger, SQL, migration, or GORM artifacts change, follow the generated
   artifact flow below and regenerate consumers in the same commit.

## Change Checklist

- Data model change: update the SQL source, regenerate GORM artifacts, then run
  `go test ./...`.
- Migration change: update migration files and SQL snapshots together, verify
  `make migrate-up`, `make migrate-status`, and at least one `make migrate-down`.
- API shape change: edit protobuf source, run `make api`, then regenerate any
  admin or uni-app clients that consume the changed Swagger output.
- Service logic change: keep the edit inside `internal/service` or the owning
  data/repository package, then run `go test ./...`.
- External-account or production-credential change: stop and confirm the target
  environment before editing checked-in config.

## Ownership Notes

Start in this template when the task changes database tables, protobuf
messages, HTTP/gRPC routes, validation, auth, permissions, jobs, or shared
response semantics. After backend generation, regenerate admin or uni-app API
clients only when the changed Swagger is consumed by those clients.

Do not use backend generation for client-only copy, page layout, local route
labels, or placeholder UI defaults. Those changes belong in the template that
users see.

## Common Commands

```bash
make run       # start the Kratos server with APP_ENV=development
make build     # build cmd/server into bin/
make migrate-up     # apply database migrations
make migrate-down   # roll back the latest migration
make migrate-status # show database migration status
make migrate-create NAME=add_table # create a new SQL migration
make api       # regenerate API, HTTP, gRPC, validation, errors, and Swagger
make gorm      # regenerate GORM model/DAO/repository code
make pbtocode  # regenerate data and service code from protobuf definitions
go test ./...  # verify backend packages
```

## Database Migrations

Database migrations live in `db/migrations` and are run with
[`pressly/goose`](https://github.com/pressly/goose). The Makefile defaults use
the local PostgreSQL settings at the top of `Makefile`; override `DB_HOST`,
`DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, or `MIGRATION_DSN` for another
environment, either as environment variables or Make variables.

```bash
make migrate-install
make migrate-up
make migrate-status
make migrate-down
make migrate-create NAME=add_example_table
```

`migrate-up` initializes an empty database with the current schema and base
seed data. `migrate-down` rolls back one migration, so from a fresh database it
first removes the seed data and leaves the schema in place.

## Generated Artifacts

API definitions start in `api/admin/v1` and `api/app/v1`. Generated Swagger
files are written to `doc/swagger`, SQL snapshots live in `doc/sql`, migration
files live in `db/migrations`, and generated GORM code lives under
`internal/data/gorm`.

When database tables change, create or update a migration first, keep the
matching SQL snapshot under `doc/sql` current, apply the migration to a local
database, then run the matching generation commands:

```bash
make migrate-create NAME=add_table_name
make migrate-up
make sqltopb admin table_name
make api
make gorm DB_TABLES=table_name
make pbtocode DB_TABLES=table_name
```

Do not generate GORM/protobuf/client artifacts from an unapplied schema change.
Do not hand-edit generated `.pb.go`, Swagger, GORM model, DAO, or repository
files unless the generator itself is being fixed.

## API Schema Tests

The Makefile includes a schema-test entry point for generated Swagger files:

```bash
make api-schema-test
make api-schema-test FILE=doc/swagger/admin/v1/user.swagger.json METHOD=GET
./scripts/api-schema-test.sh --all -m ALL
```

Set `TEST_API_URL`, `TEST_ADMIN_USER`, `TEST_ADMIN_PASS`, and `TEST_LOGIN_PATH`
when testing against a non-default local server.

The default `make api-schema-test` command runs a smoke matrix covering one
admin Swagger file and one app Swagger file against a running HTTP server. Use
`--all` for exhaustive contract coverage. If the server depends on external
accounts or production credentials, do not point the tests at that environment
without explicit approval.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Generated artifact flow: `../docs/generated-artifacts.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
