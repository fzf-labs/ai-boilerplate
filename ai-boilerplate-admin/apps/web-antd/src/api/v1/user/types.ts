/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateUserReply = {
  /** id */
  id?: string;
};

export type CreateUserReq = {
  /** 手机 */
  phone: string;
  /** 昵称 */
  nickname?: string;
  /** 性别（0未知 1男 2女） */
  gender?: number;
  /** 头像 */
  avatar?: string;
  /** 简介 */
  profile?: string;
  /** 状态 */
  status: number;
};

export type CreateUserResponses = {
  /**
   * A successful response.
   */
  200: CreateUserReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteUserReply = object;

export type DeleteUserReq = {
  /** id */
  id: string;
};

export type DeleteUserResponses = {
  /**
   * A successful response.
   */
  200: DeleteUserReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserInfoParams = {
  /** id */
  id: string;
};

export type GetUserInfoReply = {
  info?: UserInfo;
};

export type GetUserInfoResponses = {
  /**
   * A successful response.
   */
  200: GetUserInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 昵称 */
  nickname?: string;
  /** 手机 */
  phone?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetUserListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: UserInfo[];
};

export type GetUserListResponses = {
  /**
   * A successful response.
   */
  200: GetUserListReply;
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

export type UpdateUserReply = object;

export type UpdateUserReq = {
  /** id */
  id: string;
  /** 昵称 */
  nickname?: string;
  /** 性别（0未知 1男 2女） */
  gender?: number;
  /** 头像 */
  avatar?: string;
  /** 简介 */
  profile?: string;
  /** 状态 */
  status?: number;
};

export type UpdateUserResponses = {
  /**
   * A successful response.
   */
  200: UpdateUserReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateUserStatusReply = object;

export type UpdateUserStatusReq = {
  /** id */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateUserStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateUserStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UserInfo = {
  /** id */
  id?: string;
  /** 手机 */
  phone?: string;
  /** 昵称 */
  nickname?: string;
  /** 性别（0未知 1男 2女） */
  gender?: number;
  /** 头像 */
  avatar?: string;
  /** 简介 */
  profile?: string;
  /** 公众号用户Id */
  wxGzhUserId?: string;
  /** 小程序用户Id */
  wxGzhXcxId?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  userMembershipInfo?: UserMembershipInfo;
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
