# Generated Artifacts

This repository treats generated files as build output from a smaller set of
source artifacts. Update the source first, regenerate in the owning template,
then run the verification command for every affected consumer.

## Ownership

| Artifact | Source of Truth | Regenerate From |
| --- | --- | --- |
| Backend protobuf APIs | `ai-boilerplate-backend/api/**/*.proto` | `ai-boilerplate-backend` |
| Backend Go API bindings | Backend protobuf APIs | `make api` |
| Backend Swagger files | Backend protobuf APIs | `make api` |
| GORM model, DAO, repository code | SQL schema under `ai-boilerplate-backend/doc/sql` and database state | `make gorm` |
| Backend data and service scaffolding | Backend protobuf APIs and database tables | `make pbtocode` |
| Admin API clients | Backend Swagger under `ai-boilerplate-backend/doc/swagger/admin` | `pnpm api:gen` in `ai-boilerplate-admin` |
| Uni-app API clients | Backend Swagger under `ai-boilerplate-backend/doc/swagger/app` | `pnpm api:gen` in `ai-boilerplate-uniapp` |

Do not hand-edit generated `.pb.go`, Swagger, GORM model, DAO, repository, or
generated frontend API client files unless the generator itself is the change.

## Backend API Flow

Run these commands from `ai-boilerplate-backend`:

```bash
make sqltopb admin table_name
make api
make gorm DB_TABLES=table_name
make pbtocode DB_TABLES=table_name
go test ./...
```

Use `app` instead of `admin` in `make sqltopb` for app-facing APIs. If the
change spans both API surfaces, regenerate both surfaces before running tests.

## Frontend Client Flow

After backend Swagger changes, regenerate each frontend client that consumes the
changed Swagger output:

```bash
cd ai-boilerplate-admin
pnpm api:gen
pnpm check:type --filter=@vben/web-antd

cd ../ai-boilerplate-uniapp
pnpm api:gen
pnpm check:type
```

Commit the backend contract, generated backend output, generated frontend
clients, and verification evidence together so consumers do not see a partial
contract update.

## Review Checklist

- The source artifact that owns the change is included in the commit.
- Generated files were produced by the documented command, not manually patched.
- Affected frontend clients were regenerated when Swagger changed.
- Targeted verification passed for every changed template.
- `git status --short` contains only the intended source and generated files.
