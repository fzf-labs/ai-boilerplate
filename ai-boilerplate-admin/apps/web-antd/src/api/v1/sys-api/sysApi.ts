/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-接口-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_api/create */
export function createSysApi({
  body,
  options,
}: {
  body: API.CreateSysAPIReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysAPIReply>('/admin/v1/sys_api/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-接口-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_api/delete */
export function deleteSysApi({
  body,
  options,
}: {
  body: API.DeleteSysAPIReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysAPIReply>('/admin/v1/sys_api/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-接口-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_api/info */
export function getSysApiInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysAPIInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysAPIInfoReply>('/admin/v1/sys_api/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-接口-列表数据查询 返回值: An unexpected error response. POST /admin/v1/sys_api/list */
export function getSysApiList({
  body,
  options,
}: {
  body: API.GetSysAPIListReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysAPIListReply>('/admin/v1/sys_api/list', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-接口-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_api/update */
export function updateSysApi({
  body,
  options,
}: {
  body: API.UpdateSysAPIReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysAPIReply>('/admin/v1/sys_api/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
