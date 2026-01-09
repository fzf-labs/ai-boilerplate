<script setup lang="ts">
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { useAccess } from '@vben/access';
import { Page } from '@vben/common-ui';

import { Card, message, Spin, Tabs } from 'ant-design-vue';

import { sysAuthAdminInfo } from '#/api/v1/sys-auth';
import { useAuthStore } from '#/store';

import BaseInfo from './modules/base-info.vue';
import ProfileUser from './modules/profile-user.vue';
import ResetPwd from './modules/reset-pwd.vue';

const authStore = useAuthStore();
const activeName = ref('basicInfo');
const loading = ref(false);
const { hasAccessByCodes } = useAccess();
/** 加载个人信息 */
const profile = ref<SysAdminInfo>();
async function loadProfile() {
  try {
    loading.value = true;
    const result = await sysAuthAdminInfo({ options: {} });
    profile.value = result.info as unknown as SysAdminInfo;
  } catch (error: any) {
    console.error('加载个人信息失败:', error);
    message.error(error?.message || '加载个人信息失败');
  } finally {
    loading.value = false;
  }
}

/** 刷新个人信息 */
async function refreshProfile() {
  // 加载个人信息
  await loadProfile();

  // 更新 store
  await authStore.fetchAdminInfo();
}

/** 初始化 */
onMounted(loadProfile);
</script>

<template>
  <Page auto-content-height>
    <Spin :spinning="loading" tip="加载中...">
      <div class="flex flex-col gap-4 lg:flex-row">
        <!-- 左侧 个人信息 -->
        <Card class="w-full lg:w-2/5" title="个人信息">
          <ProfileUser :profile="profile" @success="refreshProfile" />
        </Card>

        <!-- 右侧 基本设置 -->
        <Card class="ml-3 w-3/5">
          <Tabs v-model:active-key="activeName" class="-mt-4">
            <Tabs.TabPane
              key="basicInfo"
              tab="基本设置"
              v-if="hasAccessByCodes(['system:person:profile:update'])"
            >
              <BaseInfo :profile="profile" @success="refreshProfile" />
            </Tabs.TabPane>
            <Tabs.TabPane
              key="resetPwd"
              tab="密码设置"
              v-if="hasAccessByCodes(['system:person:profile:update:password'])"
            >
              <ResetPwd />
            </Tabs.TabPane>
          </Tabs>
        </Card>
      </div>
    </Spin>
  </Page>
</template>
