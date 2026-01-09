/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMembershipReply = {
  /** id */
  id?: string;
};

export type CreateMembershipReq = {
  /** 会员类型名称 */
  name: string;
  /** 会员类型编码(normal,vip,svip) */
  type: string;
  /** 会员类型描述 */
  description?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type CreateMembershipResponses = {
  /**
   * A successful response.
   */
  200: CreateMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMembershipReply = object;

export type DeleteMembershipReq = {
  /** id */
  id: string;
};

export type DeleteMembershipResponses = {
  /**
   * A successful response.
   */
  200: DeleteMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipInfoParams = {
  /** id */
  id: string;
};

export type GetMembershipInfoReply = {
  info?: MembershipInfo;
};

export type GetMembershipInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 会员类型名称 */
  name?: string;
  /** 会员类型编码(normal,vip,svip) */
  type?: string;
  /** 状态(-1禁用,1启用) */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetMembershipListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MembershipInfo[];
};

export type GetMembershipListResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MembershipInfo = {
  /** id */
  id?: string;
  /** 会员类型名称 */
  name?: string;
  /** 会员类型编码(normal,vip,svip) */
  type?: string;
  /** 会员类型描述 */
  description?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMembershipReply = object;

export type UpdateMembershipReq = {
  /** id */
  id: string;
  /** 会员类型名称 */
  name: string;
  /** 会员类型编码(normal,vip,svip) */
  type: string;
  /** 会员类型描述 */
  description?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMembershipStatusReply = object;

export type UpdateMembershipStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
