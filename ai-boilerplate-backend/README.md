# AI Boilerplate Backend

## Configuration Security

The server loads `configs/config.${APP_ENV}.yaml` at startup. Keep real
`config.*.yaml` files out of git; only `*.example.yaml` templates are tracked.

Environment policy:

| APP_ENV | Template | Policy |
| --- | --- | --- |
| `local` | `configs/config.local.example.yaml` | Local-only debugging; pprof and wildcard CORS are allowed. |
| `development` | `configs/config.example.yaml` | Shared development baseline; pprof is off, CORS is limited, logging/recovery/rate limit/metrics are enabled. |
| `staging` | `configs/config.staging.example.yaml` | Pre-production baseline; use real staging secrets and restricted origins. |
| `production` | `configs/config.production.example.yaml` | Hardened baseline; pprof is off, CORS must be explicit, logging/recovery/rate limit/metrics are required. |

`testing` is accepted as a staging-compatible runtime environment for the
existing Docker build workflow.

Create runtime config from a template:

```bash
cp configs/config.production.example.yaml configs/config.production.yaml
```

Inject secrets through your deployment platform rather than committing them.
The process supports these environment variable overrides before validation:

```bash
AI_BOILERPLATE_DB_DSN
AI_BOILERPLATE_REDIS_ADDR
AI_BOILERPLATE_REDIS_PASSWORD
AI_BOILERPLATE_JWT_ADMIN_SECRET
AI_BOILERPLATE_JWT_KID_SECRET
AI_BOILERPLATE_JWT_PARENT_SECRET
AI_BOILERPLATE_WX_DEFAULT_GZH_APP_ID
AI_BOILERPLATE_WX_DEFAULT_XCX_APP_ID
AI_BOILERPLATE_BAIDU_PUSH_API_KEY
AI_BOILERPLATE_BAIDU_PUSH_SECRET_KEY
AI_BOILERPLATE_CORS_ORIGINS
```

Production startup fails fast when required secrets are missing, JWT secrets are
shorter than 32 characters, placeholders are still present, CORS uses wildcards,
pprof is enabled, rate limit/metrics/recovery/logging are disabled, or data
endpoints still point at local hosts.

Validation examples:

```bash
go test ./...
APP_ENV=production go run ./cmd/server -conf ./configs
```

The second command must fail clearly if `config.production.yaml` still contains
template placeholders or required production secrets are missing.
