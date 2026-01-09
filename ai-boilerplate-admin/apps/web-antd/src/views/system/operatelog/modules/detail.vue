<script lang="ts" setup>
import type { SysOperateLogInfo } from '#/api/v1/sys-operate-log';

import { h, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { CopyOutlined } from '@ant-design/icons-vue';
import { Button, Card, message } from 'ant-design-vue';

const formData = ref<SysOperateLogInfo>();

// 格式化JSON字符串
const formatJSON = (jsonStr: string | undefined) => {
  if (!jsonStr) return '';
  try {
    const obj = JSON.parse(jsonStr);
    return JSON.stringify(obj, null, 2);
  } catch {
    return jsonStr;
  }
};

// 复制到剪贴板
const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    message.success('复制成功');
  } catch {
    // 降级方案
    const textArea = document.createElement('textarea');
    textArea.value = text;
    document.body.append(textArea);
    textArea.select();
    document.execCommand('copy');
    textArea.remove();
    message.success('复制成功');
  }
};

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysOperateLogInfo>();
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
</script>

<template>
  <Modal
    title="操作日志详情"
    class="w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div class="space-y-6 p-4">
      <!-- 基本信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-blue-600">
            <span class="mr-2">ℹ️</span>
            基本信息
          </div>
        </template>
        <div class="rounded-lg bg-blue-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-blue-100 py-2"
            >
              <span class="font-medium text-blue-800">日志编号</span>
              <span class="text-blue-700">{{ formData?.id }}</span>
            </div>
            <div
              class="flex items-center justify-between border-b border-blue-100 py-2"
              v-if="formData?.traceId"
            >
              <span class="font-medium text-blue-800">链路编号</span>
              <span class="font-mono text-blue-700">
                {{ formData?.traceId }}
              </span>
            </div>
            <div class="flex items-center justify-between py-2">
              <span class="font-medium text-blue-800">操作时间</span>
              <span class="text-blue-700">
                {{ formatDateTime(formData?.createdAt?.toString() || '') }}
              </span>
            </div>
          </div>
        </div>
      </Card>

      <!-- 操作人信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-green-600">
            <span class="mr-2">👤</span>
            操作人信息
          </div>
        </template>
        <div class="rounded-lg bg-green-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-green-100 py-2"
            >
              <span class="font-medium text-green-800">操作人ID</span>
              <span class="text-green-700">{{ formData?.adminId }}</span>
            </div>
            <div
              class="flex items-center justify-between border-b border-green-100 py-2"
            >
              <span class="font-medium text-green-800">操作人昵称</span>
              <span class="text-green-700">{{ formData?.nickname }}</span>
            </div>
            <div
              class="flex items-center justify-between border-b border-green-100 py-2"
            >
              <span class="font-medium text-green-800">操作人IP</span>
              <span class="font-mono text-green-700">{{ formData?.IP }}</span>
            </div>
            <div class="flex items-start justify-between py-2">
              <span class="font-medium text-green-800">用户代理</span>
              <span class="max-w-md break-all text-right text-green-700">{{
                formData?.useragent
              }}</span>
            </div>
          </div>
        </div>
      </Card>

      <!-- 请求信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-orange-600">
            <span class="mr-2">📤</span>
            请求信息
          </div>
        </template>
        <div class="rounded-lg bg-orange-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-orange-100 py-2"
            >
              <span class="font-medium text-orange-800">请求URI</span>
              <span class="font-mono text-orange-700">{{ formData?.URI }}</span>
            </div>
            <div class="border-b border-orange-100 py-2">
              <div class="mb-2 flex items-center justify-between">
                <span class="font-medium text-orange-800">请求头</span>
                <Button
                  type="text"
                  size="small"
                  :icon="h(CopyOutlined)"
                  @click="copyToClipboard(formatJSON(formData?.header))"
                  class="text-orange-600 hover:text-orange-700"
                >
                  复制
                </Button>
              </div>
              <pre
                class="max-h-40 overflow-auto rounded bg-orange-100 p-3 text-sm text-orange-800"
                >{{ formatJSON(formData?.header) }}
              </pre>
            </div>
            <div class="py-2">
              <div class="mb-2 flex items-center justify-between">
                <span class="font-medium text-orange-800">请求参数</span>
                <Button
                  type="text"
                  size="small"
                  :icon="h(CopyOutlined)"
                  @click="copyToClipboard(formatJSON(formData?.req))"
                  class="text-orange-600 hover:text-orange-700"
                >
                  复制
                </Button>
              </div>
              <pre
                class="max-h-40 overflow-auto rounded bg-orange-100 p-3 text-sm text-orange-800"
                >{{ formatJSON(formData?.req) }}
              </pre>
            </div>
          </div>
        </div>
      </Card>

      <!-- 响应信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-purple-600">
            <span class="mr-2">📥</span>
            响应信息
          </div>
        </template>
        <div class="rounded-lg bg-purple-50 p-4">
          <div class="py-2">
            <div class="mb-2 flex items-center justify-between">
              <span class="font-medium text-purple-800">响应内容</span>
              <Button
                type="text"
                size="small"
                :icon="h(CopyOutlined)"
                @click="copyToClipboard(formatJSON(formData?.resp))"
                class="text-purple-600 hover:text-purple-700"
              >
                复制
              </Button>
            </div>
            <pre
              class="max-h-60 overflow-auto rounded bg-purple-100 p-3 text-sm text-purple-800"
              >{{ formatJSON(formData?.resp) }}
            </pre>
          </div>
        </div>
      </Card>
    </div>
  </Modal>
</template>

<style scoped>
/* 自定义滚动条样式 */
pre::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

pre::-webkit-scrollbar-track {
  background: rgb(0 0 0 / 10%);
  border-radius: 3px;
}

pre::-webkit-scrollbar-thumb {
  background: rgb(0 0 0 / 30%);
  border-radius: 3px;
}

pre::-webkit-scrollbar-thumb:hover {
  background: rgb(0 0 0 / 50%);
}

/* 卡片悬停效果增强 */
.ant-card {
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  transition: all 0.3s ease;
}

.ant-card:hover {
  box-shadow: 0 10px 25px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}
</style>
