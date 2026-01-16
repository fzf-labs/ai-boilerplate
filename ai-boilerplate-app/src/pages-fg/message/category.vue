<script lang="ts" setup>
import { useToast } from 'wot-design-uni'
import { getUserMessageCategoryCounts } from '@/api/v1/user-message/userMessage'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '消息中心',
  },
})

const toast = useToast()
const tokenStore = useTokenStore()

const loading = ref(false)
const counts = ref({
  transactionUnread: 0,
  systemUnread: 0,
  serviceUnread: 0,
  totalUnread: 0,
})

const categories = computed(() => [
  {
    key: 'transaction',
    title: '交易信息',
    desc: '订单、支付与资产变动',
    icon: 'wallet',
    unread: counts.value.transactionUnread,
  },
  {
    key: 'system',
    title: '系统消息',
    desc: '系统更新与活动提醒',
    icon: 'setting',
    unread: counts.value.systemUnread,
  },
  {
    key: 'service',
    title: '客服消息',
    desc: '客服通知与服务进展',
    icon: 'chat',
    unread: counts.value.serviceUnread,
  },
])

function ensureLogin() {
  if (tokenStore.hasLogin)
    return true

  toast.warning('请先登录')
  setTimeout(() => {
    uni.navigateTo({ url: LOGIN_PAGE })
  }, 1500)
  return false
}

async function fetchCounts() {
  if (!ensureLogin())
    return

  loading.value = true
  try {
    const res = await getUserMessageCategoryCounts({ options: {} })
    counts.value = {
      transactionUnread: res.transactionUnread || 0,
      systemUnread: res.systemUnread || 0,
      serviceUnread: res.serviceUnread || 0,
      totalUnread: res.totalUnread || 0,
    }
  }
  catch (error) {
    console.error('加载消息未读数失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function goToCategory(category: string) {
  if (!ensureLogin())
    return
  uni.navigateTo({
    url: `/pages-fg/message/list?category=${category}`,
  })
}

onShow(() => {
  fetchCounts()
})
</script>

<template>
  <view class="message-page">
    <view class="top-bg" />
    <view class="content">
      <view class="page-header">
        <text class="header-title">消息中心</text>
        <text class="header-subtitle">查看最新通知与服务动态</text>
      </view>

      <view v-if="loading" class="loading-box">
        <wd-loading />
      </view>

      <view v-else class="category-list">
        <view
          v-for="item in categories"
          :key="item.key"
          class="category-card"
          @click="goToCategory(item.key)"
        >
          <view class="category-info">
            <view class="category-icon">
              <wd-icon :name="item.icon" size="44rpx" color="var(--wot-color-primary)" />
            </view>
            <view class="category-text">
              <text class="category-title">{{ item.title }}</text>
              <text class="category-desc">{{ item.desc }}</text>
            </view>
          </view>
          <view class="category-meta">
            <wd-tag
              v-if="item.unread > 0"
              type="danger"
              plain
              size="small"
            >
              {{ item.unread }} 未读
            </wd-tag>
            <text v-else class="category-read">全部已读</text>
            <text class="category-arrow">></text>
          </view>
        </view>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.message-page {
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

.page-header {
  margin-bottom: 24rpx;
}

.header-title {
  display: block;
  font-size: 40rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.header-subtitle {
  display: block;
  margin-top: 8rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.loading-box {
  padding: 32rpx 0;
  display: flex;
  justify-content: center;
}

.category-list {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.category-card {
  padding: 20rpx 24rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.category-info {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.category-icon {
  width: 64rpx;
  height: 64rpx;
  border-radius: 20rpx;
  background: var(--fg-surface-muted);
  display: flex;
  align-items: center;
  justify-content: center;
}

.category-text {
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.category-title {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.category-desc {
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.category-meta {
  display: flex;
  align-items: center;
  gap: 12rpx;
}

.category-read {
  font-size: 22rpx;
  color: var(--fg-text-weak);
}

.category-arrow {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}
</style>
