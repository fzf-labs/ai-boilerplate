# Design: Activities domain (new DB table)

## Goal
Ship “活动” end-to-end (admin configurable + app Discover tab) with minimal backend change and maximum reuse of existing validated patterns.

## Key Decision
Introduce a first-class “活动(activity)” domain backed by a dedicated database table and new protobuf APIs (admin + app).

We still reuse proven patterns from the existing `banner` implementation:
- Activity list ordering + time-window filtering (start/end) mirrors `ListBanners`.
- Admin CRUD surface mirrors the structure of existing admin resources (Create/Update/UpdateStatus/Delete/Info/List).

## App Routing
Use the same routing semantics as home banners:
- If `linkType == "external"` OR `linkURL` is `http(s)://...`, open `/pages-fg/webview/index`.
- Otherwise treat `linkURL` as an internal path (allow `app://` prefix as a convenience).

## Data Model (proposed)
Create `public.activity` with the minimal fields needed by admin and app:
- `tenant_id` (for multi-tenant filtering, aligned with `banner`)
- `title`, `image_url`
- `link_type`, `link_url`
- `sort`, `status`
- `start_time`, `end_time`
- `created_at`, `updated_at`, `deleted_at`

## App Visibility Rules
The app “活动” tab only shows activities that are:
- `status = 1`
- within the configured time window (if provided)
Ordered by `sort ASC, created_at DESC`.
