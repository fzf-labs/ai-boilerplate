/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetSysNotifyMessageInfoParams = {
  /** id */
  id: string;
};

export type GetSysNotifyMessageInfoReply = {
  info?: SysNotifyMessageInfo;
};

export type GetSysNotifyMessageInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysNotifyMessageInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNotifyMessageListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 接收人 */
  receiver?: string;
  /** 发送时间 */
  sendTime?: string[];
  /** 阅读状态    1: 已读 -1: 未读 */
  readStatus?: number;
};

export type GetSysNotifyMessageListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysNotifyMessageInfo[];
};

export type GetSysNotifyMessageListResponses = {
  /**
   * A successful response.
   */
  200: GetSysNotifyMessageListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNotifyMessageMyListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 发送人 */
  sender?: string;
  /** 发送时间 */
  sendTime?: string[];
  /** 阅读状态    1: 已读 -1: 未读 */
  readStatus?: number;
};

export type GetSysNotifyMessageMyListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysNotifyMessageInfo[];
};

export type GetSysNotifyMessageMyListResponses = {
  /**
   * A successful response.
   */
  200: GetSysNotifyMessageMyListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNotifyMessageMyUnreadCountReply = {
  /** 未读数量 */
  count?: number;
};

export type GetSysNotifyMessageMyUnreadCountResponses = {
  /**
   * A successful response.
   */
  200: GetSysNotifyMessageMyUnreadCountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysNotifyMessageMyUnreadListReply = {
  /** 列表数据 */
  list?: SysNotifyMessageInfo[];
};

export type GetSysNotifyMessageMyUnreadListResponses = {
  /**
   * A successful response.
   */
  200: GetSysNotifyMessageMyUnreadListReply;
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

export type SysNotifyMessageInfo = {
  /** id */
  id?: string;
  /** 租户id */
  tenantId?: string;
  /** 消息类型 */
  type?: string;
  /** 主题 */
  subject?: string;
  /** 内容 */
  content?: string;
  /** 发送人 */
  sender?: string;
  /** 接收人 */
  receiver?: string;
  /** 发送时间 */
  sendTime?: string;
  /** 阅读时间 */
  readTime?: string;
  /** 扩展 */
  extend?: string;
  /** 发送人名称 */
  senderName?: string;
  /** 发送人头像 */
  senderAvatar?: string;
  /** 接收人名称 */
  receiverName?: string;
  /** 接收人头像 */
  receiverAvatar?: string;
};

export type UpdateSysNotifyMessageAllReadReply = object;

export type UpdateSysNotifyMessageAllReadReq = object;

export type UpdateSysNotifyMessageAllReadResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysNotifyMessageAllReadReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysNotifyMessageReadReply = object;

export type UpdateSysNotifyMessageReadReq = {
  /** id列表 */
  ids: string[];
};

export type UpdateSysNotifyMessageReadResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysNotifyMessageReadReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
