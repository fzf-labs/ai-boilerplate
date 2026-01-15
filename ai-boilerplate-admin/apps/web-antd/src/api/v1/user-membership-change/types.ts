/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetUserMembershipChangeListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 用户ID */
  userId: string;
  /** 来源类型(order,activation_code,admin) */
  sourceType: string;
  /** 来源ID(订单ID/激活码) */
  sourceId: string;
};

export type GetUserMembershipChangeListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: UserMembershipChangeInfo[];
};

export type GetUserMembershipChangeListResponses = {
  /**
   * A successful response.
   */
  200: GetUserMembershipChangeListReply;
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

export type UserMembershipChangeInfo = {
  /** id */
  id?: string;
  /** 用户ID */
  userId?: string;
  /** 来源类型(order,activation_code,admin) */
  sourceType?: string;
  /** 来源ID(订单ID/激活码) */
  sourceId?: string;
  before?: UserMembershipChangeItem;
  after?: UserMembershipChangeItem;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type UserMembershipChangeItem = {
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
