# AI Boilerplate Admin

Vue 3 admin console template for AI Boilerplate. It is based on the Vben admin
workspace and includes generated API clients, Ant Design Vue views, shared UI
packages, and Turbo-powered type checks.

## Setup

Use Node.js 20.10 or newer and pnpm. Run commands from this directory.

```bash
pnpm install
pnpm dev:antd
```

The main app is `apps/web-antd`. Environment files for the app live in
`apps/web-antd/.env*`.

## API Configuration

`VITE_GLOB_API_URL` controls the API base path used by the request client. The
development default is `/api`, which is suitable for a local proxy or mock
server.

For local backend integration, run the backend service first and point the Vite
proxy or local gateway at the backend HTTP path. Keep production hostnames and
tokens out of checked-in `.env` files.

Generated admin API clients are produced from backend Swagger files:

```bash
pnpm api:gen
```

The generator reads `../ai-boilerplate-backend/doc/swagger/admin` and writes
client modules under `apps/web-antd/src/api`.

Regenerate the client only when backend admin Swagger changed or the generator
itself changed. After generation, run:

```bash
pnpm check:type --filter=@vben/web-antd
```

If generated output changes without a matching Swagger or generator change,
inspect the local dependency install before committing the diff.

## Admin Change Checklist

- View, table, form, or route copy: edit only the owning admin files and run
  `pnpm check:type --filter=@vben/web-antd`.
- API request or response shape: update backend protobuf and Swagger first, run
  `pnpm api:gen`, then run the admin type check.
- Shared package change under `packages` or `internal`: run the targeted app
  type check first, then wider lint or unit checks before publishing.

## Verification

```bash
pnpm check:type --filter=@vben/web-antd
pnpm test:unit
pnpm lint
```

Use the targeted type check for routine template edits. Run the wider checks
before publishing changes that touch shared packages or generated clients.

## Structure

- `apps/web-antd`: admin application entry point.
- `apps/web-antd/src/views`: feature pages and CRUD screens.
- `apps/web-antd/src/api`: generated API clients.
- `packages` and `internal`: shared UI, request, lint, build, and type config.
- `scripts/api-gen`: Swagger-to-TypeScript generation script.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Generated API client flow: `../docs/generated-artifacts.md`
- Environment configuration: `../docs/environment.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
