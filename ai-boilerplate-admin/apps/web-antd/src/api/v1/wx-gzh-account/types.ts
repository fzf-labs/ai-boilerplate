/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateWxGzhAccountReply = {
  /** 编号 */
  id?: string;
};

export type CreateWxGzhAccountReq = {
  /** 公众号名称 */
  name: string;
  /** 公众号账号 */
  account: string;
  /** 公众号appid */
  appId: string;
  /** 公众号密钥 */
  appSecret: string;
  /** 公众号url */
  URL?: string;
  /** 公众号 */
  token?: string;
  /** 加密密钥 */
  encodingAesKey?: string;
  /** 二维码图片URL */
  qrCodeURL?: string;
  /** 备注 */
  remark?: string;
};

export type CreateWxGzhAccountResponses = {
  /**
   * A successful response.
   */
  200: CreateWxGzhAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteWxGzhAccountReply = object;

export type DeleteWxGzhAccountReq = {
  /** 编号 */
  id: string;
};

export type DeleteWxGzhAccountResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhAccountInfoParams = {
  /** 编号 */
  id: string;
};

export type GetWxGzhAccountInfoReply = {
  info?: WxGzhAccountInfo;
};

export type GetWxGzhAccountInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhAccountInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhAccountListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 公众号账号 */
  account?: string;
  /** 公众号appid */
  appId?: string;
};

export type GetWxGzhAccountListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhAccountInfo[];
};

export type GetWxGzhAccountListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhAccountListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhAccountSelectorReply = {
  /** 列表数据 */
  list?: WxGzhAccountSelector[];
};

export type GetWxGzhAccountSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhAccountSelectorReply;
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

export type UpdateWxGzhAccountReply = object;

export type UpdateWxGzhAccountReq = {
  /** 编号 */
  id: string;
  /** 公众号名称 */
  name: string;
  /** 公众号账号 */
  account: string;
  /** 公众号appid */
  appId: string;
  /** 公众号url */
  appSecret: string;
  /** 公众号密钥 */
  URL?: string;
  /** 公众号token */
  token?: string;
  /** 加密密钥 */
  encodingAesKey?: string;
  /** 二维码图片URL */
  qrCodeURL?: string;
  /** 备注 */
  remark?: string;
};

export type UpdateWxGzhAccountResponses = {
  /**
   * A successful response.
   */
  200: UpdateWxGzhAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type WxGzhAccountInfo = {
  /** 编号 */
  id?: string;
  /** 公众号名称 */
  name?: string;
  /** 公众号账号 */
  account?: string;
  /** 公众号appid */
  appId?: string;
  /** 公众号密钥 */
  appSecret?: string;
  /** 公众号url */
  URL?: string;
  /** 公众号token */
  token?: string;
  /** 加密密钥 */
  encodingAesKey?: string;
  /** 二维码图片URL */
  qrCodeURL?: string;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type WxGzhAccountSelector = {
  /** 编号 */
  id?: string;
  /** 公众号名称 */
  name?: string;
  /** 公众号appid */
  appId?: string;
};
