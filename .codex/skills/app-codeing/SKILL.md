---
name: app-codeing
description: "UniApp mobile application development workflow for app project. Use when users need to develop mobile app features, pages, or components. Triggers on: (1) Creating new pages (list/detail/form/tab pages) (2) Implementing business logic (3) API integration (4) Multi-platform adaptation (H5/WeChat/App) (5) Performance optimization (6) Form development (7) List with pagination (8) Complex UI interactions (9) State management with Pinia (10) Any mobile app development tasks using Vue 3 + TypeScript + UniApp stack"
---

# App Mobile Development Workflow

Complete end-to-end workflow for app mobile development using uni-app + Vue 3 + TypeScript. Supports H5, mini-programs, and native App platforms.

## Core Workflows

### Workflow A: Standard Page Development

**Use for**: List pages, detail pages, form pages, tab pages

**Steps**:
1. Confirm requirements (page type, API endpoints, target platforms, UI components)
2. In app project, run `pnpm api:gen` to generate API client code (depends on backend swagger; do not hand-write/modify generated API code)
3. Create page in `src/pages/{module}/{name}/index.vue` using Composition API and wot-design-uni components
4. Configure page routing in `pages.config.ts`
5. Add Pinia store if needed with persistence
6. Test on H5 (`pnpm dev`) and mini-program (`pnpm dev:mp`)

### Workflow B: List Page with Pagination

**Use for**: Paginated list pages

**Steps**:
1. Define list API with pagination params (page, page_size)
2. Implement page using z-paging component
3. Add pull-to-refresh and load-more
4. Handle empty and loading states
5. Create list item component with click navigation
6. Add search/filter if needed with debounce

### Workflow C: Form Page Development

**Use for**: Create/edit forms

**Steps**:
1. Design form fields and validation rules
2. Define create/update/detail APIs
3. Implement form using wot-design-uni form components
4. Add form validation
5. Handle submit logic and success/error states
6. Handle special fields (image upload, date picker, cascader)

### Workflow D: Complex Business Pages

**Use for**: Custom pages, complex interactions, multi-step flows

**Steps**:
1. Clarify requirements (use `interview` skill if unclear, `tech-decision` for technical choices)
2. Design component structure, state management, interaction flow, multi-platform adaptation
3. Create page structure, split into sub-components (`components/`), implement business logic (`composables/`), integrate APIs, add state management
4. Multi-platform adaptation using conditional compilation (#ifdef H5 / MP-WEIXIN)
5. Test on all target platforms

### Workflow E: API Integration

**Use for**: Backend API integration

**Steps**:
1. Confirm backend API is implemented and swagger is generated/updated (source of truth is backend)
2. In app project, run `pnpm api:gen` to generate/refresh API functions + TypeScript types under `src/api`
3. Import and call the generated API functions/types (do not implement API functions manually)
4. Handle loading/error/empty states in the page (and add caching/retry via `@/http` options if needed)

### Workflow F: Performance Optimization

**Use for**: Improving page performance and UX

**Steps**:
1. Identify performance issues (page load speed, bottlenecks, UX problems)
2. Design optimization strategy (use `tech-decision` for comparison if needed)
3. Implement optimizations (image lazy loading, subpackages, alova cache, component lazy loading, virtual lists)
4. Verify with performance testing and bundle analysis

### Workflow G: Multi-Platform Adaptation

**Use for**: Ensuring compatibility across platforms

**Steps**:
1. Analyze platform differences (H5 browser APIs, mini-program auth/share, App native capabilities)
2. Use conditional compilation (#ifdef) for platform-specific code
3. Adapt features (login/auth, payment, sharing, file upload)
4. Test on all platforms

## Decision Tree

```
User Request
    │
    ├─ Standard page?
    │   ├─ List page → Workflow B
    │   ├─ Detail page → Workflow A
    │   └─ Form page → Workflow C
    │
    ├─ Complex business?
    │   └─ Workflow D
    │       ├─ Unclear requirements → interview
    │       └─ Technical choice → tech-decision
    │
    ├─ API integration?
    │   └─ Workflow E
    │
    ├─ Performance issue?
    │   └─ Workflow F
    │       └─ Compare solutions → tech-decision
    │
    └─ Multi-platform adaptation?
        └─ Workflow G
```

## Tech Stack

**Core**: uni-app 3.x + Vue 3 (Composition API) + TypeScript
**UI**: wot-design-uni + z-paging
**State**: Pinia + pinia-plugin-persistedstate
**HTTP**: alova (with cache and retry)
**Styles**: UnoCSS + SCSS

## Development Standards

### Directory Structure
```
src/pages/{module}/{name}/
├── index.vue              # Main page
├── components/            # Page components
└── composables/          # Composable functions
```

### API Codegen Rule (Required)
- Use `pnpm api:gen` (in app project) to generate API client code; API availability depends on backend implementation + swagger output.
- Do not manually implement or edit generated API files under `src/api/**` (they will be overwritten on next generation).
- If an endpoint/type is missing or incorrect, request backend to update the swagger, then re-run `pnpm api:gen`.

### API Interface Pattern
```typescript
// Generated by `pnpm api:gen` (do not edit generated files)
import { deleteAccount } from '@/api/v1/user'
import type { DeleteAccountReq, DeleteAccountReply } from '@/api/v1/user'

const req: DeleteAccountReq = {
  // body fields...
}

const res: DeleteAccountReply = await deleteAccount({ body: req })
```

### Page Development Standards
- Use Composition API with `<script setup lang="ts">`
- Avoid `any` type
- Use PascalCase for component names

### Multi-Platform Adaptation
```vue
<!-- #ifdef H5 -->
<view>H5 specific content</view>
<!-- #endif -->

<!-- #ifdef MP-WEIXIN -->
<view>WeChat mini-program specific content</view>
<!-- #endif -->
```

### Performance Best Practices
- Use webp format for images
- Use virtual scrolling for long lists
- Use subpackages appropriately
- Cache API requests
- Avoid frequent setData

## Common Commands

```bash
pnpm dev              # H5 development
pnpm dev:mp           # WeChat mini-program
pnpm dev:app          # App development
pnpm build            # Build H5
pnpm build:mp         # Build mini-program
pnpm type-check       # TypeScript check
pnpm lint             # Code lint
```

## Quality Checklist

Before completion, verify:
- [ ] Complete TypeScript types, no `any`
- [ ] API error handling
- [ ] Loading and empty states
- [ ] Form validation rules
- [ ] Image lazy loading
- [ ] Pagination or virtual scrolling for long lists
- [ ] H5 platform tested
- [ ] Mini-program platform tested
- [ ] ESLint passed
- [ ] TypeScript check passed

## Related Skills

- `interview` - Explore technical solutions
- `tech-decision` - Technical decision making
- `ui-ux-pro-max` - UI/UX design
- `bug-detective` - Debugging
- `git-workflow` - Git workflow
