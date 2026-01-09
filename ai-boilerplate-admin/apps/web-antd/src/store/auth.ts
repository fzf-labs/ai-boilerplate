import type { AuthAdminInfo, Recordable } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { LOGIN_PATH } from '@vben/constants';
import { preferences } from '@vben/preferences';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';

import { notification } from 'ant-design-vue';
import { defineStore } from 'pinia';

import {
  getAdminInfoApi,
  getAdminPermissionApi,
  loginApi,
  logoutApi,
} from '#/api';
import { $t } from '#/locales';

export const useAuthStore = defineStore('auth', () => {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const router = useRouter();

  const loginLoading = ref(false);

  /**
   * 异步处理登录操作
   * Asynchronously handle the login process
   * @param params 登录表单数据
   */
  async function authLogin(
    params: Recordable<any>,
    onSuccess?: () => Promise<void> | void,
  ) {
    // 异步处理用户登录操作并获取 accessToken
    let adminInfo: AuthAdminInfo | null = null;
    try {
      loginLoading.value = true;
      const { token } = await loginApi(params);
      // 如果成功获取到 accessToken
      if (token) {
        accessStore.setAccessToken(token);
        // 获取用户信息并存储到 accessStore 中
        const fetchAdminInfoResult = await fetchAdminInfo();
        adminInfo = fetchAdminInfoResult;
        if (accessStore.loginExpired) {
          accessStore.setLoginExpired(false);
        } else {
          onSuccess
            ? await onSuccess?.()
            : await router.push(
                adminInfo?.info.homePath || preferences.app.defaultHomePath,
              );
        }
        if (adminInfo?.info.realName) {
          notification.success({
            description: `${$t('authentication.loginSuccessDesc')}:${adminInfo?.info.realName}`,
            duration: 3,
            message: $t('authentication.loginSuccess'),
          });
        }
      }
    } finally {
      loginLoading.value = false;
    }

    return {
      adminInfo,
    };
  }

  async function logout(redirect: boolean = true) {
    try {
      await logoutApi();
    } catch {
      // 不做任何处理
    }
    resetAllStores();
    accessStore.setLoginExpired(false);

    // 回登录页带上当前路由地址
    await router.replace({
      path: LOGIN_PATH,
      query: redirect
        ? {
            redirect: encodeURIComponent(router.currentRoute.value.fullPath),
          }
        : {},
    });
  }

  async function fetchAdminInfo() {
    const adminInfo = await getAdminInfoApi();
    const permission = await getAdminPermissionApi();
    userStore.setUserInfo(adminInfo.info);
    accessStore.setAccessCodes(permission.permission);
    return adminInfo;
  }

  function $reset() {
    loginLoading.value = false;
  }

  return {
    $reset,
    authLogin,
    fetchAdminInfo,
    loginLoading,
    logout,
  };
});
