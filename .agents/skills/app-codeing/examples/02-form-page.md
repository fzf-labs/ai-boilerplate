# Form Page with Validation

This example shows how to create a form page with validation using wot-design-uni.

## Complete Example

```vue
<script lang="ts" setup>
import type { FormInstance, FormRules } from 'wot-design-uni/components/wd-form/types'
import { useToast } from 'wot-design-uni'
import { updateUserInfo } from '@/api/v1/user/user'

definePage({
  style: {
    navigationBarTitleText: 'Edit Profile',
  },
})

const toast = useToast()
const formRef = ref<FormInstance | null>(null)
const loading = ref(false)

// Form data
const form = reactive({
  nickname: '',
  phone: '',
  email: '',
  bio: '',
})

// Validation rules
const rules: FormRules = {
  nickname: [
    { required: true, message: 'Please enter nickname' },
    { min: 2, max: 20, message: '2-20 characters' },
  ],
  phone: [
    { required: true, message: 'Please enter phone' },
    { pattern: /^1[3-9]\d{9}$/, message: 'Invalid phone number' },
  ],
  email: [
    { pattern: /^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/, message: 'Invalid email' },
  ],
}

/**
 * Submit form
 */
async function handleSubmit() {
  // Validate form
  const validateRes = await formRef.value?.validate()
  if (validateRes && !validateRes.valid) {
    return
  }

  try {
    loading.value = true
    await updateUserInfo({
      body: {
        nickname: form.nickname,
        phone: form.phone,
        email: form.email,
        bio: form.bio,
      },
      options: {},
    })
    toast.success('Saved successfully')
    setTimeout(() => uni.navigateBack(), 1500)
  }
  catch (error) {
    console.error('Save failed:', error)
    toast.error('Save failed')
  }
  finally {
    loading.value = false
  }
}
</script>

<template>
  <view class="form-page">
    <wd-card type="rectangle" custom-class="form-card">
      <wd-form ref="formRef" :model="form" :rules="rules" error-type="toast">
        <wd-form-item label="Nickname" prop="nickname" required>
          <wd-input
            v-model="form.nickname"
            placeholder="Enter nickname"
            clearable
            :maxlength="20"
          />
        </wd-form-item>

        <wd-form-item label="Phone" prop="phone" required>
          <wd-input
            v-model="form.phone"
            placeholder="Enter phone number"
            type="number"
            clearable
            :maxlength="11"
          />
        </wd-form-item>

        <wd-form-item label="Email" prop="email">
          <wd-input
            v-model="form.email"
            placeholder="Enter email (optional)"
            clearable
          />
        </wd-form-item>

        <wd-form-item label="Bio" prop="bio">
          <wd-textarea
            v-model="form.bio"
            placeholder="Tell us about yourself"
            :maxlength="200"
            show-word-limit
          />
        </wd-form-item>
      </wd-form>
    </wd-card>

    <view class="actions">
      <wd-button
        type="primary"
        size="large"
        block
        :loading="loading"
        :disabled="loading"
        @click="handleSubmit"
      >
        Save
      </wd-button>
    </view>

    <wd-toast />
  </view>
</template>

<style lang="scss" scoped>
.form-page {
  min-height: 100vh;
  background: var(--fg-bg);
  padding: var(--fg-page-x);
}

:deep(.form-card.is-rectangle) {
  border-radius: 28rpx;
  background: var(--fg-surface);
  border: 1px solid var(--fg-border);
}

.actions {
  margin-top: 40rpx;
  padding: 0 20rpx;
}
</style>
```

## Key Points

### 1. Form Setup

```typescript
import type { FormInstance, FormRules } from 'wot-design-uni/components/wd-form/types'

const formRef = ref<FormInstance | null>(null)
const form = reactive({ /* fields */ })
const rules: FormRules = { /* validation rules */ }
```

### 2. Validation Rules

```typescript
const rules: FormRules = {
  fieldName: [
    { required: true, message: 'Required field' },
    { min: 2, max: 20, message: 'Length constraint' },
    { pattern: /regex/, message: 'Pattern mismatch' },
  ],
}
```

### 3. Form Validation

```typescript
const validateRes = await formRef.value?.validate()
if (validateRes && !validateRes.valid) {
  return // Validation failed
}
```

### 4. Error Display

```vue
<wd-form error-type="toast">
  <!-- error-type options: toast, message, none -->
</wd-form>
```
