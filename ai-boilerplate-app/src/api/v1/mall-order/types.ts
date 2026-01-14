/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetOrderInfoParams = {
  /** 订单ID */
  orderId: string;
};

export type GetOrderInfoReply = {
  info?: MallOrderInfo;
};

export type GetOrderInfoResponses = {
  /**
   * A successful response.
   */
  200: GetOrderInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetPaymentInfoReply = {
  /** 订单ID */
  orderId?: string;
  /** 支付方式 */
  paymentMethod?: string;
  /** 实付金额 */
  actualAmount?: number;
  /** 支付跳转链接(H5/APP支付) */
  paymentUrl?: string;
  /** 支付二维码链接(扫码支付) */
  qrCodeUrl?: string;
  /** 预支付交易会话标识(微信小程序/APP支付) */
  prepayId?: string;
  /** 支付过期时间 */
  expiredTime?: string;
  /** 应用ID(微信支付需要) */
  appId?: string;
  /** 时间戳(微信支付需要) */
  timeStamp?: string;
  /** 随机字符串(微信支付需要) */
  nonceStr?: string;
  /** 签名类型(微信支付需要) */
  signType?: string;
  /** 签名(微信支付需要) */
  paySign?: string;
};

export type GetPaymentInfoReq = {
  /** 订单ID */
  orderId: string;
  /** 支付方式(wechat:微信,alipay:支付宝) */
  paymentMethod: string;
};

export type GetPaymentInfoResponses = {
  /**
   * A successful response.
   */
  200: GetPaymentInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserOrderListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 订单状态过滤(可选) */
  status?: string;
  /** 支付状态过滤(可选) */
  paymentStatus?: number;
};

export type GetUserOrderListReply = {
  /** 总数 */
  total?: number;
  /** 订单列表 */
  list?: MallOrderInfo[];
};

export type GetUserOrderListResponses = {
  /**
   * A successful response.
   */
  200: GetUserOrderListReply;
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

export type PaymentCallbackReply = {
  /** 是否处理成功 */
  success?: boolean;
  /** 处理消息 */
  message?: string;
};

export type PaymentCallbackReq = {
  /** 订单ID */
  orderId?: string;
  /** 支付方式(wechat:微信,alipay:支付宝) */
  paymentMethod?: string;
  /** 第三方支付交易ID */
  transactionId?: string;
  /** 支付状态(1已支付,2支付失败) */
  paymentStatus?: number;
  /** 回调原始数据(JSON字符串) */
  callbackData?: string;
};

export type PaymentCallbackResponses = {
  /**
   * A successful response.
   */
  200: PaymentCallbackReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type PlaceOrderReply = {
  /** 订单ID */
  orderId?: string;
  /** 实付金额 */
  actualAmount?: number;
  /** 支付信息(预留,如支付二维码、支付链接等) */
  paymentInfo?: string;
};

export type PlaceOrderReq = {
  /** 商品类型(membership:会员,service:服务,goods:商品) */
  productType: string;
  /** 商品ID */
  productId: string;
  /** 支付方式(wechat:微信,alipay:支付宝) */
  paymentMethod?: string;
  /** 备注 */
  remark?: string;
};

export type PlaceOrderResponses = {
  /**
   * A successful response.
   */
  200: PlaceOrderReply;
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
