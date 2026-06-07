# Verification Commands

Use the smallest command that covers the project you touched. Run from the
listed directory unless noted otherwise.

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

When a command cannot run because a local SDK, package manager, signing asset,
or generated workspace is missing, keep the exact command output in the task
summary and run the targeted static checks that cover the edited files.
