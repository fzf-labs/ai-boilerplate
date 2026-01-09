/* eslint-disable */
// @ts-ignore

export type AiIndexChatConversationItem = {
  /** 对话编号 */
  id?: string;
  /** 租户编号 */
  tenantId?: string;
  /** 用户编号 */
  adminId?: string;
  /** 对话标题 */
  title?: string;
  /** 是否置顶 */
  pinned?: boolean;
  /** 置顶时间 */
  pinnedTime?: string;
  promptSetting?: PromptSetting;
  modelSetting?: ModelSetting;
  knowledgeSetting?: KnowledgeSetting;
  mcpSetting?: McpSetting;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type AiIndexChatMessageItem = {
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

export type CreateAiIndexChatConversationReply = {
  /** 对话编号 */
  id?: string;
};

export type CreateAiIndexChatConversationReq = {
  /** 对话标题 */
  title: string;
  promptSetting?: PromptSetting;
  modelSetting?: ModelSetting;
  knowledgeSetting?: KnowledgeSetting;
  mcpSetting?: McpSetting;
};

export type CreateAiIndexChatConversationResponses = {
  /**
   * A successful response.
   */
  200: CreateAiIndexChatConversationReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiIndexChatConversationReply = object;

export type DeleteAiIndexChatConversationReq = {
  /** 对话编号 */
  id: string;
};

export type DeleteAiIndexChatConversationResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiIndexChatConversationReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexChatConversationItemParams = {
  /** 对话编号 */
  id: string;
};

export type GetAiIndexChatConversationItemReply = {
  info?: AiIndexChatConversationItem;
};

export type GetAiIndexChatConversationItemResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexChatConversationItemReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexChatConversationListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexChatConversationListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexChatConversationItem[];
};

export type GetAiIndexChatConversationListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexChatConversationListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexChatMessageListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexChatMessageListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexChatMessageItem[];
};

export type GetAiIndexChatMessageListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexChatMessageListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type KnowledgeSetting = {
  /** 知识库编号 */
  knowledgeIds?: string[];
};

export type McpSetting = {
  /** mcp编号 */
  mcpIds?: string[];
};

export type ModelSetting = {
  /** 平台 */
  platform?: string;
  /** 编号 */
  platformId?: string;
  /** 模型类型 */
  modelType?: string;
  /** 模型id */
  modelId?: string;
  /** 采样温度。控制了生成文本时对每个候选词的概率分布进行平滑的程度。当取值为 0 时模型仅考虑对数概率最大的一个 token。 */
  temperature?: number;
  /** 核采样概率阈值。模型会考虑概率质量在 top_p 内的 token 结果。当取值为 0 时模型仅考虑对数概率最大的一个 token。 */
  top_p?: number;
  /** 模型回答最大长度（单位 token）。 */
  max_tokens?: number;
  /** 控制模型输出的最大长度（包括模型回答和模型思维链内容长度，单位 token）。 */
  max_contexts?: number;
};

export type PinAiIndexChatConversationReply = object;

export type PinAiIndexChatConversationReq = {
  /** 对话编号 */
  id: string;
  /** 是否置顶 */
  pinned?: boolean;
};

export type PinAiIndexChatConversationResponses = {
  /**
   * A successful response.
   */
  200: PinAiIndexChatConversationReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type PromptSetting = {
  /** 名称 */
  name?: string;
  /** 描述 */
  desc?: string;
  /** 提示词 */
  prompt?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateAiIndexChatConversationReply = object;

export type UpdateAiIndexChatConversationReq = {
  /** 对话编号 */
  id: string;
  /** 对话标题 */
  title: string;
  promptSetting?: PromptSetting;
  modelSetting?: ModelSetting;
  knowledgeSetting?: KnowledgeSetting;
  mcpSetting?: McpSetting;
};

export type UpdateAiIndexChatConversationResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiIndexChatConversationReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
