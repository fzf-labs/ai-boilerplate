# Environment Configuration

Keep repository defaults safe for local development. Put real secrets,
production endpoints, signing assets, provisioning files, and account-specific
IDs in local environment files or platform secret stores.

## Shared Rules

- Do not commit production credentials, private keys, keystores, provisioning
  profiles, Firebase service files, or real third-party app secrets.
- Prefer local loopback defaults for templates that talk to the backend.
- Document any required environment value in the owning template README when it
  becomes mandatory for local development.
- Keep generated clients pointed at the same backend path shape used by the
  runtime request layer.

## Local Setup Checklist

1. Pick one template and install dependencies only inside that directory.
2. Copy or create the documented local config file for that template.
3. Keep loopback defaults for local-only browser or simulator flows.
4. Use a LAN IP or emulator host alias when a physical device or emulator cannot
   reach `127.0.0.1`.
5. Run the template verification command before changing environment-sensitive
   behavior.
6. Move real credentials into local files, CI secrets, or platform secret
   storage before release.

## Template Files

| Template | Environment Files | Notes |
| --- | --- | --- |
| Backend | `ai-boilerplate-backend/configs/config.example.yaml`, `configs/config.development.yaml` | Set PostgreSQL, Redis, JWT, WeChat, registry, tracing, and push-service values per environment. |
| Admin | `ai-boilerplate-admin/apps/web-antd/.env*` | `VITE_GLOB_API_URL` controls the admin API base path. Development defaults to `/api`. |
| Uni-app | `ai-boilerplate-uniapp/env/.env*` | `VITE_SERVER_BASEURL` controls the backend URL. WeChat mini-program builds can override it per release channel. |
| PC web | `ai-boilerplate-pc` Vite env files when added | Keep public browser endpoints distinct from admin-only APIs. |
| Electron | `ai-boilerplate-electron` env files when added | Keep package identity in `electron-builder.yml`; keep runtime secrets outside the repo. |
| Tauri | `ai-boilerplate-tauri` env files when added | Keep renderer values public and move native secrets to platform-specific secure storage. |
| Chrome extension | `ai-boilerplate-chrome-extension/public/manifest.json` and Vite env files when added | Request browser permissions only when a feature needs them. |
| iOS | `ai-boilerplate-ios/AIBoilerplate/Configurations/XCConfigs` after Tuist generation | `API_BASE_URL` defaults to the local backend path. Update bundle IDs and endpoints before shipping. |
| Android | `ai-boilerplate-android/local.properties` | Set `BASE_URL` for a reachable backend. Android emulators reach the host at `10.0.2.2`. |

## Local Backend URLs

Use these defaults unless a template README says otherwise:

| Client | Local Backend URL |
| --- | --- |
| Admin web | `/api` through a local proxy or mock server |
| Uni-app H5 | `http://127.0.0.1:8000` |
| Android emulator | `http://10.0.2.2:8000/api/` |
| iOS simulator | `http://127.0.0.1:8000/api` |

When testing on a physical device, replace loopback addresses with a LAN IP that
the device can reach.

## Override Scope

Treat endpoint overrides as local unless a release task explicitly says
otherwise:

| Scope | Where To Put It | Verification |
| --- | --- | --- |
| Local browser or simulator check | Template env file, local config, or `local.properties` | Run the template verification command after the value changes. |
| Physical device debugging | LAN IP or platform-approved test domain in local config | Confirm the device can reach the backend before changing code. |
| CI or preview environment | CI secrets or deployment configuration | Document the variable name, not the secret value. |
| Production release | Product-owned secret store or platform signing/deployment system | Stop for approval before committing product-specific identifiers. |

Do not commit a production endpoint simply to make a local build pass. Starter
defaults should stay reusable and safe for new projects.

## Device URL Rules

- Browser-based local clients can use `/api` or `127.0.0.1` when the dev server
  proxies to the backend.
- Android emulators use `10.0.2.2` to reach services running on the host.
- iOS simulators can usually use `127.0.0.1`; physical iOS devices need a LAN
  IP or deployed backend.
- Mini-program targets must use platform-approved domains for non-local
  releases.
- Do not commit production domains to starter defaults unless the product owner
  explicitly makes that endpoint part of the template.

## Quick Reference

| Template | Runtime Value | Local Default | Keep Out Of Git |
| --- | --- | --- | --- |
| Backend | Database, Redis, JWT, WeChat, tracing, push service | `configs/config.development.yaml` | Real credentials and production service endpoints. |
| Admin | `VITE_GLOB_API_URL` | `/api` | Production admin hostnames and tokens. |
| Uni-app | `VITE_SERVER_BASEURL`, platform app IDs | `http://127.0.0.1:8000` for H5 | Real mini-program IDs, platform secrets, private domains. |
| PC web | Vite public env values when added | Project-specific local value | Private API keys or admin-only endpoints. |
| Electron | Renderer env and package identity | Local development defaults | Signing identities, notarization credentials, update server secrets. |
| Tauri | Renderer env and native secure storage values | Local development defaults | Native secrets and signing credentials. |
| Chrome extension | Manifest permissions and public env values | Manifest V3 defaults | Store credentials and unnecessary browser permissions. |
| iOS | `API_BASE_URL`, bundle ID, signing config | `http://127.0.0.1:8000/api` | Provisioning profiles, certificates, real team IDs. |
| Android | `BASE_URL` in `local.properties` | `http://10.0.2.2:8000/api/` | Keystores, signing passwords, Firebase service files. |

## Secret Handling

- Backend JWT secrets should be at least 32 characters and unique per audience.
- Third-party API keys belong in local config files or deployment secrets.
- Mobile signing files and desktop signing identities must stay outside the
  repository.
- If a value is needed in CI, add it through CI secrets and document the
  variable name without exposing the value.
