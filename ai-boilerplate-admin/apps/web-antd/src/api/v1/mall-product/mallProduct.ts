/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 商品表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/mall_product/create */
export function createMallProduct({
  body,
  options,
}: {
  body: API.CreateMallProductReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMallProductReply>('/admin/v1/mall_product/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 商品表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/mall_product/delete */
export function deleteMallProduct({
  body,
  options,
}: {
  body: API.DeleteMallProductReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMallProductReply>('/admin/v1/mall_product/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 商品表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mall_product/info */
export function getMallProductInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallProductInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallProductInfoReply>('/admin/v1/mall_product/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 商品表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mall_product/list */
export function getMallProductList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallProductListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallProductListReply>('/admin/v1/mall_product/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 商品表-选择器 返回值: An unexpected error response. GET /admin/v1/mall_product/selector */
export function getMallProductSelector({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallProductSelectorParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallProductSelectorReply>(
    '/admin/v1/mall_product/selector',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 商品表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/mall_product/update */
export function updateMallProduct({
  body,
  options,
}: {
  body: API.UpdateMallProductReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMallProductReply>('/admin/v1/mall_product/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 商品表-更新状态 返回值: An unexpected error response. POST /admin/v1/mall_product/update/status */
export function updateMallProductStatus({
  body,
  options,
}: {
  body: API.UpdateMallProductStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMallProductStatusReply>(
    '/admin/v1/mall_product/update/status',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
}
