<script lang="ts" setup>
import type { FormInstance, FormRules } from 'wot-design-uni/components/wd-form/types'
import { computed, reactive, ref } from 'vue'
import { useToast } from 'wot-design-uni'
import { register } from '@/api/v1/user/user'
import { LOGIN_PAGE } from '@/router/config'
import { AGREEMENT_URLS } from './config'

definePage({
  style: {
    navigationBarTitleText: '注册',
  },
})

const toast = useToast()
const formRef = ref<FormInstance | null>(null)
const loading = ref(false)
const agreeTerms = ref(false)

const form = reactive({
  phone: '',
  password: '',
  confirmPassword: '',
  nickname: '',
})

const rules: FormRules = {
  phone: [{ required: true, message: '请输入手机号' }],
  password: [{ required: true, message: '请输入密码' }],
  confirmPassword: [{ required: true, message: '请再次输入密码' }],
}

const canSubmit = computed(() => {
  return !!form.phone.trim()
    && !!form.password
    && !!form.confirmPassword
    && agreeTerms.value
    && !loading.value
})

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

function goLogin() {
  uni.navigateTo({
    url: LOGIN_PAGE,
  })
}

async function doRegister() {
  if (loading.value)
    return

  const validateRes = await formRef.value?.validate()
  if (validateRes && !validateRes.valid)
    return

  const phone = form.phone.trim()
  const password = form.password
  const confirmPassword = form.confirmPassword
  const nickname = form.nickname.trim()
  const phoneReg = /^1[3-9]\d{9}$/

  if (!phoneReg.test(phone)) {
    toast.warning('手机号格式不正确')
    return
  }
  if (password.length < 6 || password.length > 32) {
    toast.warning('密码长度为6-32位')
    return
  }
  if (password !== confirmPassword) {
    toast.warning('两次密码不一致')
    return
  }
  if (!agreeTerms.value) {
    toast.warning('请先同意用户协议和隐私政策')
    return
  }

  loading.value = true
  try {
    await register({
      body: {
        phone,
        password,
        confirmPassword,
        nickname: nickname || undefined,
      },
    })
    uni.setStorageSync('fg.login.username', phone)
    toast.success('注册成功')
    setTimeout(() => {
      goLogin()
    }, 1200)
  }
  catch (error) {
    console.error('注册失败:', error)
    toast.error('注册失败')
  }
  finally {
    loading.value = false
  }
}
</script>

<template>
  <view class="register-page">
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
          创建账号，开始使用
        </view>
      </view>

      <wd-card type="rectangle" custom-class="register-card">
        <wd-form ref="formRef" :model="form" :rules="rules" error-type="toast">
          <wd-form-item label="手机号" prop="phone" required>
            <wd-input
              v-model="form.phone"
              placeholder="请输入手机号"
              prefix-icon="phone"
              clearable
              :maxlength="11"
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
              :maxlength="32"
              confirm-type="next"
            />
          </wd-form-item>

          <wd-form-item label="确认密码" prop="confirmPassword" required>
            <wd-input
              v-model="form.confirmPassword"
              placeholder="请再次输入密码"
              prefix-icon="lock-on"
              type="text"
              show-password
              clearable
              :maxlength="32"
              confirm-type="next"
            />
          </wd-form-item>

          <wd-form-item label="昵称" prop="nickname">
            <wd-input
              v-model="form.nickname"
              placeholder="请输入昵称（选填）"
              prefix-icon="user"
              clearable
              :maxlength="50"
              confirm-type="done"
              @confirm="doRegister"
            />
          </wd-form-item>

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
              :disabled="!canSubmit"
              :loading="loading"
              @click="doRegister"
            >
              注册
            </wd-button>
          </view>
        </wd-form>

        <template #footer>
          <view class="footer">
            <text class="muted">已有账号？</text>
            <text class="link" @click="goLogin">去登录</text>
          </view>
        </template>
      </wd-card>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.register-page {
  min-height: 100vh;
  position: relative;
  background: var(--fg-bg);
}

.bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.92), rgba(248, 249, 252, 1));
}

.content {
  position: relative;
  z-index: 1;
  padding: 96rpx 32rpx 48rpx;
}

.brand {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 40rpx;
}

.logo {
  width: 112rpx;
  height: 112rpx;
  border-radius: 28rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 24rpx;
  background: rgba(255, 255, 255, 0.72);
  box-shadow: 0 10rpx 30rpx rgba(15, 23, 42, 0.08);
}

.brand-title {
  font-size: 40rpx;
  font-weight: 700;
  color: var(--fg-text-primary);
  line-height: 1.2;
}

.brand-subtitle {
  margin-top: 12rpx;
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

.register-card {
  border-radius: 24rpx;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(18rpx);
  box-shadow: 0 24rpx 60rpx rgba(15, 23, 42, 0.08);
}

.agreement-row {
  display: flex;
  align-items: flex-start;
  padding: 24rpx 0 8rpx;
  gap: 8rpx;
}

.agreement-text {
  flex: 1;
  font-size: 24rpx;
  color: var(--fg-text-muted);
  line-height: 1.5;
}

.actions {
  margin-top: 24rpx;
}

.footer {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8rpx;
  font-size: 24rpx;
}

.muted {
  color: var(--fg-text-muted);
}

.link {
  color: var(--fg-primary);
  font-weight: 600;
}
</style>
