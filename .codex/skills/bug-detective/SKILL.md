---
name: bug-detective
description: Systematic bug detection and troubleshooting for Kratos Admin project. Use when users encounter bugs, errors, or issues requiring debugging in Go/Kratos backend, Vue frontend, PostgreSQL database, or Redis cache. Triggers on keywords like bug, debug, error, exception, issue, troubleshooting, not working, failure, or when users report unexpected behavior. Provides structured debugging workflows for HTTP/gRPC services, database queries, caching, permissions, and common Kratos framework errors.
---

# Bug Detective

## Core Debugging Workflow

Follow this systematic approach for all bug investigations:

```
1. Reproduce → 2. Check Logs → 3. Locate Code → 4. Analyze Cause → 5. Fix & Verify
```

## Backend Debugging (Go + Kratos)

### View Logs

```bash
# Start service with console output
cd ai-boilerplate-backend
make run

# Filter for errors
make run 2>&1 | grep -i "error\|panic\|fatal"

# Save logs to file
make run 2>&1 | tee logs/debug.log
```

**Log Levels:**
- `DEBUG` - Detailed debugging (recommended for development)
- `INFO` - General information
- `WARN` - Warning messages
- `ERROR` - Error messages
- `FATAL` - Fatal errors

**Configure log level:**
```yaml
# configs/config.development.yaml
logger:
  type: "zerolog"
  zap:
    level: "debug"
```

### Add Debug Logs

```go
import "github.com/go-kratos/kratos/v2/log"

log.Debugf("User ID: %d, Username: %s", user.ID, user.Username)
log.Infof("Processing request: %+v", req)
log.Errorf("Error occurred: %v", err)
```

### Breakpoint Debugging (VS Code)

**launch.json:**
```json
{
  "version": "0.2.0",
  "configurations": [{
    "name": "Launch Kratos Server",
    "type": "go",
    "request": "launch",
    "mode": "debug",
    "program": "${workspaceFolder}/ai-boilerplate-backend/cmd/server",
    "env": {"APP_ENV": "development"}
  }]
}
```

### Performance Profiling (pprof)

```bash
# CPU profile
go tool pprof http://localhost:8000/debug/pprof/profile?seconds=30

# Memory profile
go tool pprof http://localhost:8000/debug/pprof/heap

# View goroutines
go tool pprof http://localhost:8000/debug/pprof/goroutine
```

## Frontend Debugging (Vue)

### Browser Console

```javascript
// Log variables
console.log('Value:', variable)
console.table(arrayData)
console.trace()

// Watch data changes
watch(() => props.data, (newVal, oldVal) => {
  console.log('Data changed:', { newVal, oldVal })
}, { deep: true })
```

### Network Requests

```javascript
// Intercept responses
axios.interceptors.response.use(
  response => {
    console.log('Response:', response.config.url, response.data)
    return response
  },
  error => {
    console.error('Request failed:', error.config?.url, error.response?.data)
    return Promise.reject(error)
  }
)
```

## Database Debugging (PostgreSQL + GORM)

### Test Connection

```bash
psql -h 0.0.0.0 -p 5432 -U postgres -d ai_boilerplate
```

### Enable SQL Logging

```yaml
# configs/config.development.yaml
data:
  gorm:
    showLog: true
    tracing: true
```

### Debug GORM Queries

```go
// Print SQL
db.Debug().Where("id = ?", 1).First(&user)

// View generated SQL
sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
    return tx.Model(&User{}).Where("id = ?", 1).Find(&users)
})
fmt.Println(sql)

// Test connection
sqlDB, _ := db.DB()
if err := sqlDB.Ping(); err != nil {
    log.Fatal("Database connection failed:", err)
}
```

### Common GORM Errors

- **record not found** - Query returned no results (check with RowsAffected before First/Take)
- **invalid connection** - Database connection failed (check DSN and service status)
- **duplicate key value** - Unique constraint violation (check constraints and data)
- **relation does not exist** - Table not found (run migrations or check table name)

## Redis Debugging

### Connect to Redis

```bash
redis-cli -h 0.0.0.0 -p 6379 -a 123456

# View keys
KEYS ai-boilerplate:*

# Get value
GET key_name

# Check TTL
TTL key_name

# Monitor commands
MONITOR
```

### Debug Redis in Code

```go
// Test connection
if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
    log.Errorf("Redis connection failed: %v", err)
}

// View cache content
result, err := client.Do(ctx, client.B().Get().Key("cache_key").Build()).ToString()
log.Debugf("Cache content: %s", result)
```

