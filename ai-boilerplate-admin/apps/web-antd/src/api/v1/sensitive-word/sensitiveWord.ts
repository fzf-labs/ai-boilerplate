/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 敏感词-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sensitive_word/create */
export function createSensitiveWord({
  body,
  options,
}: {
  body: API.CreateSensitiveWordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSensitiveWordReply>(
    '/admin/v1/sensitive_word/create',
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

/** 敏感词-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sensitive_word/delete */
export function deleteSensitiveWord({
  body,
  options,
}: {
  body: API.DeleteSensitiveWordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSensitiveWordReply>(
    '/admin/v1/sensitive_word/delete',
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

/** 敏感词-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sensitive_word/info */
export function getSensitiveWordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSensitiveWordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSensitiveWordInfoReply>(
    '/admin/v1/sensitive_word/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 敏感词-标签选择器 返回值: An unexpected error response. GET /admin/v1/sensitive_word/labs/selector */
export function getSensitiveWordLabsSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSensitiveWordLabsSelectorReply>(
    '/admin/v1/sensitive_word/labs/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 敏感词-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sensitive_word/list */
export function getSensitiveWordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSensitiveWordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSensitiveWordListReply>(
    '/admin/v1/sensitive_word/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 敏感词-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sensitive_word/update */
export function updateSensitiveWord({
  body,
  options,
}: {
  body: API.UpdateSensitiveWordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSensitiveWordReply>(
    '/admin/v1/sensitive_word/update',
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
