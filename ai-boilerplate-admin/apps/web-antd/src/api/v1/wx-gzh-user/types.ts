/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type DeleteWxGzhUserReply = object;

export type DeleteWxGzhUserReq = {
  /** 编号 */
  id: string;
};

export type DeleteWxGzhUserResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhUserReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhUserInfoParams = {
  /** 编号 */
  id: string;
};

export type GetWxGzhUserInfoReply = {
  info?: WxGzhUserInfo;
};

export type GetWxGzhUserInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhUserInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhUserListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 微信公众号 appid */
  appId: string;
  /** 用户标识 */
  openid?: string;
  /** 昵称 */
  nickname?: string;
  /** 关注状态 */
  subscribeStatus?: number;
};

export type GetWxGzhUserListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhUserInfo[];
};

export type GetWxGzhUserListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhUserListReply;
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

export type WxGzhUserInfo = {
  /** 编号 */
  id?: string;
  /** 微信公众号 appid */
  appId?: string;
  /** 用户标识 */
  openid?: string;
  /** 微信生态唯一标识 */
  unionid?: string;
  /** 关注状态 */
  subscribeStatus?: number;
  /** 昵称 */
  nickname?: string;
  /** 头像地址 */
  avatarUrl?: string;
  /** 语言 */
  language?: string;
  /** 国家 */
  country?: string;
  /** 省份 */
  province?: string;
  /** 城市 */
  city?: string;
  /** 标签编号数组 */
  tagIds?: string;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
