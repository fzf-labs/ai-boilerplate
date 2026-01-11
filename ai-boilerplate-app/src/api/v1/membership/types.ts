/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetMembershipBenefitsParams = {
  /** 会员类型编码(可选,不传则返回当前用户会员等级的权益) */
  membership_type?: string;
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
  membership_type?: string;
  /** 会员类型名称 */
  membership_name?: string;
  /** 会员类型描述 */
  membership_description?: string;
  /** 会员状态(-1禁用,1正常) */
  status?: number;
  /** 到期时间(普通会员为空,表示永不过期) */
  expired_at?: string;
  /** 是否已过期 */
  is_expired?: boolean;
  /** 是否自动续费(0否,1是) */
  auto_renew?: number;
  /** 自动续费天数 */
  auto_renew_days?: number;
  /** 开通时间 */
  created_at?: string;
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
  benefit_key?: string;
  /** 权益名称 */
  benefit_name?: string;
  /** 权益描述 */
  benefit_desc?: string;
  /** 权益值 */
  benefit_value?: string;
  /** 权益次数 */
  benefit_num?: string;
  /** 排序 */
  sort?: number;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};
