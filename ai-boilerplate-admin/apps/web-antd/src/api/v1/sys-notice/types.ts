/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysNoticeReply = {
  /** id */
  id?: string;
};

export type CreateSysNoticeReq = {
  /** 类型 */
  type: string;
  /** 标题 */
  title: string;
  /** 内容 */
  content: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type CreateSysNoticeResponses = {
  /**
   * A successful response.
   */
  200: CreateSysNoticeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysNoticeReply = object;

export type DeleteSysNoticeReq = {
  /** id */
  id: string;
};

export type DeleteSysNoticeResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysNoticeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNoticeInfoParams = {
  /** id */
  id: string;
};

export type GetSysNoticeInfoReply = {
  info?: SysNoticeInfo;
};

export type GetSysNoticeInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysNoticeInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNoticeListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 类型 */
  type?: string;
  /** 标题 */
  title?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
};

export type GetSysNoticeListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysNoticeInfo[];
};

export type GetSysNoticeListResponses = {
  /**
   * A successful response.
   */
  200: GetSysNoticeListReply;
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

export type SysNoticeInfo = {
  /** id */
  id?: string;
  /** 类型 */
  type?: string;
  /** 标题 */
  title?: string;
  /** 内容 */
  content?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type UpdateSysNoticeReply = object;

export type UpdateSysNoticeReq = {
  /** id */
  id: string;
  /** 类型 */
  type: string;
  /** 标题 */
  title: string;
  /** 内容 */
  content: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateSysNoticeResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysNoticeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysNoticeStatusReply = object;

export type UpdateSysNoticeStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateSysNoticeStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysNoticeStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
