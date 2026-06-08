# Technical Decisions

This document records the main stack choices in the boilerplate and the
practical reason each one exists.

The repository intentionally contains several app surfaces instead of one
runtime. Each template is meant to be usable on its own, while backend contracts
and generated clients keep shared product flows aligned.

## Selection In Practice

Start with the smallest template that owns the user-visible behavior. Use the
backend when schema, API contract, auth, permissions, or background jobs
change. Use admin for internal CRUD and operator flows, uni-app for
cross-platform mobile and mini-program flows, PC web for lightweight browser
shells, and iOS or Android only when platform-specific APIs or distribution
needs justify the native maintenance cost.

When a task crosses templates, update the source artifact first and regenerate
downstream clients or generated output in the same commit.

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

Use backend generation when the change affects stored data, API shape,
authorization behavior, or shared response semantics. Do not use backend
generation for client-only copy, layout, route labels, or presentation defaults.
Those changes should stay in the client template that owns the experience.

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

## Template Tradeoffs

| Template | Choose When | Avoid When | Verification |
| --- | --- | --- | --- |
| Backend | You need typed APIs, generated HTTP/gRPC bindings, database access, jobs, or shared contracts. | The change is only client-side copy, routing, or presentation. | `go test ./...` |
| Admin console | You need internal CRUD, permissions, dense tables, operational forms, or admin-only workflows. | The experience is public marketing, lightweight web content, or mobile-first. | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | You need one cross-platform mobile or mini-program client. | You need deep native APIs, custom platform distribution, or native-only performance. | `pnpm check:type` |
| PC web | You need a lightweight browser app outside the admin console. | The workflow needs admin permissions, dense back-office tables, or native desktop APIs. | `pnpm type-check` |
| Electron | You need desktop distribution with Chromium and Node integration. | A browser app or smaller native shell can satisfy the workflow. | `npm run typecheck` |
| Tauri | You need a smaller desktop shell and can support Rust/Tauri prerequisites. | You need broad Node integration in the desktop runtime. | `npm run build` |
| Chrome extension | You need browser-side tab, page, or productivity workflows. | The feature is a normal web app or needs server-owned state only. | `npm run type-check && npm test` |
| iOS | You need a native Apple-platform app or native iOS capabilities. | Uni-app can provide the needed mobile workflow with lower maintenance. | `swift build` |
| Android | You need a native Android app or Android-specific capabilities. | Uni-app can provide the needed mobile workflow with lower maintenance. | `./gradlew detekt test assembleDebug -x validateSigningDebug` |

## Decision Rules

- Start with the backend only when the task changes shared data, contracts,
  authentication, permissions, async jobs, or API behavior.
- Start with admin for internal operator workflows and keep the UI dense,
  scannable, and permission-aware.
- Start with uni-app for the default mobile surface. Use native iOS or Android
  only when the product has platform-specific requirements.
- Start with PC web for public or lightweight browser workflows that do not
  belong in the admin console.
- Choose Electron over Tauri when Node integration is a real runtime
  requirement. Choose Tauri when a smaller desktop footprint matters more.
- Choose the Chrome extension only when browser context, tabs, or extension
  permissions are central to the feature.

When a task crosses templates, decide which artifact owns the behavior first.
Backend contracts and database schemas should lead generated client updates;
client-only visual or copy changes should stay in the client template.

## Generation Decision Points

Use this checklist before editing generated files:

| Question | If Yes | If No |
| --- | --- | --- |
| Does the API route, method, request, response, or validation rule change? | Edit protobuf source and run backend API generation. | Keep generated API files unchanged. |
| Does the database table or column shape change? | Update SQL/database source and run GORM generation. | Avoid touching generated model or DAO files. |
| Does admin or uni-app consume changed Swagger output? | Regenerate the matching frontend API client. | Do not run client generation just to refresh timestamps. |
| Is the task only UI copy, navigation, or styling? | Change the owning client template directly. | Re-evaluate whether the backend owns the behavior. |
