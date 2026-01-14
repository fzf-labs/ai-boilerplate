<script lang="ts" setup>
import type { MallProductInfo } from '@/api/v1/mall-product/types'
import type { GetUserMembershipInfoReply, MembershipBenefit } from '@/api/v1/membership/types'
import { useToast } from 'wot-design-uni'
import { getMallProductList } from '@/api/v1/mall-product/mallProduct'
import { getMembershipBenefits, getUserMembershipInfo } from '@/api/v1/membership/membership'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: 'VIP会员中心',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

// 当前会员信息
const membershipInfo = ref<GetUserMembershipInfoReply | null>(null)

// 会员类型主题配置 - 以绿色为主色调（与 me.vue 保持一致）
const membershipTheme = computed(() => {
  const type = membershipInfo.value?.membershipType || 'normal'
  switch (type) {
    case 'svip':
      // 深翠绿 + 金色点缀 - 尊贵感
      return {
        icon: '👑',
        gradient: 'linear-gradient(135deg, #065f46 0%, #047857 50%, #059669 100%)',
        shadow: 'rgba(6, 95, 70, 0.35)',
        shadowAlt: 'rgba(4, 120, 87, 0.25)',
      }
    case 'vip':
      // 翠绿色 - 经典VIP
      return {
        icon: '💎',
        gradient: 'linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%)',
        shadow: 'rgba(16, 185, 129, 0.3)',
        shadowAlt: 'rgba(5, 150, 105, 0.2)',
      }
    default:
      // 浅绿灰色 - 普通会员
      return {
        icon: '⭐',
        gradient: 'linear-gradient(135deg, #6ee7b7 0%, #34d399 50%, #10b981 100%)',
        shadow: 'rgba(110, 231, 183, 0.3)',
        shadowAlt: 'rgba(52, 211, 153, 0.2)',
      }
  }
})

// 会员描述文案
const membershipDescText = computed(() => {
  const info = membershipInfo.value
  if (!info)
    return '升级VIP解锁更多权益'

  const type = info.membershipType || 'normal'

  if (type === 'normal') {
    return info.membershipDescription || '升级VIP解锁更多权益'
  }

  // VIP/SVIP 会员显示到期时间
  if (info.expiredAt) {
    const expireDate = info.expiredAt.split('T')[0]
    return `有效期至 ${expireDate}`
  }

  return info.membershipDescription || '尊享专属会员权益'
})

// VIP 套餐列表
const productList = ref<MallProductInfo[]>([])
// 当前选中的套餐
const selectedProduct = ref<MallProductInfo | null>(null)
// 会员权益列表
const benefitList = ref<MembershipBenefit[]>([])
// 加载状态
const loading = ref(false)

/**
 * 获取用户会员信息
 */
async function fetchMembershipInfo() {
  if (!tokenStore.hasLogin)
    return

  try {
    const res = await getUserMembershipInfo({ options: {} })
    membershipInfo.value = res
  }
  catch (error) {
    console.error('获取会员信息失败:', error)
  }
}

/**
 * 获取 VIP 套餐列表
 */
async function fetchProductList() {
  try {
    loading.value = true
    const res = await getMallProductList({
      params: {
        page: 1,
        pageSize: 20,
        productType: 'membership',
        status: 1, // 在售
      },
      options: {},
    })
    productList.value = res.list || []
    // 默认选中第一个套餐
    if (productList.value.length > 0 && !selectedProduct.value) {
      selectedProduct.value = productList.value[0]
    }
  }
  catch (error) {
    console.error('获取套餐列表失败:', error)
    toast.error('获取套餐列表失败')
  }
  finally {
    loading.value = false
  }
}

/**
 * 获取会员权益列表
 */
async function fetchBenefitList() {
  try {
    // 获取 VIP 权益（不传 membershipType 获取最高等级权益展示）
    const res = await getMembershipBenefits({
      params: { membershipType: 'vip' },
      options: {},
    })
    benefitList.value = res.benefits || []
  }
  catch (error) {
    console.error('获取权益列表失败:', error)
  }
}

/**
 * 选择套餐
 */
function handleSelectProduct(product: MallProductInfo) {
  selectedProduct.value = product
}

/**
 * 立即开通
 */
function handleSubscribe() {
  if (!tokenStore.hasLogin) {
    toast.warning('请先登录')
    setTimeout(() => {
      uni.navigateTo({ url: '/pages-fg/login/login' })
    }, 1500)
    return
  }

  if (!selectedProduct.value) {
    toast.warning('请选择套餐')
    return
  }

  // 跳转到支付页面
  uni.navigateTo({
    url: `/pages-fg/vip/pay?productId=${selectedProduct.value.id}&productName=${encodeURIComponent(selectedProduct.value.productName || '')}&price=${selectedProduct.value.currentPrice}`,
  })
}

/**
 * 格式化价格
 */
