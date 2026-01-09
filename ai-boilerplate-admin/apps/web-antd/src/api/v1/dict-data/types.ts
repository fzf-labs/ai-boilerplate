/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateDictDatumReply = {
  /** id */
  id?: string;
};

export type CreateDictDatumReq = {
  /** 字典类型 */
  type: string;
  /** 字典标签 */
  label: string;
  /** 字典键 */
  key: string;
  /** 字典值 */
  value: string;
  /** 备注 */
  remark?: string;
  /** css 颜色 */
  CSSColor?: string;
  /** css 样式 */
  CSSClass?: string;
  /** 状态（0正常 1停用） */
  status: number;
};

export type CreateDictDatumResponses = {
  /**
   * A successful response.
   */
  200: CreateDictDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteDictDatumReply = object;

export type DeleteDictDatumReq = {
  /** id集合 */
  ids: string[];
};

export type DeleteDictDatumResponses = {
  /**
   * A successful response.
   */
  200: DeleteDictDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DictDatumInfo = {
  /** id */
  id?: string;
  /** 字典类型 */
  type?: string;
  /** 字典标签 */
  label?: string;
  /** 字典键 */
  key?: string;
  /** 字典值 */
  value?: string;
  /** 备注 */
  remark?: string;
  /** css 颜色 */
  CSSColor?: string;
  /** css 样式 */
  CSSClass?: string;
  /** 状态（0正常 1停用） */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type GetDictDatumInfoParams = {
  /** id */
  id: string;
};

export type GetDictDatumInfoReply = {
  info?: DictDatumInfo;
};

export type GetDictDatumInfoResponses = {
  /**
   * A successful response.
   */
  200: GetDictDatumInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetDictDatumListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 字典类型 */
  type?: string;
};

export type GetDictDatumListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: DictDatumInfo[];
};

export type GetDictDatumListResponses = {
  /**
   * A successful response.
   */
  200: GetDictDatumListReply;
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

export type UpdateDictDatumReply = object;

export type UpdateDictDatumReq = {
  /** id */
  id: string;
  /** 字典类型 */
  type: string;
  /** 字典标签 */
  label: string;
  /** 字典键 */
  key: string;
  /** 字典值 */
  value: string;
  /** 备注 */
  remark?: string;
  /** css 颜色 */
  CSSColor?: string;
  /** css 样式 */
  CSSClass?: string;
  /** 状态（0正常 1停用） */
  status: number;
};

export type UpdateDictDatumResponses = {
  /**
   * A successful response.
   */
  200: UpdateDictDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateDictDatumStatusReply = object;

export type UpdateDictDatumStatusReq = {
  /** id */
  id: string;
  /** 状态（0正常 1停用） */
  status: number;
};

export type UpdateDictDatumStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateDictDatumStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
