<script setup lang="ts">
// i-carbon-code
import type { CustomTabBarItem } from './types'
import { useToast } from 'wot-design-uni'
import { customTabbarEnable, needHideNativeTabbar, tabbarCacheEnable } from './config'
import { tabbarList, tabbarStore } from './store'

// #ifdef MP-WEIXIN
// 将自定义节点设置成虚拟的（去掉自定义组件包裹层），更加接近Vue组件的表现，能更好的使用flex属性
defineOptions({
  virtualHost: true,
})
// #endif

const toast = useToast()

/**
 * 中间的鼓包tabbarItem的点击事件
 */
function handleClickBulge() {
  toast.info('点击了中间的鼓包tabbarItem')
}

function handleClick(index: number) {
  // 点击原来的不做操作
  if (index === tabbarStore.curIdx) {
    return
  }
  if (tabbarList[index].isBulge) {
    handleClickBulge()
    return
  }
  const url = tabbarList[index].pagePath
  tabbarStore.setCurIdx(index)
  if (tabbarCacheEnable) {
    uni.switchTab({ url })
  }
  else {
    uni.navigateTo({ url })
  }
}
// #ifndef MP-WEIXIN || MP-ALIPAY
// 因为有了 custom:true， 微信里面不需要多余的hide操作
onLoad(() => {
  // 解决原生 tabBar 未隐藏导致有2个 tabBar 的问题
  needHideNativeTabbar
  && uni.hideTabBar({
    fail(err) {
      console.warn('hideTabBar fail: ', err)
    },
  })
})
// #endif

// #ifdef MP-ALIPAY
onMounted(() => {
  // 解决支付宝自定义tabbar 未隐藏导致有2个 tabBar 的问题; 注意支付宝很特别，需要在 onMounted 钩子调用
  customTabbarEnable // 另外，支付宝里面，只要是 customTabbar 都需要隐藏
  && uni.hideTabBar({
    fail(err) {
      console.warn('hideTabBar fail: ', err)
    },
  })
})
// #endif
const activeColor = 'var(--fg-primary, #007aff)'
const inactiveColor = 'var(--fg-text-muted, #8e8e93)'
function getColorByIndex(index: number) {
  return tabbarStore.curIdx === index ? activeColor : inactiveColor
}

function getImageByIndex(index: number, item: CustomTabBarItem) {
  if (!item.iconActive) {
    console.warn('image 模式下，需要配置 iconActive (高亮时的图片），否则无法切换高亮图片')
    return item.icon
  }
  return tabbarStore.curIdx === index ? item.iconActive : item.icon
}
</script>

