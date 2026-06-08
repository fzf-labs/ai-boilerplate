# Troubleshooting

Start with the smallest command that reproduces the failure. Read the full error
message before changing code, then compare the failing template with its README
and `docs/verification.md`.

## First Response Flow

1. Run the documented verification command from the template directory.
2. Confirm whether the failure is deterministic by rerunning the smallest
   failing command once.
3. Identify the owning template or generated artifact from the file path in the
   error.
4. Compare the failing file with a nearby working example before editing.
5. If generated output is involved, regenerate from the source artifact instead
   of patching the generated file by hand.
6. After the fix, rerun the targeted command and record the exact result.

Stop and ask before changing production credentials, signing setup, live third-
party account settings, or behavior whose expected product outcome is unclear.

## Escalation Path

If the first response flow does not isolate the issue:

1. Rerun the smallest failing command and save the exact error.
2. Check the most recent diff and recent commits touching the failing template.
3. Compare the failing file with a nearby working implementation.
4. Add temporary diagnostics only at component boundaries, then remove them
   before committing.
5. If three separate fix attempts fail, stop and reassess the architecture or
   task scope before making another code change.

Escalate to the user when the next step requires external accounts, production
permissions, private signing material, paid services, or a product decision not
covered by the repo docs.

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

## Documentation Drift

If a command in the docs no longer matches a template:

1. Check the template `package.json`, `Makefile`, `go.mod`, Gradle files, or
   Swift package first.
2. Update the root guide and the template README together when both mention the
   command.
3. Run the documented command after the edit, or record why local prerequisites
   prevent it.
4. Use `git diff --check` for documentation-only changes to catch whitespace and
   formatting errors.

## Native Tooling

- iOS warnings about unused package dependencies can be acceptable for a starter
  template, but `swift build` must still finish successfully.
- Android Gradle deprecation warnings are not release blockers by themselves.
  A failed task, missing SDK, signing error, or incompatible Gradle/JDK pairing
  is the actionable failure.
