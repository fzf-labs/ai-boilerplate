/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 聊天对话表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_chat_conversation/delete */
export function deleteAiChatConversation({
  body,
  options,
}: {
  body: API.DeleteAiChatConversationReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiChatConversationReply>(
    '/admin/v1/ai_chat_conversation/delete',
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

/** AI 聊天对话表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_chat_conversation/info */
export function getAiChatConversationInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiChatConversationInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiChatConversationInfoReply>(
    '/admin/v1/ai_chat_conversation/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 聊天对话表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_chat_conversation/list */
export function getAiChatConversationList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiChatConversationListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiChatConversationListReply>(
    '/admin/v1/ai_chat_conversation/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
