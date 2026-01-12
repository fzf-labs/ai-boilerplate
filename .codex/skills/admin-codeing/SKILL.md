---
name: admin-codeing
description: Admin 管理后台开发技能。当用户需要开发前端页面、对接后端接口、实现 CRUD 功能时使用此技能。触发场景包括：(1) 新增管理页面 (2) 表单和表格开发 (3) API 接口对接 (4) 权限控制实现 (5) 组件封装 (6) 完整的前端功能开发流程
---

# Admin 管理后台开发技能

本技能提供 Admin 管理后台的前端开发完整工作流程。

## 项目技术栈

- **框架**: Vue 3 + TypeScript (Composition API)
- **UI 库**: Ant Design Vue
- **表格**: Vxe-Table
- **状态**: Pinia
- **构建**: Vite

## 项目结构

```
ai-boilerplate-admin/
├── apps/web-antd/
│   ├── src/
│   │   ├── views/      # 页面组件
│   │   ├── api/        # API 接口
│   │   ├── router/     # 路由配置
│   │   └── locales/    # 国际化
│   └── package.json
└── packages/           # 共享包
```

---

## 核心工作流程

### 流程 A：标准 CRUD 页面开发

**适用场景**：为后端已有的 CRUD 接口创建前端页面

#### 步骤 1：确认需求

**询问用户**：
- 模块名称和功能描述
- 后端接口路径（如：`/admin/v1/sys_dept`）
- 特殊字段类型（树形、富文本、上传等）
- 权限代码（如：`system:dept:create`）

**自动执行**：
- 查看类似页面作为参考
- 检查后端接口文档

#### 步骤 2：创建目录结构

在 `apps/web-antd/src/views/{module}/{name}/` 创建：

```
{module}/{name}/
├── index.vue              # 列表页
├── data.ts               # 表格列定义、表单配置
└── modules/
    └── form.vue          # 创建/编辑表单
```

#### 步骤 3：创建 API 接口定义

在 `apps/web-antd/src/api/{module}/{name}/index.ts` 创建 API 接口。

**参考**：`examples/01-basic-crud.md` 中的 "API 接口定义" 部分

**关键点**：
- 使用 namespace 组织类型
- 标准函数命名：`get{Name}List`, `get{Name}Info`, `create{Name}`, `update{Name}`, `delete{Name}`

#### 步骤 4：创建列表页面

创建 `index.vue`，包含：
- 表格配置（使用 `useVbenVxeGrid`）
- 操作按钮（创建、编辑、删除）
- 权限控制（`v-access:code`）
- 刷新逻辑

**参考**：`examples/01-basic-crud.md` 中的 "列表页面" 部分

#### 步骤 5：创建表格列和表单配置

创建 `data.ts`，包含：
- `useGridColumns()` - 表格列配置
- `useFormSchema()` - 表单配置

**参考**：
- `examples/01-basic-crud.md` 中的 "表格列和表单配置" 部分
- `references/components-guide.md` 了解所有可用组件

#### 步骤 6：创建表单组件

创建 `modules/form.vue`，包含：
- 表单配置（使用 `useVbenForm`）
- 弹窗配置（使用 `useVbenModal`）
- 创建/编辑逻辑

**参考**：`examples/01-basic-crud.md` 中的 "表单组件" 部分

#### 步骤 7：配置路由

在 `apps/web-antd/src/router/routes/modules/{module}.ts` 添加路由配置。

**参考**：`examples/01-basic-crud.md` 中的 "路由配置" 部分

#### 步骤 8：测试验证

**自动验证**：
1. 运行 `pnpm dev`
2. 检查页面渲染
3. 测试 CRUD 功能
4. 验证权限控制
5. 检查控制台错误

**向用户展示**：
- 创建的文件列表
- 页面访问路径
- 权限代码列表

---

### 流程 B：复杂业务页面开发

**适用场景**：自定义页面、复杂交互、多步骤流程

#### 步骤 1：需求分析

**询问用户**：
- 页面功能和交互流程
- 需要调用的 API 接口
- 是否需要状态管理（Pinia）
- 特殊组件需求（图表、编辑器、拖拽等）

#### 步骤 2：组件拆分

将复杂页面拆分为多个子组件：

```
{module}/{name}/
├── index.vue              # 主页面（容器）
├── components/            # 业务组件
│   ├── Header.vue
│   ├── Sidebar.vue
│   └── Content.vue
└── composables/          # 组合式函数
    └── use{Name}.ts
```

#### 步骤 3：状态管理（如需要）

在 `packages/stores/src/modules/{name}.ts` 创建 Pinia Store。

