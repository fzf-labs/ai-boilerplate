/* eslint-disable */
// @ts-ignore

export type AiChatMessageInfo = {
  /** 消息编号 */
  id?: string;
  /** 租户编号 */
  tenantId?: string;
  /** 对话编号 */
  conversationId?: string;
  /** 回复编号 */
  replyId?: string;
  /** 用户编号 */
  adminId?: string;
  /** 角色编号 */
  roleId?: string;
  /** 消息类型 */
  type?: string;
  /** 模型标识 */
  model?: string;
  /** 模型编号 */
  modelId?: string;
  /** 消息内容 */
  content?: string;
  /** 是否携带上下文 */
  useContext?: boolean;
  /** 段落编号数组 */
  segmentIds?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type GetAiChatMessageListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 对话编号 */
  conversationId: string;
};

export type GetAiChatMessageListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiChatMessageInfo[];
};

export type GetAiChatMessageListResponses = {
  /**
   * A successful response.
   */
  200: GetAiChatMessageListReply;
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
