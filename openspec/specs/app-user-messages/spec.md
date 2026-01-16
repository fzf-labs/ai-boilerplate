# app-user-messages Specification

## Purpose
TBD - created by archiving change add-app-user-messages. Update Purpose after archive.
## Requirements
### Requirement: App message center entry and categories
The system SHALL provide an app message entry in the home header that navigates to a message center showing categories for transaction, system, and service messages with unread counts per category for the current user.

#### Scenario: User opens message center
- **WHEN** the user taps the home header message entry
- **THEN** the app displays the three message categories with unread counts for each category

### Requirement: Category message list
The system SHALL provide a paginated message list for a selected category, ordered by most recent first, and indicate read/unread status for each item.

#### Scenario: User views a category list
- **WHEN** the user selects a category
- **THEN** the app shows a list of messages in that category ordered by sent time descending

### Requirement: Message detail and read state
The system SHALL display message details and mark the message as read when the detail view is opened.

#### Scenario: User opens a message detail
- **GIVEN** the message is unread
- **WHEN** the user opens the detail view
- **THEN** the message is marked as read and the category unread count decreases

### Requirement: Message media and deeplink
The system SHALL support optional cover image, summary, and deeplink for each message and render them in list/detail views when present.

#### Scenario: Message has optional fields
- **GIVEN** a message includes a cover image or deeplink
- **WHEN** the user views the list or detail view
- **THEN** the app displays the cover and provides a link action to navigate

