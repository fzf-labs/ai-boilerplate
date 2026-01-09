/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 激活码管理表-批量生成激活码 返回值: An unexpected error response. POST /admin/v1/mall_activation_code/batch_generate */
export function batchGenerateMallActivationCode({
  body,
  options,
}: {
  body: API.BatchGenerateMallActivationCodeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.BatchGenerateMallActivationCodeReply>(
    '/admin/v1/mall_activation_code/batch_generate',
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

/** 激活码管理表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/mall_activation_code/delete */
export function deleteMallActivationCode({
  body,
  options,
}: {
  body: API.DeleteMallActivationCodeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMallActivationCodeReply>(
    '/admin/v1/mall_activation_code/delete',
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

/** 激活码管理表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mall_activation_code/info */
export function getMallActivationCodeInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallActivationCodeInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallActivationCodeInfoReply>(
    '/admin/v1/mall_activation_code/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 激活码管理表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mall_activation_code/list */
export function getMallActivationCodeList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMallActivationCodeListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMallActivationCodeListReply>(
    '/admin/v1/mall_activation_code/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 激活码管理表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/mall_activation_code/update */
export function updateMallActivationCode({
  body,
  options,
}: {
  body: API.UpdateMallActivationCodeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMallActivationCodeReply>(
    '/admin/v1/mall_activation_code/update',
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

/** 激活码管理表-更新状态 返回值: An unexpected error response. POST /admin/v1/mall_activation_code/update/status */
export function updateMallActivationCodeStatus({
  body,
  options,
}: {
  body: API.UpdateMallActivationCodeStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMallActivationCodeStatusReply>(
    '/admin/v1/mall_activation_code/update/status',
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
