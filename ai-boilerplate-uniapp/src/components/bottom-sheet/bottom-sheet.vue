<script lang="ts" setup>
export interface BottomSheetAction {
  name: string
  subname?: string
  color?: string
  disabled?: boolean
  loading?: boolean
}

const props = withDefaults(defineProps<{
  modelValue: boolean
  title?: string
  actions?: BottomSheetAction[]
  cancelText?: string
  confirmText?: string
  showCancel?: boolean
  showConfirm?: boolean
  confirmDisabled?: boolean
  confirmVariant?: 'primary' | 'danger'
  closeOnClickModal?: boolean
  closeOnSelect?: boolean
}>(), {
  title: '',
  actions: () => [],
  cancelText: '取消',
  confirmText: '完成',
  showCancel: true,
  showConfirm: true,
  confirmDisabled: false,
  confirmVariant: 'primary',
  closeOnClickModal: true,
  closeOnSelect: true,
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  'close': []
  'cancel': []
  'confirm': []
  'select': [payload: { index: number }]
}>()

const hasActions = computed(() => (props.actions?.length ?? 0) > 0)

function close() {
  emit('update:modelValue', false)
  emit('close')
}

function cancel() {
  emit('cancel')
  close()
}

function confirm() {
  if (props.confirmDisabled)
    return
  emit('confirm')
}

function select(index: number) {
  const item = props.actions[index]
  if (!item || item.disabled || item.loading)
    return
  emit('select', { index })
  if (props.closeOnSelect)
    close()
}
</script>

<template>
  <wd-popup
    :model-value="modelValue"
    position="bottom"
    :close-on-click-modal="closeOnClickModal"
    :safe-area-inset-bottom="false"
    custom-style="background: transparent;"
    @update:model-value="emit('update:modelValue', $event)"
  >
    <view class="bottom-sheet">
      <view class="bottom-sheet__header">
        <text class="bottom-sheet__title">{{ title }}</text>
        <wd-icon class="bottom-sheet__close" name="add" @click="close" />
      </view>

      <view class="bottom-sheet__content">
        <template v-if="hasActions">
          <button
            v-for="(action, index) in actions"
            :key="`${action.name}-${index}`"
            class="bottom-sheet__action"
            :class="action.disabled ? 'is-disabled' : ''"
            :style="action.color ? `color: ${action.color};` : ''"
            @click="select(index)"
          >
            <view class="bottom-sheet__action-name">
              {{ action.name }}
            </view>
            <view v-if="action.subname" class="bottom-sheet__action-subname">
              {{ action.subname }}
            </view>
          </button>
        </template>
        <template v-else>
          <view class="bottom-sheet__slot">
            <slot />
          </view>
        </template>
      </view>

      <template v-if="!hasActions">
        <button
          v-if="showConfirm"
          class="bottom-sheet__action bottom-sheet__confirm"
          :class="[
            confirmVariant === 'danger' ? 'bottom-sheet__confirm--danger' : 'bottom-sheet__confirm--primary',
            confirmDisabled ? 'is-disabled' : '',
          ]"
          @click="confirm"
        >
          <view class="bottom-sheet__action-name">
            {{ confirmText }}
          </view>
        </button>
      </template>

      <button v-if="showCancel" class="bottom-sheet__cancel" @click="cancel">
        {{ cancelText }}
      </button>
    </view>
  </wd-popup>
</template>

<style lang="scss" scoped>
.bottom-sheet {
  margin: 0 20rpx calc(var(--window-bottom, 0px) + 20rpx) 20rpx;
  position: relative;
  padding-top: 16rpx;
  border-radius: var(--fg-radius-card);
  background: var(--fg-surface-glass);
  overflow: hidden;
  box-shadow: var(--fg-shadow-card);
  border: 1px solid var(--fg-border);
  -webkit-backdrop-filter: blur(var(--fg-blur-strong));
  backdrop-filter: blur(var(--fg-blur-strong));
}

.bottom-sheet::before {
  content: '';
  position: absolute;
  top: 10rpx;
  left: 50%;
  width: 72rpx;
  height: 8rpx;
  border-radius: var(--fg-radius-pill);
  background: var(--fg-ink-04);
  transform: translateX(-50%);
}

@supports not (
  (
    backdrop-filter: blur(1px),
  )
) {
  .bottom-sheet {
    background: var(--fg-surface);
  }
}

.bottom-sheet__header {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 96rpx;
  padding-top: 6rpx;
  border-bottom: 1px solid var(--fg-border-weak);
}

.bottom-sheet__title {
  font-size: 30rpx;
  font-weight: 800;
  color: var(--fg-text);
}

