<script lang="ts" setup>
import type { HelpCategoryInfo } from '@/api/v1/help-category/types'
import type { SubmitFeedbackReq } from '@/api/v1/help-feedback/types'
import { useToast } from 'wot-design-uni'
import { listHelpCategories } from '@/api/v1/help-category/helpCategory'
import { submitFeedback } from '@/api/v1/help-feedback/helpFeedback'

definePage({
  style: {
    navigationBarTitleText: '问题反馈',
  },
})

const toast = useToast()

// 帮助分类列表
const categories = ref<HelpCategoryInfo[]>([])
const categoryOptions = computed(() =>
  categories.value
    .filter(c => !!c.name)
    .map(c => ({ label: c.name as string, value: c.name as string })),
)
// 表单数据
const formData = ref<SubmitFeedbackReq>({
  category: '',
  description: '',
  images: '',
  contact: '',
})
const categorySheetVisible = ref(false)

const categoryActions = computed(() => {
  const options = categoryOptions.value
  return options.map((item) => {
    return {
      name: item.label,
      color: item.value === formData.value.category ? 'var(--fg-primary)' : '',
    }
  })
})
// 图片列表
const imageList = ref<string[]>([])
// 提交中状态
const submitting = ref(false)

/**
 * 获取帮助分类
 */
async function fetchCategories() {
  try {
    const res = await listHelpCategories({ options: {} })
    categories.value = res.list || []
  }
  catch (error) {
    console.error('获取分类失败:', error)
  }
}

/**
 * 选择图片
 */
function chooseImage() {
  uni.chooseImage({
    count: 3 - imageList.value.length,
    success: (res) => {
      const paths = Array.isArray(res.tempFilePaths) ? res.tempFilePaths : [res.tempFilePaths]
      imageList.value = [...imageList.value, ...paths]
      formData.value.images = JSON.stringify(imageList.value)
    },
  })
}

/**
 * 删除图片
 */
function removeImage(index: number) {
  imageList.value.splice(index, 1)
  formData.value.images = JSON.stringify(imageList.value)
}

function openCategorySheet() {
  if (categoryOptions.value.length === 0) {
    toast.warning('暂无可选分类')
    return
  }
  categorySheetVisible.value = true
}

function handleCategorySelect(payload: { index: number }) {
  const selected = categoryOptions.value[payload.index]
  if (!selected)
    return
  formData.value.category = selected.value
}

/**
 * 提交反馈
 */
async function handleSubmit() {
  if (!formData.value.category) {
    toast.warning('请选择问题分类')
    return
  }
  if (!formData.value.description.trim()) {
    toast.warning('请输入问题描述')
    return
  }

  try {
    submitting.value = true
    await submitFeedback({
      body: formData.value,
      options: {},
    })
    toast.success('提交成功')
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  }
  catch (error) {
    console.error('提交反馈失败:', error)
    toast.error('提交失败')
  }
  finally {
    submitting.value = false
  }
}

onLoad(() => {
  fetchCategories()
})
</script>

