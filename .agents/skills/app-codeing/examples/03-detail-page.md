# Detail Page Pattern

This example shows how to create a detail page that fetches and displays a single item.

## Complete Example

```vue
<script lang="ts" setup>
import type { ContentDetail } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentDetail } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: 'Content Detail',
  },
})

const toast = useToast()
const content = ref<ContentDetail | null>(null)
const loading = ref(false)

/**
 * Fetch detail data
 */
async function fetchDetail(id: number) {
  try {
    loading.value = true
    const res = await getContentDetail({
      params: { id },
      options: {},
    })
    content.value = res.info || null
  }
  catch (error) {
    console.error('Failed to fetch detail:', error)
    toast.error('Load failed')
  }
  finally {
    loading.value = false
  }
}

onLoad((options) => {
  const idStr = (options as Record<string, string | undefined>).id
  const id = Number(idStr)

  if (!idStr || Number.isNaN(id)) {
    toast.error('Invalid parameter')
    setTimeout(() => uni.navigateBack(), 1200)
    return
  }

  fetchDetail(id)
})
</script>

<template>
  <view class="detail-page">
    <!-- Loading state -->
    <view v-if="loading" class="loading-state">
      <wd-loading />
      <text class="loading-text">Loading...</text>
    </view>

    <!-- Empty state -->
    <view v-else-if="!content" class="empty-state">
      <wd-icon name="warning" size="80rpx" />
      <text class="empty-text">Content not found</text>
    </view>

    <!-- Content -->
    <view v-else class="content">
      <image
        v-if="content.coverImage"
        :src="content.coverImage"
        class="cover"
        mode="aspectFill"
      />

      <view class="info">
        <text class="title">{{ content.title }}</text>

        <view class="meta">
          <text class="time">{{ content.publishTime }}</text>
          <view v-if="content.tags?.length" class="tags">
            <wd-tag
              v-for="tag in content.tags"
              :key="tag"
              type="primary"
              plain
              size="small"
            >
              {{ tag }}
            </wd-tag>
          </view>
        </view>

        <rich-text
          v-if="content.content"
          :nodes="content.content"
          class="body"
        />
      </view>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.detail-page {
  min-height: 100vh;
  background: var(--fg-bg);
}

.loading-state,
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 160rpx 0;
  gap: 20rpx;
}

.loading-text,
.empty-text {
  font-size: 28rpx;
  color: var(--fg-text-muted);
}

.content {
  background: var(--fg-surface);
}

.cover {
  width: 100%;
  height: 400rpx;
}

.info {
  padding: 32rpx var(--fg-page-x);
}

.title {
  font-size: 40rpx;
  font-weight: 700;
  color: var(--fg-text);
  line-height: 1.4;
  margin-bottom: 20rpx;
}

.meta {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 32rpx;
  padding-bottom: 24rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.time {
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.tags {
  display: flex;
  gap: 12rpx;
  flex-wrap: wrap;
}

.body {
  font-size: 30rpx;
  color: var(--fg-text);
  line-height: 1.8;
}
</style>
```

## Key Points

### 1. Route Parameter Handling

```typescript
onLoad((options) => {
  const idStr = (options as Record<string, string | undefined>).id
  const id = Number(idStr)

  // Validate parameter
  if (!idStr || Number.isNaN(id)) {
    toast.error('Invalid parameter')
    setTimeout(() => uni.navigateBack(), 1200)
    return
  }

  fetchDetail(id)
})
```

### 2. Three-State Pattern

Always handle these states:
- **Loading**: Show loading indicator
- **Empty/Error**: Show error message
- **Success**: Show content

```vue
<view v-if="loading">Loading...</view>
<view v-else-if="!content">Not found</view>
<view v-else>{{ content.title }}</view>
```

### 3. Navigation to Detail

From list page:
```typescript
uni.navigateTo({
  url: `/pages/content/detail?id=${item.id}`
})
```
