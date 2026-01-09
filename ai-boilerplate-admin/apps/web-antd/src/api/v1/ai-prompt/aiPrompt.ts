/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 提示词-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_prompt/info */
export function getAiPromptInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiPromptInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiPromptInfoReply>('/admin/v1/ai_prompt/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** AI 提示词-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_prompt/list */
export function getAiPromptList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiPromptListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiPromptListReply>('/admin/v1/ai_prompt/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
