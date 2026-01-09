import type { BasicUserInfo } from '@vben-core/typings';
/** 管理员信息 */
interface AdminInfo extends BasicUserInfo {
  homePath: string; // 首页地址
}

/** 菜单 */
interface Menu {
  id: string; // 菜单id
  pid: string; // 父级菜单id
  name: string; // 菜单名称
  type: string; // 菜单类型 dir menu button
  path: string; // 菜单路径
  permission: string; // 菜单权限
  icon: string; // 菜单图标
  component: string; // 菜单组件
  componentName: string; // 菜单组件名称
  sort: number; // 菜单排序
  status: number; // 菜单状态
  children?: Menu[]; // 子菜单
}

/** 管理员权限信息 */
interface AuthAdminInfo {
  info: AdminInfo; // 管理员信息
}

export type { AdminInfo, AuthAdminInfo, Menu };