<template>
  <view class="feedback-page">
    <view class="top-bg" />
    <view class="content">
      <!-- 优化后的头部区域 -->
      <view class="page-header">
        <view class="header-icon">
          <wd-icon name="edit" size="48rpx" color="var(--wot-color-primary)" />
        </view>
        <view class="header-info">
          <text class="header-title">问题反馈</text>
          <text class="header-subtitle">我们会尽快处理并改进体验</text>
        </view>
      </view>

      <view class="sheet">
        <!-- 优化后的表单区域 -->
        <view class="form-section">
          <view class="form-label">
            <wd-icon name="apps" size="32rpx" color="var(--wot-color-primary)" />
            <text class="label-text">问题分类</text>
            <text class="label-required">*</text>
          </view>
          <wd-cell-group border>
            <wd-cell is-link clickable @click="openCategorySheet">
              <text :class="formData.category ? 'value-text' : 'value-placeholder'">
                {{ formData.category || '请选择分类' }}
              </text>
            </wd-cell>
          </wd-cell-group>
        </view>

        <view class="form-section">
          <view class="form-label">
            <wd-icon name="edit-outline" size="32rpx" color="var(--wot-color-primary)" />
            <text class="label-text">问题描述</text>
            <text class="label-required">*</text>
          </view>
          <wd-textarea
            v-model="formData.description"
            placeholder="请详细描述您遇到的问题，我们会尽快为您解决"
            :maxlength="500"
            show-word-limit
          />
        </view>

        <view class="form-section">
          <view class="form-label">
            <wd-icon name="picture" size="32rpx" color="var(--wot-color-primary)" />
            <text class="label-text">上传图片</text>
            <text class="label-optional">（选填，最多3张）</text>
          </view>
          <view class="image-list">
            <view
              v-for="(image, index) in imageList"
              :key="index"
              class="image-item"
            >
              <image :src="image" class="image" mode="aspectFill" />
              <view class="image-remove" @click="removeImage(index)">
                <wd-icon name="close" size="32rpx" color="#fff" />
              </view>
            </view>
            <view
              v-if="imageList.length < 3"
              class="image-upload"
              @click="chooseImage"
            >
              <wd-icon name="add" size="64rpx" color="var(--fg-text-weak)" />
            </view>
          </view>
        </view>

        <view class="form-section">
          <view class="form-label">
            <wd-icon name="phone" size="32rpx" color="var(--wot-color-primary)" />
            <text class="label-text">联系方式</text>
            <text class="label-optional">（选填）</text>
          </view>
          <wd-input
            v-model="formData.contact"
            placeholder="请输入您的手机号或邮箱，方便我们联系您"
            clearable
          />
        </view>

        <view class="sheet-footer pb-safe">
          <wd-button
            :block="true"
            :round="true"
            size="large"
            type="primary"
            :loading="submitting"
            @click="handleSubmit"
          >
            提交反馈
          </wd-button>
        </view>
      </view>
    </view>

    <bottom-sheet
      v-model="categorySheetVisible"
      title="选择分类"
      :actions="categoryActions"
      cancel-text="取消"
      :show-confirm="false"
      @select="handleCategorySelect"
    />
    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.feedback-page {
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

.page-header {
  display: flex;
  align-items: center;
  gap: 20rpx;
  padding: 24rpx;
  background: var(--fg-surface);
  border-radius: 28rpx;
  border: 1px solid var(--fg-border);
  box-shadow: var(--fg-shadow-card);
  margin-bottom: 20rpx;
}

.header-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 50%;
  background: linear-gradient(
    135deg,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.1) 0%,
    rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.05) 100%
  );
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6rpx;
}

.header-title {
  font-size: 36rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.header-subtitle {
  font-size: 24rpx;
  color: var(--fg-text-muted);
}

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}

.form-section {
  padding: 24rpx 20rpx;
  border-bottom: 1px solid var(--fg-border-weak);

  &:last-of-type {
    border-bottom: none;
  }
}

.form-label {
  display: flex;
  align-items: center;
  gap: 10rpx;
  margin-bottom: 16rpx;
}

.label-text {
  font-size: 30rpx;
  font-weight: 700;
  color: var(--fg-text);
}

.label-required {
  color: #ff4d4f;
  font-size: 28rpx;
  font-weight: 600;
}

.label-optional {
  font-size: 24rpx;
  color: var(--fg-text-muted);
  font-weight: 400;
}

.value-text {
  color: var(--fg-text);
  font-weight: 600;
}

.value-placeholder {
  color: var(--fg-text-weak);
  font-weight: 500;
}

.image-list {
  display: flex;
  flex-wrap: wrap;
  gap: 20rpx;
}

.image-item {
  position: relative;
  width: 200rpx;
  height: 200rpx;
  border-radius: 16rpx;
  overflow: hidden;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.08);
}

.image {
  width: 100%;
  height: 100%;
  border-radius: 16rpx;
}

.image-remove {
  position: absolute;
  top: -10rpx;
  right: -10rpx;
  width: 52rpx;
  height: 52rpx;
  background: linear-gradient(135deg, #ff4d4f 0%, #ff7875 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4rpx 12rpx rgba(255, 77, 79, 0.3);
  transition: all 0.2s ease;

  &:active {
    transform: scale(0.9);
  }
}

.image-upload {
  width: 200rpx;
  height: 200rpx;
  border: 3rpx dashed var(--fg-border);
  border-radius: 16rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--fg-bg-alt);
  transition: all 0.3s ease;

  &:active {
    background: rgba(var(--wot-color-primary-rgb, 0, 122, 255), 0.05);
    border-color: var(--wot-color-primary);
  }
}

.sheet-footer {
  padding: 24rpx 20rpx;
  background: var(--fg-bg-alt);
}
</style>
