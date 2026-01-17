<script lang="ts" setup>
import type { BannerItem, ListBannersParams } from '@/api/v1/banner/types'
import { listBanners } from '@/api/v1/banner/banner'
import HomeBanner from './components/HomeBanner.vue'
import HomeArticles from './components/HomeArticles.vue'

defineOptions({
  name: 'Home',
})

definePage({
  type: 'home',
  style: {
    navigationBarTitleText: '首页',
  },
})

const bannerPosition = 'home_top'
const messageEntryPath = '/pages-fg/message/category'

interface BannerDisplayItem {
  id?: string
  title?: string
  imageUrl?: string
  linkUrl?: string
  linkType?: string
  position?: string
  platform?: string
  sort?: number
}

// 轮播图数据
const bannerList = ref<BannerDisplayItem[]>([])

function resolveBannerPlatform() {
  // #ifdef H5
  return 'web'
  // #endif
  // #ifdef APP-PLUS
  return 'app'
  // #endif
  // #ifdef MP-WEIXIN
  return 'mp-weixin'
  // #endif
  // #ifdef MP-ALIPAY
  return 'mp-alipay'
  // #endif
  return ''
}

function toBannerDisplayItem(item: BannerItem): BannerDisplayItem {
  return {
    id: item.id,
    title: item.title,
    imageUrl: item.imageURL,
    linkUrl: item.linkURL,
    linkType: item.linkType,
    position: item.position,
    platform: item.platform,
    sort: item.sort,
  }
}

/**
 * 获取轮播图数据
 */
async function fetchBannerList() {
  try {
    const params: ListBannersParams = { position: bannerPosition }
    const platform = resolveBannerPlatform()
    if (platform) {
      params.platform = platform
    }

    const res = await listBanners({ params, options: {} })
    const list = res.list || []
    bannerList.value = list.map(toBannerDisplayItem).filter(item => !!item.imageUrl)
  }
  catch (error) {
    console.error('获取轮播图失败:', error)
  }
}

onLoad(() => {
  fetchBannerList()
})

/**
 * 轮播图点击
 */
function handleBannerClick(item: BannerDisplayItem) {
  const rawUrl = item.linkUrl?.trim()
  if (!rawUrl) {
    return
  }

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
  uni.navigateTo({ url: targetUrl })
}

function goToMessages() {
  uni.navigateTo({ url: messageEntryPath })
}
</script>

<template>
  <view class="home-container">
    <view class="home-header">
      <view class="home-title">
        <view class="home-brand">
          <view class="home-brand__sticker" />
          <view class="home-brand__glow" />
          <view class="home-brand__badge">
            <image class="home-brand__image" src="/static/logo.svg" mode="aspectFill" />
          </view>
        </view>
      </view>
      <view class="message-entry" @click="goToMessages">
        <wd-icon name="chat" size="40rpx" color="var(--fg-text)" />
      </view>
    </view>
    <!-- 轮播图 -->
    <HomeBanner
      v-if="bannerList.length > 0"
      :list="bannerList"
      @click="handleBannerClick"
    />

    <HomeArticles />
  </view>
</template>

<style lang="scss" scoped>
.home-container {
  min-height: 100vh;
  background: var(--fg-bg);
}

.home-header {
  padding: calc(env(safe-area-inset-top) + 24rpx) var(--fg-page-x) 12rpx;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.home-title {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.home-brand {
  position: relative;
  width: 72rpx;
  height: 72rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.home-brand__sticker {
  position: absolute;
  inset: -8rpx;
  border-radius: 20rpx 34rpx 18rpx 30rpx;
  background: var(--fg-white, #ffffff);
  border: 2px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 10rpx 22rpx rgba(0, 0, 0, 0.12);
  transform: rotate(-6deg) translate(-2rpx, 2rpx);
  z-index: 0;
  pointer-events: none;
}

.home-brand__glow {
  position: absolute;
  inset: -14rpx;
  border-radius: 28rpx;
  border: 2px solid rgba(var(--fg-primary-rgb), 0.32);
  background: radial-gradient(circle at 30% 30%, rgba(var(--fg-primary-rgb), 0.28) 0%, rgba(255, 255, 255, 0) 62%);
  box-shadow:
    0 0 0 2rpx rgba(var(--fg-primary-rgb), 0.12),
    0 12rpx 30rpx rgba(var(--fg-primary-rgb), 0.24),
    0 18rpx 40rpx rgba(var(--fg-primary-rgb), 0.18);
  pointer-events: none;
  z-index: 1;
}

.home-brand__badge {
  position: relative;
  width: 72rpx;
  height: 72rpx;
  border-radius: 18rpx 30rpx 20rpx 28rpx;
  background: var(--fg-surface);
  border: 2px solid rgba(var(--fg-primary-rgb), 0.25);
  box-shadow:
    0 14rpx 28rpx var(--fg-ink-08),
    inset 0 0 0 1rpx rgba(255, 255, 255, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  transform: rotate(-4deg);
  z-index: 2;
}

.home-brand__image {
  width: 52rpx;
  height: 52rpx;
  border-radius: 14rpx 20rpx 12rpx 18rpx;
  box-shadow: 0 6rpx 16rpx rgba(0, 0, 0, 0.12);
  transform: rotate(4deg);
}

.message-entry {
  width: 68rpx;
  height: 68rpx;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--fg-surface);
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}
</style>
