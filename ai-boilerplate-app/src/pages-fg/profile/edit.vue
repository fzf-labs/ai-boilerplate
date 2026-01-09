<script lang="ts" setup>
import type { UpdateUserInfoReq, UserInfo } from '@/api/v1/user/types'
import { getUserInfo, updateUserInfo } from '@/api/v1/user/user'

definePage({
  style: {
    navigationBarTitleText: '个人信息',
  },
})

const loading = ref(false)
const saving = ref(false)

const formData = ref<UpdateUserInfoReq>({
  nickname: '',
  avatar: '',
  gender: 0,
  profile: '',
})

const originData = ref<UpdateUserInfoReq | null>(null)

const genderOptions = [
  { label: '保密', value: 0 },
  { label: '男', value: 1 },
  { label: '女', value: 2 },
]

const avatarSrc = computed(() => formData.value.avatar || '/static/images/default-avatar.png')
const displayName = computed(() => formData.value.nickname?.trim() || '未设置昵称')
const displayGender = computed(() => {
  const v = formData.value.gender ?? 0
  return genderOptions.find(i => i.value === v)?.label || '保密'
})
const displayProfile = computed(() => {
  const p = formData.value.profile?.trim() || ''
  if (!p)
    return '未填写'
  return p.length > 18 ? `${p.slice(0, 18)}…` : p
})

const hasProfile = computed(() => !!formData.value.profile?.trim())
const isDirty = computed(() => {
  if (!originData.value)
    return false

  const a = normalizeUserInfo(formData.value)
  const b = normalizeUserInfo(originData.value)
  return a !== b
})

type EditorField = 'nickname' | 'profile'
const editorVisible = ref(false)
const editorField = ref<EditorField>('nickname')
const editorTitle = ref('编辑昵称')
const editorValue = ref('')

const genderSheetVisible = ref(false)
const genderActions = computed(() => {
  const selected = formData.value.gender ?? 0
  return genderOptions.map((item) => {
    return {
      name: item.label,
      color: item.value === selected ? 'var(--fg-primary)' : '',
    }
  })
})

function normalizeUserInfo(data: UpdateUserInfoReq) {
  return JSON.stringify({
    nickname: (data.nickname || '').trim(),
    avatar: (data.avatar || '').trim(),
    gender: data.gender ?? 0,
    profile: (data.profile || '').trim(),
  })
}

function toReq(info?: UserInfo | null): UpdateUserInfoReq {
  return {
    nickname: info?.nickname || '',
    avatar: info?.avatar || '',
    gender: info?.gender ?? 0,
    profile: info?.profile || '',
  }
}

/**
 * 获取用户信息
 */
async function fetchUserProfile() {
  try {
    loading.value = true
    const res = await getUserInfo({ options: {} })
    const info = res.info || null
    if (!info) {
      uni.showToast({ title: '获取用户信息失败', icon: 'none' })
      return
    }
    const req = toReq(info)
    formData.value = { ...req }
    originData.value = { ...req }
  }
  catch (error) {
    console.error('获取用户信息失败:', error)
  }
  finally {
    loading.value = false
  }
}

/**
 * 选择头像
 */
function handleChooseAvatar() {
  uni.chooseImage({
    count: 1,
    sizeType: ['compressed'],
    sourceType: ['album', 'camera'],
    success: (res) => {
      const tempFilePath = res.tempFilePaths[0]
      // 这里应该上传到服务器，获取图片URL
      // 暂时使用本地路径
      formData.value.avatar = tempFilePath
      saveProfile()
    },
  })
}

async function saveProfile() {
  if (!isDirty.value)
    return
  if (saving.value)
    return

  try {
    if (formData.value.nickname && (formData.value.nickname.length < 2 || formData.value.nickname.length > 20)) {
      uni.showToast({ title: '昵称长度为2-20个字符', icon: 'none' })
      return
    }

    saving.value = true
    uni.showLoading({ title: '保存中' })
    await updateUserInfo({ body: formData.value })
    originData.value = { ...formData.value }
    uni.showToast({ title: '已保存', icon: 'success' })
  }
  catch (error) {
    console.error('保存失败:', error)
    uni.showToast({ title: '保存失败', icon: 'none' })
  }
  finally {
    saving.value = false
    uni.hideLoading()
  }
}

function openNicknameEditor() {
  editorField.value = 'nickname'
  editorTitle.value = '编辑昵称'
  editorValue.value = formData.value.nickname || ''
  editorVisible.value = true
}

function openProfileEditor() {
  editorField.value = 'profile'
  editorTitle.value = '编辑简介'
  editorValue.value = formData.value.profile || ''
  editorVisible.value = true
}

function closeEditor() {
  editorVisible.value = false
}

