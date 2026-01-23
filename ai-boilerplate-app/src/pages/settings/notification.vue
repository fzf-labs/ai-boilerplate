<script lang="ts" setup>
import { useToast } from 'wot-design-uni'
import type { NotificationCategory, NotificationSettingsInfo, UpdateNotificationSettingsReq } from '@/api/v1/user-notification-settings/types'
import { getNotificationSettings, updateNotificationSettings } from '@/api/v1/user-notification-settings/userNotificationSetting'
import { LOGIN_PAGE } from '@/router/config'
import { useTokenStore } from '@/store/token'

definePage({
  style: {
    navigationBarTitleText: '通知设置',
  },
})

const toast = useToast()
const tokenStore = useTokenStore()

const loading = ref(false)
const saving = ref(false)

const categories = ref<NotificationCategory[]>([])
const hasCategories = computed(() => categories.value.length > 0)
const dndSettings = reactive({
  dndStartTime: '',
  dndEndTime: '',
})

function handleOpenSystemSettings() {
  // #ifdef APP-PLUS
  const openSetting = (uni as any).openAppAuthorizeSetting
  if (typeof openSetting !== 'function') {
    toast.info('请在系统设置中开启通知权限')
    return
  }
  openSetting({
    fail: (error: unknown) => {
      console.error('打开系统设置失败:', error)
      toast.error('打开系统设置失败')
    },
  })
  // #endif

  // #ifndef APP-PLUS
  toast.info('请在系统设置中开启通知权限')
  // #endif
}

function ensureLogin() {
  if (tokenStore.hasLogin)
    return true

  toast.warning('请先登录')
  setTimeout(() => {
    uni.navigateTo({ url: LOGIN_PAGE })
  }, 1500)
  return false
}

function applySettings(info?: NotificationSettingsInfo | null) {
  const settings = info || null
  categories.value = (settings?.categories || []).map(item => ({
    key: item.key || '',
    title: item.title || '未命名通知',
    description: item.description || '',
    enabled: item.enabled ?? true,
  }))
  dndSettings.dndStartTime = settings?.dndStartTime || ''
  dndSettings.dndEndTime = settings?.dndEndTime || ''
}

async function fetchSettings() {
  if (!ensureLogin())
    return

  loading.value = true
  try {
    const res = await getNotificationSettings({ options: {} })
    applySettings(res.settings)
  }
  catch (error) {
    console.error('加载通知设置失败:', error)
    toast.error('加载失败')
  }
  finally {
    loading.value = false
  }
}

async function handleSave() {
  if (!ensureLogin() || saving.value)
    return

  saving.value = true
  try {
    const payload: UpdateNotificationSettingsReq = {
      preferences: categories.value.map(item => ({
        key: item.key,
        enabled: item.enabled ?? true,
      })),
    }
    if (dndSettings.dndStartTime)
      payload.dndStartTime = dndSettings.dndStartTime
    if (dndSettings.dndEndTime)
      payload.dndEndTime = dndSettings.dndEndTime

    await updateNotificationSettings({ body: payload, options: {} })
    toast.success('已保存')
  }
  catch (error) {
    console.error('保存通知设置失败:', error)
    toast.error('保存失败')
  }
  finally {
    saving.value = false
  }
}

onShow(() => {
  fetchSettings()
})
</script>

<template>
  <view class="notification-page">
    <view class="top-bg" />
    <view class="content">
      <view class="header-card">
        <view class="header-left">
          <view class="header-icon">
            <wd-icon name="setting" size="44rpx" color="var(--fg-text-inverse)" />
          </view>
          <view class="header-text">
            <text class="header-title">通知设置</text>
            <text class="header-subtitle">自定义不同类型消息的推送提醒</text>
          </view>
        </view>
      </view>

      <wd-card type="rectangle" custom-class="card">
        <template #title>
          <view class="card-title">
            消息通知
          </view>
        </template>

        <view v-if="loading" class="loading-box">
          <wd-loading />
        </view>
        <wd-cell-group v-else-if="hasCategories" border>
        <wd-cell
          v-for="item in categories"
          :key="item.key"
          :title="item.title"
          :label="item.description"
        >
          <wd-switch v-model="item.enabled" />
        </wd-cell>
        </wd-cell-group>
        <view v-else class="empty-box">
          暂无可用通知类型
        </view>
      </wd-card>

      <view class="action-bar">
        <wd-button
          :block="true"
          :round="true"
          size="large"
          type="primary"
          :loading="saving"
          :disabled="!hasCategories"
          @click="handleSave"
        >
          保存设置
        </wd-button>
      </view>
    </view>
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.notification-page {
  min-height: 100vh;
  background: var(--fg-bg);
  position: relative;
}

.top-bg {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  height: 240rpx;
  pointer-events: none;
  background: var(--fg-top-bg-gradient);
}

.content {
  position: relative;
  padding: 24rpx var(--fg-page-x) 40rpx;
}

.header-card {
  padding: 20rpx 24rpx;
  border-radius: 24rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20rpx;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.header-icon {
  width: 64rpx;
  height: 64rpx;
  border-radius: 20rpx;
  background: var(--fg-primary);
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.header-title {
  font-size: 32rpx;
  font-weight: 600;
  color: var(--fg-text);
}

.header-subtitle {
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.loading-box {
  padding: 24rpx 0;
  display: flex;
  justify-content: center;
}

.empty-box {
  padding: 24rpx 0;
  text-align: center;
  font-size: 24rpx;
  color: var(--fg-text-weak);
}

.action-bar {
  margin-top: 28rpx;
}

:deep(.card .wd-card__content) {
  padding: 0 12rpx 16rpx;
}
</style>
