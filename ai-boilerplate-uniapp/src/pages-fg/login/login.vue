<script lang="ts" setup>
import type { FormInstance, FormRules } from 'wot-design-uni/components/wd-form/types'
import { computed, reactive, ref } from 'vue'
import { useToast } from 'wot-design-uni'
import { REGISTER_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'
import { isPageTabbar } from '@/tabbar/store'
import { parseUrlToObj } from '@/utils'
import { AGREEMENT_URLS } from './config'

definePage({
  style: {
    navigationBarTitleText: '登录',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()

const form = reactive({
  username: '',
  password: '',
})
const formRef = ref<FormInstance | null>(null)
const loading = ref(false)
const rememberAccount = ref(true)
const agreeTerms = ref(false)
const redirectPath = ref('')

const rules: FormRules = {
  username: [{ required: true, message: '请输入账号' }],
  password: [{ required: true, message: '请输入密码' }],
}

const canSubmit = computed(() => {
  return !!form.username.trim() && !!form.password && agreeTerms.value
})

function restoreRememberedAccount() {
  const remembered = uni.getStorageSync('fg.login.username')
  if (typeof remembered === 'string') {
    form.username = remembered
  }
}

function persistRememberedAccount() {
  if (!rememberAccount.value) {
    uni.removeStorageSync('fg.login.username')
    return
  }
  uni.setStorageSync('fg.login.username', form.username.trim())
}

function goBackOrMeTab() {
  const pages = getCurrentPages()
  if (pages.length > 1) {
    uni.navigateBack()
    return
  }
  uni.switchTab({ url: '/pages/me/me' })
}

function handleForgotPassword() {
  toast.info('请联系管理员重置密码')
}

function goRegister() {
  uni.navigateTo({ url: REGISTER_PAGE })
}

function safeDecode(value: string) {
  try {
    return decodeURIComponent(value)
  }
  catch {
    return value
  }
}

function navigateAfterLogin() {
  const target = redirectPath.value.trim()
  if (target) {
    const { path } = parseUrlToObj(target)
    if (path && isPageTabbar(path)) {
      uni.switchTab({ url: path })
    }
    else {
      uni.reLaunch({ url: target })
    }
    return
  }
  goBackOrMeTab()
}

function handleWeixinLoginTip() {
  toast.info('暂未接入微信登录')
}

function openAgreement(type: 'userAgreement' | 'privacyPolicy') {
  const url = AGREEMENT_URLS[type]
  if (!url) {
    toast.info('请先配置协议链接')
    return
  }
  uni.navigateTo({
    url: `/pages-fg/webview/index?url=${encodeURIComponent(url)}`,
  })
}

async function doLogin() {
  if (tokenStore.hasLogin) {
    goBackOrMeTab()
    return
  }
  if (loading.value)
    return

  const validateRes = await formRef.value?.validate()
  if (validateRes && !validateRes.valid)
    return

  loading.value = true
  try {
    await tokenStore.login({
      username: form.username.trim(),
      password: form.password,
    })
    persistRememberedAccount()
    navigateAfterLogin()
  }
  catch (error) {
    console.error('登录失败', error)
  }
  finally {
    loading.value = false
  }
}

onLoad((options) => {
  restoreRememberedAccount()
  redirectPath.value = options?.redirect ? safeDecode(options.redirect) : ''
})
</script>

<template>
  <view class="login-page">
    <view class="bg" />
    <view class="content">
      <view class="brand">
        <view class="logo">
          <wd-icon name="/static/my-icons/smile.svg" size="56rpx" />
        </view>
        <view class="brand-title">
          ai-boilerplate
        </view>
        <view class="brand-subtitle">
          欢迎回来，登录继续使用
        </view>
      </view>

      <wd-card type="rectangle" custom-class="login-card">
        <wd-form ref="formRef" :model="form" :rules="rules" error-type="toast">
          <wd-form-item label="账号" prop="username" required>
            <wd-input
              v-model="form.username"
              placeholder="请输入账号/手机号"
              prefix-icon="user"
              clearable
              :maxlength="50"
              confirm-type="next"
            />
          </wd-form-item>

          <wd-form-item label="密码" prop="password" required>
            <wd-input
              v-model="form.password"
              placeholder="请输入密码"
              prefix-icon="lock-on"
              type="text"
              show-password
              clearable
              :maxlength="50"
              confirm-type="done"
              @confirm="doLogin"
            />
          </wd-form-item>

          <view class="row">
            <wd-checkbox v-model="rememberAccount" shape="square">
              记住账号
            </wd-checkbox>
            <text class="link" @click="handleForgotPassword">
              忘记密码
            </text>
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

          <view class="actions">
            <wd-button
              type="primary"
              size="large"
              :block="true"
              :disabled="!canSubmit || loading"
              :loading="loading"
              @click="doLogin"
            >
              登录
            </wd-button>
          </view>

          <!-- #ifdef MP-WEIXIN -->
          <view class="other">
            <wd-divider>其他方式</wd-divider>
            <wd-button :block="true" :plain="true" :disabled="loading" @click="handleWeixinLoginTip">
              微信登录
            </wd-button>
          </view>
          <!-- #endif -->
        </wd-form>

        <template #footer>
          <view class="footer">
            <text class="muted">没有账号？</text>
            <text class="link" @click="goRegister">去注册</text>
          </view>
        </template>
      </wd-card>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.login-page {
  min-height: 100vh;
  position: relative;
  background: var(--fg-bg);
}

.bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 320rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
  border-top-left-radius: 44rpx;
  border-top-right-radius: 44rpx;
  border-bottom-left-radius: 44rpx;
  border-bottom-right-radius: 44rpx;
  overflow: hidden;
}

.content {
  position: relative;
  padding: 16rpx var(--fg-page-x) 40rpx;
}

.brand {
  padding: 4rpx 6rpx 18rpx;
}

.logo {
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

.brand-title {
  margin-top: 18rpx;
  font-size: 44rpx;
  font-weight: 800;
  color: rgba(255, 255, 255, 0.96);
  letter-spacing: 0.5px;
}

.brand-subtitle {
  margin-top: 10rpx;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.82);
}

.row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14rpx 8rpx 0;
}

.agreement-row {
  display: flex;
  align-items: flex-start;
  padding: 16rpx 8rpx 0;
  gap: 8rpx;
}

.agreement-text {
  flex: 1;
  font-size: 24rpx;
  color: var(--fg-text-muted);
  line-height: 1.5;
}

.actions {
  margin-top: 18rpx;
}

.other {
  margin-top: 20rpx;
}

.footer {
  display: flex;
  justify-content: center;
  gap: 10rpx;
  padding: 8rpx 0 6rpx;
  font-size: 26rpx;
}

.muted {
  color: var(--fg-text-muted);
}

.link {
  color: var(--fg-primary);
  font-weight: 600;
}

:deep(.wd-card.login-card.is-rectangle) {
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

:deep(.login-card .wd-card__content) {
  padding: 18rpx 8rpx 6rpx;
}

:deep(.login-card .wd-card__footer) {
  padding: 10rpx 12rpx 14rpx;
}
</style>
