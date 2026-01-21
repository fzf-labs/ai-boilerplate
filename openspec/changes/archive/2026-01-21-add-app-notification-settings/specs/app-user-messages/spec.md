## ADDED Requirements
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
