# Change: Add user membership change records

## Why
Operators need a unified audit trail of membership changes caused by orders, activation codes, and admin updates.

## What Changes
- Add a user membership change record table capturing before/after membership type and expiry, duration days, source type, source id, remark, and timestamps.
- Write change records when membership changes due to: order payment success, activation code redemption, and admin membership updates.
- Add admin APIs to list membership change records by user and view them from the user list action menu in the admin UI.
- App UI and endpoints are out of scope for this change.

## Impact
- Affected specs: membership-change (new)
- Affected code: ai-boilerplate-backend (schema + admin APIs + change hooks), ai-boilerplate-admin (user list action + modal)
