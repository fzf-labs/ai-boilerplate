# Change: Add banner management and app delivery APIs

## Why
Provide a single source of truth for banner content so admin can manage banners and app can consume active banners.

## What Changes
- Add a banner table (`banner`) with tenant scoping and fields for title, image, link, position, schedule, status, sort, and platform.
- Add admin CRUD and status endpoints for banners.
- Add app banner list endpoint with active filtering by time/status/position/platform.
- Replace hardcoded home banner data with banner table data (default position: home).

## Impact
- Affected specs: banner (new)
- Affected code: backend DB schema, gorm, admin/app proto & services, app home banner service