.bottom-sheet__close {
  position: absolute;
  right: 18rpx;
  top: 50%;
  transform: translateY(-50%) rotate(-45deg);
  font-size: 36rpx;
  color: var(--fg-text-muted);
  line-height: 1.1;
}

.bottom-sheet__content {
  max-height: 56vh;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
}

.bottom-sheet__slot {
  padding: 18rpx 20rpx 12rpx;
}

.bottom-sheet :deep(.wd-input),
.bottom-sheet :deep(.wd-textarea) {
  background: var(--fg-bg-alt);
  border: 1px solid var(--fg-border);
  border-radius: 20rpx;
  padding: 6rpx 18rpx;
  box-shadow: inset 0 1rpx 0 rgba(255, 255, 255, 0.7);
  transition:
    border-color 0.2s ease,
    box-shadow 0.2s ease,
    background-color 0.2s ease;
}

.bottom-sheet :deep(.wd-input)::after,
.bottom-sheet :deep(.wd-textarea)::after {
  display: none;
}

.bottom-sheet :deep(.wd-input:focus-within),
.bottom-sheet :deep(.wd-textarea:focus-within) {
  background: var(--fg-surface);
  border-color: rgba(var(--fg-primary-rgb), 0.5);
  box-shadow:
    inset 0 1rpx 0 rgba(255, 255, 255, 0.8),
    0 0 0 3rpx rgba(var(--fg-primary-rgb), 0.14);
}

.bottom-sheet :deep(.wd-input__value),
.bottom-sheet :deep(.wd-textarea__value) {
  min-height: 88rpx;
}

.bottom-sheet :deep(.wd-input__inner) {
  height: 88rpx;
  font-size: 28rpx;
  color: var(--fg-text);
}

.bottom-sheet :deep(.wd-textarea__inner) {
  min-height: 160rpx;
  line-height: 1.5;
  font-size: 28rpx;
  color: var(--fg-text);
}

.bottom-sheet :deep(.wd-input__inner:focus),
.bottom-sheet :deep(.wd-textarea__inner:focus) {
  box-shadow: 0 0 0 2rpx rgba(var(--fg-primary-rgb), 0.2);
}

.bottom-sheet :deep(.wd-input__inner::-webkit-input-placeholder),
.bottom-sheet :deep(.wd-textarea__inner::-webkit-input-placeholder) {
  color: var(--fg-text-muted);
}

.bottom-sheet :deep(.wd-input__icon),
.bottom-sheet :deep(.wd-input__clear),
.bottom-sheet :deep(.wd-textarea__icon),
.bottom-sheet :deep(.wd-textarea__clear) {
  color: var(--fg-text-muted);
  background: transparent;
}

.bottom-sheet__action {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  min-height: 104rpx;
  padding: 18rpx 18rpx;
  border: none;
  background: var(--fg-surface);
  color: var(--fg-text);
  font-size: 30rpx;
  text-align: center;
  outline: none;
}

.bottom-sheet__action::after {
  display: none;
}

.bottom-sheet__action:active {
  background: var(--fg-bg-alt);
}

.bottom-sheet__action.is-disabled {
  color: var(--fg-text-weak);
}

.bottom-sheet__action.is-disabled:active {
  background: var(--fg-surface);
}

.bottom-sheet__action-name {
  display: inline-block;
  font-weight: 650;
}

.bottom-sheet__action-subname {
  display: inline-block;
  margin-left: 10rpx;
  font-size: 24rpx;
  color: var(--fg-text-weak);
  font-weight: 500;
}

.bottom-sheet__confirm {
  display: block;
  width: calc(100% - 48rpx);
  height: 96rpx;
  line-height: 96rpx;
  padding: 0;
  border-radius: var(--fg-radius-lg);
  margin: 18rpx auto 0;
}

.bottom-sheet__cancel {
  display: block;
  width: calc(100% - 48rpx);
  height: 96rpx;
  line-height: 96rpx;
  padding: 0;
  text-align: center;
  border-radius: var(--fg-radius-lg);
  border: none;
  background: var(--fg-surface-muted);
  outline: none;
  margin: 18rpx auto 24rpx;
  color: var(--fg-text);
  font-size: 30rpx;
  font-weight: 650;
}

.bottom-sheet__cancel::after {
  display: none;
}

.bottom-sheet__cancel:active {
  background: var(--fg-border-weak);
}

.bottom-sheet__confirm--primary {
  color: var(--fg-text-inverse);
  background: var(--fg-primary);
}

.bottom-sheet__confirm--primary:active {
  background: var(--fg-primary-600);
}

.bottom-sheet__confirm--danger {
  color: var(--fg-text-inverse);
  background: var(--fg-danger);
}

.bottom-sheet__confirm--danger:active {
  background: #e58a78;
}
</style>
