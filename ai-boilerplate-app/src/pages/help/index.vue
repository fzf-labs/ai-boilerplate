<script lang="ts" setup>
import type { HelpCategoryInfo } from '@/api/v1/help-category/types'
import type { HelpFaqInfo } from '@/api/v1/help-faq/types'
import { listHelpCategories } from '@/api/v1/help-category/helpCategory'
import { listHelpFaqs } from '@/api/v1/help-faq/helpFaq'

definePage({
  style: {
    navigationBarTitleText: '帮助中心',
  },
})

// 帮助分类列表
const categories = ref<HelpCategoryInfo[]>([])
// 热门问题列表
const hotFaqs = ref<HelpFaqInfo[]>([])
// 加载状态
const loading = ref(false)
// 搜索关键词
const searchKeyword = ref('')

/**
 * 获取帮助分类
 */
async function fetchCategories() {
  try {
    loading.value = true
    const res = await listHelpCategories({ options: {} })
    categories.value = res.list || []
  }
  catch (error) {
    console.error('获取帮助分类失败:', error)
    uni.showToast({
      title: '加载失败',
      icon: 'none',
    })
  }
  finally {
    loading.value = false
  }
}

/**
 * 获取热门问题
 */
async function fetchHotFaqs() {
  try {
    const res = await listHelpFaqs({
      params: {
        page: 1,
        pageSize: 5,
      },
      options: {},
    })
    hotFaqs.value = res.list || []
  }
  catch (error) {
    console.error('获取热门问题失败:', error)
  }
}

/**
 * 跳转到分类详情
 */
function goToCategoryDetail(categoryId?: string) {
  if (!categoryId) {
    uni.showToast({
      title: '分类不存在',
      icon: 'none',
    })
    return
  }
  uni.navigateTo({
    url: `/pages/help/faq?categoryId=${categoryId}`,
  })
}

function goToFaq(faq: HelpFaqInfo) {
  const keyword = faq.question?.trim()
  uni.navigateTo({
    url: keyword ? `/pages/help/faq?keyword=${encodeURIComponent(keyword)}` : '/pages/help/faq',
  })
}

/**
 * 跳转到问题反馈
 */
function goToFeedback() {
  uni.navigateTo({
    url: '/pages/help/feedback',
  })
}

/**
 * 搜索问题
 */
function handleSearch() {
  if (!searchKeyword.value.trim()) {
    uni.showToast({
      title: '请输入搜索关键词',
      icon: 'none',
    })
    return
  }
  uni.navigateTo({
    url: `/pages/help/faq?keyword=${encodeURIComponent(searchKeyword.value.trim())}`,
  })
}

onLoad(() => {
  fetchCategories()
  fetchHotFaqs()
})
</script>

<template>
  <view class="help-page">
    <view class="top-bg" />
    <view class="content">
      <view class="header-card">
        <view class="header-title">
          帮助中心
        </view>
        <view class="header-subtitle">
          搜索问题，或选择分类快速找到答案
        </view>
        <view class="header-search">
          <wd-search
            v-model="searchKeyword"
            :hide-cancel="true"
            placeholder="搜索问题（如：登录、密码、手机号）"
            @search="handleSearch"
          />
        </view>
      </view>

      <wd-card type="rectangle" custom-class="card">
        <template #title>
          <view class="card-title">
            <text class="card-title-text">帮助分类</text>
          </view>
        </template>

        <wd-skeleton
          theme="paragraph"
          :row-col="[1, 1, 1]"
          :loading="loading"
          animation="gradient"
        >
          <view v-if="categories.length === 0" class="empty">
            暂无分类
          </view>
          <wd-grid v-else :column="3" :border="false">
            <wd-grid-item
              v-for="category in categories"
              :key="category.id"
              @click="goToCategoryDetail(category.id)"
            >
              <view class="category-item">
                <view class="category-icon">
                  <wd-icon :name="category.icon || 'help'" size="46rpx" />
                </view>
                <text class="category-name">{{ category.name }}</text>
              </view>
            </wd-grid-item>
          </wd-grid>
        </wd-skeleton>
      </wd-card>

      <wd-card type="rectangle" custom-class="card card-gap">
        <template #title>
          <view class="card-title">
            <text class="card-title-text">热门问题</text>
            <wd-tag type="primary" plain>
              HOT
            </wd-tag>
          </view>
        </template>

        <wd-cell-group>
          <wd-cell
            v-for="faq in hotFaqs"
            :key="faq.id"
            :title="faq.question"
            is-link
            @click="goToFaq(faq)"
          />
          <view v-if="!loading && hotFaqs.length === 0" class="empty empty-pad">
            暂无热门问题
          </view>
        </wd-cell-group>
      </wd-card>

      <view class="feedback">
        <wd-button :block="true" :round="true" size="large" type="primary" @click="goToFeedback">
          提交问题反馈
        </wd-button>
      </view>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.help-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 260rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: 22rpx var(--fg-page-x) 40rpx;
}

.header-card {
  padding: 22rpx 18rpx 18rpx;
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  margin-bottom: 18rpx;
}

.header-title {
  font-size: 40rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.header-subtitle {
  margin-top: 10rpx;
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

.header-search {
  margin-top: 16rpx;
  background: var(--fg-bg-alt);
  border-radius: 20rpx;
  padding: 10rpx 10rpx;
  border: 1px solid var(--fg-border-weak);
}

.wd-card.card.is-rectangle {
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.card-gap {
  margin-top: 18rpx;
}

:deep(.card .wd-card__title-content) {
  padding: 18rpx 18rpx 0;
}

:deep(.card .wd-card__content) {
  padding: 16rpx 8rpx 18rpx;
}

.category-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 14rpx;
  padding: 18rpx 0;
}

.category-icon {
  width: 88rpx;
  height: 88rpx;
  border-radius: 22rpx;
  background: rgba(var(--fg-primary-rgb), 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.category-name {
  font-size: 26rpx;
  color: var(--fg-text-secondary);
  font-weight: 600;
}

.card-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.card-title-text {
  font-size: 32rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.feedback {
  margin-top: 18rpx;
}

.empty {
  color: var(--fg-text-muted);
  font-size: 26rpx;
  text-align: center;
}

.empty-pad {
  padding: 22rpx 0 8rpx;
}
</style>
