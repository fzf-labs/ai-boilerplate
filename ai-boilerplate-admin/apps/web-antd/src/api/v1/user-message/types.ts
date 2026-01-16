/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetUserMessageInfoParams = {
  /** id */
  id: string;
};

export type GetUserMessageInfoReply = {
  info?: UserMessageInfo;
};

export type GetUserMessageInfoResponses = {
  /**
   * A successful response.
   */
  200: GetUserMessageInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserMessageListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 消息分类 */
  category?: string;
  /** 标题 */
  title?: string;
  /** 用户id */
  userId?: string;
  /** 消息批次id */
  messageId?: string;
  /** 投放范围 */
  audienceType?: string;
  /** 阅读状态 1已读 -1未读 */
  readStatus?: number;
};

export type GetUserMessageListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: UserMessageInfo[];
};

export type GetUserMessageListResponses = {
  /**
   * A successful response.
   */
  200: GetUserMessageListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SendUserMessageReply = {
  /** 消息批次id */
  messageId?: string;
  /** 发送数量 */
  total?: number;
};

export type SendUserMessageReq = {
  /** 消息分类(transaction/system/service) */
  category: string;
  /** 标题 */
  title: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverURL?: string;
  /** 内容 */
  content: string;
  /** 跳转链接 */
  linkURL?: string;
  /** 投放范围(all/segment/users) */
  audienceType: string;
  /** 指定用户ID列表 */
  userIds?: string[];
  /** 会员类型筛选 */
  membershipType?: string;
  /** 活跃天数筛选 */
  activeWithinDays?: number;
};

export type SendUserMessageResponses = {
  /**
   * A successful response.
   */
  200: SendUserMessageReply;
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

export type UserMessageInfo = {
  /** id */
  id?: string;
  /** 消息批次id */
  messageId?: string;
  /** 用户id */
  userId?: string;
  /** 消息分类(transaction/system/service) */
  category?: string;
  /** 标题 */
  title?: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverURL?: string;
  /** 内容 */
  content?: string;
  /** 跳转链接 */
  linkURL?: string;
  /** 投放范围(all/segment/users) */
  audienceType?: string;
  /** 投放条件/用户列表 */
  audienceValue?: string;
  /** 发送时间 */
  sentAt?: string;
  /** 阅读时间 */
  readAt?: string;
  /** 创建人 */
  adminId?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
