/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-租户-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_tenant/create */
export function createSysTenant({
  body,
  options,
}: {
  body: API.CreateSysTenantReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysTenantReply>('/admin/v1/sys_tenant/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-租户-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_tenant/delete */
export function deleteSysTenant({
  body,
  options,
}: {
  body: API.DeleteSysTenantReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysTenantReply>('/admin/v1/sys_tenant/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-租户-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_tenant/info */
export function getSysTenantInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysTenantInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysTenantInfoReply>('/admin/v1/sys_tenant/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-租户-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_tenant/list */
export function getSysTenantList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysTenantListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysTenantListReply>('/admin/v1/sys_tenant/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-租户-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_tenant/update */
export function updateSysTenant({
  body,
  options,
}: {
  body: API.UpdateSysTenantReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysTenantReply>('/admin/v1/sys_tenant/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-租户-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_tenant/update/status */
export function updateSysTenantStatus({
  body,
  options,
}: {
  body: API.UpdateSysTenantStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysTenantStatusReply>(
    '/admin/v1/sys_tenant/update/status',
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
