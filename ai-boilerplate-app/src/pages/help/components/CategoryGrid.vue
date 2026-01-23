<script lang="ts" setup>
import type { HelpCategoryInfo } from '@/api/v1/help-category/types'

interface Props {
  categories: HelpCategoryInfo[]
  loading?: boolean
}

interface Emits {
  (e: 'select', categoryId: string): void
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
})

const emit = defineEmits<Emits>()

function handleSelect(categoryId?: string) {
  if (categoryId) {
    emit('select', categoryId)
  }
}
</script>

<template>
  <view class="category-section">
    <view class="section-header">
      <view class="section-title">
        <wd-icon name="apps" size="28rpx" color="var(--wot-color-primary)" />
        <text>分类</text>
      </view>
    </view>

    <wd-skeleton
      theme="paragraph"
      :row-col="[1, 1]"
      :loading="loading"
      animation="gradient"
    >
      <view v-if="categories.length === 0" class="empty-state">
        <wd-icon name="inbox" size="56rpx" color="var(--fg-text-muted)" />
        <text class="empty-text">暂无分类</text>
      </view>
      <view v-else class="category-grid">
        <view
          v-for="(category, index) in categories"
          :key="category.id"
          class="category-card"
          :style="{ animationDelay: `${index * 0.05}s` }"
          @click="handleSelect(category.id)"
        >
          <view class="category-icon">
            <wd-icon :name="category.icon || 'help'" size="32rpx" />
          </view>
          <text class="category-name">{{ category.name }}</text>
        </view>
      </view>
    </wd-skeleton>
  </view>
</template>

<style lang="scss" scoped>
.category-section {
  background: var(--fg-surface);
  border-radius: 20rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  padding: 24rpx;
  margin-bottom: 20rpx;
}

.section-header {
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

.category-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16rpx;
}

.category-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10rpx;
  padding: 20rpx 12rpx;
  background: var(--fg-bg-alt);
  border-radius: 16rpx;
  border: 1px solid var(--fg-border-weak);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0;
  animation: fadeInUp 0.4s ease-out forwards;

  &:active {
    transform: scale(0.95);
    background: var(--fg-bg);
  }
}

.category-icon {
  width: 64rpx;
  height: 64rpx;
  border-radius: 16rpx;
  background: linear-gradient(
    135deg,
    rgba(var(--fg-primary-rgb), 0.12) 0%,
    rgba(var(--fg-primary-rgb), 0.06) 100%
  );
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4rpx 12rpx rgba(var(--fg-primary-rgb), 0.08);
}

.category-name {
  font-size: 24rpx;
  color: var(--fg-text-secondary);
  font-weight: 600;
  text-align: center;
  line-height: 1.3;
  word-break: break-all;
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

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
