## ADDED Requirements
### Requirement: App article list
The system SHALL provide an app article list that displays published articles with paging, ordered by publish time descending, and accessible without login.

#### Scenario: User opens the article list
- **WHEN** a user navigates to the article list
- **THEN** the app shows a paginated list of published articles with title, optional summary, and optional cover image

#### Scenario: Guest opens the article list
- **GIVEN** the user is not logged in
- **WHEN** the user navigates to the article list
- **THEN** the system returns the published article list successfully

#### Scenario: No published articles
- **GIVEN** no articles are published
- **WHEN** a user opens the article list
- **THEN** the app displays an empty state

### Requirement: App article detail
The system SHALL provide an app article detail view that renders the article Markdown content and metadata, and is accessible without login.

#### Scenario: User opens a published article
- **GIVEN** an article is published
- **WHEN** a user opens the article detail
- **THEN** the app renders the Markdown content and shows title, publish time, and optional cover image

#### Scenario: Guest opens a published article
- **GIVEN** an article is published
- **AND** the user is not logged in
- **WHEN** the user opens the article detail
- **THEN** the system returns the article detail successfully

#### Scenario: Article is not published
- **GIVEN** an article is draft or unpublished
- **WHEN** a user opens the article detail
- **THEN** the system returns a not-found (or access denied) error
