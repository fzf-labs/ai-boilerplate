/* eslint-disable */
// @ts-ignore

export type ActivateMembershipByCodeReply = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 到期时间(普通会员为空,表示永不过期) */
  expiredAt?: string;
};

export type ActivateMembershipByCodeReq = {
  /** 激活码 */
  code: string;
};

export type ActivateMembershipByCodeResponses = {
  /**
   * A successful response.
   */
  200: ActivateMembershipByCodeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type ActivationCodeRedemptionInfo = {
  /** 激活码 */
  code?: string;
  /** 激活时间 */
  activatedAt?: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 到期时间(普通会员为空,表示永不过期) */
  expiredAt?: string;
  /** 有效时长(天) */
  durationDays?: number;
};

export type Any = {
  '@type'?: string;
};

export type ListActivationCodeRedemptionsParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type ListActivationCodeRedemptionsReply = {
  /** 总数 */
  total?: number;
  /** 兑换记录 */
  list?: ActivationCodeRedemptionInfo[];
};

export type ListActivationCodeRedemptionsResponses = {
  /**
   * A successful response.
   */
  200: ListActivationCodeRedemptionsReply;
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
