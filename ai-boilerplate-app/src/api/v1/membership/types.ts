/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetMembershipBenefitsParams = {
  /** 会员类型编码(可选,不传则返回当前用户会员等级的权益) */
  membershipType?: string;
};

export type GetMembershipBenefitsReply = {
  /** 权益列表 */
  benefits?: MembershipBenefit[];
};

export type GetMembershipBenefitsResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitsReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserMembershipInfoReply = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 会员类型名称 */
  membershipName?: string;
  /** 会员类型描述 */
  membershipDescription?: string;
  /** 会员状态(-1禁用,1正常) */
  status?: number;
  /** 到期时间(普通会员为空,表示永不过期) */
  expiredAt?: string;
  /** 是否已过期 */
  isExpired?: boolean;
  /** 是否自动续费(0否,1是) */
  autoRenew?: number;
  /** 自动续费天数 */
  autoRenewDays?: number;
  /** 开通时间 */
  createdAt?: string;
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

export type MembershipBenefit = {
  /** 权益标识 */
  benefitKey?: string;
  /** 权益名称 */
  benefitName?: string;
  /** 权益描述 */
  benefitDesc?: string;
  /** 权益值 */
  benefitValue?: string;
  /** 权益次数 */
  benefitNum?: string;
  /** 排序 */
  sort?: number;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};
