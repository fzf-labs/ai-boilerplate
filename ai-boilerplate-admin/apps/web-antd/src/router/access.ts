import type {
  ComponentRecordType,
  GenerateMenuAndRoutesOptions,
} from '@vben/types';

import { generateAccessible } from '@vben/access';
import { preferences } from '@vben/preferences';
import { convertServerMenuToRouteRecordStringComponent } from '@vben/utils';

import { sysAuthMenu } from '#/api/v1/sys-auth';
import { BasicLayout, IFrameView } from '#/layouts';

const forbiddenComponent = () => import('#/views/_core/fallback/forbidden.vue');

async function generateAccess(options: GenerateMenuAndRoutesOptions) {
  const pageMap: ComponentRecordType = import.meta.glob('../views/**/*.vue');
  const layoutMap: ComponentRecordType = {
    BasicLayout,
    IFrameView,
  };
  return await generateAccessible(preferences.app.accessMode, {
    ...options,
    fetchMenuListAsync: async () => {
      const list = await sysAuthMenu({});
      return convertServerMenuToRouteRecordStringComponent(
        (list.menu || []) as unknown as Parameters<
          typeof convertServerMenuToRouteRecordStringComponent
        >[0],
      );
    },
    // 可以指定没有权限跳转403页面
    forbiddenComponent,
    // 如果 route.meta.menuVisibleWithForbidden = true
    layoutMap,
    pageMap,
  });
}

export { generateAccess };
