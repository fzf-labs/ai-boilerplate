/* eslint-disable */
// @ts-ignore

export type ActivityInfo = {
  /** id */
  id?: string;
  /** 租户id */
  tenantId?: string;
  /** 标题 */
  title?: string;
  /** 图片URL */
  imageURL?: string;
  /** 跳转链接 */
  linkURL?: string;
  /** 跳转类型 */
  linkType?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 开始时间 */
  startTime?: string;
  /** 结束时间 */
  endTime?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateActivityReply = {
  /** id */
  id?: string;
};

export type CreateActivityReq = {
  /** 租户id */
  tenantId?: string;
  /** 标题 */
  title: string;
  /** 图片URL */
  imageURL: string;
  /** 跳转链接 */
  linkURL: string;
  /** 跳转类型 */
  linkType: string;
  /** 排序 */
  sort: number;
  /** 状态(-1禁用,1开启) */
  status: number;
  /** 开始时间 */
  startTime?: string;
  /** 结束时间 */
  endTime?: string;
};

export type CreateActivityResponses = {
  /**
   * A successful response.
   */
  200: CreateActivityReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteActivityReply = object;

export type DeleteActivityReq = {
  /** id */
  id: string;
};

export type DeleteActivityResponses = {
  /**
   * A successful response.
   */
  200: DeleteActivityReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetActivityInfoParams = {
  /** id */
  id: string;
};

export type GetActivityInfoReply = {
  info?: ActivityInfo;
};

export type GetActivityInfoResponses = {
  /**
   * A successful response.
   */
  200: GetActivityInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetActivityListParams = {
  /** 关键词(标题) */
  keyword?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetActivityListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: ActivityInfo[];
};

export type GetActivityListResponses = {
  /**
   * A successful response.
   */
  200: GetActivityListReply;
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

export type UpdateActivityReply = object;

export type UpdateActivityReq = {
  /** id */
  id: string;
  /** 标题 */
  title: string;
  /** 图片URL */
  imageURL: string;
  /** 跳转链接 */
  linkURL: string;
  /** 跳转类型 */
  linkType: string;
  /** 排序 */
  sort: number;
  /** 状态(-1禁用,1开启) */
  status: number;
  /** 开始时间 */
  startTime?: string;
  /** 结束时间 */
  endTime?: string;
};

export type UpdateActivityResponses = {
  /**
   * A successful response.
   */
  200: UpdateActivityReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateActivityStatusReply = object;

export type UpdateActivityStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateActivityStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateActivityStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
