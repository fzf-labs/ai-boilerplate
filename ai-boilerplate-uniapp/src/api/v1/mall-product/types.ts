/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetMallProductInfoParams = {
  /** id */
  id: string;
};

export type GetMallProductInfoReply = {
  info?: MallProductInfo;
};

export type GetMallProductInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMallProductInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMembershipProductListReply = {
  /** 列表数据 */
  list?: MallProductInfo[];
};

export type GetMembershipProductListResponses = {
  /**
   * A successful response.
   */
  200: GetMembershipProductListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MallProductInfo = {
  /** id */
  id?: string;
  /** 商品类型(membership:会员,service:增值服务,goods:商品) */
  productType?: string;
  /** 商品名称 */
  productName?: string;
  /** 商品描述 */
  productDesc?: string;
  /** 商品图片(多个用逗号分隔) */
  productImages?: string[];
  /** 商品详情(JSON格式,包含特色功能等) */
  productDetail?: string[];
  productConfig?: ProductConfig;
  /** 原价 */
  originalPrice?: number;
  /** 现价 */
  currentPrice?: number;
  /** 库存数量(-1表示无限库存) */
  stockQuantity?: number;
  /** 已售数量 */
  soldQuantity?: number;
  /** 排序 */
  sort?: number;
  /** 状态(-1下架,0待上架,1在售,2售罄) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type MembershipConfig = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 时长天数 */
  durationDays?: number;
};

export type ProductConfig = {
  membership?: MembershipConfig;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};
