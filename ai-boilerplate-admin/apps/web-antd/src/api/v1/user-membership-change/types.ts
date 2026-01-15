/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetUserMembershipChangeListParams = {
  /** 用户ID */
  userId: string;
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
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
  /** 变更前会员类型 */
  beforeMembershipType?: string;
  /** 变更后会员类型 */
  afterMembershipType?: string;
  /** 变更前到期时间 */
  beforeExpiredAt?: string;
  /** 变更后到期时间 */
  afterExpiredAt?: string;
  /** 变更时长(天) */
  durationDays?: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
