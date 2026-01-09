/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 用户反馈表-提交反馈 返回值: An unexpected error response. POST /app/v1/help/feedback */
export function submitFeedback({
  body,
  options,
}: {
  body: API.SubmitFeedbackReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.SubmitFeedbackReply>('/app/v1/help/feedback', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
