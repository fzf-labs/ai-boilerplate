/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 视频表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_video_record/create */
export function createAiVideoRecord({
  body,
  options,
}: {
  body: API.CreateAiVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiVideoRecordReply>(
    '/admin/v1/ai_video_record/create',
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

/** AI 视频表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_video_record/delete */
export function deleteAiVideoRecord({
  body,
  options,
}: {
  body: API.DeleteAiVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiVideoRecordReply>(
    '/admin/v1/ai_video_record/delete',
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

/** AI 视频表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_video_record/info */
export function getAiVideoRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiVideoRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiVideoRecordInfoReply>(
    '/admin/v1/ai_video_record/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 视频表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_video_record/list */
export function getAiVideoRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiVideoRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiVideoRecordListReply>(
    '/admin/v1/ai_video_record/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 视频表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_video_record/update */
export function updateAiVideoRecord({
  body,
  options,
}: {
  body: API.UpdateAiVideoRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiVideoRecordReply>(
    '/admin/v1/ai_video_record/update',
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

/** AI 视频表-更新状态 返回值: An unexpected error response. POST /admin/v1/ai_video_record/update/status */
export function updateAiVideoRecordStatus({
  body,
  options,
}: {
  body: API.UpdateAiVideoRecordStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiVideoRecordStatusReply>(
    '/admin/v1/ai_video_record/update/status',
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
