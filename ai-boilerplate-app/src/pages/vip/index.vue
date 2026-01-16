<script lang="ts" setup>
import type { MallProductInfo } from '@/api/v1/mall-product/types'
import type { GetUserMembershipInfoReply, MembershipBenefitCompareItem, MembershipBenefitValue } from '@/api/v1/membership/types'
import { useToast } from 'wot-design-uni'
import { getMembershipProductList } from '@/api/v1/mall-product/mallProduct'
import { getMembershipBenefitsCompare, getUserMembershipInfo } from '@/api/v1/membership/membership'
import { useTokenStore } from '@/store/token'

// 会员类型定义
type MembershipType = 'normal' | 'vip' | 'svip'

// 权益对比项（用于展示）
interface BenefitCompareDisplayItem {
  benefitKey: string
  benefitName: string
  benefitDesc?: string
  normal: string
  vip: string
  svip: string
}

// 会员等级配置
const MEMBERSHIP_CONFIG: Record<MembershipType, { name: string, icon: string, color: string, bgColor: string }> = {
  normal: { name: '普通会员', icon: 'N', color: 'var(--fg-gold-600)', bgColor: 'rgba(var(--fg-gold-rgb), 0.08)' },
  vip: { name: 'VIP会员', icon: 'V', color: 'var(--fg-gold-600)', bgColor: 'rgba(var(--fg-gold-rgb), 0.12)' },
  svip: { name: 'SVIP会员', icon: 'SV', color: 'var(--fg-gold-600)', bgColor: 'rgba(var(--fg-gold-rgb), 0.16)' },
}

