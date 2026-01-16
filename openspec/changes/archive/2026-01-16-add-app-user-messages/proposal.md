# Change: Add app user messages with admin delivery

## Why
Users need an in-app message center for transaction/system/service notifications, and admins need to create and deliver those messages to specific users or segments.

## What Changes
- Add a new app user message data model and app APIs for category unread counts, message lists, message detail, and marking messages read.
- Add admin APIs and UI for creating and delivering user messages to all users, segments, or specific users.
- Add app UI entry (home header button) plus message categories, list, and detail pages.

## Impact
- Affected specs: app-user-messages, admin-user-messages.
- Affected code: backend schema/proto/service layers, admin UI (ai-boilerplate-admin), app UI (ai-boilerplate-app), and generated API clients.
