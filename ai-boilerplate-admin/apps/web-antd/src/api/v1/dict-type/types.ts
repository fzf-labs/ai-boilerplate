/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateDictTypeReply = {
  /** id */
  id?: string;
};

export type CreateDictTypeReq = {
  /** 字典名称 */
  name: string;
  /** 字典类型 */
  type: string;
  /** 状态（1正常-1停用） */
  status: number;
  /** 备注 */
  remark?: string;
};

export type CreateDictTypeResponses = {
  /**
   * A successful response.
   */
  200: CreateDictTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteDictTypeReply = object;

export type DeleteDictTypeReq = {
  /** id集合 */
  ids: string[];
};

export type DeleteDictTypeResponses = {
  /**
   * A successful response.
   */
  200: DeleteDictTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DictTypeInfo = {
  /** id */
  id?: string;
  /** 字典名称 */
  name?: string;
  /** 字典类型 */
  type?: string;
  /** 状态（1正常-1停用） */
  status?: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type DictTypeSelector = {
  /** 字典类型 */
  type?: string;
  /** 字典名称 */
  name?: string;
};

export type GetDictTypeInfoParams = {
  /** id */
  id: string;
};

export type GetDictTypeInfoReply = {
  info?: DictTypeInfo;
};

export type GetDictTypeInfoResponses = {
  /**
   * A successful response.
   */
  200: GetDictTypeInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetDictTypeListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 字典名称 */
  name?: string;
  /** 字典类型 */
  type?: string;
  /** 状态（1正常-1停用） */
  status?: number;
};

export type GetDictTypeListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: DictTypeInfo[];
};

export type GetDictTypeListResponses = {
  /**
   * A successful response.
   */
  200: GetDictTypeListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetDictTypeSelectorParams = {
  /** 字典类型 */
  name?: string;
};

export type GetDictTypeSelectorReply = {
  /** 列表数据 */
  list?: DictTypeSelector[];
};

export type GetDictTypeSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetDictTypeSelectorReply;
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

export type UpdateDictTypeReply = object;

export type UpdateDictTypeReq = {
  /** id */
  id: string;
  /** 字典名称 */
  name: string;
  /** 字典类型 */
  type: string;
  /** 状态（1正常-1停用） */
  status: number;
  /** 备注 */
  remark?: string;
};

export type UpdateDictTypeResponses = {
  /**
   * A successful response.
   */
  200: UpdateDictTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateDictTypeStatusReply = object;

export type UpdateDictTypeStatusReq = {
  /** id */
  id: string;
  /** 状态（1正常-1停用） */
  status: number;
};

export type UpdateDictTypeStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateDictTypeStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
