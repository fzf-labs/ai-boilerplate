# AI Boilerplate Backend

Go backend template for AI Boilerplate. It uses Kratos, protobuf-first API
definitions, generated HTTP/gRPC bindings, PostgreSQL/GORM, Redis, and generated
Swagger output for frontend clients.

## Setup

Install Go 1.24 or newer, then run commands from this directory.

```bash
go mod download
go test ./...
```

The development config lives in `configs/config.development.yaml`. Copy
`configs/config.example.yaml` when creating environment-specific config and set
the database, Redis, JWT, WeChat, and push-service values for that environment.

## Common Commands

```bash
make run       # start the Kratos server with APP_ENV=development
make build     # build cmd/server into bin/
make api       # regenerate API, HTTP, gRPC, validation, errors, and Swagger
make gorm      # regenerate GORM model/DAO/repository code
make pbtocode  # regenerate data and service code from protobuf definitions
go test ./...  # verify backend packages
```

## Generated Artifacts

API definitions start in `api/admin/v1` and `api/app/v1`. Generated Swagger
files are written to `doc/swagger`, SQL snapshots live in `doc/sql`, and
generated GORM code lives under `internal/data/gorm`.

When database tables change, update the SQL source first, then run the matching
generation commands:

```bash
make sqltopb admin table_name
make api
make gorm DB_TABLES=table_name
make pbtocode DB_TABLES=table_name
```

Do not hand-edit generated `.pb.go`, Swagger, GORM model, DAO, or repository
files unless the generator itself is being fixed.

## API Schema Tests

The Makefile includes a schema-test entry point for generated Swagger files:

```bash
make api-schema-test
make api-schema-test FILE=admin/v1/user.swagger.json METHOD=GET
```

Set `TEST_API_URL`, `TEST_ADMIN_USER`, `TEST_ADMIN_PASS`, and `TEST_LOGIN_PATH`
when testing against a non-default local server.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Generated artifact flow: `../docs/generated-artifacts.md`
- Environment configuration: `../docs/environment.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
