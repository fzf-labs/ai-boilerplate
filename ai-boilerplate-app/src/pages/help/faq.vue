<script lang="ts" setup>
import type { HelpFaqInfo } from '@/api/v1/help-faq/types'
import { listHelpFaqs } from '@/api/v1/help-faq/helpFaq'

definePage({
  style: {
    navigationBarTitleText: '常见问题',
  },
})

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
  uni.showToast({
    title: '感谢您的反馈',
    icon: 'success',
  })
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
      <view class="sheet">
        <view class="sheet-search">
          <wd-search
            v-model="searchKeyword"
            placeholder="搜索问题"
            @search="fetchFaqs(undefined, searchKeyword)"
          />
        </view>

        <view class="faq-list">
          <view
            v-for="(faq, index) in faqList"
            :key="faq.id || String(index)"
            class="faq-item"
          >
            <view class="faq-question" @click="toggleExpand(faq.id)">
              <text class="question-text">{{ faq.question || '（无标题）' }}</text>
              <wd-icon
                :name="isExpanded(faq.id) ? 'arrow-up' : 'arrow-down'"
                size="32rpx"
              />
            </view>

            <view v-if="isExpanded(faq.id)" class="faq-answer">
              <text class="answer-text">{{ faq.answer || '（暂无答案）' }}</text>
              <view class="feedback-buttons">
                <text class="feedback-label">这个回答有帮助吗？</text>
                <wd-button size="small" @click="handleFeedback(faq.id, true)">
                  有帮助
                </wd-button>
                <wd-button size="small" plain @click="handleFeedback(faq.id, false)">
                  没帮助
                </wd-button>
              </view>
            </view>
          </view>
        </view>
      </view>
    </view>
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

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}

.sheet-search {
  padding: 14rpx 12rpx;
  border-bottom: 1px solid var(--fg-border-weak);
  background: var(--fg-surface);
}

.faq-list {
  padding: 0;
}

.faq-item {
  position: relative;

  & + .faq-item {
    border-top: 1px solid var(--fg-border-weak);
  }
}

.faq-question {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 22rpx 18rpx;
}

.question-text {
  flex: 1;
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.faq-answer {
  padding: 0 18rpx 22rpx;
}

.answer-text {
  font-size: 28rpx;
  color: var(--fg-text-muted);
  line-height: 1.6;
}

.feedback-buttons {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-top: 24rpx;
}

.feedback-label {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}
</style>
