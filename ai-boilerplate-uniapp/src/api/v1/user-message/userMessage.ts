/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** App-用户消息-分类未读数量 返回值: An unexpected error response. GET /app/v1/user_message/category_counts */
export function getUserMessageCategoryCounts({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetUserMessageCategoryCountsReply>(
    '/app/v1/user_message/category_counts',
    {
      method: 'GET',
      ...(options || {}),
    }
  );
}

/** App-用户消息-单条数据查询 返回值: An unexpected error response. GET /app/v1/user_message/info */
export function getUserMessageInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMessageInfoParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetUserMessageInfoReply>('/app/v1/user_message/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** App-用户消息-列表数据查询 返回值: An unexpected error response. GET /app/v1/user_message/list */
export function getUserMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMessageListParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetUserMessageListReply>('/app/v1/user_message/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** App-用户消息-标记已读 返回值: An unexpected error response. POST /app/v1/user_message/read */
export function updateUserMessageRead({
  body,
  options,
}: {
  body: API.UpdateUserMessageReadReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.UpdateUserMessageReadReply>('/app/v1/user_message/read', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
