/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 查询订单接口-单条订单查询 返回值: An unexpected error response. GET /app/v1/mall_order/order/info */
export function getOrderInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetOrderInfoParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetOrderInfoReply>('/app/v1/mall_order/order/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 支付回调接口 返回值: An unexpected error response. POST /app/v1/mall_order/payment/callback */
export function paymentCallback({
  body,
  options,
}: {
  body: API.PaymentCallbackReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.PaymentCallbackReply>(
    '/app/v1/mall_order/payment/callback',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    }
  );
}

/** 获取支付信息接口-选择支付方式并获取支付信息 返回值: An unexpected error response. POST /app/v1/mall_order/payment/info */
export function getPaymentInfo({
  body,
  options,
}: {
  body: API.GetPaymentInfoReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetPaymentInfoReply>('/app/v1/mall_order/payment/info', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 下单接口-创建订单 返回值: An unexpected error response. POST /app/v1/mall_order/place */
export function placeOrder({
  body,
  options,
}: {
  body: API.PlaceOrderReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.PlaceOrderReply>('/app/v1/mall_order/place', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 查询订单接口-用户订单列表 返回值: An unexpected error response. GET /app/v1/mall_order/user/orders */
export function getUserOrderList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserOrderListParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetUserOrderListReply>('/app/v1/mall_order/user/orders', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
