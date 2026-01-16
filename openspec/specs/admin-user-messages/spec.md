# admin-user-messages Specification

## Purpose
TBD - created by archiving change add-app-user-messages. Update Purpose after archive.
## Requirements
### Requirement: Admin create and deliver app messages
The system SHALL allow admins to create and deliver app user messages with category (transaction/system/service), title, summary, cover image, content, and optional deeplink, targeting all users, a filtered segment (e.g., membership level or activity status), or specific user IDs.

#### Scenario: Admin sends to all users
- **WHEN** an admin submits a message with audience scope "all users"
- **THEN** the system creates unread message deliveries for every eligible user

#### Scenario: Admin sends to specific users
- **WHEN** an admin submits a message with a list of user IDs
- **THEN** only those users receive unread message deliveries

### Requirement: Admin view message deliveries
The system SHALL allow admins to list and view sent messages with their audience scope and delivery metadata.

#### Scenario: Admin reviews messages
- **WHEN** an admin opens the message list
- **THEN** the system shows each message with category, audience scope, and sent time

