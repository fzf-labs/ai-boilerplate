/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type DeleteWxXcxUserReply = object;

export type DeleteWxXcxUserReq = {
  /** 编号 */
  id: string;
};

export type DeleteWxXcxUserResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxXcxUserReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxXcxUserInfoParams = {
  /** 编号 */
  id: string;
};

export type GetWxXcxUserInfoReply = {
  info?: WxXcxUserInfo;
};

export type GetWxXcxUserInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxXcxUserInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxXcxUserListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetWxXcxUserListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxXcxUserInfo[];
};

export type GetWxXcxUserListResponses = {
  /**
   * A successful response.
   */
  200: GetWxXcxUserListReply;
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

export type WxXcxUserInfo = {
  /** 编号 */
  id?: string;
  /** 微信小程序 appid */
  appId?: string;
  /** 用户标识 */
  openid?: string;
  /** 微信生态唯一标识 */
  unionid?: string;
  /** 昵称 */
  nickname?: string;
  /** 头像地址 */
  avatarURL?: string;
  /** 语言 */
  language?: string;
  /** 国家 */
  country?: string;
  /** 省份 */
  province?: string;
  /** 城市 */
  city?: string;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
