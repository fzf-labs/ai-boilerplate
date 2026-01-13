## ADDED Requirements
### Requirement: Admin refresh token
The system SHALL provide an admin refresh token API that accepts a refresh token and returns a new access token with expiration and refresh timestamps.

#### Scenario: Refresh succeeds
- **WHEN** a valid refresh token is submitted
- **THEN** the system returns a new access token, expiredAt, and refreshAt.
