---
name: backend-audit
description: 后端开发前置审计技能。用于验证开发前置条件、审计现有工件状态、确定开发起点。触发场景：(1) 开始后端开发任务前 (2) 检查表/API 工件是否存在 (3) 确定从哪个步骤开始开发
allowed-tools: Read, Glob, Grep, mcp__dbhub__search_objects
---

# Backend Audit (Step 0)

后端开发流程入口，收集输入、审计状态、确定起点。

## 输入参数

| 参数 | 说明 | 示例 |
|------|------|------|
| position | API 位置 | `admin` / `app` |
| table_name | 表名 | `user` / `user,role` |
| goal | 开发目标 | CRUD / 自定义过滤 / 业务 RPC |

## 工件检查

| Step | 路径 |
|------|------|
| 1 | `doc/sql/ai_boilerplate/{table}.sql` |
| 2 | `internal/data/gorm/ai_boilerplate_model/{table}.gen.go` |
| 3 | `api/{position}/v1/{table}.proto` |
| 5 | `internal/service/{position}_v1_{table}.go` |

## 输出

```
## 审计结论
- position: {value}
- table_name: {value}
- 建议从 Step {N} 开始
```
