/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 写作表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_write_record/create */
export function createAiWriteRecord({
  body,
  options,
}: {
  body: API.CreateAiWriteRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiWriteRecordReply>(
    '/admin/v1/ai_write_record/create',
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

/** AI 写作表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_write_record/delete */
export function deleteAiWriteRecord({
  body,
  options,
}: {
  body: API.DeleteAiWriteRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiWriteRecordReply>(
    '/admin/v1/ai_write_record/delete',
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

/** AI 写作表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_write_record/info */
export function getAiWriteRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiWriteRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiWriteRecordInfoReply>(
    '/admin/v1/ai_write_record/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 写作表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_write_record/list */
export function getAiWriteRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiWriteRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiWriteRecordListReply>(
    '/admin/v1/ai_write_record/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 写作表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_write_record/update */
export function updateAiWriteRecord({
  body,
  options,
}: {
  body: API.UpdateAiWriteRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiWriteRecordReply>(
    '/admin/v1/ai_write_record/update',
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
