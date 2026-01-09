/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type BatchGenerateMallActivationCodeReply = {
  /** 批次号 */
  batchNo?: string;
};

export type BatchGenerateMallActivationCodeReq = {
  /** 商品类型(membership:会员,service:服务) */
  productType: string;
  /** 商品ID */
  productId: string;
  /** 激活码有效期开始时间 */
  validSt: string;
  /** 激活码有效期截止时间 */
  validEd: string;
  /** 平台 */
  platform?: string;
  /** 备注 */
  remark?: string;
  /** 生成数量 */
  num: number;
};

export type BatchGenerateMallActivationCodeResponses = {
  /**
   * A successful response.
   */
  200: BatchGenerateMallActivationCodeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMallActivationCodeReply = object;

export type DeleteMallActivationCodeReq = {
  /** id */
  id: string;
};

export type DeleteMallActivationCodeResponses = {
  /**
   * A successful response.
   */
  200: DeleteMallActivationCodeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMallActivationCodeInfoParams = {
  /** id */
  id: string;
};

export type GetMallActivationCodeInfoReply = {
  info?: MallActivationCodeInfo;
};

export type GetMallActivationCodeInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMallActivationCodeInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMallActivationCodeListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 商品类型(membership:会员,service:服务) */
  productType?: string;
  /** 商品ID */
  productId?: string;
  /** 批次号 */
  batchNo?: string;
  /** 激活码 */
  code?: string;
  /** 用户ID */
  userId?: string;
  /** 平台 */
  platform?: string;
  /** 平台订单号 */
  platformOrderNo?: string;
  /** 状态(-2已退款,-1禁用,0库存,1已售出,2已激活,3已过期) */
  status?: number;
  /** 激活时间范围 */
  activatedAt?: string[];
  /** 创建时间范围 */
  createdAt?: string[];
};

export type GetMallActivationCodeListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MallActivationCodeInfo[];
};

export type GetMallActivationCodeListResponses = {
  /**
   * A successful response.
   */
  200: GetMallActivationCodeListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MallActivationCodeInfo = {
  /** id */
  id?: string;
  /** 商品类型(membership:会员,service:服务) */
  productType?: string;
  /** 商品ID */
  productId?: string;
  /** 批次号 */
  batchNo?: string;
  /** 激活码 */
  code?: string;
  /** 激活码有效期开始时间 */
  validSt?: string;
  /** 激活码有效期截止时间 */
  validEd?: string;
  /** 激活时间 */
  activatedAt?: string;
  /** 用户ID */
  userId?: string;
  userChange?: UserChange;
  /** 平台 */
  platform?: string;
  /** 平台售出时间 */
  platformSoldAt?: string;
  /** 平台订单号 */
  platformOrderNo?: string;
  /** 平台买家ID */
  platformBuyerId?: string;
  /** 平台买家昵称 */
  platformBuyerName?: string;
  /** 备注 */
  remark?: string;
  /** 状态(-2已退款,-1禁用,0库存,1已售出,2已激活,3已过期) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 商品名称 */
  productName?: string;
  /** 用户昵称 */
  userNickname?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMallActivationCodeReply = object;

export type UpdateMallActivationCodeReq = {
  /** id */
  id: string;
  /** 平台 */
  platform?: string;
  /** 平台售出时间 */
  platformSoldAt?: string;
  /** 平台订单号 */
  platformOrderNo?: string;
  /** 平台买家ID */
  platformBuyerId?: string;
  /** 平台买家昵称 */
  platformBuyerName?: string;
  /** 备注 */
  remark?: string;
};

export type UpdateMallActivationCodeResponses = {
  /**
   * A successful response.
   */
  200: UpdateMallActivationCodeReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMallActivationCodeStatusReply = object;

export type UpdateMallActivationCodeStatusReq = {
  /** id */
  id: string;
  /** 状态(-2已退款,-1禁用,0库存,1已售出,2已激活,3已过期) */
  status: number;
};

export type UpdateMallActivationCodeStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMallActivationCodeStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UserChange = {
  userMembershipChange?: UserMembershipChange;
};

export type UserMembershipChange = {
  before?: UserMembershipChangeItem;
  after?: UserMembershipChangeItem;
};

export type UserMembershipChangeItem = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 到期时间(普通会员为NULL,表示永不过期) */
  expiredAt?: string;
  /** 状态(-1禁用,1正常) */
  status?: number;
};
