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

### Requirement: Notification settings by message category
The system SHALL provide a notification settings view that lets users enable or disable push notifications per message category, including activity, service, interaction, and direct-message categories and any additional configured categories.

#### Scenario: User opens notification settings
- **WHEN** the user opens the notification settings view
- **THEN** the app displays toggle controls for activity, service, interaction, and direct-message categories with the current on/off states

#### Scenario: User changes a category setting
- **WHEN** the user disables a category and saves the setting
- **THEN** the system persists the updated preference for that category

### Requirement: Push delivery respects category settings
The system SHALL suppress push notifications for categories that the user has disabled while keeping in-app messages available in the message center.

#### Scenario: Category disabled for push
- **GIVEN** the user has disabled push notifications for the system category
- **WHEN** a system message is generated
- **THEN** the system does not send a push notification and the message remains visible in the message center

