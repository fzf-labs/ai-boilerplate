# Change: Add “My Orders” to the app

## Why
Users can already place and pay orders (e.g., VIP purchase flow), but they cannot view their historical orders in the app. This change adds a “我的订单” entry under the “我的” tab so users can browse their own orders and inspect details.

## What Changes
- Add a “我的订单” entry in the “我的” page, gated by login.
- Add an order list page with paging and basic status filtering, powered by existing app endpoints:
  - `GET /app/v1/mall_order/user/orders`
  - `GET /app/v1/mall_order/order/info`
- Add an order detail page to display core order/payment fields.

## Impact
- Affected specs: `app-orders` (new capability)
- Backend: no changes expected (reuse existing mall order APIs)
- App: new pages under `ai-boilerplate-app/src/pages/` and a new entry in `src/pages/me/me.vue`

## Non-Goals
- Order operations: cancel, refund, confirm receipt, reorder, invoice, after-sales workflows.
- Admin-side order management changes.

## Open Questions
- None (confirmed: show all `productType`; filters use 全部/待付款/待发货/待收货/已完成/已取消/已退款).
