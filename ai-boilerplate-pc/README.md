# ai-boilerplate-pc

`ai-boilerplate-pc` is the browser-based desktop shell for AI Boilerplate.
It ships as a product-facing Vue 3 app instead of a generated starter.
Use Node.js 20.19+ and pnpm 10.14+ from this directory.

## Project Setup

```sh
pnpm install
```

### Compile and Hot-Reload for Development

```sh
pnpm dev
```

### Type-Check, Compile and Minify for Production

```sh
pnpm type-check
pnpm build
```

Use `pnpm type-check` as the first verification command for routine PC web
changes. Run `pnpm build`, unit tests, or Playwright only when the change touches
bundling, runtime behavior, or browser workflows.

For release evidence, record the exact command that matches the changed surface:
type-only copy or route work needs `pnpm type-check`, bundle configuration needs
`pnpm build`, and browser behavior changes need the relevant unit or Playwright
test in addition to the type check.

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
pnpm test:unit
```

### Run End-to-End Tests with [Playwright](https://playwright.dev)

```sh
# Install browsers for the first run
npx playwright install

# When testing on CI, must build the project first
pnpm build

# Runs the end-to-end tests
pnpm test:e2e
# Runs the tests only on Chromium
pnpm test:e2e -- --project=chromium
# Runs the app smoke test only
pnpm test:e2e -- e2e/vue.spec.ts
# Runs the tests in debug mode
pnpm test:e2e -- --debug
```

### Lint with [ESLint](https://eslint.org/)

```sh
pnpm lint
```

`pnpm lint` runs ESLint with auto-fix enabled, so it may rewrite files. Use
`pnpm type-check` for a read-only routine verification check, and use
`pnpm lint` or `pnpm format` when you intend to accept formatting or lint-fix
changes.

## Included Surface

- `HomeView` is the main product shell.
- `BackToTop` stays mounted globally for long operator flows.
- Router entries are intentionally small so product routes can be added without starter noise.

Choose this template for a product-facing browser app that does not need admin
permissions or desktop-native APIs. Use the admin template for internal CRUD and
dense operator workflows.

## Ownership Notes

Start here for lightweight public browser routes, product-facing copy, simple
Vue/Vite interaction, and web-only presentation that does not require admin
permissions. Keep browser shell work separate from admin CRUD so the public
surface stays small and easy to verify.

Start in the backend first when the page needs a new shared API contract or
server-owned behavior. Start in Electron or Tauri only when the same workflow
requires desktop packaging, native APIs, local filesystem access, or runtime
capabilities that a browser page cannot provide.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
