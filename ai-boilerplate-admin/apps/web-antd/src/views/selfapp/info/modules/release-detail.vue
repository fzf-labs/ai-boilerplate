<script lang="ts" setup>
import type { SelfAppReleaseInfo } from '#/api/v1/self-app-release';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Tag } from 'ant-design-vue';

const releaseInfo = ref<SelfAppReleaseInfo>();

const getTitle = computed(() => {
  const info = releaseInfo.value;
  return info?.title ? `版本发布详情 - ${info.title}` : '版本发布详情';
});

const updateTypeMap: Record<number, { color: string; text: string }> = {
  1: { color: 'red', text: '强制更新' },
  2: { color: 'orange', text: '提示更新' },
  3: { color: 'blue', text: '静默更新' },
};

const grayStrategyMap: Record<number, { color: string; text: string }> = {
  1: { color: 'green', text: '全量发布' },
  2: { color: 'blue', text: '自定义设备' },
};

const formatGraySns = computed(() => {
  if (!releaseInfo.value?.graySns?.length) return '-';
  return releaseInfo.value.graySns.join(', ');
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      releaseInfo.value = undefined;
      return;
    }
    // 获取传入的数据
    const data = modalApi.getData<SelfAppReleaseInfo>();
    if (data) {
      releaseInfo.value = data;
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal :title="getTitle" :width="1200" class="release-detail-modal">
    <div v-if="releaseInfo" class="release-detail-content">
      <!-- 版本发布头部信息 -->
      <div
        class="release-header mb-6 rounded-lg bg-gradient-to-r from-blue-50 to-indigo-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div class="shrink-0">
            <!-- 版本图标 -->
            <div
              class="flex h-24 w-24 items-center justify-center rounded-2xl bg-gradient-to-br from-blue-400 to-indigo-500 text-3xl font-bold text-white shadow-lg"
            >
              🚀
            </div>
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <p class="text-lg font-medium text-gray-600">
                {{ releaseInfo.packageName }}
              </p>
            </div>
            <div class="flex items-center space-x-4">
              <Tag
                :color="releaseInfo.status === 1 ? 'success' : 'error'"
                class="rounded-full px-3 py-1"
              >
                {{ releaseInfo.status === 1 ? '启用' : '禁用' }}
              </Tag>
            </div>
          </div>
        </div>
      </div>

      <!-- 详细信息 -->
      <div class="space-y-6">
        <!-- 更新信息 -->
        <Card title="🔄 更新信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-emerald-50 p-4">
              <div class="mb-2">
                <span class="font-medium text-gray-600">更新标题</span>
              </div>
              <div class="text-lg font-semibold text-gray-800">
                {{ releaseInfo.title || '-' }}
              </div>
            </div>
            <div
              v-if="releaseInfo.changelog"
              class="rounded-lg border-l-4 border-emerald-400 bg-emerald-50 p-4"
            >
              <div class="mb-2">
                <span class="font-medium text-gray-600">更新日志</span>
              </div>
              <div
                class="max-h-48 overflow-y-auto whitespace-pre-wrap leading-relaxed text-gray-700"
              >
                {{ releaseInfo.changelog }}
              </div>
            </div>
          </div>
        </Card>

        <!-- 基本信息 -->
        <Card title="📋 基本信息" size="small">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <div class="rounded-lg bg-blue-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">包名</span>
                <span class="font-semibold text-gray-900">{{
                  releaseInfo.packageName || '-'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-green-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">发布渠道</span>
                <span class="text-gray-800">{{
                  releaseInfo.channel || '-'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-purple-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">版本号</span>
                <span class="font-mono text-gray-800">{{
                  releaseInfo.version || '-'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-orange-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">Build号</span>
                <span class="font-mono text-gray-800">{{
                  releaseInfo.buildNum || '-'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-pink-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">最低系统版本</span>
                <span class="text-gray-800">{{
                  releaseInfo.minOsVersion || '-'
                }}</span>
              </div>
            </div>
            <div class="rounded-lg bg-indigo-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">发布时间</span>
                <span class="text-gray-800">{{
                  formatDateTime(releaseInfo.publishTime || '') || '-'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 安装包信息 -->
        <Card title="📦 安装包信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-emerald-50 p-4">
              <div class="mb-2">
                <span class="font-medium text-gray-600">安装包地址</span>
              </div>
              <div>
                <a
                  v-if="releaseInfo.packageURL"
                  :href="releaseInfo.packageURL"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="break-all text-blue-600 hover:text-blue-800 hover:underline"
                >
                  {{ releaseInfo.packageURL }}
                </a>
                <span v-else class="text-gray-500">-</span>
              </div>
            </div>
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div class="rounded-lg bg-teal-50 p-4">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">安装包大小</span>
                  <span class="font-semibold text-gray-800">{{
                    releaseInfo.packageSize
                      ? `${releaseInfo.packageSize} MB`
                      : '-'
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-cyan-50 p-4">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">安装包MD5</span>
                  <span class="font-mono text-sm text-gray-800">{{
                    releaseInfo.packageMd5 || '-'
                  }}</span>
                </div>
              </div>
            </div>
          </div>
        </Card>

        <!-- 灰度策略 -->
        <Card title="🎯 发布策略" size="small">
          <div class="space-y-4">
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div class="rounded-lg bg-blue-50 p-4">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">更新类型</span>
                  <Tag
                    v-if="releaseInfo.updateType"
                    :color="updateTypeMap[releaseInfo.updateType]?.color"
                    class="rounded"
                  >
                    {{ updateTypeMap[releaseInfo.updateType]?.text }}
                  </Tag>
                  <span v-else class="text-gray-500">-</span>
                </div>
              </div>
              <div class="rounded-lg bg-slate-50 p-4">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">灰度策略</span>
                  <Tag
                    v-if="releaseInfo.grayStrategy"
                    :color="grayStrategyMap[releaseInfo.grayStrategy]?.color"
                    class="rounded"
                  >
                    {{ grayStrategyMap[releaseInfo.grayStrategy]?.text }}
                  </Tag>
                  <span v-else class="text-gray-500">-</span>
                </div>
              </div>
            </div>
            <div
              v-if="releaseInfo.grayStrategy === 2"
              class="rounded-lg bg-lime-50 p-4"
            >
              <div class="mb-2">
                <span class="font-medium text-gray-600">灰度设备</span>
              </div>
              <div class="max-h-32 overflow-y-auto">
                <span class="whitespace-pre-wrap text-gray-800">{{
                  formatGraySns
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 时间记录 -->
        <Card title="⏰ 时间记录" size="small">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <div class="rounded-lg bg-violet-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">创建时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(releaseInfo.createdAt || '') }}
              </div>
            </div>
            <div class="rounded-lg bg-rose-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">更新时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(releaseInfo.updatedAt || '') }}
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
.release-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.release-detail-content {
  padding: 1.5rem;
}

.release-header {
  border: 1px solid rgb(209 213 219 / 30%);
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%);
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

.grid > div {
  transition: all 0.2s ease;
}

.grid > div:hover {
  transform: translateX(4px);
}
</style>
