/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-通知消息-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sys_notify_message/info */
export function getSysNotifyMessageInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysNotifyMessageInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNotifyMessageInfoReply>(
    '/admin/v1/sys_notify_message/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 系统-通知消息-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_notify_message/list */
export function getSysNotifyMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysNotifyMessageListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNotifyMessageListReply>(
    '/admin/v1/sys_notify_message/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 系统-通知消息-我的-列表 返回值: An unexpected error response. GET /admin/v1/sys_notify_message/my/list */
export function getSysNotifyMessageMyList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysNotifyMessageMyListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNotifyMessageMyListReply>(
    '/admin/v1/sys_notify_message/my/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 系统-通知消息-我的-未读数量 返回值: An unexpected error response. GET /admin/v1/sys_notify_message/my/unread_count */
export function getSysNotifyMessageMyUnreadCount({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNotifyMessageMyUnreadCountReply>(
    '/admin/v1/sys_notify_message/my/unread_count',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 系统-通知消息-我的-未读列表 返回值: An unexpected error response. GET /admin/v1/sys_notify_message/my/unread_list */
export function getSysNotifyMessageMyUnreadList({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysNotifyMessageMyUnreadListReply>(
    '/admin/v1/sys_notify_message/my/unread_list',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 系统-通知消息-我的-全部已读 返回值: An unexpected error response. POST /admin/v1/sys_notify_message/my/update_all_read */
export function updateSysNotifyMessageAllRead({
  body,
  options,
}: {
  body: API.UpdateSysNotifyMessageAllReadReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysNotifyMessageAllReadReply>(
    '/admin/v1/sys_notify_message/my/update_all_read',
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

/** 系统-通知消息-我的-指定消息已读 返回值: An unexpected error response. POST /admin/v1/sys_notify_message/my/update_read */
export function updateSysNotifyMessageRead({
  body,
  options,
}: {
  body: API.UpdateSysNotifyMessageReadReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSysNotifyMessageReadReply>(
    '/admin/v1/sys_notify_message/my/update_read',
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
