<script lang="ts" setup>
interface Emits {
  (e: 'retry'): void
}

const emit = defineEmits<Emits>()

function handleRetry() {
  emit('retry')
}
</script>

<template>
  <view class="error-state">
    <view class="error-icon">
      <wd-icon name="close-circle" size="80rpx" color="var(--wot-color-danger)" />
    </view>
    <text class="error-title">加载失败</text>
    <text class="error-desc">网络连接异常，请检查后重试</text>
    <wd-button type="primary" size="medium" @click="handleRetry">
      <wd-icon name="refresh" size="28rpx" custom-style="margin-right: 8rpx" />
      重新加载
    </wd-button>
  </view>
</template>

<style lang="scss" scoped>
.error-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 100rpx 40rpx;
  text-align: center;
  animation: fadeIn 0.4s ease-out;
}

.error-icon {
  margin-bottom: 24rpx;
  animation: shake 0.5s ease-in-out;
}

.error-title {
  font-size: 32rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 12rpx;
}

.error-desc {
  font-size: 26rpx;
  color: var(--fg-text-muted);
  margin-bottom: 32rpx;
  line-height: 1.6;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes shake {
  0%,
  100% {
    transform: translateX(0);
  }
  25% {
    transform: translateX(-10rpx);
  }
  75% {
    transform: translateX(10rpx);
  }
}
</style>
