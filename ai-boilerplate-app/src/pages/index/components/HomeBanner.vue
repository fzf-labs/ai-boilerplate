<script lang="ts" setup>
import type { PropType } from 'vue'
import { computed } from 'vue'

interface BannerDisplayItem {
  id?: string
  title?: string
  imageUrl?: string
  linkUrl?: string
  linkType?: string
  position?: string
  platform?: string
  sort?: number
}

const props = defineProps({
  list: {
    type: Array as PropType<BannerDisplayItem[]>,
    default: () => [],
  },
  height: {
    type: [Number, String] as PropType<number | string>,
    default: '460rpx',
  },
  interval: {
    type: Number,
    default: 4500,
  },
})

const emit = defineEmits<{
  (event: 'click', item: BannerDisplayItem): void
}>()

const autoplay = computed(() => props.list.length > 1)
</script>

<template>
  <view class="banner-section">
    <view class="banner-shell">
      <wd-swiper
        custom-class="banner-swiper"
        :list="list"
        :autoplay="autoplay"
        :interval="interval"
        indicator-position="bottom"
        value-key="imageUrl"
        :height="height"
      >
        <template #default="{ item }">
          <view class="banner-slide" @click="emit('click', item as BannerDisplayItem)">
            <image class="banner-slide__image" :src="(item as BannerDisplayItem).imageUrl" mode="aspectFill" />
            <view class="banner-slide__shade" />
            <view class="banner-slide__content">
              <text class="banner-slide__title">{{ (item as BannerDisplayItem).title || '精选活动' }}</text>
              <view class="banner-slide__meta">
                <text class="banner-slide__tag">精选推荐</text>
                <view class="banner-slide__cta">
                  <text class="banner-slide__cta-text">立即查看</text>
                </view>
              </view>
            </view>
          </view>
        </template>
        <template #indicator="{ current, total }">
          <view class="banner-indicator">
            <view class="banner-indicator__dots">
              <view
                v-for="index in total"
                :key="index"
                class="banner-indicator__dot" :class="[index - 1 === current ? 'is-active' : '']"
              />
            </view>
          </view>
        </template>
      </wd-swiper>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.banner-section {
  width: 100%;
  padding: 12rpx var(--fg-page-x) var(--fg-section-gap);
  box-sizing: border-box;
}

.banner-shell {
  position: relative;
  z-index: 1;
}

:deep(.banner-swiper) {
  position: relative;
  z-index: 1;
}

.banner-slide {
  position: relative;
  width: 100%;
  height: 100%;
  border-radius: var(--fg-radius-card-lg);
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  transform: translateZ(0);
}

.banner-slide__image {
  width: 100%;
  height: 100%;
  display: block;
}

.banner-slide__shade {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, var(--fg-glass-0) 0%, var(--fg-glass-50) 60%, var(--fg-glass-85) 100%);
}

.banner-slide__content {
  position: absolute;
  left: 22rpx;
  right: 22rpx;
  bottom: 22rpx;
  display: flex;
  flex-direction: column;
  gap: 12rpx;
  padding: 18rpx 20rpx;
  border-radius: var(--fg-radius-card);
  background: var(--fg-surface-glass);
  border: 1rpx solid var(--fg-glass-70);
  box-shadow: var(--fg-shadow-soft);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
}

.banner-slide__title {
  font-size: 34rpx;
  font-weight: 700;
  color: var(--fg-text);
  line-height: 1.4;
  letter-spacing: -0.2rpx;
}

.banner-slide__meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16rpx;
}

.banner-slide__tag {
  padding: 6rpx 16rpx;
  border-radius: 999rpx;
  background: var(--fg-ink-04);
  border: 1rpx solid var(--fg-border);
  color: var(--fg-text-secondary);
  font-size: 22rpx;
  font-weight: 600;
}

.banner-slide__cta {
  padding: 10rpx 22rpx;
  border-radius: 999rpx;
  background: var(--fg-primary);
  box-shadow: 0 8rpx 18rpx rgba(var(--fg-primary-rgb), 0.3);
}

.banner-slide__cta-text {
  font-size: 24rpx;
  color: var(--fg-text-inverse);
  font-weight: 600;
}

.banner-indicator {
  position: absolute;
  left: 50%;
  bottom: 22rpx;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  padding: 8rpx 14rpx;
  border-radius: 999rpx;
  background: var(--fg-glass-70);
  border: 1rpx solid var(--fg-border);
  box-shadow: var(--fg-shadow-soft);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
}

.banner-indicator__dots {
  display: flex;
  align-items: center;
  gap: 8rpx;
}

.banner-indicator__dot {
  width: 10rpx;
  height: 10rpx;
  border-radius: 999rpx;
  background: var(--fg-text-weak);
  transition: all 0.25s ease-out;
}

.banner-indicator__dot.is-active {
  width: 18rpx;
  background: var(--fg-primary);
  box-shadow: 0 4rpx 10rpx rgba(var(--fg-primary-rgb), 0.3);
}

@supports not (
  (
    backdrop-filter: blur(1px),
  )
) {
  .banner-slide__content,
  .banner-indicator {
    background: var(--fg-surface);
  }
}
</style>
