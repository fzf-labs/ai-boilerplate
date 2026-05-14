---
name: admin-codeing
description: Admin 管理后台前端开发技能。基于 Vue 3 + TypeScript + Ant Design Vue 技术栈，提供完整的 CRUD 页面开发工作流程。触发场景：(1) 开发管理后台页面 (2) 实现表单和表格功能 (3) 对接后端 API 接口 (4) 配置菜单和权限 (5) 实现复杂业务页面 (6) 组件封装和状态管理。关键词：admin、管理后台、前端开发、Vue、表单、表格、CRUD、页面开发。
---

# Admin 管理后台开发技能

Admin 管理后台前端开发工作流程（Vue 3 + TypeScript + Ant Design Vue）。

## 快速决策指南

**你要开发什么类型的页面？**

1. **标准 CRUD 页面**（列表+表单）→ 使用标准流程（参考 `examples/01-basic-crud.md`）
2. **带搜索/批量操作的列表** → 参考 `examples/02-advanced-features.md`
3. **详情展示页面** → 参考 `examples/03-detail-page.md`
4. **复杂业务页面** → 需要自定义组件和状态管理

**是否需要生成 API？**
- 后端 API 已更新？ → 是，执行 Step 2
- API 已存在？ → 跳到 Step 3

**需要组件用法？** → 查看 `references/components-guide.md`

**需要代码规范？** → 查看 `references/best-practices.md`

---

## 开发工作流程

### Step 1: 确认需求

**询问用户**：
- 模块名称和功能描述
- 页面类型：标准 CRUD / 复杂业务页面
- 后端接口路径（如 `/admin/v1/sys_dept`）
- 特殊字段类型（树形、富文本、上传等）
- 权限代码（如 `system:dept:create`）

**复杂页面额外询问**：
- 页面功能和交互流程
- 是否需要状态管理（Pinia）
- 特殊组件需求（图表、编辑器、拖拽等）

**自动执行**：
- 查看类似页面作为参考
- 检查后端接口文档

### Step 2: 创建目录结构

**标准 CRUD 页面**：
```
views/{module}/{name}/
├── index.vue       # 列表页
├── data.ts         # 表格列定义、表单配置
└── modules/
    └── form.vue    # 创建/编辑表单
```

**复杂业务页面**：
```
views/{module}/{name}/
├── index.vue       # 主页面（容器）
├── components/     # 业务组件
│   ├── Header.vue
│   ├── Sidebar.vue
│   └── Content.vue
└── composables/    # 组合式函数
    └── use{Name}.ts
```

### Step 3: 生成 API 接口代码

**前置条件**：后端 API 已实现并生成 Swagger 文件

**生成命令**：
```bash
cd ai-boilerplate-admin && pnpm api:gen
```

**生成文件**：
- `apps/web-antd/src/api/v1/{module}/index.ts` - API 函数
- `apps/web-antd/src/api/v1/{module}/types.ts` - TypeScript 类型

**重要提示**：
- ✅ 使用 `pnpm api:gen` 自动生成，禁止手动创建或修改
- ✅ 生成的文件会被下次生成覆盖，不要手动编辑
- ✅ 如果 API 缺失或类型错误，需要后端更新 Swagger 后重新生成

**参考**：详细说明请查看 `admin-api-gen` 技能

### Step 4: 实现页面功能

**标准 CRUD 页面**：

1. **列表页面** (`index.vue`) - 表格配置、操作按钮、权限控制
2. **表格列和表单配置** (`data.ts`) - `useGridColumns()`、`useFormSchema()`
3. **表单组件** (`modules/form.vue`) - 表单配置、弹窗配置、创建/编辑逻辑

**参考完整示例**：[examples/01-basic-crud.md](./examples/01-basic-crud.md)

**复杂业务页面**：

1. **主页面** (`index.vue`) - 页面布局、子组件引用、状态管理集成
2. **业务组件** (`components/`) - 按功能拆分的子组件、组件间通信
3. **状态管理** (可选) - 在 `packages/stores/src/modules/{name}.ts` 创建 Pinia Store
4. **组合式函数** (`composables/`) - 可复用的业务逻辑

**何时使用 Pinia**：跨页面共享数据、需要持久化的数据、复杂业务逻辑

**组件使用指南**：需要了解表单、表格、弹窗等组件的详细用法时，查看 [references/components-guide.md](./references/components-guide.md)

**代码规范**：需要了解 TypeScript、API、样式等代码规范时，查看 [references/best-practices.md](./references/best-practices.md)

### Step 5: 配置菜单

本项目使用后端动态路由模式，菜单配置存储在数据库的 `sys_menu` 表中。

**使用 dbhub MCP 工具插入菜单数据**：

