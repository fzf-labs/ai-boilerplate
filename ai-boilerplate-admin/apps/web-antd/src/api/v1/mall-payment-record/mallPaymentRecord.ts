/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 支付记录表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mall_payment_record/success/list/order_id */
export function getMallPaymentRecordListByOrderId({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallPaymentRecordListByOrderIdParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallPaymentRecordListByOrderIdReply>(
    '/admin/v1/mall_payment_record/success/list/order_id',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 支付记录表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mall_payment_record/success/order_id */
export function getMallPaymentRecordSuccessByOrderId({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallPaymentRecordSuccessByOrderIdParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallPaymentRecordSuccessByOrderIdReply>(
    '/admin/v1/mall_payment_record/success/order_id',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
