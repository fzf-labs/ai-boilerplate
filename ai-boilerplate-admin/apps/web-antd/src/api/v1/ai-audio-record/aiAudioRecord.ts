/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 音乐表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_audio_record/delete */
export function deleteAiAudioRecord({
  body,
  options,
}: {
  body: API.DeleteAiAudioRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiAudioRecordReply>(
    '/admin/v1/ai_audio_record/delete',
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

/** AI 音乐表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_audio_record/info */
export function getAiAudioRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiAudioRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiAudioRecordInfoReply>(
    '/admin/v1/ai_audio_record/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 音乐表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_audio_record/list */
export function getAiAudioRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiAudioRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiAudioRecordListReply>(
    '/admin/v1/ai_audio_record/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
