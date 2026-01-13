/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type BannerInfo = {
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
  /** 展示位置 */
  position?: string;
  /** 平台 */
  platform?: string;
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

export type CreateBannerReply = {
  /** id */
  id?: string;
};

export type CreateBannerReq = {
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
  /** 展示位置 */
  position: string;
  /** 平台 */
  platform: string;
  /** 排序 */
  sort: number;
  /** 状态(-1禁用,1开启) */
  status: number;
  /** 开始时间 */
  startTime?: string;
  /** 结束时间 */
  endTime?: string;
};

export type CreateBannerResponses = {
  /**
   * A successful response.
   */
  200: CreateBannerReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteBannerReply = object;

export type DeleteBannerReq = {
  /** id */
  id: string;
};

export type DeleteBannerResponses = {
  /**
   * A successful response.
   */
  200: DeleteBannerReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetBannerInfoParams = {
  /** id */
  id: string;
};

export type GetBannerInfoReply = {
  info?: BannerInfo;
};

export type GetBannerInfoResponses = {
  /**
   * A successful response.
   */
  200: GetBannerInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetBannerListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 标题 */
  title?: string;
  /** 展示位置 */
  position?: string;
  /** 平台 */
  platform?: string;
  /** 跳转类型 */
  linkType?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
};

export type GetBannerListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: BannerInfo[];
};

export type GetBannerListResponses = {
  /**
   * A successful response.
   */
  200: GetBannerListReply;
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

export type UpdateBannerReply = object;

export type UpdateBannerReq = {
  /** id */
  id: string;
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
  /** 展示位置 */
  position: string;
  /** 平台 */
  platform: string;
  /** 排序 */
  sort: number;
  /** 状态(-1禁用,1开启) */
  status: number;
  /** 开始时间 */
  startTime?: string;
  /** 结束时间 */
  endTime?: string;
};

export type UpdateBannerResponses = {
  /**
   * A successful response.
   */
  200: UpdateBannerReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateBannerStatusReply = object;

export type UpdateBannerStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateBannerStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateBannerStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
