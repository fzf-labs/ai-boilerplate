<script lang="ts" setup>
import type { WxGzhUserInfo } from '#/api/v1/wx-gzh-user';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Spin, Tag } from 'ant-design-vue';

import { getWxGzhUserInfo } from '#/api/v1/wx-gzh-user';

const userInfo = ref<WxGzhUserInfo>();
const loading = ref(false);

const getTitle = computed(() => {
  const info = userInfo.value;
  return info
    ? `公众号粉丝详情 - ${info.nickname || info.openid}`
    : '公众号粉丝详情';
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      userInfo.value = undefined;
      return;
    }

    // 加载数据
    const data = modalApi.getData<WxGzhUserInfo>();
    if (!data?.id) {
      return;
    }

    modalApi.lock();
    try {
      loading.value = true;
      const res = await getWxGzhUserInfo({ params: { id: data.id } });
      userInfo.value = res.info;
    } finally {
      loading.value = false;
      modalApi.lock(false);
    }
  },
});

// 格式化关注状态
const formatSubscribeStatus = (status?: number) => {
  return status === 1 ? '已关注' : '未关注';
};

// 获取关注状态标签颜色
const getSubscribeStatusColor = (status?: number) => {
  return status === 1 ? 'success' : 'default';
};
</script>

<template>
  <Modal :title="getTitle" class="gzh-user-detail-modal w-full max-w-4xl">
    <div v-if="userInfo" class="user-detail-content">
      <Spin :spinning="loading">
        <div class="space-y-4">
          <!-- 身份信息 -->
          <Card title="🆔 身份信息" size="small">
            <div class="space-y-3">
              <div class="rounded-lg bg-blue-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">用户ID</span>
                  <span class="font-mono text-gray-800">{{ userInfo.id }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-green-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">公众号AppID</span>
                  <span class="font-mono text-gray-800">{{
                    userInfo.appId
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-purple-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">OpenID</span>
                  <span class="font-mono text-gray-800">{{
                    userInfo.openid
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-orange-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">UnionID</span>
                  <span class="font-mono text-gray-800">{{
                    userInfo.unionid || '无'
                  }}</span>
                </div>
              </div>
            </div>
          </Card>

          <!-- 用户信息 -->
          <Card title="👤 用户信息" size="small">
            <div class="space-y-3">
              <div class="rounded-lg bg-pink-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">昵称</span>
                  <span class="text-gray-800">{{
                    userInfo.nickname || '无'
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-indigo-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">语言</span>
                  <span class="text-gray-800">{{
                    userInfo.language || '无'
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-yellow-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">标签ID</span>
                  <span class="text-gray-800">{{
                    userInfo.tagIds || '无'
                  }}</span>
                </div>
              </div>
            </div>
          </Card>

          <!-- 关注信息 -->
          <Card title="💝 关注信息" size="small">
            <div class="space-y-3">
              <div class="rounded-lg bg-emerald-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">关注状态</span>
                  <Tag
                    :color="getSubscribeStatusColor(userInfo.subscribeStatus)"
                    class="rounded-full px-3 py-1"
                  >
                    {{ formatSubscribeStatus(userInfo.subscribeStatus) }}
                  </Tag>
                </div>
              </div>
              <div class="rounded-lg bg-lime-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">关注时间</span>
                  <span class="text-gray-800">
                    {{
                      userInfo.createdAt
                        ? formatDateTime(userInfo.createdAt)
                        : '-'
                    }}
                  </span>
                </div>
              </div>
            </div>
          </Card>

          <!-- 地理位置 -->
          <Card title="📍 地理位置" size="small">
            <div class="space-y-3">
              <div class="rounded-lg bg-red-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">国家</span>
                  <span class="text-gray-800">{{
                    userInfo.country || '无'
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-blue-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">省份</span>
                  <span class="text-gray-800">{{
                    userInfo.province || '无'
                  }}</span>
                </div>
              </div>
              <div class="rounded-lg bg-green-50 p-3">
                <div class="flex items-center justify-between">
                  <span class="font-medium text-gray-600">城市</span>
                  <span class="text-gray-800">{{ userInfo.city || '无' }}</span>
                </div>
              </div>
            </div>
          </Card>

          <!-- 备注信息 -->
          <Card v-if="userInfo.remark" title="📝 备注信息" size="small">
            <div class="rounded-lg border-l-4 border-amber-400 bg-amber-50 p-4">
              <p class="leading-relaxed text-gray-700">
                {{ userInfo.remark }}
              </p>
            </div>
          </Card>

          <!-- 时间信息 -->
          <Card title="⏰ 时间记录" size="small">
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div class="rounded-lg bg-cyan-50 p-4 text-center">
                <div class="mb-2 text-sm text-gray-600">创建时间</div>
                <div class="font-medium text-gray-800">
                  {{ formatDateTime(userInfo.createdAt || '') }}
                </div>
              </div>
              <div class="rounded-lg bg-emerald-50 p-4 text-center">
                <div class="mb-2 text-sm text-gray-600">更新时间</div>
                <div class="font-medium text-gray-800">
                  {{ formatDateTime(userInfo.updatedAt || '') }}
                </div>
              </div>
            </div>
          </Card>
        </div>
      </Spin>
    </div>
  </Modal>
</template>

<style scoped>
/* 自定义样式 */
</style>
