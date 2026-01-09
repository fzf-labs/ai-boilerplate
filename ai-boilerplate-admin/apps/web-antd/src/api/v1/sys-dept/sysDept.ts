/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-部门-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_dept/create */
export function createSysDept({
  body,
  options,
}: {
  body: API.CreateSysDeptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysDeptReply>('/admin/v1/sys_dept/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-部门-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_dept/delete */
export function deleteSysDept({
  body,
  options,
}: {
  body: API.DeleteSysDeptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysDeptReply>('/admin/v1/sys_dept/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-部门-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_dept/info */
export function getSysDeptInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysDeptInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysDeptInfoReply>('/admin/v1/sys_dept/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-部门-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_dept/list */
export function getSysDeptList({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysDeptListReply>('/admin/v1/sys_dept/list', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 系统-部门-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_dept/update */
export function updateSysDept({
  body,
  options,
}: {
  body: API.UpdateSysDeptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysDeptReply>('/admin/v1/sys_dept/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-部门-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_dept/update/status */
export function updateSysDeptStatus({
  body,
  options,
}: {
  body: API.UpdateSysDeptStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysDeptStatusReply>(
    '/admin/v1/sys_dept/update/status',
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
