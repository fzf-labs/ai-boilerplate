<script lang="ts" setup>
import type { GetUserMembershipInfoReply, MembershipBenefit } from '@/api/v1/membership/types'
import { useToast } from 'wot-design-uni'
import { getMembershipBenefits, getUserMembershipInfo } from '@/api/v1/membership/membership'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '会员详情',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

// 会员信息
const membershipInfo = ref<GetUserMembershipInfoReply | null>(null)
// 会员权益列表
const membershipBenefits = ref<MembershipBenefit[]>([])
const loading = ref(false)

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

const expiredText = computed(() => {
  const info = membershipInfo.value
  if (!info?.expiredAt)
    return '永久有效'
  return formatDate(info.expiredAt)
})

const benefitCount = computed(() => membershipBenefits.value.length)

const autoRenewText = computed(() => {
  const info = membershipInfo.value
  if (!info || info.autoRenew !== 1)
    return '关'
  if (info.autoRenewDays)
    return `开（每${info.autoRenewDays}天）`
  return '开'
})

/**
 * 获取用户会员信息
 */
async function fetchMembershipInfo() {
  if (!tokenStore.hasLogin)
    return

  loading.value = true
  try {
    const res = await getUserMembershipInfo({ options: {} })
    membershipInfo.value = res
  }
  catch (error) {
    console.error('获取会员信息失败:', error)
    toast.error('获取会员信息失败')
  }
  finally {
    loading.value = false
  }
}

/**
 * 获取会员权益列表
 */
async function fetchMembershipBenefits() {
  if (!tokenStore.hasLogin)
    return

  try {
    const res = await getMembershipBenefits({ params: {}, options: {} })
    membershipBenefits.value = res.benefits || []
  }
  catch (error) {
    console.error('获取会员权益失败:', error)
  }
}

/**
 * 格式化日期
 */
function formatDate(dateStr?: string) {
  if (!dateStr)
    return '-'
  if (dateStr.includes('T'))
    return dateStr.split('T')[0]
  return dateStr.split(' ')[0]
}

/**
 * 页面加载时获取会员信息
 */
onLoad(() => {
  fetchMembershipInfo()
  fetchMembershipBenefits()
})
</script>

<template>
  <view class="membership-detail-page">
    <view class="top-bg" />

    <view class="content">
      <!-- 会员卡片 -->
      <view
        class="membership-card"
        :style="{
          background: membershipTheme.gradient,
          boxShadow: `0 20rpx 60rpx ${membershipTheme.shadow}, 0 8rpx 16rpx ${membershipTheme.shadowAlt}`,
        }"
      >
        <!-- 背景装饰 -->
        <view class="card-decoration">
          <view class="decoration-circle circle-1" />
          <view class="decoration-circle circle-2" />
          <view class="decoration-circle circle-3" />
        </view>

        <!-- 会员信息 -->
        <view v-if="membershipInfo" class="card-content">
          <view class="card-header">
            <view class="member-icon">
              <text class="icon-text">{{ membershipTheme.icon }}</text>
            </view>
            <view class="member-info">
              <text class="member-name">{{ membershipInfo.membershipName || '普通会员' }}</text>
              <text class="member-type">{{ membershipInfo.membershipType?.toUpperCase() || 'NORMAL' }}</text>
            </view>
          </view>
          <text class="member-desc">{{ membershipInfo.membershipDescription || '享受基础服务' }}</text>
          <view class="meta-grid">
            <view class="meta-item">
              <text class="meta-label">到期时间</text>
              <text class="meta-value">{{ expiredText }}</text>
            </view>
            <view class="meta-item">
              <text class="meta-label">权益数量</text>
              <text class="meta-value">{{ benefitCount || '-' }}</text>
            </view>
            <view class="meta-item">
              <text class="meta-label">自动续费</text>
              <text class="meta-value">{{ autoRenewText }}</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 会员权益 -->
      <view v-if="membershipBenefits.length > 0" class="detail-section">
        <view class="section-title">
          会员权益
        </view>

        <view class="benefits-grid">
          <view v-for="benefit in membershipBenefits" :key="benefit.benefitKey" class="benefit-card">
            <view class="benefit-header">
              <text class="benefit-name">{{ benefit.benefitName }}</text>
              <text v-if="benefit.benefitNum" class="benefit-num">{{ benefit.benefitNum }}次</text>
            </view>
            <text v-if="benefit.benefitDesc" class="benefit-desc">{{ benefit.benefitDesc }}</text>
          </view>
        </view>
      </view>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.membership-detail-page {
  min-height: 100vh;
  background: var(--fg-bg);
  padding-bottom: 40rpx;
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

.content {
  position: relative;
  padding: 32rpx var(--fg-page-x) 40rpx;
  z-index: 1;
}

.membership-card {
  position: relative;
  border-radius: var(--fg-radius-card-lg);
  overflow: hidden;
  border: 1px solid var(--fg-border);
  margin-bottom: 24rpx;
}

.card-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}

