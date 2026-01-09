<script lang="ts" setup>
import type { BindPhoneReq, ChangePasswordReq, DeleteAccountReq } from '@/api/v1/user/types'
import { bindPhone, changePassword, deleteAccount, sendVerifyCode } from '@/api/v1/user/user'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '账号安全',
  },
})

const tokenStore = useTokenStore()
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
    label: '不可恢复，请谨慎操作',
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

async function submitChangePassword() {
  if (submitting.value)
    return

  const oldPwd = changePasswordForm.oldPassword.trim()
  const newPwd = changePasswordForm.newPassword.trim()
  const confirmPwd = changePasswordForm.confirmPassword.trim()
  if (!oldPwd || !newPwd || !confirmPwd) {
    uni.showToast({ title: '请填写完整', icon: 'none' })
    return
  }
  if (newPwd.length < 6 || newPwd.length > 20) {
    uni.showToast({ title: '密码长度为6-20位', icon: 'none' })
    return
  }
  if (newPwd !== confirmPwd) {
    uni.showToast({ title: '两次密码不一致', icon: 'none' })
    return
  }

  submitting.value = true
  try {
    await changePassword({ body: { ...changePasswordForm } })
    uni.showToast({ title: '修改成功', icon: 'success' })
    changePasswordSheetVisible.value = false
  }
  catch (error) {
    console.error('修改密码失败:', error)
    uni.showToast({ title: '修改失败', icon: 'none' })
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
    uni.showToast({ title: '手机号格式不正确', icon: 'none' })
    return
  }

  submitting.value = true
  try {
    await sendVerifyCode({ body: { phone } })
    uni.showToast({ title: '验证码已发送', icon: 'success' })
    startCodeCountdown()
  }
  catch (error) {
    console.error('发送验证码失败:', error)
    uni.showToast({ title: '发送失败', icon: 'none' })
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
    uni.showToast({ title: '手机号格式不正确', icon: 'none' })
    return
  }
  if (!code) {
    uni.showToast({ title: '请输入验证码', icon: 'none' })
    return
  }

  submitting.value = true
  try {
    await bindPhone({ body: { phone, code } })
    uni.showToast({ title: '绑定成功', icon: 'success' })
    bindPhoneSheetVisible.value = false
    stopCodeCountdown()
  }
  catch (error) {
    console.error('绑定手机号失败:', error)
    uni.showToast({ title: '绑定失败', icon: 'none' })
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
    uni.showToast({ title: '请输入密码确认', icon: 'none' })
    return
  }

  submitting.value = true
  try {
    await deleteAccount({ body: { password } })
    uni.showToast({ title: '注销成功', icon: 'success' })
    deleteAccountSheetVisible.value = false
    setTimeout(() => {
      tokenStore.logout()
      uni.reLaunch({ url: '/pages/index/index' })
    }, 1200)
  }
  catch (error) {
    console.error('注销账号失败:', error)
    uni.showToast({ title: '注销失败', icon: 'none' })
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
      <view class="sheet">
        <view class="sheet-header">
          <view class="header-left">
            <view class="header-icon">
              <wd-icon name="shield" size="44rpx" color="#fff" />
            </view>
            <view class="header-text">
              <text class="header-title">账号安全</text>
              <text class="header-subtitle">保护你的账号与隐私信息</text>
            </view>
          </view>
          <wd-tag type="success" plain>
            已开启基础防护
          </wd-tag>
        </view>

        <wd-cell-group border>
          <wd-cell
            v-for="item in menuList"
            :key="item.title"
            :title="item.title"
            :label="item.label"
            :icon="item.icon"
            is-link
            clickable
            :custom-title-class="item.danger ? 'danger-title' : ''"
            :custom-icon-class="item.danger ? 'danger-icon' : ''"
            @click="handleMenuClick(item.action)"
          />
        </wd-cell-group>
      </view>

      <view class="tips">
        <wd-notice-bar
          text="安全提示：定期修改密码；绑定手机号用于验证与找回；注销账号后数据将被永久删除。"
          type="warning"
          :scrollable="false"
        />
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
          注销后，账号数据将被永久删除且无法恢复
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
  height: 260rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: 22rpx var(--fg-page-x) 40rpx;
}

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}

.sheet-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16rpx;
  padding: 22rpx 18rpx 18rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.header-icon {
  width: 88rpx;
  height: 88rpx;
  border-radius: 22rpx;
  background: linear-gradient(135deg, var(--fg-primary-600) 0%, var(--fg-primary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10rpx 26rpx rgba(var(--fg-primary-rgb), 0.22);
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 8rpx;
}

.header-title {
  font-size: 40rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.header-subtitle {
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

:deep(.wd-cell) {
  transition: background-color 0.2s ease;

  &:active {
    background-color: var(--fg-bg-alt);
  }
}

:deep(.wd-cell__icon) {
  color: var(--fg-primary);
}

:deep(.danger-title) {
  color: #ef4444;
  font-weight: 700;
}

:deep(.danger-icon) {
  color: #ef4444;
}

.tips-section {
  margin-top: 18rpx;
}

.tips {
  margin-top: 18rpx;
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
  font-size: 26rpx;
  line-height: 1.5;
  color: var(--fg-text-muted);
  background: var(--fg-bg-alt);
  border-radius: 16rpx;
  border: 1px solid var(--fg-border-weak);
}
</style>
