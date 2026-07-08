# Verification Commands

The repository CI gate lives in
[.github/workflows/quality-gates.yml](../.github/workflows/quality-gates.yml).
Mark that workflow as a required branch-protection check when merges must be
blocked by verification failures.

Use the smallest command that covers the project you touched. Run from the
listed directory unless noted otherwise.

Install dependencies inside each template directory before running Node-based
checks. Backend checks require Go, Android checks require an Android SDK, and
iOS checks require Swift Package Manager/Xcode tooling.

## Fast Local Checks

| Project | Command |
| --- | --- |
| Backend | `go test ./...` |
| Admin | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | `pnpm check:type` |
| PC web | `pnpm type-check` |
| Electron | `npm run typecheck` |
| Tauri web shell | `npm run build` |
| Chrome extension | `npm run type-check && npm test` |
| Android | `./gradlew detekt test assembleDebug -x validateSigningDebug` |
| iOS | `swift build` |

For iOS, `swift build` is the reusable static check for the Swift package
workspace. Changes to generated Xcode targets, app assets, signing,
entitlements, or simulator/device behavior also need `tuist generate` and an
Xcode or Tuist build for the affected target.

When a command cannot run because a local SDK, package manager, signing asset,
or generated workspace is missing, keep the exact command output in the task
summary and run the targeted static checks that cover the edited files.

## CI Gate Commands

These commands mirror the quality-gates workflow. Keep this table aligned with
the workflow when changing required checks.

| Template | Core verification |
| --- | --- |
| Backend | `go test ./...` |
| Admin | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | `pnpm check:type` |
| PC web | `pnpm type-check` |
| Chrome extension | `npm run type-check`, `npm test` |
| Electron | `npm run typecheck` |
| Tauri | `npm run build` |
| Android | `./gradlew --no-daemon detekt test assembleDebug -x validateSigningDebug` |
| iOS | `swift build` |

Stricter lint, production build, Tuist/Fastlane, and platform packaging checks
are release checks. Run them before publishing when the touched template and
local prerequisites require them.

## Backend Core Test Matrix

Backend P0 changes should keep these flows covered by automated tests before
running broader integration checks.

| Flow | Coverage target | Local command |
| --- | --- | --- |
| Auth/token | Admin and app user token generation/parsing, token claim shape | `go test ./internal/data` |
| RBAC | Auth whitelist matcher, menu tree filtering, permission de-duplication | `go test ./internal/middleware/auth ./internal/data` |
| Payment orders | Payment callback input normalization, order status transitions, payment record mapping | `go test ./internal/service` |
| SMS verification codes | Scene config, code generation, send callback behavior | `go test ./internal/data` |
| Account deletion | Deleted-user reference and anonymized retained user fields | `go test ./internal/service ./internal/data` |
| Secret masking | SMS channel API key/secret response masking | `go test ./internal/service` |
| API schema smoke | One admin Swagger and one app Swagger against a running server | `make api-schema-test` |
| API schema full | All Swagger files against a running server | `./scripts/api-schema-test.sh --all -m ALL` |

Run backend commands from `ai-boilerplate-backend`. The Go tests use local
fixtures under `internal/testfixture` and must not depend on production
accounts. Schema tests require the HTTP server to be running and use
`TEST_API_URL`, `TEST_ADMIN_USER`, `TEST_ADMIN_PASS`, and `TEST_LOGIN_PATH` for
non-default environments.

## Documentation-Only Changes

For root documentation edits that do not change code or generated artifacts,
run:

```bash
git diff --check
```

If the edit changes a documented command, also run that command from the listed
directory or record the missing local prerequisite. Documentation that changes
release, generated-artifact, or cross-template workflow guidance should still
be reviewed against the full verification table before publishing.

## Full Verification Set

For cross-template changes, run every command in the table when the local SDKs
are available. Record both successful checks and checks skipped because of a
missing local prerequisite.

## Targeted Verification

For a scoped change, run the command for the touched template first. Examples:

- Admin-only view or API client change: `pnpm check:type --filter=@vben/web-antd`
- Uni-app page or config change: `pnpm check:type`
- Backend service or generated API change: `go test ./...`
- Chrome extension runtime message or popup change: `npm run type-check && npm test`
- Desktop shell copy/config change: the corresponding `npm run typecheck` or
  `npm run build`
