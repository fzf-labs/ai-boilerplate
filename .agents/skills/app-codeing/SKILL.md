---
name: app-codeing
description: App 移动端前端开发技能,基于 uni-app + Vue 3 + TypeScript + wot-design-uni 技术栈。触发场景包括:(1) 开发移动端页面 (2) 实现列表/表单/详情页 (3) 对接后端 API (4) 使用 wot-design-uni 组件 (5) 状态管理和路由导航 (6) 完整的移动端 CRUD 功能开发。关键词:app、移动端、uni-app、Vue、页面开发、表单、列表、详情页、wot-design-uni、z-paging。
---

# App Development Skill

uni-app 移动端应用开发技能 (ai-boilerplate-uniapp),使用 Vue 3 + TypeScript + wot-design-uni。

## 快速决策指南

**你要开发什么类型的页面?**

1. **列表页(带分页)** → 使用 z-paging 模式 → 查看 `examples/01-list-page.md`
2. **表单页(带验证)** → 使用 wd-form 模式 → 查看 `examples/02-form-page.md`
3. **详情/展示页** → 使用基础页面模式 → 查看 `examples/03-detail-page.md`
4. **需要组件参考?** → 查看 `references/components-guide.md`
5. **需要最佳实践?** → 查看 `references/best-practices.md`

**是否需要生成 API?**
- 后端新增了 API? → 是,执行 Step 1
- API 已存在? → 否,跳到 Step 2

---

## 开发工作流

### Step 1: 生成 API 客户端代码 (条件执行)

**何时执行**: 后端 API 已更新或新增接口。

**何时跳过**: 功能所需的 API 类型已存在。

```bash
cd ai-boilerplate-uniapp && pnpm api:gen
```

**重要**: 禁止手动编辑 `src/api/v1/*` - 这些文件由 Swagger 自动生成。

---

### Step 2: 选择页面类型和位置

**选择目录**:
- Tabbar 页面 → `src/pages/{module}/`
- 子页面 → `src/pages-fg/{module}/`

**选择模式** (完整代码见 examples/):
- 列表+分页 → `examples/01-list-page.md`
- 表单+验证 → `examples/02-form-page.md`
- 详情/展示 → `examples/03-detail-page.md`

---

### Step 3: 实现页面功能

**基础页面结构**:

```vue
<script lang="ts" setup>
import type { DataType } from '@/api/v1/module/types'
import { useToast } from 'wot-design-uni'
import { apiMethod } from '@/api/v1/module/module'

definePage({
  style: { navigationBarTitleText: 'Page Title' },
})

const toast = useToast()
const loading = ref(false)
const data = ref<DataType | null>(null)

async function fetchData() {
  try {
    loading.value = true
    const res = await apiMethod({ options: {} })
    data.value = res.info || null
  }
  catch (error) {
    console.error('Error:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

onLoad(() => fetchData())
</script>

<template>
  <view class="page-container">
    <!-- 页面内容 -->
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.page-container {
  min-height: 100vh;
  background: var(--fg-bg);
}
</style>
```

**详细指南**:
- 组件用法 → `references/components-guide.md`
- 编码规范 → `references/best-practices.md`

---

### Step 4: 质量检查和测试

**运行检查**:

```bash
cd ai-boilerplate-uniapp && pnpm type-check && pnpm lint:fix
```

**本地测试**:

```bash
cd ai-boilerplate-uniapp && pnpm dev
```

修复所有 TypeScript 和 ESLint 错误后再继续。

---

## 技术栈

| 类别 | 技术 |
|------|------|
| 框架 | Vue 3 + TypeScript (Composition API) |
| UI 库 | wot-design-uni |
| 列表组件 | z-paging |
| 状态管理 | Pinia |
| HTTP 客户端 | Alova |
| 构建工具 | Vite + uni-app |

---

## 项目结构

```
ai-boilerplate-uniapp/src/
├── pages/           # Tabbar 页面
├── pages-fg/        # 子页面(非 tabbar)
├── api/v1/          # 生成的 API (禁止编辑)
├── store/           # Pinia stores
├── components/      # 共享组件
└── http/            # HTTP 客户端配置
```

---

## 参考文件

**何时使用**:

- **需要组件语法?** → `references/components-guide.md`
  - wot-design-uni 组件 (wd-form, wd-button, wd-card 等)
  - z-paging 列表组件
  - 组件属性和事件

- **需要代码模式?** → `references/best-practices.md`
  - TypeScript 规范
  - API 调用模式
  - 页面生命周期钩子
  - CSS 变量
  - 认证模式

- **需要完整示例?** → `examples/`
  - `01-list-page.md` - 带分页的列表
  - `02-form-page.md` - 带验证的表单
  - `03-detail-page.md` - 带参数的详情页

---

## 常用命令

```bash
pnpm dev          # 启动开发服务器
pnpm api:gen      # 从 Swagger 生成 API
pnpm type-check   # TypeScript 检查
pnpm lint:fix     # Lint 并自动修复
pnpm build        # 构建 H5
pnpm build:mp     # 构建微信小程序
```

---

## 验证清单

完成前检查:

- [ ] API 已生成 (如果后端有变更)
- [ ] 页面在正确目录 (pages/ 或 pages-fg/)
- [ ] 使用 `definePage` 设置 navigationBarTitleText
- [ ] 处理了加载/错误状态
- [ ] 模板中包含 `<wd-toast />`
- [ ] 使用 CSS 变量 (var(--fg-*))
- [ ] Type check 通过
- [ ] Lint check 通过
- [ ] 在开发服务器中测试通过

---

## 故障排查

**API 类型未找到**:
```bash
cd ai-boilerplate-uniapp && pnpm api:gen
```

**TypeScript 错误**:
```bash
cd ai-boilerplate-uniapp && pnpm type-check
```

**Lint 错误**:
```bash
cd ai-boilerplate-uniapp && pnpm lint:fix
```

**组件不工作**: 查看 `references/components-guide.md` 了解正确用法。