function formatPrice(price?: number) {
  if (price === undefined || price === null)
    return '0.00'
  return price.toFixed(2)
}

/**
 * 解析商品详情
 */
function parseProductDetail(detail?: string): { features?: string[], duration?: string } {
  if (!detail)
    return {}
  try {
    return JSON.parse(detail)
  }
  catch {
    return {}
  }
}

// 页面加载
onLoad(() => {
  fetchMembershipInfo()
  fetchProductList()
  fetchBenefitList()
})
</script>

<template>
  <view class="vip-page">
    <view class="top-bg" />

    <!-- 当前会员状态卡片 - 与 me.vue 保持一致的样式 -->
    <view class="status-section">
      <view
        class="status-card"
        :style="{
          background: membershipTheme.gradient,
          boxShadow: `0 20rpx 60rpx ${membershipTheme.shadow}, 0 8rpx 16rpx ${membershipTheme.shadowAlt}`,
        }"
      >
        <!-- 背景装饰层 -->
        <view class="status-decoration">
          <view class="decoration-circle c1" />
          <view class="decoration-circle c2" />
          <view class="decoration-circle c3" />
        </view>

        <view class="status-content">
          <view class="status-icon">
            <text class="icon-emoji">{{ membershipTheme.icon }}</text>
          </view>
          <view class="status-info">
            <view class="status-title-row">
              <view class="status-name">
                {{ membershipInfo?.membershipName || '普通会员' }}
              </view>
              <view class="status-type">
                {{ (membershipInfo?.membershipType || 'normal').toUpperCase() }}
              </view>
            </view>
            <view class="status-desc">
              {{ membershipDescText }}
            </view>
          </view>
        </view>
      </view>
    </view>

    <!-- 套餐选择 -->
    <view class="section">
      <view class="section-title">
        <text class="title-text">选择套餐</text>
      </view>

      <view v-if="loading" class="loading-box">
        <wd-loading type="ring" />
      </view>

      <view v-else class="product-grid">
        <view
          v-for="product in productList"
          :key="product.id"
          class="product-card"
          :class="{ active: selectedProduct?.id === product.id }"
          @click="handleSelectProduct(product)"
        >
          <!-- 推荐标签 -->
          <view v-if="product.sort === 1" class="recommend-tag">
            推荐
          </view>

          <view class="product-name">
            {{ product.productName }}
          </view>
          <view class="product-price">
            <text class="price-symbol">¥</text>
            <text class="price-value">{{ formatPrice(product.currentPrice) }}</text>
          </view>
          <view v-if="product.originalPrice && product.originalPrice > (product.currentPrice || 0)" class="original-price">
            ¥{{ formatPrice(product.originalPrice) }}
          </view>
          <view class="product-desc">
            {{ product.productDesc }}
          </view>

          <!-- 选中标记 -->
          <view v-if="selectedProduct?.id === product.id" class="check-mark">
            <wd-icon name="check" size="28rpx" color="#fff" />
          </view>
        </view>
      </view>
    </view>

    <!-- VIP 权益展示 -->
    <view class="section">
      <view class="section-title">
        <text class="title-text">VIP专属权益</text>
      </view>

      <view class="benefit-grid">
        <view v-for="benefit in benefitList" :key="benefit.benefitKey" class="benefit-item">
          <view class="benefit-icon">
            <wd-icon name="check-circle" size="40rpx" color="#10b981" />
          </view>
          <view class="benefit-content">
            <view class="benefit-name">
              {{ benefit.benefitName }}
            </view>
            <view v-if="benefit.benefitDesc" class="benefit-desc">
              {{ benefit.benefitDesc }}
            </view>
          </view>
          <view v-if="benefit.benefitNum" class="benefit-value">
            {{ benefit.benefitNum }}次
          </view>
        </view>
      </view>
    </view>

    <!-- 底部操作栏 -->
    <view class="bottom-bar">
      <view class="bar-content">
        <view class="price-info">
          <text class="label">应付金额</text>
          <view class="total-price">
            <text class="symbol">¥</text>
            <text class="amount">{{ formatPrice(selectedProduct?.currentPrice) }}</text>
          </view>
        </view>
        <wd-button
          type="primary"
          size="large"
          round
          :disabled="!selectedProduct"
          custom-class="subscribe-btn"
          @click="handleSubscribe"
        >
          立即开通
        </wd-button>
      </view>
    </view>

    <!-- 底部安全距离 -->
    <view class="safe-bottom" />

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.vip-page {
  min-height: 100vh;
  background: var(--fg-bg);
  padding-bottom: 200rpx;
}

.top-bg {
  position: fixed;
  left: 0;
  top: 0;
  right: 0;
  height: 400rpx;
  background: linear-gradient(180deg, #10b981 0%, var(--fg-bg) 100%);
  z-index: 0;
}

/* 会员状态卡片 - 与 me.vue 保持一致 */
.status-section {
  position: relative;
  padding: 32rpx var(--fg-page-x);
  z-index: 1;
}

.status-card {
  position: relative;
  border-radius: 32rpx;
  overflow: hidden;
  border: none;
  transition: all 0.3s ease-out;
}

.status-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0) 100%);
  pointer-events: none;
}

