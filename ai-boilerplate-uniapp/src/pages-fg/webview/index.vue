<script lang="ts" setup>
import { useToast } from 'wot-design-uni'

definePage({
  style: {
    navigationBarTitleText: '',
  },
})

const toast = useToast()
const url = ref('')
const title = ref('')
const loading = ref(true)

onLoad((options) => {
  const targetUrl = options?.url ? decodeURIComponent(options.url) : ''
  const pageTitle = options?.title ? decodeURIComponent(options.title) : ''

  // 如果没有 URL，跳转到 404 页面
  if (!targetUrl) {
    uni.redirectTo({
      url: '/pages-fg/404/index',
    })
    return
  }

  url.value = targetUrl
  title.value = pageTitle

  // 设置导航栏标题
  if (pageTitle) {
    uni.setNavigationBarTitle({ title: pageTitle })
  }
})

function handleLoad() {
  loading.value = false
}

function handleError() {
  loading.value = false
  toast.error('页面加载失败')
}
</script>

<template>
  <view class="webview-page">
    <wd-loading v-if="loading" custom-class="loading" />
    <web-view
      v-if="url"
      :src="url"
      @load="handleLoad"
      @error="handleError"
    />
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.webview-page {
  width: 100%;
  height: 100vh;
}

:deep(.loading) {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
}
</style>
