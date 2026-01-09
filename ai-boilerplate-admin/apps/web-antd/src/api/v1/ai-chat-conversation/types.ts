/* eslint-disable */
// @ts-ignore

export type AiChatConversationInfo = {
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

export type Any = {
  '@type'?: string;
};

export type DeleteAiChatConversationReply = object;

export type DeleteAiChatConversationReq = {
  /** 对话编号 */
  id: string;
};

export type DeleteAiChatConversationResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiChatConversationReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiChatConversationInfoParams = {
  /** 对话编号 */
  id: string;
};

export type GetAiChatConversationInfoReply = {
  info?: AiChatConversationInfo;
};

export type GetAiChatConversationInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiChatConversationInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiChatConversationListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiChatConversationListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiChatConversationInfo[];
};

export type GetAiChatConversationListResponses = {
  /**
   * A successful response.
   */
  200: GetAiChatConversationListReply;
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
