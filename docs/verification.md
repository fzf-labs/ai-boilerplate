# Verification

The repository CI gate lives in [.github/workflows/quality-gates.yml](../.github/workflows/quality-gates.yml).

## Commands

| Template | Core verification |
| --- | --- |
| Backend | `golangci-lint run --config .golangci.yml ./... -v`, `go test ./...` |
| Admin | `pnpm lint`, `pnpm check`, `pnpm build` |
| UniApp | `pnpm lint`, `pnpm check:type`, `pnpm build` |
| PC | `pnpm lint`, `pnpm test:unit`, `pnpm build` |
| Chrome extension | `npm ci`, `npm run build` |
| Electron | `npm install --no-audit --no-fund`, `npm run lint`, `npm run build` |
| Tauri | `cargo check --manifest-path src-tauri/Cargo.toml`, `npm install --no-audit --no-fund`, `npm run build` |
| Android | `./gradlew --no-daemon detekt test assembleDebug -x validateSigningDebug` |
| iOS | `mise install`, `mise exec -- bundle install`, `mise exec -- bundle exec arkana`, `mise exec -- tuist generate --no-open`, `mise exec -- bundle exec fastlane buildAndTestLane` |

## Backend Core Test Matrix

Run from `ai-boilerplate-backend`:

| Area | Coverage | Command |
| --- | --- | --- |
| auth/token | User/admin JWT claims, token expiry window, password hash verification | `go test ./internal/data` |
| RBAC | Menu permission collection, sorting, de-duplication, menu tree filtering | `go test ./internal/data` |
| payment order | Auth guard before repo access, WeChat/Alipay payment info response mapping, unsupported method rejection | `go test ./internal/service` |
| SMS verify code | Public auth whitelist, numeric code generation, reset-code TTL fixture | `go test ./internal/middleware/auth ./internal/data` |
| account logout/delete | Admin logout whitelist, app delete-account auth requirement, current no-op service reply contracts | `go test ./internal/middleware/auth ./internal/service` |
| secret masking | SMS channel `APISecret` response masking | `go test ./internal/service` |

Full backend regression:

```bash
cd ai-boilerplate-backend
go test ./...
```

## API Schema Test

The schema test target runs against an already running backend server. Defaults are `TEST_API_URL=http://localhost:8000`, `TEST_ADMIN_USER=admin`, `TEST_ADMIN_PASS=123456`, and `TEST_LOGIN_PATH=/admin/v1/sys_auth/login`; override them for local or CI fixtures instead of using production accounts.

```bash
cd ai-boilerplate-backend
make api-schema-test FILE=doc/swagger/admin/v1/sys_auth.swagger.json METHOD=GET
make api-schema-test FILE=doc/swagger/app/v1/home.swagger.json METHOD=GET NO_AUTH=true
```

To run every Swagger file:

```bash
cd ai-boilerplate-backend
make api-schema-test METHOD=GET
```
