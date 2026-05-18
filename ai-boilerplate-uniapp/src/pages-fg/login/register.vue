<script lang="ts" setup>
import { ref } from 'vue'
import { useToast } from 'wot-design-uni'
import { LOGIN_PAGE } from '@/router/config'

definePage({
  style: {
    navigationBarTitleText: '注册',
  },
})

const toast = useToast()
const agreeTerms = ref(false)

// 协议链接配置
const AGREEMENT_URLS = {
  userAgreement: 'https://example.com/user-agreement',
  privacyPolicy: 'https://example.com/privacy-policy',
}

function openAgreement(type: 'userAgreement' | 'privacyPolicy') {
  const url = AGREEMENT_URLS[type]
  uni.navigateTo({
    url: `/pages-fg/webview/index?url=${encodeURIComponent(url)}`,
  })
}

function doRegister() {
  if (!agreeTerms.value) {
    toast.warning('请先同意用户协议和隐私政策')
    return
  }
  toast.success('注册成功')
  // 注册成功后跳转到登录页
  uni.navigateTo({
    url: LOGIN_PAGE,
  })
}
</script>

<template>
  <view class="register-page">
    <view class="text-center">
      注册页
    </view>

    <view class="agreement-row">
      <wd-checkbox v-model="agreeTerms" shape="square" />
      <view class="agreement-text">
        <text>我已阅读并同意</text>
        <text class="link" @click.stop="openAgreement('userAgreement')">《用户协议》</text>
        <text>和</text>
        <text class="link" @click.stop="openAgreement('privacyPolicy')">《隐私政策》</text>
      </view>
    </view>

    <button class="mt-4 w-40 text-center" :disabled="!agreeTerms" @click="doRegister">
      点击模拟注册
    </button>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.register-page {
  padding: 32rpx;
}

.agreement-row {
  display: flex;
  align-items: flex-start;
  padding: 24rpx 0;
  gap: 8rpx;
}

.agreement-text {
  flex: 1;
  font-size: 24rpx;
  color: var(--fg-text-muted);
  line-height: 1.5;
}

.link {
  color: var(--fg-primary);
  font-weight: 600;
}
</style>
