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
          <view class="row head">
            <view class="left">
              <text class="title">订单号</text>
              <text class="value">{{ item.id }}</text>
            </view>
            <text class="status">{{ getOrderStatusText(item.status) }}</text>
          </view>

          <view class="row">
            <text class="label">类型</text>
            <text class="value">{{ getProductTypeText(item.productType) }}</text>
          </view>

          <view class="row">
            <text class="label">支付</text>
            <text class="value">{{ getPaymentStatusText(item.paymentStatus) }}</text>
          </view>

          <view class="row">
            <text class="label">金额</text>
            <text class="value price">¥{{ formatPrice(item.actualAmount) }}</text>
          </view>

          <view class="row">
            <text class="label">创建</text>
            <text class="value">{{ formatDateTime(item.createdAt) }}</text>
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
  margin-bottom: 18rpx;
}

.status-row {
  display: flex;
  gap: 12rpx;
  padding-right: 12rpx;
}

.status-pill {
  height: 56rpx;
  padding: 0 20rpx;
  border-radius: 28rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  flex-shrink: 0;
}

.status-pill.active {
  background: rgba(var(--fg-primary-rgb), 0.1);
  border-color: rgba(var(--fg-primary-rgb), 0.25);
}

.status-text {
  font-size: 24rpx;
  color: var(--fg-text);
}

.status-pill.active .status-text {
  color: var(--fg-primary-600);
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
  gap: 16rpx;
}

.order-item {
  padding: 20rpx 22rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16rpx;
  padding: 10rpx 0;
}

.row.head {
  padding-top: 0;
  padding-bottom: 14rpx;
  border-bottom: 1px solid var(--fg-border);
}

.left {
  display: flex;
  flex: 1;
  min-width: 0;
  gap: 10rpx;
  align-items: center;
}

.title {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  flex-shrink: 0;
}

.label {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.value {
  font-size: 24rpx;
  color: var(--fg-text);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: right;
}

.status {
  font-size: 24rpx;
  color: var(--fg-primary-600);
  flex-shrink: 0;
}

.price {
  font-weight: 700;
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
