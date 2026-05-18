/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 帮助中心常见问题表-列表数据查询 返回值: An unexpected error response. GET /app/v1/help/faqs */
export function listHelpFaqs({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.ListHelpFaqsParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.ListHelpFaqsReply>('/app/v1/help/faqs', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
