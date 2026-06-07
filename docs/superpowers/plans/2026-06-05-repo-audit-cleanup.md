# Repo Audit Cleanup Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Remove the highest-priority leftover placeholder/example content from the audited admin, uni-app, and iOS files, then verify the touched targets with the smallest useful checks and commit the result.

**Architecture:** Keep the changes local to the existing files, preserve current behavior where it is already correct, and only tighten behavior where the current placeholder content is visibly broken or misleading. Prefer small, explicit edits over broader refactors so each subsystem stays easy to validate independently.

**Tech Stack:** Vue 3 + TypeScript, Uni-App, Ant Design Vue, SwiftUI, pnpm, Swift Package Manager / Tuist.

---

### Task 1: Access control component

**Files:**
- Modify: `ai-boilerplate-admin/packages/effects/access/src/access-control.vue`

- [x] Remove the stale TODO-only comment block and make the empty-codes case render consistently with the rest of the access helpers.
- [x] Add support for an explicit all-vs-any match mode without changing the existing default behavior.

### Task 2: Native SwiftUI settings screen

**Files:**
- Modify: `ai-boilerplate-ios/Tuist/Interfaces/SwiftUI/Sources/Presentation/Modules/Settings/SettingsView.swift`

- [x] Replace the placeholder paragraph with a simple, real settings-style layout that reads like a finished destination screen.

### Task 3: Uni-app route exclusions

**Files:**
- Modify: `ai-boilerplate-uniapp/src/router/config.ts`

- [x] Remove the sample route exclusions so the config only contains intentional paths.

### Task 4: Uni-app native tabbar defaults

**Files:**
- Modify: `ai-boilerplate-uniapp/src/tabbar/config.ts`

- [x] Replace the remaining example image references with a real asset from `src/static/tabbar`.

### Task 5: Profile avatar upload flow

**Files:**
- Modify: `ai-boilerplate-uniapp/src/pages/profile/edit.vue`

- [x] Stop saving the temporary local file path directly and upload the chosen image before persisting the user profile.

### Task 6: Uni-app dev config noise

**Files:**
- Modify: `ai-boilerplate-uniapp/vite.config.ts`

- [x] Remove the startup debug logging and keep the config output clean.

### Task 7: Shared tree helper

**Files:**
- Modify: `ai-boilerplate-admin/apps/web-antd/src/utils/tree.ts`

- [x] Replace the duplicated implementation with a thin compatibility re-export to the shared helper.

### Task 8: Self-app schema copy

**Files:**
- Modify: `ai-boilerplate-admin/apps/web-antd/src/views/selfapp/info/data.ts`

- [x] Replace the generic package-name placeholder with a template-specific example.

### Task 9: WeChat menu form hints

**Files:**
- Modify: `ai-boilerplate-admin/apps/web-antd/src/views/gzh/menu/index.vue`

- [x] Remove the example AppId and page-path hints from the help text and keep the guidance factual.

### Task 10: Uni-app settings/legal copy

**Files:**
- Modify: `ai-boilerplate-uniapp/src/pages/settings/index.vue`

- [x] Remove the fake support email and hide or replace the empty legal/contact placeholders so the page does not advertise broken links.

### Verification

**Files to re-check:**
- `ai-boilerplate-admin/packages/effects/access/src/access-control.vue`
- `ai-boilerplate-ios/Tuist/Interfaces/SwiftUI/Sources/Presentation/Modules/Settings/SettingsView.swift`
- `ai-boilerplate-uniapp/src/pages/profile/edit.vue`
- `ai-boilerplate-uniapp/src/pages/settings/index.vue`

- [x] Run `pnpm check:type --filter=@vben/web-antd` from `ai-boilerplate-admin`.
- [x] Run `pnpm check:type` from `ai-boilerplate-uniapp`.
- [x] Run `swift build` (or the closest package build check available in `ai-boilerplate-ios`).
- [x] Run `npm run typecheck` from `ai-boilerplate-electron`.
- [x] Run `npm run build` from `ai-boilerplate-tauri`.
- [x] Run `pnpm type-check` from `ai-boilerplate-pc`.
- [x] Run `./gradlew detekt test assembleDebug -x validateSigningDebug` from `ai-boilerplate-android`.
- [x] Commit the finished changes with a concise message.
