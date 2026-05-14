---
name: backend-database
description: PostgreSQL 数据库表设计技能。触发场景：(1) 创建新表 (2) 修改现有表 (3) 设计表关系 (4) 查询表结构
allowed-tools: Read, Write, Glob, mcp__dbhub__execute_sql, mcp__dbhub__search_objects
---

# Backend Database (Step 1)

设计并创建数据库表。

## 流程

1. 收集需求
2. 生成 SQL（参考 `references/schema-guide.md`）
3. 保存到 `doc/sql/ai_boilerplate/{table}.sql`
4. 执行 `mcp__dbhub__execute_sql`

## 必需字段

每表必须包含：`id`, `created_at`, `updated_at`, `deleted_at`

## 参考

详细规范见 [references/schema-guide.md](references/schema-guide.md)

## 输出

```
## Step 1: 数据库表
- 表名: {table}
- 状态: ✅ 已创建
```
