<script lang="ts" setup>
import type { GetPaymentInfoReply, PlaceOrderReply } from '@/api/v1/mall-order/types'
import { useToast } from 'wot-design-uni'
import { getPaymentInfo, placeOrder } from '@/api/v1/mall-order/mallOrder'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '确认订单',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

// 路由参数
const productId = ref('')
const productName = ref('')
const price = ref(0)

// 订单信息
const orderInfo = ref<PlaceOrderReply | null>(null)
// 支付信息
const paymentInfo = ref<GetPaymentInfoReply | null>(null)
// 选中的支付方式
const paymentMethod = ref('wechat')
// 加载状态
const loading = ref(false)
const paying = ref(false)

// 支付方式列表
const paymentMethods = [
  {
    value: 'wechat',
    label: '微信支付',
    icon: '/static/images/wechat-pay.svg',
    desc: '推荐使用',
  },
]

/**
 * 格式化价格
 */
function formatPrice(p?: number) {
  if (p === undefined || p === null)
    return '0.00'
  return p.toFixed(2)
}

/**
 * 创建订单
 */
async function handleCreateOrder() {
  if (!tokenStore.hasLogin) {
    toast.warning('请先登录')
    return
  }

  if (!productId.value) {
    toast.error('商品信息错误')
    return
  }

  try {
    loading.value = true
    const res = await placeOrder({
      body: {
        productType: 'membership',
        productId: productId.value,
        paymentMethod: paymentMethod.value,
      },
      options: {},
    })
    orderInfo.value = res
    // 订单创建成功，获取支付信息
    if (res.orderId) {
      await fetchPaymentInfo(res.orderId)
    }
  }
  catch (error) {
    console.error('创建订单失败:', error)
    toast.error('创建订单失败')
  }
  finally {
    loading.value = false
  }
}

/**
 * 获取支付信息
 */
async function fetchPaymentInfo(orderId: string) {
  try {
    const res = await getPaymentInfo({
      body: {
        orderId,
        paymentMethod: paymentMethod.value,
      },
      options: {},
    })
    paymentInfo.value = res
  }
  catch (error) {
    console.error('获取支付信息失败:', error)
    toast.error('获取支付信息失败')
  }
}

/**
 * 发起支付
 */
async function handlePay() {
  if (!orderInfo.value?.orderId) {
    // 先创建订单
    await handleCreateOrder()
    if (!orderInfo.value?.orderId)
      return
  }

  paying.value = true

  try {
    // 调起微信支付
    // #ifdef MP-WEIXIN
    if (paymentInfo.value) {
      await new Promise((resolve, reject) => {
        uni.requestPayment({
          provider: 'wxpay',
          timeStamp: paymentInfo.value!.timeStamp || '',
          nonceStr: paymentInfo.value!.nonceStr || '',
          package: `prepay_id=${paymentInfo.value!.prepayId}`,
          signType: (paymentInfo.value!.signType as 'MD5' | 'HMAC-SHA256' | 'RSA') || 'RSA',
          paySign: paymentInfo.value!.paySign || '',
          success: () => resolve(true),
          fail: (err) => {
            if (err.errMsg?.includes('cancel')) {
              reject(new Error('cancel'))
            }
            else {
              reject(err)
            }
          },
        })
      })

      // 支付成功
      uni.redirectTo({
        url: `/pages/vip/result?status=success&orderId=${orderInfo.value.orderId}`,
      })
    }
    // #endif

    // #ifndef MP-WEIXIN
    // H5/APP 使用支付链接
    if (paymentInfo.value?.paymentUrl) {
      // 打开支付页面
      uni.navigateTo({
        url: `/pages-fg/webview/index?url=${encodeURIComponent(paymentInfo.value.paymentUrl)}`,
      })
    }
    else {
      // 模拟支付成功（开发测试用）
      toast.loading('正在处理...')
      setTimeout(() => {
        toast.close()
        uni.redirectTo({
          url: `/pages/vip/result?status=success&orderId=${orderInfo.value?.orderId}`,
        })
      }, 1500)
    }
    // #endif
  }
  catch (error: any) {
    if (error?.message === 'cancel') {
      toast.warning('已取消支付')
    }
    else {
      console.error('支付失败:', error)
      uni.redirectTo({
        url: `/pages/vip/result?status=fail&orderId=${orderInfo.value?.orderId}`,
      })
    }
  }
  finally {
    paying.value = false
  }
}

// 页面加载
onLoad((options) => {
  const opts = options as Record<string, string>
  productId.value = opts.productId || ''
  productName.value = decodeURIComponent(opts.productName || '')
  price.value = Number.parseFloat(opts.price || '0')

  if (!productId.value) {
    toast.error('商品信息错误')
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  }
})
</script>

