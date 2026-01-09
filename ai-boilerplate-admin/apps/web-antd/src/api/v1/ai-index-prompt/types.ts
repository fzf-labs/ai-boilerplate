/* eslint-disable */
// @ts-ignore

export type AiIndexPromptInfo = {
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

export type CreateAiIndexPromptReply = {
  /** Id */
  id?: string;
};

export type CreateAiIndexPromptReq = {
  /** 名称 */
  name: string;
  /** 描述 */
  desc: string;
  /** 提示词 */
  prompt: string;
  /** 排序 */
  sort: number;
};

export type CreateAiIndexPromptResponses = {
  /**
   * A successful response.
   */
  200: CreateAiIndexPromptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiIndexPromptReply = object;

export type DeleteAiIndexPromptReq = {
  /** Id */
  id: string;
};

export type DeleteAiIndexPromptResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiIndexPromptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexPromptInfoParams = {
  /** Id */
  id: string;
};

export type GetAiIndexPromptInfoReply = {
  info?: AiIndexPromptInfo;
};

export type GetAiIndexPromptInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexPromptInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexPromptListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexPromptListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexPromptInfo[];
};

export type GetAiIndexPromptListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexPromptListReply;
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

export type UpdateAiIndexPromptReply = object;

export type UpdateAiIndexPromptReq = {
  /** Id */
  id: string;
  /** 名称 */
  name: string;
  /** 描述 */
  desc: string;
  /** 提示词 */
  prompt: string;
  /** 排序 */
  sort: number;
};

export type UpdateAiIndexPromptResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiIndexPromptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
