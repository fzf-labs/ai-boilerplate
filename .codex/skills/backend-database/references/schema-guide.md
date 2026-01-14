# 数据库表设计参考

## 必需字段

```sql
id uuid DEFAULT gen_random_uuid() NOT NULL,
created_at timestamp with time zone NOT NULL,
updated_at timestamp with time zone NOT NULL,
deleted_at timestamp with time zone
```

## 表名前缀

| 前缀 | 模块 |
|------|------|
| `sys_` | 系统管理 |
| `user_` | 用户相关 |
| `ai_` | AI 功能 |
| `mall_` | 电商 |
| `dict_` | 字典数据 |
| `file_` | 文件管理 |

## 字段命名

- 使用 `snake_case`
- 外键：`{table}_id`
- 时间：`_at` 或 `_time` 后缀

## 索引命名

- 普通索引：`{table}_{column}_idx`
- 唯一索引：`{table}_{column}_idx` (UNIQUE)
- 主键：`{table}_pkey`

## 常用类型

| 用途 | 类型 |
|------|------|
| 主键 | `uuid DEFAULT gen_random_uuid()` |
| 短文本 | `varchar(N)` |
| 长文本 | `text` |
| 时间戳 | `timestamp with time zone` |
| 金额 | `numeric(10,2)` |
| JSON | `jsonb` |

## SQL 模板

```sql
CREATE TABLE public.{table} (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    -- 业务字段
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);

COMMENT ON TABLE public.{table} IS '表描述';
COMMENT ON COLUMN public.{table}.id IS 'id';

ALTER TABLE ONLY public.{table} ADD CONSTRAINT {table}_pkey PRIMARY KEY (id);
CREATE INDEX {table}_{column}_idx ON public.{table} USING btree ({column});
```
