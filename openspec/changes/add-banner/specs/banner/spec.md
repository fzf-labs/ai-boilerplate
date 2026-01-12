## ADDED Requirements
### Requirement: Banner management (admin)
The system SHALL provide admin APIs to create, update, delete, enable/disable, and list tenant-scoped banners with basic filters.

#### Scenario: Admin lists banners by position
- **WHEN** an admin requests a banner list with a position filter
- **THEN** the system returns only banners belonging to the tenant, ordered by sort

### Requirement: Banner delivery (app)
The system SHALL provide app APIs to list active banners for a given position and platform, filtering by status and time window, scoped by tenant.

#### Scenario: App requests active banners
- **WHEN** the app requests banners for position "home"
- **THEN** the system returns only enabled banners within the time window for the tenant, ordered by sort

### Requirement: Home banner endpoint compatibility
The system SHALL return banner table data from the existing app home banner list endpoint using position "home".

#### Scenario: Home banner list
- **WHEN** the app calls GET /app/v1/home/banner/list
- **THEN** the system returns the same data set as the banner list for position "home"
