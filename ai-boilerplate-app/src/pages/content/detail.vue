<script lang="ts" setup>
import type { ContentDetail } from '@/api/v1/home/types'
import { getContentDetail } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: '内容详情',
  },
})

const content = ref<ContentDetail | null>(null)
const loading = ref(false)

async function fetchContentDetail(id: number) {
  try {
    loading.value = true
    const res = await getContentDetail({
      params: { id },
      options: {},
    })
    content.value = res.info || null
  }
  catch (error) {
    console.error('获取内容详情失败:', error)
    uni.showToast({
      title: '加载失败',
      icon: 'none',
    })
  }
  finally {
    loading.value = false
  }
}

onLoad((options) => {
  const idStr = (options as Record<string, string | undefined>).id
  const id = Number(idStr)
  if (!idStr || Number.isNaN(id)) {
    uni.showToast({
      title: '参数错误',
      icon: 'none',
    })
    setTimeout(() => uni.navigateBack(), 1200)
    return
  }

  fetchContentDetail(id)
})
</script>

<template>
  <view class="detail-container">
    <view v-if="loading" class="loading">
      加载中...
    </view>

    <view v-else-if="!content" class="empty">
      <text>内容不存在</text>
    </view>

    <view v-else class="content">
      <image v-if="content.coverImage" :src="content.coverImage" class="cover" mode="aspectFill" />
      <view class="title">
        {{ content.title }}
      </view>
      <view class="meta">
        <text v-if="content.publishTime">{{ content.publishTime }}</text>
        <view v-if="content.tags && content.tags.length" class="tags">
          <wd-tag v-for="tag in content.tags" :key="tag" type="primary" plain size="small">
            {{ tag }}
          </wd-tag>
        </view>
      </view>
      <rich-text v-if="content.content" :nodes="content.content" />
    </view>
  </view>
</template>

<style lang="scss" scoped>
.detail-container {
  min-height: 100vh;
  background: var(--fg-bg-alt);
  padding: var(--fg-page-x);
}

.loading,
.empty {
  padding: 120rpx 0;
  text-align: center;
  color: var(--fg-text-muted);
  font-size: 28rpx;
}

.content {
  background: var(--fg-surface);
  border-radius: var(--fg-radius-card);
  overflow: hidden;
  padding: 32rpx;
}

.cover {
  width: 100%;
  height: 360rpx;
  border-radius: var(--fg-radius-lg);
  margin-bottom: 24rpx;
}

.title {
  font-size: 36rpx;
  font-weight: 700;
  color: var(--fg-text);
  line-height: 1.4;
  margin-bottom: 16rpx;
}

.meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 24rpx;
  color: var(--fg-text-weak);
  font-size: 24rpx;
}

.tags {
  display: flex;
  gap: 12rpx;
  flex-wrap: wrap;
}
</style>
