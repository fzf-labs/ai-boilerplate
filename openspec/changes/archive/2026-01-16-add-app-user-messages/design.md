## Context
The app needs a message center with categorized notifications (transaction, system, service) and per-category unread counts. Admins need to create and deliver those messages to all users, specific users, or filtered segments.

## Goals / Non-Goals
- Goals:
  - Provide app message entry in the home header, category list with unread counts, per-category message list, and message detail.
  - Mark messages as read when the detail view is opened.
  - Provide admin creation and delivery for app messages, supporting all users, segmented users, and explicit user lists.
- Non-Goals:
  - Push notifications, realtime delivery, or websocket-based updates.
  - Two-way chat or user replies.
  - Complex segmentation rules beyond basic filters (e.g., membership level or activity status).

## Decisions
- Data model: introduce a `user_message` table storing per-user deliveries with message content and read state.
  - Core fields: id, user_id, category, title, summary, cover_url, content, link_url, sent_at, read_at, created_by, created_at, updated_at, deleted_at.
  - Indexes: (user_id, category, read_at, sent_at) for unread counts and list queries.
- Delivery: admin create/send endpoints fan out messages into per-user rows. For large audiences (all/segment), use background job batching to avoid timeouts.
- API shape:
  - App: category unread counts, list by category (paged), message detail, mark read.
  - Admin: create/send message, list messages, view message details, and track audience scope.

## Risks / Trade-offs
- Per-user storage duplicates message content; this is simpler to query but increases storage. Acceptable for initial release and can be refactored later if needed.
- Broadcast fan-out can be expensive; mitigate with background jobs and batch inserts.

## Migration Plan
1. Add `user_message` table and indexes.
2. Generate GORM models/repos, Protobuf APIs, and server code.
3. Implement admin and app endpoints.
4. Add admin UI and app UI.

## Open Questions
- Define the exact "activity" filter for segmented delivery.
- Final API route naming (e.g., `/app/v1/user_message/*` vs `/app/v1/message/*`).
