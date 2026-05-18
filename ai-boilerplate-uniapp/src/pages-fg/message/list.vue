<script lang="ts" setup>
import type { UserMessageInfo } from '@/api/v1/user-message/types'
import { useToast } from 'wot-design-uni'
import { getUserMessageList } from '@/api/v1/user-message/userMessage'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '消息列表',
  },
})

const toast = useToast()
const tokenStore = useTokenStore()

const category = ref('system')
const messages = ref<UserMessageInfo[]>([])
const page = ref(1)
const pageSize = 20
const total = ref(0)
const loading = ref(false)
const lastFetchCount = ref(0)

const categoryTitle = computed(() => {
  switch (category.value) {
    case 'transaction':
      return '交易信息'
    case 'service':
      return '客服消息'
    default:
      return '系统消息'
  }
})

const hasMore = computed(() => {
  if (total.value > 0)
    return messages.value.length < total.value
  return lastFetchCount.value === pageSize
})

function ensureLogin() {
  if (tokenStore.hasLogin)
    return true

  toast.warning('请先登录')
  setTimeout(() => {
    uni.navigateTo({ url: LOGIN_PAGE })
  }, 1500)
  return false
}

function formatDateTime(value?: string) {
  if (!value)
    return ''
  const date = new Date(value)
  if (Number.isNaN(date.getTime()))
    return value
  const pad = (num: number) => String(num).padStart(2, '0')
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
}

async function fetchList(reset = false) {
  if (loading.value)
    return
  if (!ensureLogin())
    return

  loading.value = true
  const currentPage = reset ? 1 : page.value
  try {
    const res = await getUserMessageList({
      params: {
        page: currentPage,
        pageSize,
        category: category.value,
      },
      options: {},
    })
    const list = res.list || []
    total.value = res.total ?? total.value
    lastFetchCount.value = list.length
    if (reset) {
      messages.value = list
    }
    else {
      messages.value = messages.value.concat(list)
    }
    page.value = currentPage + 1
  }
  catch (error) {
    console.error('加载消息列表失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function goDetail(item: UserMessageInfo) {
  if (!item.id)
    return
  uni.navigateTo({
    url: `/pages-fg/message/detail?id=${item.id}`,
  })
}

onLoad((options) => {
  category.value = options?.category || 'system'
  uni.setNavigationBarTitle({ title: categoryTitle.value })
})

onShow(() => {
  fetchList(true)
})

onReachBottom(() => {
  if (loading.value || !hasMore.value)
    return
  fetchList(false)
})
</script>

<template>
  <view class="message-list-page">
    <view class="top-bg" />
    <view class="content">
      <view class="list-header">
        <text class="list-title">{{ categoryTitle }}</text>
        <text class="list-subtitle">共 {{ total || messages.length }} 条</text>
      </view>

      <view v-if="loading && messages.length === 0" class="loading-box">
        <wd-loading />
      </view>

      <view v-else-if="messages.length === 0" class="empty-box">
        <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
        <text class="empty-text">暂无消息</text>
      </view>

      <view v-else class="message-list">
        <view
          v-for="item in messages"
          :key="item.id"
          class="message-item"
          @click="goDetail(item)"
        >
          <view class="item-head">
            <text
              class="item-title"
              :class="{ unread: !item.readAt }"
            >
              {{ item.title || '未命名消息' }}
            </text>
            <text class="item-time">{{ formatDateTime(item.sentAt) }}</text>
          </view>
          <text v-if="item.summary" class="item-summary">{{ item.summary }}</text>
          <view class="item-footer">
            <text class="item-status" :class="{ unread: !item.readAt }">
              {{ item.readAt ? '已读' : '未读' }}
            </text>
          </view>
        </view>
      </view>

      <view v-if="messages.length > 0" class="list-footer">
        <view v-if="loading" class="loading-more">
          <wd-loading size="20rpx" />
          <text>加载中...</text>
        </view>
        <text v-else-if="!hasMore" class="no-more">没有更多了</text>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.message-list-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 240rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: 24rpx var(--fg-page-x) 40rpx;
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

.message-list {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.message-item {
  padding: 20rpx 24rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.item-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12rpx;
}

.item-title {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
  max-width: 70%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.item-title.unread {
  color: var(--fg-primary-600);
}

.item-time {
  font-size: 22rpx;
  color: var(--fg-text-weak);
}

.item-summary {
  display: block;
  margin-top: 12rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.item-footer {
  margin-top: 12rpx;
  display: flex;
  justify-content: flex-end;
}

.item-status {
  font-size: 22rpx;
  color: var(--fg-text-weak);
}

.item-status.unread {
  color: var(--fg-primary-600);
}

.list-footer {
  padding: 20rpx 0 0;
  text-align: center;
  color: var(--fg-text-weak);
  font-size: 22rpx;
}

.loading-more {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8rpx;
}
</style>
