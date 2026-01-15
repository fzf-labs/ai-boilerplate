## ADDED Requirements
### Requirement: Activate membership with activation code
The system SHALL provide an authenticated app activation-code endpoint to redeem a membership activation code for the current user.

#### Scenario: Redeem a valid code
- **GIVEN** the activation code exists and is not disabled, refunded, expired, or already activated
- **AND** the activation code is within its valid time window
- **AND** the linked product is a membership product with a membershipType and duration_days in product_config.membership
- **WHEN** the user submits the activation code
- **THEN** the system updates the activation code status to activated, sets activated_at, and binds user_id
- **AND** the system updates user_membership to the product membershipType
- **AND** the system sets expired_at to now + duration_days, or extends from current expired_at if it is later than now

#### Scenario: Code is invalid
- **GIVEN** the activation code does not exist
- **WHEN** the user submits the activation code
- **THEN** the system returns an activation-code-not-found error

#### Scenario: Code is not redeemable
- **GIVEN** the activation code is disabled, refunded, expired, or already activated
- **WHEN** the user submits the activation code
- **THEN** the system returns an activation-code-not-redeemable error

#### Scenario: Product config is missing membership data
- **GIVEN** the activation code links to a product without membership config
- **WHEN** the user submits the activation code
- **THEN** the system returns a product-config-invalid error