```sql
-- 1. 插入父菜单（如果需要）
INSERT INTO sys_menu (pid, name, type, path, permission, icon, component, component_name, sort, status, created_at, updated_at)
VALUES (
  '父菜单ID',          -- pid: 父菜单ID，顶级菜单使用 '0'
  '模块名称',          -- name: 菜单显示名称
  'menu',             -- type: 'menu' 表示菜单项
  'module-path',      -- path: 路由路径
  '',                 -- permission: 菜单权限代码（可为空）
  'icon-name',        -- icon: 图标名称（如 'fa:user'）
  'module/index',     -- component: 组件路径（相对于 views 目录）
  'ModuleName',       -- component_name: 组件名称（PascalCase）
  100,                -- sort: 排序值（数字越小越靠前）
  1,                  -- status: 状态（1=启用，0=禁用）
  NOW(),
  NOW()
);

-- 2. 插入按钮权限（CRUD 操作）
INSERT INTO sys_menu (pid, name, type, path, permission, icon, component, component_name, sort, status, created_at, updated_at)
VALUES
  ('菜单ID', '查询', 'button', '', 'module:query', '', '', NULL, 0, 1, NOW(), NOW()),
  ('菜单ID', '创建', 'button', '', 'module:create', '', '', NULL, 1, 1, NOW(), NOW()),
  ('菜单ID', '更新', 'button', '', 'module:update', '', '', NULL, 2, 1, NOW(), NOW()),
  ('菜单ID', '删除', 'button', '', 'module:delete', '', '', NULL, 3, 1, NOW(), NOW());
```

**字段说明**：
- `pid`: 父菜单ID，顶级菜单使用 `'0'`
- `type`: `'menu'` 表示菜单项，`'button'` 表示按钮权限
- `path`: 路由路径（菜单项必填，按钮为空）
- `permission`: 权限代码（格式：`module:action`）
- `component`: 组件路径（相对于 `apps/web-antd/src/views/` 目录）

**注意事项**：
- 使用 dbhub MCP 工具的 `execute_sql` 功能执行 SQL 语句
- 菜单插入后，前端会自动从后端 API 获取最新的菜单配置
- 权限代码需要与后端 API 的权限验证保持一致

### Step 6: 质量检查与测试

**代码质量检查**：
```bash
cd ai-boilerplate-admin && pnpm check:type
cd ai-boilerplate-admin && pnpm lint:fix
```

**功能测试**：
1. 运行 `cd ai-boilerplate-admin && pnpm dev`
2. 检查页面渲染
3. 测试功能（CRUD / 业务流程）
4. 验证权限控制
5. 检查控制台错误

**向用户展示**：
- 创建的文件列表
- 页面访问路径
- 权限代码列表

---

## 快速参考

### 常用命令

```bash
# 启动开发服务器
cd ai-boilerplate-admin && pnpm dev

# 生成 API 代码
cd ai-boilerplate-admin && pnpm api:gen

# 代码检查并自动修复
cd ai-boilerplate-admin && pnpm lint:fix

# 类型检查
cd ai-boilerplate-admin && pnpm check:type

# 构建项目
cd ai-boilerplate-admin && pnpm build
```

### 验证检查清单

- [ ] API 接口定义已创建
- [ ] 列表页面已创建
- [ ] 表单组件已创建
- [ ] 菜单配置已添加
- [ ] 权限代码已配置
- [ ] 页面正常渲染
- [ ] CRUD 功能正常
- [ ] 无 TypeScript 错误
- [ ] 无 ESLint 警告

---

## 进阶学习资源

### 完整示例

- **[examples/01-basic-crud.md](./examples/01-basic-crud.md)** - 基础 CRUD 页面开发
- **[examples/02-advanced-features.md](./examples/02-advanced-features.md)** - 进阶功能实现（搜索、批量操作、导入导出）
- **[examples/03-detail-page.md](./examples/03-detail-page.md)** - 详情页实现
- **[examples/README.md](./examples/README.md)** - 示例索引和学习路径

### 详细参考文档

- **[references/components-guide.md](./references/components-guide.md)** - 组件使用指南（表单、表格、弹窗、权限、国际化、文件上传等）
- **[references/best-practices.md](./references/best-practices.md)** - 开发最佳实践（代码组织、TypeScript、性能优化、状态管理、样式规范等）

### 常见场景实现

需要实现以下场景时，查看对应文档：

1. **树形表格** - 设置 `treeConfig`（参考 `references/components-guide.md`）
2. **搜索表单** - 使用 inline 布局的 Form（参考 `examples/02-advanced-features.md`）
3. **批量操作** - 使用 `checkboxConfig`（参考 `examples/02-advanced-features.md`）
4. **详情页** - 两种风格：简单版（Descriptions）和华丽版（Card）（参考 `examples/03-detail-page.md`）
5. **文件上传** - 使用 `uploadFile` 工具函数（参考 `references/components-guide.md`）

---

## 与其他技能协作

- **admin-api-gen**: 生成 API 客户端代码
- **backend-dev**: 先开发后端接口，再用 admin-codeing 开发前端页面
- **interview**: 复杂功能先用 interview 探索方案
- **tech-decision**: 需要技术选型时使用
