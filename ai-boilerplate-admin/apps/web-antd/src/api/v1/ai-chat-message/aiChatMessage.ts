/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 聊天消息表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_chat_message/list */
export function getAiChatMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiChatMessageListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiChatMessageListReply>(
    '/admin/v1/ai_chat_message/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
