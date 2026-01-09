/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMallProductReply = {
  /** id */
  id?: string;
};

export type CreateMallProductReq = {
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType: string;
  /** 商品名称 */
  productName: string;
  /** 商品描述 */
  productDesc?: string;
  /** 商品图片(多个用逗号分隔) */
  productImages?: string[];
  /** 商品详情(JSON格式,包含特色功能等) */
  productDetail?: string[];
  productConfig?: ProductConfig;
  /** 原价 */
  originalPrice: number;
  /** 现价 */
  currentPrice: number;
  /** 库存数量(-1表示无限库存) */
  stockQuantity?: number;
  /** 已售数量 */
  soldQuantity?: number;
  /** 排序 */
  sort?: number;
  /** 状态(-1下架,0待上架,1在售,2售罄) */
  status: number;
};

export type CreateMallProductResponses = {
  /**
   * A successful response.
   */
  200: CreateMallProductReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMallProductReply = object;

export type DeleteMallProductReq = {
  /** id */
  id: string;
};

export type DeleteMallProductResponses = {
  /**
   * A successful response.
   */
  200: DeleteMallProductReply;
  /**
   * An unexpected error response.
   */
  default: Status;
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

export type GetMallProductListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType?: string;
  /** 商品名称 */
  productName?: string;
};

export type GetMallProductListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MallProductInfo[];
};

export type GetMallProductListResponses = {
  /**
   * A successful response.
   */
  200: GetMallProductListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMallProductSelectorParams = {
  /** 搜索商品名称 */
  searchName?: string;
};

export type GetMallProductSelectorReply = {
  /** 列表数据 */
  list?: MallProductSelector[];
};

export type GetMallProductSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetMallProductSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MallProductInfo = {
  /** id */
  id?: string;
  /** 商品类型(membership:会员,service:服务,goods:商品) */
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

export type MallProductSelector = {
  /** id */
  id?: string;
  /** 商品名称 */
  productName?: string;
};

export type MembershipConfig = {
  /** 会员类型编码(normal,vip,svip) */
  membershipType?: string;
  /** 时长天数 */
  duration_days?: number;
};

export type ProductConfig = {
  membership?: MembershipConfig;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMallProductReply = object;

export type UpdateMallProductReq = {
  /** id */
  id: string;
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType: string;
  /** 商品名称 */
  productName: string;
  /** 商品描述 */
  productDesc?: string;
  /** 商品图片(多个用逗号分隔) */
  productImages?: string[];
  /** 商品详情(JSON格式,包含特色功能等) */
  productDetail?: string[];
  productConfig?: ProductConfig;
  /** 原价 */
  originalPrice: number;
  /** 现价 */
  currentPrice: number;
  /** 库存数量(-1表示无限库存) */
  stockQuantity?: number;
  /** 已售数量 */
  soldQuantity?: number;
  /** 排序 */
  sort?: number;
  /** 状态(-1下架,0待上架,1在售,2售罄) */
  status: number;
};

export type UpdateMallProductResponses = {
  /**
   * A successful response.
   */
  200: UpdateMallProductReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMallProductStatusReply = object;

export type UpdateMallProductStatusReq = {
  /** id */
  id: string;
  /** 状态(-1下架,0待上架,1在售,2售罄) */
  status: number;
};

export type UpdateMallProductStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMallProductStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
