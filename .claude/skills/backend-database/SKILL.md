---
name: backend-database
description: >-
  PostgreSQL database table design for ai-boilerplate project. Use when users need
  to design new database tables, add fields, create indexes, or query table
  structures. Triggers on: (1) Create new table structure (2) Modify existing
  table (3) Design table relationships (4) Create index strategy (5) Choose field
  types (6) Naming convention consultation (7) Database schema questions
---

# Database Table Design

PostgreSQL table design guide with interactive workflows for the ai-boilerplate project.

## Core Principles

Every table MUST follow these mandatory requirements:

1. **Required Fields** - Every table must have:
   - `id uuid DEFAULT gen_random_uuid() NOT NULL` - Primary key
   - `created_at timestamp with time zone NOT NULL` - Creation timestamp
   - `updated_at timestamp with time zone NOT NULL` - Update timestamp
   - `deleted_at timestamp with time zone` - Soft delete timestamp (nullable)

2. **Table Naming** - Use module prefix (sys_, user_, ai_, mall_, etc.)

3. **Field Order** - Follow strict order:
   - id (first)
   - Business fields
   - status/sort fields
   - Timestamp fields (created_at, updated_at, deleted_at - last)

4. **Comments Required** - Every table and field must have COMMENT

5. **Index Strategy**:
   - Foreign key fields must have index
   - Unique constraint fields must have unique index
   - High-frequency query fields should have index

## Project Structure

- SQL files: `ai-boilerplate-backend/doc/sql/ai_boilerplate/`
- MCP tools: `mcp__dbhub__execute_sql`, `mcp__dbhub__search_objects`

## Workflows

### Workflow A: Create New Table

**Step 1: Gather Requirements & Analyze References**

Ask user:
- Table purpose and module (sys, user, ai, mall, etc.)
- Required business fields and their purpose
- Multi-tenant support needed? (tenant_id)
- Foreign key relationships?
- Special needs (status, sort, soft delete)

Simultaneously:
- Use Glob to view existing tables in `ai-boilerplate-backend/doc/sql/ai_boilerplate/*.sql`
- Find tables with same module prefix as reference
- Analyze similar tables for field naming and index patterns

**Step 2: Generate SQL Design**

Generate complete SQL including:
- CREATE TABLE with all required fields
- COMMENT for table and all fields
- Primary key constraint
- Necessary indexes (foreign keys, unique, query optimization)

