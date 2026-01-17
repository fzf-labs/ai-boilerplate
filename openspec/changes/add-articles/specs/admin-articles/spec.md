## ADDED Requirements
### Requirement: Admin manage articles
The system SHALL allow authenticated admins to create, edit, list, and delete articles with title, optional summary, optional cover image, and Markdown content.

#### Scenario: Admin creates a draft article
- **WHEN** an admin submits a new article with title and Markdown content
- **THEN** the system creates the article in draft state
- **AND** the article is not visible to the app until published

#### Scenario: Admin updates an existing article
- **GIVEN** an article exists
- **WHEN** an admin updates its title, summary, cover, or Markdown content
- **THEN** the system persists the changes

#### Scenario: Admin lists articles with filters
- **WHEN** an admin queries the article list with paging and optional filters (status, keyword)
- **THEN** the system returns a paginated list ordered by publish time (or updated time) descending

### Requirement: Admin publish and unpublish articles
The system SHALL allow authenticated admins to publish and unpublish articles, controlling app visibility via publish status and publish time.

#### Scenario: Admin publishes an article
- **GIVEN** an article is in draft state
- **WHEN** an admin publishes the article (optionally providing a publish time)
- **THEN** the article becomes visible in the app article list and detail endpoints

#### Scenario: Admin unpublishes an article
- **GIVEN** an article is published
- **WHEN** an admin unpublishes the article
- **THEN** the article no longer appears in the app article list
- **AND** the app detail endpoint does not expose the unpublished article
