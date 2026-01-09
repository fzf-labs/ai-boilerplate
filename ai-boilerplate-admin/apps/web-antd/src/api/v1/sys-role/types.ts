/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysRoleReply = {
  /** 编号 */
  id?: string;
};

export type CreateSysRoleReq = {
  /** 名称 */
  name: string;
  /** 备注 */
  remark?: string;
  /** 数据范围 */
  dataScope: string;
  /** 菜单id集合 */
  menuIds?: string[];
  /** 排序值 */
  sort: number;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type CreateSysRoleResponses = {
  /**
   * A successful response.
   */
  200: CreateSysRoleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysRoleReply = object;

export type DeleteSysRoleReq = {
  /** 编号集合 */
  id: string;
};

export type DeleteSysRoleResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysRoleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysRoleInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSysRoleInfoReply = {
  info?: SysRoleInfo;
};

export type GetSysRoleInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysRoleInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysRoleListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 名称 */
  name?: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSysRoleListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysRoleInfo[];
};

export type GetSysRoleListResponses = {
  /**
   * A successful response.
   */
  200: GetSysRoleListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysRoleSelectorItem = {
  /** 编号 */
  id?: string;
  /** 名称 */
  name?: string;
};

export type GetSysRoleSelectorReply = {
  /** 列表数据 */
  list?: GetSysRoleSelectorItem[];
};

export type GetSysRoleSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSysRoleSelectorReply;
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

export type SysRoleInfo = {
  /** 编号 */
  id?: string;
  /** 租户id */
  tenantId?: string;
  /** 名称 */
  name?: string;
  /** 备注 */
  remark?: string;
  /** 数据范围 */
  dataScope?: string;
  /** 菜单id集合 */
  menuIds?: string[];
  /** 排序值 */
  sort?: number;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type UpdateSysRoleReply = object;

export type UpdateSysRoleReq = {
  /** 编号 */
  id: string;
  /** 名称 */
  name: string;
  /** 备注 */
  remark?: string;
  /** 数据范围 */
  dataScope: string;
  /** 菜单id集合 */
  menuIds?: string[];
  /** 排序值 */
  sort: number;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysRoleResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysRoleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysRoleStatusReply = object;

export type UpdateSysRoleStatusReq = {
  /** 编号 */
  id: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysRoleStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysRoleStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
