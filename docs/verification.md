# Verification

The repository CI gate lives in [.github/workflows/quality-gates.yml](../.github/workflows/quality-gates.yml).

## Commands

| Template | Core verification |
| --- | --- |
| Backend | `golangci-lint run --config .golangci.yml ./... -v`, `go test ./...` |
| Admin | `pnpm lint`, `pnpm check`, `pnpm build` |
| UniApp | `pnpm lint`, `pnpm check:type`, `pnpm build` |
| PC | `pnpm lint`, `pnpm test:unit`, `pnpm build` |
| Chrome extension | `npm ci`, `npm run build` |
| Electron | `npm install --no-audit --no-fund`, `npm run lint`, `npm run build` |
| Tauri | `cargo check --manifest-path src-tauri/Cargo.toml`, `npm install --no-audit --no-fund`, `npm run build` |
| Android | `./gradlew --no-daemon detekt test assembleDebug -x validateSigningDebug` |
| iOS | `mise install`, `mise exec -- bundle install`, `mise exec -- bundle exec arkana`, `mise exec -- tuist generate --no-open`, `mise exec -- bundle exec fastlane buildAndTestLane` |
