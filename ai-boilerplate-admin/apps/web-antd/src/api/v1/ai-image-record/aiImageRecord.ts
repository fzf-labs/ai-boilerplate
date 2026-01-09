/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 绘画表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_image_record/create */
export function createAiImageRecord({
  body,
  options,
}: {
  body: API.CreateAiImageRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiImageRecordReply>(
    '/admin/v1/ai_image_record/create',
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

/** AI 绘画表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_image_record/delete */
export function deleteAiImageRecord({
  body,
  options,
}: {
  body: API.DeleteAiImageRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiImageRecordReply>(
    '/admin/v1/ai_image_record/delete',
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

/** AI 绘画表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_image_record/info */
export function getAiImageRecordInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiImageRecordInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiImageRecordInfoReply>(
    '/admin/v1/ai_image_record/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 绘画表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_image_record/list */
export function getAiImageRecordList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiImageRecordListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiImageRecordListReply>(
    '/admin/v1/ai_image_record/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 绘画表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_image_record/update */
export function updateAiImageRecord({
  body,
  options,
}: {
  body: API.UpdateAiImageRecordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiImageRecordReply>(
    '/admin/v1/ai_image_record/update',
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

/** AI 绘画表-更新状态 返回值: An unexpected error response. POST /admin/v1/ai_image_record/update/status */
export function updateAiImageRecordStatus({
  body,
  options,
}: {
  body: API.UpdateAiImageRecordStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiImageRecordStatusReply>(
    '/admin/v1/ai_image_record/update/status',
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
