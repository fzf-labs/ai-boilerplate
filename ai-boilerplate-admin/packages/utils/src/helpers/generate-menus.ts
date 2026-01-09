import type { Router, RouteRecordRaw } from 'vue-router';

// eslint-disable-next-line no-restricted-imports
import type { Menu } from '@vben/types';

import type {
  ExRouteRecordRaw,
  MenuRecordRaw,
  RouteMeta,
  RouteRecordStringComponent,
} from '@vben-core/typings';

import { filterTree, mapTree } from '@vben-core/shared/utils';

/**
 * 根据 routes 生成菜单列表
 * @param routes - 路由配置列表
 * @param router - Vue Router 实例
 * @returns 生成的菜单列表
 */
function generateMenus(
  routes: RouteRecordRaw[],
  router: Router,
): MenuRecordRaw[] {
  // 将路由列表转换为一个以 name 为键的对象映射
  const finalRoutesMap: { [key: string]: string } = Object.fromEntries(
    router.getRoutes().map(({ name, path }) => [name, path]),
  );

  let menus = mapTree<ExRouteRecordRaw, MenuRecordRaw>(routes, (route) => {
    // 获取最终的路由路径
    const path = finalRoutesMap[route.name as string] ?? route.path ?? '';

    const {
      meta = {} as RouteMeta,
      name: routeName,
      redirect,
      children = [],
    } = route;
    const {
      activeIcon,
      badge,
      badgeType,
      badgeVariants,
      hideChildrenInMenu = false,
      icon,
      link,
      order,
      title = '',
    } = meta;

    // 确保菜单名称不为空
    const name = (title || routeName || '') as string;

    // 处理子菜单
    const resultChildren = hideChildrenInMenu
      ? []
      : ((children as MenuRecordRaw[]) ?? []);

    // 设置子菜单的父子关系
    if (resultChildren.length > 0) {
      resultChildren.forEach((child) => {
        child.parents = [...(route.parents ?? []), path];
        child.parent = path;
      });
    }

    // 确定最终路径
    const resultPath = hideChildrenInMenu ? redirect || path : link || path;

    return {
      activeIcon,
      badge,
      badgeType,
      badgeVariants,
      icon,
      name,
      order,
      parent: route.parent,
      parents: route.parents,
      path: resultPath,
      show: !meta.hideInMenu,
      children: resultChildren,
    };
  });

  // 对菜单进行排序，避免order=0时被替换成999的问题
  menus = menus.sort((a, b) => (a?.order ?? 999) - (b?.order ?? 999));

  // 过滤掉隐藏的菜单项
  return filterTree(menus, (menu) => !!menu.show);
}

/**
 * 转换后端菜单数据为路由数据
 * @param menus 后端菜单数据
 * @param parent 父级菜单路径
 * @returns 路由数据
 */
function convertServerMenuToRouteRecordStringComponent(
  menus: Menu[],
  parent = '',
): RouteRecordStringComponent[] {
  const routes: RouteRecordStringComponent[] = [];
  if (!menus || menus.length === 0) {
    return routes;
  }
  menus.forEach((menu) => {
    const {
      id,
      pid,
      name,
      path,
      icon,
      component,
      permission,
      sort,
      status,
      children,
    } = menu;

    const isExternalLink =
      path.startsWith('http://') || path.startsWith('https://');
    if (isExternalLink && pid === '0') {
      const urlRoute: RouteRecordStringComponent = {
        component: 'IFrameView',
        meta: {
          hideInMenu: status !== 1,
          icon,
          link: path,
          order: sort,
          permission,
          title: name,
        },
        name,
        path: `/iframe/${id}`,
      };
      routes.push(urlRoute);
      return;
    }

    // 处理组件
    let componentName = component;

    // 处理布局组件
    if (children && children.length > 0 && pid === '') {
      componentName = 'BasicLayout';
    } else if (!children || children.length === 0) {
      componentName = component;
    }

    // 特殊处理 Layout 组件
    if (componentName === 'Layout') {
      componentName = 'BasicLayout';
    }

    // 父级菜单但不是顶级菜单时，组件为空
    if (children && children.length > 0 && pid !== '') {
      componentName = '';
    }

    // 处理路径
    let fullPath = path;
    if (parent) {
      fullPath = `${parent}/${path}`;
    }

    if (!fullPath.startsWith('/')) {
      fullPath = `/${fullPath}`;
    }

    // 构建路由对象
    const route: RouteRecordStringComponent = {
      component: componentName,
      meta: {
        hideInMenu: status !== 1,
        icon,
        keepAlive: true, // 默认开启缓存，可根据实际需求调整
        order: sort,
        permission,
        title: name,
      },
      name: `${name}${id}`, // 防止名称重复，加上ID
      path: fullPath,
    };

    // 递归处理子菜单
    if (children && children.length > 0) {
      route.children = convertServerMenuToRouteRecordStringComponent(
        children,
        fullPath,
      );
    }

    routes.push(route);
  });

  return routes;
}

export { convertServerMenuToRouteRecordStringComponent, generateMenus };
