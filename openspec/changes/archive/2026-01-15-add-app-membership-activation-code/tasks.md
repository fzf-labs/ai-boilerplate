## 1. Implementation
- [x] Add ActivateMembershipByCode RPC and messages in api/app/v1/mall_activation_code.proto with auth header metadata.
- [x] Define activation code error reasons in api/app/v1/error_reason.proto and regenerate API bindings.
- [x] Implement activation flow in internal/service/app_v1_membership_activate_membership_by_code.go.
- [x] Wire MallActivationCodeRepo and MallProductRepo into AppV1MembershipService if missing.
- [x] Update activation code status/user fields and user membership expiration as specified.

## 2. Verification
- [x] Run make api (or equivalent) to regenerate pb/swagger.
- [x] Run targeted tests if available (or document manual verification). Manual verification: redeem a valid code via `/app/v1/mall_activation_code/activate` and confirm activation code status/user fields plus membership expiration updates.
