<script lang="ts" setup>
import type { DeviceInfo } from '#/api/v1/device';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Avatar, Card, Tag } from 'ant-design-vue';

import { getDeviceInfo } from '#/api/v1/device';

const deviceData = ref<DeviceInfo>();

const getTitle = computed(() => {
  const data = deviceData.value;
  return data ? `设备详情 - ${data.name || data.sn}` : '设备详情';
});

const statusMap: Record<number, { color: string; icon: string; text: string }> =
  {
    1: { color: 'success', icon: '✅', text: '启用' },
    [-1]: { color: 'error', icon: '❌', text: '禁用' },
  };

const getDeviceStatus = computed(() => {
  const status = deviceData.value?.status ?? 0;
  return statusMap[status] || { color: 'default', icon: '❓', text: '未知' };
});

const getDeviceInitials = computed(() => {
  const data = deviceData.value;
  if (data?.name) {
    return data.name.slice(0, 2).toUpperCase();
  }
  if (data?.sn) {
    return data.sn.slice(0, 2).toUpperCase();
  }
  return 'DE';
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      deviceData.value = undefined;
      return;
    }
    const data = modalApi.getData<DeviceInfo>();
    if (!data || !data.sn) return;

    modalApi.lock();
    try {
      const res = await getDeviceInfo({ params: { sn: data.sn } });
      deviceData.value = res.info;
    } catch (error) {
      console.error('加载设备详情失败:', error);
    } finally {
      modalApi.lock(false);
    }
  },
});

/** 格式化存储大小 */
function formatSize(size: number) {
  return size ? `${size}MB` : '未设置';
}

defineExpose({ modalApi });
</script>

