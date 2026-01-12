# Best Practices

Code standards and best practices for app development.

## Project Structure

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

## TypeScript Standards

### Type Imports

```typescript
// Separate type imports
import type { UserInfo } from '@/api/v1/user/types'
import { getUserInfo } from '@/api/v1/user/user'
```

### Ref Types

```typescript
// Explicit type for refs
const user = ref<UserInfo | null>(null)
const list = ref<ContentInfo[]>([])
const loading = ref(false)
```

### Computed Properties

```typescript
const displayName = computed(() =>
  user.value?.nickname || 'Guest'
)
```

---

## API Call Patterns

### Standard Pattern

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

### With Parameters

```typescript
const res = await getList({
  params: { page: 1, pageSize: 10 },
  options: {},
})
```

### With Body

```typescript
const res = await createItem({
  body: { name: 'test', value: 123 },
  options: {},
})
```

---

## Page Lifecycle

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

## CSS Variables

### Colors

```scss
var(--fg-text)          // Primary text
var(--fg-text-muted)    // Secondary text
var(--fg-text-weak)     // Tertiary text
var(--fg-primary)       // Primary color
var(--fg-bg)            // Background
var(--fg-surface)       // Card background
var(--fg-border)        // Border color
```

### Spacing

```scss
var(--fg-page-x)        // Page horizontal padding
var(--fg-page-y)        // Page vertical padding
var(--fg-section-gap)   // Section gap
```

### Radius & Shadow

```scss
var(--fg-radius-card)   // Card radius
var(--fg-radius-lg)     // Large radius
var(--fg-shadow-card)   // Card shadow
var(--fg-shadow-soft)   // Soft shadow
```

---

## Error Handling

### Toast Messages

```typescript
toast.success('Saved')
toast.error('Failed')
toast.warning('Warning')
toast.info('Info')
```

### Form Validation

```typescript
const validateRes = await formRef.value?.validate()
if (validateRes && !validateRes.valid) {
  return // Stop if invalid
}
```

---

## Authentication

### Check Login

```typescript
import { useTokenStore } from '@/store/token'

const tokenStore = useTokenStore()

if (!tokenStore.hasLogin) {
  uni.navigateTo({ url: LOGIN_PAGE })
  return
}
```

### Protected Actions

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
