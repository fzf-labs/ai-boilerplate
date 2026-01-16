## Context
The backend already supports activation code redemption via `/app/v1/mall_activation_code/activate`, but the app lacks a My page membership card entry, redemption UI, and a redemption history list. The user also wants a history page that shows activation code, activated time, membership type, and expiration/duration.

## Goals / Non-Goals
- Goals:
  - Add My page membership card entries to redeem activation codes and a separate history page.
  - Support infinite scroll for redemption history.
  - Provide an authenticated app API to list successful redemption records.
  - Preserve backend as the source of truth for redemption results and error messages.
- Non-Goals:
  - QR/scan activation codes.
  - Display failed or pending redemption attempts.
  - Admin-side changes.

## Decisions
- Add a new app RPC in `app/v1/mall_activation_code.proto` (e.g., `ListActivationCodeRedemptions`) with `page`/`pageSize` and a list of successful redemptions.
- Filter redemption records by `status=activated` and `user_id` from the activation code table, ordered by `activated_at` desc.
- Capture a membership change snapshot into `mall_activation_code.user_change` during activation (membership type, duration days, expiredAt) so history can show the fields requested by the app without relying on current membership state.
- Add a new app page for redemption history under `pages/vip/` and link to it from the My page membership card; use infinite scroll to load more pages.
- Activation code input trims whitespace and blocks submission when the normalized code exceeds 64 characters.

## Risks / Trade-offs
- If `user_change` is missing (older activations), history may need to fall back to product config or show partial data.

## Migration Plan
1. Add the new list RPC and types in the app proto and regenerate API bindings.
2. Update activation flow to populate `user_change` when a code is successfully redeemed.
3. Implement the list endpoint and wire repositories.
4. Generate the app API client and implement UI pages/entry.

## Open Questions
- None.
