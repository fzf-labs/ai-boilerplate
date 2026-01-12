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
            <view class="banner-slide__glow" />
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
  padding: calc(env(safe-area-inset-top) + 12rpx) var(--fg-page-x) var(--fg-section-gap);
  box-sizing: border-box;
}

.banner-shell {
  position: relative;
  z-index: 1;
}

.banner-shell::before,
.banner-shell::after {
  content: '';
  position: absolute;
  width: 220rpx;
  height: 220rpx;
  border-radius: 50%;
  filter: blur(30rpx);
  opacity: 0.7;
  z-index: 0;
}

.banner-shell::before {
  top: -60rpx;
  left: -30rpx;
  background: radial-gradient(circle at 30% 30%, rgba(255, 205, 155, 0.9), rgba(255, 205, 155, 0) 70%);
}

.banner-shell::after {
  bottom: -50rpx;
  right: -20rpx;
  background: radial-gradient(circle at 70% 70%, rgba(145, 198, 255, 0.9), rgba(145, 198, 255, 0) 70%);
}

:deep(.banner-swiper) {
  position: relative;
  z-index: 1;
}

.banner-slide {
  position: relative;
  width: 100%;
  height: 100%;
  border-radius: 32rpx;
  overflow: hidden;
  box-shadow: 0 24rpx 60rpx rgba(11, 18, 32, 0.28);
  background: #0b1220;
  transform: translateZ(0);
}

.banner-slide__image {
  width: 100%;
  height: 100%;
  display: block;
}

.banner-slide__glow {
  position: absolute;
  inset: -20% auto auto -15%;
  width: 60%;
  height: 60%;
  background: radial-gradient(circle at 30% 30%, rgba(255, 255, 255, 0.35), rgba(255, 255, 255, 0));
  opacity: 0.7;
}

.banner-slide__shade {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(8, 12, 20, 0) 0%, rgba(8, 12, 20, 0.45) 55%, rgba(8, 12, 20, 0.85) 100%);
}

.banner-slide__content {
  position: absolute;
  left: 28rpx;
  right: 28rpx;
  bottom: 28rpx;
  display: flex;
  flex-direction: column;
  gap: 16rpx;
}

.banner-slide__title {
  font-size: 38rpx;
  font-weight: 700;
  color: #ffffff;
  line-height: 1.4;
  letter-spacing: 0.5rpx;
  text-shadow: 0 8rpx 20rpx rgba(0, 0, 0, 0.35);
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
  background: rgba(255, 255, 255, 0.22);
  border: 1rpx solid rgba(255, 255, 255, 0.3);
  color: #ffffff;
  font-size: 22rpx;
  font-weight: 500;
}

.banner-slide__cta {
  padding: 10rpx 22rpx;
  border-radius: 999rpx;
  background: rgba(255, 255, 255, 0.95);
  box-shadow: 0 10rpx 24rpx rgba(13, 20, 34, 0.2);
}

.banner-slide__cta-text {
  font-size: 24rpx;
  color: #1f2a44;
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
  background: rgba(8, 12, 20, 0.35);
  border: 1rpx solid rgba(255, 255, 255, 0.18);
  box-shadow: 0 8rpx 22rpx rgba(0, 0, 0, 0.25);
  backdrop-filter: blur(10rpx);
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
  background: rgba(255, 255, 255, 0.35);
  transition: all 0.3s ease;
}

.banner-indicator__dot.is-active {
  width: 22rpx;
  background: #ffffff;
  box-shadow: 0 4rpx 12rpx rgba(255, 255, 255, 0.35);
}
</style>
