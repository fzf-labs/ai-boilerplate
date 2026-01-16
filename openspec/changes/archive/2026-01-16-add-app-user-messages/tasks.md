## 1. Backend
- [x] 1.1 Add `user_message` table + indexes for user/category/unread queries.
- [x] 1.2 Generate GORM model/repo for the new table.
- [x] 1.3 Generate Protobuf APIs and edit for app endpoints (category counts, list, detail, mark read) and admin endpoints (create/send, list, detail, audience scope).
- [x] 1.4 Generate API server code and register new HTTP routes.
- [x] 1.5 Implement service logic for message delivery, list queries, unread counts, and read updates.

## 2. Admin Web
- [x] 2.1 Regenerate admin API client for new endpoints.
- [x] 2.2 Add admin pages: message list, create/send form, message detail; wire permissions and menu entry.

## 3. App
- [x] 3.1 Regenerate app API client for new endpoints.
- [x] 3.2 Add home header message button and message category page with unread counts.
- [x] 3.3 Add per-category list page and message detail page; mark read on detail view.

## 4. Validation
- [ ] 4.1 Add or update backend tests for unread counts, list filters, and mark read.
- [ ] 4.2 Run `go test ./...`.
- [ ] 4.3 Run `pnpm -C ai-boilerplate-admin test:unit` (if affected).
- [ ] 4.4 Run `pnpm -C ai-boilerplate-app lint`.
