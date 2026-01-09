---
name: backend-codeing
description: Backend development skill for this repo. Use when implementing backend features, generating CRUD code, or writing service/data business logic (including after table schema changes).
---

# Backend Code Development Skill

This skill defines the **end-to-end backend workflow** from database tables to protobuf APIs and service/business logic.

## Strict Mode (No Step Skipping)

To prevent “skipping steps”, this skill uses a **mandatory execution protocol** (state-machine style). Even if the user says “just write the service” or “only tweak proto”, you MUST start from **Step 0: Preconditions & Audit**, and you may only advance after each step’s validation passes.

### Execution Protocol (MUST follow)

1. **Advance only in 0→7 order**: You MUST NOT start directly from Step 3/6/7. After Step 0 audit, proceed in order; a step may be marked `N/A` only if its audit/validation items are still satisfied.
2. **Provide “completion evidence” per step**: Include (a) commands run, (b) key file existence checks, (c) if failed: the blocking reason + the next action to unblock.
3. **Hard stop on failed validation**: If any validation checklist item fails, stop and fix it before continuing (no “write now, fix later”).
4. **Edit only editable files**: You may edit only `*.proto`, `internal/service/*.go`, and `internal/data/{table}.go`. Never edit generated artifacts (see “Do Not Edit” below).
5. **Proto changes trigger regeneration**: Any `.proto` change requires rerunning Steps 4/5/7 (and re-validating), otherwise the work is incomplete.
6. **Lock proto before writing logic**: Before Step 6 begins, confirm proto is “locked”; otherwise `make pbtocode` may overwrite/change service file layout and cause rework.

### Do Not Edit (Will Be Overwritten)

- `internal/data/gorm/**` `*.gen.go`
- `api/**/v1/*.pb.go`, `*_grpc.pb.go`, `*_http.pb.go`, `*.pb.validate.go`
- `doc/swagger/**` `*.swagger.json`

## Project Layout

```
ai-boilerplate-backend/
├── api/
│   ├── admin/v1/              # Admin Protobuf API definitions
│   └── app/v1/                # App Protobuf API definitions
├── internal/
│   ├── service/               # Service layer (business logic)
│   └── data/
│       ├── {table}.go         # Data layer extensions (editable)
│       └── gorm/
│           ├── ai_boilerplate_model/   # GORM models (generated, do not edit)
│           ├── ai_boilerplate_dao/     # DAO (generated, do not edit)
│           └── ai_boilerplate_repo/    # Repos (generated, do not edit)
└── doc/sql/ai_boilerplate/     # SQL table definitions
```

## Core Workflow Principles

**Critical rule: generate first, then edit.**

1. **Follow the sequence**: no cross-step execution (allow `N/A` only with audit/validation satisfied).
2. **Validate each step**: confirm outputs exist before moving on.
3. **Edit only where allowed**: Step 3 (proto) and Step 6 (service/data).
4. **Regenerate after proto edits**: after modifying proto, rerun `make api` and `make pbtocode`.

---

## Step 0–7 Workflow (Strict Order)

### Step 0: Preconditions & Audit (MANDATORY)

#### 0.1 Minimal required inputs (all required)

- `position`: `admin` or `app`
- `table_name`: single table or comma-separated tables (ASCII comma, no spaces)
- Goal: standard CRUD / CRUD + custom filters / additional business RPCs (list RPC names)

**Multi-table rule:** if `table_name` contains multiple tables, Steps 2–6 MUST be executed **per table sequentially** (produce completion evidence per table). Steps 1 and 7 may be run once globally.

#### 0.2 Repo path confirmation

- Backend directory: `ai-boilerplate-backend/`
- All commands assume: `cd ai-boilerplate-backend && ...`

#### 0.3 State audit (decide where you may start / mark `N/A`)

If the user insists on “starting from the middle”, you MUST verify artifacts exist; if missing, you MUST go back and fill earlier steps:

- Claim Step 1 done: check `internal/data/gorm/ai_boilerplate_model/`, `ai_boilerplate_dao/`, `ai_boilerplate_repo/` for the target table artifacts
- Claim Step 2 done: check `api/{position}/v1/{table}.proto` exists
- Claim Step 4 done: check `api/{position}/v1/{table}.pb.go`, `*_http.pb.go`, etc. exist
- Claim Step 5 done: check `internal/service/{position}_v1_{table}_*.go` exists

