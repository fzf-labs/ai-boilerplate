<script lang="ts" setup>
import type { UserInfo, UserMembershipInfo } from '#/api/v1/user';
import type { WxGzhUserInfo } from '#/api/v1/wx-gzh-user';
import type { WxXcxUserInfo } from '#/api/v1/wx-xcx-user';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Card, Image, Result, Spin, TabPane, Tabs, Tag } from 'ant-design-vue';

import { getUserInfo } from '#/api/v1/user';
import { getWxGzhUserInfo } from '#/api/v1/wx-gzh-user';
import { getWxXcxUserInfo } from '#/api/v1/wx-xcx-user';

const userInfo = ref<UserInfo>();
const gzhUserInfo = ref<WxGzhUserInfo>();
const xcxUserInfo = ref<WxXcxUserInfo>();
const membershipInfo = ref<null | UserMembershipInfo>();

const activeTab = ref('user');
const gzhLoading = ref(false);
const xcxLoading = ref(false);
const gzhError = ref(false);
const xcxError = ref(false);

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      userInfo.value = undefined;
      gzhUserInfo.value = undefined;
      xcxUserInfo.value = undefined;
      membershipInfo.value = undefined;
      activeTab.value = 'user';
      gzhError.value = false;
      xcxError.value = false;
      return;
    }

    const data = modalApi.getData<UserInfo>();
    if (!data || !data.id) {
      return;
    }

    modalApi.lock();
    try {
      const res = await getUserInfo({ params: { id: data.id } });
      userInfo.value = res.info;
      // 从用户信息中直接获取会员信息
      membershipInfo.value = res.info?.userMembershipInfo || null;
    } finally {
      modalApi.lock(false);
    }
  },
});

// 加载公众号用户信息
const loadGzhUserInfo = async () => {
  if (!userInfo.value?.wxGzhUserId || gzhUserInfo.value) {
    return;
  }

  gzhLoading.value = true;
  gzhError.value = false;
  try {
    const res = await getWxGzhUserInfo({ params: { id: userInfo.value.wxGzhUserId } });
    gzhUserInfo.value = res.info;
  } catch (error) {
    console.error('Failed to load WeChat Official Account user info:', error);
    gzhError.value = true;
  } finally {
    gzhLoading.value = false;
  }
};

// 加载小程序用户信息
const loadXcxUserInfo = async () => {
  if (!userInfo.value?.wxGzhXcxId || xcxUserInfo.value) {
    return;
  }

  xcxLoading.value = true;
  xcxError.value = false;
  try {
    const res = await getWxXcxUserInfo({ params: { id: userInfo.value.wxGzhXcxId } });
    xcxUserInfo.value = res.info;
  } catch (error) {
    console.error('Failed to load WeChat Mini Program user info:', error);
    xcxError.value = true;
  } finally {
    xcxLoading.value = false;
  }
};

// 处理 tab 切换
const handleTabChange = (key: number | string) => {
  const tabKey = String(key);
  activeTab.value = tabKey;
  switch (tabKey) {
    case 'gzh': {
      loadGzhUserInfo();
      break;
    }
    case 'membership': {
      // 会员信息已在用户信息加载时获取，无需额外操作
      break;
    }
    case 'xcx': {
      loadXcxUserInfo();
      break;
    }
  }
};

const genderText = computed(() => {
  switch (userInfo.value?.gender) {
    case 1: {
      return '男';
    }
    case 2: {
      return '女';
    }
    default: {
      return '未知';
    }
  }
});

const statusText = computed(() => {
  return userInfo.value?.status === 1 ? '启用' : '禁用';
});

// 判断是否有相关联的微信信息
const hasGzhUser = computed(() => !!userInfo.value?.wxGzhUserId);
const hasXcxUser = computed(() => !!userInfo.value?.wxGzhXcxId);

// 判断是否有会员信息
const hasMembership = computed(() => !!membershipInfo.value?.id);

// 会员信息相关计算属性
const membershipTypeText = computed(() => {
  switch (membershipInfo.value?.membershipType) {
    case 'normal': {
      return '普通会员';
    }
    case 'svip': {
      return 'SVIP会员';
    }
    case 'vip': {
      return 'VIP会员';
    }
    default: {
      return '未知类型';
    }
  }
});

const membershipStatusText = computed(() => {
  return membershipInfo.value?.status === 1 ? '正常' : '禁用';
});

const autoRenewText = computed(() => {
  return membershipInfo.value?.autoRenew === 1 ? '是' : '否';
});

const autoRenewDaysText = computed(() => {
  if (!membershipInfo.value?.autoRenewDays) {
    return '无';
  }
  return `${membershipInfo.value.autoRenewDays}天`;
});

