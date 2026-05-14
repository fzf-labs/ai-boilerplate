# Best Practices

App 开发的代码规范和最佳实践。

## 目录

- [项目结构](#项目结构)
- [TypeScript 规范](#typescript-规范)
- [API 调用模式](#api-调用模式)
- [页面生命周期](#页面生命周期)
- [CSS 变量](#css-变量)
- [错误处理](#错误处理)
- [身份认证](#身份认证)

---

## 项目结构

```
src/
├── pages/           # Tabbar pages only
│   ├── index/       # Home page
│   ├── me/          # Profile page
│   └── help/        # Help pages
├── pages-fg/        # Sub-pages (non-tabbar)
│   ├── login/       # Auth pages
│   ├── profile/     # Profile edit
│   └── settings/    # Settings
├── components/      # Shared components
├── api/v1/          # Generated API (DO NOT EDIT)
├── store/           # Pinia stores
└── http/            # HTTP client
```

## TypeScript 规范

### 类型导入

```typescript
// Separate type imports
import type { UserInfo } from '@/api/v1/user/types'
import { getUserInfo } from '@/api/v1/user/user'
```

### Ref 类型

```typescript
// Explicit type for refs
const user = ref<UserInfo | null>(null)
const list = ref<ContentInfo[]>([])
const loading = ref(false)
```

### 计算属性

```typescript
const displayName = computed(() =>
  user.value?.nickname || 'Guest'
)
```

---

## API 调用模式

### 标准模式

```typescript
async function fetchData() {
  try {
    loading.value = true
    const res = await someApi({ options: {} })
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
```

### 带参数

```typescript
const res = await getList({
  params: { page: 1, pageSize: 10 },
  options: {},
})
```

### 带请求体

```typescript
const res = await createItem({
  body: { name: 'test', value: 123 },
  options: {},
})
```

---

## 页面生命周期

### definePage

```typescript
definePage({
  style: {
    navigationBarTitleText: 'Page Title',
  },
})
```

### onLoad

```typescript
onLoad((options) => {
  const { id } = options as Record<string, string>
  if (id) fetchDetail(id)
})
```

### onShow

```typescript
onShow(() => {
  // Refresh data when page shows
  if (tokenStore.hasLogin) {
    fetchUserData()
  }
})
```

---

## CSS 变量

### 颜色

```scss
var(--fg-text)          // Primary text
var(--fg-text-muted)    // Secondary text
var(--fg-text-weak)     // Tertiary text
var(--fg-primary)       // Primary color
var(--fg-bg)            // Background
var(--fg-surface)       // Card background
var(--fg-border)        // Border color
```

### 间距

```scss
var(--fg-page-x)        // Page horizontal padding
var(--fg-page-y)        // Page vertical padding
var(--fg-section-gap)   // Section gap
```

### 圆角和阴影

```scss
var(--fg-radius-card)   // Card radius
var(--fg-radius-lg)     // Large radius
var(--fg-shadow-card)   // Card shadow
var(--fg-shadow-soft)   // Soft shadow
```

---

## 错误处理

### Toast 消息

```typescript
toast.success('Saved')
toast.error('Failed')
toast.warning('Warning')
toast.info('Info')
```

### 表单验证

```typescript
const validateRes = await formRef.value?.validate()
if (validateRes && !validateRes.valid) {
  return // Stop if invalid
}
```

---

## 身份认证

### 检查登录状态

```typescript
import { useTokenStore } from '@/store/token'

const tokenStore = useTokenStore()

if (!tokenStore.hasLogin) {
  uni.navigateTo({ url: LOGIN_PAGE })
  return
}
```

### 受保护的操作

```typescript
function handleAction() {
  if (!tokenStore.hasLogin) {
    toast.warning('Please login first')
    setTimeout(() => {
      uni.navigateTo({ url: LOGIN_PAGE })
    }, 1500)
    return
  }
  // Continue action
}
```
