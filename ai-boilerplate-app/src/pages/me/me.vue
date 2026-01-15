<script lang="ts" setup>
import type { GetUserMembershipInfoReply } from '@/api/v1/membership/types'
import type { UserInfo } from '@/api/v1/user/types'
import { useToast } from 'wot-design-uni'
import { getUserMembershipInfo } from '@/api/v1/membership/membership'
import { getUserInfo } from '@/api/v1/user/user'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '我的',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

// 用户详细信息
const userProfile = ref<UserInfo | null>(null)

// 会员信息
const membershipInfo = ref<GetUserMembershipInfoReply | null>(null)

// 会员类型主题配置 - 清新薄荷为主色调
const membershipTheme = computed(() => {
  const type = membershipInfo.value?.membershipType || 'normal'
  switch (type) {
    case 'svip':
      // 深绿 + 清透薄荷 - 更清爽的尊贵感
      return {
        icon: '👑',
        gradient: 'linear-gradient(135deg, #064e3b 0%, #059669 55%, #6ee7b7 100%)',
        shadow: 'rgba(5, 150, 105, 0.32)',
        shadowAlt: 'rgba(110, 231, 183, 0.22)',
      }
    case 'vip':
      // 薄荷渐变 - 清新VIP
      return {
        icon: '💎',
        gradient: 'linear-gradient(135deg, #059669 0%, #10b981 50%, #a7f3d0 100%)',
        shadow: 'rgba(16, 185, 129, 0.3)',
        shadowAlt: 'rgba(167, 243, 208, 0.24)',
      }
    default:
      // 轻白 + 浅薄荷 - 普通会员
      return {
        icon: '⭐',
        gradient: 'linear-gradient(135deg, #ffffff 0%, #ecfdf5 55%, #d1fae5 100%)',
        shadow: 'rgba(15, 23, 42, 0.08)',
        shadowAlt: 'rgba(16, 185, 129, 0.14)',
      }
  }
})

