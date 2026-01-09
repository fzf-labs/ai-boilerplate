/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号标签表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_tag/create */
export function createWxGzhTag({
  body,
  options,
}: {
  body: API.CreateWxGzhTagReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateWxGzhTagReply>('/admin/v1/wx_gzh_tag/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 公众号标签表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_tag/delete */
export function deleteWxGzhTag({
  body,
  options,
}: {
  body: API.DeleteWxGzhTagReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhTagReply>('/admin/v1/wx_gzh_tag/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 公众号标签表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_tag/info */
export function getWxGzhTagInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhTagInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhTagInfoReply>('/admin/v1/wx_gzh_tag/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 公众号标签表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_tag/list */
export function getWxGzhTagList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhTagListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhTagListReply>('/admin/v1/wx_gzh_tag/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 公众号标签表-选择器 返回值: An unexpected error response. GET /admin/v1/wx_gzh_tag/selector */
export function getWxGzhTagSelector({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhTagSelectorParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhTagSelectorReply>(
    '/admin/v1/wx_gzh_tag/selector',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号标签表-同步标签 返回值: An unexpected error response. POST /admin/v1/wx_gzh_tag/sync */
export function syncWxGzhTag({
  body,
  options,
}: {
  body: API.SyncWxGzhTagReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SyncWxGzhTagReply>('/admin/v1/wx_gzh_tag/sync', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 公众号标签表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_tag/update */
export function updateWxGzhTag({
  body,
  options,
}: {
  body: API.UpdateWxGzhTagReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateWxGzhTagReply>('/admin/v1/wx_gzh_tag/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
