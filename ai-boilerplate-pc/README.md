# ai-boilerplate-pc

`ai-boilerplate-pc` is the browser-based desktop shell for AI Boilerplate.
It ships as a product-facing Vue 3 app instead of a generated starter.

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Run End-to-End Tests with [Playwright](https://playwright.dev)

```sh
# Install browsers for the first run
npx playwright install

# When testing on CI, must build the project first
npm run build

# Runs the end-to-end tests
npm run test:e2e
# Runs the tests only on Chromium
npm run test:e2e -- --project=chromium
# Runs the app smoke test only
npm run test:e2e -- e2e/vue.spec.ts
# Runs the tests in debug mode
npm run test:e2e -- --debug
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

## Included Surface

- `HomeView` is the main product shell.
- `BackToTop` stays mounted globally for long operator flows.
- Router entries are intentionally small so product routes can be added without starter noise.
