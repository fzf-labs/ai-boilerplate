<script lang="ts" setup>
interface ICacheInfo {
  sizeKB: number
  sizeText: string
}

interface IAppVersion {
  currentVersion: string
  latestVersion?: string
  hasUpdate: boolean
  updateDesc?: string
  downloadUrl?: string
}

definePage({
  style: {
    navigationBarTitleText: '通用设置',
  },
})

// 缓存信息
const cacheInfo = ref<ICacheInfo | null>(null)
// 版本信息
const versionInfo = ref<IAppVersion | null>(null)
// 当前语言
const currentLanguage = ref('zh-CN')
const languageSheetVisible = ref(false)
const clearCacheSheetVisible = ref(false)
const aboutSheetVisible = ref(false)
const updateSheetVisible = ref(false)

// 语言选项
const languageOptions = [
  { label: '简体中文', value: 'zh-CN' },
  { label: 'English', value: 'en-US' },
]

const languageActions = computed(() => {
  return languageOptions.map((item) => {
    return {
      name: item.label,
      color: item.value === currentLanguage.value ? 'var(--fg-primary)' : '',
    }
  })
})

const clearCacheActions = [
  {
    name: '清除缓存（可能导致退出登录）',
    subname: '不可撤销',
    color: 'var(--fg-danger)',
  },
]

// 菜单列表
const menuList = [
  {
    title: '语言切换',
    icon: 'translate-bold',
    label: '切换应用显示语言',
    action: 'changeLanguage',
    showValue: true,
    value: '简体中文',
  },
  {
    title: '清除缓存',
    icon: 'delete1',
    label: '释放存储空间（可能导致退出登录）',
    action: 'clearCache',
    showValue: true,
    value: '0 MB',
  },
]

const legalList = [
  {
    title: '隐私协议',
    icon: 'file',
    label: '了解我们如何保护你的隐私',
    action: 'privacy',
    showValue: false,
  },
  {
    title: '用户协议',
    icon: 'file',
    label: '服务条款与使用规范',
    action: 'terms',
    showValue: false,
  },
]

const aboutList = [
  {
    title: '关于我们',
    icon: 'info-circle',
    label: '版本信息与联系方式',
    action: 'about',
    showValue: false,
  },
  {
    title: '版本更新',
    icon: 'refresh',
    label: '检查是否有新版本',
    action: 'checkVersion',
    showValue: true,
    value: 'v1.0.0',
  },
]

function formatSize(kb: number): string {
  if (!Number.isFinite(kb) || kb <= 0)
    return '0 KB'
  if (kb < 1024)
    return `${kb.toFixed(0)} KB`
  const mb = kb / 1024
  if (mb < 1024)
    return `${mb.toFixed(2)} MB`
  const gb = mb / 1024
  return `${gb.toFixed(2)} GB`
}

async function getCurrentVersion(): Promise<string> {
  // #ifdef APP-PLUS
  return await new Promise((resolve) => {
    plus.runtime.getProperty(plus.runtime.appid, (info) => {
      resolve(info?.version || '1.0.0')
    })
  })
  // #endif

  // #ifndef APP-PLUS
  return '1.0.0'
  // #endif
}

/**
 * 获取缓存信息
 */
async function fetchCacheInfo() {
  try {
    const info = uni.getStorageInfoSync()
    const sizeKB = Number(info.currentSize || 0)
    const res: ICacheInfo = {
      sizeKB,
      sizeText: formatSize(sizeKB),
    }
    cacheInfo.value = res
    const cacheItem = menuList.find(item => item.action === 'clearCache')
    if (cacheItem) {
      cacheItem.value = res.sizeText
    }
  }
  catch (error) {
    console.error('获取缓存信息失败:', error)
  }
}

/**
 * 检查版本更新
 */
async function fetchVersionInfo() {
  try {
    const currentVersion = await getCurrentVersion()
    const res: IAppVersion = {
      currentVersion,
      hasUpdate: false,
    }
    versionInfo.value = res
    const versionItem = aboutList.find(item => item.action === 'checkVersion')
    if (versionItem) {
      versionItem.value = `v${res.currentVersion}`
    }
  }
  catch (error) {
    console.error('检查版本失败:', error)
  }
}

/**
 * 菜单点击
 */
function handleMenuClick(action: string) {
  switch (action) {
    case 'changeLanguage':
      languageSheetVisible.value = true
      break
    case 'clearCache':
      clearCacheSheetVisible.value = true
      break
    case 'privacy':
      uni.navigateTo({
        url: '/pages-fg/webview/index?url=https://example.com/privacy',
      })
      break
    case 'terms':
      uni.navigateTo({
        url: '/pages-fg/webview/index?url=https://example.com/terms',
      })
      break
    case 'about':
      aboutSheetVisible.value = true
      break
    case 'checkVersion':
      handleCheckVersion()
      break
  }
}

function handleLanguageSelect(payload: { index: number }) {
  const selected = languageOptions[payload.index]
  if (!selected)
    return

  currentLanguage.value = selected.value
  const languageItem = menuList.find(item => item.action === 'changeLanguage')
  if (languageItem) {
    languageItem.value = selected.label
  }
  uni.showToast({ title: '切换成功', icon: 'success' })
}

/**
 * 清除缓存
 */
async function handleClearCache() {
  try {
    uni.clearStorageSync()
    uni.showToast({
      title: '清除成功',
      icon: 'success',
    })
    fetchCacheInfo()
  }
  catch (error) {
    console.error('清除缓存失败:', error)
    uni.showToast({ title: '清除失败', icon: 'none' })
  }
}

