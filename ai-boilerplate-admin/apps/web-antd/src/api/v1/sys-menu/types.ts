/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysMenuReply = {
  /** id */
  id?: string;
};

export type CreateSysMenuReq = {
  /** 上级菜单 */
  pid?: string;
  /** 菜单名称 */
  name: string;
  /** 菜单类型(dir,menu,button) */
  type: string;
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
  sort: number;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type CreateSysMenuResponses = {
  /**
   * A successful response.
   */
  200: CreateSysMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysMenuReply = object;

export type DeleteSysMenuReq = {
  /** id集合 */
  id: string;
};

export type DeleteSysMenuResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysMenuInfoParams = {
  /** id */
  id: string;
};

export type GetSysMenuInfoReply = {
  info?: SysMenuInfo;
};

export type GetSysMenuInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysMenuInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysMenuListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysMenuInfo[];
};

export type GetSysMenuListResponses = {
  /**
   * A successful response.
   */
  200: GetSysMenuListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type SysMenuInfo = {
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
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type UpdateSysMenuReply = object;

export type UpdateSysMenuReq = {
  /** id */
  id: string;
  /** 上级菜单 */
  pid?: string;
  /** 菜单名称 */
  name: string;
  /** 菜单类型(dir,menu,button) */
  type: string;
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
  sort: number;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysMenuResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysMenuStatusReply = object;

export type UpdateSysMenuStatusReq = {
  /** id */
  id: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysMenuStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysMenuStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
