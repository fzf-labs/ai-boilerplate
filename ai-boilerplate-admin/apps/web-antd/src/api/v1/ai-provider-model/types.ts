/* eslint-disable */
// @ts-ignore

export type AiProviderModelInfo = {
  /** 编号 */
  id?: string;
  /** 平台编号 */
  platformId?: string;
  /** 模型类型 */
  modelType?: string;
  /** 模型id */
  modelId?: string;
  /** 模型名字 */
  modelName?: string;
  modelConfig?: ModelConfig;
  /** 排序 */
  sort?: number;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type AiProviderModelLabelSelectorItem = {
  /** 标签 */
  label?: string;
  /** 值 */
  value?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateAiProviderModelReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiProviderModelReq = {
  /** 平台编号 */
  platformId: string;
  /** 模型类型 */
  modelType: string;
  /** 模型id */
  modelId: string;
  /** 模型名字 */
  modelName: string;
  modelConfig?: ModelConfig;
  /** 排序 */
  sort: number;
  /** 状态 */
  status: number;
};

export type CreateAiProviderModelResponses = {
  /**
   * A successful response.
   */
  200: CreateAiProviderModelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiProviderModelReply = object;

export type DeleteAiProviderModelReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiProviderModelResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiProviderModelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderModelInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiProviderModelInfoReply = {
  info?: AiProviderModelInfo;
};

export type GetAiProviderModelInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderModelInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderModelLabelSelectorReply = {
  /** 列表数据 */
  list?: AiProviderModelLabelSelectorItem[];
};

export type GetAiProviderModelLabelSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderModelLabelSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderModelListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiProviderModelListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiProviderModelInfo[];
};

export type GetAiProviderModelListResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderModelListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type ModelConfig = object;

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateAiProviderModelReply = object;

export type UpdateAiProviderModelReq = {
  /** 编号 */
  id: string;
  /** 平台编号 */
  platformId: string;
  /** 模型类型 */
  modelType: string;
  /** 模型id */
  modelId: string;
  /** 模型名字 */
  modelName: string;
  modelConfig?: ModelConfig;
  /** 排序 */
  sort: number;
  /** 状态 */
  status: number;
};

export type UpdateAiProviderModelResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiProviderModelReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateAiProviderModelStatusReply = object;

export type UpdateAiProviderModelStatusReq = {
  /** 编号 */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateAiProviderModelStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiProviderModelStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
