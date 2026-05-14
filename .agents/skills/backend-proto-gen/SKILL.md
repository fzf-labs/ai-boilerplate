---
name: backend-proto-gen
description: 后端 Protobuf API 定义生成技能（必选步骤）。基于 sqltopb 从数据库表自动生成 Proto 文件，禁止手动创建。触发场景：(1) 新建表后生成 API (2) 开发后端 CRUD 功能 (3) 需要 Proto/protobuf 文件 (4) 创建 gRPC/HTTP 接口定义 (5) 后端开发流程中的 Step 3
allowed-tools: Bash, Read, Glob
---

# Backend Proto Gen (Step 3)

从 SQL 生成 Proto 定义。

## 核心规范

### 强制要求

1. **必须使用 sqltopb 命令生成 Proto 文件**
   - 禁止手动创建或自定义 .proto 文件
   - 所有 Proto 定义必须通过 `make sqltopb` 命令从 SQL 表结构自动生成

2. **文件生命周期管理**
   - Proto 文件只需创建一次，创建后不再重复创建
   - 如需修改，使用 `backend-proto-edit` 技能进行更改
   - 表结构变更时，重新运行 sqltopb 会更新现有文件

### 禁止行为

- ❌ 手动编写 .proto 文件内容
- ❌ 直接创建新的 .proto 文件
- ❌ 绕过 sqltopb 命令生成 Proto

## 命令

```bash
# Admin
cd ai-boilerplate-backend && make sqltopb admin {table}

# App
cd ai-boilerplate-backend && make sqltopb app {table}
```

## 生成文件

`api/{position}/v1/{table}.proto`

## 工作流程

1. 确认数据库表已存在
2. 检查 Proto 文件是否已存在
   - 不存在：运行 sqltopb 生成
   - 已存在：如需修改，使用 `backend-proto-edit` 技能
3. 验证生成结果

## 输出

```
## Step 3: Proto 生成
- 命令: make sqltopb {position} {table}
- 文件: api/{position}/v1/{table}.proto
- 状态: ✅ 已生成 / ⏭️ 已存在(跳过)
```
