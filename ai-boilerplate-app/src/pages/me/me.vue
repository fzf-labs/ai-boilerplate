<script lang="ts" setup>
import type { UserInfo } from '@/api/v1/user/types'
import { getUserInfo } from '@/api/v1/user/user'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '我的',
  },
})

const tokenStore = useTokenStore()

// 用户详细信息
const userProfile = ref<UserInfo | null>(null)

const displayName = computed(() => userProfile.value?.nickname || userProfile.value?.phone || '用户')
const displayAvatar = computed(() => userProfile.value?.avatar || '/static/images/default-avatar.png')
const displayPhone = computed(() => userProfile.value?.phone || '')

// 菜单列表
const menuList = [
  {
    title: '个人信息',
    icon: 'edit',
    label: '完善头像昵称',
    path: '/pages-fg/profile/edit',
    needLogin: true,
  },
  {
    title: '账号安全',
    icon: 'lock-on',
    label: '密码/手机号/注销',
    path: '/pages-fg/security/index',
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
    path: '/pages-fg/settings/index',
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
 * 退出登录
 */
function handleLogout() {
  uni.showModal({
    title: '提示',
    content: '确定要退出登录吗？',
    success: (res) => {
      if (res.confirm) {
        tokenStore.logout()
        userProfile.value = null
        uni.showToast({
          title: '退出登录成功',
          icon: 'success',
        })
      }
    },
  })
}

/**
 * 菜单点击
 */
function handleMenuClick(item: typeof menuList[0]) {
  if (item.needLogin && !tokenStore.hasLogin) {
    uni.showToast({
      title: '请先登录',
      icon: 'none',
    })
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
 * 页面展示时刷新用户信息（登录后返回会触发）
 */
onShow(() => {
  if (tokenStore.hasLogin) {
    fetchUserProfile()
  }
  else {
    userProfile.value = null
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
      <wd-card type="rectangle" custom-class="quick-card">
        <wd-grid :column="4" :border="false">
          <wd-grid-item v-for="item in menuList" :key="item.title" @click="handleMenuClick(item)">
            <view class="quick-item">
              <wd-icon :name="item.icon" size="44rpx" />
              <text class="quick-text">{{ item.title }}</text>
            </view>
          </wd-grid-item>
        </wd-grid>
      </wd-card>

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

      <view v-if="tokenStore.hasLogin" class="logout">
        <wd-button :block="true" :round="true" size="large" type="error" @click="handleLogout">
          退出登录
        </wd-button>
      </view>
    </view>
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
  padding: 0 var(--fg-page-x) 40rpx;
}

.quick-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 14rpx;
  padding: 10rpx 0;
  color: var(--fg-text);
}

.quick-text {
  font-size: 24rpx;
  color: var(--fg-text-secondary);
}

.wd-card.quick-card.is-rectangle {
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

:deep(.quick-card .wd-card__content) {
  padding: 22rpx 8rpx;
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

.logout {
  margin-top: 22rpx;

  :deep(.wd-button) {
    border-radius: 16rpx;
    font-weight: 600;
    box-shadow: 0 6rpx 18rpx rgba(239, 68, 68, 0.18);
  }
}
</style>