You MUST output the audit conclusion:
- `Start from step X` (why: which artifacts exist/missing)
- `Missing user inputs` (e.g., table schema, API requirements)

---

### Step 1: Verify DB table and generate GORM code

#### 1.1 Check table existence

Prefer MCP dbhub to check whether the table exists. If DB access is not available, you may skip DB checking, but you MUST confirm `doc/sql/ai_boilerplate/{table}.sql` exists or the user provides the CREATE TABLE SQL.

```
Tool: mcp__dbhub__search_objects
Params:
  object_type: "table"
  pattern: "{table}"
  schema: "public"
  detail_level: "summary"
```

#### 1.2 Handle outcomes

| Outcome | Action |
|---|---|
| ✅ Table exists and schema is correct | Proceed to Step 1.3 |
| ⚠️ Table missing | Check `doc/sql/ai_boilerplate/{table}.sql`<br>- Exists: run via `mcp__dbhub__execute_sql`<br>- Missing: ask user to create/provide SQL |
| ⚠️ Table schema needs changes | Ask user to update schema first |

#### 1.3 Generate GORM artifacts

```bash
cd ai-boilerplate-backend && make gorm
```

#### 1.4 Validate outputs

Confirm all files exist:
- `internal/data/gorm/ai_boilerplate_model/{table}.gen.go`
- `internal/data/gorm/ai_boilerplate_dao/{table}.gen.go`
- `internal/data/gorm/ai_boilerplate_repo/{table}.gen.go`

**Done criteria:** all required files exist. Otherwise, stop here and diagnose (common: table missing / DB connection / generator failure).

---

### Step 2: Generate Protobuf definition

#### 2.1 Run generation

Choose by API position:

```bash
# Admin APIs
cd ai-boilerplate-backend && make sqltopb admin {table_name}

# App APIs
cd ai-boilerplate-backend && make sqltopb app {table_name}

# Multiple tables (comma-separated, no spaces)
cd ai-boilerplate-backend && make sqltopb admin table1,table2,table3
```

#### 2.2 Validate outputs

Confirm file exists:
- `api/{position}/v1/{table}.proto`

**Done criteria:** proto exists and is parsable. Otherwise, stop here (do not proceed to Steps 3/4).

---

### Step 3: Edit Protobuf API (AI-editable)

**Rule:** You may ONLY edit the generated proto file; do not create new proto files.

#### 3.1 Read and evaluate the generated proto

Read: `api/{position}/v1/{table}.proto`, decide whether changes are needed.

#### 3.2 Decide whether to modify

| Case | Action |
|---|---|
| Standard CRUD is sufficient | Mark Step 3 as `N/A` (no changes), explicitly state “proto locked”, then proceed to Step 4 |
| Need extra list filters | Modify `Get{Table}ListReq` |
| Need different validation | Adjust `buf.validate` rules |
| Need new business method(s) | Add new rpc(s) in the proto `service` |
| Need to remove unused methods | Remove corresponding rpc(s) |

#### 3.3 Common edits

- Add query filters: update `Get{Table}ListReq`
- Validation: adjust `buf.validate`
- Add business RPCs: add rpc(s) in proto `service`
- Remove RPCs: delete the rpc definitions

**Done criteria:** proto is saved and syntactically correct. Next step MUST be Step 4 (do not start writing service code yet).

---

### Step 4: Generate API code

```bash
cd ai-boilerplate-backend && make api
```

Validate generated outputs (do not edit):
- `api/{position}/v1/{table}.pb.go`
- `api/{position}/v1/{table}.pb.validate.go`
- `api/{position}/v1/{table}_http.pb.go`
- `api/{position}/v1/{table}_grpc.pb.go`
- `doc/swagger/{position}/v1/{table}.swagger.json`

**Done criteria:** artifacts exist and generation had no errors. If failed, go back to Step 3 and fix proto.

---

### Step 5: Generate Service stubs

```bash
cd ai-boilerplate-backend && make pbtocode
```

