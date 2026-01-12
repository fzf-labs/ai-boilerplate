<script lang="ts" setup>
import type { HelpFaqInfo } from '@/api/v1/help-faq/types'
import { useToast } from 'wot-design-uni'
import { listHelpFaqs } from '@/api/v1/help-faq/helpFaq'

definePage({
  style: {
    navigationBarTitleText: '常见问题',
  },
})

const toast = useToast()

// FAQ列表
const faqList = ref<HelpFaqInfo[]>([])
// 展开的问题ID
const expandedIds = ref<Set<string>>(new Set())
// 加载状态
const loading = ref(false)
// 搜索关键词
const searchKeyword = ref('')

/**
 * 获取FAQ列表
 */
async function fetchFaqs(categoryId?: string, keyword?: string) {
  try {
    loading.value = true
    const res = await listHelpFaqs({
      params: {
        categoryId,
        keyword,
        page: 1,
        pageSize: 50,
      },
      options: {},
    })
    faqList.value = res.list || []
  }
  catch (error) {
    console.error('获取FAQ失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

/**
 * 切换问题展开/折叠
 */
function toggleExpand(id?: string) {
  if (!id)
    return
  if (expandedIds.value.has(id)) {
    expandedIds.value.delete(id)
  }
  else {
    expandedIds.value.add(id)
  }
}

function isExpanded(id?: string) {
  return !!id && expandedIds.value.has(id)
}

/**
 * 提交反馈
 */
async function handleFeedback(faqId?: string, isHelpful?: boolean) {
  // TODO: 等待后端提供 FAQ 反馈接口
  toast.success('感谢您的反馈')
}

onLoad(async (options) => {
  const { categoryId, keyword, id } = options as Record<string, string | undefined>
  if (keyword) {
    searchKeyword.value = keyword
    await fetchFaqs(undefined, keyword)
  }
  else if (categoryId) {
    await fetchFaqs(categoryId)
  }
  else {
    await fetchFaqs()
  }

  if (id)
    expandedIds.value.add(id)
})
</script>

<template>
  <view class="faq-page">
    <view class="top-bg" />
    <view class="content">
      <!-- 优化后的标题区域 -->
      <view class="page-header">
        <view class="header-icon">
          <wd-icon name="chat" size="48rpx" color="var(--wot-color-primary)" />
        </view>
        <view class="header-info">
          <text class="header-title">常见问题</text>
          <text class="header-subtitle">找到您需要的答案</text>
        </view>
      </view>

      <view class="sheet">
        <view class="sheet-search">
          <wd-search
            v-model="searchKeyword"
            placeholder="搜索问题关键词"
            @search="fetchFaqs(undefined, searchKeyword)"
          />
        </view>

        <view v-if="loading" class="loading-container">
          <wd-loading />
        </view>

        <view v-else-if="faqList.length === 0" class="empty-container">
          <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
          <text class="empty-text">暂无相关问题</text>
        </view>

        <view v-else class="faq-list">
          <view
            v-for="(faq, index) in faqList"
            :key="faq.id || String(index)"
            class="faq-item"
            :class="{ 'is-expanded': isExpanded(faq.id) }"
          >
            <view class="faq-question" @click="toggleExpand(faq.id)">
              <view class="question-left">
                <view class="question-icon">
                  <wd-icon name="help-circle" size="36rpx" color="var(--wot-color-primary)" />
                </view>
                <text class="question-text">{{ faq.question || '（无标题）' }}</text>
              </view>
              <wd-icon
                :name="isExpanded(faq.id) ? 'arrow-up' : 'arrow-down'"
                size="32rpx"
                color="var(--fg-text-muted)"
              />
            </view>

            <view v-if="isExpanded(faq.id)" class="faq-answer">
              <view class="answer-content">
                <view class="answer-icon">
                  <wd-icon name="check-circle" size="32rpx" color="#52c41a" />
                </view>
                <text class="answer-text">{{ faq.answer || '（暂无答案）' }}</text>
              </view>
              <view class="feedback-section">
                <text class="feedback-label">这个回答有帮助吗？</text>
                <view class="feedback-buttons">
                  <wd-button size="small" type="success" plain @click="handleFeedback(faq.id, true)">
                    <wd-icon name="thumb-up" size="28rpx" />
                    有帮助
                  </wd-button>
                  <wd-button size="small" plain @click="handleFeedback(faq.id, false)">
                    <wd-icon name="thumb-down" size="28rpx" />
                    没帮助
                  </wd-button>
                </view>
              </view>
            </view>
          </view>
        </view>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.faq-page {
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

.page-header {
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 24rpx;
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  margin-bottom: 20rpx;
}

.header-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.1) 0%, rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.05) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.header-title {
  font-size: 36rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.header-subtitle {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}

.sheet-search {
  padding: 20rpx 16rpx;
  border-bottom: 1px solid var(--fg-border-weak);
  background: var(--fg-surface);
}

.loading-container {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 80rpx 0;
}

.empty-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80rpx 0;
  gap: 20rpx;
}

.empty-text {
  font-size: 28rpx;
  color: var(--fg-text-muted);
}

.faq-list {
  padding: 0;
}

.faq-item {
  position: relative;
  transition: all 0.3s ease;

  & + .faq-item {
    border-top: 1px solid var(--fg-border-weak);
  }

  &.is-expanded {
    background: rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.02);
  }
}

.faq-question {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24rpx 20rpx;
  cursor: pointer;
  transition: all 0.2s ease;

  &:active {
    background: rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.05);
  }
}

.question-left {
  flex: 1;
  display: flex;
  align-items: flex-start;
  gap: 16rpx;
}

.question-icon {
  flex-shrink: 0;
  margin-top: 4rpx;
}

.question-text {
  flex: 1;
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
  line-height: 1.5;
}

.faq-answer {
  padding: 0 20rpx 24rpx 20rpx;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.answer-content {
  display: flex;
  gap: 16rpx;
  padding: 20rpx;
  background: rgba(82, 196, 26, 0.05);
  border-radius: 16rpx;
  border-left: 4rpx solid #52c41a;
}

.answer-icon {
  flex-shrink: 0;
  margin-top: 4rpx;
}

.answer-text {
  flex: 1;
  font-size: 28rpx;
  color: var(--fg-text-secondary);
  line-height: 1.7;
  white-space: pre-wrap;
}

.feedback-section {
  margin-top: 24rpx;
  padding-top: 20rpx;
  border-top: 1px dashed var(--fg-border-weak);
}

.feedback-label {
  display: block;
  font-size: 26rpx;
  color: var(--fg-text-muted);
  margin-bottom: 16rpx;
  font-weight: 500;
}

.feedback-buttons {
  display: flex;
  gap: 16rpx;
}
</style>
