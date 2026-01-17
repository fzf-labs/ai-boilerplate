<script lang="ts" setup>
import type { NotificationItem } from '@vben/layouts';

import { computed, onMounted, ref, watch } from 'vue';

import { useAccess } from '@vben/access';
import { AuthenticationLoginExpiredModal } from '@vben/common-ui';
import { useWatermark } from '@vben/hooks';
import { AntdProfileOutlined } from '@vben/icons';
import {
  BasicLayout,
  LockScreen,
  Notification,
  UserDropdown,
} from '@vben/layouts';
import { $t } from '@vben/locales';
import { preferences } from '@vben/preferences';
import { useAccessStore, useUserStore } from '@vben/stores';
import { formatDateTime } from '@vben/utils';

import {
  getSysNotifyMessageMyUnreadCount,
  getSysNotifyMessageMyUnreadList,
  updateSysNotifyMessageAllRead,
  updateSysNotifyMessageRead,
} from '#/api/v1/sys-notify-message';
import { router } from '#/router';
import { useAuthStore } from '#/store';
import LoginForm from '#/views/_core/authentication/login.vue';

const notifications = ref<(NotificationItem & { id?: string })[]>([]);
const unreadCount = ref(0);
const showDot = computed(() => unreadCount.value > 0);

const userStore = useUserStore();
const authStore = useAuthStore();
const accessStore = useAccessStore();
const { destroyWatermark, updateWatermark } = useWatermark();
const { hasAccessByCodes } = useAccess();
const menus = computed(() => [
  {
    handler: () => {
      router.push({ path: '/system/person/profile' });
    },
    icon: AntdProfileOutlined,
    text: $t('ui.widgets.profile'),
    show: hasAccessByCodes(['system:person:profile:query']),
  },
]);

const avatar = computed(() => {
  return userStore.userInfo?.avatar ?? preferences.app.defaultAvatar;
});

async function handleLogout() {
  await authStore.logout(false);
}

/** 获得未读消息数 */
async function handleNotificationGetUnreadCount() {
  const res = await getSysNotifyMessageMyUnreadCount({});
  unreadCount.value = res.count ?? 0;
}

/** 获得消息列表 */
async function handleNotificationGetList() {
  const res = await getSysNotifyMessageMyUnreadList({});
  notifications.value = (res.list ?? []).map((item) => {
    const sendTime = item.sendTime ?? '';
    const readTime = item.readTime ?? '';
    return {
      avatar: preferences.app.defaultAvatar,
      date: formatDateTime(sendTime) as string,
      isRead: readTime !== '',
      id: item.id,
      message: item.content ?? '',
      title: item.subject ?? '',
    };
  });
}

/** 跳转我的站内信 */
function handleNotificationViewAll() {
  router.push({
    path: '/system/person/message',
  });
}

/** 标记所有已读 */
async function handleNotificationMakeAll() {
  await updateSysNotifyMessageAllRead({ body: {} });
  unreadCount.value = 0;
  notifications.value = [];
}

/** 清空通知 */
async function handleNotificationClear() {
  handleNotificationMakeAll();
}

/** 标记单个已读 */
async function handleNotificationRead(item: NotificationItem) {
  // 从数组中找到对应的项（包含 id）
  const found = notifications.value.find(
    (n) =>
      n.title === item.title &&
      n.message === item.message &&
      n.date === item.date,
  );

  if (!found?.id) {
    return;
  }

  await updateSysNotifyMessageRead({ body: { ids: [found.id] } });
  await handleNotificationGetUnreadCount();
  notifications.value = notifications.value.filter((n) => n.id !== found.id);
}

/** 处理通知打开 */
function handleNotificationOpen(open: boolean) {
  if (!open) {
    return;
  }
  handleNotificationGetList();
  handleNotificationGetUnreadCount();
}

onMounted(() => {
  // 首次加载未读数量
  handleNotificationGetUnreadCount();
  // 轮询刷新未读数量
  setInterval(
    () => {
      if (userStore.userInfo) {
        handleNotificationGetUnreadCount();
      }
    },
    1000 * 60 * 2,
  );
});

watch(
  () => preferences.app.watermark,
  async (enable) => {
    if (enable) {
      await updateWatermark({
        content: `${userStore.userInfo?.username} - ${userStore.userInfo?.realName}`,
      });
    } else {
      destroyWatermark();
    }
  },
  {
    immediate: true,
  },
);
</script>

<template>
  <BasicLayout @clear-preferences-and-logout="handleLogout">
    <template #user-dropdown>
      <UserDropdown
        :avatar="avatar"
        :menus="menus"
        :text="userStore.userInfo?.nickname"
        :description="userStore.userInfo?.username"
        @logout="handleLogout"
      />
    </template>
    <template #notification>
      <Notification
        :dot="showDot"
        :notifications="notifications"
        @clear="handleNotificationClear"
        @make-all="handleNotificationMakeAll"
        @view-all="handleNotificationViewAll"
        @open="handleNotificationOpen"
        @read="handleNotificationRead"
      />
    </template>
    <template #extra>
      <AuthenticationLoginExpiredModal
        v-model:open="accessStore.loginExpired"
        :avatar="avatar"
      >
        <LoginForm />
      </AuthenticationLoginExpiredModal>
    </template>
    <template #lock-screen>
      <LockScreen :avatar="avatar" @to-login="handleLogout" />
    </template>
  </BasicLayout>
</template>