Validate generated outputs:
- `internal/service/{position}_v1_{table}.go` (service struct)
- `internal/service/{position}_v1_{table}_{method}.go` (one file per rpc)

**Done criteria:** target rpc stub files exist. If missing, stop here and diagnose (common: Step 4 failed / proto missing rpc / generator error).

---

### Step 6: Implement business logic (AI-editable)

Goal: implement logic for each `internal/service/{position}_v1_{table}_*.go` method file.

#### 6.0 Pre-check (MUST pass)

- Step 3 is explicitly “proto locked” (or user explicitly accepts that later proto changes will cause rework)
- If you need to add/remove rpc: go back to Step 3, then rerun Steps 4/5/7

#### 6.1 Locate files to implement

Find: `internal/service/{position}_v1_{table}_*.go`

#### 6.2 Implement logic

Implementation approach: **clone patterns from existing services**.

References to start from:
- Single record: `internal/service/app_v1_user_getuserinfo.go`
- List + filters: `internal/service/admin_v1_user_getuserlist.go`
- Create/update: use create/update implementations in other tables

Workflow:
1. Use Glob to find similar service files
2. Use Read to open reference implementations
3. Use Edit to implement your service methods by pattern
4. Keep style consistent (error wrapping, mapping, naming, time formatting)

**Independent-file examples (based on existing code)**

- **ONLY when you are in Step 6 and you need a copyable “per-file skeleton”**: use Read to open `references/service-examples.md`; otherwise do not read it (avoid loading irrelevant context).

#### 6.3 Common implementation scenarios

**A. Single record**
- Reference: `internal/service/app_v1_user_getuserinfo.go`
- Key points: `FindOneCacheByID`, user ID from context metadata, time formatting

**B. List query (pagination + filters)**
- Reference: `internal/service/admin_v1_user_getuserlist.go`
- Key points: build `condition.Req`, use `FindMultiCacheByCondition`

**C. Join/related lookup**
- Reference: `internal/service/admin_v1_user_getuserlist.go` (search for `FindMultiCacheByUserIDS` / `make(map`)
- Key points: query main list, collect IDs, batch query related rows, build a map, fill into response

**D. Simple list (no pagination)**
- Reference: `internal/service/app_v1_helpcategory_listhelpcategories.go`
- Key points: call custom repo method, straightforward mapping

#### 6.4 Add data-layer methods (if needed)

If generated repo methods don’t match the requirement, add custom methods in `internal/data/{table}.go`.

References:
- `internal/data/helpcategory.go` (custom query example)
- `internal/data/user.go` (complex business example)

Workflow:
1. Read reference data files
2. Add a method in `internal/data/{table}.go`
3. Use `ai_boilerplate_dao.Use(r.data.gorm).{Table}` for DB operations
4. Keep style consistent (naming, errors)

**Done criteria:** all target rpc methods are implemented and the code compiles. Only then proceed to Step 7.

---

### Step 7: Dependency injection and code quality

#### 7.1 Wire

```bash
cd ai-boilerplate-backend && make wire
```

Validate: `internal/data/wire_gen.go` and `internal/service/wire_gen.go` updated.

#### 7.2 Import formatting

```bash
cd ai-boilerplate-backend && make gci
```

#### 7.3 Lint

```bash
cd ai-boilerplate-backend && make lint
```

If lint fails: read errors → fix → rerun.

#### 7.4 Final build

```bash
cd ai-boilerplate-backend && go build ./cmd/...
```

**Done criteria:** `wire` / `gci` / `lint` / `go build` all pass. If any fails, fix first before delivering.

---

## Per-Step Output Template (Anti-Skipping)

For every step, output using this template (do not omit “Validation” or “Blockers”):

1. `Step`: Step N (name)
2. `Action`: commands run / files edited
3. `Validation`: checklist with ✅/❌ results
4. `Blockers`: if any, the failure reason + what input/command unblocks it
5. `Next`: only Step N+1 (or rollback to a prior step to fix); never skip forward

---

## Skill Triggers

Use this skill when the user asks for:

- Generate backend code for a table
- Implement backend API(s) / business logic
- Add/update CRUD endpoints
- Update backend code after table schema changes
- Full backend development flow (DB → proto → API codegen → service)
- Or explicitly invokes `/backend-codeing`

