# Database Migration Tool

## Decision

Use `github.com/pressly/goose/v3/cmd/goose` for backend SQL migrations.

## Context

The backend already keeps PostgreSQL SQL snapshots under `doc/sql` and generates
GORM/protobuf/client artifacts from database state, but it did not have a
versioned migration and rollback entry point for empty-environment bootstrap.

## Options

| Option | Fit | Notes |
| --- | --- | --- |
| `pressly/goose` | High | Provides `up`, `down`, `status`, and `create`; supports plain SQL files and the existing PostgreSQL DSN format. |
| `golang-migrate/migrate` | Medium | Mature and widely used, but the CLI does not provide a native `status` command matching this repository's acceptance criteria. |
| Custom scripts | Low | Could reuse snapshots, but would create project-specific migration behavior and higher maintenance risk. |

## Consequences

- Migration files live in `db/migrations`.
- `make migrate-up`, `make migrate-down`, `make migrate-status`, and
  `make migrate-create NAME=...` are the standard local and CI entry points.
- Database schema changes must be applied through migrations before running
  GORM/protobuf/client generation.
- SQL snapshots under `doc/sql` remain generator input documentation and must
  stay aligned with applied migrations.

## Rollback

If goose becomes unsuitable, the SQL migration files can be adapted to another
tool because they use ordinary PostgreSQL statements and explicit up/down
sections.
