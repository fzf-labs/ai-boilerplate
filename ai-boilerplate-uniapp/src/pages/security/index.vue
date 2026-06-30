<script lang="ts" setup>
import type { BindPhoneReq, ChangePasswordReq, DeleteAccountReq } from '@/api/v1/user/types'
import { useMessage, useToast } from 'wot-design-uni'
import { bindPhone, changePassword, deleteAccount, sendVerifyCode } from '@/api/v1/user/user'
import { useTokenStore } from '@/store/token'
import { DELETE_ACCOUNT_POLICY_COPY } from './delete-account-policy'

definePage({
  style: {
    navigationBarTitleText: '账号管理',
  },
})

const tokenStore = useTokenStore()
const toast = useToast()
const message = useMessage()
const submitting = ref(false)

// 菜单列表
const menuList = [
  {
    title: '修改密码',
    icon: 'lock-on',
    label: '定期更换密码更安全',
    action: 'changePassword',
  },
  {
    title: '绑定手机号',
    icon: 'phone',
    label: '用于验证与找回账号',
    action: 'bindPhone',
  },
  {
    title: '注销账号',
    icon: 'delete1',
    label: '清理资料与绑定，保留脱敏账务记录',
    action: 'deleteAccount',
    danger: true,
  },
]

const changePasswordSheetVisible = ref(false)
const bindPhoneSheetVisible = ref(false)
const deleteAccountSheetVisible = ref(false)

const changePasswordForm = reactive<ChangePasswordReq>({
  oldPassword: '',
  newPassword: '',
  confirmPassword: '',
})

const bindPhoneForm = reactive<BindPhoneReq>({
  phone: '',
  code: '',
})

const deleteAccountForm = reactive<DeleteAccountReq>({
  password: '',
})

const codeCountdown = ref(0)
let codeTimer: number | null = null

function resetChangePasswordForm() {
  changePasswordForm.oldPassword = ''
  changePasswordForm.newPassword = ''
  changePasswordForm.confirmPassword = ''
}

function resetBindPhoneForm() {
  bindPhoneForm.phone = ''
  bindPhoneForm.code = ''
  stopCodeCountdown()
}

function resetDeleteAccountForm() {
  deleteAccountForm.password = ''
}

function stopCodeCountdown() {
  if (codeTimer) {
    clearInterval(codeTimer)
    codeTimer = null
  }
  codeCountdown.value = 0
}

function startCodeCountdown() {
  stopCodeCountdown()
  codeCountdown.value = 60
  codeTimer = setInterval(() => {
    codeCountdown.value -= 1
    if (codeCountdown.value <= 0) {
      stopCodeCountdown()
    }
  }, 1000) as any
}

/**
 * 菜单点击
 */
function handleMenuClick(action: string) {
  switch (action) {
    case 'changePassword':
      resetChangePasswordForm()
      changePasswordSheetVisible.value = true
      break
    case 'bindPhone':
      resetBindPhoneForm()
      bindPhoneSheetVisible.value = true
      break
    case 'deleteAccount':
      resetDeleteAccountForm()
      deleteAccountSheetVisible.value = true
      break
  }
}

/**
 * 退出登录
 */
async function handleLogout() {
  try {
    await message.confirm({
      title: '提示',
      msg: '确定要退出登录吗？',
      confirmButtonText: '确定',
      cancelButtonText: '取消',
    })
    tokenStore.logout()
    toast.success('退出登录成功')
    setTimeout(() => {
      uni.switchTab({ url: '/pages/me/me' })
    }, 1200)
  }
  catch {
    // 用户取消
  }
}

