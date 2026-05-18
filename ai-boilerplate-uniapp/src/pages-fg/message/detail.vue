<script lang="ts" setup>
import type { UserMessageInfo } from '@/api/v1/user-message/types'
import { useToast } from 'wot-design-uni'
import { getUserMessageInfo } from '@/api/v1/user-message/userMessage'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '消息详情',
  },
})

const toast = useToast()
const tokenStore = useTokenStore()

const message = ref<UserMessageInfo | null>(null)
const loading = ref(false)

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

async function fetchDetail(id: string) {
  if (!ensureLogin())
    return

  loading.value = true
  try {
    const res = await getUserMessageInfo({
      params: { id },
      options: {},
    })
    message.value = res.info || null
  }
  catch (error) {
    console.error('加载消息详情失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function handleLink() {
  const rawUrl = message.value?.linkURL?.trim()
  if (!rawUrl)
    return

  const normalizedUrl = rawUrl.startsWith('app://')
    ? rawUrl.replace('app://', '/')
    : rawUrl
  const isExternal = /^https?:\/\//i.test(normalizedUrl)

  if (isExternal) {
    const encodedUrl = encodeURIComponent(normalizedUrl)
    const encodedTitle = encodeURIComponent(message.value?.title || '')
    uni.navigateTo({
      url: `/pages-fg/webview/index?url=${encodedUrl}&title=${encodedTitle}`,
    })
    return
  }

  const targetUrl = normalizedUrl.startsWith('/')
    ? normalizedUrl
    : `/${normalizedUrl}`
  uni.navigateTo({ url: targetUrl })
}

onLoad((options) => {
  const id = options?.id
  if (!id) {
    toast.error('消息不存在')
    return
  }
  fetchDetail(id)
})
</script>

<template>
  <view class="message-detail-page">
    <view class="top-bg" />
    <view class="content">
      <view v-if="loading" class="loading-box">
        <wd-loading />
      </view>

      <view v-else-if="!message" class="empty-box">
        <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
        <text class="empty-text">暂无消息内容</text>
      </view>

      <view v-else class="detail-card">
        <text class="detail-title">{{ message.title || '未命名消息' }}</text>
        <text class="detail-time">{{ formatDateTime(message.sentAt) }}</text>
        <text v-if="message.summary" class="detail-summary">{{ message.summary }}</text>

        <image
          v-if="message.coverURL"
          class="detail-cover"
          :src="message.coverURL"
          mode="aspectFill"
        />

        <text class="detail-content">{{ message.content || '-' }}</text>

        <view v-if="message.linkURL" class="link-section">
          <wd-button :block="true" type="primary" @click="handleLink">
            查看详情
          </wd-button>
        </view>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.message-detail-page {
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

.detail-card {
  padding: 24rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.detail-title {
  font-size: 34rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.detail-time {
  font-size: 22rpx;
  color: var(--fg-text-weak);
}

.detail-summary {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}

.detail-cover {
  width: 100%;
  height: 320rpx;
  border-radius: 20rpx;
}

.detail-content {
  font-size: 28rpx;
  color: var(--fg-text);
  line-height: 1.6;
  white-space: pre-wrap;
}

.link-section {
  margin-top: 8rpx;
}
</style>
