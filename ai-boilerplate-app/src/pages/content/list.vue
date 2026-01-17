<script lang="ts" setup>
import type { ContentInfo } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentList } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: '文章',
  },
})

const toast = useToast()

const items = ref<ContentInfo[]>([])
const page = ref(1)
const pageSize = 20
const total = ref(0)
const loading = ref(false)
const lastFetchCount = ref(0)

const hasMore = computed(() => {
  if (total.value > 0)
    return items.value.length < total.value
  return lastFetchCount.value === pageSize
})

function formatDateTime(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
}

async function fetchList(reset = false) {
  if (loading.value)
    return

  loading.value = true
  const currentPage = reset ? 1 : page.value
  try {
    const res = await getContentList({
      params: {
        page: currentPage,
        pageSize,
      },
      options: {},
    })
    const list = res.list || []
    total.value = res.total ?? total.value
    lastFetchCount.value = list.length

    if (reset) {
      items.value = list
    }
    else {
      items.value = items.value.concat(list)
    }
    page.value = currentPage + 1
  }
  catch (error) {
    console.error('加载文章列表失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function goDetail(item: ContentInfo) {
  if (!item.id)
    return
  uni.navigateTo({
    url: `/pages/content/detail?id=${encodeURIComponent(item.id)}`,
  })
}

onShow(() => {
  fetchList(true)
})

onReachBottom(() => {
  if (loading.value || !hasMore.value)
    return
  fetchList(false)
})
</script>

<template>
  <view class="article-list-page">
    <view class="top-bg" />
    <view class="content">
      <view v-if="loading && items.length === 0" class="loading-box">
        <wd-loading />
      </view>

      <view v-else-if="items.length === 0" class="empty-box">
        <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
        <text class="empty-text">暂无文章</text>
      </view>

      <view v-else class="article-list">
        <view
          v-for="item in items"
          :key="item.id"
          class="article-item"
          @click="goDetail(item)"
        >
          <image
            v-if="item.coverImage"
            class="item-cover"
            :src="item.coverImage"
            mode="aspectFill"
          />
          <view class="item-body">
            <text class="item-title">{{ item.title || '未命名文章' }}</text>
            <text v-if="item.summary" class="item-summary">{{ item.summary }}</text>
            <text v-if="item.publishTime" class="item-time">{{ formatDateTime(item.publishTime) }}</text>
          </view>
        </view>
      </view>

      <view v-if="items.length > 0" class="list-footer">
        <view v-if="loading" class="loading-more">
          <wd-loading size="20rpx" />
          <text>加载中...</text>
        </view>
        <text v-else-if="!hasMore" class="no-more">没有更多了</text>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.article-list-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 240rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: 24rpx var(--fg-page-x) 40rpx;
}

.list-header {
  margin-bottom: 16rpx;
}

.list-title {
  display: block;
  font-size: 38rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.list-subtitle {
  display: block;
  margin-top: 6rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.loading-box {
  padding: 32rpx 0;
  display: flex;
  justify-content: center;
}

.empty-box {
  padding: 40rpx 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12rpx;
}

.empty-text {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}

.article-list {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.article-item {
  display: flex;
  gap: 16rpx;
  padding: 20rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.item-cover {
  width: 180rpx;
  height: 140rpx;
  border-radius: 18rpx;
  flex-shrink: 0;
}

.item-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8rpx;
}

.item-title {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.item-summary {
  font-size: 24rpx;
  color: var(--fg-text-weak);
  line-height: 1.5;
}

.item-time {
  margin-top: auto;
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.list-footer {
  padding: 24rpx 0 0;
  display: flex;
  justify-content: center;
}

.loading-more {
  display: flex;
  align-items: center;
  gap: 12rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.no-more {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}
</style>
