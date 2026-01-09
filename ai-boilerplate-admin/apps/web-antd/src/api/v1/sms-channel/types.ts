/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSmsChannelReply = {
  /** id */
  id?: string;
};

export type CreateSmsChannelReq = {
  /** 渠道名称 */
  name: string;
  /** 运营商 */
  operator: string;
  /** 备注 */
  remark?: string;
  /** 短信 API 的账号 */
  APIKey: string;
  /** 短信 API 的秘钥 */
  APISecret?: string;
  /** 短信发送回调 URL */
  callbackURL?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
};

export type CreateSmsChannelResponses = {
  /**
   * A successful response.
   */
  200: CreateSmsChannelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSmsChannelReply = object;

export type DeleteSmsChannelReq = {
  /** id */
  id: string;
};

export type DeleteSmsChannelResponses = {
  /**
   * A successful response.
   */
  200: DeleteSmsChannelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsChannelInfoParams = {
  /** id */
  id: string;
};

export type GetSmsChannelInfoReply = {
  info?: SmsChannelInfo;
};

export type GetSmsChannelInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSmsChannelInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsChannelListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 渠道名称 */
  name?: string;
  /** 运营商编码 */
  operator?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSmsChannelListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SmsChannelInfo[];
};

export type GetSmsChannelListResponses = {
  /**
   * A successful response.
   */
  200: GetSmsChannelListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsChannelOperatorReply = {
  /** 列表数据 */
  list?: SmsChannelOperator[];
};

export type GetSmsChannelOperatorResponses = {
  /**
   * A successful response.
   */
  200: GetSmsChannelOperatorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsChannelSelectorReply = {
  /** 列表数据 */
  list?: SmsChannelSelector[];
};

export type GetSmsChannelSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSmsChannelSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SmsChannelInfo = {
  /** id */
  id?: string;
  /** 渠道名称 */
  name?: string;
  /** 运营商 */
  operator?: string;
  /** 备注 */
  remark?: string;
  /** 短信 API 的账号 */
  APIKey?: string;
  /** 短信 API 的秘钥 */
  APISecret?: string;
  /** 短信发送回调 URL */
  callbackURL?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 运营商名称 */
  operatorName?: string;
};

export type SmsChannelOperator = {
  /** 运营商名称 */
  name?: string;
  /** 运营商编码 */
  operator?: string;
};

export type SmsChannelSelector = {
  /** id */
  id?: string;
  /** 渠道名称 */
  name?: string;
  /** 运营商编码 */
  operator?: string;
  /** 运营商名称 */
  operatorName?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateSmsChannelReply = object;

export type UpdateSmsChannelReq = {
  /** id */
  id: string;
  /** 渠道名称 */
  name: string;
  /** 运营商 */
  operator: string;
  /** 备注 */
  remark?: string;
  /** 短信 API 的账号 */
  APIKey: string;
  /** 短信 API 的秘钥 */
  APISecret?: string;
  /** 短信发送回调 URL */
  callbackURL?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
};

export type UpdateSmsChannelResponses = {
  /**
   * A successful response.
   */
  200: UpdateSmsChannelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSmsChannelStatusReply = object;

export type UpdateSmsChannelStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateSmsChannelStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSmsChannelStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
