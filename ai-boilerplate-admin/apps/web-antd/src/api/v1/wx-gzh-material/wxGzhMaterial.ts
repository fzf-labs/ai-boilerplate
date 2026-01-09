/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号素材表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_material/delete */
export function deleteWxGzhMaterial({
  body,
  options,
}: {
  body: API.DeleteWxGzhMaterialReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhMaterialReply>(
    '/admin/v1/wx_gzh_material/delete',
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

/** 公众号素材表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_material/info */
export function getWxGzhMaterialInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMaterialInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMaterialInfoReply>(
    '/admin/v1/wx_gzh_material/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号素材表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_material/list */
export function getWxGzhMaterialList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMaterialListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMaterialListReply>(
    '/admin/v1/wx_gzh_material/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号素材表-统计 返回值: An unexpected error response. GET /admin/v1/wx_gzh_material/stats */
export function getWxGzhMaterialStats({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMaterialStatsParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMaterialStatsReply>(
    '/admin/v1/wx_gzh_material/stats',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号素材表-素材同步 返回值: An unexpected error response. POST /admin/v1/wx_gzh_material/sync */
export function syncWxGzhMaterial({
  body,
  options,
}: {
  body: API.SyncWxGzhMaterialReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SyncWxGzhMaterialReply>('/admin/v1/wx_gzh_material/sync', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
