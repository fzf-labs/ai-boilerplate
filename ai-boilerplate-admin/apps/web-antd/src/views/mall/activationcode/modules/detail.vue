<script lang="ts" setup>
import type { MallActivationCodeInfo } from '#/api/v1/mall-activation-code';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Tag } from 'ant-design-vue';

const formData = ref<MallActivationCodeInfo>();

const getTitle = computed(() => {
  return `激活码详情 - ${formData.value?.code || ''}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 获取数据
    const data = modalApi.getData<MallActivationCodeInfo>();
    if (data) {
      formData.value = data;
    }
  },
});

// 商品类型映射
const productTypeMap: Record<string, string> = {
  membership: '会员',
  service: '服务',
};

// 状态映射
const statusMap: Record<number, { color: string; icon: string; text: string }> =
  {
    [-2]: { color: 'red', icon: '💸', text: '已退款' },
    [-1]: { color: 'red', icon: '❌', text: '禁用' },
    0: { color: 'orange', icon: '📦', text: '库存' },
    1: { color: 'blue', icon: '🛒', text: '已售出' },
    2: { color: 'green', icon: '✅', text: '已激活' },
    3: { color: 'gray', icon: '⏰', text: '已过期' },
  };

// 平台映射
const platformMap: Record<string, string> = {
  taobao: '淘宝',
  jd: '京东',
  pdd: '拼多多',
  official: '官网',
  other: '其他',
};

const getStatusInfo = computed(() => {
  const status = formData.value?.status ?? 0;
  return (
    statusMap[status] || { color: 'default', icon: '❓', text: '未知状态' }
  );
});

defineExpose({ modalApi });
</script>

<template>
  <Modal
    :title="getTitle"
    class="activation-code-detail-modal w-full max-w-4xl"
  >
    <div v-if="formData" class="activation-code-detail-content">
      <!-- Header Section -->
      <div
        class="code-header mb-6 rounded-lg bg-gradient-to-r from-blue-50 via-purple-50 to-pink-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div
            class="flex h-24 w-24 shrink-0 items-center justify-center rounded-lg bg-white text-4xl shadow-lg ring-4 ring-white"
          >
            🎫
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <h2 class="mb-2 text-2xl font-bold text-gray-800">
                {{ formData.code }}
              </h2>
              <p class="mb-2 text-sm text-gray-600">
                {{
                  formData.productType
                    ? productTypeMap[formData.productType] ||
                      formData.productType
                    : ''
                }}
              </p>
              <div class="flex items-center gap-3">
                <Tag :color="getStatusInfo.color" class="rounded-full">
                  {{ getStatusInfo.icon }} {{ getStatusInfo.text }}
                </Tag>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <!-- 🏷️ 基础信息 -->
        <Card title="🏷️ 基础信息" size="small" class="basic-info-card">
          <div class="space-y-4">
            <div class="info-item bg-blue-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">激活码</span>
                <span class="font-semibold text-gray-900">{{
                  formData.code
                }}</span>
              </div>
            </div>
            <div class="info-item bg-sky-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">批次号</span>
                <span class="font-semibold text-gray-900">{{
                  formData.batchNo
                }}</span>
              </div>
            </div>
            <div class="info-item bg-indigo-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">商品类型</span>
                <Tag color="purple" class="rounded">
                  {{
                    formData.productType
                      ? productTypeMap[formData.productType] ||
                        formData.productType
                      : ''
                  }}
                </Tag>
              </div>
            </div>
            <div class="info-item bg-violet-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">商品ID</span>
                <span class="font-semibold text-gray-900">{{
                  formData.productId
                }}</span>
              </div>
            </div>
            <div class="info-item bg-purple-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">状态</span>
                <Tag :color="getStatusInfo.color" class="rounded">
                  {{ getStatusInfo.icon }} {{ getStatusInfo.text }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- ⏰ 时间信息 -->
        <Card title="⏰ 时间信息" size="small" class="time-card">
          <div class="space-y-4">
            <div class="info-item bg-cyan-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">有效期开始</span>
                <span class="font-semibold text-gray-900">
                  {{ formatDateTime(formData.validSt || '') }}
                </span>
              </div>
            </div>
            <div class="info-item bg-teal-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">有效期截止</span>
                <span class="font-semibold text-gray-900">
                  {{ formatDateTime(formData.validEd || '') }}
                </span>
              </div>
            </div>
            <div class="info-item bg-emerald-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">激活时间</span>
                <span class="font-semibold text-gray-900">
                  {{
                    formData.activatedAt
                      ? formatDateTime(formData.activatedAt)
                      : '未激活'
                  }}
                </span>
              </div>
            </div>
            <div class="info-item bg-lime-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">平台售出时间</span>
                <span class="font-semibold text-gray-900">
                  {{
                    formData.platformSoldAt
                      ? formatDateTime(formData.platformSoldAt)
                      : '未售出'
                  }}
                </span>
              </div>
            </div>
            <div class="info-item bg-green-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-semibold text-gray-900">
                  {{ formatDateTime(formData.createdAt || '') }}
                </span>
              </div>
            </div>
            <div class="info-item bg-teal-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-semibold text-gray-900">
                  {{ formatDateTime(formData.updatedAt || '') }}
                </span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 👤 用户信息 -->
        <Card
          v-if="formData.userId"
          title="👤 用户信息"
          size="small"
          class="contact-card"
        >
          <div class="space-y-4">
            <div class="info-item bg-orange-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">用户ID</span>
                <span class="font-semibold text-gray-900">{{
                  formData.userId
                }}</span>
              </div>
            </div>
            <div
              v-if="formData.userChange?.userMembershipChange"
              class="info-item bg-amber-50"
            >
              <div class="flex flex-col space-y-2">
                <span class="font-medium text-gray-600">用户权益变化</span>
                <div class="rounded bg-white p-3">
                  <div
                    v-if="formData.userChange.userMembershipChange.before"
                    class="mb-2"
                  >
                    <div class="text-xs text-gray-500">变更前:</div>
                    <div class="text-sm">
                      会员类型:
                      {{
                        formData.userChange.userMembershipChange.before
                          .membershipType
                      }}
                      <span
                        v-if="
                          formData.userChange.userMembershipChange.before
                            .expiredAt
                        "
                      >
                        | 到期时间:
                        {{
                          formatDateTime(
                            formData.userChange.userMembershipChange.before
                              .expiredAt,
                          )
                        }}
                      </span>
                    </div>
                  </div>
                  <div v-if="formData.userChange.userMembershipChange.after">
                    <div class="text-xs text-gray-500">变更后:</div>
                    <div class="text-sm">
                      会员类型:
                      {{
                        formData.userChange.userMembershipChange.after
                          .membershipType
                      }}
                      <span
                        v-if="
                          formData.userChange.userMembershipChange.after
                            .expiredAt
                        "
                      >
                        | 到期时间:
                        {{
                          formatDateTime(
                            formData.userChange.userMembershipChange.after
                              .expiredAt,
                          )
                        }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Card>

        <!-- 🛒 平台信息 -->
        <Card
          v-if="formData.platform"
          title="🛒 平台信息"
          size="small"
          class="org-card"
        >
          <div class="space-y-4">
            <div class="info-item bg-indigo-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">平台</span>
                <Tag color="blue" class="rounded">
                  {{ platformMap[formData.platform] || formData.platform }}
                </Tag>
              </div>
            </div>
            <div
              v-if="formData.platformOrderNo"
              class="info-item bg-emerald-50"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">平台订单号</span>
                <span class="font-semibold text-gray-900">{{
                  formData.platformOrderNo
                }}</span>
              </div>
            </div>
            <div v-if="formData.platformBuyerId" class="info-item bg-violet-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">平台买家ID</span>
                <span class="font-semibold text-gray-900">{{
                  formData.platformBuyerId
                }}</span>
              </div>
            </div>
            <div
              v-if="formData.platformBuyerName"
              class="info-item bg-purple-50"
            >
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">平台买家昵称</span>
                <span class="font-semibold text-gray-900">{{
                  formData.platformBuyerName
                }}</span>
              </div>
            </div>
          </div>
        </Card>

        <!-- 📝 备注信息 -->
        <Card
          v-if="formData.remark"
          title="📝 备注信息"
          size="small"
          class="system-card"
        >
          <div class="space-y-4">
            <div class="info-item bg-gray-50">
              <div class="flex flex-col space-y-2">
                <span class="font-medium text-gray-600">备注</span>
                <span class="text-sm text-gray-900">{{ formData.remark }}</span>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
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

.basic-info-card {
  border-left: 4px solid #3b82f6;
}

.contact-card {
  border-left: 4px solid #10b981;
}

.org-card {
  border-left: 4px solid #8b5cf6;
}

.time-card {
  border-left: 4px solid #06b6d4;
}

.system-card {
  border-left: 4px solid #6b7280;
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

.bg-blue-50:hover {
  background-color: rgb(239 246 255);
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

.bg-cyan-50:hover {
  background-color: rgb(236 254 255);
}

.bg-sky-50:hover {
  background-color: rgb(240 249 255);
}

.bg-green-50:hover {
  background-color: rgb(240 253 244);
}

.bg-purple-50:hover {
  background-color: rgb(250 245 255);
}

.bg-gray-50:hover {
  background-color: rgb(249 250 251);
}
</style>
