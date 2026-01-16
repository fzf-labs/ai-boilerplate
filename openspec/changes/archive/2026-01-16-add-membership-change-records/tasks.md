## 1. Database
- [ ] 1.1 Add `user_membership_change` table SQL with required fields and indexes.
- [ ] 1.2 Regenerate GORM models/DAO/repo for the new table.

## 2. Admin API
- [ ] 2.1 Add admin proto for membership change list (by user_id, paging, fields).
- [ ] 2.2 Regenerate admin API bindings and Swagger artifacts.
- [ ] 2.3 Implement list handler + repo query (order by created_at desc).

## 3. Change Capture
- [ ] 3.1 Record membership changes on activation code redemption.
- [ ] 3.2 Record membership changes on order payment success (membership product only).
- [ ] 3.3 Record membership changes on admin membership updates.

## 4. Admin UI
- [ ] 4.1 Add "会员变更记录" action in user list operation menu.
- [ ] 4.2 Add modal UI to display records for selected user (columns: source_type, source_id, before/after membership_type, before/after expired_at, duration_days, remark, created_at).

## 5. Validation
- [ ] 5.1 Manual: trigger change via activation code and confirm admin records appear.
- [ ] 5.2 Manual: trigger change via admin update and confirm records appear.
- [ ] 5.3 Manual: trigger order payment success (membership product) and confirm records appear.
