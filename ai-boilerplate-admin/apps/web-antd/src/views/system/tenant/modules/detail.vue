<script lang="ts" setup>
import type { SysMenuInfo } from '#/api/v1/sys-menu';
import type { SysTenantInfo } from '#/api/v1/sys-tenant';

import { computed, ref } from 'vue';

import { useVbenModal, VbenTree } from '@vben/common-ui';
import { formatDateTime, handleTree } from '@vben/utils';

import { Card, Tag } from 'ant-design-vue';

import { getSysMenuList } from '#/api/v1/sys-menu';

const tenantData = ref<SysTenantInfo & { adminName?: string }>();
const menuTree = ref<SysMenuInfo[]>([]); // 菜单树
const menuLoading = ref(false); // 加载菜单列表

const getTitle = computed(() => {
  const data = tenantData.value;
  return data ? `租户详情 - ${data.name}` : '租户详情';
});

// 状态映射
const statusMap: Record<number, { color: string; icon: string; text: string }> =
  {
    1: { color: 'success', icon: '✅', text: '启用' },
    0: { color: 'error', icon: '❌', text: '禁用' },
  };

// 获取状态信息
const getStatusInfo = computed(() => {
  const status = tenantData.value?.status ?? 0;
  return statusMap[status] || statusMap[0];
});

// 检查是否已过期
const isExpired = computed(() => {
  if (!tenantData.value?.expireTime) return false;
  return new Date(tenantData.value.expireTime) < new Date();
});

/** 加载菜单树 */
async function loadMenuTree() {
  menuLoading.value = true;
  try {
    const res = await getSysMenuList({ options: {} });
    menuTree.value = handleTree(res.list || []) as SysMenuInfo[];
  } catch (error) {
    console.error('加载菜单树失败:', error);
    menuTree.value = [];
  } finally {
    menuLoading.value = false;
  }
}

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      tenantData.value = undefined;
      menuTree.value = [];
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysTenantInfo & { adminName?: string }>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      tenantData.value = data;
      // 并行加载菜单树，提高性能
      if (data.menuIds && data.menuIds.length > 0) {
        await loadMenuTree();
      }
    } catch (error) {
      console.error('加载租户详情失败:', error);
    } finally {
      modalApi.lock(false);
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal :title="getTitle" class="tenant-detail-modal w-full max-w-4xl">
    <div v-if="tenantData" class="tenant-detail-content">
      <!-- 详细信息 -->
      <div class="space-y-6">
        <!-- 基本信息 -->
        <Card title="📋 基本信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-blue-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">租户ID</span>
                <span class="font-mono text-gray-800">
                  {{ tenantData.id }}
                </span>
              </div>
            </div>
            <div class="rounded-lg bg-green-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">租户名称</span>
                <span class="font-semibold text-gray-900">
                  {{ tenantData.name }}
                </span>
              </div>
            </div>
            <div class="rounded-lg bg-cyan-50 p-4">
              <div class="flex flex-col space-y-2">
                <span class="font-medium text-gray-600">租户描述</span>
                <p class="whitespace-pre-line leading-relaxed text-gray-700">
                  {{ tenantData.remark || '暂无描述信息' }}
                </p>
              </div>
            </div>
          </div>
        </Card>

        <!-- 管理员信息 -->
        <Card title="👤 管理员信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-purple-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">管理员ID</span>
                <span class="font-mono text-gray-800">
                  {{ tenantData.adminId || '未分配' }}
                </span>
              </div>
            </div>
            <div
              v-if="tenantData.adminName"
              class="rounded-lg bg-indigo-50 p-4"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">管理员名称</span>
                <span class="font-semibold text-gray-900">
                  {{ tenantData.adminName }}
                </span>
              </div>
            </div>
            <div v-else class="rounded-lg bg-gray-50 p-4 text-center">
              <span class="text-gray-500">管理员信息未设置</span>
            </div>
          </div>
        </Card>

        <!-- 状态信息 -->
        <Card title="🔄 状态信息" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-orange-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">租户状态</span>
                <Tag :color="getStatusInfo?.color" class="rounded">
                  {{ getStatusInfo?.icon }} {{ getStatusInfo?.text }}
                </Tag>
              </div>
            </div>
            <div
              class="rounded-lg p-4"
              :class="[isExpired ? 'bg-red-50' : 'bg-blue-50']"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">过期时间</span>
                <div class="text-right">
                  <div class="font-semibold text-gray-900">
                    {{
                      tenantData.expireTime
                        ? formatDateTime(String(tenantData.expireTime))
                        : '永不过期'
                    }}
                  </div>
                  <div
                    v-if="tenantData.expireTime"
                    class="text-sm"
                    :class="[isExpired ? 'text-red-600' : 'text-green-600']"
                  >
                    {{ isExpired ? '已过期' : '正常使用中' }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Card>

        <!-- 菜单权限 -->
        <Card title="🔐 菜单权限" size="small" class="menus-card">
          <div class="rounded-lg border-l-4 border-indigo-400 bg-indigo-50 p-4">
            <div class="mb-4 flex items-center gap-2">
              <span class="text-sm font-medium text-indigo-600">权限配置</span>
            </div>
            <div v-if="menuLoading" class="py-8 text-center">
              <div class="inline-flex items-center gap-2 text-gray-500">
                <div
                  class="h-4 w-4 animate-spin rounded-full border-2 border-gray-300 border-t-blue-500"
                ></div>
                正在加载权限配置...
              </div>
            </div>
            <div
              v-else-if="tenantData.menuIds && tenantData.menuIds.length > 0"
            >
              <VbenTree
                class="max-h-[400px] overflow-y-auto"
                :loading="menuLoading"
                :tree-data="menuTree"
                :model-value="tenantData.menuIds"
                checkable
                :selectable="false"
                :disabled="true"
                value-field="id"
                label-field="name"
              />
            </div>
            <div v-else class="py-8 text-center text-gray-500">
              <div class="mb-2">📋</div>
              暂无权限配置
            </div>
          </div>
        </Card>

        <!-- 时间信息 -->
        <Card title="⏰ 时间记录" size="small">
          <div class="space-y-4">
            <div class="rounded-lg bg-cyan-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(tenantData.createdAt || '')) }}
                </span>
              </div>
            </div>
            <div class="rounded-lg bg-emerald-50 p-4">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(tenantData.updatedAt || '')) }}
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
.tenant-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.tenant-detail-content {
  padding: 1.5rem;
}

.menus-card {
  border-color: rgb(99 102 241 / 20%);
}

.menus-card :deep(.ant-card-head) {
  border-bottom-color: rgb(99 102 241 / 20%);
}

/* 菜单权限样式 */
.menus-card .border-indigo-400 {
  position: relative;
  overflow: hidden;
}

.menus-card .border-indigo-400::before {
  position: absolute;
  top: 0;
  right: 0;
  left: 0;
  height: 2px;
  content: '';
  background: linear-gradient(90deg, rgb(99 102 241), rgb(139 92 246));
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

/* Tag 样式优化 */
:deep(.ant-tag) {
  font-weight: 500;
  border: none;
}

/* 单列布局优化 */
.tenant-detail-content .space-y-6 > * {
  width: 100%;
}
</style>
