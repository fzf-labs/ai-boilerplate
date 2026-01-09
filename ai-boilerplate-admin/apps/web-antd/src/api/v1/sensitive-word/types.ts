/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSensitiveWordReply = {
  /** id */
  id?: string;
};

export type CreateSensitiveWordReq = {
  /** 敏感词 */
  word: string;
  /** 标签 */
  lab?: string;
};

export type CreateSensitiveWordResponses = {
  /**
   * A successful response.
   */
  200: CreateSensitiveWordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSensitiveWordReply = object;

export type DeleteSensitiveWordReq = {
  /** id */
  id: string;
};

export type DeleteSensitiveWordResponses = {
  /**
   * A successful response.
   */
  200: DeleteSensitiveWordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSensitiveWordInfoParams = {
  /** id */
  id: string;
};

export type GetSensitiveWordInfoReply = {
  info?: SensitiveWordInfo;
};

export type GetSensitiveWordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSensitiveWordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSensitiveWordLabsSelectorReply = {
  /** 标签 */
  list?: SensitiveWordLabsSelector[];
};

export type GetSensitiveWordLabsSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSensitiveWordLabsSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSensitiveWordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 敏感词 */
  word?: string;
  /** 标签 */
  lab?: string;
};

export type GetSensitiveWordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SensitiveWordInfo[];
};

export type GetSensitiveWordListResponses = {
  /**
   * A successful response.
   */
  200: GetSensitiveWordListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SensitiveWordInfo = {
  /** id */
  id?: string;
  /** 敏感词 */
  word?: string;
  /** 标签 */
  lab?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type SensitiveWordLabsSelector = {
  /** 标签键 */
  key?: string;
  /** 标签值 */
  value?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateSensitiveWordReply = object;

export type UpdateSensitiveWordReq = {
  /** id */
  id: string;
  /** 敏感词 */
  word: string;
  /** 标签 */
  lab?: string;
};

export type UpdateSensitiveWordResponses = {
  /**
   * A successful response.
   */
  200: UpdateSensitiveWordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
