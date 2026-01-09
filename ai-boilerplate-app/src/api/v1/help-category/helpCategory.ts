/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 帮助分类表-列表数据查询 返回值: An unexpected error response. GET /app/v1/help/categories */
export function listHelpCategories({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.ListHelpCategoriesReply>('/app/v1/help/categories', {
    method: 'GET',
    ...(options || {}),
  });
}
