/* eslint-disable */
// @ts-ignore

export type AiProviderPlatformInfo = {
  /** 编号 */
  id?: string;
  /** 平台 */
  platform?: string;
  /** 名称 */
  name?: string;
  /** API 地址 */
  APIURL?: string;
  /** API KEY */
  APIKey?: string;
  /** 文档地址 */
  docURL?: string;
  /** 排序 */
  sort?: number;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type AiProviderPlatformSelectorItem = {
  /** 标签 */
  label?: string;
  /** 值 */
  value?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateAiProviderPlatformReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiProviderPlatformReq = {
  /** 平台 */
  platform: string;
  /** 名称 */
  name: string;
  /** API 地址 */
  APIURL?: string;
  /** API KEY */
  APIKey?: string;
  /** 文档地址 */
  docURL?: string;
  /** 排序 */
  sort?: number;
  /** 状态 */
  status: number;
};

export type CreateAiProviderPlatformResponses = {
  /**
   * A successful response.
   */
  200: CreateAiProviderPlatformReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiProviderPlatformReply = object;

export type DeleteAiProviderPlatformReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiProviderPlatformResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiProviderPlatformReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderPlatformInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiProviderPlatformInfoReply = {
  info?: AiProviderPlatformInfo;
};

export type GetAiProviderPlatformInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderPlatformInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderPlatformListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  platform?: string;
  status?: number;
};

export type GetAiProviderPlatformListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiProviderPlatformInfo[];
};

export type GetAiProviderPlatformListResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderPlatformListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiProviderPlatformSelectorReply = {
  /** 列表数据 */
  list?: AiProviderPlatformSelectorItem[];
};

export type GetAiProviderPlatformSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetAiProviderPlatformSelectorReply;
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

export type UpdateAiProviderPlatformReply = object;

export type UpdateAiProviderPlatformReq = {
  /** 编号 */
  id: string;
  /** 平台 */
  platform: string;
  /** 名称 */
  name: string;
  /** API 地址 */
  APIURL?: string;
  /** API KEY */
  APIKey?: string;
  /** 文档地址 */
  docURL?: string;
  /** 排序 */
  sort?: number;
  /** 状态 */
  status: number;
};

export type UpdateAiProviderPlatformResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiProviderPlatformReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateAiProviderPlatformStatusReply = object;

export type UpdateAiProviderPlatformStatusReq = {
  /** 编号 */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateAiProviderPlatformStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiProviderPlatformStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
