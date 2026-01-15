<script lang="ts" setup>
import type { ActivationCodeRedemptionInfo } from '@/api/v1/mall-activation-code/types'
import { useToast } from 'wot-design-uni'
import { activateMembershipByCode, listActivationCodeRedemptions } from '@/api/v1/mall-activation-code/mallActivationCode'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '激活码兑换',
  },
})

const toast = useToast()
const tokenStore = useTokenStore()

const code = ref('')
const submitting = ref(false)
const records = ref<ActivationCodeRedemptionInfo[]>([])
const page = ref(1)
const pageSize = 20
const total = ref(0)
const historyLoading = ref(false)
const lastFetchCount = ref(0)

const hasMore = computed(() => {
  if (total.value > 0)
    return records.value.length < total.value
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

async function handleSubmit() {
  if (!ensureLogin())
    return

  const normalized = code.value.trim()
  if (!normalized) {
    toast.warning('请输入激活码')
    return
  }
  if (normalized.length > 64) {
    toast.warning('激活码长度不能超过64位')
    return
  }
  if (submitting.value)
    return

  submitting.value = true
  code.value = normalized
  try {
    await activateMembershipByCode({
      body: {
        code: normalized,
      },
    })
    toast.success('兑换成功')
    setTimeout(() => {
      uni.navigateTo({
        url: '/pages/membership/detail',
      })
    }, 300)
  }
  catch (error: any) {
    const message = error?.data?.message || error?.data?.msg || error?.message || '兑换失败'
    toast.error(message)
  }
  finally {
    submitting.value = false
  }
}

function formatDate(value?: string) {
  if (!value)
    return '-'
  return value.split('T')[0]
}

async function fetchHistory(reset = false) {
  if (historyLoading.value)
    return
  if (!ensureLogin())
    return

  historyLoading.value = true
  const currentPage = reset ? 1 : page.value
  try {
    const res = await listActivationCodeRedemptions({
      params: {
        page: currentPage,
        pageSize,
      },
      options: {},
    })
    const list = res.list || []
    total.value = res.total ?? total.value
    lastFetchCount.value = list.length
    if (reset) {
      records.value = list
    }
    else {
      records.value = records.value.concat(list)
    }
    page.value = currentPage + 1
  }
  catch (error) {
    console.error('获取兑换记录失败:', error)
    toast.error('加载失败')
  }
  finally {
    historyLoading.value = false
  }
}

onLoad(() => {
  ensureLogin()
})

onShow(() => {
  if (tokenStore.hasLogin)
    fetchHistory(true)
})

onReachBottom(() => {
  if (historyLoading.value || !hasMore.value)
    return
  fetchHistory(false)
})
</script>

<template>
  <view class="activation-page">
    <view class="top-bg" />
    <view class="content">
      <view class="page-header">
        <view class="header-icon">
          <wd-icon name="edit" size="48rpx" color="var(--wot-color-primary)" />
        </view>
        <view class="header-info">
          <text class="header-title">兑换激活码</text>
          <text class="header-subtitle">输入会员激活码即可开通权益</text>
        </view>
      </view>

      <view class="sheet">
        <view class="form-section">
          <view class="form-label">
            <text class="label-text">激活码</text>
          </view>
          <wd-input
            v-model="code"
            placeholder="请输入激活码"
            clearable
            :maxlength="64"
          />
          <text class="form-tip">激活码区分大小写，请确认输入无误</text>
        </view>

        <view class="sheet-footer pb-safe">
          <wd-button
            :block="true"
            :round="true"
            size="large"
            type="primary"
            :loading="submitting"
            @click="handleSubmit"
          >
            立即兑换
          </wd-button>
        </view>

        <view class="history-section">
          <view class="history-title">
            兑换记录
          </view>

          <view v-if="historyLoading && records.length === 0" class="loading-container">
            <wd-loading />
          </view>

          <view v-else-if="records.length === 0" class="empty-container">
            <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
            <text class="empty-text">暂无兑换记录</text>
          </view>

          <view v-else class="history-list">
            <view v-for="(item, index) in records" :key="item.code || `${index}`" class="history-card">
              <view class="card-header">
                <text class="code-text">{{ item.code || '-' }}</text>
                <text class="activated-text">{{ formatDate(item.activatedAt) }}</text>
              </view>
            </view>
          </view>

          <view v-if="records.length > 0" class="history-footer">
            <view v-if="historyLoading" class="loading-more">
              <wd-loading size="20rpx" />
              <text>加载中...</text>
            </view>
            <text v-else-if="!hasMore" class="no-more">没有更多了</text>
          </view>
        </view>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.activation-page {
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
  border-radius: 24rpx;
  background: linear-gradient(
    135deg,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.12) 0%,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.04) 100%
  );
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-info {
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.header-title {
  font-size: 32rpx;
  font-weight: 700;
  color: var(--fg-text);
}

.header-subtitle {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  padding: 24rpx;
}

.form-section {
  display: flex;
  flex-direction: column;
  gap: 14rpx;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 10rpx;
}

.label-text {
  font-size: 26rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.form-tip {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.sheet-footer {
  margin-top: 28rpx;
}

.history-section {
  margin-top: 32rpx;
  padding-top: 24rpx;
  border-top: 1px solid var(--fg-border);
}

.history-title {
  font-size: 26rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 16rpx;
}

.loading-container,
.empty-container {
  padding: 40rpx 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16rpx;
}

.empty-text {
  color: var(--fg-text-muted);
  font-size: 24rpx;
}

.history-list {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.history-card {
  background: rgba(255, 255, 255, 0.02);
  border-radius: 20rpx;
  border: 1px solid var(--fg-border);
  padding: 18rpx 20rpx;
  display: flex;
  flex-direction: column;
  gap: 14rpx;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12rpx;
}

.code-text {
  font-size: 26rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.activated-text {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.history-footer {
  padding: 20rpx 0 8rpx;
  text-align: center;
}

.loading-more {
  display: inline-flex;
  align-items: center;
  gap: 8rpx;
  color: var(--fg-text-muted);
  font-size: 22rpx;
}

.no-more {
  color: var(--fg-text-muted);
  font-size: 22rpx;
}
</style>
