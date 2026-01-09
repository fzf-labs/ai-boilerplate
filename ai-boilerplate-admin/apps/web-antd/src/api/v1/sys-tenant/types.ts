/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysTenantReply = {
  /** id */
  id?: string;
};

export type CreateSysTenantReq = {
  /** 名称 */
  name: string;
  /** 备注 */
  remark?: string;
  /** 过期时间 */
  expireTime?: string;
  /** 菜单 */
  menuIds?: string[];
  /** 状态 */
  status?: number;
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
};

export type CreateSysTenantResponses = {
  /**
   * A successful response.
   */
  200: CreateSysTenantReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysTenantReply = object;

export type DeleteSysTenantReq = {
  /** id */
  id: string;
};

export type DeleteSysTenantResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysTenantReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysTenantInfoParams = {
  /** id */
  id: string;
};

export type GetSysTenantInfoReply = {
  info?: SysTenantInfo;
};

export type GetSysTenantInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysTenantInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysTenantListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 名称 */
  name?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSysTenantListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysTenantInfo[];
};

export type GetSysTenantListResponses = {
  /**
   * A successful response.
   */
  200: GetSysTenantListReply;
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

export type SysTenantInfo = {
  /** id */
  id?: string;
  /** 名称 */
  name?: string;
  /** 备注 */
  remark?: string;
  /** 租户管理员Id */
  adminId?: string;
  /** 过期时间 */
  expireTime?: string;
  /** 菜单集合 */
  menuIds?: string[];
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 租户管理员名称 */
  adminName?: string;
};

export type UpdateSysTenantReply = object;

export type UpdateSysTenantReq = {
  /** id */
  id: string;
  /** 名称 */
  name: string;
  /** 备注 */
  remark?: string;
  /** 过期时间 */
  expireTime?: string;
  /** 菜单 */
  menuIds?: string[];
  /** 状态 */
  status?: number;
};

export type UpdateSysTenantResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysTenantReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysTenantStatusReply = object;

export type UpdateSysTenantStatusReq = {
  /** id */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateSysTenantStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysTenantStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
