# AI Boilerplate Tauri

`ai-boilerplate-tauri` is the Tauri desktop starter for this repository. It
uses Vue 3 for the renderer and Rust/Tauri for the native shell.

## Project Setup

```bash
npm install
npm run build
```

For native development and packaging, install the Tauri prerequisites for your
platform, then run:

```bash
npm run tauri dev
```

Choose Tauri when a smaller desktop runtime matters and the native shell can be
implemented through Tauri/Rust APIs. Use Electron instead when broad Node.js
integration is a real runtime requirement.

`npm run build` verifies the Vue renderer and production web bundle. Run the
Tauri native build on the target platform before publishing a desktop artifact,
and keep signing credentials out of the repository.

## Ownership Notes

Start here when desktop distribution is required and the native work can live in
Tauri/Rust APIs without broad Node.js integration. If the requirement is only a
browser route, use `ai-boilerplate-pc`; if the runtime needs Node modules or
Electron-specific APIs, use `ai-boilerplate-electron`.

Record the target operating system before changing Tauri capabilities,
installer metadata, updater settings, or signing-related configuration. Use
`npm run build` for renderer changes and the platform Tauri build before
publishing native artifacts.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
