<script lang="ts" setup>
import type { SysNotifyMessageInfo } from '#/api/v1/sys-notify-message';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card } from 'ant-design-vue';

const formData = ref<SysNotifyMessageInfo>();

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysNotifyMessageInfo>();
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
    title="站内信详情"
    class="w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div class="space-y-6 p-4">
      <!-- 基本信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-blue-600">
            <span class="mr-2">📋</span>
            基本信息
          </div>
        </template>
        <div class="rounded-lg bg-blue-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-blue-100 py-2"
            >
              <span class="font-medium text-blue-800">消息编号</span>
              <span class="text-blue-700">{{ formData?.id || '-' }}</span>
            </div>
            <div
              class="flex items-center justify-between border-b border-blue-100 py-2"
            >
              <span class="font-medium text-blue-800">消息类型</span>
              <span class="text-blue-700">{{ formData?.type || '-' }}</span>
            </div>
            <div class="flex items-center justify-between py-2">
              <span class="font-medium text-blue-800">消息主题</span>
              <span class="text-blue-700">{{ formData?.subject || '-' }}</span>
            </div>
          </div>
        </div>
      </Card>

      <!-- 消息内容 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-green-600">
            <span class="mr-2">💬</span>
            消息内容
          </div>
        </template>
        <div class="rounded-lg bg-green-50 p-4">
          <div class="py-2">
            <div class="mb-2">
              <span class="font-medium text-green-800">内容详情</span>
            </div>
            <div
              class="min-h-[60px] rounded bg-green-100 p-3 text-sm text-green-800"
            >
              {{ formData?.content || '-' }}
            </div>
          </div>
        </div>
      </Card>

      <!-- 接收人信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-purple-600">
            <span class="mr-2">👤</span>
            接收人信息
          </div>
        </template>
        <div class="rounded-lg bg-purple-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-purple-100 py-2"
            >
              <span class="font-medium text-purple-800">接收人</span>
              <span class="font-mono text-purple-700">{{
                formData?.receiver || '-'
              }}</span>
            </div>
            <div class="flex items-center justify-between py-2">
              <span class="font-medium text-purple-800">接收人名称</span>
              <span class="text-purple-700">{{
                formData?.receiverName || '-'
              }}</span>
            </div>
          </div>
        </div>
      </Card>

      <!-- 时间信息 -->
      <Card class="transition-shadow duration-300 hover:shadow-lg">
        <template #title>
          <div class="flex items-center text-orange-600">
            <span class="mr-2">⏰</span>
            时间信息
          </div>
        </template>
        <div class="rounded-lg bg-orange-50 p-4">
          <div class="grid grid-cols-1 gap-4">
            <div
              class="flex items-center justify-between border-b border-orange-100 py-2"
            >
              <span class="font-medium text-orange-800">发送时间</span>
              <span class="text-orange-700">{{
                formatDateTime(formData?.sendTime?.toString() || '') || '-'
              }}</span>
            </div>
            <div class="flex items-center justify-between py-2">
              <span class="font-medium text-orange-800">阅读时间</span>
              <span class="text-orange-700">{{
                formatDateTime(formData?.readTime?.toString() || '') || '-'
              }}</span>
            </div>
          </div>
        </div>
      </Card>

      <!-- 扩展信息 -->
      <Card
        v-if="formData?.extend"
        class="transition-shadow duration-300 hover:shadow-lg"
      >
        <template #title>
          <div class="flex items-center text-gray-600">
            <span class="mr-2">🔧</span>
            扩展信息
          </div>
        </template>
        <div class="rounded-lg bg-gray-50 p-4">
          <div class="py-2">
            <div class="mb-2">
              <span class="font-medium text-gray-800">扩展数据</span>
            </div>
            <div
              class="min-h-[60px] rounded bg-gray-100 p-3 text-sm text-gray-800"
            >
              {{ formData?.extend }}
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
  box-shadow: 0 10px 25px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}
</style>
