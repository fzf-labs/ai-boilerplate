## ADDED Requirements

### Requirement: Admin manage activities
The system SHALL allow authenticated admins to create, edit, list, and delete activities that will be displayed in the app Discover page, including title, poster image, link type/url, scheduling (start/end time), sort order, and enable/disable status.

#### Scenario: Admin creates an enabled activity
- **WHEN** an admin creates an activity with title, poster image, link type/url, and status enabled
- **THEN** the system persists the activity
- **AND** the activity becomes eligible to be shown in the app Discover “活动” tab (subject to scheduling rules)

#### Scenario: Admin schedules an activity
- **GIVEN** an activity exists
- **WHEN** an admin sets a start time and/or end time window
- **THEN** the activity is only eligible for display during the configured time window

#### Scenario: Admin disables an activity
- **GIVEN** an activity is enabled
- **WHEN** an admin disables the activity
- **THEN** the activity is no longer eligible to be shown in the app Discover “活动” tab

### Requirement: Activities are scoped to Discover placement
The system SHALL scope “活动管理” to only manage activities intended for the app Discover page.

#### Scenario: Admin lists activities
- **WHEN** an admin opens the activities list
- **THEN** the system returns only activities belonging to the Discover activities dataset