.decoration-circle {
  position: absolute;
  border-radius: 50%;
  background: var(--fg-ink-04);
  -webkit-backdrop-filter: blur(24rpx);
  backdrop-filter: blur(24rpx);
}

.circle-1 {
  width: 200rpx;
  height: 200rpx;
  top: -80rpx;
  right: -60rpx;
  background: var(--fg-ink-04);
}

.circle-2 {
  width: 150rpx;
  height: 150rpx;
  bottom: -40rpx;
  left: -40rpx;
  background: var(--fg-ink-03);
}

.circle-3 {
  width: 100rpx;
  height: 100rpx;
  top: 50%;
  right: 20rpx;
  background: var(--fg-ink-03);
}

.card-content {
  position: relative;
  padding: 48rpx 40rpx;
  z-index: 1;
}

.card-header {
  display: flex;
  align-items: flex-start;
  gap: 24rpx;
  margin-bottom: 24rpx;
}

.member-icon {
  width: 96rpx;
  height: 96rpx;
  border-radius: 50%;
  background: var(--fg-surface-glass);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
}

.icon-text {
  font-size: 32rpx;
  font-weight: 700;
  letter-spacing: 1rpx;
  color: var(--fg-gold-600);
}

.member-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8rpx;
}

.member-name {
  font-size: 44rpx;
  font-weight: 800;
  color: var(--fg-text);
  line-height: 1.2;
  letter-spacing: 0.2rpx;
}

.member-type {
  font-size: 26rpx;
  font-weight: 600;
  color: var(--fg-text-muted);
  letter-spacing: 2rpx;
}

.member-desc {
  font-size: 28rpx;
  color: var(--fg-text-secondary);
  line-height: 1.6;
  font-weight: 400;
  letter-spacing: 0.2rpx;
}

.meta-grid {
  margin-top: 28rpx;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16rpx;
}

.meta-item {
  display: flex;
  flex-direction: column;
  gap: 6rpx;
  padding: 16rpx 18rpx;
  border-radius: 16rpx;
  background: var(--fg-glass-70);
  border: 1px solid var(--fg-border);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
}

.meta-label {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.meta-value {
  font-size: 24rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.detail-section {
  position: relative;
  margin-top: 28rpx;
  z-index: 1;
}

.section-title {
  font-size: 32rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 24rpx;
  padding-left: 8rpx;
}

/* 会员权益网格 */
.benefits-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16rpx;
}

.benefit-card {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
  padding: 28rpx 24rpx;
  border-radius: var(--fg-radius-card);
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  transition: all 0.3s ease;
}

.benefit-card:active {
  transform: translateY(-4rpx);
  box-shadow: 0 8rpx 24rpx var(--fg-ink-08);
}

.benefit-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 8rpx;
}

.benefit-name {
  flex: 1;
  font-size: 28rpx;
  font-weight: 600;
  color: var(--fg-text);
  letter-spacing: 0.3rpx;
}

.benefit-num {
  flex-shrink: 0;
  font-size: 32rpx;
  font-weight: 700;
  color: var(--fg-primary);
  line-height: 1.2;
}

.benefit-desc {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  line-height: 1.5;
}
</style>
