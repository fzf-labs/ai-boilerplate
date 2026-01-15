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

/** 获取会员商品列表 返回值: An unexpected error response. GET /app/v1/mall_product/membership/list */
export function getMembershipProductList({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetMembershipProductListReply>(
    '/app/v1/mall_product/membership/list',
    {
      method: 'GET',
      ...(options || {}),
    }
  );
}
