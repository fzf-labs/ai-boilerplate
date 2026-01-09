/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type ConfigDatumInfo = {
  /** id */
  id?: string;
  /** 名称 */
  name?: string;
  /** 健 */
  key?: string;
  /** 值 */
  value?: string;
  /** 备注 */
  remark?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type CreateConfigDatumReply = {
  /** id */
  id?: string;
};

export type CreateConfigDatumReq = {
  /** 名称 */
  name: string;
  /** 健 */
  key: string;
  /** 值 */
  value?: string;
  /** 备注 */
  remark?: string;
  /** 状态 */
  status?: number;
};

export type CreateConfigDatumResponses = {
  /**
   * A successful response.
   */
  200: CreateConfigDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteConfigDatumReply = object;

export type DeleteConfigDatumReq = {
  /** id */
  id: string;
};

export type DeleteConfigDatumResponses = {
  /**
   * A successful response.
   */
  200: DeleteConfigDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetConfigDatumInfoParams = {
  /** id */
  id: string;
};

export type GetConfigDatumInfoReply = {
  info?: ConfigDatumInfo;
};

export type GetConfigDatumInfoResponses = {
  /**
   * A successful response.
   */
  200: GetConfigDatumInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetConfigDatumListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 名称 */
  name?: string;
  /** 健 */
  key?: string;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetConfigDatumListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: ConfigDatumInfo[];
};

export type GetConfigDatumListResponses = {
  /**
   * A successful response.
   */
  200: GetConfigDatumListReply;
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

export type UpdateConfigDatumReply = object;

export type UpdateConfigDatumReq = {
  /** id */
  id: string;
  /** 名称 */
  name: string;
  /** 健 */
  key: string;
  /** 值 */
  value?: string;
  /** 备注 */
  remark?: string;
  /** 状态 */
  status?: number;
};

export type UpdateConfigDatumResponses = {
  /**
   * A successful response.
   */
  200: UpdateConfigDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateConfigDatumStatusReply = object;

export type UpdateConfigDatumStatusReq = {
  /** id */
  id: string;
  /** 状态 */
  status?: number;
};

export type UpdateConfigDatumStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateConfigDatumStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