<template>
  <view class="pay-page">
    <!-- 商品信息 -->
    <view class="section">
      <view class="section-title">
        商品信息
      </view>
      <view class="product-card">
        <view class="product-icon">
          <text class="icon-emoji">💎</text>
        </view>
        <view class="product-info">
          <view class="product-name">
            {{ productName }}
          </view>
          <view class="product-type">
            VIP会员
          </view>
        </view>
        <view class="product-price">
          <text class="symbol">¥</text>
          <text class="amount">{{ formatPrice(price) }}</text>
        </view>
      </view>
    </view>

    <!-- 支付方式 -->
    <view class="section">
      <view class="section-title">
        支付方式
      </view>
      <view class="payment-list">
        <view
          v-for="method in paymentMethods"
          :key="method.value"
          class="payment-item"
          :class="{ active: paymentMethod === method.value }"
          @click="paymentMethod = method.value"
        >
          <view class="payment-icon">
            <image
              v-if="method.icon"
              :src="method.icon"
              class="icon-image"
              mode="aspectFit"
            />
            <wd-icon v-else name="wallet" size="48rpx" color="var(--fg-primary-600)" />
          </view>
          <view class="payment-info">
            <view class="payment-label">
              {{ method.label }}
            </view>
            <view v-if="method.desc" class="payment-desc">
              {{ method.desc }}
            </view>
          </view>
          <view class="payment-check">
            <wd-icon
              :name="paymentMethod === method.value ? 'check-circle-filled' : 'circle'"
              size="44rpx"
              :color="paymentMethod === method.value ? 'var(--fg-primary-600)' : 'var(--fg-text-weak)'"
            />
          </view>
        </view>
      </view>
    </view>

    <!-- 订单金额 -->
    <view class="section">
      <view class="amount-card">
        <view class="amount-row">
          <text class="label">商品金额</text>
          <text class="value">¥{{ formatPrice(price) }}</text>
        </view>
        <view class="amount-row">
          <text class="label">优惠金额</text>
          <text class="value discount">-¥0.00</text>
        </view>
        <view class="divider" />
        <view class="amount-row total">
          <text class="label">应付金额</text>
          <view class="total-value">
            <text class="symbol">¥</text>
            <text class="amount">{{ formatPrice(price) }}</text>
          </view>
        </view>
      </view>
    </view>

    <!-- 底部按钮 -->
    <view class="bottom-bar">
      <view class="bar-content">
        <view class="price-info">
          <text class="label">应付</text>
          <view class="total-price">
            <text class="symbol">¥</text>
            <text class="amount">{{ formatPrice(price) }}</text>
          </view>
        </view>
        <wd-button
          type="primary"
          size="large"
          round
          :loading="loading || paying"
          :disabled="!productId"
          custom-class="pay-btn"
          @click="handlePay"
        >
          {{ loading ? '创建订单中...' : paying ? '支付中...' : '立即支付' }}
        </wd-button>
      </view>
    </view>

    <!-- 底部安全距离 -->
    <view class="safe-bottom" />

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.pay-page {
  min-height: 100vh;
  background: var(--fg-bg);
  padding-bottom: 180rpx;
}

.section {
  padding: 24rpx var(--fg-page-x);
}

.section-title {
  font-size: 30rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 20rpx;
}

/* 商品卡片 */
.product-card {
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 28rpx 24rpx;
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
}

.product-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 16rpx;
  background: linear-gradient(135deg, var(--fg-primary-600) 0%, var(--fg-primary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-emoji {
  font-size: 40rpx;
}

.product-info {
  flex: 1;
  min-width: 0;
}

.product-name {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
  margin-bottom: 6rpx;
}

.product-type {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.product-price {
  display: flex;
  align-items: baseline;
  color: var(--fg-danger);
}

.product-price .symbol {
  font-size: 24rpx;
  font-weight: 600;
}

.product-price .amount {
  font-size: 36rpx;
  font-weight: 800;
}

/* 支付方式列表 */
.payment-list {
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
  overflow: hidden;
}

.payment-item {
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 28rpx 24rpx;
  border-bottom: 1px solid var(--fg-border);
  transition: background-color 0.2s ease;
}

.payment-item:last-child {
  border-bottom: none;
}

.payment-item:active {
  background: var(--fg-bg-alt);
}

.payment-item.active {
  background: rgba(var(--fg-primary-rgb), 0.06);
}

.payment-icon {
  width: 48rpx;
  height: 48rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-image {
  width: 48rpx;
  height: 48rpx;
}

.payment-info {
  flex: 1;
}

.payment-label {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.payment-desc {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  margin-top: 4rpx;
}

.payment-check {
  flex-shrink: 0;
}

/* 金额卡片 */
.amount-card {
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
  padding: 24rpx;
}

.amount-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12rpx 0;
}

.amount-row .label {
  font-size: 28rpx;
  color: var(--fg-text-secondary);
}

.amount-row .value {
  font-size: 28rpx;
  font-weight: 500;
  color: var(--fg-text);
}

.amount-row .discount {
  color: var(--fg-primary-600);
}

.amount-row.total .label {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.divider {
  height: 1px;
  background: var(--fg-border);
  margin: 12rpx 0;
}

.total-value {
  display: flex;
  align-items: baseline;
  color: var(--fg-danger);
}

.total-value .symbol {
  font-size: 26rpx;
  font-weight: 600;
}

.total-value .amount {
  font-size: 40rpx;
  font-weight: 800;
}

/* 底部栏 */
.bottom-bar {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  background: var(--fg-surface);
  border-top: 1px solid var(--fg-border);
  padding: 20rpx var(--fg-page-x);
  padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
  z-index: 100;
}

.bar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.price-info {
  display: flex;
  align-items: baseline;
  gap: 8rpx;
}

.price-info .label {
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

.total-price {
  display: flex;
  align-items: baseline;
  color: var(--fg-danger);
}

.total-price .symbol {
  font-size: 26rpx;
  font-weight: 600;
}

.total-price .amount {
  font-size: 44rpx;
  font-weight: 800;
}

:deep(.pay-btn) {
  min-width: 260rpx;
}

.safe-bottom {
  height: env(safe-area-inset-bottom);
}
</style>
