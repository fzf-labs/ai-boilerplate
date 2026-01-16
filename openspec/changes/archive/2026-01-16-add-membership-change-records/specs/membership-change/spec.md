## ADDED Requirements
### Requirement: Record user membership changes
The system SHALL create a user membership change record when a user's membership type or expiry changes due to an order payment success, activation code redemption, or admin update.

#### Scenario: Membership change from order payment
- **GIVEN** a membership order payment succeeds
- **WHEN** the system applies the membership change to the user
- **THEN** the system records a change with source_type=order and source_id=order_id
- **AND** the record includes before/after membership_type, before/after expired_at, duration_days, remark, and created_at

#### Scenario: Membership change from activation code
- **GIVEN** a membership activation code is redeemed successfully
- **WHEN** the system applies the membership change to the user
- **THEN** the system records a change with source_type=activation_code and source_id=activation code
- **AND** the record includes before/after membership_type, before/after expired_at, duration_days, remark, and created_at

#### Scenario: Membership change from admin update
- **GIVEN** an admin updates a user's membership
- **WHEN** the membership change is applied
- **THEN** the system records a change with source_type=admin
- **AND** the record includes before/after membership_type, before/after expired_at, duration_days, remark, and created_at

### Requirement: Admin can list membership change records by user
The system SHALL provide an authenticated admin endpoint to list membership change records for a specified user with paging.

#### Scenario: List membership change records
- **GIVEN** the admin provides a user_id, page, and pageSize
- **WHEN** the admin requests the membership change list
- **THEN** the system returns records ordered by created_at desc
- **AND** each record includes source_type, source_id, before/after membership_type, before/after expired_at, duration_days, remark, and created_at

#### Scenario: No membership change records
- **GIVEN** the user has no membership change records
- **WHEN** the admin requests the list
- **THEN** the system returns an empty list