async function confirmEditor() {
  if (editorField.value === 'nickname') {
    const v = editorValue.value.trim()
    if (v && (v.length < 2 || v.length > 20)) {
      uni.showToast({ title: '昵称长度为2-20个字符', icon: 'none' })
      return
    }
    formData.value.nickname = v
    await saveProfile()
    closeEditor()
    return
  }

  formData.value.profile = editorValue.value.trim()
  await saveProfile()
  closeEditor()
}

function openGenderSheet() {
  genderSheetVisible.value = true
}

async function handleGenderSelect(payload: { index: number }) {
  const item = genderOptions[payload.index]
  formData.value.gender = item?.value ?? 0
  await saveProfile()
}

onLoad(() => {
  fetchUserProfile()
})
</script>

<template>
  <view class="profile-edit-page">
    <view class="top-bg" />

    <view class="content">
      <wd-skeleton
        theme="paragraph"
        :row-col="[1, 1, 1]"
        :loading="loading"
        animation="gradient"
      >
        <view class="sheet">
          <view class="sheet-header">
            <view class="avatar-wrap" @click="handleChooseAvatar">
              <image :src="avatarSrc" class="avatar" mode="aspectFill" />
              <view class="avatar-badge">
                <wd-icon name="edit" size="28rpx" color="#fff" />
              </view>
            </view>
            <view class="header-meta">
              <text class="header-title">{{ displayName }}</text>
              <text class="header-sub">点击下方项进行编辑</text>
            </view>
          </view>

          <wd-cell-group border>
            <wd-cell title="昵称" is-link clickable @click="openNicknameEditor">
              <text class="value-text">{{ displayName }}</text>
            </wd-cell>

            <wd-cell title="性别" is-link clickable @click="openGenderSheet">
              <text class="value-text">{{ displayGender }}</text>
            </wd-cell>

            <wd-cell title="简介" is-link clickable @click="openProfileEditor">
              <text :class="hasProfile ? 'value-text' : 'value-placeholder'">{{ displayProfile }}</text>
            </wd-cell>
          </wd-cell-group>
        </view>
      </wd-skeleton>
    </view>

    <bottom-sheet
      v-model="editorVisible"
      :title="editorTitle"
      confirm-text="完成"
      cancel-text="取消"
      :close-on-click-modal="true"
      :close-on-select="false"
      @confirm="confirmEditor"
      @cancel="closeEditor"
      @close="closeEditor"
    >
      <wd-input
        v-if="editorField === 'nickname'"
        v-model="editorValue"
        placeholder="请输入昵称"
        :maxlength="20"
        clearable
        confirm-type="done"
        @confirm="confirmEditor"
      />
      <wd-textarea
        v-else
        v-model="editorValue"
        placeholder="一句话介绍自己"
        :maxlength="200"
        show-word-limit
      />
    </bottom-sheet>

    <bottom-sheet
      v-model="genderSheetVisible"
      title="选择性别"
      :actions="genderActions"
      cancel-text="取消"
      :show-confirm="false"
      @select="handleGenderSelect"
    />
  </view>
</template>

<style lang="scss" scoped>
.profile-edit-page {
  min-height: 100vh;
  background: var(--fg-bg);
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
  align-items: center;
  gap: 18rpx;
  padding: 22rpx 18rpx 18rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.avatar-wrap {
  position: relative;
  width: 96rpx;
  height: 96rpx;
  border-radius: 48rpx;
  overflow: hidden;
  border: 3rpx solid rgba(var(--fg-primary-rgb), 0.18);
  box-shadow: 0 10rpx 26rpx rgba(var(--fg-primary-rgb), 0.16);

  &:active {
    opacity: 0.92;
  }
}

.avatar {
  width: 100%;
  height: 100%;
  border-radius: 48rpx;
}

.avatar-badge {
  position: absolute;
  right: 6rpx;
  bottom: 6rpx;
  width: 40rpx;
  height: 40rpx;
  border-radius: 20rpx;
  background: rgba(var(--fg-primary-rgb), 0.92);
  border: 1px solid rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-meta {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6rpx;
  min-width: 0;
}

.header-title {
  font-size: 34rpx;
  font-weight: 800;
  color: var(--fg-text);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.header-sub {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

:deep(.wd-cell) {
  padding-top: 22rpx;
  padding-bottom: 22rpx;
}

:deep(.sheet .wd-cell.is-border .wd-cell__wrapper::after) {
  display: none;
}

:deep(.sheet .wd-cell.is-border::after) {
  position: absolute;
  display: block;
  content: '';
  width: 100%;
  height: 1px;
  left: 0;
  top: 0;
  transform: scaleY(0.5);
  transform-origin: top;
  background: var(--fg-border-weak);
  pointer-events: none;
}

.value-text {
  color: var(--fg-text);
  font-weight: 600;
}

.value-placeholder {
  color: var(--fg-text-weak);
  font-weight: 500;
}
</style>
