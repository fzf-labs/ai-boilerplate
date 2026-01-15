# Change: Add app membership activation by code

## Why
Users need to redeem membership activation codes directly in the app without creating orders.

## What Changes
- Add an authenticated app API to redeem a membership activation code.
- Update activation code state and user membership based on the linked product config.
- Add app error reasons for activation code validation failures.

## Impact
- Affected specs: app-membership
- Affected code: api/app/v1/membership.proto, api/app/v1/error_reason.proto, internal/service/app_v1_membership_*.go, internal/data/mallactivationcode.go, internal/data/*repo wiring
