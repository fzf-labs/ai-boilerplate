# Operations Guide

This repository is a multi-template workspace. Use the command that matches the
project you changed, then run the full verification set before publishing a
cross-template change.

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

## Release Checklist

Before merging a change that touches more than one template:

1. Run each relevant targeted verification command.
2. Run the full verification set from `docs/verification.md` when local SDKs are
   available.
3. Confirm generated artifacts are current if backend API or database files
   changed.
4. Confirm `git status --short` only contains intended changes.
5. Commit with a concise message that names the affected template or workflow.
