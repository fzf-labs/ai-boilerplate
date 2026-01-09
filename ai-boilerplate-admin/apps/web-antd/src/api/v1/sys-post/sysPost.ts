/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-工作岗位-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_post/create */
export function createSysPost({
  body,
  options,
}: {
  body: API.CreateSysPostReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysPostReply>('/admin/v1/sys_post/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-工作岗位-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_post/delete */
export function deleteSysPost({
  body,
  options,
}: {
  body: API.DeleteSysPostReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysPostReply>('/admin/v1/sys_post/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-工作岗位-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_post/info */
export function getSysPostInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysPostInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysPostInfoReply>('/admin/v1/sys_post/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-工作岗位-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_post/list */
export function getSysPostList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysPostListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysPostListReply>('/admin/v1/sys_post/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-工作岗位-选择器 返回值: An unexpected error response. GET /admin/v1/sys_post/selector */
export function getSysPostSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysPostSelectorReply>('/admin/v1/sys_post/selector', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 系统-工作岗位-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_post/update */
export function updateSysPost({
  body,
  options,
}: {
  body: API.UpdateSysPostReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysPostReply>('/admin/v1/sys_post/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-工作岗位-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_post/update/status */
export function updateSysPostStatus({
  body,
  options,
}: {
  body: API.UpdateSysPostStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysPostStatusReply>(
    '/admin/v1/sys_post/update/status',
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