<template>
  <Modal
    :title="getTitle"
    class="device-detail-modal w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div class="device-detail-container-compact">
      <!-- Header Section -->
      <div
        class="device-header animate-fade-in-up mb-4 rounded-lg bg-gradient-to-r from-blue-50 via-purple-50 to-pink-50 p-4"
      >
        <div class="flex items-center gap-4">
          <Avatar
            :size="64"
            class="flex-shrink-0 shadow-lg"
            style="
              background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            "
          >
            <span class="text-xl font-bold text-white">{{
              getDeviceInitials
            }}</span>
          </Avatar>
          <div class="flex-1">
            <h2 class="mb-1 text-xl font-bold text-gray-800">
              {{ deviceData?.name || deviceData?.sn || '未命名设备' }}
            </h2>
            <div class="flex items-center gap-4">
              <Tag
                :color="getDeviceStatus.color"
                class="flex items-center gap-1 px-3 py-1 text-sm font-medium"
              >
                <span>{{ getDeviceStatus.icon }}</span>
                {{ getDeviceStatus.text }}
              </Tag>
            </div>
          </div>
        </div>
      </div>

      <!-- Content Cards -->
      <div class="grid grid-cols-1 gap-4">
        <!-- 设备标识 -->
        <Card
          title="🏷️ 设备标识"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #3b82f6;
            animation-delay: 0.1s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-blue-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">设备序列号</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.sn || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-sky-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">设备名称</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.name || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-cyan-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">设备品牌</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.brand || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-blue-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">设备型号</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.model || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 硬件规格 -->
        <Card
          title="💾 硬件规格"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #10b981;
            animation-delay: 0.2s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-green-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">CPU型号</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.CPU || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-emerald-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">RAM大小</span>
                <span class="font-semibold text-gray-900">{{
                  formatSize(deviceData?.RAMSize || 0)
                }}</span>
              </div>
            </div>
            <div class="info-item bg-teal-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">DDR大小</span>
                <span class="font-semibold text-gray-900">{{
                  formatSize(deviceData?.ddrSize || 0)
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 网络连接 -->
        <Card
          title="🌐 网络连接"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #f59e0b;
            animation-delay: 0.3s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-orange-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">IMEI</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.imei || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-yellow-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">MAC地址</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.mac || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-amber-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">入网型号</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.network || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 系统版本 -->
        <Card
          title="📱 系统版本"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #8b5cf6;
            animation-delay: 0.4s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-purple-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">APP版本</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.appVersion || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-violet-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">安卓版本</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.androidVersion || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 推送服务 -->
        <Card
          title="📢 推送服务"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #f59e0b;
            animation-delay: 0.5s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-orange-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">推送通道ID</span>
                <span class="font-semibold text-gray-900">{{
                  deviceData?.push?.channelID || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 设备激活 -->
        <Card
          title="🚀 设备激活"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #10b981;
            animation-delay: 0.6s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-green-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">激活时间</span>
                <span class="font-semibold text-gray-900">{{
                  formatDateTime(deviceData?.registryTime || '') || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 时间记录 -->
        <Card
          title="⏰ 时间记录"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #06b6d4;
            animation-delay: 0.7s;
          "
        >
          <div class="space-y-3">
            <div class="info-item bg-cyan-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-semibold text-gray-900">{{
                  formatDateTime(deviceData?.createdAt || '') || '未设置'
                }}</span>
              </div>
            </div>
            <div class="info-item bg-sky-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-semibold text-gray-900">{{
                  formatDateTime(deviceData?.updatedAt || '') || '未设置'
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 设备描述 -->
        <Card
          title="📝 设备描述"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #ef4444;
            animation-delay: 0.8s;
          "
        >
          <div class="info-item bg-rose-50">
            <div class="max-h-32 overflow-y-auto">
              <div class="whitespace-pre-wrap leading-relaxed text-gray-900">
                {{ deviceData?.desc || '暂无描述信息' }}
              </div>
            </div>
          </div>
        </Card>

        <!-- 安全信息 -->
        <Card
          title="🔐 安全信息"
          class="device-card animate-fade-in-up"
          style="
            background: linear-gradient(145deg, #fff, #f8fafc);
            border-left: 4px solid #84cc16;
            animation-delay: 0.9s;
          "
        >
          <div class="space-y-6">
            <div>
              <h4 class="mb-2 text-sm font-medium text-gray-700">设备证书</h4>
              <div class="info-item bg-lime-50">
                <div class="max-h-32 overflow-y-auto">
                  <pre
                    class="whitespace-pre-wrap break-all text-xs leading-relaxed text-gray-700"
                    >{{ deviceData?.certificate || '暂无证书信息' }} 
                  </pre>
                </div>
              </div>
            </div>
            <div>
              <h4 class="mb-2 text-sm font-medium text-gray-700">设备密钥</h4>
              <div class="info-item bg-green-50">
                <div class="max-h-32 overflow-y-auto">
                  <pre
                    class="whitespace-pre-wrap break-all text-xs leading-relaxed text-gray-700"
                    >{{ deviceData?.secureKey || '暂无密钥信息' }}
                  </pre>
                </div>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
@keyframes fade-in-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .device-detail-container,
  .device-detail-container-compact {
    padding: 0.75rem;
  }

  .device-header {
    padding: 0.75rem;
  }

  .device-header .flex {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }

  .info-item {
    padding: 0.5rem;
  }

  .info-item .flex {
    flex-direction: column;
    gap: 0.5rem;
    align-items: flex-start;
  }
}

@media (max-width: 480px) {
  .device-detail-container,
  .device-detail-container-compact {
    padding: 0.5rem;
  }

  .device-header {
    padding: 0.5rem;
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}

/* 容器样式 */
.device-detail-container {
  max-height: 80vh;
  padding: 1.5rem;
  overflow-y: auto;
}

/* 紧凑容器样式 */
.device-detail-container-compact {
  max-height: 75vh;
  padding: 1rem;
  overflow-y: auto;
}

/* 头部样式 */
.device-header {
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.device-header::before {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: '';
  background: linear-gradient(
    135deg,
    rgb(255 255 255 / 10%) 0%,
    rgb(255 255 255 / 5%) 100%
  );
}

/* 卡片样式 */
.device-card {
  overflow: hidden;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgb(0 0 0 / 6%);
  transition: all 0.3s ease;
}

.device-card:hover {
  box-shadow: 0 8px 25px rgb(0 0 0 / 12%);
  transform: translateY(-2px);
}

/* 信息项样式 */
.info-item {
  padding: 0.75rem;
  cursor: default;
  border-radius: 6px;
  transition: all 0.3s ease;
}

.info-item:hover {
  box-shadow:
    0 10px 15px -3px rgb(0 0 0 / 10%),
    0 4px 6px -2px rgb(0 0 0 / 5%);
  transform: translateX(0.25rem) scale(1.02);
}

/* 滚动条样式 */
.device-detail-container::-webkit-scrollbar,
.device-detail-container-compact::-webkit-scrollbar {
  width: 6px;
}

.device-detail-container::-webkit-scrollbar-track,
.device-detail-container-compact::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 3px;
}

.device-detail-container::-webkit-scrollbar-thumb,
.device-detail-container-compact::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.device-detail-container::-webkit-scrollbar-thumb:hover,
.device-detail-container-compact::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* 内容区域滚动条 */
.max-h-32::-webkit-scrollbar,
.max-h-40::-webkit-scrollbar {
  width: 4px;
}

.max-h-32::-webkit-scrollbar-track,
.max-h-40::-webkit-scrollbar-track {
  background: #f8fafc;
  border-radius: 2px;
}

.max-h-32::-webkit-scrollbar-thumb,
.max-h-40::-webkit-scrollbar-thumb {
  background: #e2e8f0;
  border-radius: 2px;
}

.max-h-32::-webkit-scrollbar-thumb:hover,
.max-h-40::-webkit-scrollbar-thumb:hover {
  background: #cbd5e1;
}

/* 标签样式增强 */
.ant-tag {
  font-weight: 500;
  border: none;
  border-radius: 6px;
}

/* 头像样式增强 */
.ant-avatar {
  border: 3px solid rgb(255 255 255 / 30%);
}

/* 卡片标题样式 */
:deep(.ant-card-head-title) {
  font-size: 1.1rem;
  font-weight: 600;
  color: #374151;
}

/* 预格式化文本样式 */
pre {
  font-family: 'SF Mono', Monaco, Inconsolata, 'Roboto Mono', monospace;
  line-height: 1.4;
}

/* 动画定义 */
</style>
