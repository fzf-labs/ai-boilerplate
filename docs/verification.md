# Verification Commands

Use the smallest command that covers the project you touched. Run from the
listed directory unless noted otherwise.

Install dependencies inside each template directory before running Node-based
checks. Backend checks require Go, Android checks require an Android SDK, and
iOS checks require Swift Package Manager/Xcode tooling.

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

When a command cannot run because a local SDK, package manager, signing asset,
or generated workspace is missing, keep the exact command output in the task
summary and run the targeted static checks that cover the edited files.

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