<template>
  <view v-if="customTabbarEnable" class="tabbar-root pb-safe">
    <view class="border-and-fixed" @touchmove.stop.prevent>
      <view class="tabbar-row flex items-center">
        <view
          v-for="(item, index) in tabbarList" :key="index"
          class="tabbar-item flex flex-1 flex-col items-center justify-center"
          :class="{ 'is-active': tabbarStore.curIdx === index }"
          :style="{ color: getColorByIndex(index) }"
          @click="handleClick(index)"
        >
          <view v-if="item.isBulge" class="relative">
            <!-- 中间一个鼓包tabbarItem的处理 -->
            <view class="bulge">
              <!-- 中间鼓包 tabbarItem：通常是图片或 icon，点击后触发业务逻辑 -->
              <!-- 常见的是：扫描按钮、发布按钮、更多按钮等 -->
              <image class="mt-6rpx h-200rpx w-200rpx" src="/static/tabbar/scan.png" />
            </view>
          </view>
          <view v-else class="tabbar-item__content relative text-center">
            <template v-if="item.iconType === 'uiLib'">
              <!-- 以下内容请根据选择的 UI 库自行替换 -->
              <!-- 如：<wd-icon name="home" /> (https://wot-design-uni.cn/component/icon.html) -->
              <!-- 如：<uv-icon name="home" /> (https://www.uvui.cn/components/icon.html) -->
              <!-- 如：<sar-icon name="image" /> (https://sard.wzt.zone/sard-uniapp-docs/components/icon)(sar没有home图标^_^) -->
              <!-- <wd-icon :name="item.icon" size="20" /> -->
            </template>
            <template v-if="item.iconType === 'unocss' || item.iconType === 'iconfont'">
              <view :class="item.icon" class="tabbar-icon" />
            </template>
            <template v-if="item.iconType === 'image'">
              <image :src="getImageByIndex(index, item)" mode="scaleToFill" class="tabbar-icon" />
            </template>
            <view class="tabbar-label">
              {{ item.text }}
            </view>
            <!-- 角标显示 -->
            <view v-if="item.badge">
              <template v-if="item.badge === 'dot'">
                <view class="tabbar-badge-dot absolute right-0 top-0 h-2 w-2 rounded-full" />
              </template>
              <template v-else>
                <view class="tabbar-badge-count absolute top-0 box-border h-5 min-w-5 center rounded-full px-1 text-center text-xs text-white -right-3">
                  {{ item.badge > 99 ? '99+' : item.badge }}
                </view>
              </template>
            </view>
          </view>
        </view>
      </view>

      <view class="pb-safe" />
    </view>
    <wd-toast />
  </view>
</template>

<style scoped lang="scss">
.tabbar-root {
  height: 50px;
}

.tabbar-row {
  height: 50px;
  padding-top: 6rpx;
  padding-bottom: 2rpx;
  box-sizing: border-box;
}

.tabbar-item {
  transition: color 0.2s ease;
}

.tabbar-item.is-active .tabbar-icon {
  transform: translateY(-1rpx);
}

.tabbar-item__content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-width: 88rpx;
  padding: 0 18rpx;
}

.tabbar-icon {
  width: 24px;
  height: 24px;
  font-size: 24px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition:
    transform 0.2s ease,
    opacity 0.2s ease;
}

.tabbar-label {
  margin-top: 4rpx;
  font-size: 20rpx;
  line-height: 1.1;
  letter-spacing: 0.2rpx;
}

.tabbar-item.is-active .tabbar-label {
  font-weight: 600;
}

.border-and-fixed {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 1000;

  border-top: 1px solid var(--fg-border, var(--fg-border-light));
  background: linear-gradient(
    180deg,
    var(--fg-glass-50, rgba(255, 255, 255, 0.5)) 0%,
    var(--fg-glass-85, rgba(255, 255, 255, 0.85)) 100%
  );
  box-shadow: 0 -4rpx 16rpx var(--fg-ink-05);
  -webkit-backdrop-filter: blur(var(--fg-blur-strong, 24rpx));
  backdrop-filter: blur(var(--fg-blur-strong, 24rpx));
  box-sizing: border-box;
}

@supports not (
  (
    backdrop-filter: blur(1px),
  )
) {
  .border-and-fixed {
    background: var(--fg-surface, var(--fg-white));
  }
}
// 中间鼓包的样式
.bulge {
  position: absolute;
  top: -20px;
  left: 50%;
  transform-origin: top center;
  transform: translateX(-50%) scale(0.5) translateY(-33%);
  display: flex;
  justify-content: center;
  align-items: center;
  width: 250rpx;
  height: 250rpx;
  border-radius: 50%;
  background-color: var(--fg-surface-glass, var(--fg-white));
  border: 1px solid var(--fg-border, var(--fg-border-light));
  box-shadow: var(--fg-shadow-soft, 0 4rpx 12rpx var(--fg-ink-08));

  &:active {
    // opacity: 0.8;
  }
}

.tabbar-badge-dot,
.tabbar-badge-count {
  background: var(--fg-danger, #ff3b30);
  border: 2rpx solid var(--fg-surface, #ffffff);
  box-shadow: 0 2rpx 6rpx var(--fg-ink-04);
}

.tabbar-badge-count {
  font-size: 18rpx;
  font-weight: 600;
  line-height: 22rpx;
}
</style>
