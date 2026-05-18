<script lang="ts" setup>
import { HOME_PAGE } from '@/utils'

definePage({
  style: {
    // 'custom' 表示开启自定义导航栏，默认 'default'
    navigationStyle: 'custom',
  },
})

function goHome() {
  uni.switchTab({
    url: HOME_PAGE,
    fail() {
      uni.reLaunch({ url: HOME_PAGE })
    },
  })
}

function goBackOrHome() {
  const pages = getCurrentPages?.() ?? []
  if (pages.length > 1) {
    uni.navigateBack()
    return
  }
  goHome()
}
</script>

<template>
  <view class="not-found-page">
    <view class="top-bg" />

    <view class="content">
      <view class="hero">
        <view class="hero-icon">
          <wd-icon name="/static/my-icons/smile.svg" size="56rpx" />
        </view>
        <view class="hero-title">
          页面走丢了
        </view>
        <view class="hero-subtitle">
          你访问的页面不存在或已被移除
        </view>
      </view>

      <wd-card type="rectangle" custom-class="not-found-card">
        <view class="card-body">
          <view class="code">
            404
          </view>
          <view class="tips">
            <view class="tip-item">
              · 检查是否输入了正确的路径
            </view>
            <view class="tip-item">
              · 尝试返回上一页，或回到首页继续使用
            </view>
          </view>

          <view class="actions">
            <wd-button :block="true" :round="true" size="large" type="primary" @click="goHome">
              返回首页
            </wd-button>
            <wd-button :block="true" :round="true" size="large" :plain="true" @click="goBackOrHome">
              返回上一页
            </wd-button>
          </view>
        </view>
      </wd-card>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.not-found-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 320rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: calc(var(--status-bar-height) + 28rpx) var(--fg-page-x) 40rpx;
}

.hero {
  padding: 0 6rpx 18rpx;
}

.hero-icon {
  width: 112rpx;
  height: 112rpx;
  border-radius: 56rpx;
  background: rgba(255, 255, 255, 0.18);
  border: 1px solid rgba(255, 255, 255, 0.22);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10rpx 26rpx rgba(2, 6, 23, 0.16);
}

.hero-title {
  margin-top: 18rpx;
  font-size: 44rpx;
  font-weight: 900;
  color: rgba(255, 255, 255, 0.96);
  letter-spacing: 0.5px;
}

.hero-subtitle {
  margin-top: 10rpx;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.82);
}

.card-body {
  padding: 22rpx 18rpx 18rpx;
}

.code {
  font-size: 76rpx;
  font-weight: 900;
  color: var(--fg-text);
  letter-spacing: 2px;
}

.tips {
  margin-top: 14rpx;
  padding: 14rpx 14rpx;
  border-radius: 20rpx;
  background: var(--fg-bg-alt);
  border: 1px solid var(--fg-border-weak);
}

.tip-item {
  font-size: 26rpx;
  color: var(--fg-text-secondary);
  line-height: 1.8;
}

.actions {
  margin-top: 18rpx;
  display: flex;
  flex-direction: column;
  gap: 14rpx;
  padding-bottom: calc(env(safe-area-inset-bottom) + 6rpx);
  padding-bottom: calc(constant(safe-area-inset-bottom) + 6rpx);
}

:deep(.wd-card.not-found-card.is-rectangle) {
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}
</style>
