<script lang="ts" setup>
import type { MallOrderInfo } from '@/api/v1/mall-order/types'
import { useToast } from 'wot-design-uni'
import { getUserOrderList } from '@/api/v1/mall-order/mallOrder'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '我的订单',
  },
  excludeLoginPath: true,
})

type OrderStatus = NonNullable<MallOrderInfo['status']>

const toast = useToast()
const tokenStore = useTokenStore()

const statusOptions: Array<{ label: string, value: '' | OrderStatus }> = [
  { label: '全部', value: '' },
  { label: '待付款', value: 'pendingPayment' },
  { label: '待发货', value: 'pendingDelivery' },
  { label: '待收货', value: 'pendingReceipt' },
  { label: '已完成', value: 'completed' },
  { label: '已取消', value: 'canceled' },
  { label: '已退款', value: 'refunded' },
]

const activeStatus = ref<'' | OrderStatus>('')

const items = ref<MallOrderInfo[]>([])
const page = ref(1)
const pageSize = 20
const total = ref(0)
const loading = ref(false)
const lastFetchCount = ref(0)

const hasMore = computed(() => {
  if (total.value > 0)
    return items.value.length < total.value
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

async function fetchList(reset = false) {
  if (loading.value)
    return
  if (!ensureLogin())
    return

  loading.value = true
  const currentPage = reset ? 1 : page.value
  try {
    const res = await getUserOrderList({
      params: {
        page: currentPage,
        pageSize,
        status: activeStatus.value || undefined,
      },
      options: {},
    })

    const list = res.list || []
    total.value = res.total ?? total.value
    lastFetchCount.value = list.length

    if (reset) {
      items.value = list
    }
    else {
      items.value = items.value.concat(list)
    }
    page.value = currentPage + 1
  }
  catch (error) {
    console.error('加载订单列表失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

function handleStatusChange(status: '' | OrderStatus) {
  if (activeStatus.value === status)
    return
  activeStatus.value = status
  fetchList(true)
}

function goDetail(item: MallOrderInfo) {
  if (!item.id)
    return
  uni.navigateTo({
    url: `/pages/orders/detail?orderId=${encodeURIComponent(item.id)}`,
  })
}

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
  <view class="order-list-page">
    <view class="top-bg" />

    <view class="content">
      <scroll-view class="status-bar" scroll-x :show-scrollbar="false">
        <view class="status-row">
          <view
            v-for="s in statusOptions"
            :key="s.value"
            class="status-pill"
            :class="{ active: activeStatus === s.value }"
            @click="handleStatusChange(s.value)"
          >
            <text class="status-text">{{ s.label }}</text>
          </view>
        </view>
      </scroll-view>

      <view v-if="loading && items.length === 0" class="loading-box">
        <wd-loading />
      </view>

      <view v-else-if="items.length === 0" class="empty-box">
        <wd-icon name="inbox" size="120rpx" color="var(--fg-text-disabled)" />
        <text class="empty-text">暂无订单</text>
      </view>

      <view v-else class="order-list">
        <view
          v-for="item in items"
          :key="item.id"
          class="order-item"
          @click="goDetail(item)"
        >
          <view class="order-head">
            <view class="order-id">
              <text class="order-id-label">订单号</text>
              <text class="order-id-value">{{ item.id }}</text>
            </view>
            <view class="order-status" :class="`status-${item.status || 'unknown'}`">
              <text class="order-status-text">{{ getOrderStatusText(item.status) }}</text>
            </view>
          </view>

          <view class="order-body">
            <view class="meta-row">
              <view class="meta-cell">
                <text class="meta-label">类型</text>
                <text class="meta-value">{{ getProductTypeText(item.productType) }}</text>
              </view>
              <view class="meta-cell">
                <text class="meta-label">支付</text>
                <text class="meta-value">{{ getPaymentStatusText(item.paymentStatus) }}</text>
              </view>
            </view>
            <view class="meta-row">
              <view class="meta-cell">
                <text class="meta-label">金额</text>
                <text class="meta-value price">¥{{ formatPrice(item.actualAmount) }}</text>
              </view>
              <view class="meta-cell">
                <text class="meta-label">创建</text>
                <text class="meta-value">{{ formatDateTime(item.createdAt) }}</text>
              </view>
            </view>
          </view>
        </view>
      </view>

      <view v-if="items.length > 0" class="list-footer">
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
.order-list-page {
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
  padding: 18rpx var(--fg-page-x) 40rpx;
}

.status-bar {
  margin: 6rpx 0 20rpx;
  padding: 8rpx;
  border-radius: 999rpx;
  background: var(--fg-surface-glass);
  border: 1px solid var(--fg-border-weak);
  box-shadow: var(--fg-shadow-soft);
  box-sizing: border-box;
}

.status-row {
  display: inline-flex;
  align-items: center;
  gap: 8rpx;
  padding-right: 6rpx;
  flex-wrap: nowrap;
}

.status-pill {
  height: 52rpx;
  padding: 0 18rpx;
  border-radius: 999rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: 1px solid transparent;
  flex-shrink: 0;
  transition: all 0.2s ease;
}

.status-pill.active {
  background: var(--fg-surface);
  border-color: var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
}

.status-text {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.status-pill.active .status-text {
  color: var(--fg-text);
  font-weight: 600;
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

.order-list {
  display: flex;
  flex-direction: column;
  gap: 18rpx;
}

.order-item {
  padding: 22rpx 22rpx 18rpx;
  border-radius: 28rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border-weak);
  box-shadow: var(--fg-shadow-card);
  transition:
    transform 0.18s ease,
    box-shadow 0.18s ease,
    border-color 0.18s ease;
}

.order-item:active {
  transform: scale(0.99);
  box-shadow: var(--fg-shadow-soft);
  border-color: rgba(var(--fg-primary-rgb), 0.2);
}

.order-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 14rpx;
  padding-bottom: 14rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.order-id {
  display: flex;
  align-items: center;
  gap: 10rpx;
  min-width: 0;
}

.order-id-label {
  font-size: 22rpx;
  color: var(--fg-text-muted);
  flex-shrink: 0;
}

.order-id-value {
  font-size: 24rpx;
  font-weight: 600;
  color: var(--fg-text);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.order-status {
  padding: 6rpx 14rpx;
  border-radius: 999rpx;
  font-size: 22rpx;
  font-weight: 600;
  background: rgba(var(--fg-primary-rgb), 0.12);
  color: var(--fg-primary-600);
  flex-shrink: 0;
}

.order-status.status-pendingPayment,
.order-status.status-pendingDelivery,
.order-status.status-pendingReceipt {
  background: rgba(255, 159, 10, 0.14);
  color: var(--fg-warning);
}

.order-status.status-completed {
  background: rgba(52, 199, 89, 0.16);
  color: var(--fg-success);
}

.order-status.status-canceled,
.order-status.status-refunded {
  background: rgba(255, 59, 48, 0.14);
  color: var(--fg-danger);
}

.order-body {
  padding-top: 14rpx;
  display: flex;
  flex-direction: column;
  gap: 12rpx;
}

.meta-row {
  display: flex;
  gap: 12rpx;
}

.meta-cell {
  flex: 1;
  min-width: 0;
  padding: 12rpx 14rpx;
  border-radius: 18rpx;
  background: var(--fg-surface-glass);
  border: 1px solid var(--fg-border-weak);
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.meta-label {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.meta-value {
  font-size: 24rpx;
  color: var(--fg-text);
  font-weight: 600;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.meta-value.price {
  font-size: 28rpx;
  color: var(--fg-text);
}

.list-footer {
  padding: 24rpx 0 0;
  display: flex;
  justify-content: center;
}

.loading-more {
  display: flex;
  align-items: center;
  gap: 12rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.no-more {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}
</style>