definePage({
  style: {
    navigationBarTitleText: 'VIP会员中心',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

// 当前会员信息
const membershipInfo = ref<GetUserMembershipInfoReply | null>(null)

const isActiveMember = computed(() => {
  const info = membershipInfo.value
  if (!info)
    return false
  if (info.status !== 1)
    return false
  if (info.isExpired)
    return false
  return (info.membershipType || 'normal') !== 'normal'
})

const subscribeButtonText = computed(() => (isActiveMember.value ? '续费' : '立即开通'))

// 会员类型主题配置 - iOS 中性色调（与 me.vue 保持一致）
const membershipTheme = computed(() => {
  const type = membershipInfo.value?.membershipType || 'normal'
  switch (type) {
    case 'svip':
      return {
        icon: 'SV',
        gradient: 'linear-gradient(135deg, var(--fg-gold-100) 0%, var(--fg-gold-200) 60%, var(--fg-gold-300) 100%)',
        shadow: 'var(--fg-ink-08)',
        shadowAlt: 'var(--fg-ink-05)',
      }
    case 'vip':
      return {
        icon: 'V',
        gradient: 'linear-gradient(135deg, var(--fg-gold-50) 0%, var(--fg-gold-100) 55%, var(--fg-gold-200) 100%)',
        shadow: 'var(--fg-ink-08)',
        shadowAlt: 'var(--fg-ink-05)',
      }
    default:
      return {
        icon: 'N',
        gradient: 'linear-gradient(135deg, var(--fg-gold-50) 0%, var(--fg-gold-100) 60%, var(--fg-gold-200) 100%)',
        shadow: 'var(--fg-ink-08)',
        shadowAlt: 'var(--fg-ink-05)',
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
// 权益对比列表
const benefitCompareList = ref<BenefitCompareDisplayItem[]>([])
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
    const res = await getMembershipProductList({ options: {} })
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
 * 获取会员权益对比列表（单次请求）
 */
async function fetchBenefitCompare() {
  try {
    const res = await getMembershipBenefitsCompare({ options: {} })
    const items = res.items || []

    // 转换为展示用的数据结构
    benefitCompareList.value = items
      .map((item: MembershipBenefitCompareItem) => ({
        benefitKey: item.benefitKey || '',
        benefitName: item.benefitName || '',
        benefitDesc: item.benefitDesc,
        normal: formatBenefitValue(item.normal),
        vip: formatBenefitValue(item.vip),
        svip: formatBenefitValue(item.svip),
      }))
      .sort((a, b) => {
        const aItem = items.find(i => i.benefitKey === a.benefitKey)
        const bItem = items.find(i => i.benefitKey === b.benefitKey)
        return (aItem?.sort || 0) - (bItem?.sort || 0)
      })
  }
  catch (error) {
    console.error('获取权益对比列表失败:', error)
  }
}

/**
 * 格式化权益值显示
 */
function formatBenefitValue(benefit?: MembershipBenefitValue): string {
  if (!benefit || !benefit.supported) {
    return '✗' // 不支持
  }

  // 如果有次数，显示次数
  if (benefit.num) {
    const num = Number.parseInt(benefit.num, 10)
    if (num === -1 || benefit.num === 'unlimited') {
      return '无限'
    }
    return `${benefit.num}次`
  }

  // 如果有值，显示值
  if (benefit.value) {
    if (benefit.value === 'true' || benefit.value === '1') {
      return '✓'
    }
    if (benefit.value === 'false' || benefit.value === '0') {
      return '✗'
    }
    return benefit.value
  }

  // 默认表示支持
  return '✓'
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
    url: `/pages/vip/pay?productId=${selectedProduct.value.id}&productName=${encodeURIComponent(selectedProduct.value.productName || '')}&price=${selectedProduct.value.currentPrice}`,
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

// 页面加载
onLoad(() => {
  fetchMembershipInfo()
  fetchProductList()
  fetchBenefitCompare()
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
            <wd-icon name="check" size="28rpx" color="var(--fg-text-inverse)" />
          </view>
        </view>
      </view>
    </view>

    <!-- 会员权益对比 -->
    <view class="section">
      <view class="section-title">
        <text class="title-text">会员权益对比</text>
      </view>

      <view class="compare-table">
        <!-- 表头 - 会员等级 -->
        <view class="compare-header">
          <view class="compare-cell header-label">
            权益
          </view>
          <view
            v-for="type in (['normal', 'vip', 'svip'] as MembershipType[])"
            :key="type"
            class="compare-cell header-type"
            :class="[`type-${type}`, { 'is-current': membershipInfo?.membershipType === type }]"
          >
            <view class="type-icon">
              {{ MEMBERSHIP_CONFIG[type].icon }}
            </view>
            <view class="type-name">
              {{ MEMBERSHIP_CONFIG[type].name }}
            </view>
            <view v-if="membershipInfo?.membershipType === type" class="current-tag">
              当前
            </view>
          </view>
        </view>

        <!-- 权益行 -->
        <view
          v-for="item in benefitCompareList"
          :key="item.benefitKey"
          class="compare-row"
        >
          <view class="compare-cell row-label">
            <view class="label-name">
              {{ item.benefitName }}
            </view>
            <view v-if="item.benefitDesc" class="label-desc">
              {{ item.benefitDesc }}
            </view>
          </view>
          <view
            v-for="type in (['normal', 'vip', 'svip'] as MembershipType[])"
            :key="type"
            class="compare-cell row-value"
            :class="[`type-${type}`]"
          >
            <template v-if="item[type] === '✓'">
              <wd-icon name="check" size="36rpx" :color="MEMBERSHIP_CONFIG[type].color" />
            </template>
            <template v-else-if="item[type] === '✗'">
              <wd-icon name="close" size="36rpx" color="var(--fg-text-weak)" />
            </template>
            <template v-else>
              <text class="value-text" :style="{ color: MEMBERSHIP_CONFIG[type].color }">
                {{ item[type] }}
              </text>
            </template>
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
          {{ subscribeButtonText }}
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
  background: var(--fg-top-bg-gradient);
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
  border-radius: var(--fg-radius-card-lg);
  overflow: hidden;
  border: 1px solid var(--fg-border);
  transition: all 0.3s ease-out;
}

.status-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, var(--fg-glass-50) 0%, var(--fg-glass-0) 100%);
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
  background: var(--fg-ink-04);
  -webkit-backdrop-filter: blur(24rpx);
  backdrop-filter: blur(24rpx);
}

.c1 {
  width: 200rpx;
  height: 200rpx;
  top: -80rpx;
  right: -60rpx;
  background: var(--fg-ink-04);
}

.c2 {
  width: 150rpx;
  height: 150rpx;
  bottom: -40rpx;
  left: -40rpx;
  background: var(--fg-ink-03);
}

.c3 {
  width: 100rpx;
  height: 100rpx;
  top: 50%;
  right: 20rpx;
  background: var(--fg-ink-03);
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
  background: var(--fg-surface-glass);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: var(--fg-shadow-soft);
  border: 1px solid var(--fg-border);
}

.icon-emoji {
  font-size: 32rpx;
  font-weight: 700;
  letter-spacing: 1rpx;
  color: var(--fg-gold-600);
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
  color: var(--fg-text);
  line-height: 1.2;
  letter-spacing: 0.2rpx;
}

.status-type {
  font-size: 24rpx;
  font-weight: 600;
  color: var(--fg-text-muted);
  letter-spacing: 2rpx;
  text-transform: uppercase;
}

.status-desc {
  font-size: 28rpx;
  color: var(--fg-text-secondary);
  line-height: 1.6;
  font-weight: 400;
  letter-spacing: 0.2rpx;
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
  border-radius: var(--fg-radius-card);
  border: 2px solid var(--fg-border);
  transition: all 0.3s ease;
}

.product-card.active {
  border-color: var(--fg-gold-600);
  background: linear-gradient(180deg, rgba(var(--fg-gold-rgb), 0.12) 0%, var(--fg-surface) 100%);
}

.recommend-tag {
  position: absolute;
  top: -2px;
  right: -2px;
  padding: 4rpx 16rpx;
  font-size: 20rpx;
  font-weight: 600;
  color: var(--fg-text-inverse);
  background: var(--fg-gold-600);
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
  color: var(--fg-danger);
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
  background: var(--fg-gold-600);
  border-radius: 12rpx 0 18rpx 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 权益对比表格 */
.compare-table {
  background: var(--fg-surface);
  border-radius: var(--fg-radius-card);
  border: 1px solid var(--fg-border);
  overflow: hidden;
}

.compare-header {
  display: flex;
  background: linear-gradient(180deg, rgba(var(--fg-gold-rgb), 0.1) 0%, var(--fg-surface) 100%);
  border-bottom: 1px solid var(--fg-border);
}

.compare-row {
  display: flex;
  border-bottom: 1px solid var(--fg-border);
}

.compare-row:last-child {
  border-bottom: none;
}

.compare-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20rpx 12rpx;
}

.header-label {
  flex: 1.2;
  justify-content: flex-start;
  padding-left: 24rpx;
  font-size: 24rpx;
  font-weight: 600;
  color: var(--fg-text-muted);
}

.header-type {
  flex: 1;
  flex-direction: column;
  gap: 8rpx;
  padding: 24rpx 12rpx;
  position: relative;
}

.header-type.is-current {
  background: rgba(var(--fg-gold-rgb), 0.14);
}

.type-icon {
  font-size: 32rpx;
}

.type-name {
  font-size: 22rpx;
  font-weight: 600;
  text-align: center;
  line-height: 1.2;
  color: var(--fg-gold-600);
}

.type-normal .type-name {
  color: var(--fg-gold-600);
}

.type-vip .type-name {
  color: var(--fg-gold-600);
}

.type-svip .type-name {
  color: var(--fg-gold-600);
}

.current-tag {
  position: absolute;
  top: 4rpx;
  right: 4rpx;
  padding: 2rpx 10rpx;
  font-size: 18rpx;
  font-weight: 600;
  color: var(--fg-text-inverse);
  background: var(--fg-gold-600);
  border-radius: 0 0 0 12rpx;
}

.row-label {
  flex: 1.2;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  gap: 4rpx;
  padding: 20rpx 24rpx;
}

.label-name {
  font-size: 26rpx;
  font-weight: 500;
  color: var(--fg-text);
}

.label-desc {
  font-size: 20rpx;
  color: var(--fg-text-muted);
  line-height: 1.3;
}

.row-value {
  flex: 1;
  min-height: 80rpx;
}

.row-value.type-normal {
  background: rgba(var(--fg-gold-rgb), 0.06);
}

.row-value.type-vip {
  background: rgba(var(--fg-gold-rgb), 0.1);
}

.row-value.type-svip {
  background: rgba(var(--fg-gold-rgb), 0.08);
}

.value-text {
  font-size: 24rpx;
  font-weight: 600;
  text-align: center;
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

:deep(.subscribe-btn) {
  min-width: 280rpx;
}

:deep(.subscribe-btn.wd-button--primary) {
  background: var(--fg-gold-600);
  border-color: var(--fg-gold-600);
}

.safe-bottom {
  height: env(safe-area-inset-bottom);
}
</style>
