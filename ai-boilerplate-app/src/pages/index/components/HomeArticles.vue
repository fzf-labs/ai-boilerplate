<script lang="ts" setup>
import type { ContentInfo } from '@/api/v1/home/types'
import { useToast } from 'wot-design-uni'
import { getContentList } from '@/api/v1/home/home'

const props = defineProps({
  entryPath: {
    type: String,
    default: '/pages/content/list',
  },
  detailPath: {
    type: String,
    default: '/pages/content/detail',
  },
  pageSize: {
    type: Number,
    default: 10,
  },
  title: {
    type: String,
    default: '文章推荐',
  },
})

const toast = useToast()
const articleList = ref<ContentInfo[]>([])
const articleLoading = ref(false)

function formatShortDate(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}.${pad(date.getMonth() + 1)}.${pad(date.getDate())}`
}

async function fetchArticleList() {
  if (articleLoading.value)
    return

  try {
    articleLoading.value = true
    const res = await getContentList({
      params: {
        page: 1,
        pageSize: props.pageSize,
      },
      options: {},
    })
    articleList.value = res.list || []
  }
  catch (error) {
    console.error('加载首页文章失败:', error)
    toast.error('加载文章失败')
  }
  finally {
    articleLoading.value = false
  }
}

function goToArticles() {
  uni.navigateTo({ url: props.entryPath })
}

function goToArticleDetail(item: ContentInfo) {
  if (!item.id)
    return
  uni.navigateTo({
    url: `${props.detailPath}?id=${encodeURIComponent(item.id)}`,
  })
}

onMounted(() => {
  fetchArticleList()
})
</script>

<template>
  <view class="article-section">
    <view class="section-head">
      <text class="section-title">{{ title }}</text>
      <text class="section-link" @click="goToArticles">更多</text>
    </view>

    <view v-if="articleLoading && articleList.length === 0" class="article-loading">
      <wd-loading size="28rpx" />
    </view>
    <view v-else-if="articleList.length === 0" class="article-empty">
      暂无文章
    </view>
    <scroll-view v-else scroll-x class="article-scroll" :show-scrollbar="false">
      <view class="article-track">
        <view
          v-for="item in articleList"
          :key="item.id"
          class="article-card"
          @click="goToArticleDetail(item)"
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

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.article-section {
  padding: 24rpx var(--fg-page-x) 40rpx;
}

.section-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16rpx;
}

.section-title {
  font-size: 36rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.section-link {
  font-size: 26rpx;
  color: var(--fg-primary);
}

.article-loading,
.article-empty {
  padding: 20rpx 0;
  display: flex;
  justify-content: center;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.article-scroll {
  width: 100%;
  -ms-overflow-style: none;
  scrollbar-width: none;
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
</style>
