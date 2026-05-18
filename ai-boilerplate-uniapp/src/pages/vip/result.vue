<script lang="ts" setup>
import type { MallOrderInfo } from '@/api/v1/mall-order/types'
import { useToast } from 'wot-design-uni'
import { getOrderInfo } from '@/api/v1/mall-order/mallOrder'

definePage({
  style: {
    navigationBarTitleText: '支付结果',
  },
})

const toast = useToast()

// 支付状态
const status = ref<'success' | 'fail'>('success')
// 订单ID
const orderId = ref('')
// 订单信息
const orderInfo = ref<MallOrderInfo | null>(null)
// 加载状态
const loading = ref(false)

/**
 * 获取订单详情
 */
async function fetchOrderInfo() {
  if (!orderId.value)
    return

  try {
    loading.value = true
    const res = await getOrderInfo({
      params: { orderId: orderId.value },
      options: {},
    })
    orderInfo.value = res.info || null
  }
  catch (error) {
    console.error('获取订单信息失败:', error)
  }
  finally {
    loading.value = false
  }
}

/**
 * 格式化价格
 */
function formatPrice(p?: number) {
  if (p === undefined || p === null)
    return '0.00'
  return p.toFixed(2)
}

/**
 * 查看会员详情
 */
function handleViewMembership() {
  uni.redirectTo({
    url: '/pages/membership/detail',
  })
}

/**
 * 返回首页
 */
function handleGoHome() {
  uni.switchTab({
    url: '/pages/index/index',
  })
}

/**
 * 重新支付
 */
function handleRetry() {
  uni.navigateBack()
}

// 页面加载
onLoad((options) => {
  const opts = options as Record<string, string>
  status.value = (opts.status as 'success' | 'fail') || 'success'
  orderId.value = opts.orderId || ''

  if (orderId.value) {
    fetchOrderInfo()
  }
})
</script>

<template>
  <view class="result-page">
    <!-- 结果图标 -->
    <view class="result-header" :class="{ success: status === 'success', fail: status === 'fail' }">
      <view class="result-icon">
        <wd-icon
          :name="status === 'success' ? 'check' : 'close'"
          size="80rpx"
          color="var(--fg-text-inverse)"
        />
      </view>
      <view class="result-title">
        {{ status === 'success' ? '支付成功' : '支付失败' }}
      </view>
      <view class="result-desc">
        <template v-if="status === 'success'">
          恭喜您成功开通VIP会员
        </template>
        <template v-else>
          支付过程中出现问题，请重试
        </template>
      </view>
    </view>

    <!-- 订单信息 -->
    <view v-if="orderInfo" class="order-section">
      <view class="section-title">
        订单信息
      </view>
      <view class="order-card">
        <view class="order-row">
          <text class="label">订单编号</text>
          <text class="value">{{ orderInfo.id }}</text>
        </view>
        <view class="order-row">
          <text class="label">商品类型</text>
          <text class="value">VIP会员</text>
        </view>
        <view class="order-row">
          <text class="label">支付金额</text>
          <text class="value price">¥{{ formatPrice(orderInfo.actualAmount) }}</text>
        </view>
        <view class="order-row">
          <text class="label">支付方式</text>
          <text class="value">微信支付</text>
        </view>
        <view v-if="orderInfo.paymentTime" class="order-row">
          <text class="label">支付时间</text>
          <text class="value">{{ orderInfo.paymentTime }}</text>
        </view>
      </view>
    </view>

    <!-- 操作按钮 -->
    <view class="action-section">
      <template v-if="status === 'success'">
        <wd-button
          type="primary"
          size="large"

          round block
          @click="handleViewMembership"
        >
          查看会员权益
        </wd-button>
        <wd-button
          type="text"
          size="large"
          block
          custom-class="secondary-btn"
          @click="handleGoHome"
        >
          返回首页
        </wd-button>
      </template>
      <template v-else>
        <wd-button
          type="primary"
          size="large"

          round block
          @click="handleRetry"
        >
          重新支付
        </wd-button>
        <wd-button
          type="text"
          size="large"
          block
          custom-class="secondary-btn"
          @click="handleGoHome"
        >
          返回首页
        </wd-button>
      </template>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.result-page {
  min-height: 100vh;
  background: var(--fg-bg);
  padding: 0 var(--fg-page-x);
}

/* 结果头部 */
.result-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 80rpx 0 60rpx;
}

.result-icon {
  width: 140rpx;
  height: 140rpx;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 32rpx;
}

.result-header.success .result-icon {
  background: rgba(var(--fg-primary-rgb), 0.12);
  box-shadow: var(--fg-shadow-soft);
}

.result-header.fail .result-icon {
  background: rgba(255, 59, 48, 0.12);
  box-shadow: 0 12rpx 28rpx rgba(255, 59, 48, 0.2);
}

.result-title {
  font-size: 40rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 16rpx;
}

.result-desc {
  font-size: 28rpx;
  color: var(--fg-text-muted);
  text-align: center;
}

/* 订单信息 */
.order-section {
  margin-top: 20rpx;
}

.section-title {
  font-size: 30rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 20rpx;
}

.order-card {
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
  padding: 8rpx 24rpx;
}

.order-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24rpx 0;
  border-bottom: 1px solid var(--fg-border);
}

.order-row:last-child {
  border-bottom: none;
}

.order-row .label {
  font-size: 28rpx;
  color: var(--fg-text-muted);
}

.order-row .value {
  font-size: 28rpx;
  font-weight: 500;
  color: var(--fg-text);
  max-width: 60%;
  text-align: right;
  word-break: break-all;
}

.order-row .price {
  font-weight: 700;
  color: var(--fg-danger);
}

/* 操作按钮 */
.action-section {
  margin-top: 60rpx;
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

:deep(.secondary-btn) {
  color: var(--fg-text-muted) !important;
}
</style>
