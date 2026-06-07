<script lang="ts" setup>
import type { HelpCategoryInfo } from '@/api/v1/help-category/types'
import type { HelpFaqInfo } from '@/api/v1/help-faq/types'
import { useToast } from 'wot-design-uni'
import { listHelpCategories } from '@/api/v1/help-category/helpCategory'
import { listHelpFaqs } from '@/api/v1/help-faq/helpFaq'
import CategoryGrid from './components/CategoryGrid.vue'
import ErrorState from './components/ErrorState.vue'
import HelpHeader from './components/HelpHeader.vue'
import HotFaqList from './components/HotFaqList.vue'

definePage({
  style: {
    navigationBarTitleText: '帮助中心',
  },
})

const toast = useToast()

// 状态管理
const state = reactive({
  categories: [] as HelpCategoryInfo[],
  hotFaqs: [] as HelpFaqInfo[],
  loading: false,
  refreshing: false,
  loadError: false,
  searchKeyword: '',
})

/**
 * 加载帮助分类
 */
async function loadCategories() {
  try {
    state.loading = true
    state.loadError = false
    const res = await listHelpCategories({ options: {} })
    state.categories = res.list || []
  }
  catch (error) {
    console.error('获取帮助分类失败:', error)
    state.loadError = true
    toast.error('加载失败，请重试')
  }
  finally {
    state.loading = false
  }
}

/**
 * 加载热门问题
 */
async function loadHotFaqs() {
  try {
    const res = await listHelpFaqs({
      params: {
        page: 1,
        pageSize: 5,
      },
      options: {},
    })
    state.hotFaqs = res.list || []
  }
  catch (error) {
    console.error('获取热门问题失败:', error)
  }
}

/**
 * 加载所有数据
 */
async function loadData() {
  await Promise.all([
    loadCategories(),
    loadHotFaqs(),
  ])
}

/**
 * 下拉刷新
 */
async function handleRefresh() {
  state.refreshing = true
  try {
    await loadData()
    toast.success('刷新成功')
  }
  catch (error) {
    console.error('刷新失败:', error)
  }
  finally {
    state.refreshing = false
  }
}

/**
 * 重试加载
 */
async function handleRetry() {
  await loadData()
}

/**
 * 搜索问题
 */
function handleSearch() {
  if (!state.searchKeyword.trim()) {
    toast.warning('请输入搜索关键词')
    return
  }
  uni.navigateTo({
    url: `/pages/help/faq?keyword=${encodeURIComponent(state.searchKeyword.trim())}`,
  })
}

/**
 * 选择分类
 */
function handleCategorySelect(categoryId: string) {
  const category = state.categories.find(c => c.id === categoryId)
  const categoryName = category?.name || '常见问题'
  uni.navigateTo({
    url: `/pages/help/faq?categoryId=${categoryId}&categoryName=${encodeURIComponent(categoryName)}`,
  })
}

/**
 * 选择FAQ
 */
function handleFaqSelect(faq: HelpFaqInfo) {
  const keyword = faq.question?.trim()
  uni.navigateTo({
    url: keyword ? `/pages/help/faq?keyword=${encodeURIComponent(keyword)}` : '/pages/help/faq',
  })
}

/**
 * 跳转到问题反馈
 */
function handleFeedback() {
  uni.navigateTo({
    url: '/pages/help/feedback',
  })
}

// 页面加载
onLoad(() => {
  loadData()
})
</script>

<template>
  <view class="help-page">
    <!-- 背景装饰 -->
    <view class="page-bg" />

    <!-- 下拉刷新容器 -->
    <scroll-view
      class="scroll-container"
      scroll-y
      refresher-enabled
      :refresher-triggered="state.refreshing"
      @refresherrefresh="handleRefresh"
    >
      <view class="page-content">
        <!-- 头部搜索 -->
        <HelpHeader
          v-model="state.searchKeyword"
          @search="handleSearch"
        />

        <!-- 错误状态 -->
        <ErrorState
          v-if="state.loadError && !state.loading"
          @retry="handleRetry"
        />

        <!-- 正常内容 -->
        <template v-else>
          <!-- 分类网格 -->
          <CategoryGrid
            :categories="state.categories"
            :loading="state.loading"
            @select="handleCategorySelect"
          />

          <!-- 热门问题 -->
          <HotFaqList
            :faqs="state.hotFaqs"
            :loading="state.loading"
            @select="handleFaqSelect"
          />

          <!-- 反馈按钮 -->
          <view class="feedback-section">
            <wd-button
              :block="true"
              :round="true"
              size="large"
              type="primary"
              @click="handleFeedback"
            >
              <wd-icon name="edit" size="32rpx" custom-style="margin-right: 8rpx" />
              提交问题反馈
            </wd-button>
          </view>
        </template>
      </view>
    </scroll-view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.help-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.page-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 240rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.scroll-container {
  height: 100vh;
}

.page-content {
  position: relative;
  padding: 20rpx var(--fg-page-x) 40rpx;
}

.feedback-section {
  margin-top: 20rpx;
  animation: fadeInUp 0.4s ease-out 0.3s backwards;

  :deep(.wd-button) {
    box-shadow: 0 8rpx 24rpx rgba(var(--wot-color-primary-rgb), 0.25);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

    &:active {
      transform: scale(0.98);
      box-shadow: 0 4rpx 12rpx rgba(var(--wot-color-primary-rgb), 0.2);
    }
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
