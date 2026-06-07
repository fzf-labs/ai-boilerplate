# Technical Decisions

This document records the main stack choices in the boilerplate and the
practical reason each one exists.

The repository intentionally contains several app surfaces instead of one
runtime. Each template is meant to be usable on its own, while backend contracts
and generated clients keep shared product flows aligned.

## Backend

- **Go + Kratos**: gives the backend a typed service structure, generated HTTP
  and gRPC bindings, and middleware patterns that fit API-heavy products.
- **PostgreSQL + GORM**: keeps the default data layer relational and easy to
  inspect while still supporting generated models and repositories.
- **Protobuf-first APIs**: keeps admin and app API contracts explicit, allows
  Swagger generation, and reduces drift between backend and clients.
- **Redis / Rueidis / Asynq**: provides cache, queue, and background job
  building blocks without forcing every starter app to use external SaaS.

Generated artifacts are treated as implementation output, not source of truth.
Database schema changes should flow through SQL, protobuf generation, Go API
generation, and client generation in that order when a feature crosses the
backend/frontend boundary.

## Admin Console

- **Vue 3 + TypeScript**: matches the existing Vben-based admin architecture
  and keeps generated API types useful at the component boundary.
- **Ant Design Vue**: provides dense operational UI controls for CRUD, forms,
  tables, modals, and admin workflows.
- **pnpm workspace + Turbo**: keeps shared packages, app code, and internal
  tooling in one install graph while allowing targeted builds and type checks.

The admin template favors operational density over marketing-style layout.
Tables, filter forms, modals, and permission-aware actions should follow the
existing Vben and Ant Design Vue patterns before adding new UI abstractions.

## Mobile App

- **Uni-app + Vue 3 + TypeScript**: supports H5, App, and mini-program targets
  from the same client codebase.
- **wot-design-uni**: provides mobile-ready UI controls that work across the
  uni-app targets already configured in the template.
- **Generated API clients**: keep app requests aligned with the backend Swagger
  output and reduce hand-written request surface area.

The uni-app template is the default cross-platform mobile client. Native iOS or
Android should be used when a product needs platform-specific capability,
distribution control, or native performance that would be awkward in uni-app.

## Desktop and Web Shells

- **PC web shell**: a lightweight Vue/Vite browser entry for public or
  operational web experiences outside the admin console.
- **Electron shell**: used when desktop distribution needs Chromium and Node.js
  integration.
- **Tauri shell**: used when desktop distribution should keep the runtime
  footprint smaller and can rely on a Rust-backed shell.
- **Chrome extension**: Manifest V3 is the current Chrome extension model and
  fits tab inspection or browser-side productivity workflows.

These shells are intentionally small. Product teams can choose the lightest
surface that matches their distribution target instead of forcing every web or
desktop workflow through the admin console.

## Native Clients

- **SwiftUI iOS starter**: keeps iOS screens declarative and aligns with modern
  Apple platform guidance.
- **Native Android starter**: provides a Kotlin Android baseline for teams that
  need native mobile capabilities beyond the uni-app client.

Native templates should keep signing, bundle IDs, endpoints, and third-party app
IDs configurable. Repository defaults should remain safe starter values rather
than product secrets.

## Verification Policy

Each template owns its own verification command. The repository-level standard
is to run the smallest relevant check for local changes and the full list in
`docs/verification.md` before publishing broad cross-template work.

## Updating Decisions

Record a new decision here when a template changes its primary framework,
package manager, generated-artifact strategy, runtime target, or verification
standard. Routine copy edits, generated file refreshes, and narrow page changes
should stay in the affected template README or code comments instead.

## Selection Guide

Use this repository as a menu of starting points:

| Need | Start With |
| --- | --- |
| Typed API service, generated HTTP/gRPC, data layer | Backend |
| Internal operations, CRUD, permissions, dense dashboards | Admin console |
| Cross-platform mobile app or mini-program | Uni-app |
| Public browser app or lightweight web shell | PC web |
| Desktop app with Node integration | Electron |
| Desktop app with smaller runtime footprint | Tauri |
| Browser extension workflow | Chrome extension |
| Native iOS app | SwiftUI iOS |
| Native Android app | Kotlin Android |

Prefer one primary client per product workflow. Add another template only when a
specific platform or distribution requirement justifies the extra maintenance.
