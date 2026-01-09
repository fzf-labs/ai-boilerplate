<script lang="ts" setup>
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Avatar, Card, Tag } from 'ant-design-vue';

import { getSysAdminInfo } from '#/api/v1/sys-admin';

const adminData = ref<SysAdminInfo>();

const getTitle = computed(() => {
  const data = adminData.value;
  return data ? `用户详情 - ${data.nickname || data.username}` : '用户详情';
});

// 性别映射
const sexMap: Record<number, { color: string; icon: string; text: string }> = {
  1: { color: 'blue', icon: '👨', text: '男' },
  2: { color: 'pink', icon: '👩', text: '女' },
};

// 状态映射
const statusMap: Record<number, { color: string; icon: string; text: string }> =
  {
    1: { color: 'success', icon: '✅', text: '启用' },
    [-1]: { color: 'error', icon: '❌', text: '禁用' },
  };

// 获取性别信息
const getSexInfo = computed(() => {
  const sex = adminData.value?.sex ?? 0;
  return sexMap[sex] || { color: 'default', icon: '❓', text: '未知' };
});

// 获取状态信息
const getStatusInfo = computed(() => {
  const status = adminData.value?.status ?? -1;
  return statusMap[status] || statusMap[-1];
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      adminData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysAdminInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getSysAdminInfo({ params: { id: data.id } });
      adminData.value = res.info;
    } catch (error) {
      console.error('加载用户详情失败:', error);
    } finally {
      modalApi.lock(false);
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal :title="getTitle" class="admin-detail-modal w-full max-w-4xl">
    <div v-if="adminData" class="admin-detail-content">
      <!-- 用户头部信息 -->
      <div
        class="user-header mb-6 rounded-lg bg-gradient-to-r from-blue-50 via-purple-50 to-pink-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div class="shrink-0">
            <Avatar
              :size="96"
              :src="adminData.avatar"
              class="shadow-lg ring-4 ring-white"
            >
              <template v-if="!adminData.avatar">
                {{
                  adminData.nickname?.charAt(0) ||
                  adminData.username?.charAt(0) ||
                  'U'
                }}
              </template>
            </Avatar>
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <h2 class="mb-2 text-2xl font-bold text-gray-800">
                {{ adminData.nickname || adminData.username }}
              </h2>
              <p class="mb-2 font-mono text-sm text-gray-600">
                @{{ adminData.username }}
              </p>
              <div class="flex items-center gap-3">
                <Tag :color="getStatusInfo?.color" class="rounded-full">
                  {{ getStatusInfo?.icon }} {{ getStatusInfo?.text }}
                </Tag>
                <Tag :color="getSexInfo?.color" class="rounded-full">
                  {{ getSexInfo?.icon }} {{ getSexInfo?.text }}
                </Tag>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <!-- 联系方式 -->
        <Card title="📞 联系方式" size="small" class="contact-card">
          <div class="space-y-4">
            <div class="info-item bg-orange-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">手机号码</span>
                <span class="font-semibold text-gray-900">
                  {{ adminData.mobile || '未设置' }}
                </span>
              </div>
            </div>
            <div class="info-item bg-teal-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">邮箱地址</span>
                <span class="font-semibold text-gray-900">
                  {{ adminData.email || '未设置' }}
                </span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 组织架构 -->
        <Card title="🏢 组织架构" size="small" class="org-card">
          <div class="space-y-4">
            <div class="info-item bg-indigo-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">角色</span>
                <Tag color="blue" class="rounded">
                  🎭 {{ adminData.roleName || '未分配' }}
                </Tag>
              </div>
            </div>
            <div class="info-item bg-emerald-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">部门</span>
                <Tag color="green" class="rounded">
                  🏛️ {{ adminData.deptName || '未分配' }}
                </Tag>
              </div>
            </div>
            <div class="info-item bg-yellow-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">岗位</span>
                <Tag color="orange" class="rounded">
                  💼 {{ adminData.postName || '未分配' }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- 用户状态 -->
        <Card title="🔄 用户状态" size="small" class="status-card">
          <div class="space-y-4">
            <div class="info-item bg-rose-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">账户状态</span>
                <Tag :color="getStatusInfo?.color" class="rounded">
                  {{ getStatusInfo?.icon }} {{ getStatusInfo?.text }}
                </Tag>
              </div>
            </div>
            <div class="info-item bg-sky-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">用户性别</span>
                <Tag :color="getSexInfo?.color" class="rounded">
                  {{ getSexInfo?.icon }} {{ getSexInfo?.text }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- 系统信息 -->
        <Card title="🔧 系统信息" size="small" class="system-card">
          <div class="space-y-4">
            <div class="info-item bg-blue-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">用户ID</span>
                <span class="font-mono text-gray-800">
                  {{ adminData.id }}
                </span>
              </div>
            </div>
            <div class="info-item bg-violet-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">角色ID</span>
                <span class="font-mono text-sm text-gray-800">
                  {{ adminData.roleId || '-' }}
                </span>
              </div>
            </div>
            <div class="info-item bg-lime-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">部门ID</span>
                <span class="font-mono text-sm text-gray-800">
                  {{ adminData.deptId || '-' }}
                </span>
              </div>
            </div>
            <div class="info-item bg-amber-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">岗位ID</span>
                <span class="font-mono text-sm text-gray-800">
                  {{ adminData.postId || '-' }}
                </span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 时间记录 -->
        <Card title="⏰ 时间记录" size="small" class="time-card">
          <div class="space-y-4">
            <div class="info-item bg-cyan-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(adminData.createdAt || '')) }}
                </span>
              </div>
            </div>
            <div class="info-item bg-emerald-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(adminData.updatedAt || '')) }}
                </span>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
/* 动画效果 */
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

/* 响应式优化 */
@media (max-width: 768px) {
  .user-header {
    padding: 1rem;
  }

  .user-header .flex {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }

  .admin-detail-content {
    padding: 1rem;
  }
}

.admin-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.admin-detail-content {
  padding: 1.5rem;
}

/* 用户头部渐变背景增强 */
.user-header {
  position: relative;
  overflow: hidden;
}

.user-header::before {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: '';
  background: linear-gradient(
    135deg,
    rgb(59 130 246 / 10%),
    rgb(147 51 234 / 10%),
    rgb(236 72 153 / 10%)
  );
}

/* 信息项样式 */
.info-item {
  padding: 1rem;
  cursor: pointer;
  border-radius: 0.5rem;
  transition: all 0.2s;
}

.info-item:hover {
  box-shadow:
    0 10px 15px -3px rgb(0 0 0 / 10%),
    0 4px 6px -2px rgb(0 0 0 / 5%);
  transform: translateX(0.25rem) scale(1.05);
}

.bg-blue-50:hover {
  background-color: rgb(239 246 255);
}

.bg-green-50:hover {
  background-color: rgb(240 253 244);
}

.bg-purple-50:hover {
  background-color: rgb(250 245 255);
}

.bg-cyan-50:hover {
  background-color: rgb(236 254 255);
}

.bg-orange-50:hover {
  background-color: rgb(255 247 237);
}

.bg-teal-50:hover {
  background-color: rgb(240 253 250);
}

.bg-indigo-50:hover {
  background-color: rgb(238 242 255);
}

.bg-emerald-50:hover {
  background-color: rgb(236 253 245);
}

.bg-yellow-50:hover {
  background-color: rgb(254 249 195);
}

.bg-rose-50:hover {
  background-color: rgb(255 241 242);
}

.bg-sky-50:hover {
  background-color: rgb(240 249 255);
}

.bg-violet-50:hover {
  background-color: rgb(245 243 255);
}

.bg-lime-50:hover {
  background-color: rgb(247 254 231);
}

.bg-amber-50:hover {
  background-color: rgb(255 251 235);
}

/* 卡片悬停效果 */
:deep(.ant-card) {
  background: linear-gradient(145deg, #fff, #f8fafc);
  border-color: rgb(229 231 235);
  transition: all 0.3s ease;
}

:deep(.ant-card:hover) {
  border-color: rgb(99 102 241 / 30%);
  box-shadow: 0 8px 25px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}

/* 卡片标题样式 */
:deep(.ant-card-head-title) {
  font-size: 1rem;
  font-weight: 600;
}

/* Tag 样式优化 */
:deep(.ant-tag) {
  font-weight: 500;
  border: none;
  box-shadow: 0 2px 4px rgb(0 0 0 / 10%);
}

/* Avatar 样式 */
:deep(.ant-avatar) {
  font-size: 2rem;
  font-weight: bold;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: 3px solid white;
}

/* 特定卡片样式 */
.info-card {
  border-left: 4px solid #3b82f6;
}

.contact-card {
  border-left: 4px solid #10b981;
}

.org-card {
  border-left: 4px solid #8b5cf6;
}

.status-card {
  border-left: 4px solid #f59e0b;
}

.system-card {
  border-left: 4px solid #6366f1;
}

.time-card {
  border-left: 4px solid #06b6d4;
}

.space-y-6 > * {
  animation: fade-in-up 0.3s ease-out;
}

.space-y-6 > *:nth-child(2) {
  animation-delay: 0.1s;
}

.space-y-6 > *:nth-child(3) {
  animation-delay: 0.2s;
}

.space-y-6 > *:nth-child(4) {
  animation-delay: 0.3s;
}

.space-y-6 > *:nth-child(5) {
  animation-delay: 0.4s;
}

.space-y-6 > *:nth-child(6) {
  animation-delay: 0.5s;
}
</style>