async function submitChangePassword() {
  if (submitting.value)
    return

  const oldPwd = changePasswordForm.oldPassword.trim()
  const newPwd = changePasswordForm.newPassword.trim()
  const confirmPwd = changePasswordForm.confirmPassword.trim()
  if (!oldPwd || !newPwd || !confirmPwd) {
    toast.warning('请填写完整')
    return
  }
  if (newPwd.length < 6 || newPwd.length > 20) {
    toast.warning('密码长度为6-20位')
    return
  }
  if (newPwd !== confirmPwd) {
    toast.warning('两次密码不一致')
    return
  }

  submitting.value = true
  try {
    await changePassword({ body: { ...changePasswordForm } })
    toast.success('修改成功')
    changePasswordSheetVisible.value = false
  }
  catch (error) {
    console.error('修改密码失败:', error)
    toast.error('修改失败')
  }
  finally {
    submitting.value = false
  }
}

async function sendCode() {
  if (submitting.value || codeCountdown.value > 0)
    return

  const phone = bindPhoneForm.phone.trim()
  const phoneReg = /^1[3-9]\d{9}$/
  if (!phoneReg.test(phone)) {
    toast.warning('手机号格式不正确')
    return
  }

  submitting.value = true
  try {
    await sendVerifyCode({ body: { phone } })
    toast.success('验证码已发送')
    startCodeCountdown()
  }
  catch (error) {
    console.error('发送验证码失败:', error)
    toast.error('发送失败')
  }
  finally {
    submitting.value = false
  }
}

async function submitBindPhone() {
  if (submitting.value)
    return

  const phone = bindPhoneForm.phone.trim()
  const code = bindPhoneForm.code.trim()
  const phoneReg = /^1[3-9]\d{9}$/
  if (!phoneReg.test(phone)) {
    toast.warning('手机号格式不正确')
    return
  }
  if (!code) {
    toast.warning('请输入验证码')
    return
  }

  submitting.value = true
  try {
    await bindPhone({ body: { phone, code } })
    toast.success('绑定成功')
    bindPhoneSheetVisible.value = false
    stopCodeCountdown()
  }
  catch (error) {
    console.error('绑定手机号失败:', error)
    toast.error('绑定失败')
  }
  finally {
    submitting.value = false
  }
}

async function submitDeleteAccount() {
  if (submitting.value)
    return

  const password = deleteAccountForm.password.trim()
  if (!password) {
    toast.warning('请输入密码确认')
    return
  }

  submitting.value = true
  try {
    await deleteAccount({ body: { password } })
    toast.success('注销成功')
    deleteAccountSheetVisible.value = false
    setTimeout(() => {
      tokenStore.logout()
      uni.reLaunch({ url: '/pages/index/index' })
    }, 1200)
  }
  catch (error) {
    console.error('注销账号失败:', error)
    toast.error('注销失败')
  }
  finally {
    submitting.value = false
  }
}

onUnload(() => {
  stopCodeCountdown()
})
</script>

