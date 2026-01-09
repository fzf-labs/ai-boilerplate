/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type DeleteMailLogReply = object;

export type DeleteMailLogReq = {
  /** id */
  id: string;
};

export type DeleteMailLogResponses = {
  /**
   * A successful response.
   */
  200: DeleteMailLogReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailLogInfoParams = {
  /** id */
  id: string;
};

export type GetMailLogInfoReply = {
  info?: MailLogInfo;
};

export type GetMailLogInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMailLogInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailLogListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 邮箱账号编号 */
  accountId?: string;
  /** 发送邮箱地址 */
  fromMail?: string;
  /** 接收邮箱地址 */
  toMail?: string;
  /** 模板编号 */
  templateId?: string;
};

export type GetMailLogListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MailLogInfo[];
};

export type GetMailLogListResponses = {
  /**
   * A successful response.
   */
  200: GetMailLogListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MailLogInfo = {
  /** id */
  id?: string;
  /** 邮箱账号编号 */
  accountId?: string;
  /** 发送邮箱地址 */
  fromMail?: string;
  /** 接收邮箱地址 */
  toMail?: string;
  /** 模板编号 */
  templateId?: string;
  /** 模板编码 */
  templateCode?: string;
  /** 模版发送人名称 */
  templateNickname?: string;
  /** 邮件标题 */
  templateTitle?: string;
  /** 邮件内容 */
  templateContent?: string;
  /** 邮件参数 */
  templateParams?: string;
  /** 发送状态 */
  sendStatus?: number;
  /** 发送时间 */
  sendTime?: string;
  /** 发送返回的消息 ID */
  sendMessageId?: string;
  /** 发送异常 */
  sendException?: string;
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
