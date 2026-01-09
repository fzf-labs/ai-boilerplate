/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 绘画表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_image/create */
export function createAiIndexImageRecord({
  body,
  options,
}: {
  body: API.CreateAiIndexImageRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiIndexImageRecordReply>(
    '/admin/v1/ai_index_image/create',
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

/** AI 绘画表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_image/delete */
export function deleteAiIndexImageRecord({
  body,
  options,
}: {
  body: API.DeleteAiIndexImageRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiIndexImageRecordReply>(
    '/admin/v1/ai_index_image/delete',
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

/** AI 绘画表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_image/list */
export function getAiIndexImageRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexImageRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexImageRecordListReply>(
    '/admin/v1/ai_index_image/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
