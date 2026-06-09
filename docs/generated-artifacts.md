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

## Contract Operation Checklist

Use this order for backend contract work:

1. Update the source artifact: SQL schema for stored data, protobuf for API
   shape and validation, or generator code when the tool itself is changing.
2. Regenerate backend output with the Makefile target that matches the source
   change.
3. Regenerate admin and/or uni-app clients only when the Swagger output they
   consume changed.
4. Run backend verification before frontend verification so contract failures
   are caught at the owner first.
5. Commit source and generated output together.

Acceptance for a contract change means the source file, generated backend
output, generated client output when required, and all affected verification
commands are included in the same task result.

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

## Frontend Regeneration Triggers

Regenerate a frontend client when any of these source artifacts changed:

| Consumer | Trigger | Command | Verification |
| --- | --- | --- | --- |
| Admin | Files under `ai-boilerplate-backend/doc/swagger/admin` changed. | `pnpm api:gen` from `ai-boilerplate-admin` | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | Files under `ai-boilerplate-backend/doc/swagger/app` changed. | `pnpm api:gen` from `ai-boilerplate-uniapp` | `pnpm check:type` |

Do not regenerate clients for UI-only changes, route-label copy, local env
defaults, or documentation updates. If generated client output changes without a
matching Swagger change, inspect the generator version and local install before
committing the diff.

When a task does not require regeneration, record that as evidence instead of
running a generator just to prove a no-op. The usual evidence is: the source
Swagger or protobuf path did not change, the edited files are not generated
client output, and the owning template verification command still passes.

These changes are explicit non-triggers for generated output:

| Change | Regeneration Needed |
| --- | --- |
| Root docs, template README, or release checklist edits | No. Run `git diff --check` and any changed documented command. |
| Client-only copy, labels, routes, styles, or empty-state defaults | No. Run the owning client type check. |
| Local endpoint examples or environment documentation | No, unless runtime config files also changed. |
| Package metadata that does not affect generated API imports | No. Run the affected package check. |

## Generated Drift Triage

When generated files change unexpectedly:

1. Check whether the source artifact changed in the same diff.
2. Confirm the generator command was run from the documented template directory.
3. Confirm dependencies were installed with the template's package manager.
4. Revert unrelated timestamp, formatting, or dependency churn unless the
   generator intentionally owns it.
5. If drift remains unexplained, stop and record the generator version and exact
   command before committing.

## Review Checklist

- The source artifact that owns the change is included in the commit.
- Generated files were produced by the documented command, not manually patched.
- Affected frontend clients were regenerated when Swagger changed.
- Targeted verification passed for every changed template.
- `git status --short` contains only the intended source and generated files.
