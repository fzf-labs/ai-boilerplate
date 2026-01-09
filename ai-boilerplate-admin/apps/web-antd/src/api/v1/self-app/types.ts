/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSelfAppReply = {
  /** ID */
  id?: string;
};

export type CreateSelfAppReq = {
  /** 包名 */
  packageName: string;
  /** 应用名称 */
  name: string;
  /** 平台类型(1Android 2iOS) */
  platform: number;
  /** 应用描述 */
  description?: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type CreateSelfAppResponses = {
  /**
   * A successful response.
   */
  200: CreateSelfAppReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSelfAppReply = object;

export type DeleteSelfAppReq = {
  /** ID */
  id: string;
};

export type DeleteSelfAppResponses = {
  /**
   * A successful response.
   */
  200: DeleteSelfAppReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSelfAppInfoParams = {
  /** ID */
  id: string;
};

export type GetSelfAppInfoReply = {
  info?: SelfAppInfo;
};

export type GetSelfAppInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSelfAppInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSelfAppListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 包名 */
  packageName?: string;
  /** 应用名称 */
  name?: string;
  /** 平台类型(1Android 2iOS) */
  platform?: number;
  /** 状态(-1禁用 1启用) */
  status?: number;
};

export type GetSelfAppListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SelfAppInfo[];
};

export type GetSelfAppListResponses = {
  /**
   * A successful response.
   */
  200: GetSelfAppListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SelfAppInfo = {
  /** ID */
  id?: string;
  /** 包名 */
  packageName?: string;
  /** 应用名称 */
  name?: string;
  /** 平台类型(1Android 2iOS) */
  platform?: number;
  /** 应用描述 */
  description?: string;
  /** 状态(-1禁用 1启用) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateSelfAppReply = object;

export type UpdateSelfAppReq = {
  /** ID */
  id: string;
  /** 包名 */
  packageName: string;
  /** 应用名称 */
  name: string;
  /** 平台类型(1Android 2iOS) */
  platform: number;
  /** 应用描述 */
  description?: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type UpdateSelfAppResponses = {
  /**
   * A successful response.
   */
  200: UpdateSelfAppReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSelfAppStatusReply = object;

export type UpdateSelfAppStatusReq = {
  /** ID */
  id: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type UpdateSelfAppStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSelfAppStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
