/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 订单信息表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mall_order/info */
export function getMallOrderInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallOrderInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallOrderInfoReply>('/admin/v1/mall_order/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 订单信息表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mall_order/list */
export function getMallOrderList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallOrderListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallOrderListReply>('/admin/v1/mall_order/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
