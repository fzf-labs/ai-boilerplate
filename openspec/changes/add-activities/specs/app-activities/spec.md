## ADDED Requirements

### Requirement: App Discover activities tab
The app SHALL provide a “活动” tab in the Discover page that lists active activities with paging, ordered by configured sort order, and accessible without login.

#### Scenario: User switches to the activities tab
- **WHEN** a user opens Discover and switches to “活动”
- **THEN** the app fetches and displays the current active activities

#### Scenario: No active activities
- **GIVEN** there are no active activities
- **WHEN** a user opens the “活动” tab
- **THEN** the app displays an empty state

#### Scenario: Activities are ordered
- **GIVEN** multiple active activities exist with different sort values
- **WHEN** a user opens the “活动” tab
- **THEN** the app displays activities ordered by sort ascending (and stable ordering for equal sort)

### Requirement: Activity navigation
The app SHALL navigate users according to each activity’s link type and link URL.

#### Scenario: Activity opens an external link
- **GIVEN** an activity has an external link URL
- **WHEN** a user taps the activity
- **THEN** the app opens the URL in an in-app webview

#### Scenario: Activity opens an internal path
- **GIVEN** an activity has an internal app path
- **WHEN** a user taps the activity
- **THEN** the app navigates to the target page within the app
