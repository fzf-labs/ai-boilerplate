# Change: Add Articles (admin publish, app display)

## Why
The app already has “内容(content)” endpoints and a detail page, but backend responses are currently stubbed demo data and there is no admin workflow to edit/publish real articles. This change adds a minimal CMS-style article feature so admins can publish articles and the app can list and display them.

## What Changes
- Add a persistent article/content table and repositories in backend.
- Implement `Home.GetContentList` and `Home.GetContentDetail` to read published articles from DB (replace stub demo data) and make them accessible without login.
- Add admin APIs to create/edit/publish/unpublish/list/delete articles.
- Add an app article list entry/page that uses the existing content list endpoint; reuse existing content detail page for article display.
- Use Markdown as the source content format for authoring; render safely for app display.

## Impact
- Affected specs: `admin-articles`, `app-articles`
- Backend: new DB table + new admin API surface; app `home` content endpoints switch from stubbed data to real DB data.
- Admin web: new “文章管理” screens + generated API client for new endpoints.
- App: add article list UX and entry point; existing `src/pages/content/detail.vue` becomes part of the article flow.

## Non-Goals
- Comments, likes interaction, sharing, full-text search, multi-language, attachments management.
- Category taxonomy and advanced permissions beyond existing admin RBAC patterns.

## Open Questions
- None.
