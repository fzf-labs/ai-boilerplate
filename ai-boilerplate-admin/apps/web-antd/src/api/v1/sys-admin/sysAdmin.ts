/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-用户-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_admin/create */
export function createSysAdmin({
  body,
  options,
}: {
  body: API.CreateSysAdminReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysAdminReply>('/admin/v1/sys_admin/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-用户-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_admin/delete */
export function deleteSysAdmin({
  body,
  options,
}: {
  body: API.DeleteSysAdminReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysAdminReply>('/admin/v1/sys_admin/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-用户-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_admin/info */
export function getSysAdminInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysAdminInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysAdminInfoReply>('/admin/v1/sys_admin/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-用户-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_admin/list */
export function getSysAdminList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysAdminListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysAdminListReply>('/admin/v1/sys_admin/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-用户-选择器 返回值: An unexpected error response. GET /admin/v1/sys_admin/selector */
export function getSysAdminSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysAdminSelectorReply>('/admin/v1/sys_admin/selector', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 系统-用户-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_admin/update */
export function updateSysAdmin({
  body,
  options,
}: {
  body: API.UpdateSysAdminReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysAdminReply>('/admin/v1/sys_admin/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-用户-重置密码 返回值: An unexpected error response. POST /admin/v1/sys_admin/update/password */
export function updateSysAdminPassword({
  body,
  options,
}: {
  body: API.UpdateSysAdminPasswordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysAdminPasswordReply>(
    '/admin/v1/sys_admin/update/password',
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

/** 系统-用户-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_admin/update/status */
export function updateSysAdminStatus({
  body,
  options,
}: {
  body: API.UpdateSysAdminStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysAdminStatusReply>(
    '/admin/v1/sys_admin/update/status',
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
