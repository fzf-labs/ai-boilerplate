# wot-design-uni Components Guide

Quick reference for commonly used wot-design-uni components.

## Layout Components

### wd-card

```vue
<wd-card
  title="Card Title"
  thumb="https://example.com/image.jpg"
  type="rectangle"
  custom-class="my-card"
>
  <template #content>
    <text>Card content</text>
  </template>
  <template #footer>
    <text>Footer</text>
  </template>
</wd-card>
```

### wd-cell / wd-cell-group

```vue
<wd-cell-group>
  <wd-cell
    title="Title"
    label="Description"
    value="Value"
    icon="user"
    is-link
    @click="handleClick"
  />
</wd-cell-group>
```

### wd-divider

```vue
<wd-divider>Section Title</wd-divider>
<wd-divider :hairline="true" />
```

---

## Form Components

### wd-form

```vue
<wd-form ref="formRef" :model="form" :rules="rules" error-type="toast">
  <wd-form-item label="Name" prop="name" required>
    <wd-input v-model="form.name" placeholder="Enter name" />
  </wd-form-item>
</wd-form>
```

### wd-input

```vue
<wd-input
  v-model="value"
  placeholder="Placeholder"
  type="text"
  prefix-icon="user"
  clearable
  show-password
  :maxlength="50"
  @confirm="handleConfirm"
/>
```

### wd-textarea

```vue
<wd-textarea
  v-model="value"
  placeholder="Enter content"
  :maxlength="200"
  show-word-limit
  :rows="4"
/>
```

### wd-picker

```vue
<wd-picker
  v-model="value"
  :columns="options"
  label="Select"
  placeholder="Please select"
/>
```

### wd-datetime-picker

```vue
<wd-datetime-picker
  v-model="date"
  type="date"
  label="Date"
  placeholder="Select date"
/>
```

### wd-checkbox / wd-radio

```vue
<wd-checkbox v-model="checked" shape="square">
  Remember me
</wd-checkbox>

<wd-radio-group v-model="selected">
  <wd-radio value="1">Option 1</wd-radio>
  <wd-radio value="2">Option 2</wd-radio>
</wd-radio-group>
```

### wd-switch

```vue
<wd-switch v-model="enabled" />
```

---

## Feedback Components

### wd-toast

```typescript
import { useToast } from 'wot-design-uni'

const toast = useToast()
toast.success('Success')
toast.error('Error')
toast.warning('Warning')
toast.info('Info')
toast.loading('Loading...')
```

```vue
<template>
  <!-- Must include in template -->
  <wd-toast />
</template>
```

### wd-button

```vue
<wd-button
  type="primary"
  size="large"
  :block="true"
  :loading="loading"
  :disabled="disabled"
  round
  @click="handleClick"
>
  Submit
</wd-button>
```

Button types: `primary`, `success`, `warning`, `danger`, `default`

### wd-loading

```vue
<wd-loading type="ring" color="var(--wot-color-primary)" />
```

### wd-tag

```vue
<wd-tag type="primary" plain size="small">
  Tag Text
</wd-tag>
```

### wd-icon

```vue
<wd-icon name="user" size="48rpx" color="var(--fg-primary)" />
```

---

## Display Components

### wd-swiper

```vue
<wd-swiper
  :list="bannerList"
  :autoplay="true"
  :interval="4000"
  indicator-position="bottom"
  value-key="imageUrl"
  height="400rpx"
  @click="handleClick"
/>
```

### wd-collapse

```vue
<wd-collapse v-model="activeNames">
  <wd-collapse-item title="Section 1" name="1">
    Content 1
  </wd-collapse-item>
  <wd-collapse-item title="Section 2" name="2">
    Content 2
  </wd-collapse-item>
</wd-collapse>
```

### wd-search

```vue
<wd-search
  v-model="keyword"
  placeholder="Search..."
  @search="handleSearch"
/>
```

---

## Navigation Components

### wd-navbar

```vue
<wd-navbar
  title="Page Title"
  left-arrow
  @click-left="uni.navigateBack()"
/>
```

---

## z-paging (List Component)

```vue
<z-paging
  ref="pagingRef"
  v-model="list"
  @query="queryList"
>
  <!-- List items -->
  <view v-for="item in list" :key="item.id">
    {{ item.name }}
  </view>

  <!-- Empty slot -->
  <template #empty>
    <view>No data</view>
  </template>
</z-paging>
```

Methods:
- `pagingRef.value?.completeByTotal(list, total)`
- `pagingRef.value?.complete(list)`
- `pagingRef.value?.complete(false)` - on error
- `pagingRef.value?.reload()` - refresh
