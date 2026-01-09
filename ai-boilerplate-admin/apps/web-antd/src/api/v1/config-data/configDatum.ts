/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 配置管理-创建一条数据 返回值: An unexpected error response. POST /admin/v1/config_data/create */
export function createConfigDatum({
  body,
  options,
}: {
  body: API.CreateConfigDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateConfigDatumReply>('/admin/v1/config_data/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 配置管理-删除一条数据 返回值: An unexpected error response. POST /admin/v1/config_data/delete */
export function deleteConfigDatum({
  body,
  options,
}: {
  body: API.DeleteConfigDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteConfigDatumReply>('/admin/v1/config_data/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 配置管理-单条数据查询 返回值: An unexpected error response. GET /admin/v1/config_data/info */
export function getConfigDatumInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetConfigDatumInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetConfigDatumInfoReply>('/admin/v1/config_data/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 配置管理-列表数据查询 返回值: An unexpected error response. GET /admin/v1/config_data/list */
export function getConfigDatumList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetConfigDatumListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetConfigDatumListReply>('/admin/v1/config_data/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 配置管理-更新一条数据 返回值: An unexpected error response. POST /admin/v1/config_data/update */
export function updateConfigDatum({
  body,
  options,
}: {
  body: API.UpdateConfigDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateConfigDatumReply>('/admin/v1/config_data/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 配置管理-更新状态 返回值: An unexpected error response. POST /admin/v1/config_data/update/status */
export function updateConfigDatumStatus({
  body,
  options,
}: {
  body: API.UpdateConfigDatumStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateConfigDatumStatusReply>(
    '/admin/v1/config_data/update/status',
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
