<script lang="ts" setup>
import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { Card, Image } from 'ant-design-vue';

interface GzhAccountDetail {
  id?: number;
  name?: string;
  account?: string;
  appId?: string;
  appSecret?: string;
  URL?: string;
  token?: string;
  encodingAesKey?: string;
  qrCodeURL?: string;
  remark?: string;
}

const formData = ref<GzhAccountDetail>();

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<GzhAccountDetail>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      formData.value = data;
    } finally {
      modalApi.lock(false);
    }
  },
});

// 复制到剪贴板
const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    // 复制成功，静默处理
  } catch {
    // 降级方案
    const textArea = document.createElement('textarea');
    textArea.value = text;
    document.body.append(textArea);
    textArea.select();
    document.execCommand('copy');
    textArea.remove();
    // 复制成功，静默处理
  }
};

// 格式化显示文本
const formatText = (text: string | undefined) => {
  return text || '-';
};

// 截断长文本
const truncateText = (text: string | undefined, maxLength: number = 30) => {
  if (!text) return '-';
  return text.length > maxLength ? `${text.slice(0, maxLength)}...` : text;
};
</script>

<template>
  <Modal
    title="公众号账户详情"
    class="gzh-account-detail-modal w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div class="space-y-6 p-4">
      <!-- 基本信息 -->
      <Card
        class="transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-blue-600">
            <span class="mr-2">📋</span>
            基本信息
          </div>
        </template>
        <div class="rounded-lg bg-blue-50 p-4">
          <div class="space-y-4">
            <div
              class="flex items-center justify-between border-b border-blue-100 pb-3"
            >
              <span class="font-medium text-blue-800">账户ID</span>
              <span class="text-blue-700">{{
                formatText(formData?.id?.toString())
              }}</span>
            </div>
            <div
              class="flex items-center justify-between border-b border-blue-100 pb-3"
            >
              <span class="font-medium text-blue-800">账户名称</span>
              <span class="text-blue-700">{{
                formatText(formData?.name)
              }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="font-medium text-blue-800">微信号</span>
              <div class="flex items-center space-x-2">
                <span class="font-mono text-blue-700">{{
                  formatText(formData?.account)
                }}</span>
                <button
                  v-if="formData?.account"
                  @click="copyToClipboard(formData.account)"
                  class="rounded px-2 py-1 text-xs text-blue-600 transition-colors hover:bg-blue-100"
                  title="复制微信号"
                >
                  📋
                </button>
              </div>
            </div>
          </div>
        </div>
      </Card>

      <!-- 开发者配置 -->
      <Card
        class="transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-green-600">
            <span class="mr-2">⚙️</span>
            开发者配置
          </div>
        </template>
        <div class="rounded-lg bg-green-50 p-4">
          <div class="space-y-4">
            <div
              class="flex items-center justify-between border-b border-green-100 pb-3"
            >
              <span class="font-medium text-green-800">开发者ID (AppID)</span>
              <div class="flex items-center space-x-2">
                <span class="font-mono text-green-700">{{
                  formatText(formData?.appId)
                }}</span>
                <button
                  v-if="formData?.appId"
                  @click="copyToClipboard(formData.appId)"
                  class="rounded px-2 py-1 text-xs text-green-600 transition-colors hover:bg-green-100"
                  title="复制AppID"
                >
                  📋
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="font-medium text-green-800">
                开发者密码 (AppSecret)
              </span>
              <div class="flex items-center space-x-2">
                <span class="font-mono text-green-700">{{
                  truncateText(formData?.appSecret, 20)
                }}</span>
                <button
                  v-if="formData?.appSecret"
                  @click="copyToClipboard(formData.appSecret)"
                  class="rounded px-2 py-1 text-xs text-green-600 transition-colors hover:bg-green-100"
                  title="复制AppSecret"
                >
                  📋
                </button>
              </div>
            </div>
          </div>
        </div>
      </Card>

      <!-- 服务器配置 -->
      <Card
        class="transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-purple-600">
            <span class="mr-2">🔗</span>
            服务器配置
          </div>
        </template>
        <div class="rounded-lg bg-purple-50 p-4">
          <div class="space-y-4">
            <div
              class="flex items-start justify-between border-b border-purple-100 pb-3"
            >
              <span class="font-medium text-purple-800">服务器地址 (URL)</span>
              <div class="flex max-w-md items-center space-x-2">
                <span class="break-all text-right font-mono text-purple-700">{{
                  formatText(formData?.URL)
                }}</span>
                <button
                  v-if="formData?.URL"
                  @click="copyToClipboard(formData.URL)"
                  class="flex-shrink-0 rounded px-2 py-1 text-xs text-purple-600 transition-colors hover:bg-purple-100"
                  title="复制URL"
                >
                  📋
                </button>
              </div>
            </div>
            <div
              class="flex items-center justify-between border-b border-purple-100 pb-3"
            >
              <span class="font-medium text-purple-800">令牌 (Token)</span>
              <div class="flex items-center space-x-2">
                <span class="font-mono text-purple-700">{{
                  truncateText(formData?.token, 15)
                }}</span>
                <button
                  v-if="formData?.token"
                  @click="copyToClipboard(formData.token)"
                  class="rounded px-2 py-1 text-xs text-purple-600 transition-colors hover:bg-purple-100"
                  title="复制Token"
                >
                  📋
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="font-medium text-purple-800">消息加解密密钥</span>
              <div class="flex items-center space-x-2">
                <span class="font-mono text-purple-700">{{
                  truncateText(formData?.encodingAesKey, 15)
                }}</span>
                <button
                  v-if="formData?.encodingAesKey"
                  @click="copyToClipboard(formData.encodingAesKey)"
                  class="rounded px-2 py-1 text-xs text-purple-600 transition-colors hover:bg-purple-100"
                  title="复制密钥"
                >
                  📋
                </button>
              </div>
            </div>
          </div>
        </div>
      </Card>

      <!-- 二维码信息 -->
      <Card
        v-if="formData?.qrCodeURL"
        class="transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-orange-600">
            <span class="mr-2">📱</span>
            二维码信息
          </div>
        </template>
        <div class="rounded-lg bg-orange-50 p-4">
          <div class="flex flex-col items-center space-y-4">
            <div class="text-center">
              <span class="font-medium text-orange-800">微信公众号二维码</span>
            </div>
            <div class="flex justify-center">
              <Image
                :src="formData?.qrCodeURL"
                :width="200"
                :height="200"
                class="rounded-lg border-2 border-orange-200"
                :preview="true"
                fallback="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPhfDwAChwGA60e6kgAAAABJRU5ErkJggg=="
              />
            </div>
            <div class="text-center">
              <button
                @click="copyToClipboard(formData?.qrCodeURL || '')"
                class="rounded-lg bg-orange-100 px-4 py-2 text-sm text-orange-700 transition-colors hover:bg-orange-200"
              >
                📋 复制二维码链接
              </button>
            </div>
          </div>
        </div>
      </Card>

      <!-- 备注信息 -->
      <Card
        v-if="formData?.remark"
        class="transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-gray-600">
            <span class="mr-2">📝</span>
            备注信息
          </div>
        </template>
        <div class="rounded-lg bg-gray-50 p-4">
          <div class="py-2">
            <div class="rounded-lg bg-gray-100 p-4 text-gray-800">
              {{ formData?.remark }}
            </div>
          </div>
        </div>
      </Card>
    </div>
  </Modal>
</template>

<style scoped>
/* 卡片悬停效果增强 */
.ant-card {
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  transition: all 0.3s ease;
}

.ant-card:hover {
  border-color: #d1d5db;
  box-shadow: 0 10px 25px rgb(0 0 0 / 10%);
}

/* 复制按钮悬停效果 */
button:hover {
  transform: scale(1.05);
}

/* 图片容器样式 */
.ant-image {
  overflow: hidden;
  border-radius: 8px;
}
</style>
