/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 检测App版本更新 返回值: An unexpected error response. POST /app/v1/app_release/check_update */
export function checkUpdate({
  body,
  options,
}: {
  body: API.CheckUpdateReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.CheckUpdateReply>('/app/v1/app_release/check_update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
