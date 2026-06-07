# Operations Guide

This repository is a multi-template workspace. Use the command that matches the
project you changed, then run the full verification set before publishing a
cross-template change.

Most work should stay inside one template directory. Cross-template changes are
reserved for shared API contracts, generated client syncs, documentation, or
release workflow updates.

## Repository Layout

| Path | Purpose |
| --- | --- |
| `ai-boilerplate-backend` | Go backend service and generated API artifacts |
| `ai-boilerplate-admin` | Vue 3 admin console |
| `ai-boilerplate-uniapp` | Uni-app mobile client |
| `ai-boilerplate-pc` | Browser PC/web shell |
| `ai-boilerplate-electron` | Electron desktop shell |
| `ai-boilerplate-tauri` | Tauri desktop shell |
| `ai-boilerplate-chrome-extension` | Manifest V3 Chrome extension |
| `ai-boilerplate-ios` | SwiftUI iOS starter |
| `ai-boilerplate-android` | Native Android starter |

## Environment

- Node.js 20 or newer for frontend templates.
- pnpm for `ai-boilerplate-admin`, `ai-boilerplate-uniapp`, and `ai-boilerplate-pc`.
- npm for `ai-boilerplate-electron`, `ai-boilerplate-tauri`, and `ai-boilerplate-chrome-extension`.
- Go 1.24 for `ai-boilerplate-backend`.
- Xcode and Swift Package Manager for `ai-boilerplate-ios`.
- Android Studio or a compatible Android SDK for `ai-boilerplate-android`.

Install dependencies inside the template directory you are working on. Do not
mix package managers within a template.

Recommended setup:

1. Clone the repository and enter the root directory.
2. Install dependencies only for the template you plan to run.
3. Run that template's verification command before making changes if you need a
   clean baseline.
4. Keep generated files in sync with the source artifact that owns them.

Template dependency commands:

| Project | Install |
| --- | --- |
| Backend | `go mod download` |
| Admin | `pnpm install` |
| Uni-app | `pnpm install` |
| PC web | `pnpm install` |
| Electron | `npm install` |
| Tauri web shell | `npm install` |
| Chrome extension | `npm install` |
| iOS | `swift package resolve` |
| Android | `./gradlew help` |

## Local Verification

Use the smallest command that covers the edited project:

| Project | Command |
| --- | --- |
| Backend | `go test ./...` |
| Admin | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | `pnpm check:type` |
| PC web | `pnpm type-check` |
| Electron | `npm run typecheck` |
| Tauri web shell | `npm run build` |
| Chrome extension | `npm run type-check` |
| Android | `./gradlew detekt test assembleDebug -x validateSigningDebug` |
| iOS | `swift build` |

The same commands are tracked in `docs/verification.md`.

## Daily Development

- Use this priority order when choosing unattended cleanup or maintenance work:
  open `P0`/`P1`/`blocker` issues or PRs, failing verification commands, core
  workflow regressions, user-visible template issues, then low-risk
  documentation or metadata drift.
- Start from the template README when working in a specific surface.
- Use root documentation for cross-template policy, release checks, and subtree
  operations.
- Keep backend API changes protobuf-first. Do not manually patch generated Go,
  Swagger, admin API client, or uni-app API client files unless the generator
  itself is the subject of the change.
- For admin and uni-app API changes, regenerate backend Swagger first, then run
  the matching frontend API generator.
- Prefer targeted verification while developing, then run every affected
  template check before committing.
- Keep app-specific secrets, signing files, and production endpoints out of the
  repository. Use local environment files or platform secrets.

## Generated Artifacts

Backend API, GORM, and Swagger files are generated from the backend toolchain.
When API definitions or database tables change, regenerate from
`ai-boilerplate-backend` with the matching Makefile target instead of editing
generated files by hand:

- `make sqltopb <admin|app> <tables>` creates protobuf definitions from SQL.
- `make api` regenerates Go API, HTTP, gRPC, validation, errors, and Swagger
  artifacts.
- `make gorm` regenerates GORM models and repository code.
- `make pbtocode` regenerates data and service code from protobuf definitions.

After backend API generation, run `go test ./...` from
`ai-boilerplate-backend`. If the admin or uni-app clients consume the changed
Swagger output, run their `api:gen` script and type check the changed client.

## Frontend API Clients

Admin and uni-app API clients are generated from backend Swagger artifacts:

| Client | Generate | Verify |
| --- | --- | --- |
| Admin | `pnpm api:gen` from `ai-boilerplate-admin` | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | `pnpm api:gen` from `ai-boilerplate-uniapp` | `pnpm check:type` |

Regenerate clients in the same change as the backend contract update so request
types, response types, and runtime paths stay aligned.

## Native Configuration

- iOS environment values live under
  `ai-boilerplate-ios/AIBoilerplate/Configurations/XCConfigs`.
- Android build configuration lives under `ai-boilerplate-android/app` and
  `ai-boilerplate-android/gradle`.
- Keep placeholder bundle IDs and endpoint values local to starter templates
  unless a product-specific app identifier is intentionally being added.
- Do not commit signing assets, private provisioning profiles, keystores, or
  Firebase service files.

## Subtree Workflow

The root `Makefile` manages subtree templates for backend, admin, and uni-app.
Run subtree commands only from the main worktree root. The Makefile blocks
subtree operations from linked worktrees because subtree state is easiest to
corrupt there.

Common commands:

- `make subtree-status`
- `make subtree-pull-backend`
- `make subtree-pull-admin`
- `make subtree-pull-app`
- `make subtree-push-backend`
- `make subtree-push-admin`
- `make subtree-push-app`

Do not use `make git-clean` unless the repository owner explicitly asks for a
history rewrite.

## Troubleshooting

- If a Node check cannot find a package, run the template's documented install
  command from that template directory.
- If a backend generator is missing, run `make init` from
  `ai-boilerplate-backend` and retry the generator target.
- If an iOS command cannot resolve packages, run `swift package resolve` from
  `ai-boilerplate-ios`.
- If Android cannot find an SDK, open the project in Android Studio or set
  `ANDROID_HOME` to a compatible local SDK path.
- If a verification command is unavailable on the current machine, record the
  exact command and failure output in the task summary, then run the closest
  static check for the edited files.

## Release Checklist

Before merging a change that touches more than one template:

1. Run each relevant targeted verification command.
2. Run the full verification set from `docs/verification.md` when local SDKs are
   available.
3. Confirm generated artifacts are current if backend API or database files
   changed.
4. Confirm `git status --short` only contains intended changes.
5. Commit with a concise message that names the affected template or workflow.
