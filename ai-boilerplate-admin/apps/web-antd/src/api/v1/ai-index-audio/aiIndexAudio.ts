/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 音乐表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_audio/create */
export function createAiIndexAudioRecord({
  body,
  options,
}: {
  body: API.CreateAiIndexAudioRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiIndexAudioRecordReply>(
    '/admin/v1/ai_index_audio/create',
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

/** AI 音乐表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_audio/delete */
export function deleteAiIndexAudioRecord({
  body,
  options,
}: {
  body: API.DeleteAiIndexAudioRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiIndexAudioRecordReply>(
    '/admin/v1/ai_index_audio/delete',
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

/** AI 音乐表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_audio/info */
export function getAiIndexAudioRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexAudioRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexAudioRecordInfoReply>(
    '/admin/v1/ai_index_audio/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 音乐表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_audio/list */
export function getAiIndexAudioRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexAudioRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexAudioRecordListReply>(
    '/admin/v1/ai_index_audio/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
