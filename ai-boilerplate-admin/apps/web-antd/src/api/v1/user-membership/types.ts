/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateUserMembershipReply = {
  /** id */
  id?: string;
};

export type CreateUserMembershipReq = {
  /** 用户ID */
  userId: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType: string;
  /** 到期时间(普通会员为NULL,表示永不过期) */
  expiredAt?: string;
  /** 是否自动续费(0否,1是) */
  autoRenew?: number;
  /** 自动续费天数 */
  autoRenewDays?: number;
  /** 状态(-1禁用,1正常) */
  status: number;
};

export type CreateUserMembershipResponses = {
  /**
   * A successful response.
   */
  200: CreateUserMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteUserMembershipReply = object;

export type DeleteUserMembershipReq = {
  /** id */
  id: string;
};

export type DeleteUserMembershipResponses = {
  /**
   * A successful response.
   */
  200: DeleteUserMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserMembershipInfoByUserIdParams = {
  /** 用户ID */
  userId: string;
};

export type GetUserMembershipInfoByUserIdReply = {
  info?: UserMembershipInfo;
};

export type GetUserMembershipInfoByUserIdResponses = {
  /**
   * A successful response.
   */
  200: GetUserMembershipInfoByUserIdReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserMembershipInfoParams = {
  /** id */
  id: string;
};

export type GetUserMembershipInfoReply = {
  info?: UserMembershipInfo;
};

export type GetUserMembershipInfoResponses = {
  /**
   * A successful response.
   */
  200: GetUserMembershipInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserMembershipListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetUserMembershipListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: UserMembershipInfo[];
};

export type GetUserMembershipListResponses = {
  /**
   * A successful response.
   */
  200: GetUserMembershipListReply;
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

export type UpdateUserMembershipReply = object;

export type UpdateUserMembershipReq = {
  /** id */
  id: string;
  /** 用户ID */
  userId: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType: string;
  /** 到期时间(普通会员为NULL,表示永不过期) */
  expiredAt?: string;
  /** 是否自动续费(0否,1是) */
  autoRenew?: number;
  /** 自动续费天数 */
  autoRenewDays?: number;
  /** 状态(-1禁用,1正常) */
  status: number;
};

export type UpdateUserMembershipResponses = {
  /**
   * A successful response.
   */
  200: UpdateUserMembershipReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateUserMembershipStatusReply = object;

export type UpdateUserMembershipStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1正常) */
  status: number;
};

export type UpdateUserMembershipStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateUserMembershipStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UserMembershipInfo = {
  /** id */
  id?: string;
  /** 用户ID */
  userId?: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 到期时间(普通会员为NULL,表示永不过期) */
  expiredAt?: string;
  /** 是否自动续费(0否,1是) */
  autoRenew?: number;
  /** 自动续费天数 */
  autoRenewDays?: number;
  /** 状态(-1禁用,1正常) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
