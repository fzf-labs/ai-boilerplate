# Project Context

## Purpose
AI boilerplate monorepo providing a backend API plus admin web and mobile app
clients for an AI-enabled product (chat/image/audio/video, prompts), with
membership and e-commerce features, and integrations like WeChat and messaging.

## Tech Stack
- Backend: Go 1.24, Kratos v2, gRPC + grpc-gateway, Protobuf + OpenAPI/Swagger
- Data: PostgreSQL (primary), Redis (cache/queues), GORM + gorm/gen
- Infra: OpenTelemetry, Asynq, optional service registry (etcd/consul/nacos)
- Admin web: Vue 3 + TypeScript, Vite, Ant Design Vue (vben admin monorepo)
- Mobile app: uni-app (Vue 3), Vite, Pinia, wot-design-uni, alova, z-paging
- Tooling: pnpm workspaces, ESLint/Prettier, Vitest (admin)

## Project Conventions

### Code Style
- Go: gofmt, idiomatic Go; keep API handlers in `internal/service` and data
  access in `internal/data`; generated files keep `admin_v1_*` / `app_v1_*`
  naming.
- Frontend: follow repo lint configs (ESLint/Prettier in admin, ESLint in app);
  use Vue 3 composition API and TypeScript.

### Architecture Patterns
- Monorepo with git subtrees for `ai-boilerplate-backend`, `ai-boilerplate-admin`,
  and `ai-boilerplate-app`.
- Backend follows Kratos layering: service (transport handlers) and data (DB,
  repos, external integrations). GORM repos live under
  `internal/data/gorm/...` and are generated.
- API definitions live in Protobuf and are published as Swagger in
  `ai-boilerplate-backend/doc/swagger`.

### Testing Strategy
- Backend: `go test ./...` for unit/integration tests as needed.
- Admin web: Vitest unit tests (`pnpm test:unit`).
- App: lint + type checks (`pnpm lint`, `pnpm type-check`) before release.

### Git Workflow
- Subtrees are synced via `make subtree-pull-*` / `make subtree-push-*`.
- Default branch is `master` for subtree remotes.
- Prefer Conventional Commits when working in the admin frontend (commitlint).

## Domain Context
Core domains include users/admins/roles/menus, tenants, notices, notifications,
membership and activation codes, mall products/orders/payments, AI providers and
models, AI content records (chat/image/audio/video/write), and WeChat (GZH/mini
program) operations plus SMS/email services.

## Important Constraints
- Requires PostgreSQL and Redis (see `docker-compose.example.yml`).
- Config is managed via `ai-boilerplate-backend/configs/config.yaml` (example in
  `configs/config.example.yaml`).
- Node >= 20 and pnpm >= 9 for the app; Go 1.24 for the backend.

## External Dependencies
- PostgreSQL, Redis, optional etcd/consul/nacos registry
- OpenTelemetry collector (OTLP)
- WeChat Official Account / Mini Program (PowerWeChat)
- SMS/email providers and Baidu Push
