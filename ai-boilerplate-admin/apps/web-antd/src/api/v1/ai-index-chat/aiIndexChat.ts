/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 聊天对话表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_chat/conversation/create */
export function createAiIndexChatConversation({
  body,
  options,
}: {
  body: API.CreateAiIndexChatConversationReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiIndexChatConversationReply>(
    '/admin/v1/ai_index_chat/conversation/create',
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

/** AI 聊天对话表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_chat/conversation/delete */
export function deleteAiIndexChatConversation({
  body,
  options,
}: {
  body: API.DeleteAiIndexChatConversationReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiIndexChatConversationReply>(
    '/admin/v1/ai_index_chat/conversation/delete',
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

/** AI 聊天对话表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_chat/conversation/info */
export function getAiIndexChatConversationItem({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexChatConversationItemParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexChatConversationItemReply>(
    '/admin/v1/ai_index_chat/conversation/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 聊天对话表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_chat/conversation/list */
export function getAiIndexChatConversationList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexChatConversationListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexChatConversationListReply>(
    '/admin/v1/ai_index_chat/conversation/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 聊天对话表-置顶和取消置顶 返回值: An unexpected error response. POST /admin/v1/ai_index_chat/conversation/pin */
export function pinAiIndexChatConversation({
  body,
  options,
}: {
  body: API.PinAiIndexChatConversationReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.PinAiIndexChatConversationReply>(
    '/admin/v1/ai_index_chat/conversation/pin',
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

/** AI 聊天对话表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_index_chat/conversation/update */
export function updateAiIndexChatConversation({
  body,
  options,
}: {
  body: API.UpdateAiIndexChatConversationReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiIndexChatConversationReply>(
    '/admin/v1/ai_index_chat/conversation/update',
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

/** AI 聊天消息表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_index_chat/message/list */
export function getAiIndexChatMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiIndexChatMessageListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiIndexChatMessageListReply>(
    '/admin/v1/ai_index_chat/message/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
