---
name: backend-audit
description: 后端开发前置审计技能。用于验证开发前置条件、审计现有工件状态、确定开发起点。触发场景：(1) 开始后端开发任务前 (2) 检查表/API 工件是否存在 (3) 确定从哪个步骤开始开发
allowed-tools: Read, Glob, Grep, mcp__dbhub__search_objects
---

# Backend Audit (Step 0)

后端开发流程入口，收集输入、审计状态、确定起点。

## 完整开发流程

| Step | 技能 | 说明 | 工件路径 |
|------|------|------|----------|
| 0 | backend-audit | 审计现状 | - |
| 1 | backend-database | 数据库表设计 | `doc/sql/ai_boilerplate/{table}.sql` |
| 2 | backend-gorm | GORM 代码生成 | `internal/data/gorm/ai_boilerplate_model/{table}.gen.go` |
| 3 | backend-proto-gen | Proto 生成（必须用 sqltopb） | `api/{position}/v1/{table}.proto` |
| 4 | backend-proto-edit | Proto 编辑（可选） | 同上 |
| 5 | backend-api-gen | API 代码生成 | `api/{position}/v1/{table}.pb.go` |
| 6 | backend-codeing | 业务逻辑实现 | `internal/service/{position}_v1_{table}.go` |
| 7 | backend-quality | 质量检查 | - |

## 输入参数

| 参数 | 说明 | 示例 |
|------|------|------|
| position | API 位置 | `admin` / `app` |
| table_name | 表名 | `user` / `user,role` |
| goal | 开发目标 | CRUD / 自定义过滤 / 业务 RPC |

## 审计检查项

按顺序检查以下工件是否存在：

| Step | 检查项 | 路径 | 检查方式 |
|------|--------|------|----------|
| 1 | SQL 文件 | `doc/sql/ai_boilerplate/{table}.sql` | Glob |
| 1 | 数据库表 | public.{table} | mcp__dbhub__search_objects |
| 2 | GORM Model | `internal/data/gorm/ai_boilerplate_model/{table}.gen.go` | Glob |
| 3 | Proto 文件 | `api/{position}/v1/{table}.proto` | Glob |
| 5 | API pb.go | `api/{position}/v1/{table}.pb.go` | Glob |
| 5 | Service 文件 | `internal/service/{position}_v1_{table}.go` | Glob |

## 审计逻辑

```
如果 SQL 文件不存在 → 从 Step 1 开始
如果 GORM Model 不存在 → 从 Step 2 开始
如果 Proto 文件不存在 → 从 Step 3 开始
如果 API pb.go 不存在 → 从 Step 5 开始
如果 Service 文件不存在 → 从 Step 6 开始
否则 → 检查是否需要修改现有代码
```

## 输出格式

```
## 审计结论

### 输入参数
- position: {value}
- table_name: {value}
- goal: {value}

### 工件状态
| Step | 工件 | 状态 |
|------|------|------|
| 1 | SQL 文件 | ✅/❌ |
| 2 | 数据库表 | ✅/❌ |
| 3 | GORM Model | ✅/❌ |
| 4 | Proto 文件 | ✅/❌ |
| 5 | API pb.go | ✅/❌ |
| 6 | Service 文件 | ✅/❌ |

### 建议
- 建议从 Step {N} 开始
- 需要执行的技能: {skill_list}
```