<template>
  <view class="security-page">
    <view class="top-bg" />
    <view class="content">
      <view class="hero">
        <view class="hero-main">
          <view class="hero-icon">
            <wd-icon name="user" size="44rpx" color="var(--fg-text-inverse)" />
          </view>
          <view class="hero-text">
            <text class="hero-title">账号管理</text>
            <text class="hero-subtitle">管理你的账号与安全设置</text>
          </view>
        </view>
        <view class="hero-status">
          <view class="status-dot" />
          <text class="status-text">已开启基础防护</text>
        </view>
      </view>

      <view class="section">
        <view class="section-card">
          <wd-cell-group border>
            <wd-cell
              v-for="item in menuList"
              :key="item.title"
              :title="item.title"
              :label="item.label"
              :icon="item.icon"

              is-link clickable center
              :custom-title-class="item.danger ? 'danger-title' : ''"
              :custom-icon-class="item.danger ? 'cell-icon danger-icon' : 'cell-icon'"
              @click="handleMenuClick(item.action)"
            />
          </wd-cell-group>
        </view>
      </view>

      <view class="section logout">
        <view class="section-card action-card">
          <wd-button class="logout-button" :block="true" :round="true" size="large" type="error" @click="handleLogout">
            退出登录
          </wd-button>
        </view>
      </view>
    </view>

    <bottom-sheet
      v-model="changePasswordSheetVisible"
      title="修改密码"
      confirm-text="确认修改"
      cancel-text="取消"
      :close-on-select="false"
      :confirm-disabled="submitting"
      @confirm="submitChangePassword"
    >
      <view class="sheet-form">
        <wd-input
          v-model="changePasswordForm.oldPassword"
          placeholder="请输入旧密码"
          type="safe-password"
          show-password
          clearable
        />
        <wd-input
          v-model="changePasswordForm.newPassword"
          placeholder="请输入新密码（6-20位）"
          type="safe-password"
          show-password
          clearable
        />
        <wd-input
          v-model="changePasswordForm.confirmPassword"
          placeholder="请确认新密码"
          type="safe-password"
          show-password
          clearable
        />
      </view>
    </bottom-sheet>

    <bottom-sheet
      v-model="bindPhoneSheetVisible"
      title="绑定手机号"
      confirm-text="确认绑定"
      cancel-text="取消"
      :close-on-select="false"
      :confirm-disabled="submitting"
      @confirm="submitBindPhone"
      @close="resetBindPhoneForm"
      @cancel="resetBindPhoneForm"
    >
      <view class="sheet-form">
        <wd-input
          v-model="bindPhoneForm.phone"
          placeholder="请输入手机号"
          type="tel"
          inputmode="tel"
          clearable
          :maxlength="11"
        />

        <view class="code-row">
          <view class="code-input">
            <wd-input
              v-model="bindPhoneForm.code"
              placeholder="请输入验证码"
              type="number"
              inputmode="numeric"
              clearable
              :maxlength="6"
            />
          </view>
          <wd-button
            size="small"
            type="primary"
            :disabled="submitting || codeCountdown > 0"
            @click="sendCode"
          >
            {{ codeCountdown > 0 ? `${codeCountdown}s` : '发送验证码' }}
          </wd-button>
        </view>
      </view>
    </bottom-sheet>

    <bottom-sheet
      v-model="deleteAccountSheetVisible"
      title="注销账号"
      confirm-text="确认注销"
      cancel-text="取消"
      confirm-variant="danger"
      :close-on-select="false"
      :confirm-disabled="submitting"
      @confirm="submitDeleteAccount"
    >
      <view class="sheet-form">
        <view class="danger-tip">
          {{ DELETE_ACCOUNT_POLICY_COPY }}
        </view>
        <wd-input
          v-model="deleteAccountForm.password"
          placeholder="请输入密码确认"
          type="safe-password"
          show-password
          clearable
        />
      </view>
    </bottom-sheet>

    <wd-message-box />
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.security-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 320rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient-strong);

  &::after {
    content: '';
    position: absolute;
    right: -40rpx;
    top: -120rpx;
    width: 320rpx;
    height: 320rpx;
    border-radius: 50%;
    background: radial-gradient(circle at 30% 30%, rgba(var(--fg-primary-rgb), 0.18), rgba(255, 255, 255, 0) 70%);
    opacity: 0.9;
  }
}