**何时使用 Pinia**：
- 跨页面共享的数据
- 需要持久化的数据
- 复杂的业务逻辑

**参考**：`references/best-practices.md` 中的 "状态管理" 部分

#### 步骤 4：实现业务逻辑

使用 Composition API 组织代码。

**参考**：`references/best-practices.md` 中的 "代码组织" 部分

---

## 常用组件速查

详细的组件使用方法请参考 `references/components-guide.md`

### 表单组件 (useVbenForm)

支持的组件类型：
- Input / InputNumber / Textarea
- Select / TreeSelect
- DatePicker / RangePicker
- RadioGroup / CheckboxGroup
- Switch / Upload

### 表格组件 (useVbenVxeGrid)

支持的功能：
- 分页 / 排序 / 筛选
- 树形结构
- 多选 / 单选
- 自定义列渲染
- 操作按钮

### 弹窗组件 (useVbenModal)

用于创建/编辑表单弹窗。

### 权限控制

```vue
<!-- 指令方式 -->
<Button v-access:code="['system:user:create']">创建</Button>

<!-- 函数方式 -->
<script setup>
import { useAccess } from '@vben/access';
const { hasAccessByCode } = useAccess();
</script>
```

---

## 代码规范

详细的代码规范请参考 `references/best-practices.md`

### TypeScript 类型定义

```typescript
// ✅ 推荐：使用 namespace 组织类型
export namespace SystemUserApi {
  export interface User {
    id: string;
    username: string;
  }
}
```

### API 函数命名

- 列表：`get{Name}List`
- 详情：`get{Name}Info`
- 创建：`create{Name}`
- 更新：`update{Name}`
- 删除：`delete{Name}`
- 状态更新：`update{Name}Status`

### 组件命名

- 页面组件：PascalCase
- 文件名：kebab-case 或 camelCase

---

## 常见场景实现

以下场景的详细实现请参考对应文档：

1. **树形表格** - 设置 `treeConfig`（参考 `references/components-guide.md`）
2. **搜索表单** - 使用 inline 布局的 Form（参考 `examples/02-advanced-features.md`）
3. **批量操作** - 使用 `checkboxConfig`（参考 `examples/02-advanced-features.md`）
4. **详情页** - 两种风格：简单版（Descriptions）和华丽版（Card）（参考 `examples/03-detail-page.md`）
5. **文件上传** - 使用 `uploadFile` 工具函数（参考 `references/components-guide.md`）

---

## 性能优化

参考 `references/best-practices.md` 中的 "性能优化" 部分：

- 虚拟滚动（大列表）
- 防抖节流
- 组件懒加载
- 计算属性缓存

---

## 完整示例

查看 `examples/` 目录获取完整的开发示例：

- **[examples/01-basic-crud.md](./examples/01-basic-crud.md)** - 基础 CRUD 页面开发
  - API 接口定义
  - 表格列配置
  - 表单配置
  - 列表页面
  - 表单组件
  - 路由配置

- **[examples/02-advanced-features.md](./examples/02-advanced-features.md)** - 进阶功能实现
  - 搜索表单
  - 批量操作
  - 数据导入/导出
  - 高级筛选
  - 排序功能
  - 刷新策略

- **[examples/03-detail-page.md](./examples/03-detail-page.md)** - 详情页实现
  - 简单版详情页（Descriptions）
  - 华丽版详情页（Card + 自定义样式）
  - 列表页集成
  - Tab 切换
  - 关联数据展示

- **[examples/README.md](./examples/README.md)** - 示例索引和学习路径

---

## 工作流程图

```
需求确认 → 创建目录 → API 定义 → 列表页面 → 表单组件 → 路由配置 → 测试验证
```

---

## 快速命令

```bash
# 启动开发服务器
cd ai-boilerplate-admin && pnpm dev

# 代码检查
pnpm lint

# 类型检查
pnpm type-check

# 构建项目
pnpm build
```

---

## 验证检查清单

- [ ] API 接口定义已创建
- [ ] 列表页面已创建
- [ ] 表单组件已创建
- [ ] 路由配置已添加
- [ ] 权限代码已配置
- [ ] 页面正常渲染
- [ ] CRUD 功能正常
- [ ] 无 TypeScript 错误
- [ ] 无 ESLint 警告

---

## 与其他技能协作

- **backend-dev**: 先开发后端接口，再用 admin-codeing 开发前端页面
- **interview**: 复杂功能先用 interview 探索方案
- **tech-decision**: 需要技术选型时使用
- **bug-detective**: 遇到问题时使用
