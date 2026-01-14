---
name: app-codeing
description: App development skill for uni-app mobile application. Use when developing mobile pages, integrating backend APIs, implementing features with wot-design-uni components. Triggers include：(1) Creating new pages (2) Form and list development (3) API integration (4) State management (5) Component usage (6) Complete mobile app development workflow
---

# App Development Skill

Complete workflow for developing the uni-app mobile application (ai-boilerplate-app).

## Tech Stack

- **Framework**: Vue 3 + TypeScript (Composition API)
- **UI Library**: wot-design-uni (preferred)
- **List Component**: z-paging
- **State Management**: Pinia
- **HTTP Client**: Alova
- **Build Tool**: Vite + uni-app

## Project Structure

```
ai-boilerplate-app/src/
├── pages/           # Tabbar pages
├── pages-fg/        # Sub-pages (non-tabbar)
├── api/v1/          # Generated API (DO NOT EDIT)
├── store/           # Pinia stores
├── components/      # Shared components
└── http/            # HTTP client config
```

---

## Core Workflow

### Step 1: Generate API (if needed)

```bash
cd ai-boilerplate-app && pnpm api:gen
```

**Important**: NEVER manually edit files in `src/api/` - they are auto-generated.

### Step 2: Create Page File

- **Tabbar pages**: `src/pages/{module}/`
- **Sub-pages**: `src/pages-fg/{module}/`

Basic structure:

```vue
<script lang="ts" setup>
import type { SomeType } from '@/api/v1/module/types'
import { useToast } from 'wot-design-uni'
import { someApi } from '@/api/v1/module/module'

definePage({
  style: { navigationBarTitleText: 'Page Title' },
})

const toast = useToast()
const loading = ref(false)
const data = ref<SomeType | null>(null)

async function fetchData() {
  try {
    loading.value = true
    const res = await someApi({ options: {} })
    data.value = res.info || null
  }
  catch (error) {
    console.error('Failed:', error)
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

### Step 3: Test

```bash
cd ai-boilerplate-app && pnpm dev
```

---

## Page Type Examples

| Type | Reference | Key Component |
|------|-----------|---------------|
| List with pagination | `examples/01-list-page.md` | z-paging |
| Form with validation | `examples/02-form-page.md` | wd-form |
| Detail page | `examples/03-detail-page.md` | onLoad params |

---

## References

- **Components**: `references/components-guide.md` - wot-design-uni usage
- **Best Practices**: `references/best-practices.md` - Code standards, CSS variables, patterns

---

## Quick Commands

```bash
pnpm dev          # Start dev server
pnpm api:gen      # Generate API from Swagger
pnpm build        # Build for H5
pnpm build:mp     # Build for WeChat Mini Program
pnpm type-check   # Type check
pnpm lint         # Lint
```

---

## Verification Checklist

- [ ] API types generated (if needed)
- [ ] Page in correct directory (pages/ or pages-fg/)
- [ ] `definePage` configured
- [ ] Loading/error states handled
- [ ] `<wd-toast />` in template
- [ ] CSS variables used
- [ ] No TypeScript/ESLint errors
