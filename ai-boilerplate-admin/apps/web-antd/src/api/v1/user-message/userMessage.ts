/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** App-用户消息-单条数据查询 返回值: An unexpected error response. GET /admin/v1/user_message/info */
export function getUserMessageInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMessageInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMessageInfoReply>('/admin/v1/user_message/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** App-用户消息-列表数据查询 返回值: An unexpected error response. GET /admin/v1/user_message/list */
export function getUserMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMessageListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMessageListReply>('/admin/v1/user_message/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** App-用户消息-发送 返回值: An unexpected error response. POST /admin/v1/user_message/send */
export function sendUserMessage({
  body,
  options,
}: {
  body: API.SendUserMessageReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SendUserMessageReply>('/admin/v1/user_message/send', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
