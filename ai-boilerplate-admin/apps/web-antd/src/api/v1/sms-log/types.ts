/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetSmsLogInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSmsLogInfoReply = {
  info?: SmsLogInfo;
};

export type GetSmsLogInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSmsLogInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsLogListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetSmsLogListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SmsLogInfo[];
};

export type GetSmsLogListResponses = {
  /**
   * A successful response.
   */
  200: GetSmsLogListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SmsLogInfo = {
  /** 编号 */
  id?: string;
  /** 短信渠道编号 */
  smsChannelId?: string;
  /** 模板编号 */
  smsTemplateId?: string;
  /** 短信参数内容 */
  smsParamsContent?: string;
  /** 手机号 */
  mobile?: string;
  /** 用户id */
  userId?: string;
  /** 发送状态 */
  sendStatus?: string;
  /** 发送时间 */
  sendTime?: string;
  /** 接收状态 */
  receiveStatus?: string;
  /** 接收时间 */
  receiveTime?: string;
  /** 短信 API 发送结果的编码 */
  APISendCode?: string;
  /** 短信 API 发送失败的提示 */
  APISendMsg?: string;
  /** 短信 API 发送返回的唯一请求 ID */
  APIRequestID?: string;
  /** 短信 API 发送返回的序号 */
  APISerialNo?: string;
  /** API 接收结果的编码 */
  APIReceiveCode?: string;
  /** API 接收结果的说明 */
  APIReceiveMsg?: string;
  /** 创建时间 */
  createdAt?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};
