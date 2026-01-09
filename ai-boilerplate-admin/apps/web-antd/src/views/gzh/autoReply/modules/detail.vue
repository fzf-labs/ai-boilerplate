<script lang="ts" setup>
import type { WxGzhAutoReplyInfo } from '#/api/v1/wx-gzh-auto-reply';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card } from 'ant-design-vue';

import { getWxGzhAutoReplyInfo } from '#/api/v1/wx-gzh-auto-reply';
import { CommonStatusEnum } from '#/utils/constants';

// Constants for auto reply types
const AutoReplyType = {
  KEYWORD: 1, // 关键词回复
  MESSAGE: 2, // 收到消息回复
  SUBSCRIBE: 3, // 被关注回复
} as const;

const KeywordMatchType = {
  EXACT: 1, // 全匹配
  PARTIAL: 2, // 半匹配
} as const;

const ResponseMessageType = {
  TEXT: 'text',
  IMAGE: 'image',
  VIDEO: 'video',
  VOICE: 'voice',
} as const;

const autoReplyData = ref<WxGzhAutoReplyInfo>();

// Utility functions
const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
  } catch {
    const textArea = document.createElement('textarea');
    textArea.value = text;
    document.body.append(textArea);
    textArea.select();
    document.execCommand('copy');
    textArea.remove();
  }
};

const formatText = (text: string | undefined) => {
  return text || '未设置';
};

const getTitle = computed(() => {
  const data = autoReplyData.value;
  return data ? `自动回复详情 - ${getTypeInfo.value.text}` : '自动回复详情';
});

// 状态映射
const statusMap: Record<number, { color: string; icon: string; text: string }> =
  {
    [CommonStatusEnum.ENABLE]: {
      color: 'success',
      icon: '✅',
      text: '启用',
    },
    [CommonStatusEnum.DISABLE]: {
      color: 'error',
      icon: '❌',
      text: '禁用',
    },
  };

// 回复类型映射
const typeMap: Record<number, { color: string; icon: string; text: string }> = {
  [AutoReplyType.KEYWORD]: {
    color: 'blue',
    icon: '🔑',
    text: '关键词回复',
  },
  [AutoReplyType.MESSAGE]: {
    color: 'green',
    icon: '💬',
    text: '收到消息回复',
  },
  [AutoReplyType.SUBSCRIBE]: {
    color: 'orange',
    icon: '👋',
    text: '被关注回复',
  },
};

// 匹配类型映射
const matchTypeMap: Record<
  number,
  { color: string; icon: string; text: string }
> = {
  [KeywordMatchType.EXACT]: {
    color: 'purple',
    icon: '🎯',
    text: '全匹配',
  },
  [KeywordMatchType.PARTIAL]: {
    color: 'cyan',
    icon: '🔍',
    text: '半匹配',
  },
};

// 消息类型映射
const messageTypeMap: Record<
  string,
  { color: string; icon: string; text: string }
> = {
  [ResponseMessageType.TEXT]: {
    color: 'blue',
    icon: '📝',
    text: '文本消息',
  },
  [ResponseMessageType.IMAGE]: {
    color: 'green',
    icon: '🖼️',
    text: '图片消息',
  },
  [ResponseMessageType.VIDEO]: {
    color: 'red',
    icon: '🎥',
    text: '视频消息',
  },
  [ResponseMessageType.VOICE]: {
    color: 'orange',
    icon: '🎵',
    text: '音频消息',
  },
};

// 获取状态信息
const getStatusInfo = computed(() => {
  const status = autoReplyData.value?.status ?? CommonStatusEnum.DISABLE;
  return statusMap[status] || { color: 'default', icon: '❓', text: '未知' };
});

// 获取回复类型信息
const getTypeInfo = computed(() => {
  const type = autoReplyData.value?.type;
  if (type === undefined) return { color: 'default', icon: '❓', text: '未知' };
  return typeMap[type] || { color: 'default', icon: '❓', text: '未知' };
});

// 获取匹配类型信息
const getMatchTypeInfo = computed(() => {
  if (
    !autoReplyData.value ||
    autoReplyData.value.type !== AutoReplyType.KEYWORD
  ) {
    return { color: 'default', icon: '➖', text: '-' };
  }
  const matchType = autoReplyData.value.requestKeywordMatch;
  if (matchType === undefined) {
    return { color: 'default', icon: '❓', text: '未知' };
  }
  return (
    matchTypeMap[matchType] || { color: 'default', icon: '❓', text: '未知' }
  );
});

// 获取消息类型信息
const getMessageTypeInfo = computed(() => {
  const messageType = autoReplyData.value?.responseMessageType;
  if (!messageType) return { color: 'default', icon: '❓', text: '未知' };
  return (
    messageTypeMap[messageType] || {
      color: 'default',
      icon: '❓',
      text: messageType,
    }
  );
});

