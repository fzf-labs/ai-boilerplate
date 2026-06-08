# Operations Guide

This repository is a multi-template workspace. Use the command that matches the
project you changed, then run the full verification set before publishing a
cross-template change.

Most work should stay inside one template directory. Cross-template changes are
reserved for shared API contracts, generated client syncs, documentation, or
release workflow updates.

When a request is not already mapped to a template, start with the selection
guide in the root README or `docs/technical-decisions.md` before editing code.
That keeps the work anchored to the owning artifact and avoids drifting into
downstream consumers too early.

## Choosing Change Scope

Start by naming the artifact that owns the behavior:

| Change Type | Owning Artifact | Required Follow-up |
| --- | --- | --- |
| Database column, table, or relation | Backend SQL and generated backend code | Regenerate protobuf/GORM outputs and run backend tests. |
| API request, response, route, or permission | Backend protobuf source | Regenerate backend Swagger, then regenerate affected frontend clients. |
| Admin-only table, form, or operator flow | Admin template | Type-check `@vben/web-antd`; regenerate clients only when Swagger changed. |
| Mobile app page or mini-program behavior | Uni-app template | Type-check uni-app; update env notes for platform-specific URLs. |
| Native app package identity or endpoint | iOS or Android template | Keep signing and production identifiers out of git. |
| Root policy, verification, or release workflow | Root docs and Makefile when needed | Run doc/example checks plus any touched template verification. |

If a change appears to require more than one owning artifact, update the source
artifact first and treat downstream generated files as consumers. Stop and ask
before changing product behavior that depends on external accounts, production
permissions, real signing identities, or customer-specific platform IDs.

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
| Chrome extension | `npm run type-check && npm test` |
| Android | `./gradlew detekt test assembleDebug -x validateSigningDebug` |
| iOS | `swift build` |

The same commands are tracked in `docs/verification.md`.

## Start Work

Use this workflow before changing code or documentation:

1. Identify the owning template from the root README selection table.
2. Check open issues, pull requests, and local requirement documents for
   `P0`, `P1`, `blocker`, or `critical` signals.
3. Read any linked issue, requirement, or recent failing command output before
   editing.
4. Read the owning template README and any relevant root guide.
5. Install dependencies only inside the owning template.
6. Run the smallest verification command from `docs/verification.md` when a
   clean baseline matters.
7. Make the smallest change that satisfies the task.
8. Regenerate artifacts from their source when contracts or database schemas
   changed.
9. Run the targeted verification command for every touched template.
10. Check `git status --short` before committing.

Use this priority order when choosing unattended cleanup or maintenance work:

| Rank | Signal | Action |
| --- | --- | --- |
| 1 | Open `P0`, `P1`, `blocker`, or `critical` issues or PRs | Fix or summarize the blocking item first. |
| 2 | Failing documented verification commands | Reproduce, find root cause, fix, and rerun the command. |
| 3 | Core workflow regressions | Prioritize setup, generation, auth, API, build, and release flows. |
| 4 | User-visible template issues | Fix broken copy, links, defaults, and UI paths that a starter user sees. |
| 5 | Low-risk documentation or metadata drift | Improve docs only when higher-priority signals are absent. |
| 6 | Product decisions or external-account work | Stop and ask for confirmation before changing behavior. |

If no priority labels or failing checks exist, use open issues and root
documentation gaps as the next source of truth. Record the source, rationale,
acceptance criteria, and verification evidence for each task.

For broad documentation requests such as "add operations and technical decision
docs", treat each missing or stale doc link, template README handoff, metadata
default, and verification-reference mismatch as its own task. Keep the source
issue attached to each task so future unattended runs can distinguish completed
documentation coverage from new product work.

Stop and ask before starting when the highest-priority item requires production
credentials, external account access, signing assets, customer-specific IDs, or
a product decision that is not already documented.

## Issue And Requirement Intake

Use this checklist when an issue, PR, or requirement document is the source of a
task:

1. Record the source link or file path before editing.
2. Capture the exact acceptance criteria or infer the smallest observable
   outcome when the source is brief.
3. Map the request to one owning artifact from the change-scope table.
4. Check whether any generated output must be refreshed from that source.
5. Pick the targeted verification command before changing files.
6. Stop if the requested outcome depends on undocumented product behavior,
   production accounts, signing credentials, or customer-specific identifiers.

For unattended maintenance, write each chosen task in this form:

| Field | Required Content |
| --- | --- |
| Task name | One concrete outcome, not an activity label. |
| Source | Issue, PR, failing command, document, or file path that justifies the work. |
| Priority rationale | Which priority rank made this task more important than alternatives. |
| Acceptance criteria | A binary or observable condition plus the verification command. |

## Template Runbook Quick Reference

| Template | Start Command | Verify Command | Publish Note |
| --- | --- | --- | --- |
| Backend | `make run` | `go test ./...` | Regenerate API/GORM output before committing contract changes. |
| Admin | `pnpm dev:antd` | `pnpm check:type --filter=@vben/web-antd` | Keep CRUD and permission workflows in Ant Design Vue/Vben patterns. |
| Uni-app | `pnpm dev:h5` | `pnpm check:type` | Check platform-specific URL behavior before app or mini-program release. |
| PC web | `pnpm dev` | `pnpm type-check` | Keep public browser flows separate from admin-only APIs. |
| Electron | `npm run dev` | `npm run typecheck` | Keep signing and update-server secrets outside the repo. |
| Tauri | `npm run tauri dev` | `npm run build` | Treat renderer secrets as public; native secrets belong in platform storage. |
| Chrome extension | `npm run dev` | `npm run type-check && npm test` | Add manifest permissions only when a feature requires them. |
| iOS | `tuist generate` | `swift build` | Do not commit real provisioning profiles or signing identities. |
| Android | `./gradlew assembleDebug -x validateSigningDebug` | `./gradlew detekt test assembleDebug -x validateSigningDebug` | Keep keystores and Firebase service files out of git. |

## Daily Development

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

For the complete generation order and review checklist, see
`docs/generated-artifacts.md`.

## Environment Configuration

Repository defaults are for local development. Real deployment values belong in
local config files, CI secrets, platform secret stores, or mobile signing
systems. See `docs/environment.md` for the per-template file list and local
endpoint defaults.

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

For a template-by-template troubleshooting matrix, see `docs/troubleshooting.md`.

## Release Checklist

Before merging a change that touches more than one template:

1. Run each relevant targeted verification command.
2. Run the full verification set from `docs/verification.md` when local SDKs are
   available.
3. Confirm generated artifacts are current if backend API or database files
   changed.
4. Confirm `git status --short` only contains intended changes.
5. Commit with a concise message that names the affected template or workflow.

For the expanded release and publishing checklist, see `docs/release.md`.
