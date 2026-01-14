/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 获取商品详情 返回值: An unexpected error response. GET /app/v1/mall_product/info */
export function getMallProductInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallProductInfoParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetMallProductInfoReply>('/app/v1/mall_product/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取商品列表 返回值: An unexpected error response. GET /app/v1/mall_product/list */
export function getMallProductList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallProductListParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetMallProductListReply>('/app/v1/mall_product/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
