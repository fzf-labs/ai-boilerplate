<script lang="ts" setup>
import type { ActivityItem } from '@/api/v1/activity/types'
import type { ContentInfo } from '@/api/v1/home/types'

import { useToast } from 'wot-design-uni'

import { listActivities } from '@/api/v1/activity/activity'
import { getContentList } from '@/api/v1/home/home'
import { isPageTabbar } from '@/tabbar/store'

definePage({
  style: {
    navigationBarTitleText: '发现',
  },
})

const toast = useToast()

const tabs = [
  { key: 'activity', label: '活动' },
  { key: 'article', label: '文章' },
] as const
const activeTab = ref<(typeof tabs)[number]['key']>(tabs[0].key)

const articleItems = ref<ContentInfo[]>([])
const articlePage = ref(1)
const activityItems = ref<ActivityItem[]>([])
const activityPage = ref(1)
const pageSize = 20
const articleTotal = ref(0)
const activityTotal = ref(0)
const articleLoading = ref(false)
const activityLoading = ref(false)
const articleLastFetchCount = ref(0)
const activityLastFetchCount = ref(0)

const hasMoreArticles = computed(() => {
  if (articleTotal.value > 0)
    return articleItems.value.length < articleTotal.value
  return articleLastFetchCount.value === pageSize
})

const hasMoreActivities = computed(() => {
  if (activityTotal.value > 0)
    return activityItems.value.length < activityTotal.value
  return activityLastFetchCount.value === pageSize
})

const featuredArticleItems = computed(() => articleItems.value.slice(0, 2))
const listArticleItems = computed(() => articleItems.value.slice(2))

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

async function fetchArticleList(reset = false) {
  if (articleLoading.value)
    return

  articleLoading.value = true
  const currentPage = reset ? 1 : articlePage.value
  try {
    const res = await getContentList({
      params: {
        page: currentPage,
        pageSize,
      },
      options: {},
    })
    const list = res.list || []
    articleTotal.value = res.total ?? articleTotal.value
    articleLastFetchCount.value = list.length

    if (reset) {
      articleItems.value = list
    }
    else {
      articleItems.value = articleItems.value.concat(list)
    }
    articlePage.value = currentPage + 1
  }
  catch (error) {
    console.error('加载文章列表失败:', error)
    toast.error('加载失败')
  }
  finally {
    articleLoading.value = false
  }
}

async function fetchActivityList(reset = false) {
  if (activityLoading.value)
    return

  activityLoading.value = true
  const currentPage = reset ? 1 : activityPage.value
  try {
    const res = await listActivities({
      params: {
        page: currentPage,
        pageSize,
      },
      options: {},
    })
    const list = res.list || []
    activityTotal.value = res.total ?? activityTotal.value
    activityLastFetchCount.value = list.length

    if (reset) {
      activityItems.value = list
    }
    else {
      activityItems.value = activityItems.value.concat(list)
    }
    activityPage.value = currentPage + 1
  }
  catch (error) {
    console.error('加载活动列表失败:', error)
    toast.error('加载失败')
  }
  finally {
    activityLoading.value = false
  }
}

function goDetail(item: ContentInfo) {
  if (!item.id)
    return
  uni.navigateTo({
    url: `/pages/content/detail?id=${encodeURIComponent(item.id)}`,
  })
}

function goActivity(item: ActivityItem) {
  const rawUrl = item.linkURL?.trim()
  if (!rawUrl)
    return

  const linkType = item.linkType?.toLowerCase() || ''
  const normalizedUrl = rawUrl.startsWith('app://') ? rawUrl.replace('app://', '/') : rawUrl
  const isExternal = /^https?:\/\//i.test(normalizedUrl)

  if (linkType === 'external' || isExternal) {
    const encodedUrl = encodeURIComponent(normalizedUrl)
    const encodedTitle = encodeURIComponent(item.title || '')
    uni.navigateTo({
      url: `/pages-fg/webview/index?url=${encodedUrl}&title=${encodedTitle}`,
    })
    return
  }

  const targetUrl = normalizedUrl.startsWith('/') ? normalizedUrl : `/${normalizedUrl}`
  if (isPageTabbar(targetUrl)) {
    uni.switchTab({ url: targetUrl })
    return
  }
  uni.navigateTo({ url: targetUrl })
}

function onTabClick(key: (typeof tabs)[number]['key']) {
  if (activeTab.value === key)
    return
  activeTab.value = key
  if (key === 'activity') {
    if (activityItems.value.length === 0)
      fetchActivityList(true)
  }
  else {
    if (articleItems.value.length === 0)
      fetchArticleList(true)
  }
}

onShow(() => {
  if (activeTab.value === 'activity') {
    fetchActivityList(true)
  }
  else {
    fetchArticleList(true)
  }
})