.status-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
  z-index: 0;
}

.decoration-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(40rpx);
}

.c1 {
  width: 200rpx;
  height: 200rpx;
  top: -80rpx;
  right: -60rpx;
  background: rgba(255, 255, 255, 0.15);
}

.c2 {
  width: 150rpx;
  height: 150rpx;
  bottom: -40rpx;
  left: -40rpx;
  background: rgba(255, 255, 255, 0.1);
}

.c3 {
  width: 100rpx;
  height: 100rpx;
  top: 50%;
  right: 20rpx;
  background: rgba(255, 255, 255, 0.08);
}

.status-content {
  position: relative;
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 40rpx 32rpx;
  z-index: 2;
}

.status-icon {
  flex-shrink: 0;
  width: 88rpx;
  height: 88rpx;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: blur(20rpx);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8rpx 24rpx rgba(0, 0, 0, 0.15);
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.icon-emoji {
  font-size: 44rpx;
}

.status-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10rpx;
  min-width: 0;
}

.status-title-row {
  display: flex;
  align-items: center;
  gap: 12rpx;
  flex-wrap: wrap;
}

.status-name {
  font-size: 40rpx;
  font-weight: 800;
  color: #ffffff;
  line-height: 1.2;
  letter-spacing: 0.5rpx;
  text-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.15);
}

.status-type {
  font-size: 24rpx;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.85);
  letter-spacing: 2rpx;
  text-transform: uppercase;
}

.status-desc {
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.6;
  font-weight: 400;
  letter-spacing: 0.3rpx;
}

/* 通用 section */
.section {
  position: relative;
  padding: 0 var(--fg-page-x);
  margin-top: 32rpx;
  z-index: 1;
}

.section-title {
  display: flex;
  align-items: center;
  margin-bottom: 24rpx;
}

.title-text {
  font-size: 32rpx;
  font-weight: 700;
  color: var(--fg-text);
}

/* 套餐选择 */
.loading-box {
  display: flex;
  justify-content: center;
  padding: 60rpx 0;
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16rpx;
}

.product-card {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 28rpx 16rpx;
  background: var(--fg-surface);
  border-radius: 20rpx;
  border: 2px solid var(--fg-border);
  transition: all 0.3s ease;
}

.product-card.active {
  border-color: #10b981;
  background: linear-gradient(180deg, rgba(16, 185, 129, 0.08) 0%, var(--fg-surface) 100%);
}

.recommend-tag {
  position: absolute;
  top: -2px;
  right: -2px;
  padding: 4rpx 16rpx;
  font-size: 20rpx;
  font-weight: 600;
  color: #fff;
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  border-radius: 0 18rpx 0 12rpx;
}

.product-name {
  font-size: 28rpx;
  font-weight: 600;
  color: var(--fg-text);
  margin-bottom: 12rpx;
}

.product-price {
  display: flex;
  align-items: baseline;
  color: #10b981;
  margin-bottom: 4rpx;
}

.price-symbol {
  font-size: 24rpx;
  font-weight: 600;
}

.price-value {
  font-size: 40rpx;
  font-weight: 800;
}

.original-price {
  font-size: 22rpx;
  color: var(--fg-text-muted);
  text-decoration: line-through;
  margin-bottom: 8rpx;
}

.product-desc {
  font-size: 22rpx;
  color: var(--fg-text-muted);
  text-align: center;
}

.check-mark {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 40rpx;
  height: 40rpx;
  background: #10b981;
  border-radius: 12rpx 0 18rpx 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 权益列表 */
.benefit-grid {
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
  overflow: hidden;
}

.benefit-item {
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 28rpx 24rpx;
  border-bottom: 1px solid var(--fg-border);
}

.benefit-item:last-child {
  border-bottom: none;
}

.benefit-icon {
  flex-shrink: 0;
}

.benefit-content {
  flex: 1;
  min-width: 0;
}

.benefit-name {
  font-size: 28rpx;
  font-weight: 600;
  color: var(--fg-text);
  margin-bottom: 4rpx;
}

.benefit-desc {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.benefit-value {
  flex-shrink: 0;
  font-size: 28rpx;
  font-weight: 700;
  color: #10b981;
}

/* 底部操作栏 */
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
  flex-direction: column;
  gap: 4rpx;
}

.label {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.total-price {
  display: flex;
  align-items: baseline;
  color: #ef4444;
}

.total-price .symbol {
  font-size: 26rpx;
  font-weight: 600;
}

.total-price .amount {
  font-size: 44rpx;
  font-weight: 800;
}

:deep(.subscribe-btn) {
  min-width: 280rpx;
}

.safe-bottom {
  height: env(safe-area-inset-bottom);
}
</style>
