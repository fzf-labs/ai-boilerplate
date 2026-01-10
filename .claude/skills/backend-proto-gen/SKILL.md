---
name: backend-proto-gen
description: 后端 Protobuf 定义生成技能。用于从 SQL 表结构自动生成 Proto 文件。触发场景：(1) 新表需要生成 API 定义 (2) 表结构变更后重新生成 Proto
allowed-tools: Bash, Read, Glob
---

# Backend Proto Gen (Step 3)

从 SQL 生成 Proto 定义。

## 命令

```bash
# Admin
cd ai-boilerplate-backend && make sqltopb admin {table}

# App
cd ai-boilerplate-backend && make sqltopb app {table}
```

## 生成文件

`api/{position}/v1/{table}.proto`

## 输出

```
## Step 3: Proto 生成
- 命令: make sqltopb {position} {table}
- 状态: ✅ 已生成
```
