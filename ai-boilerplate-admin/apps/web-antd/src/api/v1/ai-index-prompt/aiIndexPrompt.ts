/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 提示词-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_prompt/create */
export function createAiIndexPrompt({
  body,
  options,
}: {
  body: API.CreateAiIndexPromptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiIndexPromptReply>('/admin/v1/ai_prompt/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** AI 提示词-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_prompt/delete */
export function deleteAiIndexPrompt({
  body,
  options,
}: {
  body: API.DeleteAiIndexPromptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiIndexPromptReply>('/admin/v1/ai_prompt/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** AI 提示词-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_prompt/info */
export function getAiIndexPromptInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexPromptInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexPromptInfoReply>('/admin/v1/ai_prompt/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** AI 提示词-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_prompt/list */
export function getAiIndexPromptList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexPromptListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexPromptListReply>('/admin/v1/ai_prompt/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** AI 提示词-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_prompt/update */
export function updateAiIndexPrompt({
  body,
  options,
}: {
  body: API.UpdateAiIndexPromptReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiIndexPromptReply>('/admin/v1/ai_prompt/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
