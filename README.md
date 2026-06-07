# AI Boilerplate

AI Boilerplate is a collection of starter projects for backend, admin, app,
desktop, and iOS product development.

## Quick Start

Use the template directory that matches the surface you want to work on. Each
template owns its dependency install and verification commands.

```bash
# Backend
cd ai-boilerplate-backend
go test ./...

# Admin
cd ai-boilerplate-admin
pnpm install
pnpm check:type --filter=@vben/web-antd

# Uni-app
cd ai-boilerplate-uniapp
pnpm install
pnpm check:type
```

The desktop, browser extension, native Android, and iOS templates have their own
README files with the same setup pattern.

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

- [Operations guide](docs/operations.md)
- [Technical decisions](docs/technical-decisions.md)
- [Verification commands](docs/verification.md)

## Local Development Notes

- Keep dependency installs inside the template directory.
- Backend API, Swagger, and GORM files are generated; use the backend Makefile
  targets instead of editing generated artifacts by hand.
- Admin and uni-app API clients are generated from backend Swagger output.
- Run the smallest relevant verification command after local changes, and run
  the full verification set before publishing cross-template changes.
