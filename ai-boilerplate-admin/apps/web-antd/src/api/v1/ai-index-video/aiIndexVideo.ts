/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 视频表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_video/create */
export function createAiIndexVideoRecord({
  body,
  options,
}: {
  body: API.CreateAiIndexVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiIndexVideoRecordReply>(
    '/admin/v1/ai_index_video/create',
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

/** AI 视频表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_video/delete */
export function deleteAiIndexVideoRecord({
  body,
  options,
}: {
  body: API.DeleteAiIndexVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiIndexVideoRecordReply>(
    '/admin/v1/ai_index_video/delete',
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

/** AI 视频表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_video/info */
export function getAiIndexVideoRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexVideoRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexVideoRecordInfoReply>(
    '/admin/v1/ai_index_video/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 视频表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_video/list */
export function getAiIndexVideoRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexVideoRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexVideoRecordListReply>(
    '/admin/v1/ai_index_video/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 视频表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_video/update */
export function updateAiIndexVideoRecord({
  body,
  options,
}: {
  body: API.UpdateAiIndexVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiIndexVideoRecordReply>(
    '/admin/v1/ai_index_video/update',
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

/** AI 视频表-更新状态 返回值: An unexpected error response. POST /admin/v1/ai_index_video/update/status */
export function updateAiIndexVideoRecordStatus({
  body,
  options,
}: {
  body: API.UpdateAiIndexVideoRecordStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiIndexVideoRecordStatusReply>(
    '/admin/v1/ai_index_video/update/status',
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
