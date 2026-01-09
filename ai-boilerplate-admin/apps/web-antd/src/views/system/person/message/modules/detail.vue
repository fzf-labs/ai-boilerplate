<script lang="ts" setup>
import type { SysNotifyMessageInfo } from '#/api/v1/sys-notify-message';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Tag } from 'ant-design-vue';

const formData = ref<SysNotifyMessageInfo>();

const getTitle = computed(() => {
  const data = formData.value;
  return data ? `消息详情 - ${data.subject || '无主题'}` : '消息详情';
});

// 消息类型映射
const messageTypeMap: Record<string, { color: string; icon: string }> = {
  SYSTEM: { color: 'blue', icon: '🔧' },
  USER: { color: 'green', icon: '👤' },
  NOTICE: { color: 'orange', icon: '📢' },
  WARNING: { color: 'red', icon: '⚠️' },
};

// 获取消息类型信息
const getMessageTypeInfo = computed(() => {
  const type = formData.value?.type || 'SYSTEM';
  return messageTypeMap[type] || messageTypeMap.SYSTEM;
});

// 判断是否已读
const isRead = computed(() => {
  return formData.value?.readTime && formData.value.readTime !== '';
});

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
  <Modal :title="getTitle" class="message-detail-modal w-full max-w-4xl">
    <div v-if="formData" class="message-detail-content">
      <!-- 详细信息 -->
      <div class="space-y-6">
        <!-- 基本信息 -->
        <Card title="📋 基本信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-blue-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">发送人</span>
                <span class="font-semibold text-gray-900">
                  {{ formData.senderName }}
                </span>
              </div>
            </div>
            <div class="rounded-lg bg-green-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">接收人</span>
                <span class="font-semibold text-gray-900">
                  {{ formData.receiverName }}
                </span>
              </div>
            </div>
            <div class="rounded-lg bg-purple-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">消息类型</span>
                <Tag :color="getMessageTypeInfo?.color" class="rounded">
                  {{ getMessageTypeInfo?.icon }} {{ formData?.type }}
                </Tag>
              </div>
            </div>
            <div class="rounded-lg bg-orange-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">已读状态</span>
                <Tag :color="isRead ? 'success' : 'warning'" class="rounded">
                  {{ isRead ? '✅ 已读' : '📬 未读' }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- 消息内容 -->
        <Card title="📄 消息详情" size="small" class="content-card">
          <div class="space-y-4">
            <!-- 消息主题 -->
            <div
              class="rounded-lg border-l-4 border-indigo-400 bg-indigo-50 p-4"
            >
              <div class="mb-2 flex items-center gap-2">
                <span class="text-sm font-medium text-indigo-600">📌 主题</span>
              </div>
              <h3 class="text-lg font-semibold leading-relaxed text-gray-800">
                {{ formData?.subject || '无主题' }}
              </h3>
            </div>
            <!-- 消息内容 -->
            <div class="rounded-lg border-l-4 border-blue-400 bg-blue-50 p-4">
              <div class="mb-2 flex items-center gap-2">
                <span class="text-sm font-medium text-blue-600">📝 内容</span>
              </div>
              <div class="prose prose-sm max-w-none">
                <p class="whitespace-pre-line leading-relaxed text-gray-700">
                  {{ formData?.content || '暂无内容' }}
                </p>
              </div>
            </div>
          </div>
        </Card>

        <!-- 扩展信息 -->
        <Card
          v-if="formData.extend"
          title="📎 扩展信息"
          size="small"
          class="extend-card"
        >
          <div class="rounded-lg border-l-4 border-purple-400 bg-purple-50 p-4">
            <pre class="whitespace-pre-line leading-relaxed text-gray-700">
              {{ formData.extend }}
            </pre>
          </div>
        </Card>

        <!-- 时间信息 -->
        <Card title="⏰ 时间记录" size="small">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <div class="rounded-lg bg-cyan-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">发送时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(formData.sendTime || '') }}
              </div>
            </div>
            <div v-if="isRead" class="rounded-lg bg-emerald-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">阅读时间</div>
              <div class="font-medium text-gray-800">
                {{ formatDateTime(formData.readTime || '') }}
              </div>
            </div>
            <div v-else class="rounded-lg bg-yellow-50 p-4 text-center">
              <div class="mb-2 text-sm text-gray-600">状态</div>
              <div class="font-medium text-orange-600">等待阅读</div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
/* 响应式优化 */
@media (max-width: 768px) {
  .message-header .flex {
    flex-direction: column;
    text-align: center;
  }

  .message-header .space-x-6 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;

    margin-top: 1.5rem;
    margin-right: calc(0px * var(--tw-space-x-reverse));
    margin-left: calc(0px * (1 - var(--tw-space-x-reverse)));
  }
}

.message-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.message-detail-content {
  padding: 1.5rem;
}

.message-header {
  border: 1px solid rgb(209 213 219 / 30%);
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%);
}

.content-card {
  border-color: rgb(59 130 246 / 20%);
}

.content-card :deep(.ant-card-head) {
  border-bottom-color: rgb(59 130 246 / 20%);
}

/* 消息主题样式 */
.content-card .border-indigo-400 {
  position: relative;
  overflow: hidden;
}

.content-card .border-indigo-400::before {
  position: absolute;
  top: 0;
  right: 0;
  left: 0;
  height: 2px;
  content: '';
  background: linear-gradient(90deg, rgb(99 102 241), rgb(139 92 246));
}

/* 消息内容样式 */
.content-card .border-blue-400 {
  position: relative;
  overflow: hidden;
}

.content-card .border-blue-400::before {
  position: absolute;
  top: 0;
  right: 0;
  left: 0;
  height: 2px;
  content: '';
  background: linear-gradient(90deg, rgb(59 130 246), rgb(37 99 235));
}

.extend-card {
  border-color: rgb(147 51 234 / 20%);
}

.extend-card :deep(.ant-card-head) {
  border-bottom-color: rgb(147 51 234 / 20%);
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

/* 渐变图标动画 */
.message-header .shrink-0 > div {
  transition: all 0.3s ease;
}

.message-header .shrink-0 > div:hover {
  transform: scale(1.05) rotate(5deg);
}

/* Tag 样式优化 */
:deep(.ant-tag) {
  font-weight: 500;
  border: none;
}
</style>
