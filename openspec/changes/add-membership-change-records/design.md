## Context
Membership changes are currently applied in multiple places (activation code redemption, admin updates, order/payment flows), but there is no unified, queryable change log. Admin operators need to see the history per user.

## Goals / Non-Goals
- Goals:
  - Persist membership change records with before/after membership type and expiry, duration days, source type, source id, remark, and timestamps.
  - Record changes triggered by: order payment success, activation code redemption, admin membership updates.
  - Provide admin list API and UI entry from user list actions to view records.
- Non-Goals:
  - App endpoints or app UI for change records.
  - Complex analytics or reporting beyond basic list.

## Decisions
- Add `user_membership_change` table with fields:
  - `user_id`, `source_type` (order/activation_code/admin), `source_id` (order id or activation code code),
    `before_membership_type`, `after_membership_type`, `before_expired_at`, `after_expired_at`, `duration_days`,
    `remark`, `created_at`.
- Only create a record when membership changes or expiry extends (including extension-only changes).
- `source_id` for activation code uses the activation code **code** (not id).
- Admin API: list records by user id with paging and default sort by created_at desc.
- Admin UI: add a user list action "会员变更记录" that opens a modal consistent with existing user detail patterns.

## Risks / Trade-offs
- Membership updates triggered by orders may not be centralized; hooking requires locating or adding change capture at the point where user_membership is updated.

## Migration Plan
1. Add the new table and regenerate GORM artifacts.
2. Add admin proto for listing membership change records, generate API stubs.
3. Implement service + repo query and wire to admin UI modal.
4. Add record creation hooks for order payment success, activation code redemption, and admin membership updates.
