<script lang="ts" setup>
interface Props {
  modelValue: string
}

interface Emits {
  (e: 'update:modelValue', value: string): void
  (e: 'search'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const searchValue = computed({
  get: () => props.modelValue,
  set: value => emit('update:modelValue', value),
})

function handleSearch() {
  emit('search')
}
</script>

<template>
  <view class="help-header">
    <view class="header-content">
      <view class="header-icon">
        <view class="icon-bg" />
        <wd-icon name="help-circle" size="48rpx" color="var(--wot-color-primary)" />
      </view>
      <view class="header-text">
        <text class="header-title">帮助中心</text>
        <text class="header-subtitle">有问题？我们来帮你</text>
      </view>
    </view>
    <view class="search-wrapper">
      <wd-search
        v-model="searchValue"
        :hide-cancel="true"
        placeholder="搜索你的问题..."
        @search="handleSearch"
      />
    </view>
  </view>
</template>

<style lang="scss" scoped>
.help-header {
  padding: 24rpx;
  background: var(--fg-surface);
  border-radius: 24rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  margin-bottom: 20rpx;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 20rpx;
  margin-bottom: 20rpx;
}

.header-icon {
  position: relative;
  width: 80rpx;
  height: 80rpx;
  border-radius: 20rpx;
  background: linear-gradient(
    135deg,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.1) 0%,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.05) 100%
  );
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.icon-bg {
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 20rpx;
  background: rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.2);
  animation: pulse 2s ease-out infinite;
}

.header-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4rpx;
}

.header-title {
  font-size: 36rpx;
  font-weight: 700;
  color: var(--fg-text);
  line-height: 1.2;
}

.header-subtitle {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  line-height: 1.4;
}

.search-wrapper {
  background: var(--fg-bg-alt);
  border-radius: 16rpx;
  padding: 8rpx;
  border: 1px solid var(--fg-border-weak);
  transition: all 0.3s ease;

  &:focus-within {
    border-color: var(--wot-color-primary);
    box-shadow: 0 0 0 4rpx rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.1);
  }
}

@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 0.2;
  }
  50% {
    transform: scale(1.1);
    opacity: 0;
  }
  100% {
    transform: scale(1.1);
    opacity: 0;
  }
}
</style>
