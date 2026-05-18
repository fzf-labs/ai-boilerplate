<script lang="ts" setup>
import type { MallOrderInfo } from '@/api/v1/mall-order/types'
import { useToast } from 'wot-design-uni'
import { getOrderInfo } from '@/api/v1/mall-order/mallOrder'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '订单详情',
  },
  excludeLoginPath: true,
})

const toast = useToast()
const tokenStore = useTokenStore()

const orderId = ref('')
const orderInfo = ref<MallOrderInfo | null>(null)
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

function formatPrice(value?: number) {
  if (value === undefined || value === null)
    return '0.00'
  return value.toFixed(2)
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

function getProductTypeText(productType?: string) {
  switch (productType) {
    case 'membership':
      return '会员'
    case 'service':
      return '服务'
    case 'goods':
      return '商品'
    default:
      return productType || '未知'
  }
}

function getPaymentMethodText(method?: string) {
  switch (method) {
    case 'wechat':
      return '微信'
    case 'alipay':
      return '支付宝'
    default:
      return method || '未知'
  }
}

function getPaymentStatusText(status?: number) {
  switch (status) {
    case 0:
      return '待支付'
    case 1:
      return '已支付'
    case 2:
      return '支付失败'
    case 3:
      return '已退款'
    default:
      return status === undefined ? '' : String(status)
  }
}

function getOrderStatusText(status?: string) {
  switch (status) {
    case 'pendingPayment':
      return '待付款'
    case 'pendingDelivery':
      return '待发货'
    case 'pendingReceipt':
      return '待收货'
    case 'completed':
      return '已完成'
    case 'canceled':
      return '已取消'
    case 'refunded':
      return '已退款'
    default:
      return status || '未知'
  }
}

async function fetchOrder() {
  if (loading.value)
    return
  if (!ensureLogin())
    return
  if (!orderId.value) {
    toast.error('订单ID错误')
    return
  }

  loading.value = true
  try {
    const res = await getOrderInfo({
      params: { orderId: orderId.value },
      options: {},
    })
    orderInfo.value = res.info || null
    if (!orderInfo.value) {
      toast.error('订单不存在')
    }
  }
  catch (error) {
    console.error('加载订单详情失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function handleBack() {
  uni.navigateBack()
}

onLoad((options) => {
  orderId.value = (options as Record<string, string> | undefined)?.orderId || ''
})

onShow(() => {
  fetchOrder()
})
</script>

<template>
  <view class="order-detail-page">
    <view class="top-bg" />

    <view class="content">
      <view v-if="loading && !orderInfo" class="loading-box">
        <wd-loading />
      </view>

      <view v-else-if="!orderInfo" class="empty-box">
        <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
        <text class="empty-text">未找到订单信息</text>
        <wd-button type="primary" size="large" round block @click="handleBack">
          返回
        </wd-button>
      </view>

      <view v-else class="detail">
        <view class="summary-card">
          <view class="summary-head">
            <view class="summary-status" :class="`status-${orderInfo.status || 'unknown'}`">
              <text class="summary-status-text">{{ getOrderStatusText(orderInfo.status) }}</text>
            </view>
            <view class="summary-amount">
              <text class="amount-label">支付金额</text>
              <text class="amount-value">¥{{ formatPrice(orderInfo.actualAmount) }}</text>
            </view>
          </view>
          <view class="summary-meta">
            <view class="meta-item">
              <text class="meta-label">订单号</text>
              <text class="meta-value">{{ orderInfo.id }}</text>
            </view>
            <view class="meta-item">
              <text class="meta-label">创建时间</text>
              <text class="meta-value">{{ formatDateTime(orderInfo.createdAt) }}</text>
            </view>
          </view>
        </view>

        <view class="section-card">
          <view class="section-title">
            订单信息
          </view>
          <view class="cell-list">
            <view class="cell-row">
              <text class="cell-label">订单状态</text>
              <text class="cell-value strong">{{ getOrderStatusText(orderInfo.status) }}</text>
            </view>
            <view class="cell-row">
              <text class="cell-label">类型</text>
              <text class="cell-value">{{ getProductTypeText(orderInfo.productType) }}</text>
            </view>
            <view v-if="orderInfo.remark" class="cell-row multiline">
              <text class="cell-label">备注</text>
              <text class="cell-value wrap">{{ orderInfo.remark }}</text>
            </view>
          </view>
        </view>

        <view class="section-card">
          <view class="section-title">
            支付信息
          </view>
          <view class="cell-list">
            <view class="cell-row">
              <text class="cell-label">支付状态</text>
              <text class="cell-value">{{ getPaymentStatusText(orderInfo.paymentStatus) }}</text>
            </view>
            <view class="cell-row">
              <text class="cell-label">支付方式</text>
              <text class="cell-value">{{ getPaymentMethodText(orderInfo.paymentMethod) }}</text>
            </view>
            <view class="cell-row">
              <text class="cell-label">支付金额</text>
              <text class="cell-value price">¥{{ formatPrice(orderInfo.actualAmount) }}</text>
            </view>
            <view v-if="orderInfo.paymentTime" class="cell-row">
              <text class="cell-label">支付时间</text>
              <text class="cell-value">{{ formatDateTime(orderInfo.paymentTime) }}</text>
            </view>
            <view v-if="orderInfo.expiredTime" class="cell-row">
              <text class="cell-label">过期时间</text>
              <text class="cell-value">{{ formatDateTime(orderInfo.expiredTime) }}</text>
            </view>
          </view>
        </view>
      </view>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.order-detail-page {
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
  gap: 18rpx;
}

.empty-text {
  font-size: 26rpx;
  color: var(--fg-text-weak);
}

.detail {
  display: flex;
  flex-direction: column;
  gap: 18rpx;
}

.summary-card {
  padding: 22rpx 22rpx 20rpx;
  border-radius: 28rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border-weak);
  box-shadow: var(--fg-shadow-card);
}

.summary-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16rpx;
  padding-bottom: 16rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.summary-status {
  padding: 8rpx 18rpx;
  border-radius: 999rpx;
  background: rgba(var(--fg-primary-rgb), 0.12);
  color: var(--fg-primary-600);
  font-size: 24rpx;
  font-weight: 600;
  flex-shrink: 0;
}

.summary-status.status-pendingPayment,
.summary-status.status-pendingDelivery,
.summary-status.status-pendingReceipt {
  background: rgba(255, 159, 10, 0.14);
  color: var(--fg-warning);
}

.summary-status.status-completed {
  background: rgba(52, 199, 89, 0.16);
  color: var(--fg-success);
}

.summary-status.status-canceled,
.summary-status.status-refunded {
  background: rgba(255, 59, 48, 0.14);
  color: var(--fg-danger);
}

.summary-amount {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 6rpx;
}

.amount-label {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.amount-value {
  font-size: 34rpx;
  font-weight: 700;
  color: var(--fg-text);
}

.summary-meta {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
  padding: 16rpx;
  margin-top: 16rpx;
  border-radius: 20rpx;
  background: var(--fg-surface-glass);
  border: 1px solid var(--fg-border-weak);
}

.meta-item {
  display: flex;
  justify-content: space-between;
  gap: 12rpx;
  align-items: center;
}

.meta-label {
  font-size: 22rpx;
  color: var(--fg-text-muted);
  flex-shrink: 0;
}

.meta-value {
  font-size: 24rpx;
  color: var(--fg-text);
  font-weight: 600;
  text-align: right;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.section-card {
  padding: 18rpx 20rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border-weak);
  box-shadow: var(--fg-shadow-card);
}

.section-title {
  font-size: 26rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 8rpx;
}

.cell-list {
  display: flex;
  flex-direction: column;
}

.cell-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16rpx;
  padding: 14rpx 0;
}

.cell-row + .cell-row {
  border-top: 1px solid var(--fg-border-weak);
}

.cell-row.multiline {
  align-items: flex-start;
}

.cell-label {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  flex-shrink: 0;
}

.cell-value {
  font-size: 24rpx;
  color: var(--fg-text);
  text-align: right;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.cell-value.wrap {
  white-space: normal;
  text-align: left;
}

.cell-value.strong {
  font-weight: 700;
  color: var(--fg-primary-600);
}

.cell-value.price {
  font-weight: 700;
  font-size: 28rpx;
}
</style>
