/* eslint-disable */
// @ts-ignore

export type AiPromptInfo = {
  /** Id */
  id?: string;
  /** 租户编号 */
  tenantId?: string;
  /** 用户编号 */
  adminId?: string;
  /** 名称 */
  name?: string;
  /** 描述 */
  desc?: string;
  /** 提示词 */
  prompt?: string;
  /** 排序 */
  sort?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type GetAiPromptInfoParams = {
  /** Id */
  id: string;
};

export type GetAiPromptInfoReply = {
  info?: AiPromptInfo;
};

export type GetAiPromptInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiPromptInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiPromptListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiPromptListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiPromptInfo[];
};

export type GetAiPromptListResponses = {
  /**
   * A successful response.
   */
  200: GetAiPromptListReply;
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
