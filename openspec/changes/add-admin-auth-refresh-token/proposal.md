# Change: Add admin auth refresh token endpoint

## Why
Admin clients need a refresh flow to obtain a new access token without forcing re-login.

## What Changes
- Add SysAuthRefreshToken RPC with HTTP POST /admin/v1/sys_auth/refresh_token.
- Add request/response messages for refresh token exchange.
- Document admin auth refresh behavior in the spec delta.

## Impact
- Affected specs: admin-auth
- Affected code: ai-boilerplate-backend/api/admin/v1/sys_auth.proto
