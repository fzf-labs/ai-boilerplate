## ADDED Requirements
### Requirement: Redeem membership activation codes from the My page membership card
The system SHALL allow authenticated users to redeem a membership activation code from the My page membership card.

#### Scenario: Navigate to the redemption screen
- **GIVEN** the user is authenticated
- **WHEN** the user taps the activation code entry in the My page membership card
- **THEN** the system navigates to the activation code redemption screen

#### Scenario: Trim and validate the activation code
- **GIVEN** the user enters a code with leading or trailing whitespace
- **WHEN** the user submits the activation code
- **THEN** the system trims whitespace before validation and submission
- **AND** the system blocks submission if the normalized code exceeds 64 characters

#### Scenario: Redemption succeeds
- **GIVEN** the activation endpoint returns success
- **WHEN** the user submits the activation code
- **THEN** the system navigates to the membership detail page
- **AND** the membership detail reflects the updated membership and expiration

#### Scenario: Redemption fails
- **GIVEN** the activation endpoint returns an error
- **WHEN** the user submits the activation code
- **THEN** the system displays the backend error message

### Requirement: Provide activation code redemption history
The system SHALL provide an authenticated app endpoint to list successful activation code redemptions for the current user.

#### Scenario: List redemption records
- **GIVEN** the user is authenticated
- **WHEN** the user requests redemption history with page and pageSize
- **THEN** the system returns records containing activation code, activatedAt, membershipType, expiredAt, and durationDays

#### Scenario: No redemption records
- **GIVEN** the user has no successful redemptions
- **WHEN** the user requests redemption history
- **THEN** the system returns an empty list

### Requirement: Show redemption history in the app
The system SHALL provide a redemption history page accessible from the My page membership card with infinite scrolling.

#### Scenario: Open redemption history
- **WHEN** the user taps the redemption history entry in the My page membership card
- **THEN** the system navigates to the redemption history page

#### Scenario: Load more history
- **GIVEN** more redemption records exist
- **WHEN** the user reaches the end of the list
- **THEN** the system fetches the next page and appends results
