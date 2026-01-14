---
name: backend-codeing
description: 后端业务逻辑开发技能。触发场景：(1) 实现 Service 业务逻辑 (2) 定义业务错误码 (3) 注册 HTTP Server (4) 编写 Data 层代码 (5) 表结构变更后实现逻辑 (6) 后端 CRUD 功能开发
---

# Backend Codeing (Step 6)

实现 Service 业务逻辑、定义错误码并注册到 HTTP Server。

## 可编辑文件

- `internal/service/{position}_v1_{table}_*.go` - Service 方法实现
- `internal/data/{table}.go` - Data 层代码
- `internal/server/http.go` - Server 注入
- `api/{position}/v1/error_reason.proto` - 错误码定义

## 禁止编辑

- `internal/data/gorm/**/*.gen.go`
- `api/**/*.pb.go`

## 流程

1. 如需自定义错误码，编辑 `error_reason.proto` 并执行 `make api`
2. 实现 Service 方法
3. 检查 `internal/server/http.go` 是否已注入
4. 如未注入，添加参数和注册代码

## 参考文档

- **错误码定义**: 见 [references/error-code.md](references/error-code.md)
- **Server 注入**: 见 [references/server-injection.md](references/server-injection.md)
- **Service 实现**: 见 [references/service-examples.md](references/service-examples.md)

## 输出

```
## Step 6: 业务逻辑
- 错误码: ✅ 已定义 / ⏭️ 使用默认
- Service: ✅ 已实现
- Server 注入: ✅ 已注入
```
