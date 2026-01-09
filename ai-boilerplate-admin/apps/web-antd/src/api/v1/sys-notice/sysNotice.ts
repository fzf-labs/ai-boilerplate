/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-公告-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_notice/create */
export function createSysNotice({
  body,
  options,
}: {
  body: API.CreateSysNoticeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysNoticeReply>('/admin/v1/sys_notice/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-公告-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_notice/delete */
export function deleteSysNotice({
  body,
  options,
}: {
  body: API.DeleteSysNoticeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysNoticeReply>('/admin/v1/sys_notice/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-公告-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_notice/info */
export function getSysNoticeInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysNoticeInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNoticeInfoReply>('/admin/v1/sys_notice/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-公告-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_notice/list */
export function getSysNoticeList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysNoticeListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNoticeListReply>('/admin/v1/sys_notice/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 系统-公告-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_notice/update */
export function updateSysNotice({
  body,
  options,
}: {
  body: API.UpdateSysNoticeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysNoticeReply>('/admin/v1/sys_notice/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 系统-公告-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_notice/update/status */
export function updateSysNoticeStatus({
  body,
  options,
}: {
  body: API.UpdateSysNoticeStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysNoticeStatusReply>(
    '/admin/v1/sys_notice/update/status',
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