/**
 * 检查版本更新
 */
async function handleCheckVersion() {
  try {
    const res = versionInfo.value || { currentVersion: await getCurrentVersion(), hasUpdate: false }
    if (res.hasUpdate) {
      updateSheetVisible.value = true
    }
    else {
      uni.showToast({
        title: `已是最新版本(v${res.currentVersion})`,
        icon: 'success',
      })
    }
  }
  catch (error) {
    console.error('检查版本失败:', error)
  }
}

function openDownloadUrl(url?: string) {
  if (!url)
    return

  // #ifdef H5
  window.open(url)
  // #endif
  // #ifndef H5
  uni.navigateTo({
    url: `/pages-fg/webview/index?url=${encodeURIComponent(url)}`,
  })
  // #endif
}

function handleUpdateConfirm() {
  openDownloadUrl(versionInfo.value?.downloadUrl)
  updateSheetVisible.value = false
}

onLoad(() => {
  fetchCacheInfo()
  fetchVersionInfo()
})
</script>

<template>
  <view class="settings-page">
    <view class="top-bg" />
    <view class="content">
      <view class="header-card">
        <view class="header-left">
          <view class="header-icon">
            <wd-icon name="setting" size="44rpx" color="#fff" />
          </view>
          <view class="header-text">
            <text class="header-title">通用设置</text>
          </view>
        </view>
        <wd-tag v-if="versionInfo" plain>
          v{{ versionInfo.currentVersion }}
        </wd-tag>
      </view>

      <wd-card type="rectangle" custom-class="card">
        <template #title>
          <view class="card-title">
            通用
          </view>
        </template>
        <wd-cell-group border>
          <wd-cell
            v-for="item in menuList"
            :key="item.title"
            :title="item.title"
            :icon="item.icon"
            :value="item.showValue ? item.value : ''"
            is-link
            clickable
            @click="handleMenuClick(item.action)"
          />
        </wd-cell-group>
      </wd-card>

      <wd-card type="rectangle" custom-class="card card-gap">
        <template #title>
          <view class="card-title">
            协议与政策
          </view>
        </template>
        <wd-cell-group border>
          <wd-cell
            v-for="item in legalList"
            :key="item.title"
            :title="item.title"
            :icon="item.icon"
            is-link
            clickable
            @click="handleMenuClick(item.action)"
          />
        </wd-cell-group>
      </wd-card>

      <wd-card type="rectangle" custom-class="card card-gap">
        <template #title>
          <view class="card-title">
            关于
          </view>
        </template>
        <wd-cell-group border>
          <wd-cell
            v-for="item in aboutList"
            :key="item.title"
            :title="item.title"
            :icon="item.icon"
            :value="item.showValue ? item.value : ''"
            is-link
            clickable
            @click="handleMenuClick(item.action)"
          />
        </wd-cell-group>
      </wd-card>
    </view>

    <bottom-sheet
      v-model="languageSheetVisible"
      title="语言切换"
      :actions="languageActions"
      cancel-text="取消"
      :show-confirm="false"
      @select="handleLanguageSelect"
    />

    <bottom-sheet
      v-model="clearCacheSheetVisible"
      title="清除缓存"
      :actions="clearCacheActions"
      cancel-text="取消"
      :show-confirm="false"
      @select="handleClearCache"
    />

    <bottom-sheet
      v-model="aboutSheetVisible"
      title="关于我们"
      :show-confirm="false"
      cancel-text="知道了"
    >
      <view class="about-content">
        <view class="about-line">
          这是一款基于 UniApp 开发的移动应用
        </view>
        <view class="about-line">
          版本号：v{{ versionInfo?.currentVersion || '1.0.0' }}
        </view>
        <view class="about-line">
          联系方式：support@example.com
        </view>
      </view>
    </bottom-sheet>

    <bottom-sheet
      v-model="updateSheetVisible"
      title="发现新版本"
      confirm-text="立即更新"
      cancel-text="取消"
      :confirm-disabled="!versionInfo?.downloadUrl"
      @confirm="handleUpdateConfirm"
    >
      <view class="about-content">
        <view class="about-line">
          最新版本：v{{ versionInfo?.latestVersion || '-' }}
        </view>
        <view class="about-line">
          当前版本：v{{ versionInfo?.currentVersion || '-' }}
        </view>
        <view v-if="versionInfo?.updateDesc" class="about-line about-desc">
          {{ versionInfo.updateDesc }}
        </view>
      </view>
    </bottom-sheet>
  </view>
</template>

<style lang="scss" scoped>
.settings-page {
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

.header-card {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16rpx;
  padding: 22rpx 18rpx 18rpx;
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  margin-bottom: 18rpx;
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

.card-title {
  font-size: 30rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.wd-card.card.is-rectangle {
  border-radius: 28rpx;
  overflow: hidden;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
}

:deep(.card .wd-card__title-content) {
  padding: 18rpx 18rpx 0;
}

:deep(.card .wd-card__content) {
  padding: 0;
}

.card-gap {
  margin-top: 18rpx;
}

.about-content {
  padding: 8rpx 0;
}

.about-line {
  font-size: 28rpx;
  color: var(--fg-text);
  line-height: 1.7;
  padding: 6rpx 0;
}

.about-desc {
  color: var(--fg-text-muted);
}
</style>
