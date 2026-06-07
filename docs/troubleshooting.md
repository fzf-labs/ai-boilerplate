# Troubleshooting

Start with the smallest command that reproduces the failure. Read the full error
message before changing code, then compare the failing template with its README
and `docs/verification.md`.

## Dependency Install Failures

| Symptom | Check | Fix |
| --- | --- | --- |
| `pnpm` command is missing | `corepack --version` or `pnpm --version` | Enable Corepack or install pnpm, then rerun from the template directory. |
| Admin or uni-app install rejects npm/yarn | `package.json` `preinstall` script | Use `pnpm install`; do not mix package managers. |
| npm template cannot find packages | `node --version` and `npm install` output | Use Node.js 20 or newer, run `npm install` in that template only. |
| Go packages fail to resolve | `go env GOPROXY` and network access | Run `go mod download` from `ai-boilerplate-backend`. |

## Verification Failures

| Template | Common Cause | First Command |
| --- | --- | --- |
| Backend | Missing generated files, stale protobuf output, or local Go version mismatch | `go test ./...` |
| Admin | Stale generated API client or workspace package type drift | `pnpm check:type --filter=@vben/web-antd` |
| Uni-app | Missing generated app API client or env type mismatch | `pnpm check:type` |
| PC web | Vue type error or stale dependency install | `pnpm type-check` |
| Electron | Node/web tsconfig drift | `npm run typecheck` |
| Tauri | Renderer type error or missing Vite dependency | `npm run build` |
| Chrome extension | Manifest or runtime message type mismatch | `npm run type-check && npm test` |
| iOS | Missing SPM dependencies, Tuist setup, or Xcode toolchain mismatch | `swift build` |
| Android | Missing Android SDK, incompatible JDK, or Gradle cache issue | `./gradlew detekt test assembleDebug -x validateSigningDebug` |

## API Integration Problems

- If admin requests 404 locally, confirm `VITE_GLOB_API_URL` and the local proxy
  route point at the backend path used by generated Swagger.
- If uni-app H5 works but a mini-program build fails, check the per-channel
  `VITE_SERVER_BASEURL_WEIXIN_*` values and platform domain allowlists.
- If native mobile clients cannot reach the backend, replace loopback addresses
  with an emulator host alias or LAN IP as documented in `docs/environment.md`.
- If generated frontend types do not match backend responses, regenerate backend
  Swagger first, then regenerate the affected frontend client.

## Generated File Drift

Generated drift usually shows up as type errors in frontend clients or missing
Go methods after protobuf edits. Use `docs/generated-artifacts.md` to rerun the
full generation chain from the source artifact instead of patching generated
files by hand.

## Native Tooling

- iOS warnings about unused package dependencies can be acceptable for a starter
  template, but `swift build` must still finish successfully.
- Android Gradle deprecation warnings are not release blockers by themselves.
  A failed task, missing SDK, signing error, or incompatible Gradle/JDK pairing
  is the actionable failure.
