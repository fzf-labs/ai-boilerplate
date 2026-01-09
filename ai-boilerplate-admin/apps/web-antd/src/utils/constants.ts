// todo @芋艿：要不要共享
/**
 * Created by 芋道源码
 *
 * 枚举类
 */

// ========== COMMON 模块 ==========
// 全局通用状态枚举
export const CommonStatusEnum = {
  ENABLE: 1, // 开启
  DISABLE: -1, // 禁用
};

// 全局用户类型枚举
export const UserTypeEnum = {
  MEMBER: 1, // 会员
  ADMIN: 2, // 管理员
};

// ========== SYSTEM 模块 ==========
/**
 * 菜单的类型枚举
 */
export const SystemMenuTypeEnum = {
  DIR: 'dir', // 目录
  MENU: 'menu', // 菜单
  BUTTON: 'button', // 按钮
};

/**
 * 角色的类型枚举
 */
export const SystemRoleTypeEnum = {
  SYSTEM: 1, // 内置角色
  CUSTOM: 2, // 自定义角色
};

/**
 * 数据权限的范围枚举
 */
export const SystemDataScopeEnum = {
  ALL: 'all', // 全部数据权限
  DEPT_ONLY: 'dept_only', // 部门数据权限
  DEPT_AND_CHILD: 'dept_and_child', // 部门及以下数据权限
  SELF: 'self', // 仅本人数据权限
};
