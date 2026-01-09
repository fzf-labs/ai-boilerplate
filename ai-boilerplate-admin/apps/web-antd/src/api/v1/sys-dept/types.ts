/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysDeptReply = {
  /** 编号 */
  id?: string;
};

export type CreateSysDeptReq = {
  /** 父级id */
  pid?: string;
  /** 部门简称 */
  name: string;
  /** 负责人 */
  adminId?: string;
  /** 排序值 */
  sort?: number;
  /** -1=禁用 1=开启 */
  status?: number;
};

export type CreateSysDeptResponses = {
  /**
   * A successful response.
   */
  200: CreateSysDeptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysDeptReply = object;

export type DeleteSysDeptReq = {
  /** 编号集合 */
  id: string;
};

export type DeleteSysDeptResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysDeptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysDeptInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSysDeptInfoReply = {
  info?: SysDeptInfo;
};

export type GetSysDeptInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysDeptInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysDeptListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysDeptInfo[];
};

export type GetSysDeptListResponses = {
  /**
   * A successful response.
   */
  200: GetSysDeptListReply;
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

export type SysDeptInfo = {
  /** 编号 */
  id?: string;
  /** 父级id */
  pid?: string;
  /** 部门简称 */
  name?: string;
  /** 负责人id */
  adminId?: string;
  /** -1=禁用 1=开启 */
  status?: number;
  /** 排序值 */
  sort?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 负责人名称 */
  adminName?: string;
};

export type UpdateSysDeptReply = object;

export type UpdateSysDeptReq = {
  /** 编号 */
  id: string;
  /** 父级id */
  pid?: string;
  /** 部门简称 */
  name: string;
  /** 负责人id */
  adminId?: string;
  /** 排序值 */
  sort?: number;
  /** -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysDeptResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysDeptReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysDeptStatusReply = object;

export type UpdateSysDeptStatusReq = {
  /** 编号 */
  id: string;
  /** -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysDeptStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysDeptStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
