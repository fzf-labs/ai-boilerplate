---
name: backend-codeing
description: Backend development skill for this repo. Use when implementing backend features, generating CRUD code, or writing service/data business logic (including after table schema changes).
allowed-tools: Read, Edit, Glob, Grep
---

# Backend Codeing (Step 6)

实现 Service 业务逻辑并注册到 HTTP Server。

## 可编辑文件

- `internal/service/{position}_v1_{table}_*.go`
- `internal/data/{table}.go`
- `internal/server/http.go`

## 禁止编辑

- `internal/data/gorm/**/*.gen.go`
- `api/**/*.pb.go`

## 流程

1. 实现 Service 方法
2. 检查 `internal/server/http.go` 是否已注入
3. 如未注入，添加参数和注册代码

## Server 注入

见 [references/server-injection.md](references/server-injection.md)

## Service 实现

见 [references/service-examples.md](references/service-examples.md)

## 输出

```
## Step 6: 业务逻辑
- Service: ✅ 已实现
- Server 注入: ✅ 已注入
```