## Common Issues Checklist

### API 404 Not Found

- [ ] Route registered in `api/admin/v1/*.proto`
- [ ] Service implemented in `internal/service/*.go`
- [ ] Run `make api` to regenerate code
- [ ] Check HTTP server is running on correct port (default: 8000)
- [ ] Verify route format: `/api.admin.v1.ServiceName/MethodName`

### API 500 Internal Server Error

- [ ] Check backend error logs in console
- [ ] Verify database connection in `configs/config.development.yaml`
- [ ] Validate request parameters against proto definitions
- [ ] Run `make wire` to regenerate dependency injection
- [ ] Run `make gorm` to regenerate GORM models

### Data Not Displaying

- [ ] Check API response in Network tab
- [ ] Verify response format is correct
- [ ] Check Vue component data binding
- [ ] Look for v-if conditions blocking render
- [ ] Verify permissions with v-access directive

### Permission Issues

- [ ] Check JWT token validity and expiration
- [ ] Verify user role and menu associations in database (sys_role, sys_menu tables)
- [ ] Match permission codes between frontend (v-access) and backend
- [ ] Confirm Authorization header format: `Bearer <token>`

### gRPC Issues

- [ ] gRPC service running on correct port (default: 9000)
- [ ] Proto files synced with `make api`
- [ ] Service registry configured (etcd/consul/nacos)
- [ ] Test with grpcurl

## Common Kratos Errors

### Wire Dependency Injection

**Error:** `wire: no provider found for ...`
**Solution:**
1. Check `cmd/server/wire.go` for wire.Build
2. Ensure all dependencies in ProviderSet
3. Run `make wire`

### Protobuf Compilation

**Error:** `protoc-gen-go: program not found`
**Solution:**
1. Run `make init` to install tools
2. Ensure `$GOPATH/bin` in PATH

### Database Migration

**Error:** `relation "xxx" does not exist`
**Solution:**
1. Import SQL: `psql -U postgres -d ai_boilerplate < doc/sql/xxx.sql`
2. Check database connection config
3. Verify table name (case-sensitive)

## Debugging Tools Reference

| Tool | Purpose | Usage |
|------|---------|-------|
| Chrome DevTools | Frontend debugging, network | F12 |
| Vue DevTools | Vue component state | Chrome extension |
| Postman/Apifox | API testing | HTTP/gRPC support |
| VS Code Debugger | Go breakpoint debugging | Configure launch.json |
| Delve (dlv) | Go CLI debugger | `dlv debug ./cmd/server` |
| pprof | Go performance profiling | Built-in tool |
| grpcurl | gRPC testing | CLI gRPC client |
| psql | PostgreSQL client | Database queries |

## Quick Diagnostic Commands

```bash
# Health check
curl http://localhost:8000/health

# Check service process
ps aux | grep server

# Check port usage
lsof -i :8000
lsof -i :9000

# Regenerate all code
cd ai-boilerplate-backend
make wire && make api && make gorm

# Clean and rebuild
make clean && make build

# Run code checks
make lint

# Test database connection
psql -h 0.0.0.0 -p 5432 -U postgres -d ai_boilerplate -c "SELECT 1"

# Test Redis connection
redis-cli -h 0.0.0.0 -p 6379 -a 123456 ping

# View Docker logs
docker ps
docker logs <container_id>
```

## Bug Report Template

When reporting bugs, include:

```markdown
## Problem Description
Brief description of the issue

## Steps to Reproduce
1. Step one
2. Step two
3. Step three

## Expected Result
What should happen

## Actual Result
What actually happens

## Error Messages
```
Paste complete error logs:
- Backend: console output from `make run`
- Frontend: browser console errors
- Network: API response details
```

## Environment
- OS: macOS / Linux / Windows
- Go version: `go version`
- Node version: `node -v`
- Browser: Chrome 120
- Database: PostgreSQL 14
- Redis: 6.0

## Related Code
```go
// Paste relevant code snippets
```

## Attempted Solutions
- Attempt 1: ...
- Attempt 2: ...
```

## Best Practices

**DO:**
- ✅ Always check logs first
- ✅ Reproduce the issue consistently
- ✅ Test fixes thoroughly
- ✅ Run code generation commands (wire/api/gorm) after changes

**DON'T:**
- ❌ Guess without checking logs
- ❌ Commit code without testing
- ❌ Delete useful debug logs
- ❌ Modify generated code (*.gen.go)
- ❌ Change proto files without running `make api`
- ❌ Modify dependency injection without running `make wire`
