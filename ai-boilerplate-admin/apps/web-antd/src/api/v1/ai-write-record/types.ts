/* eslint-disable */
// @ts-ignore

export type AiWriteRecordInfo = {
  /** 编号 */
  id?: string;
  /** 用户编号 */
  adminId?: string;
  /** 写作类型 */
  type?: number;
  /** 平台 */
  platform?: string;
  /** 模型编号 */
  modelId?: string;
  /** 模型 */
  model?: string;
  /** 生成内容提示 */
  prompt?: string;
  /** 生成的内容 */
  generatedContent?: string;
  /** 原文 */
  originalContent?: string;
  /** 长度提示词 */
  length?: number;
  /** 格式提示词 */
  format?: number;
  /** 语气提示词 */
  tone?: number;
  /** 语言提示词 */
  language?: number;
  /** 错误信息 */
  errorMessage?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateAiWriteRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiWriteRecordReq = {
  /** 用户编号 */
  adminId: string;
  /** 写作类型 */
  type?: number;
  /** 平台 */
  platform: string;
  /** 模型编号 */
  modelId: string;
  /** 模型 */
  model: string;
  /** 生成内容提示 */
  prompt: string;
  /** 生成的内容 */
  generatedContent?: string;
  /** 原文 */
  originalContent?: string;
  /** 长度提示词 */
  length?: number;
  /** 格式提示词 */
  format?: number;
  /** 语气提示词 */
  tone?: number;
  /** 语言提示词 */
  language?: number;
  /** 错误信息 */
  errorMessage?: string;
};

export type CreateAiWriteRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiWriteRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiWriteRecordReply = object;

export type DeleteAiWriteRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiWriteRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiWriteRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiWriteRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiWriteRecordInfoReply = {
  info?: AiWriteRecordInfo;
};

export type GetAiWriteRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiWriteRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiWriteRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiWriteRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiWriteRecordInfo[];
};

export type GetAiWriteRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiWriteRecordListReply;
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

export type UpdateAiWriteRecordReply = object;

export type UpdateAiWriteRecordReq = {
  /** 编号 */
  id: string;
  /** 用户编号 */
  adminId: string;
  /** 写作类型 */
  type?: number;
  /** 平台 */
  platform: string;
  /** 模型编号 */
  modelId: string;
  /** 模型 */
  model: string;
  /** 生成内容提示 */
  prompt: string;
  /** 生成的内容 */
  generatedContent?: string;
  /** 原文 */
  originalContent?: string;
  /** 长度提示词 */
  length?: number;
  /** 格式提示词 */
  format?: number;
  /** 语气提示词 */
  tone?: number;
  /** 语言提示词 */
  language?: number;
  /** 错误信息 */
  errorMessage?: string;
};

export type UpdateAiWriteRecordResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiWriteRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
