/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMembershipBenefitTypeReply = {
  /** id */
  id?: string;
};

export type CreateMembershipBenefitTypeReq = {
  /** 权益标识 */
  benefitKey: string;
  /** 权益名称 */
  benefitName: string;
  /** 权益图标 */
  benefitIcon?: string;
  /** 权益描述 */
  benefitDesc?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type CreateMembershipBenefitTypeResponses = {
  /**
   * A successful response.
   */
  200: CreateMembershipBenefitTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMembershipBenefitTypeReply = object;

export type DeleteMembershipBenefitTypeReq = {
  /** id */
  id: string;
};

export type DeleteMembershipBenefitTypeResponses = {
  /**
   * A successful response.
   */
  200: DeleteMembershipBenefitTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipBenefitTypeInfoParams = {
  /** id */
  id: string;
};

export type GetMembershipBenefitTypeInfoReply = {
  info?: MembershipBenefitTypeInfo;
};

export type GetMembershipBenefitTypeInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitTypeInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipBenefitTypeListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 权益标识 */
  benefitKey?: string;
  /** 权益名称 */
  benefitName?: string;
};

export type GetMembershipBenefitTypeListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MembershipBenefitTypeInfo[];
};

export type GetMembershipBenefitTypeListResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipBenefitTypeListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MembershipBenefitTypeInfo = {
  /** id */
  id?: string;
  /** 权益标识 */
  benefitKey?: string;
  /** 权益名称 */
  benefitName?: string;
  /** 权益图标 */
  benefitIcon?: string;
  /** 权益描述 */
  benefitDesc?: string;
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

export type UpdateMembershipBenefitTypeReply = object;

export type UpdateMembershipBenefitTypeReq = {
  /** id */
  id: string;
  /** 权益标识 */
  benefitKey: string;
  /** 权益名称 */
  benefitName: string;
  /** 权益图标 */
  benefitIcon?: string;
  /** 权益描述 */
  benefitDesc?: string;
  /** 排序 */
  sort?: number;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipBenefitTypeResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipBenefitTypeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMembershipBenefitTypeStatusReply = object;

export type UpdateMembershipBenefitTypeStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1启用) */
  status: number;
};

export type UpdateMembershipBenefitTypeStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMembershipBenefitTypeStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
