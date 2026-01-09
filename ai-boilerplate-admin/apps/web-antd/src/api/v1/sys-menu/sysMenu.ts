/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 菜单-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sys_menu/create */
export function createSysMenu({
  body,
  options,
}: {
  body: API.CreateSysMenuReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSysMenuReply>('/admin/v1/sys_menu/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 菜单-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sys_menu/delete */
export function deleteSysMenu({
  body,
  options,
}: {
  body: API.DeleteSysMenuReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSysMenuReply>('/admin/v1/sys_menu/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 菜单-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_menu/info */
export function getSysMenuInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysMenuInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysMenuInfoReply>('/admin/v1/sys_menu/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 菜单-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_menu/list */
export function getSysMenuList({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysMenuListReply>('/admin/v1/sys_menu/list', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 菜单-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sys_menu/update */
export function updateSysMenu({
  body,
  options,
}: {
  body: API.UpdateSysMenuReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysMenuReply>('/admin/v1/sys_menu/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 菜单-更新状态 返回值: An unexpected error response. POST /admin/v1/sys_menu/update/status */
export function updateSysMenuStatus({
  body,
  options,
}: {
  body: API.UpdateSysMenuStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysMenuStatusReply>(
    '/admin/v1/sys_menu/update/status',
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
