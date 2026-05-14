---
name: backend-gorm
description: 后端 GORM 代码生成技能。用于验证数据库表存在性并生成 GORM 模型、DAO、Repo 代码。触发场景：(1) 新建表后生成 GORM 代码 (2) 表结构变更后重新生成 (3) 检查 GORM 工件状态
allowed-tools: Bash, Read, Glob
---

# Backend GORM (Step 2)

生成 GORM 代码（Model、DAO、Repo）。

## 命令

```bash
cd ai-boilerplate-backend && make gorm
```

## 生成文件

- `internal/data/gorm/ai_boilerplate_model/{table}.gen.go`
- `internal/data/gorm/ai_boilerplate_dao/{table}.gen.go`
- `internal/data/gorm/ai_boilerplate_repo/{table}.gen.go`

## 输出

```
## Step 2: GORM 生成
- 命令: make gorm
- 状态: ✅ 已生成
```
