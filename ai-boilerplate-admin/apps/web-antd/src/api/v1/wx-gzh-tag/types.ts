/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateWxGzhTagReply = {
  /** 主键 */
  id?: string;
};

export type CreateWxGzhTagReq = {
  /** 公众号 appId */
  appId: string;
  /** 标签名称 */
  name?: string;
};

export type CreateWxGzhTagResponses = {
  /**
   * A successful response.
   */
  200: CreateWxGzhTagReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteWxGzhTagReply = object;

export type DeleteWxGzhTagReq = {
  /** 主键 */
  id: string;
};

export type DeleteWxGzhTagResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhTagReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhTagInfoParams = {
  /** 主键 */
  id: string;
};

export type GetWxGzhTagInfoReply = {
  info?: WxGzhTagInfo;
};

export type GetWxGzhTagInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhTagInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhTagListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 公众号appId */
  appId: string;
};

export type GetWxGzhTagListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhTagInfo[];
};

export type GetWxGzhTagListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhTagListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhTagSelectorParams = {
  /** 公众号 appId */
  appId: string;
};

export type GetWxGzhTagSelectorReply = {
  /** 列表数据 */
  list?: WxGzhTagSelector[];
};

export type GetWxGzhTagSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhTagSelectorReply;
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

export type SyncWxGzhTagReply = object;

export type SyncWxGzhTagReq = {
  /** 公众号 appId */
  appId: string;
};

export type SyncWxGzhTagResponses = {
  /**
   * A successful response.
   */
  200: SyncWxGzhTagReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateWxGzhTagReply = object;

export type UpdateWxGzhTagReq = {
  /** 主键 */
  id: string;
  /** 标签名称 */
  name?: string;
};

export type UpdateWxGzhTagResponses = {
  /**
   * A successful response.
   */
  200: UpdateWxGzhTagReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type WxGzhTagInfo = {
  /** 主键 */
  id?: string;
  /** 公众号 appId */
  appId?: string;
  /** 公众号标签 id */
  tagId?: number;
  /** 标签名称 */
  name?: string;
  /** 粉丝数量 */
  count?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type WxGzhTagSelector = {
  /** 主键 */
  id?: string;
  /** 公众号标签 id */
  tagId?: number;
  /** 标签名称 */
  name?: string;
};
