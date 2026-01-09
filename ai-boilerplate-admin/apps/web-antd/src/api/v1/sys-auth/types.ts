/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type SysAdminInfo = {
  /** 编号 */
  id?: string;
  /** 用户名 */
  username?: string;
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 0=保密 1=女 2=男 */
  sex?: number;
  /** 邮件 */
  email?: string;
  /** 手机号 */
  mobile?: string;
  /** 角色 */
  roleId?: string;
  /** 部门 */
  deptId?: string;
  /** 岗位 */
  postId?: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 角色名称 */
  roleName?: string;
  /** 部门名称 */
  deptName?: string;
  /** 岗位名称 */
  postName?: string;
};

export type SysAuthAdminInfoReply = {
  info?: SysAdminInfo;
};

export type SysAuthAdminInfoResponses = {
  /**
   * A successful response.
   */
  200: SysAuthAdminInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthLoginReply = {
  /** token */
  token?: string;
  /** 过期时间 */
  expiredAt?: number;
  /** 刷新时间 */
  refreshAt?: number;
};

export type SysAuthLoginReq = {
  /** 账号 */
  username: string;
  /** 密码 */
  password: string;
};

export type SysAuthLoginResponses = {
  /**
   * A successful response.
   */
  200: SysAuthLoginReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthLogoutReply = object;

export type SysAuthLogoutReq = object;

export type SysAuthLogoutResponses = {
  /**
   * A successful response.
   */
  200: SysAuthLogoutReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthMenuReply = {
  /** 菜单 */
  menu?: SysMenuItem[];
};

export type SysAuthMenuResponses = {
  /**
   * A successful response.
   */
  200: SysAuthMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthPermissionReply = {
  /** 权限 */
  permission?: string[];
};

export type SysAuthPermissionResponses = {
  /**
   * A successful response.
   */
  200: SysAuthPermissionReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthUpdateAdminInfoReply = object;

export type SysAuthUpdateAdminInfoReq = {
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 0=保密 1=女 2=男 */
  sex?: number;
};

export type SysAuthUpdateAdminInfoResponses = {
  /**
   * A successful response.
   */
  200: SysAuthUpdateAdminInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysAuthUpdateAdminPasswordReply = object;

export type SysAuthUpdateAdminPasswordReq = {
  /** 旧密码 */
  oldPassword: string;
  /** 新密码 */
  newPassword: string;
};

export type SysAuthUpdateAdminPasswordResponses = {
  /**
   * A successful response.
   */
  200: SysAuthUpdateAdminPasswordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SysMenuItem = {
  /** id */
  id?: string;
  /** 上级菜单 */
  pid?: string;
  /** 菜单名称 */
  name?: string;
  /** 菜单类型(dir,menu,button) */
  type?: string;
  /** 路由路径 */
  path?: string;
  /** 权限标识 */
  permission?: string;
  /** 图标 */
  icon?: string;
  /** 组件路径 */
  component?: string;
  /** 组件名 */
  componentName?: string;
  /** 权重(排序) */
  sort?: number;
  /** 0=禁用 1=开启 */
  status?: number;
  /** 子级菜单 */
  children?: SysMenuItem[];
};
