# ai-boilerplate-pc

`ai-boilerplate-pc` is the browser-based desktop shell for AI Boilerplate.
It ships as a product-facing Vue 3 app instead of a generated starter.

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

## Included Surface

- `HomeView` is the main product shell.
- `BackToTop` stays mounted globally for long operator flows.
- Router entries are intentionally small so product routes can be added without starter noise.

Choose this template for a product-facing browser app that does not need admin
permissions or desktop-native APIs. Use the admin template for internal CRUD and
dense operator workflows.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
