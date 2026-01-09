<script lang="ts" setup>
import type { HelpCategoryInfo } from '@/api/v1/help-category/types'
import type { SubmitFeedbackReq } from '@/api/v1/help-feedback/types'
import { listHelpCategories } from '@/api/v1/help-category/helpCategory'
import { submitFeedback } from '@/api/v1/help-feedback/helpFeedback'

definePage({
  style: {
    navigationBarTitleText: '问题反馈',
  },
})

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
    uni.showToast({ title: '暂无可选分类', icon: 'none' })
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
    uni.showToast({
      title: '请选择问题分类',
      icon: 'none',
    })
    return
  }
  if (!formData.value.description.trim()) {
    uni.showToast({
      title: '请输入问题描述',
      icon: 'none',
    })
    return
  }

  try {
    submitting.value = true
    await submitFeedback({
      body: formData.value,
      options: {},
    })
    uni.showToast({
      title: '提交成功',
      icon: 'success',
    })
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  }
  catch (error) {
    console.error('提交反馈失败:', error)
    uni.showToast({
      title: '提交失败',
      icon: 'none',
    })
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
      <view class="sheet">
        <view class="sheet-header">
          <view class="sheet-title">
            问题反馈
          </view>
          <view class="sheet-subtitle">
            我们会尽快处理并改进体验
          </view>
        </view>

        <wd-cell-group border>
          <wd-cell title="问题分类" is-link clickable @click="openCategorySheet">
            <text :class="formData.category ? 'value-text' : 'value-placeholder'">
              {{ formData.category || '请选择' }}
            </text>
          </wd-cell>
        </wd-cell-group>

        <view class="section">
          <view class="section-title">
            问题描述
          </view>
          <wd-textarea
            v-model="formData.description"
            placeholder="请详细描述您遇到的问题"
            :maxlength="500"
            show-word-limit
          />
        </view>

        <view class="section">
          <view class="section-title">
            上传图片（选填，最多3张）
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

        <view class="section">
          <view class="section-title">
            联系方式（选填）
          </view>
          <wd-input
            v-model="formData.contact"
            placeholder="请输入您的手机号或邮箱"
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

.sheet {
  background: var(--fg-surface);
  border-radius: 28rpx;
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
}

.sheet-header {
  padding: 22rpx 18rpx 18rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.sheet-title {
  font-size: 40rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.sheet-subtitle {
  margin-top: 10rpx;
  font-size: 26rpx;
  color: var(--fg-text-muted);
}

.section {
  padding: 18rpx 18rpx 16rpx;
  border-top: 1px solid var(--fg-border-weak);
}

.section-title {
  font-size: 28rpx;
  font-weight: 700;
  color: var(--fg-text);
  margin-bottom: 12rpx;
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
  gap: 16rpx;
}

.image-item {
  position: relative;
  width: 200rpx;
  height: 200rpx;
}

.image {
  width: 100%;
  height: 100%;
  border-radius: 12rpx;
}

.image-remove {
  position: absolute;
  top: -8rpx;
  right: -8rpx;
  width: 48rpx;
  height: 48rpx;
  background: rgba(0, 0, 0, 0.6);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.image-upload {
  width: 200rpx;
  height: 200rpx;
  border: 2rpx dashed var(--fg-border);
  border-radius: 12rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--fg-bg-alt);
}

.sheet-footer {
  padding: 18rpx 18rpx 18rpx;
  border-top: 1px solid var(--fg-border-weak);
}
</style>
