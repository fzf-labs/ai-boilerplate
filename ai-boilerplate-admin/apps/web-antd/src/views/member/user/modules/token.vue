<script lang="ts" setup>
import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { Copy } from '@vben/icons';

import { Button, Card, Input, message } from 'ant-design-vue';

import { generateParentTestToken } from '#/api/v1/user';

const token = ref('');
const loading = ref(false);

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      token.value = '';
      loading.value = false;
      return;
    }

    const data = modalApi.getData<{ id: string; nickname?: string }>();
    if (!data || !data.id) {
      return;
    }

    loading.value = true;
    try {
      const res = await generateParentTestToken({ body: { id: data.id } });
      token.value = res.token || '';
    } catch (error) {
      console.error('生成测试Token失败:', error);
      message.error('生成测试Token失败');
      modalApi.close();
    } finally {
      loading.value = false;
    }
  },
});

// 复制 Token
async function copyToken() {
  if (!token.value) {
    return;
  }

  try {
    await navigator.clipboard.writeText(token.value);
    message.success('Token已复制到剪贴板');
  } catch {
    // 降级方案：使用传统方式复制
    const textarea = document.createElement('textarea');
    textarea.value = token.value;
    textarea.style.position = 'fixed';
    textarea.style.opacity = '0';
    document.body.append(textarea);
    textarea.select();
    try {
      document.execCommand('copy');
      message.success('Token已复制到剪贴板');
    } catch {
      message.error('复制失败，请手动复制');
    }
    textarea.remove();
  }
}

defineExpose({ modalApi });
</script>

<template>
  <Modal title="生成测试Token" class="token-modal w-full max-w-2xl">
    <div v-if="loading" class="flex justify-center py-12">
      <div class="text-center">
        <div
          class="mb-4 inline-block h-12 w-12 animate-spin rounded-full border-4 border-blue-500 border-t-transparent"
        ></div>
        <p class="text-gray-600">正在生成测试Token...</p>
      </div>
    </div>

    <div v-else-if="token" class="token-content">
      <!-- Token 显示区域 -->
      <Card title="🔑 测试Token" size="small" class="token-card">
        <div class="space-y-4">
          <!-- Token 输入框 -->
          <div class="token-display">
            <Input.TextArea
              :value="token"
              :rows="6"
              readonly
              class="font-mono text-sm"
              placeholder="Token将显示在这里"
            />
          </div>

          <!-- 复制按钮 -->
          <div class="flex justify-center">
            <Button
              type="primary"
              size="large"
              @click="copyToken"
              class="copy-button"
            >
              <Copy class="mr-2 size-5" />
              复制Token
            </Button>
          </div>
        </div>
      </Card>
    </div>
  </Modal>
</template>

<style scoped>
.token-card {
  border-left: 4px solid #3b82f6;
}

.token-display :deep(.ant-input) {
  background-color: #f9fafb;
  border: 2px solid #e5e7eb;
  transition: all 0.3s;
}

.token-display :deep(.ant-input:hover) {
  background-color: #fff;
  border-color: #3b82f6;
}

.copy-button {
  min-width: 160px;
  height: 44px;
  font-size: 16px;
  font-weight: 600;
  border-radius: 8px;
  box-shadow:
    0 4px 6px -1px rgb(0 0 0 / 10%),
    0 2px 4px -1px rgb(0 0 0 / 6%);
  transition: all 0.3s;
}

.copy-button:hover {
  box-shadow:
    0 10px 15px -3px rgb(0 0 0 / 10%),
    0 4px 6px -2px rgb(0 0 0 / 5%);
  transform: translateY(-2px);
}

.copy-button:active {
  transform: translateY(0);
}

@media (max-width: 768px) {
  .token-content {
    padding: 0.5rem;
  }

  .copy-button {
    width: 100%;
  }
}
</style>