// 会员按钮文案
const membershipActionText = computed(() => {
  const type = membershipInfo.value?.membershipType || 'normal'
  if (type === 'normal') {
    return '开通VIP'
  }
  if (type === 'svip') {
    return '续费'
  }
  return '续费/升级'
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

const displayName = computed(() => userProfile.value?.nickname || userProfile.value?.phone || '用户')
const displayAvatar = computed(() => userProfile.value?.avatar || '/static/images/default-avatar.png')
const displayPhone = computed(() => userProfile.value?.phone || '')

// 菜单列表
const menuList = [
  {
    title: '个人信息',
    icon: 'edit',
    label: '完善头像昵称',
    path: '/pages/profile/edit',
    needLogin: true,
  },
  {
    title: '账号管理',
    icon: 'user',
    label: '密码/手机号/退出登录',
    path: '/pages/security/index',
    needLogin: true,
  },
  {
    title: '帮助中心',
    icon: 'help',
    label: '常见问题与反馈',
    path: '/pages/help/index',
    needLogin: false,
  },
  {
    title: '通用设置',
    icon: 'setting',
    label: '语言/缓存/关于',
    path: '/pages/settings/index',
    needLogin: false,
  },
]

/**
 * 获取用户详细信息
 */
async function fetchUserProfile() {
  if (!tokenStore.hasLogin)
    return

  try {
    const res = await getUserInfo({ options: {} })
    userProfile.value = res.info || null
  }
  catch (error) {
    console.error('获取用户信息失败:', error)
  }
}

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
 * 登录
 */
async function handleLogin() {
  // #ifdef MP-WEIXIN
  await tokenStore.wxLogin()
  // #endif
  // #ifndef MP-WEIXIN
  uni.navigateTo({
    url: `${LOGIN_PAGE}`,
  })
  // #endif
}

/**
 * 菜单点击
 */
function handleMenuClick(item: typeof menuList[0]) {
  if (item.needLogin && !tokenStore.hasLogin) {
    toast.warning('请先登录')
    setTimeout(() => {
      handleLogin()
    }, 1500)
    return
  }

  uni.navigateTo({
    url: item.path,
  })
}

/**
 * 进入VIP中心
 */
function handleGoVipCenter() {
  uni.navigateTo({
    url: '/pages/vip/index',
  })
}

/**
 * 查看会员详情
 */
function handleViewMembershipDetail() {
  uni.navigateTo({
    url: '/pages/membership/detail',
  })
}

/**
 * 兑换激活码
 */
function handleRedeemCode() {
  if (!tokenStore.hasLogin) {
    toast.warning('请先登录')
    setTimeout(() => {
      handleLogin()
    }, 1500)
    return
  }

  uni.navigateTo({
    url: '/pages/vip/activation-code',
  })
}

/**
 * 页面展示时刷新用户信息（登录后返回会触发）
 */
onShow(() => {
  if (tokenStore.hasLogin) {
    fetchUserProfile()
    fetchMembershipInfo()
  }
  else {
    userProfile.value = null
    membershipInfo.value = null
  }
})
</script>

<template>
  <view class="me-page">
    <view class="top-bg" />
    <view class="hero">
      <view class="hero-card">
        <image
          :src="tokenStore.hasLogin ? displayAvatar : '/static/images/default-avatar.png'"
          class="avatar"
          mode="aspectFill"
        />
        <view class="hero-text">
          <view class="name-row">
            <text class="name">{{ tokenStore.hasLogin ? displayName : '未登录' }}</text>
          </view>
          <text class="sub">{{ tokenStore.hasLogin ? (displayPhone || '欢迎回来') : '登录后可同步个人信息' }}</text>
        </view>
        <wd-button v-if="!tokenStore.hasLogin" type="primary" size="small" round @click="handleLogin">
          登录/注册
        </wd-button>
      </view>
    </view>

    <view class="panel">
      <!-- 会员信息卡片 - Liquid Glass 风格 -->
      <view
        v-if="tokenStore.hasLogin"
        class="membership-card"
        :style="{
          background: membershipTheme.gradient,
          boxShadow: `0 20rpx 60rpx ${membershipTheme.shadow}, 0 8rpx 16rpx ${membershipTheme.shadowAlt}`,
        }"
        @click="handleViewMembershipDetail"
      >
        <!-- 背景装饰层 -->
        <view class="membership-bg-decoration">
          <view class="decoration-circle circle-1" />
          <view class="decoration-circle circle-2" />
          <view class="decoration-circle circle-3" />
        </view>

        <view class="membership-content">
          <!-- 顶部徽章区域 -->
          <view class="membership-header">
            <view class="membership-title-row">
              <view class="membership-icon-wrapper">
                <view class="membership-icon">
                  <text class="icon-text">{{ membershipTheme.icon }}</text>
                </view>
              </view>
              <view class="membership-title-content">
                <view class="title-main">
                  <text class="membership-name">{{ membershipInfo?.membershipName || '普通会员' }}</text>
                </view>
                <view class="title-sub">
                  <text class="membership-type-code">{{ (membershipInfo?.membershipType || 'normal').toUpperCase() }}</text>
                </view>
              </view>
              <!-- VIP 操作按钮 -->
              <view class="membership-action">
                <view class="action-btn" @click.stop="handleGoVipCenter">
                  <text class="action-text">{{ membershipActionText }}</text>
                </view>
                <view class="activation-btn" @click.stop="handleRedeemCode">
                  <text class="activation-text">激活码</text>
                </view>
              </view>
            </view>
            <text class="membership-desc">{{ membershipDescText }}</text>
          </view>
        </view>
      </view>

      <wd-card type="rectangle" custom-class="menu-card">
        <wd-cell-group>
          <wd-cell
            v-for="item in menuList"
            :key="item.title"
            :title="item.title"
            :label="item.label"
            :icon="item.icon"
            is-link
            @click="handleMenuClick(item)"
          />
        </wd-cell-group>
      </wd-card>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.me-page {
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

.hero {
  position: relative;
  padding: 22rpx var(--fg-page-x) 18rpx;
}

.hero-card {
  display: flex;
  align-items: center;
  gap: 18rpx;
  padding: 22rpx 18rpx 18rpx;
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.avatar {
  width: 96rpx;
  height: 96rpx;
  border-radius: 48rpx;
  border: 3rpx solid rgba(var(--fg-primary-rgb), 0.18);
  box-shadow: 0 10rpx 26rpx rgba(var(--fg-primary-rgb), 0.14);
}

.hero-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10rpx;
}

.name-row {
  display: flex;
  align-items: center;
  gap: 12rpx;
}

.name {
  font-size: 38rpx;
  font-weight: 700;
  color: var(--fg-text);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.sub {
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

.panel {
  position: relative;
  padding: 0 var(--fg-page-x) 40rpx;
}

/* 会员信息卡片样式 - Liquid Glass 风格 */
.membership-card {
  position: relative;
  border-radius: 32rpx;
  overflow: hidden;
  border: none;
  margin-bottom: 18rpx;
  padding: 40rpx 32rpx;
  transition: all 0.3s ease-out;
  cursor: pointer;
}

.membership-card:active {
  transform: scale(0.98);
  opacity: 0.95;
}

.membership-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0) 100%);
  pointer-events: none;
}

.membership-content {
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 32rpx;
  z-index: 2;
}

/* 背景装饰层 */
.membership-bg-decoration {
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

.circle-1 {
  width: 200rpx;
  height: 200rpx;
  top: -80rpx;
  right: -60rpx;
  background: rgba(255, 255, 255, 0.15);
}

.circle-2 {
  width: 150rpx;
  height: 150rpx;
  bottom: -40rpx;
  left: -40rpx;
  background: rgba(255, 255, 255, 0.1);
}

.circle-3 {
  width: 100rpx;
  height: 100rpx;
  top: 50%;
  right: 20rpx;
  background: rgba(255, 255, 255, 0.08);
}

.membership-header {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.membership-title-row {
  display: flex;
  align-items: flex-start;
  gap: 20rpx;
}

/* VIP 操作按钮 */
.membership-action {
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 12rpx;
}

.action-btn,
.activation-btn {
  min-width: 140rpx;
  height: 56rpx;
  padding: 0 24rpx;
  border-radius: 32rpx;
  border: 1px solid rgba(255, 255, 255, 0.4);
  backdrop-filter: blur(10rpx);
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
}

.action-btn {
  background: rgba(255, 255, 255, 0.25);
}

.action-btn:active {
  background: rgba(255, 255, 255, 0.35);
  transform: scale(0.95);
}

.action-text {
  font-size: 24rpx;
  font-weight: 600;
  color: #fff;
  white-space: nowrap;
}

.activation-btn {
  background: rgba(255, 255, 255, 0.12);
}

.activation-btn:active {
  background: rgba(255, 255, 255, 0.22);
  transform: scale(0.96);
}

.activation-text {
  font-size: 22rpx;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.9);
}

/* 会员图标 */
.membership-icon-wrapper {
  flex-shrink: 0;
  margin-top: 4rpx;
}

.membership-icon {
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

.icon-text {
  font-size: 44rpx;
}

/* 会员标题内容 */
.membership-title-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10rpx;
  min-width: 0;
}

.title-main {
  display: flex;
  align-items: center;
}

.membership-name {
  font-size: 40rpx;
  font-weight: 800;
  color: #ffffff;
  line-height: 1.2;
  letter-spacing: 0.5rpx;
  text-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.15);
}

/* 副标题行 */
.title-sub {
  display: flex;
  align-items: center;
  gap: 12rpx;
  flex-wrap: wrap;
}

.membership-type-code {
  font-size: 24rpx;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.85);
  letter-spacing: 2rpx;
  text-transform: uppercase;
}

.separator {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.5);
  font-weight: 300;
}

.status-inline {
  font-size: 24rpx;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.9);
}

.status-inline.active {
  color: var(--fg-success);
}

.status-inline.expired {
  color: var(--fg-danger);
}

.status-inline.disabled {
  color: var(--fg-text-weak);
}

.membership-desc {
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.6;
  padding-left: 108rpx;
  font-weight: 400;
  letter-spacing: 0.3rpx;
}

.wd-card.menu-card.is-rectangle {
  margin-top: 18rpx;
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

:deep(.menu-card .wd-card__content) {
  padding: 0;
}

:deep(.wd-cell) {
  transition: background-color 0.2s ease;

  &:active {
    background-color: var(--fg-bg-alt);
  }
}

:deep(.wd-cell__title) {
  font-size: 30rpx;
  font-weight: 600;
  color: var(--fg-text);
}

:deep(.wd-cell__label) {
  margin-top: 6rpx;
  color: var(--fg-text-muted);
}

:deep(.wd-cell__icon) {
  color: var(--fg-primary);
}
</style>
