<script lang="ts" setup>
import type { ContentDetail } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentDetail } from '@/api/v1/home/home'

definePage({
  style: {
    navigationBarTitleText: '内容详情',
  },
})

const toast = useToast()
const content = ref<ContentDetail | null>(null)
const loading = ref(false)

function formatDateTime(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
}

function formatCount(value?: number) {
  if (value == null)
    return ''
  if (value < 1000)
    return String(value)
  const formatted = (value / 1000).toFixed(value >= 10000 ? 0 : 1).replace(/\.0$/, '')
  return `${formatted}k`
}

const formattedPublishTime = computed(() => formatDateTime(content.value?.publishTime))

async function fetchContentDetail(id: string) {
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
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

onLoad((options) => {
  const id = (options as Record<string, string | undefined>).id
  if (!id) {
    toast.error('参数错误')
    setTimeout(() => uni.navigateBack(), 1200)
    return
  }

  fetchContentDetail(id)
})
</script>

<template>
  <view class="detail-container">
    <view class="page-glow" />
    <view v-if="loading" class="loading">
      加载中...
    </view>

    <view v-else-if="!content" class="empty">
      <text>内容不存在</text>
    </view>

    <view v-else class="content">
      <view class="hero-card">
        <image
          v-if="content.coverImage"
          :src="content.coverImage"
          class="hero-cover"
          mode="aspectFill"
        />
        <view v-if="content.coverImage" class="hero-overlay" />
        <view class="hero-body" :class="{ 'is-overlay': !!content.coverImage }">
          <text class="hero-title">{{ content.title || '未命名内容' }}</text>
          <view
            v-if="formattedPublishTime || content.viewCount !== undefined || content.likeCount !== undefined || (content.tags && content.tags.length)"
            class="hero-meta"
          >
            <view
              v-if="formattedPublishTime || content.viewCount !== undefined || content.likeCount !== undefined"
              class="meta-row"
            >
              <view v-if="formattedPublishTime" class="meta-pill">
                <wd-icon name="time" size="26rpx" color="var(--fg-text-weak)" />
                <text>{{ formattedPublishTime }}</text>
              </view>
              <view v-if="content.viewCount !== undefined" class="meta-pill">
                <wd-icon name="view" size="26rpx" color="var(--fg-text-weak)" />
                <text>{{ formatCount(content.viewCount) }}</text>
              </view>
              <view v-if="content.likeCount !== undefined" class="meta-pill">
                <wd-icon name="thumb-up" size="26rpx" color="var(--fg-text-weak)" />
                <text>{{ formatCount(content.likeCount) }}</text>
              </view>
            </view>
            <view v-if="content.tags && content.tags.length" class="tags">
              <wd-tag v-for="tag in content.tags" :key="tag" type="primary" plain size="small">
                {{ tag }}
              </wd-tag>
            </view>
          </view>
        </view>
      </view>

      <view class="body-card">
        <rich-text v-if="content.content" :nodes="content.content" class="rich-content" />
        <view v-else class="empty-body">暂无内容</view>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.detail-container {
  min-height: 100vh;
  background: var(--fg-bg-alt);
  padding: 24rpx var(--fg-page-x) 40rpx;
  position: relative;
}

.loading,
.empty {
  padding: 120rpx 0;
  text-align: center;
  color: var(--fg-text-muted);
  font-size: 28rpx;
  position: relative;
  z-index: 1;
}

.page-glow {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  height: 240rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
  z-index: 0;
}

.content {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
  position: relative;
  z-index: 1;
}

.hero-card {
  position: relative;
  border-radius: var(--fg-radius-card);
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  transform: translateZ(0);
}

.hero-cover {
  width: 100%;
  height: 460rpx;
  display: block;
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(0, 0, 0, 0.05) 0%, rgba(0, 0, 0, 0.42) 100%);
}

.hero-body {
  position: relative;
  padding: 28rpx;
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.hero-body.is-overlay {
  position: absolute;
  left: 20rpx;
  right: 20rpx;
  bottom: 20rpx;
  padding: 20rpx 22rpx;
  border-radius: var(--fg-radius-card);
  background: var(--fg-surface-glass);
  border: 1rpx solid var(--fg-glass-70);
  box-shadow: var(--fg-shadow-soft);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
}

.hero-title {
  font-size: 40rpx;
  font-weight: 700;
  color: var(--fg-text);
  line-height: 1.4;
  letter-spacing: -0.2rpx;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.hero-meta {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
}

.meta-row {
  display: flex;
  align-items: center;
  gap: 10rpx;
  flex-wrap: wrap;
}

.meta-pill {
  display: flex;
  align-items: center;
  gap: 8rpx;
  padding: 6rpx 14rpx;
  border-radius: 999rpx;
  background: var(--fg-ink-04);
  border: 1rpx solid var(--fg-border);
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.hero-body.is-overlay .meta-pill {
  background: var(--fg-glass-70);
  border-color: var(--fg-glass-70);
}

.tags {
  display: flex;
  gap: 10rpx;
  flex-wrap: wrap;
}

.body-card {
  background: var(--fg-surface);
  border-radius: var(--fg-radius-card);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  padding: 28rpx;
}

.rich-content {
  font-size: 28rpx;
  color: var(--fg-text);
  line-height: 1.8;
}

.rich-content :deep(p) {
  margin: 0 0 22rpx;
}

.rich-content :deep(h1),
.rich-content :deep(h2),
.rich-content :deep(h3) {
  margin: 28rpx 0 16rpx;
  color: var(--fg-text);
  font-weight: 700;
  line-height: 1.4;
}

.rich-content :deep(h1) {
  font-size: 36rpx;
}

.rich-content :deep(h2) {
  font-size: 32rpx;
}

.rich-content :deep(h3) {
  font-size: 30rpx;
}

.rich-content :deep(ul),
.rich-content :deep(ol) {
  padding-left: 32rpx;
  margin: 0 0 22rpx;
}

.rich-content :deep(li) {
  margin: 10rpx 0;
}

.rich-content :deep(blockquote) {
  margin: 24rpx 0;
  padding: 20rpx 22rpx;
  background: var(--fg-ink-04);
  border-left: 6rpx solid var(--fg-primary);
  border-radius: 16rpx;
  color: var(--fg-text-secondary);
}

.rich-content :deep(img) {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 16rpx 0;
  border-radius: 16rpx;
}

.empty-body {
  padding: 24rpx 0;
  text-align: center;
  color: var(--fg-text-weak);
  font-size: 26rpx;
}
</style>
