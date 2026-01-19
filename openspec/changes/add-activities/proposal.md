# Change: Add Activities (admin configure, app Discover tab)

## Why
运营需要可配置的“活动”入口；目前 App 的发现页只有“文章”tab，且管理后台没有针对“活动”的配置入口。

## What Changes
- Define “活动(activities)” as a curated list of marketing entries displayed in the app Discover page.
- Add a dedicated `activity` database table in the backend to store activities (title, poster image, link type/url, scheduling, sort, status).
- Add backend APIs:
  - Admin: `/admin/v1/activity/*` CRUD + enable/disable + list filters.
  - App: `/app/v1/activity/list` for Discover to fetch active activities with paging, ordered by sort.
- Admin web adds a dedicated “活动管理” screen (under 内容管理) backed by the new admin activity APIs.
- App Discover adds a new “活动” tab which lists the activities and routes users by `linkType/linkURL`.

## Impact
- Affected specs: `admin-activities`, `app-activities`
- Backend: new DB table + new admin/app API surface; app Discover becomes public for activities list
- Admin web: new route + views; generated API client for new endpoints
- App: update Discover UI to add the Activities tab and call the new activity list endpoint

## Non-Goals
- Activity detail CMS pages, registration/报名流程, coupons, analytics/埋点报表.
- Multi-slot placements beyond Discover activities.

## Open Questions
- Confirm `linkType` conventions: `external` opens webview, otherwise treat `linkURL` as an app path (supports `app://` prefix like home banners).
