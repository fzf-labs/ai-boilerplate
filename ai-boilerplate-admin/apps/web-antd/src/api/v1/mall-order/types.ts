/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetMallOrderInfoParams = {
  /** id */
  id: string;
};

export type GetMallOrderInfoReply = {
  info?: MallOrderInfo;
};

export type GetMallOrderInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMallOrderInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMallOrderListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 用户ID */
  userId?: string;
  /** 订单ID */
  orderId?: string;
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType?: string;
  /** 商品ID */
  productId?: string;
};

export type GetMallOrderListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MallOrderInfo[];
};

export type GetMallOrderListResponses = {
  /**
   * A successful response.
   */
  200: GetMallOrderListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MallOrderInfo = {
  /** id */
  id?: string;
  /** 用户ID */
  userId?: string;
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType?: string;
  /** 商品ID */
  productId?: string;
  /** 原价 */
  originalAmount?: number;
  /** 优惠金额 */
  discountAmount?: number;
  /** 实付金额 */
  actualAmount?: number;
  /** 退款金额 */
  refundAmount?: number;
  /** 币种 */
  currency?: string;
  /** 支付方式(微信,支付宝) */
  paymentMethod?: string;
  /** 支付状态(0待支付,1已支付,2支付失败,3已退款) */
  paymentStatus?: number;
  /** 支付时间 */
  paymentTime?: string;
  /** 确认时间 */
  deliveryTime?: string;
  /** 订单过期时间 */
  expiredTime?: string;
  /** 备注 */
  remark?: string;
  /** 状态(待付款pendingPayment,待发货pendingDelivery,待收货pendingReceipt,已完成completed,已取消canceled,已退款refunded) */
  status?: string;
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
