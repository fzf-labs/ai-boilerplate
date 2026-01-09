/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 自应用信息表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/self_app/create */
export function createSelfApp({
  body,
  options,
}: {
  body: API.CreateSelfAppReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSelfAppReply>('/admin/v1/self_app/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 自应用信息表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/self_app/delete */
export function deleteSelfApp({
  body,
  options,
}: {
  body: API.DeleteSelfAppReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSelfAppReply>('/admin/v1/self_app/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 自应用信息表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/self_app/info */
export function getSelfAppInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSelfAppInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSelfAppInfoReply>('/admin/v1/self_app/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 自应用信息表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/self_app/list */
export function getSelfAppList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSelfAppListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSelfAppListReply>('/admin/v1/self_app/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 自应用信息表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/self_app/update */
export function updateSelfApp({
  body,
  options,
}: {
  body: API.UpdateSelfAppReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSelfAppReply>('/admin/v1/self_app/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 自应用信息表-更新状态 返回值: An unexpected error response. POST /admin/v1/self_app/update/status */
export function updateSelfAppStatus({
  body,
  options,
}: {
  body: API.UpdateSelfAppStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSelfAppStatusReply>(
    '/admin/v1/self_app/update/status',
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
