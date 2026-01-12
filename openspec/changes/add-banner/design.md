## Context
We need a general banner capability that can be managed in admin and delivered to app clients, replacing the current hardcoded app home banners.

## Goals / Non-Goals
- Goals:
  - Centralize banner content in a dedicated table.
  - Provide admin CRUD and status control endpoints.
  - Provide app list endpoint filtering by position/platform/status/time window.
  - Keep the existing app home banner endpoint functional by reading from the banner table.
- Non-Goals:
  - Banner analytics or impression tracking.
  - A/B testing or personalization.

## Decisions
- Table name: `banner` (no `sys_` prefix).
- Core fields: tenant_id, title, image_url, link_url, link_type, position, platform, sort, status, start_time, end_time.
- App list returns only active banners (status enabled, within time window), scoped by tenant and ordered by sort.
- Home banner endpoint maps to position `home` and reuses the same list logic.

## Risks / Trade-offs
- Time window filtering relies on server timezone; use UTC in storage and compare in UTC.
- Adding a new table requires migration in environments without an auto-migrate workflow.

## Migration Plan
1. Add `sys_banner` SQL and deploy migration.
2. Generate GORM model and API stubs.
3. Implement admin and app services.
4. Update app home banner endpoint to read from the new table.

## Open Questions
- Confirm the exact set of link types/platform values and their validation rules.
