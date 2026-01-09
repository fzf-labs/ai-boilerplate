<script lang="ts" setup>
import type { BannerInfo, ContentInfo } from '@/api/v1/home/types'
import { getBannerList, getContentList } from '@/api/v1/home/home'

defineOptions({
  name: 'Home',
})

definePage({
  type: 'home',
  style: {
    navigationBarTitleText: '首页',
  },
})

const pagingRef = ref<any>(null)

// 轮播图数据
const bannerList = ref<BannerInfo[]>([])
// 内容列表数据
const contentList = ref<ContentInfo[]>([])

/**
 * 获取轮播图数据
 */
async function fetchBannerList() {
  try {
    const res = await getBannerList({ options: {} })
    bannerList.value = res.list || []
  }
  catch (error) {
    console.error('获取轮播图失败:', error)
  }
}

/**
 * 获取内容列表
 */
async function queryContentList(pageNo: number, pageSize: number) {
  try {
    if (pageNo === 1) {
      await fetchBannerList()
    }

    const res = await getContentList({
      params: {
        page: pageNo,
        pageSize,
      },
      options: {},
    })

    pagingRef.value?.completeByTotal(res.list || [], res.total ?? -1)
  }
  catch (error) {
    console.error('获取内容列表失败:', error)
    uni.showToast({
      title: '加载失败',
      icon: 'none',
    })
    pagingRef.value?.complete(false)
  }
}

/**
 * 轮播图点击
 */
function handleBannerClick(item: BannerInfo) {
  if (item.linkUrl) {
    // 跳转到详情页或外部链接
    console.log('跳转到:', item.linkUrl)
  }
}

/**
 * 内容卡片点击
 */
function handleContentClick(item: ContentInfo) {
  if (!item.id) {
    uni.showToast({
      title: '内容不存在',
      icon: 'none',
    })
    return
  }

  uni.navigateTo({ url: `/pages/content/detail?id=${item.id}` })
}

/**
 * 格式化时间
 */
function formatTime(time?: string) {
  if (!time)
    return ''

  const now = Date.now()
  const publishTime = new Date(time).getTime()
  if (Number.isNaN(publishTime))
    return time

  const diff = now - publishTime

  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour

  if (diff < minute) {
    return '刚刚'
  }
  else if (diff < hour) {
    return `${Math.floor(diff / minute)}分钟前`
  }
  else if (diff < day) {
    return `${Math.floor(diff / hour)}小时前`
  }
  else if (diff < 7 * day) {
    return `${Math.floor(diff / day)}天前`
  }
  else {
    return time.includes(' ') ? time.split(' ')[0] : time
  }
}
</script>

<template>
  <view class="home-container">
    <!-- 轮播图 -->
    <view v-if="bannerList.length > 0" class="banner-section">
      <wd-swiper
        :list="bannerList"
        :autoplay="bannerList.length > 1"
        :interval="4000"
        indicator-position="bottom"
        value-key="imageUrl"
        height="400rpx"
        @click="handleBannerClick"
      />
    </view>

    <!-- 内容列表 -->
    <z-paging
      ref="pagingRef"
      v-model="contentList"
      @query="queryContentList"
    >
      <view class="content-list">
        <wd-card
          v-for="item in contentList"
          :key="item.id"
          :title="item.title"
          :thumb="item.coverImage"
          @click="handleContentClick(item)"
        >
          <template #content>
            <view class="card-content">
              <text class="content-summary">{{ item.summary }}</text>
              <view class="content-meta">
                <text class="content-time">{{ formatTime(item.publishTime) }}</text>
                <view v-if="item.tags && item.tags.length > 0" class="content-tags">
                  <wd-tag
                    v-for="tag in item.tags"
                    :key="tag"
                    type="primary"
                    plain
                    size="small"
                  >
                    {{ tag }}
                  </wd-tag>
                </view>
              </view>
            </view>
          </template>
        </wd-card>
      </view>

      <!-- 空状态 -->
      <template #empty>
        <view class="empty-state">
          <text class="empty-text">暂无内容</text>
        </view>
      </template>
    </z-paging>
  </view>
</template>

<style lang="scss" scoped>
.home-container {
  min-height: 100vh;
  background: linear-gradient(180deg, var(--fg-bg-alt) 0%, var(--fg-surface-muted) 100%);
}

.banner-section {
  width: 100%;
  height: 400rpx;
  padding: var(--fg-page-x);
  box-sizing: border-box;

  :deep(.wd-swiper) {
    border-radius: var(--fg-radius-card-lg);
    overflow: hidden;
    box-shadow: var(--fg-shadow-card);
  }
}

.content-list {
  padding: 0 var(--fg-page-x) var(--fg-page-y);
  display: flex;
  flex-direction: column;
  gap: var(--fg-section-gap);

  :deep(.wd-card) {
    border-radius: var(--fg-radius-card);
    overflow: hidden;
    box-shadow: var(--fg-shadow-soft);
    transition: all 0.3s ease;
    background: var(--fg-surface);
    border: none;

    &:active {
      transform: scale(0.98);
      box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.08);
    }
  }

  :deep(.wd-card__thumb) {
    border-radius: var(--fg-radius-lg);
    overflow: hidden;
  }

  :deep(.wd-card__title) {
    font-size: 32rpx;
    font-weight: 600;
    color: var(--fg-text);
    line-height: 1.4;
  }
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
  padding-top: 12rpx;
}

.content-summary {
  font-size: 28rpx;
  color: var(--fg-text-muted);
  line-height: 1.6;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.content-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 8rpx;
  border-top: 1rpx solid var(--fg-border-weak);
}

.content-time {
  font-size: 24rpx;
  color: var(--fg-text-weak);
  font-weight: 500;
}

.content-tags {
  display: flex;
  gap: 12rpx;
  flex-wrap: wrap;

  :deep(.wd-tag) {
    border-radius: 12rpx;
    font-weight: 500;
  }
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 160rpx 0;
}

.empty-text {
  font-size: 28rpx;
  color: var(--fg-text-weak);
  font-weight: 500;
}
</style>
