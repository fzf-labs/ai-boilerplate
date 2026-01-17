---
name: app-codeing
description: App development skill for uni-app mobile application. Use when developing mobile pages, integrating backend APIs, implementing features with wot-design-uni components. Triggers include：(1) Creating new pages (2) Form and list development (3) API integration (4) State management (5) Component usage (6) Complete mobile app development workflow
---

# App Development Skill

Workflow for developing uni-app mobile application (ai-boilerplate-app) with Vue 3 + TypeScript.

## Quick Start Decision Guide

**What are you building?**

1. **List page with pagination** → Use z-paging pattern (see `examples/01-list-page.md`)
2. **Form page with validation** → Use wd-form pattern (see `examples/02-form-page.md`)
3. **Detail/display page** → Use basic page pattern (see `examples/03-detail-page.md`)
4. **Need component reference?** → See `references/components-guide.md`

**Do you need to generate API?**
- New backend API added? → Yes, run Step 1
- API already exists? → Skip to Step 2

---

## Development Workflow

### Step 1: Generate API (Conditional)

**When to run**: Backend API has been updated or new endpoints added.

**When to skip**: API types already exist for your feature.

```bash
cd ai-boilerplate-app && pnpm api:gen
```

**Important**: NEVER manually edit `src/api/v1/*` - files are auto-generated from Swagger.

---

### Step 2: Select Page Type & Location

**Choose directory**:
- Tabbar pages → `src/pages/{module}/`
- Sub-pages → `src/pages-fg/{module}/`

**Choose pattern** (see examples/ for full code):
- List with pagination → `examples/01-list-page.md`
- Form with validation → `examples/02-form-page.md`
- Detail/display → `examples/03-detail-page.md`

---

### Step 3: Implement Page

**Basic page structure**:

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
    toast.error('Load failed')
  }
  finally {
    loading.value = false
  }
}

onLoad(() => fetchData())
</script>

<template>
  <view class="page-container">
    <!-- Content -->
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

**For component usage**: See `references/components-guide.md`

**For patterns & standards**: See `references/best-practices.md`

---

### Step 4: Quality Check & Test

**Run checks**:

```bash
cd ai-boilerplate-app && pnpm type-check && pnpm lint:fix
```

**Test locally**:

```bash
cd ai-boilerplate-app && pnpm dev
```

Fix any TypeScript or ESLint errors before proceeding.

---

## Tech Stack Reference

| Category | Technology |
|----------|-----------|
| Framework | Vue 3 + TypeScript (Composition API) |
| UI Library | wot-design-uni |
| List Component | z-paging |
| State Management | Pinia |
| HTTP Client | Alova |
| Build Tool | Vite + uni-app |

---

## Project Structure

```
ai-boilerplate-app/src/
├── pages/           # Tabbar pages only
├── pages-fg/        # Sub-pages (non-tabbar)
├── api/v1/          # Generated API (DO NOT EDIT)
├── store/           # Pinia stores
├── components/      # Shared components
└── http/            # HTTP client config
```

---

## Reference Files

**When to use**:

- **Need component syntax?** → `references/components-guide.md`
  - wot-design-uni components (wd-form, wd-button, wd-card, etc.)
  - z-paging list component
  - Component props and events

- **Need code patterns?** → `references/best-practices.md`
  - TypeScript standards
  - API call patterns
  - Page lifecycle hooks
  - CSS variables
  - Authentication patterns

- **Need complete examples?** → `examples/`
  - `01-list-page.md` - Paginated list with z-paging
  - `02-form-page.md` - Form with validation
  - `03-detail-page.md` - Detail page with params

---

## Common Commands

```bash
pnpm dev          # Start dev server
pnpm api:gen      # Generate API from Swagger
pnpm type-check   # TypeScript check
pnpm lint:fix     # Lint and auto-fix
pnpm build        # Build for H5
pnpm build:mp     # Build for WeChat Mini Program
```

---

## Verification Checklist

Before completing:

- [ ] API generated (if backend changed)
- [ ] Page in correct directory (pages/ or pages-fg/)
- [ ] `definePage` with navigationBarTitleText
- [ ] Loading/error states handled
- [ ] `<wd-toast />` in template
- [ ] CSS variables used (var(--fg-*))
- [ ] Type check passes
- [ ] Lint check passes
- [ ] Tested in dev server

---

## Troubleshooting

**API types not found**:
```bash
cd ai-boilerplate-app && pnpm api:gen
```

**TypeScript errors**:
```bash
cd ai-boilerplate-app && pnpm type-check
```

**Lint errors**:
```bash
cd ai-boilerplate-app && pnpm lint:fix
```

**Component not working**: Check `references/components-guide.md` for correct usage.
