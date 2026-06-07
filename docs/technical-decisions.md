# Technical Decisions

This document records the main stack choices in the boilerplate and the
practical reason each one exists.

## Backend

- **Go + Kratos**: gives the backend a typed service structure, generated HTTP
  and gRPC bindings, and middleware patterns that fit API-heavy products.
- **PostgreSQL + GORM**: keeps the default data layer relational and easy to
  inspect while still supporting generated models and repositories.
- **Protobuf-first APIs**: keeps admin and app API contracts explicit, allows
  Swagger generation, and reduces drift between backend and clients.
- **Redis / Rueidis / Asynq**: provides cache, queue, and background job
  building blocks without forcing every starter app to use external SaaS.

## Admin Console

- **Vue 3 + TypeScript**: matches the existing Vben-based admin architecture
  and keeps generated API types useful at the component boundary.
- **Ant Design Vue**: provides dense operational UI controls for CRUD, forms,
  tables, modals, and admin workflows.
- **pnpm workspace + Turbo**: keeps shared packages, app code, and internal
  tooling in one install graph while allowing targeted builds and type checks.

## Mobile App

- **Uni-app + Vue 3 + TypeScript**: supports H5, App, and mini-program targets
  from the same client codebase.
- **wot-design-uni**: provides mobile-ready UI controls that work across the
  uni-app targets already configured in the template.
- **Generated API clients**: keep app requests aligned with the backend Swagger
  output and reduce hand-written request surface area.

## Desktop and Web Shells

- **PC web shell**: a lightweight Vue/Vite browser entry for public or
  operational web experiences outside the admin console.
- **Electron shell**: used when desktop distribution needs Chromium and Node.js
  integration.
- **Tauri shell**: used when desktop distribution should keep the runtime
  footprint smaller and can rely on a Rust-backed shell.
- **Chrome extension**: Manifest V3 is the current Chrome extension model and
  fits tab inspection or browser-side productivity workflows.

## Native Clients

- **SwiftUI iOS starter**: keeps iOS screens declarative and aligns with modern
  Apple platform guidance.
- **Native Android starter**: provides a Kotlin Android baseline for teams that
  need native mobile capabilities beyond the uni-app client.

## Verification Policy

Each template owns its own verification command. The repository-level standard
is to run the smallest relevant check for local changes and the full list in
`docs/verification.md` before publishing broad cross-template work.