onReachBottom(() => {
  if (activeTab.value === 'activity') {
    if (activityLoading.value || !hasMoreActivities.value)
      return
    fetchActivityList(false)
    return
  }

  if (articleLoading.value || !hasMoreArticles.value)
    return
  fetchArticleList(false)
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
          @click="onTabClick(tab.key)"
        >
          <text class="tab-label">{{ tab.label }}</text>
        </view>
      </view>

      <view v-if="activeTab === 'activity'" class="tab-panel">
        <view class="activity-list-section">
          <view v-if="activityLoading && activityItems.length === 0" class="loading-box">
            <wd-loading />
          </view>

          <view v-else-if="activityItems.length === 0" class="empty-box">
            <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
            <text class="empty-text">暂无活动</text>
          </view>

          <template v-else>
            <view class="activity-list">
              <view
                v-for="item in activityItems"
                :key="item.id || item.linkURL"
                class="activity-item"
                @click="goActivity(item)"
              >
                <image
                  v-if="item.imageURL"
                  class="activity-cover"
                  :src="item.imageURL"
                  mode="aspectFill"
                />
                <view v-else class="activity-cover activity-cover--fallback">
                  <wd-icon name="picture" size="32rpx" color="var(--fg-text-weak)" />
                </view>
                <view class="activity-body">
                  <text class="activity-title">{{ item.title || '未命名活动' }}</text>
                  <text v-if="item.linkType" class="activity-subtitle">
                    {{ item.linkType.toLowerCase() === 'external' ? '外链活动' : '站内活动' }}
                  </text>
                </view>
              </view>
            </view>

            <view v-if="activityItems.length > 0" class="list-footer">
              <view v-if="activityLoading" class="loading-more">
                <wd-loading size="20rpx" />
                <text>加载中...</text>
              </view>
              <text v-else-if="!hasMoreActivities" class="no-more">没有更多了</text>
            </view>
          </template>
        </view>
      </view>

      <view v-if="activeTab === 'article'" class="tab-panel">
        <view class="article-list-section">
          <view v-if="articleLoading && articleItems.length === 0" class="loading-box">
            <wd-loading />
          </view>

          <view v-else-if="articleItems.length === 0" class="empty-box">
            <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
            <text class="empty-text">暂无文章</text>
          </view>

          <template v-else>
            <scroll-view
              v-if="featuredArticleItems.length > 0"
              scroll-x
              class="article-scroll"
              :show-scrollbar="false"
            >
              <view class="article-track">
                <view
                  v-for="item in featuredArticleItems"
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

            <view v-if="listArticleItems.length > 0" class="article-list">
              <view
                v-for="item in listArticleItems"
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

            <view v-if="articleItems.length > 0" class="list-footer">
              <view v-if="articleLoading" class="loading-more">
                <wd-loading size="20rpx" />
                <text>加载中...</text>
              </view>
              <text v-else-if="!hasMoreArticles" class="no-more">没有更多了</text>
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
  height: 280rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient-strong);
}

.content {
  position: relative;
  padding: 20rpx 0 48rpx;
}

.tab-header {
  display: flex;
  align-items: center;
  margin: 0 var(--fg-page-x);
  padding: 8rpx;
  gap: 8rpx;
  border-radius: var(--fg-radius-pill);
  background: var(--fg-surface-glass);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
}

.tab-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  padding: 10rpx 0;
  border-radius: var(--fg-radius-pill);
  transition:
    background 0.2s ease,
    color 0.2s ease;
}

.tab-label {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}

.tab-item.is-active .tab-label {
  color: var(--fg-text);
  font-weight: 600;
}

.tab-item.is-active {
  background: var(--fg-surface);
  box-shadow: 0 10rpx 20rpx var(--fg-ink-06);
}

.tab-panel {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
  padding-top: 16rpx;
}

.article-list-section {
  padding: 0 var(--fg-page-x);
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.activity-list-section {
  padding: 0 var(--fg-page-x);
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.activity-item {
  display: flex;
  flex-direction: column;
  gap: 14rpx;
  padding: 18rpx;
  border-radius: var(--fg-radius-card);
  background: linear-gradient(180deg, var(--fg-white) 0%, var(--fg-surface-muted) 100%);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
}

.activity-cover {
  width: 100%;
  height: 280rpx;
  border-radius: 22rpx;
  flex-shrink: 0;
  background: var(--fg-bg);
}

.activity-cover--fallback {
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed var(--fg-border);
  background: var(--fg-bg-alt);
}

.activity-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10rpx;
}

.activity-title {
  font-size: 32rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.activity-subtitle {
  align-self: flex-start;
  font-size: 22rpx;
  color: var(--fg-primary-600);
  line-height: 1.4;
  padding: 6rpx 16rpx;
  border-radius: 999rpx;
  background: rgba(var(--fg-primary-rgb), 0.12);
}

.article-scroll {
  width: 100%;
  -ms-overflow-style: none;
  scrollbar-width: none;
  margin-bottom: 4rpx;
}

.article-track {
  display: flex;
  gap: 18rpx;
  padding-bottom: 8rpx;
}

:deep(.article-scroll::-webkit-scrollbar) {
  width: 0;
  height: 0;
  display: none;
}

.article-card {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
  padding: 18rpx;
  width: 320rpx;
  min-height: 340rpx;
  border-radius: 28rpx;
  background: linear-gradient(180deg, var(--fg-white) 0%, var(--fg-surface-muted) 100%);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
  flex-shrink: 0;
}

.article-cover {
  width: 100%;
  height: 190rpx;
  border-radius: 20rpx;
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
  font-size: 28rpx;
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
  color: var(--fg-text-secondary);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-time {
  margin-top: auto;
  font-size: 20rpx;
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
  gap: 18rpx;
}

.article-item {
  display: flex;
  gap: 16rpx;
  padding: 18rpx;
  border-radius: 26rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
}

.item-cover {
  width: 168rpx;
  height: 124rpx;
  border-radius: 16rpx;
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
  color: var(--fg-text-secondary);
  line-height: 1.5;
}

.item-time {
  margin-top: auto;
  font-size: 20rpx;
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
