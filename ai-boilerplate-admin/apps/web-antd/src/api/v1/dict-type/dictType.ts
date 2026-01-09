/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 字典类型表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/dict_type/create */
export function createDictType({
  body,
  options,
}: {
  body: API.CreateDictTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateDictTypeReply>('/admin/v1/dict_type/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典类型表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/dict_type/delete */
export function deleteDictType({
  body,
  options,
}: {
  body: API.DeleteDictTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteDictTypeReply>('/admin/v1/dict_type/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典类型表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/dict_type/info */
export function getDictTypeInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDictTypeInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDictTypeInfoReply>('/admin/v1/dict_type/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 字典类型表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/dict_type/list */
export function getDictTypeList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDictTypeListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDictTypeListReply>('/admin/v1/dict_type/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 字典类型表-选择器 返回值: An unexpected error response. GET /admin/v1/dict_type/selector */
export function getDictTypeSelector({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDictTypeSelectorParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDictTypeSelectorReply>('/admin/v1/dict_type/selector', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 字典类型表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/dict_type/update */
export function updateDictType({
  body,
  options,
}: {
  body: API.UpdateDictTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateDictTypeReply>('/admin/v1/dict_type/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 字典类型表-更新状态 返回值: An unexpected error response. POST /admin/v1/dict_type/update/status */
export function updateDictTypeStatus({
  body,
  options,
}: {
  body: API.UpdateDictTypeStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateDictTypeStatusReply>(
    '/admin/v1/dict_type/update/status',
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