Internal validation (don't show to user):
- ✓ Table name uses correct module prefix
- ✓ Contains required fields (id, created_at, updated_at, deleted_at)
- ✓ All fields have COMMENT
- ✓ Foreign key fields have indexes
- ✓ Unique fields have unique indexes
- ✓ Field order is correct
- ✓ Data types are appropriate

Show user:
1. Complete SQL code
2. Design explanation (data types, index strategy, relationships)
3. Wait for confirmation

**Step 3: Save & Execute**

After user confirmation:

1. **Save SQL file**
   - Create file at `ai-boilerplate-backend/doc/sql/ai_boilerplate/{table_name}.sql`
   - Filename must match table name exactly

2. **Auto-execute SQL (if MCP available)**
   - Check for `mcp__dbhub__execute_sql` tool
   - If available:
     - Execute complete SQL (CREATE TABLE + COMMENT + INDEX)
     - Verify table created using `mcp__dbhub__search_objects`
     - Display execution results
   - If not available:
     - Prompt user to manually execute SQL file

### Workflow B: Modify Existing Table

**Step 1: Read & Analyze**

1. Use Read to load `ai-boilerplate-backend/doc/sql/ai_boilerplate/{table_name}.sql`
2. Analyze existing table structure
3. Ask user what modifications are needed

**Step 2: Generate Modification Plan**

Generate based on user needs:
- ALTER TABLE statements (add/modify fields, indexes)
- COMMENT statements for new fields
- Updated complete CREATE TABLE statement (for updating SQL file)

Show user:
1. ALTER TABLE statements
2. Modification explanation
3. Wait for confirmation

**Step 3: Update & Execute**

After user confirmation:

1. **Update SQL file**
   - Replace original file content with complete CREATE TABLE

2. **Auto-execute ALTER TABLE (if MCP available)**
   - Check for `mcp__dbhub__execute_sql` tool
   - If available:
     - Execute ALTER TABLE statements
     - Execute COMMENT for new fields
     - Execute new indexes
     - Verify modifications successful
   - If not available:
     - Prompt user to manually execute ALTER TABLE statements

### Workflow C: Query Table Structure

When user asks about table structure or field information:

1. **Prefer MCP (if available)**
   - Use `mcp__dbhub__search_objects` to query tables, fields, indexes
   - Get real-time database structure

2. **Fallback: Read SQL file**
   - Use Read to load `ai-boilerplate-backend/doc/sql/ai_boilerplate/{table_name}.sql`
   - Parse SQL file content

3. **Display results**
   - Table structure overview
   - Field list (name, type, comment)
   - Index list
   - Relationships

## Naming Conventions

### Table Prefixes

| Prefix | Module | Examples |
|--------|--------|----------|
| `sys_` | System management | sys_admin, sys_role, sys_menu |
| `user_` | User related | user, user_membership |
| `ai_` | AI features | ai_chat_message, ai_image_record |
| `mall_` | E-commerce | mall_order, mall_product |
| `wx_gzh_` | WeChat Official Account | wx_gzh_user, wx_gzh_menu |
| `wx_xcx_` | WeChat Mini Program | wx_xcx_user |
| `dict_` | Dictionary data | dict_type, dict_data |
| `file_` | File management | file_config, file_data |
| `mail_` | Email | mail_account, mail_template |
| `sms_` | SMS | sms_channel, sms_template |
| `self_` | Self-owned apps | self_app, self_app_release |

### Field Naming Rules

- Use `snake_case`
- Foreign keys: `{related_table}_id` format (e.g., `user_id`, `tenant_id`)
- Boolean: avoid `is_` prefix, use semantic names (e.g., `pinned`, `public_status`)
- Time fields: use `_at` or `_time` suffix (e.g., `created_at`, `payment_time`)

### Index Naming Rules

- Regular index: `{table}_{column}_idx`
- Unique index: `{table}_{column}_idx` (with UNIQUE)
- Composite index: `{table}_{column1}_{column2}_idx`
- Primary key: `{table}_pkey`

## SQL File Format

Each table has one `.sql` file, named exactly as the table name.

### Standard SQL Template

```sql
CREATE TABLE public.{table_name} (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    -- Business fields here
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);

-- Table comment
COMMENT ON TABLE public.{table_name} IS 'Table description';

-- Field comments (required for every field)
COMMENT ON COLUMN public.{table_name}.id IS 'id';
COMMENT ON COLUMN public.{table_name}.status IS 'Status';
COMMENT ON COLUMN public.{table_name}.created_at IS 'Created at';
COMMENT ON COLUMN public.{table_name}.updated_at IS 'Updated at';
COMMENT ON COLUMN public.{table_name}.deleted_at IS 'Deleted at';

-- Primary key constraint
ALTER TABLE ONLY public.{table_name} ADD CONSTRAINT {table_name}_pkey PRIMARY KEY (id);

-- Indexes (add as needed)
CREATE INDEX {table_name}_{column}_idx ON public.{table_name} USING btree ({column});
CREATE UNIQUE INDEX {table_name}_{column}_idx ON public.{table_name} USING btree ({column});
```


## Common Field Patterns

### Required Fields (every table must have)

```sql
id uuid DEFAULT gen_random_uuid() NOT NULL,     -- Primary key
created_at timestamp with time zone NOT NULL,   -- Creation time
updated_at timestamp with time zone NOT NULL,   -- Update time
deleted_at timestamp with time zone             -- Soft delete time
```

### Multi-tenant Field

```sql
tenant_id character varying(64) NOT NULL,       -- Tenant ID
```

### Status Field

```sql
status integer DEFAULT 1 NOT NULL,              -- 1=enabled, -1=disabled
status smallint DEFAULT 1 NOT NULL,             -- Same, more space efficient
```

### Sort Field

```sql
sort integer DEFAULT 0 NOT NULL,                -- Sort value
sort bigint NOT NULL,                           -- Large range sorting
```

### User Association

```sql
admin_id character varying(64) NOT NULL,        -- Admin ID
user_id character varying(64) NOT NULL,         -- User ID
```


## Common Data Types

| Purpose | Type | Notes |
|---------|------|-------|
| Primary key | `uuid DEFAULT gen_random_uuid()` | Auto-generated UUID |
| Short text | `character varying(N)` | N is max length |
| Long text | `text` | No length limit |
| Integer | `integer` / `bigint` / `smallint` | Choose by range |
| Boolean | `boolean` | true/false |
| Timestamp | `timestamp with time zone` | With timezone |
| Money | `numeric(10,2)` | Exact decimal |
| JSON | `jsonb` | Binary JSON, indexable |
| Float | `double precision` | Inexact decimal |

**Common length reference:**
- Username/nickname: `varchar(64)` or `varchar(100)`
- Phone: `varchar(20)`
- Email: `varchar(128)`
- Password hash: `varchar(128)`
- URL: `varchar(512)` or `varchar(2048)`
- Title: `varchar(256)`
- Notes: `varchar(500)` or `text`


## Index Strategy

### Required Indexes

- Primary key auto-creates unique index
- Foreign key fields: `{table}_{fk_column}_idx`
- Unique constraint fields: `{table}_{column}_idx` (UNIQUE)

### Common Index Scenarios

```sql
-- Foreign key indexes (improve join performance)
CREATE INDEX {table}_user_id_idx ON public.{table} USING btree (user_id);
CREATE INDEX {table}_tenant_id_idx ON public.{table} USING btree (tenant_id);

-- Unique indexes (ensure data uniqueness)
CREATE UNIQUE INDEX {table}_username_idx ON public.{table} USING btree (username);
CREATE UNIQUE INDEX {table}_phone_idx ON public.{table} USING btree (phone);

-- Composite unique index (combined uniqueness)
CREATE UNIQUE INDEX {table}_type_key_idx ON public.{table} USING btree (type, key);

-- Time index (query optimization)
CREATE INDEX {table}_created_at_idx ON public.{table} USING btree (created_at);

-- Status index (high-frequency queries)
CREATE INDEX {table}_status_idx ON public.{table} USING btree (status);
```

### Index Design Principles

- Foreign key fields must have indexes
- Unique constraint fields must have unique indexes
- High-frequency query fields should have indexes
- Time range query fields should have indexes
- Avoid too many indexes (affects write performance)


## Table Relationship Patterns

### One-to-Many Relationship

```sql
-- Parent table
CREATE TABLE public.dict_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type character varying(100) NOT NULL,
    ...
);

-- Child table
CREATE TABLE public.dict_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type character varying(100) NOT NULL,  -- References dict_type.type
    ...
);
CREATE INDEX dict_data_type_idx ON public.dict_data USING btree (type);
```

### Many-to-Many Relationship (JSON approach)

```sql
-- Role table stores menu ID array
CREATE TABLE public.sys_role (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "menuIds" jsonb,  -- Menu ID array ["id1", "id2", ...]
    ...
);
```

**Note:** This project prefers using jsonb to store related ID arrays rather than creating junction tables.


## Common Questions

### Q: How to handle enum types?

A: Use integer or varchar, specify enum values in comment:

```sql
status integer DEFAULT 1 NOT NULL,
-- COMMENT: Status(0=pending,1=processing,2=completed,3=cancelled)

product_type character varying(20) NOT NULL,
-- COMMENT: Product type(membership:membership,service:service,goods:goods)
```

### Q: When to use text vs varchar?

A:
- Fixed format with clear length limit: use `varchar(n)`
- User input long text, uncertain length: use `text`
- Prompts, content, descriptions: use `text`

### Q: What type for money fields?

A: Use `numeric(10,2)`, precise to cents:

```sql
original_amount numeric(10,2) NOT NULL,  -- Original price
discount_amount numeric(10,2) DEFAULT 0.00,  -- Discount amount
actual_amount numeric(10,2) NOT NULL,  -- Actual amount
```

### Q: How to design soft delete?

A: Use `deleted_at` field (nullable):

```sql
deleted_at timestamp with time zone  -- NULL=not deleted, value=deletion time
```

GORM automatically handles soft delete logic.

### Q: Should we create foreign key constraints?

A: This project **does not use database foreign key constraints**, instead:
- Maintain relationships at application layer
- Create regular indexes for foreign key fields (not FK constraints)
- Document relationships in comments


## MCP Tools Usage

### Available MCP Tools

1. **mcp__dbhub__execute_sql**
   - Execute SQL statements (CREATE TABLE, ALTER TABLE, COMMENT, INDEX)
   - For automated table creation and modification

2. **mcp__dbhub__search_objects**
   - Query database objects (schema, table, column, procedure, index)
   - For verifying table creation and querying table structure

### Usage Examples

```javascript
// Execute CREATE TABLE
mcp__dbhub__execute_sql({
  sql: "CREATE TABLE public.test_table (...)"
})

// Execute COMMENT
mcp__dbhub__execute_sql({
  sql: "COMMENT ON TABLE public.test_table IS 'Test table'"
})

// Verify table exists
mcp__dbhub__search_objects({
  object_type: "table",
  pattern: "test_table"
})

// Query all fields of table
mcp__dbhub__search_objects({
  object_type: "column",
  schema: "public",
  table: "test_table",
  detail_level: "full"
})
```

### When to Use MCP

**Recommended:**
- Development environment rapid iteration
- Need immediate SQL syntax validation
- Want to automate entire workflow
- Need quick table structure view

**Not recommended:**
- Production environment (should use migration scripts)
- Changes requiring approval process
- Complex data migration operations
- Operations involving sensitive data


## Design Checklist

When creating or modifying tables, ensure:

- [ ] Table name uses correct module prefix
- [ ] Uses UUID primary key: `id uuid DEFAULT gen_random_uuid() NOT NULL`
- [ ] Contains timestamp fields: `created_at`, `updated_at`, `deleted_at`
- [ ] If multi-tenant: add `tenant_id` field
- [ ] If status management: add `status` field
- [ ] Every field has COMMENT
- [ ] Table has COMMENT
- [ ] Foreign key fields have indexes
- [ ] Unique constraint fields have unique indexes
- [ ] High-frequency query fields have indexes
- [ ] Field order is correct (id → business fields → status/sort → timestamp fields)
- [ ] Data types are appropriate
- [ ] Index naming follows conventions


## Quick Reference

### Complete Flow for Creating New Table

1. Ask requirements + view existing tables
2. Generate SQL + internal validation
3. Show design + wait for confirmation
4. Save file + execute SQL (MCP)

### Complete Flow for Modifying Existing Table

1. Read existing table + ask modification needs
2. Generate ALTER TABLE + show plan
3. Update file + execute ALTER (MCP)

### Complete Flow for Querying Table Structure

1. Prefer MCP query
2. Fallback: read SQL file
3. Display table structure, fields, indexes

**Remember**: Think in English, respond to user in Chinese. Quality over speed—iterate until database table design are truly clear.
