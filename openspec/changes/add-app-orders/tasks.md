## 1. Spec / UX
- [x] 1.1 Confirm “我的订单” scope (product types + filters)
- [x] 1.2 Confirm the minimal fields shown in list and detail

## 2. App (Pages)
- [x] 2.1 Add `pages/orders/list` (paged list + empty/loading states)
- [x] 2.2 Add `pages/orders/detail` (order info + payment info presentation)
- [x] 2.3 Ensure pages require login (router strategy + in-page guard)

## 3. App (Entry)
- [x] 3.1 Add “我的订单” entry under `src/pages/me/me.vue`
- [ ] 3.2 (Optional) Link from VIP payment result page to the order detail page

## 4. Validation
- [x] 4.1 Run `pnpm -C ai-boilerplate-app lint`
- [ ] 4.2 Run `pnpm -C ai-boilerplate-app type-check` (currently fails due to pre-existing typecheck issues in the app + wot-design-uni source types)
