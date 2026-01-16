<script lang="ts" setup>
import type { BannerItem, ListBannersParams } from '@/api/v1/banner/types'
import { listBanners } from '@/api/v1/banner/banner'
import HomeBanner from './components/HomeBanner.vue'

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
</script>

<template>
  <view class="home-container">
    <view class="home-header">
      <text class="home-title fg-large-title">首页</text>
    </view>
    <!-- 轮播图 -->
    <HomeBanner
      v-if="bannerList.length > 0"
      :list="bannerList"
      @click="handleBannerClick"
    />

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.home-container {
  min-height: 100vh;
  background: var(--fg-bg);
}

.home-header {
  padding: calc(env(safe-area-inset-top) + 24rpx) var(--fg-page-x) 12rpx;
}

.home-title {
  line-height: 1.15;
}
</style>
