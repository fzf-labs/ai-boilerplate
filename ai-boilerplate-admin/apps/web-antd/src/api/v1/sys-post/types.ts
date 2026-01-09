/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysPostReply = {
  /** 编号 */
  id?: string;
};

export type CreateSysPostReq = {
  /** 租户Id */
  tenantId?: string;
  /** 岗位名称 */
  name: string;
  /** 岗位编码 */
  code?: string;
  /** 备注 */
  remark?: string;
  /** 排序值 */
  sort: number;
  /** 0=禁用 1=开启 */
  status: number;
};

export type CreateSysPostResponses = {
  /**
   * A successful response.
   */
  200: CreateSysPostReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysPostReply = object;

export type DeleteSysPostReq = {
  /** 编号 */
  id: string;
};

export type DeleteSysPostResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysPostReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysPostInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSysPostInfoReply = {
  info?: SysPostInfo;
};

export type GetSysPostInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysPostInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysPostListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 岗位名称 */
  name?: string;
  /** 岗位编码 */
  code?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSysPostListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysPostInfo[];
};

export type GetSysPostListResponses = {
  /**
   * A successful response.
   */
  200: GetSysPostListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysPostSelectorReply = {
  /** 列表数据 */
  list?: SysPostSelectorItem[];
};

export type GetSysPostSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSysPostSelectorReply;
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

export type SysPostInfo = {
  /** 编号 */
  id?: string;
  /** 租户Id */
  tenantId?: string;
  /** 岗位名称 */
  name?: string;
  /** 岗位编码 */
  code?: string;
  /** 备注 */
  remark?: string;
  /** 排序值 */
  sort?: number;
  /** 0=禁用 1=开启 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type SysPostSelectorItem = {
  /** 编号 */
  id?: string;
  /** 岗位名称 */
  name?: string;
};

export type UpdateSysPostReply = object;

export type UpdateSysPostReq = {
  /** 编号 */
  id: string;
  /** 租户Id */
  tenantId?: string;
  /** 岗位名称 */
  name: string;
  /** 岗位编码 */
  code?: string;
  /** 备注 */
  remark?: string;
  /** 排序值 */
  sort: number;
  /** 0=禁用 1=开启 */
  status: number;
};

export type UpdateSysPostResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysPostReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysPostStatusReply = object;

export type UpdateSysPostStatusReq = {
  /** 编号 */
  id: string;
  /** 0=禁用 1=开启 */
  status: number;
};

export type UpdateSysPostStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysPostStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
