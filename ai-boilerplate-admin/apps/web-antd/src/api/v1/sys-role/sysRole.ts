/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-角色-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_role/create */
export function createSysRole({
  body,
  options,
}: {
  body: API.CreateSysRoleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysRoleReply>('/admin/v1/sys_role/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-角色-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_role/delete */
export function deleteSysRole({
  body,
  options,
}: {
  body: API.DeleteSysRoleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysRoleReply>('/admin/v1/sys_role/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-角色-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_role/info */
export function getSysRoleInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysRoleInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysRoleInfoReply>('/admin/v1/sys_role/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-角色-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_role/list */
export function getSysRoleList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysRoleListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysRoleListReply>('/admin/v1/sys_role/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-角色-选择器 返回值: An unexpected error response. GET /admin/v1/sys_role/selector */
export function getSysRoleSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysRoleSelectorReply>('/admin/v1/sys_role/selector', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 系统-角色-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_role/update */
export function updateSysRole({
  body,
  options,
}: {
  body: API.UpdateSysRoleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysRoleReply>('/admin/v1/sys_role/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-角色-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_role/update/status */
export function updateSysRoleStatus({
  body,
  options,
}: {
  body: API.UpdateSysRoleStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysRoleStatusReply>(
    '/admin/v1/sys_role/update/status',
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