.content {
  position: relative;
  padding: 22rpx var(--fg-page-x) 40rpx;
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.hero {
  display: flex;
  flex-direction: column;
  gap: 16rpx;
  padding: 20rpx 18rpx;
  background: var(--fg-surface-glass);
  border-radius: 28rpx;
  border: 1px solid var(--fg-glass-70);
  box-shadow: var(--fg-shadow-soft);
  -webkit-backdrop-filter: blur(var(--fg-blur-soft));
  backdrop-filter: blur(var(--fg-blur-soft));
  animation: ios-fade-up 0.45s ease both;
}

.hero-main {
  display: flex;
  align-items: center;
  gap: 16rpx;
  width: 100%;
}

.hero-icon {
  width: 88rpx;
  height: 88rpx;
  border-radius: 22rpx;
  background: linear-gradient(140deg, rgba(var(--fg-primary-rgb), 0.25), rgba(var(--fg-primary-rgb), 0.08));
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10rpx 20rpx rgba(var(--fg-primary-rgb), 0.2);
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: 8rpx;
  flex: 1;
}

.hero-title {
  font-size: 40rpx;
  font-weight: 700;
  letter-spacing: -0.4rpx;
  color: var(--fg-text);
}

.hero-subtitle {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

.hero-status {
  align-self: flex-start;
  display: inline-flex;
  align-items: center;
  gap: 10rpx;
  padding: 10rpx 16rpx;
  border-radius: 999rpx;
  background: rgba(52, 199, 89, 0.12);
  color: var(--fg-text-secondary);
  font-size: 22rpx;
}

.status-dot {
  width: 14rpx;
  height: 14rpx;
  border-radius: 50%;
  background: var(--fg-success);
  box-shadow: 0 0 0 6rpx rgba(52, 199, 89, 0.12);
}

.status-text {
  font-weight: 600;
}

.section {
  display: flex;
  flex-direction: column;
  gap: 12rpx;
  animation: ios-fade-up 0.45s ease both;
}

.section:nth-of-type(2) {
  animation-delay: 0.06s;
}

.section:nth-of-type(3) {
  animation-delay: 0.12s;
}

.section:nth-of-type(4) {
  animation-delay: 0.18s;
}

.section-card {
  background: var(--fg-surface);
  border-radius: 24rpx;
  overflow: hidden;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

.tips-card {
  padding: 8rpx;
}

.action-card {
  padding: 12rpx;
}

.sheet-form {
  display: flex;
  flex-direction: column;
  gap: 14rpx;
}

.code-row {
  display: flex;
  align-items: center;
  gap: 12rpx;
}

.code-input {
  flex: 1;
  min-width: 0;
}

.danger-tip {
  padding: 14rpx 12rpx;
  font-size: 24rpx;
  line-height: 1.5;
  color: var(--fg-text-muted);
  background: var(--fg-bg-alt);
  border-radius: 16rpx;
  border: 1px solid var(--fg-border-weak);
}

:deep(.wd-cell) {
  transition: background-color 0.2s ease;

  &:active {
    background-color: var(--fg-bg-alt);
  }
}

:deep(.wd-cell__title) {
  font-size: 28rpx;
  font-weight: 600;
}

:deep(.wd-cell__label) {
  font-size: 22rpx;
  color: var(--fg-text-muted);
}

:deep(.wd-cell__right) {
  color: var(--fg-text-weak);
}

:deep(.wd-cell__arrow-right) {
  color: var(--fg-text-weak);
  font-size: 26rpx;
}

:deep(.wd-cell__icon) {
  width: 56rpx;
  height: 56rpx;
  border-radius: 16rpx;
  background: rgba(var(--fg-primary-rgb), 0.12);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

:deep(.cell-icon) {
  color: var(--fg-primary);
}

:deep(.danger-title) {
  color: var(--fg-danger);
  font-weight: 700;
}

:deep(.danger-icon) {
  background: rgba(255, 59, 48, 0.12);
  color: var(--fg-danger);
}

:deep(.wd-notice-bar) {
  border-radius: 16rpx;
  border: 1px solid rgba(255, 159, 10, 0.18);
  background: rgba(255, 159, 10, 0.12);
}

:deep(.wd-notice-bar__content) {
  color: var(--fg-text-secondary);
}

:deep(.logout-button) {
  border-radius: 18rpx;
  font-weight: 600;
  background: var(--fg-surface);
  color: var(--fg-danger);
  border: 1px solid rgba(255, 59, 48, 0.24);
  box-shadow: none;
}

@keyframes ios-fade-up {
  from {
    opacity: 0;
    transform: translateY(14rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (prefers-reduced-motion: reduce) {
  .hero,
  .section {
    animation: none;
  }
}
</style>
