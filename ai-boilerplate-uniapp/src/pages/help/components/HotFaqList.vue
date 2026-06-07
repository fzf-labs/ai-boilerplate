<script lang="ts" setup>
import type { HelpFaqInfo } from '@/api/v1/help-faq/types'

interface Props {
  faqs: HelpFaqInfo[]
  loading?: boolean
}

interface Emits {
  (e: 'select', faq: HelpFaqInfo): void
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
})

const emit = defineEmits<Emits>()

function handleSelect(faq: HelpFaqInfo) {
  emit('select', faq)
}
</script>

<template>
  <view class="faq-section">
    <view class="section-header">
      <view class="section-title">
        <view class="fire-icon">
          <wd-icon name="fire" size="28rpx" color="#ff6b6b" />
        </view>
        <text>热门问题</text>
      </view>
      <wd-tag type="danger" plain size="small">
        HOT
      </wd-tag>
    </view>

    <view v-if="!loading && faqs.length === 0" class="empty-state">
      <wd-icon name="inbox" size="56rpx" color="var(--fg-text-muted)" />
      <text class="empty-text">暂无热门问题</text>
    </view>

    <view v-else class="faq-list">
      <view
        v-for="(faq, index) in faqs"
        :key="faq.id"
        class="faq-item"
        :style="{ animationDelay: `${index * 0.05}s` }"
        @click="handleSelect(faq)"
      >
        <view class="faq-index">
          {{ index + 1 }}
        </view>
        <text class="faq-question">{{ faq.question }}</text>
        <wd-icon name="arrow-right" size="28rpx" color="var(--fg-text-muted)" />
      </view>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.faq-section {
  background: var(--fg-surface);
  border-radius: 20rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  padding: 24rpx;
  margin-bottom: 20rpx;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20rpx;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 8rpx;
  font-size: 28rpx;
  font-weight: 700;
  color: var(--fg-text);
}

.fire-icon {
  animation: fireGlow 2s ease-in-out infinite;
}

.faq-list {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
}

.faq-item {
  display: flex;
  align-items: center;
  gap: 16rpx;
  padding: 20rpx 16rpx;
  background: var(--fg-bg-alt);
  border-radius: 16rpx;
  border: 1px solid var(--fg-border-weak);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0;
  animation: fadeInRight 0.4s ease-out forwards;

  &:active {
    transform: scale(0.98);
    background: var(--fg-bg);
  }
}

.faq-index {
  width: 36rpx;
  height: 36rpx;
  border-radius: 10rpx;
  background: linear-gradient(135deg, var(--wot-color-primary) 0%, rgba(var(--wot-color-primary-rgb), 0.8) 100%);
  color: #fff;
  font-size: 22rpx;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 4rpx 12rpx rgba(var(--wot-color-primary-rgb), 0.25);
}

.faq-question {
  flex: 1;
  font-size: 26rpx;
  color: var(--fg-text);
  line-height: 1.5;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12rpx;
  padding: 40rpx 0;
}

.empty-text {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

@keyframes fadeInRight {
  from {
    opacity: 0;
    transform: translateX(10rpx);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes fireGlow {
  0%,
  100% {
    filter: drop-shadow(0 0 4rpx rgba(255, 107, 107, 0.3));
  }
  50% {
    filter: drop-shadow(0 0 8rpx rgba(255, 107, 107, 0.6));
  }
}
</style>
