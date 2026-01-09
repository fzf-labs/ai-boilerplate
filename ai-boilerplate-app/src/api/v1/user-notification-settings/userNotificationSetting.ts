/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 获取用户通知设置 返回值: An unexpected error response. GET /app/v1/notification/settings */
export function getNotificationSettings({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetNotificationSettingsReply>(
    '/app/v1/notification/settings',
    {
      method: 'GET',
      ...(options || {}),
    }
  );
}

/** 更新用户通知设置 返回值: An unexpected error response. PUT /app/v1/notification/settings */
export function updateNotificationSettings({
  body,
  options,
}: {
  body: API.UpdateNotificationSettingsReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.UpdateNotificationSettingsReply>(
    '/app/v1/notification/settings',
    {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    }
  );
}