const membershipExpiredText = computed(() => {
  if (!membershipInfo.value?.expiredAt) {
    return '永不过期';
  }
  const expiredDate = new Date(membershipInfo.value.expiredAt);
  const now = new Date();
  const isExpired = expiredDate < now;
  return {
    text: formatDateTime(membershipInfo.value.expiredAt),
    isExpired,
  };
});
</script>

<template>
  <Modal title="用户详情" class="user-detail-modal w-full max-w-4xl">
    <div v-if="userInfo" class="user-detail-content">
      <!-- 用户头部信息 -->
      <div
        class="user-header mb-6 rounded-lg bg-gradient-to-r from-blue-50 to-indigo-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div class="shrink-0">
            <Image
              :src="userInfo.avatar"
              :width="96"
              :height="96"
              fallback="/default-avatar.png"
              class="rounded-2xl shadow-lg"
            />
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <h2 class="mb-1 text-2xl font-bold text-gray-800">
                {{ userInfo.nickname || '未设置昵称' }}
              </h2>
              <p class="text-lg font-medium text-gray-600">
                {{ userInfo.phone }}
              </p>
            </div>
            <div class="flex items-center space-x-4">
              <Tag
                :color="userInfo.status === 1 ? 'success' : 'error'"
                class="rounded-full px-3 py-1"
              >
                {{ statusText }}
              </Tag>
              <Tag color="blue" class="rounded-full px-3 py-1">
                {{ genderText }}
              </Tag>
            </div>
          </div>
        </div>
      </div>

      <!-- Tab 导航 -->
      <Tabs
        :active-key="activeTab"
        @change="handleTabChange"
        type="card"
        class="user-detail-tabs"
      >
        <!-- 基础用户信息 Tab -->
        <TabPane key="user">
          <template #tab>
            <span class="flex items-center gap-2">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                <path
                  d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"
                />
              </svg>
              基础信息
            </span>
          </template>

          <div class="space-y-4">
            <!-- 基本信息 -->
            <Card title="基本信息" size="small">
              <div class="space-y-4">
                <div class="rounded-lg bg-blue-50 p-4">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-600">ID</span>
                    <span class="font-mono text-gray-800">{{
                      userInfo.id
                    }}</span>
                  </div>
                </div>
                <div class="rounded-lg bg-green-50 p-4">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-600">手机号</span>
                    <span class="font-semibold text-gray-900">{{
                      userInfo.phone
                    }}</span>
                  </div>
                </div>
                <div class="rounded-lg bg-purple-50 p-4">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-600">昵称</span>
                    <span class="text-gray-800">{{
                      userInfo.nickname || '未设置'
                    }}</span>
                  </div>
                </div>
                <div class="rounded-lg bg-orange-50 p-4">
                  <div class="flex items-center justify-between">
                    <span class="font-medium text-gray-600">性别</span>
                    <span class="text-gray-800">{{ genderText }}</span>
                  </div>
                </div>
              </div>
            </Card>
            <!-- 个人简介 -->
            <Card
              v-if="userInfo.profile"
              title="个人简介"
              size="small"
              class="profile-card"
            >
              <div class="prose prose-sm max-w-none">
                <p class="whitespace-pre-line leading-relaxed text-gray-700">
                  {{ userInfo.profile }}
                </p>
              </div>
            </Card>
            <!-- 时间信息 -->
            <Card title="时间信息" size="small">
              <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
                <div class="rounded-lg bg-blue-50 p-4 text-center">
                  <div class="mb-2 text-sm text-gray-600">创建时间</div>
                  <div class="font-medium text-gray-800">
                    {{ formatDateTime(userInfo.createdAt || '') }}
                  </div>
                </div>
                <div class="rounded-lg bg-green-50 p-4 text-center">
                  <div class="mb-2 text-sm text-gray-600">更新时间</div>
                  <div class="font-medium text-gray-800">
                    {{ formatDateTime(userInfo.updatedAt || '') }}
                  </div>
                </div>
              </div>
            </Card>
          </div>
        </TabPane>

        <!-- 会员信息 Tab -->
        <TabPane key="membership" :disabled="!hasMembership">
          <template #tab>
            <span class="flex items-center gap-2">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                <path
                  d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
                />
              </svg>
              会员信息
              <span v-if="!hasMembership" class="text-xs text-gray-400">
                (无)
              </span>
            </span>
          </template>

          <div v-if="membershipInfo" class="space-y-6">
            <!-- 详细信息 -->
            <div class="space-y-4">
              <!-- 基本信息 -->
              <Card title="💎 会员基本信息" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-blue-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">ID</span>
                      <span class="font-mono text-gray-800">{{
                        membershipInfo.id
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-green-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">用户ID</span>
                      <span class="font-mono text-gray-800">{{
                        membershipInfo.userId
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-yellow-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">会员类型</span>
                      <span class="text-gray-800">{{
                        membershipTypeText
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-purple-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">状态</span>
                      <span class="text-gray-800">{{
                        membershipStatusText
                      }}</span>
                    </div>
                  </div>
                  <div
                    class="rounded-lg p-3"
                    :class="
                      typeof membershipExpiredText === 'object' &&
                      membershipExpiredText.isExpired
                        ? 'bg-red-50'
                        : 'bg-orange-50'
                    "
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">到期时间</span>
                      <span
                        class="font-medium"
                        :class="
                          typeof membershipExpiredText === 'object' &&
                          membershipExpiredText.isExpired
                            ? 'text-red-600'
                            : 'text-gray-800'
                        "
                      >
                        {{
                          typeof membershipExpiredText === 'string'
                            ? membershipExpiredText
                            : membershipExpiredText.text
                        }}
                      </span>
                    </div>
                  </div>
                </div>
              </Card>

              <!-- 续费信息 -->
              <Card title="🔄 自动续费信息" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-indigo-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">自动续费</span>
                      <Tag
                        :color="
                          membershipInfo.autoRenew === 1 ? 'success' : 'default'
                        "
                        class="rounded-full px-3 py-1"
                      >
                        {{ autoRenewText }}
                      </Tag>
                    </div>
                  </div>
                  <div class="rounded-lg bg-pink-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">续费天数</span>
                      <span class="text-gray-800">{{ autoRenewDaysText }}</span>
                    </div>
                  </div>
                </div>
              </Card>

              <!-- 时间信息 -->
              <Card title="⏰ 时间记录" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-cyan-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">创建时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(membershipInfo.createdAt || '') }}
                      </span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-emerald-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">更新时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(membershipInfo.updatedAt || '') }}
                      </span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>
          </div>

          <div v-else class="py-8">
            <Result
              status="info"
              title="无会员信息"
              sub-title="该用户暂无会员信息"
            />
          </div>
        </TabPane>

        <!-- 微信公众号信息 Tab -->
        <TabPane key="gzh" :disabled="!hasGzhUser">
          <template #tab>
            <span class="flex items-center gap-2">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                <path
                  d="M8.691 2.188C8.691 1.533 8.158 1 7.503 1s-1.188.533-1.188 1.188.533 1.188 1.188 1.188 1.188-.533 1.188-1.188zM22 7.5c0 1.933-1.567 3.5-3.5 3.5S15 9.433 15 7.5 16.567 4 18.5 4 22 5.567 22 7.5z"
                />
              </svg>
              微信公众号
              <span v-if="!hasGzhUser" class="text-xs text-gray-400">(无)</span>
            </span>
          </template>

          <div v-if="gzhLoading" class="flex justify-center py-12">
            <Spin size="large" tip="加载公众号信息中..." />
          </div>

          <div v-else-if="gzhError" class="py-8">
            <Result
              status="error"
              title="加载失败"
              sub-title="无法加载微信公众号用户信息"
            >
              <template #extra>
                <button
                  @click="loadGzhUserInfo"
                  class="rounded-lg bg-blue-500 px-4 py-2 text-white transition-colors hover:bg-blue-600"
                >
                  重试
                </button>
              </template>
            </Result>
          </div>

          <div v-else-if="gzhUserInfo" class="space-y-6">
            <!-- 详细信息 -->
            <div class="space-y-4">
              <!-- 基本信息 -->
              <Card title="🔑 身份信息" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-blue-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">ID</span>
                      <span class="font-mono text-gray-800">{{
                        gzhUserInfo.id
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-green-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">公众号AppID</span>
                      <span class="font-mono text-gray-800">{{
                        gzhUserInfo.appId
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-purple-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">OpenID</span>
                      <span class="font-mono text-gray-800">{{
                        gzhUserInfo.openid
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-orange-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">UnionID</span>
                      <span class="font-mono text-gray-800">{{
                        gzhUserInfo.unionid || '无'
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
                        gzhUserInfo.nickname || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-indigo-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">语言</span>
                      <span class="text-gray-800">{{
                        gzhUserInfo.language || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-yellow-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">标签ID</span>
                      <span class="text-gray-800">{{
                        gzhUserInfo.tagIds || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-teal-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">头像</span>
                      <div>
                        <Image
                          v-if="gzhUserInfo.avatarUrl"
                          :src="gzhUserInfo.avatarUrl"
                          :width="40"
                          :height="40"
                          :preview="true"
                          class="rounded-lg shadow-sm"
                        />
                        <span v-else class="text-gray-400">无</span>
                      </div>
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
                        gzhUserInfo.country || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-blue-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">省份</span>
                      <span class="text-gray-800">{{
                        gzhUserInfo.province || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-green-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">城市</span>
                      <span class="text-gray-800">{{
                        gzhUserInfo.city || '无'
                      }}</span>
                    </div>
                  </div>
                </div>
              </Card>

              <!-- 时间信息 -->
              <Card title="⏰ 时间记录" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-cyan-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">创建时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(gzhUserInfo.createdAt || '') }}
                      </span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-emerald-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">更新时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(gzhUserInfo.updatedAt || '') }}
                      </span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>

            <!-- 备注 -->
            <Card v-if="gzhUserInfo.remark" title="📝 备注信息" size="small">
              <div
                class="rounded-lg border-l-4 border-amber-400 bg-amber-50 p-4"
              >
                <p class="leading-relaxed text-gray-700">
                  {{ gzhUserInfo.remark }}
                </p>
              </div>
            </Card>
          </div>
        </TabPane>

        <!-- 微信小程序信息 Tab -->
        <TabPane key="xcx" :disabled="!hasXcxUser">
          <template #tab>
            <span class="flex items-center gap-2">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="currentColor">
                <path
                  d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 17.93c-3.94-.49-7-3.85-7-7.93 0-.62.08-1.21.21-1.79L9 15v1c0 1.1.9 2 2 2v1.93zm6.9-2.54c-.26-.81-1-1.39-1.9-1.39h-1v-3c0-.55-.45-1-1-1H8v-2h2c.55 0 1-.45 1-1V7h2c1.1 0 2-.9 2-2v-.41c2.93 1.19 5 4.06 5 7.41 0 2.08-.8 3.97-2.1 5.39z"
                />
              </svg>
              微信小程序
              <span v-if="!hasXcxUser" class="text-xs text-gray-400">(无)</span>
            </span>
          </template>

          <div v-if="xcxLoading" class="flex justify-center py-12">
            <Spin size="large" tip="加载小程序信息中..." />
          </div>

          <div v-else-if="xcxError" class="py-8">
            <Result
              status="error"
              title="加载失败"
              sub-title="无法加载微信小程序用户信息"
            >
              <template #extra>
                <button
                  @click="loadXcxUserInfo"
                  class="rounded-lg bg-blue-500 px-4 py-2 text-white transition-colors hover:bg-blue-600"
                >
                  重试
                </button>
              </template>
            </Result>
          </div>

          <div v-else-if="xcxUserInfo" class="space-y-6">
            <!-- 详细信息 -->
            <div class="space-y-4">
              <!-- 身份信息 -->
              <Card title="🆔 身份信息" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-blue-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">ID</span>
                      <span class="font-mono text-gray-800">{{
                        xcxUserInfo.id
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-green-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">小程序AppID</span>
                      <span class="font-mono text-gray-800">{{
                        xcxUserInfo.appId
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-purple-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">OpenID</span>
                      <span class="font-mono text-gray-800">{{
                        xcxUserInfo.openid
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-orange-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">UnionID</span>
                      <span class="font-mono text-gray-800">{{
                        xcxUserInfo.unionid || '无'
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
                        xcxUserInfo.nickname || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-indigo-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">语言</span>
                      <span class="text-gray-800">{{
                        xcxUserInfo.language || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-teal-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">头像</span>
                      <div>
                        <Image
                          v-if="xcxUserInfo.avatarURL"
                          :src="xcxUserInfo.avatarURL"
                          :width="40"
                          :height="40"
                          :preview="true"
                          class="rounded-lg shadow-sm"
                        />
                        <span v-else class="text-gray-400">无</span>
                      </div>
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
                        xcxUserInfo.country || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-blue-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">省份</span>
                      <span class="text-gray-800">{{
                        xcxUserInfo.province || '无'
                      }}</span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-green-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">城市</span>
                      <span class="text-gray-800">{{
                        xcxUserInfo.city || '无'
                      }}</span>
                    </div>
                  </div>
                </div>
              </Card>

              <!-- 时间信息 -->
              <Card title="⏰ 时间记录" size="small">
                <div class="space-y-3">
                  <div class="rounded-lg bg-cyan-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">创建时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(xcxUserInfo.createdAt || '') }}
                      </span>
                    </div>
                  </div>
                  <div class="rounded-lg bg-emerald-50 p-3">
                    <div class="flex items-center justify-between">
                      <span class="font-medium text-gray-600">更新时间</span>
                      <span class="text-gray-800">
                        {{ formatDateTime(xcxUserInfo.updatedAt || '') }}
                      </span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>

            <!-- 备注 -->
            <Card v-if="xcxUserInfo.remark" title="📝 备注信息" size="small">
              <div
                class="rounded-lg border-l-4 border-amber-400 bg-amber-50 p-4"
              >
                <p class="leading-relaxed text-gray-700">
                  {{ xcxUserInfo.remark }}
                </p>
              </div>
            </Card>
          </div>
        </TabPane>
      </Tabs>
    </div>
  </Modal>
</template>
