<script lang="ts" setup>
import type { ContentInfo } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentList } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: '发现',
  },
})

const toast = useToast()

const tabs = [
  { key: 'article', label: '文章' },
] as const
const activeTab = ref<(typeof tabs)[number]['key']>(tabs[0].key)

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
const featuredItems = computed(() => items.value.slice(0, 2))
const listItems = computed(() => items.value.slice(2))

function formatDateTime(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
}

function formatShortDate(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}.${pad(date.getMonth() + 1)}.${pad(date.getDate())}`
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
      <view class="tab-header">
        <view
          v-for="tab in tabs"
          :key="tab.key"
          class="tab-item"
          :class="{ 'is-active': activeTab === tab.key }"
          @click="activeTab = tab.key"
        >
          <text class="tab-label">{{ tab.label }}</text>
          <view v-if="activeTab === tab.key" class="tab-underline" />
        </view>
      </view>

      <view v-if="activeTab === 'article'" class="tab-panel">
        <view class="article-list-section">
          <view v-if="loading && items.length === 0" class="loading-box">
            <wd-loading />
          </view>

          <view v-else-if="items.length === 0" class="empty-box">
            <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
            <text class="empty-text">暂无文章</text>
          </view>

          <template v-else>
            <scroll-view
              v-if="featuredItems.length > 0"
              scroll-x
              class="article-scroll"
              :show-scrollbar="false"
            >
              <view class="article-track">
                <view
                  v-for="item in featuredItems"
                  :key="item.id"
                  class="article-card"
                  @click="goDetail(item)"
                >
                  <image
                    v-if="item.coverImage"
                    class="article-cover"
                    :src="item.coverImage"
                    mode="aspectFill"
                  />
                  <view v-else class="article-cover article-cover--fallback">
                    <wd-icon name="file" size="32rpx" color="var(--fg-text-weak)" />
                  </view>
                  <text class="article-title">{{ item.title || '未命名文章' }}</text>
                  <text v-if="item.summary" class="article-summary">{{ item.summary }}</text>
                  <text v-if="item.publishTime" class="article-time">{{ formatShortDate(item.publishTime) }}</text>
                </view>
              </view>
            </scroll-view>

            <view v-if="listItems.length > 0" class="article-list">
              <view
                v-for="item in listItems"
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
          </template>
        </view>
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
  padding: 24rpx 0 40rpx;
}

.tab-header {
  display: flex;
  gap: 28rpx;
  padding: 0 var(--fg-page-x) 16rpx;
}

.tab-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 8rpx 0;
  gap: 6rpx;
}

.tab-label {
  font-size: 28rpx;
  color: var(--fg-text-weak);
}

.tab-item.is-active .tab-label {
  color: var(--fg-text);
  font-weight: 600;
}

.tab-underline {
  width: 42rpx;
  height: 6rpx;
  border-radius: 999rpx;
  background: var(--fg-primary);
  box-shadow: 0 6rpx 12rpx rgba(var(--fg-primary-rgb), 0.24);
}

.tab-panel {
  display: flex;
  flex-direction: column;
  gap: 8rpx;
}

.article-list-section {
  padding: 0 var(--fg-page-x);
}

.article-scroll {
  width: 100%;
  -ms-overflow-style: none;
  scrollbar-width: none;
  margin-bottom: 12rpx;
}

.article-track {
  display: flex;
  gap: 16rpx;
  padding-bottom: 6rpx;
}

:deep(.article-scroll::-webkit-scrollbar) {
  width: 0;
  height: 0;
  display: none;
}

.article-card {
  display: flex;
  flex-direction: column;
  gap: 10rpx;
  padding: 16rpx;
  width: 300rpx;
  min-height: 320rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  flex-shrink: 0;
}

.article-cover {
  width: 100%;
  height: 170rpx;
  border-radius: 18rpx;
  background: var(--fg-bg);
}

.article-cover--fallback {
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed var(--fg-border);
  background: var(--fg-bg-alt);
}

.article-title {
  font-size: 26rpx;
  font-weight: 600;
  color: var(--fg-text);
  line-height: 1.35;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-summary {
  font-size: 22rpx;
  color: var(--fg-text-weak);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-time {
  margin-top: auto;
  font-size: 22rpx;
  color: var(--fg-text-muted);
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
