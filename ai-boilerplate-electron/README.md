# ai-boilerplate-electron

Electron desktop shell for AI Boilerplate. It uses Electron Vite, Vue 3, and
TypeScript to provide a native desktop entry point for the same product modules
as the backend, admin, and mobile templates.

## Recommended IDE Setup

- [VSCode](https://code.visualstudio.com/) + [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint) + [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar)

## Project Setup

### Install

```bash
npm install
```

### Development

```bash
npm run dev
```

### Verification

```bash
npm run typecheck
npm run build
```

### Build

```bash
# For windows
npm run build:win

# For macOS
npm run build:mac

# For Linux
npm run build:linux
```

## Packaging Notes

`electron-builder.yml` owns package identity, artifact names, and the Electron
download mirror. Keep mirror settings there so npm does not warn about unknown
project-level config during type checks and installs.

Choose Electron when the desktop app needs Chromium plus Node.js integration at
runtime. Keep signing identities, notarization credentials, update-server
secrets, and production endpoints outside the repository.

Run `npm run typecheck` for routine code changes. Run the platform build command
for the target OS before publishing an installer.

## Ownership Notes

Start here when desktop packaging, tray/window behavior, native dialogs, local
filesystem access, or Node.js runtime integration is part of the requirement.
If the workflow can run as a normal browser page, use `ai-boilerplate-pc`
instead; if it needs a smaller shell without broad Node integration, compare it
with `ai-boilerplate-tauri`.

Document the target operating system before changing installer metadata,
auto-update behavior, or signing-related configuration. Verification for routine
code is `npm run typecheck`; installer work also needs the matching platform
build command.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
