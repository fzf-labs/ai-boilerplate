# List Page with Pagination

This example shows how to create a paginated list page using z-paging.

## Complete Example

```vue
<script lang="ts" setup>
import type { ContentInfo } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentList } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: 'Content List',
  },
})

const toast = useToast()
const pagingRef = ref<any>(null)
const contentList = ref<ContentInfo[]>([])

/**
 * Query list data - called by z-paging
 */
async function queryList(pageNo: number, pageSize: number) {
  try {
    const res = await getContentList({
      params: {
        page: pageNo,
        pageSize,
      },
      options: {},
    })

    // Complete with total count for accurate pagination
    pagingRef.value?.completeByTotal(res.list || [], res.total ?? -1)
  }
  catch (error) {
    console.error('Failed to fetch list:', error)
    toast.error('Load failed')
    pagingRef.value?.complete(false)
  }
}

/**
 * Handle item click
 */
function handleItemClick(item: ContentInfo) {
  if (!item.id) {
    toast.warning('Invalid item')
    return
  }
  uni.navigateTo({ url: `/pages/content/detail?id=${item.id}` })
}
</script>

<template>
  <view class="list-page">
    <z-paging
      ref="pagingRef"
      v-model="contentList"
      @query="queryList"
    >
      <view class="list-container">
        <wd-card
          v-for="item in contentList"
          :key="item.id"
          :title="item.title"
          :thumb="item.coverImage"
          @click="handleItemClick(item)"
        >
          <template #content>
            <view class="card-content">
              <text class="summary">{{ item.summary }}</text>
              <view class="meta">
                <text class="time">{{ item.publishTime }}</text>
              </view>
            </view>
          </template>
        </wd-card>
      </view>

      <!-- Empty state slot -->
      <template #empty>
        <view class="empty-state">
          <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
          <text class="empty-text">No content yet</text>
        </view>
      </template>
    </z-paging>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.list-page {
  min-height: 100vh;
  background: var(--fg-bg);
}

.list-container {
  padding: var(--fg-page-x);
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.summary {
  font-size: 28rpx;
  color: var(--fg-text-muted);
  line-height: 1.6;
}

.meta {
  display: flex;
  justify-content: space-between;
}

.time {
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 160rpx 0;
  gap: 20rpx;
}

.empty-text {
  font-size: 28rpx;
  color: var(--fg-text-muted);
}
</style>
```

## Key Points

### 1. z-paging Setup

```vue
<z-paging
  ref="pagingRef"
  v-model="contentList"
  @query="queryList"
>
```

- `ref`: Reference for calling methods
- `v-model`: Binds to list data
- `@query`: Called on refresh/load-more

### 2. Query Function Signature

```typescript
async function queryList(pageNo: number, pageSize: number) {
  // pageNo: Current page (1-based)
  // pageSize: Items per page (default 10)
}
```

### 3. Completion Methods

```typescript
// With total count (recommended)
pagingRef.value?.completeByTotal(list, total)

// Without total (uses list length)
pagingRef.value?.complete(list)

// On error
pagingRef.value?.complete(false)
```

### 4. Manual Refresh

```typescript
// Trigger refresh programmatically
pagingRef.value?.reload()
```
