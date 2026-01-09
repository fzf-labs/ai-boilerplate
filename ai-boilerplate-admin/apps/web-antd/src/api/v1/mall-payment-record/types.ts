/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetMallPaymentRecordListByOrderIdParams = {
  /** 订单ID */
  orderId: string;
};

export type GetMallPaymentRecordListByOrderIdReply = {
  /** 列表数据 */
  list?: MallPaymentRecordInfo[];
};

export type GetMallPaymentRecordListByOrderIdResponses = {
  /**
   * A successful response.
   */
  200: GetMallPaymentRecordListByOrderIdReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMallPaymentRecordSuccessByOrderIdParams = {
  /** id */
  orderId: string;
};

export type GetMallPaymentRecordSuccessByOrderIdReply = {
  info?: MallPaymentRecordInfo;
};

export type GetMallPaymentRecordSuccessByOrderIdResponses = {
  /**
   * A successful response.
   */
  200: GetMallPaymentRecordSuccessByOrderIdReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MallPaymentRecordInfo = {
  /** id */
  id?: string;
  /** 订单ID */
  orderId?: string;
  /** 交易流水号 */
  transactionId?: string;
  /** 支付渠道(wechat,alipay) */
  paymentChannel?: string;
  /** 支付方式(mini_program,h5,native,jsapi) */
  paymentMethod?: string;
  /** 支付金额 */
  amount?: number;
  /** 币种 */
  currency?: string;
  /** 支付状态(0待支付,1支付成功,2支付失败,3已退款) */
  paymentStatus?: number;
  /** 第三方订单号 */
  thirdPartyOrderNo?: string;
  /** 第三方交易号 */
  thirdPartyTransactionId?: string;
  /** 回调数据 */
  callbackData?: string;
  /** 回调时间 */
  callbackTime?: string;
  /** 错误代码 */
  errorCode?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 状态(-1无效,1正常) */
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
