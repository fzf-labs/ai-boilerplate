# Change: Add app notification settings

## Why
Users need control over which message categories send push notifications so they can reduce unwanted alerts.

## What Changes
- Add per-category push notification settings in the app (activity/service/interaction/direct), with future categories supported.
- Replace fixed boolean columns with JSON-based preference storage for extensibility (**BREAKING** API/schema change).
- Persist user preferences and apply them when delivering push notifications.
- Keep in-app message visibility unchanged when a category is disabled.

## Impact
- Affected specs: app-user-messages
- Affected code: app settings UI, message/push notification delivery logic, user notification preference storage
