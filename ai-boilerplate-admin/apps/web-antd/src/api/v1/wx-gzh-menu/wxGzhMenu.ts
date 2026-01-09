/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号菜单表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_menu/delete */
export function deleteWxGzhMenu({
  body,
  options,
}: {
  body: API.DeleteWxGzhMenuReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhMenuReply>('/admin/v1/wx_gzh_menu/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 公众号菜单表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_menu/info */
export function getWxGzhMenuInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMenuInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMenuInfoReply>('/admin/v1/wx_gzh_menu/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 公众号菜单表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_menu/list */
export function getWxGzhMenuList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMenuListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMenuListReply>('/admin/v1/wx_gzh_menu/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 公众号菜单表-创建/更新一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_menu/store */
export function storeWxGzhMenu({
  body,
  options,
}: {
  body: API.StoreWxGzhMenuReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.StoreWxGzhMenuReply>('/admin/v1/wx_gzh_menu/store', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
