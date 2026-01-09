<script lang="ts" setup>
import type { SelfAppInfo } from '#/api/v1/self-app';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Tag } from 'ant-design-vue';

const appInfo = ref<SelfAppInfo>();

const getTitle = computed(() => {
  const info = appInfo.value;
  return info ? `应用详情 - ${info.name}` : '应用详情';
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      appInfo.value = undefined;
      return;
    }
    // 获取传入的数据
    const data = modalApi.getData<SelfAppInfo>();
    if (data) {
      appInfo.value = data;
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal :title="getTitle" class="selfapp-detail-modal w-full max-w-4xl">
    <div v-if="appInfo" class="selfapp-detail-content">
      <!-- 应用头部信息 -->
      <div
        class="app-header mb-6 rounded-lg bg-gradient-to-r from-emerald-50 to-teal-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div class="shrink-0">
            <!-- 应用图标 -->
            <div
              class="flex h-24 w-24 items-center justify-center rounded-2xl bg-gradient-to-br from-emerald-400 to-teal-500 text-3xl font-bold text-white shadow-lg"
            >
              📱
            </div>
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <h2 class="mb-1 text-2xl font-bold text-gray-800">
                {{ appInfo.name || '未设置应用名称' }}
              </h2>
              <p class="text-lg font-medium text-gray-600">
                {{ appInfo.packageName }}
              </p>
            </div>
            <div class="flex items-center space-x-4">
              <Tag
                :color="appInfo.status === 1 ? 'success' : 'error'"
                class="rounded-full px-3 py-1"
              >
                {{ appInfo.status === 1 ? '启用' : '禁用' }}
              </Tag>
            </div>
          </div>
        </div>
      </div>

      <!-- 详细信息 -->
      <div class="space-y-6">
        <!-- 基本信息 -->
        <Card title="📋 基本信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-blue-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">应用ID</span>
                <span class="font-mono text-gray-800">{{ appInfo.id }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-green-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">包名</span>
                <span class="font-semibold text-gray-900">{{
                  appInfo.packageName
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-purple-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">应用名称</span>
                <span class="text-gray-800">{{
                  appInfo.name || '未设置'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-pink-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">状态</span>
                <Tag
                  :color="appInfo.status === 1 ? 'success' : 'error'"
                  class="rounded"
                >
                  {{ appInfo.status === 1 ? '启用' : '禁用' }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- 应用描述 -->
        <Card
          v-if="appInfo.description"
          title="📝 应用描述"
          size="small"
          class="description-card"
        >
          <div
            class="rounded-lg border-l-4 border-emerald-400 bg-emerald-50 p-4"
          >
            <p class="whitespace-pre-line leading-relaxed text-gray-700">
              {{ appInfo.description }}
            </p>
          </div>
        </Card>

        <!-- 时间信息 -->
        <Card title="⏰ 时间记录" size="small">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <div class="rounded-lg bg-blue-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">创建时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(appInfo.createdAt || '') }}
              </div>
            </div>
            <div class="rounded-lg bg-green-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">更新时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(appInfo.updatedAt || '') }}
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
.selfapp-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.selfapp-detail-content {
  padding: 1.5rem;
}

.app-header {
  border: 1px solid rgb(209 213 219 / 30%);
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%);
}

.description-card {
  border-color: rgb(16 185 129 / 20%);
}

.description-card :deep(.ant-card-head) {
  border-bottom-color: rgb(16 185 129 / 20%);
}

/* 卡片悬停效果 */
:deep(.ant-card) {
  border-color: rgb(229 231 235);
  transition: all 0.3s ease;
}

:deep(.ant-card:hover) {
  box-shadow: 0 4px 12px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}

/* 信息项动画效果 */
.space-y-4 > div {
  transition: all 0.2s ease;
}

.space-y-4 > div:hover {
  transform: translateX(4px);
}
</style>
