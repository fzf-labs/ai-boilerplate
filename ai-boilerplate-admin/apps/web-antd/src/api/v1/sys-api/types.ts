/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysAPIReply = {
  /** 编号 */
  id?: string;
};

export type CreateSysAPIReq = {
  /** 权限Id */
  permissionId: string;
  /** 方法 */
  method: string;
  /** 路径 */
  path: string;
  /** 描述 */
  desc: string;
};

export type CreateSysAPIResponses = {
  /**
   * A successful response.
   */
  200: CreateSysAPIReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysAPIReply = object;

export type DeleteSysAPIReq = {
  /** 编号集合 */
  ids: string[];
};

export type DeleteSysAPIResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysAPIReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysAPIInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSysAPIInfoReply = {
  info?: SysAPIInfo;
};

export type GetSysAPIInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysAPIInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysAPIListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysAPIInfo[];
};

export type GetSysAPIListReq = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetSysAPIListResponses = {
  /**
   * A successful response.
   */
  200: GetSysAPIListReply;
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

export type SysAPIInfo = {
  /** 编号 */
  id?: string;
  /** 权限Id */
  permissionId?: string;
  /** 方法 */
  method?: string;
  /** 路径 */
  path?: string;
  /** 描述 */
  desc?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type UpdateSysAPIReply = object;

export type UpdateSysAPIReq = {
  /** 编号 */
  id: string;
  /** 权限Id */
  permissionId: string;
  /** 方法 */
  method: string;
  /** 路径 */
  path: string;
  /** 描述 */
  desc: string;
};

export type UpdateSysAPIResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysAPIReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
