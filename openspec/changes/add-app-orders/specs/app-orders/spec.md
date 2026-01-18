## ADDED Requirements
### Requirement: App “My Orders” entry
The system SHALL provide a “我的订单” entry point in the app “我的” tab that is accessible only to authenticated users.

#### Scenario: Guest taps “我的订单”
- **GIVEN** the user is not logged in
- **WHEN** the user taps “我的订单”
- **THEN** the app prompts the user to log in
- **AND** navigates to the login page after confirmation (or a short delay)

#### Scenario: Logged-in user opens “我的订单”
- **GIVEN** the user is logged in
- **WHEN** the user taps “我的订单”
- **THEN** the app navigates to the order list page

### Requirement: App order list
The system SHALL provide an order list page that shows the user’s orders with paging and basic status filtering.

#### Scenario: User opens the order list
- **GIVEN** the user is logged in
- **WHEN** the user opens the order list page
- **THEN** the app loads orders from `GET /app/v1/mall_order/user/orders` with paging
- **AND** the app displays an empty state when the list is empty

#### Scenario: User filters by status
- **GIVEN** the user is logged in
- **WHEN** the user selects a status filter (e.g., pendingPayment / completed)
- **THEN** the app reloads the list using the `status` parameter

### Requirement: App order detail
The system SHALL provide an order detail page that displays core order and payment fields for a single order.

#### Scenario: User views an order detail
- **GIVEN** the user is logged in
- **WHEN** the user opens an order detail page for an order id
- **THEN** the app loads the order from `GET /app/v1/mall_order/order/info`
- **AND** the app displays order status, amounts, payment status/time, and created time when available

#### Scenario: Order id is invalid
- **GIVEN** the user is logged in
- **WHEN** the user opens an order detail page with an invalid or missing order id
- **THEN** the app shows an error state and allows navigating back
