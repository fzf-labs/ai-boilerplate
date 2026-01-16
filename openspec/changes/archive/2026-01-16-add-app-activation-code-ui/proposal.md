# Change: Add app activation code redemption UI and history

## Why
Users need to redeem membership activation codes from the My page membership card and review their successful redemption history in the app.

## What Changes
- Add My page membership card entries and a redemption screen to submit activation codes (trim whitespace, enforce max length 64).
- Add a redemption history page with infinite scroll showing code, activated time, membership type, and expiration/duration, accessible from the My page membership card.
- Add an authenticated app API to list successful activation code redemptions for the current user and capture membership change snapshot for display.
- Regenerate the app API client for the new endpoint.

## Impact
- Affected specs: app-membership
- Affected code: ai-boilerplate-app/src/pages/me/me.vue, ai-boilerplate-app/src/pages/vip/*, ai-boilerplate-app/src/pages.json, ai-boilerplate-app/src/api/v1, ai-boilerplate-backend/api/app/v1/mall_activation_code.proto, ai-boilerplate-backend/internal/service/app_v1_mallactivationcode_*.go, ai-boilerplate-backend/internal/data/*
