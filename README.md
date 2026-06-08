# AI Boilerplate

AI Boilerplate is a collection of starter projects for backend, admin, app,
desktop, and iOS product development.

## Quick Start

Use the template directory that matches the surface you want to work on. Each
template owns its dependency install and verification commands.

| Template | Install | First Run | First Verification |
| --- | --- | --- | --- |
| Backend | `go mod download` | `make run` | `go test ./...` |
| Admin | `pnpm install` | `pnpm dev:antd` | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | `pnpm install` | `pnpm dev:h5` | `pnpm check:type` |
| PC web | `pnpm install` | `pnpm dev` | `pnpm type-check` |
| Electron | `npm install` | `npm run dev` | `npm run typecheck` |
| Tauri | `npm install` | `npm run tauri dev` | `npm run build` |
| Chrome extension | `npm install` | `npm run dev` | `npm run type-check && npm test` |
| iOS | `swift package resolve` | `tuist generate` | `swift build` |
| Android | `./gradlew help` | `./gradlew assembleDebug -x validateSigningDebug` | `./gradlew detekt test assembleDebug -x validateSigningDebug` |

Run commands from the template directory. The detailed setup notes live in each
template README, and the full verification matrix is in
[Verification commands](docs/verification.md).

## Choosing A Template

Pick the template that matches the primary surface of the product first, then
consult the template README for setup details.

| Need | Start With |
| --- | --- |
| Backend API, data layer, generated contracts | `ai-boilerplate-backend` |
| Internal operations dashboard | `ai-boilerplate-admin` |
| Cross-platform mobile app or mini-program | `ai-boilerplate-uniapp` |
| Public browser app or lightweight web shell | `ai-boilerplate-pc` |
| Desktop app with Chromium and Node integration | `ai-boilerplate-electron` |
| Smaller-footprint desktop shell | `ai-boilerplate-tauri` |
| Browser extension workflow | `ai-boilerplate-chrome-extension` |
| Native iOS app | `ai-boilerplate-ios` |
| Native Android app | `ai-boilerplate-android` |

For the stack rationale behind those choices, see
[Technical decisions](docs/technical-decisions.md).

## Templates

- `ai-boilerplate-backend`: Go backend boilerplate.
- `ai-boilerplate-admin`: Admin frontend boilerplate.
- `ai-boilerplate-uniapp`: Uni-app mobile app boilerplate.
- `ai-boilerplate-pc`: PC/web frontend boilerplate.
- `ai-boilerplate-electron`: Electron desktop boilerplate.
- `ai-boilerplate-tauri`: Tauri desktop boilerplate.
- `ai-boilerplate-chrome-extension`: Chrome extension boilerplate based on Manifest V3.
- `ai-boilerplate-ios`: SwiftUI iOS boilerplate based on `nimblehq/ios-templates`.
- `ai-boilerplate-android`: Native Android boilerplate based on
  `its-me-debk007/kotlin-android-mvvm-template`.

## Documentation

- [Operations guide](docs/operations.md): start-work workflow, task priority, generated artifacts, subtree workflow, and release checks.
- [Technical decisions](docs/technical-decisions.md): stack choices, template tradeoffs, and selection guidance.
- [Verification commands](docs/verification.md): targeted and full verification commands for every template.
- [Generated artifacts](docs/generated-artifacts.md): backend API, Swagger, GORM, and frontend client generation flow.
- [Environment configuration](docs/environment.md): local endpoints, secrets, and per-template environment files.
- [Troubleshooting](docs/troubleshooting.md): common setup, dependency, verification, and API integration failures.
- [Release checklist](docs/release.md): pre-publish checks for single-template and cross-template changes.

## Development Workflow

1. Choose the template that owns the product surface you are changing.
2. Read that template README plus the matching root guide in the documentation list above.
3. Install dependencies only inside the touched template directory.
4. Run the smallest verification command from `docs/verification.md` before and after the change when you need a clean baseline.
5. Keep generated files in sync with their source artifacts.
6. Before committing, confirm `git status --short` only contains intentional changes.

## Local Development Notes

- Keep dependency installs inside the template directory.
- Backend API, Swagger, and GORM files are generated; use the backend Makefile
  targets instead of editing generated artifacts by hand.
- Admin and uni-app API clients are generated from backend Swagger output.
- Run the smallest relevant verification command after local changes, and run
  the full verification set before publishing cross-template changes.