const [Modal, modalApi] = useVbenModal({
  showCancelButton: false,
  showConfirmButton: false,
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      autoReplyData.value = undefined;
      return;
    }

    const data = modalApi.getData<WxGzhAutoReplyInfo>();
    if (!data || !data.id) return;

    modalApi.lock();
    try {
      const res = await getWxGzhAutoReplyInfo({ params: { id: data.id } });
      autoReplyData.value = res.info;
    } catch (error) {
      console.error('加载自动回复详情失败:', error);
    } finally {
      modalApi.lock(false);
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal class="auto-reply-detail-modal w-full max-w-4xl" :title="getTitle">
    <div class="space-y-6">
      <div v-if="autoReplyData" class="space-y-6">
        <!-- 基本信息卡片 -->
        <Card
          class="animate-fade-in-up transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            animation-delay: 0.1s;
          "
        >
          <template #title>
            <div class="flex items-center text-blue-600">
              <span class="mr-2">📋</span>
              基本信息
            </div>
          </template>
          <div class="space-y-4">
            <div
              class="info-item rounded-lg bg-cyan-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">ID</span>
                <div class="flex items-center space-x-2">
                  <span class="font-mono font-semibold text-gray-900">
                    {{ formatText(autoReplyData.id?.toString()) }}
                  </span>
                  <button
                    v-if="autoReplyData.id"
                    @click="copyToClipboard(autoReplyData.id.toString())"
                    class="rounded-md bg-white/50 px-2 py-1 text-xs text-blue-600 transition-all duration-200 hover:scale-105 hover:bg-white"
                    title="复制ID"
                  ></button>
                </div>
              </div>
            </div>
            <div
              class="info-item rounded-lg bg-indigo-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">AppId</span>
                <div class="flex items-center space-x-2">
                  <span class="font-mono font-semibold text-gray-900">
                    {{ formatText(autoReplyData.appId) }}
                  </span>
                  <button
                    v-if="autoReplyData.appId"
                    @click="copyToClipboard(autoReplyData.appId)"
                    class="rounded-md bg-white/50 px-2 py-1 text-xs text-blue-600 transition-all duration-200 hover:scale-105 hover:bg-white"
                    title="复制 AppId"
                  ></button>
                </div>
              </div>
            </div>
            <div
              class="info-item rounded-lg bg-blue-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">回复类型</span>
                <span class="font-semibold text-gray-900">{{
                  getTypeInfo.text
                }}</span>
              </div>
            </div>
            <div
              class="info-item rounded-lg bg-sky-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">状态</span>
                <span class="font-semibold text-gray-900">{{
                  getStatusInfo.text
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 触发条件卡片 -->
        <Card
          class="animate-fade-in-up transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            animation-delay: 0.2s;
          "
        >
          <template #title>
            <div class="flex items-center text-green-600">
              <span class="mr-2">🎯</span>
              触发条件
            </div>
          </template>
          <div class="space-y-4">
            <div
              v-if="autoReplyData.type === AutoReplyType.KEYWORD"
              class="space-y-4"
            >
              <div
                class="info-item rounded-lg bg-green-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
              >
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">请求关键字</span>
                  <div class="flex items-center space-x-2">
                    <span class="font-mono font-semibold text-gray-900">
                      {{ formatText(autoReplyData.requestKeyword) }}
                    </span>
                    <button
                      v-if="autoReplyData.requestKeyword"
                      @click="copyToClipboard(autoReplyData.requestKeyword)"
                      class="rounded-md bg-white/50 px-2 py-1 text-xs text-green-600 transition-all duration-200 hover:scale-105 hover:bg-white"
                      title="复制关键字"
                    ></button>
                  </div>
                </div>
              </div>
              <div
                class="info-item rounded-lg bg-emerald-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
              >
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">匹配类型</span>
                  <span class="font-semibold text-gray-900">{{
                    getMatchTypeInfo.text
                  }}</span>
                </div>
              </div>
            </div>
            <div
              v-else
              class="info-item rounded-lg bg-teal-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-center">
                <div class="text-center">
                  <p class="font-semibold text-gray-900">
                    {{
                      autoReplyData.type === AutoReplyType.MESSAGE
                        ? '收到任意消息时触发'
                        : '用户关注时触发'
                    }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </Card>

        <!-- 回复内容卡片 -->
        <Card
          class="animate-fade-in-up transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            animation-delay: 0.3s;
          "
        >
          <template #title>
            <div class="flex items-center text-purple-600">
              <span class="mr-2">💬</span>
              回复内容
            </div>
          </template>
          <div class="space-y-4">
            <div
              class="info-item rounded-lg bg-purple-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">消息类型</span>
                <span class="font-semibold text-gray-900">{{
                  getMessageTypeInfo.text
                }}</span>
              </div>
            </div>
            <div
              v-if="
                autoReplyData.responseMessageType === ResponseMessageType.TEXT
              "
              class="space-y-4"
            >
              <div
                class="info-item rounded-lg bg-violet-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
              >
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-600">回复内容</span>
                    <button
                      v-if="autoReplyData.responseContent"
                      @click="copyToClipboard(autoReplyData.responseContent)"
                      class="rounded-md bg-white/50 px-2 py-1 text-xs text-purple-600 transition-all duration-200 hover:scale-105 hover:bg-white"
                      title="复制回复内容"
                    ></button>
                  </div>
                  <div
                    class="max-h-32 overflow-y-auto rounded-lg border border-purple-200 bg-white p-4 shadow-sm"
                  >
                    <pre
                      class="whitespace-pre-wrap text-sm leading-relaxed text-gray-800"
                      >{{ formatText(autoReplyData.responseContent) }}
                    </pre>
                  </div>
                </div>
              </div>
            </div>
            <div
              v-else
              class="info-item rounded-lg bg-indigo-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">媒体文件ID</span>
                <div class="flex items-center space-x-2">
                  <span class="font-mono font-semibold text-gray-900">
                    {{ formatText(autoReplyData.responseMediaId) }}
                  </span>
                  <button
                    v-if="autoReplyData.responseMediaId"
                    @click="copyToClipboard(autoReplyData.responseMediaId)"
                    class="rounded-md bg-white/50 px-2 py-1 text-xs text-purple-600 transition-all duration-200 hover:scale-105 hover:bg-white"
                    title="复制媒体文件ID"
                  >
                    📋
                  </button>
                </div>
              </div>
            </div>
          </div>
        </Card>

        <!-- 时间记录卡片 -->
        <Card
          class="animate-fade-in-up transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            animation-delay: 0.4s;
          "
        >
          <template #title>
            <div class="flex items-center text-gray-600">
              <span class="mr-2">⏰</span>
              时间记录
            </div>
          </template>
          <div class="space-y-4">
            <div
              class="info-item rounded-lg bg-orange-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-semibold text-gray-900">
                  {{
                    formatDateTime(autoReplyData.createdAt || '') || '未设置'
                  }}
                </span>
              </div>
            </div>
            <div
              class="info-item rounded-lg bg-yellow-50 p-4 transition-all duration-300 hover:scale-105 hover:shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-semibold text-gray-900">
                  {{
                    formatDateTime(autoReplyData.updatedAt || '') || '未设置'
                  }}
                </span>
              </div>
            </div>
          </div>
        </Card>
      </div>

      <!-- 加载状态 -->
      <div
        v-else
        class="animate-fade-in-up flex flex-col items-center justify-center py-16"
      >
        <div
          class="mb-6 h-16 w-16 animate-spin rounded-full border-4 border-blue-200 border-t-blue-500 shadow-lg"
        ></div>
        <div class="text-center">
          <h3 class="mb-2 text-xl font-semibold text-gray-700">加载中...</h3>
          <p class="text-gray-500">正在获取自动回复详情</p>
        </div>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
@keyframes fade-in-up {
  from {
    opacity: 0;
    transform: translateY(30px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Responsive design */
@media (max-width: 768px) {
  .space-y-6 > * + * {
    margin-top: 1rem;
  }

  .info-item {
    padding: 0.75rem;
  }

  .info-item:hover {
    transform: scale(1.02);
  }
}

@media (max-width: 640px) {
  .space-y-4 > * + * {
    margin-top: 0.75rem;
  }
}

/* 动画类 */
.animate-fade-in-up {
  opacity: 0;
  animation: fade-in-up 0.6s ease-out forwards;
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Info item hover effects */
.info-item {
  cursor: pointer;
  transition: all 0.3s ease;
}

.info-item:hover {
  box-shadow:
    0 10px 15px -3px rgb(0 0 0 / 10%),
    0 4px 6px -2px rgb(0 0 0 / 5%);
  transform: translateX(0.25rem) scale(1.05);
}

/* Card styling */
.ant-card {
  overflow: hidden;
  border: 1px solid #e5e7eb;
  border-radius: 16px;
  transition: all 0.3s ease;
}

.ant-card:hover {
  border-color: #d1d5db;
  box-shadow:
    0 20px 25px -5px rgb(0 0 0 / 10%),
    0 10px 10px -5px rgb(0 0 0 / 4%);
}

/* Button styling */
button {
  transition: all 0.2s ease;
}

button:hover {
  transform: scale(1.05);
}

button:active {
  transform: scale(0.95);
}

/* Tag styling */
:deep(.ant-tag) {
  padding: 4px 12px;
  font-weight: 500;
  border: none;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgb(0 0 0 / 5%);
  transition: all 0.2s ease;
}

:deep(.ant-tag):hover {
  box-shadow: 0 4px 8px rgb(0 0 0 / 10%);
  transform: translateY(-1px);
}

/* Typography */
.font-mono {
  font-family:
    ui-monospace, SFMono-Regular, 'SF Mono', Monaco, Consolas,
    'Liberation Mono', 'Courier New', monospace;
}

/* Loading animation enhancement */
.h-16.w-16 {
  filter: drop-shadow(0 4px 8px rgb(59 130 246 / 15%));
}

/* Smooth scrolling for content areas */
.overflow-y-auto {
  scrollbar-color: rgb(156 163 175 / 50%) transparent;
  scrollbar-width: thin;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 4px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: rgb(156 163 175 / 50%);
  border-radius: 2px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background-color: rgb(156 163 175 / 70%);
}

/* 动画定义 */
</style>
