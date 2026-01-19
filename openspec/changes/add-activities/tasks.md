## 1. Spec / UX
- [x] 1.1 Confirm activity card fields in app (title + poster image; optional subtitle not supported)
- [x] 1.2 Confirm `linkType` handling rules (external vs internal)

## 2. Backend (DB + Repos)
- [x] 2.1 Create `activity` table SQL under `ai-boilerplate-backend/doc/sql/ai_boilerplate/`
- [x] 2.2 Implement `ActivityRepo` (GORM)

## 3. Backend (App API)
- [x] 3.1 Add `app/v1/activity.proto` (list) and regenerate protobuf + swagger artifacts
- [x] 3.2 Implement `Activity.ListActivities` to return only active activities with paging, ordering, and time-window filtering
- [x] 3.3 Make activities list public (add to app auth whitelist)

## 4. Backend (Admin API)
- [x] 4.1 Add `admin/v1/activity.proto` (CRUD + update status) and regenerate protobuf + swagger artifacts
- [x] 4.2 Implement admin service handlers in `ai-boilerplate-backend/internal/service/`
- [ ] 4.3 Register sys_menu/sys_api and permission codes for “活动管理” as needed (requires DB-side menu/button permissions configuration)

## 5. Admin Web (Activities management)
- [x] 5.1 Generate admin TS API client for new `activity` endpoints
- [x] 5.2 Add router entry under 内容管理: `/content/activity` → “活动管理”
- [x] 5.3 Implement list page backed by `GET /admin/v1/activity/list`
- [x] 5.4 Implement create/edit form backed by `/admin/v1/activity/create` + `/admin/v1/activity/update`
- [x] 5.5 Implement enable/disable backed by `/admin/v1/activity/update/status`
- [x] 5.6 Add `v-access:code` permissions (codes: `content:activity:create|update|enable|disable|delete`)

## 6. App (Discover tab)
- [x] 6.1 Generate app TS API client for new activity list endpoint
- [x] 6.2 Add “活动” tab in `ai-boilerplate-app/src/pages/discover/index.vue`
- [x] 6.3 Fetch activities via `/app/v1/activity/list`
- [x] 6.4 Render list/empty/loading states; clicking routes via `linkType/linkURL` (reuse home banner routing logic)

## 7. Validation
- [x] 7.1 Run `go test ./...` in `ai-boilerplate-backend`
- [ ] 7.2 Run `pnpm -C ai-boilerplate-admin lint && pnpm -C ai-boilerplate-admin test:unit` (test:unit fails due to pre-existing jsdom/localStorage env)
- [ ] 7.3 Run `pnpm -C ai-boilerplate-app lint && pnpm -C ai-boilerplate-app type-check` (type-check fails due to pre-existing wot-design-uni + app typing issues)
