/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMembershipBenefitReply = {
  /** id */
  id?: string;
};

export type CreateMembershipBenefitReq = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType: string;
  /** 权益标识 */
  benefitKey: string;
  /** 权益名称 */
  benefitName: string;
  /** 权益描述 */
  benefitDesc?: string;
  /** 权益值 */
  benefitValue?: string;
  /** 权益次数 */
  benefitNum?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type CreateMembershipBenefitResponses = {
  /**
   * A successful response.
   */
  200: CreateMembershipBenefitReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMembershipBenefitReply = object;

export type DeleteMembershipBenefitReq = {
  /** id */
  id: string;
};

export type DeleteMembershipBenefitResponses = {
  /**
   * A successful response.
   */
  200: DeleteMembershipBenefitReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipBenefitInfoParams = {
  /** id */
  id: string;
};

export type GetMembershipBenefitInfoReply = {
  info?: MembershipBenefitInfo;
};

export type GetMembershipBenefitInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipBenefitKeySelectReply = {
  /** 列表数据 */
  list?: MembershipBenefitKeySelect[];
};

export type GetMembershipBenefitKeySelectResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitKeySelectReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipBenefitListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 会员类型编码(normal,vip,svip) */
  membershipType: string;
  /** 权益标识 */
  benefitKey?: string;
  /** 权益名称 */
  benefitName?: string;
};

export type GetMembershipBenefitListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MembershipBenefitInfo[];
};

export type GetMembershipBenefitListResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MembershipBenefitInfo = {
  /** id */
  id?: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
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
  /** 状态(-1禁用,1启用) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type MembershipBenefitKeySelect = {
  /** 权益标识 */
  key?: string;
  /** 权益名称 */
  name?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMembershipBenefitReply = object;

export type UpdateMembershipBenefitReq = {
  /** id */
  id: string;
  /** 会员类型编码(normal,vip,svip) */
  membershipType: string;
  /** 权益标识 */
  benefitKey: string;
  /** 权益名称 */
  benefitName: string;
  /** 权益描述 */
  benefitDesc?: string;
  /** 权益值 */
  benefitValue?: string;
  /** 权益次数 */
  benefitNum?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipBenefitResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipBenefitReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMembershipBenefitStatusReply = object;

export type UpdateMembershipBenefitStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipBenefitStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipBenefitStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
