## 1. Backend API
- [ ] 1.1 Add a list-redemptions RPC and messages in `api/app/v1/mall_activation_code.proto` with paging and fields: code, activatedAt, membershipType, expiredAt, durationDays.
- [ ] 1.2 Regenerate app API bindings and Swagger artifacts for the new RPC.
- [ ] 1.3 Populate `mall_activation_code.user_change` with a membership snapshot (type, expiredAt, durationDays) during activation.
- [ ] 1.4 Implement the list-redemptions handler and repo query (filter by user_id + activated status, order by activated_at desc).

## 2. App API Client
- [ ] 2.1 Run app API generation to add client wrappers/types for the new list endpoint.

## 3. App UI
- [ ] 3.1 Add My page membership card entries for activation code redemption and redemption history.
- [ ] 3.2 Implement the activation code form with whitespace trimming, max length 64, login guard, and submit to `/app/v1/mall_activation_code/activate`.
- [ ] 3.3 On success, navigate to membership detail and refresh membership info; on failure, show backend error.
- [ ] 3.4 Create a redemption history page with infinite scroll; display code, activatedAt, membershipType, expiredAt, and durationDays.

## 4. Validation
- [ ] 4.1 Manual verification: redeem a valid code, confirm membership detail refreshes, and confirm the redemption record appears in history.
- [ ] 4.2 Manual verification: redeem an invalid/unredeemable code and confirm the backend error is shown.
