## Context
The app already renders HTML content via `rich-text` (`src/pages/content/detail.vue`) and backend currently returns hard-coded demo content from `AppV1HomeService`. We need a real persisted article model, admin authoring workflow, and safe rendering.

## Goals / Non-Goals
- Goals:
  - Support admin authoring (draft) and publishing (visible to app).
  - Render article content on app with the existing `rich-text` approach.
  - Keep the API surface minimal and consistent with existing `admin/v1/*` CRUD patterns.
- Non-Goals:
  - Full CMS features (comments/likes workflows, taxonomy, WYSIWYG blocks, search).

## Decisions
- Decision: Store article `content` as Markdown (source of truth).
  - Why: Markdown is easy to author/review/version, and matches the chosen content format for v1.
  - Alternatives considered:
    - HTML-only storage: higher XSS risk and worse editing ergonomics.
    - Plain text only: insufficient for “文章” use cases (images/formatting, links).
- Decision: Publish gate via `status` + `publish_time`.
  - Why: Enables drafts and scheduled/past publish ordering while keeping app filtering simple.
 - Decision: App article browsing is public (no login required).
   - Why: Articles are marketing/knowledge content and should be accessible to guests.

## Data Model (Proposed)
Single table (name TBD, e.g. `home_content` or `article`) with at least:
- `id` (int64/bigserial or uuid; must match app-facing ID type)
- `title`, `summary`, `cover_image`, `content_markdown` (Markdown)
- `content_html` (optional; derived from Markdown via server-side rendering, if we keep app consuming HTML)
- `status` (draft/published/unpublished) and `publish_time`
- `tags` (optional; e.g. `text[]` or `jsonb`)
- `is_recommend`, `is_hot` (optional)
- `view_count`, `like_count` (optional; `view_count` can increment on detail view)
- `created_at`, `updated_at`, `deleted_at` (soft delete consistent with existing tables)

## API Surface (Proposed)
- App:
  - Keep using existing endpoints:
    - `GET /app/v1/home/content/list` (published-only, paginated)
    - `GET /app/v1/home/content/detail` (published-only, by id)
  - Authentication:
    - No `Authorization` required for article browsing (header may be present but is optional)
- Admin (new):
  - `POST /admin/v1/article/create`
  - `POST /admin/v1/article/update`
  - `POST /admin/v1/article/update/status` (publish/unpublish)
  - `POST /admin/v1/article/delete`
  - `GET /admin/v1/article/info`
  - `GET /admin/v1/article/list`

## Risks / Trade-offs
- XSS/content safety risk (Markdown rendering) → Mitigation: disable raw HTML in Markdown and sanitize the generated HTML before returning to the app.
- Backfill risk: none (new table), but existing app users may see empty list until content is published.

## Migration Plan
1. Add new table and generate repos.
2. Implement app read APIs against DB (behind publish filter).
3. Add admin CRUD + publish endpoints.
4. Ship admin UI, then enable app entry point.

## Open Questions
- Exact Markdown-to-HTML policy (library choice + allowlist).
