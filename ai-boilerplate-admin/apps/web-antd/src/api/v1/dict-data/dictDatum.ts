/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 字典数据表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/dict_data/create */
export function createDictDatum({
  body,
  options,
}: {
  body: API.CreateDictDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateDictDatumReply>('/admin/v1/dict_data/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典数据表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/dict_data/delete */
export function deleteDictDatum({
  body,
  options,
}: {
  body: API.DeleteDictDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteDictDatumReply>('/admin/v1/dict_data/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典数据表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/dict_data/info */
export function getDictDatumInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDictDatumInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDictDatumInfoReply>('/admin/v1/dict_data/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 字典数据表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/dict_data/list */
export function getDictDatumList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDictDatumListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDictDatumListReply>('/admin/v1/dict_data/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 字典数据表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/dict_data/update */
export function updateDictDatum({
  body,
  options,
}: {
  body: API.UpdateDictDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateDictDatumReply>('/admin/v1/dict_data/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典数据表-更新状态 返回值: An unexpected error response. POST /admin/v1/dict_data/update/status */
export function updateDictDatumStatus({
  body,
  options,
}: {
  body: API.UpdateDictDatumStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateDictDatumStatusReply>(
    '/admin/v1/dict_data/update/status',
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
