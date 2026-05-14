---
name: backend-api-gen
description: 后端 API 代码生成技能。用于从 Proto 文件生成 Go 代码（pb.go、grpc.pb.go、http.pb.go 等）。触发场景：(1) Proto 文件修改后 (2) 需要重新生成 API 代码
allowed-tools: Bash, Read, Glob
---

# Backend API Gen (Step 5)

生成 API 代码和 Service 桩。

## 命令

```bash
cd ai-boilerplate-backend && make api
cd ai-boilerplate-backend && make pbtocode
```

## 生成文件

**API：**
- `api/{position}/v1/{table}.pb.go`
- `api/{position}/v1/{table}_http.pb.go`

**Service：**
- `internal/service/{position}_v1_{table}.go`
- `internal/service/{position}_v1_{table}_{method}.go`

## 输出

```
## Step 5: API + Service 生成
- 命令: make api && make pbtocode
- 状态: ✅ 已生成
```
