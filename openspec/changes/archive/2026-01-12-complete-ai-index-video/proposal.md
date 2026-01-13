# Change: Complete AI index video endpoints

## Why
AI index video endpoints are currently stubbed, blocking admin video features and leaving responses inconsistent with other AI index modules.

## What Changes
- Implement AI index video create/info/list/update/update-status/delete using existing video record storage with tenant/admin scoping.
- Ensure AI index video responses map record fields consistently with other AI index APIs.

## Impact
- Affected specs: ai-index-backend (new)
- Affected code: AI index video service methods
