## 1. Backend (DB + Repos)
- [x] 1.1 Create article table SQL under `ai-boilerplate-backend/doc/sql/ai_boilerplate/` (content fields + publish status/time)
- [x] 1.2 Apply schema to Postgres and run `godb ormgen` for the new table
- [x] 1.3 Add a data wrapper repo in `ai-boilerplate-backend/internal/data/` for the generated repo

## 2. Backend (App API)
- [x] 2.1 Implement `Home.GetContentList` to return only published articles with paging and ordering (publish_time desc)
- [x] 2.2 Implement `Home.GetContentDetail` to fetch one published article; optionally increment `view_count`
- [x] 2.3 Ensure error handling matches existing `pb.ErrorReason...` conventions
- [x] 2.4 Make article browsing public: remove auth requirement for `Home` content endpoints (proto OpenAPI + auth middleware)

## 3. Backend (Admin API)
- [x] 3.1 Add `admin/v1/article.proto` (CRUD + publish/unpublish) and regenerate protobuf + swagger artifacts
- [x] 3.2 Implement admin service handlers in `ai-boilerplate-backend/internal/service/` using generated repos
- [ ] 3.3 Update auth middleware and admin RBAC permission/menu data as needed for new endpoints
- [x] 3.4 Accept Markdown content on write and render/sanitize for app consumption (disable raw HTML in Markdown)

## 4. Admin Web
- [x] 4.1 Add admin TS API module for new `article` endpoints
- [x] 4.2 Add “文章管理” list page (filter by status/keyword) and editor modal/page (title/summary/cover/content/publish)
- [x] 4.3 Add access codes for create/update/publish/delete actions consistent with existing `v-access:code` usage

## 5. App
- [x] 5.1 Add an article list page that calls `/app/v1/home/content/list` and navigates to `/src/pages/content/detail.vue`
- [x] 5.2 Add an entry point (e.g., home section/button) to open the article list
- [x] 5.3 Render article content from Markdown (either client-side Markdown renderer or consume server-rendered/sanitized HTML)

## 6. Validation
- [x] 6.1 Run `go test ./...` in `ai-boilerplate-backend`
- [ ] 6.2 Run `pnpm -C ai-boilerplate-admin/apps/web-antd lint && pnpm -C ai-boilerplate-admin/apps/web-antd test:unit`
- [ ] 6.3 Run `pnpm -C ai-boilerplate-app lint && pnpm -C ai-boilerplate-app type-check`
