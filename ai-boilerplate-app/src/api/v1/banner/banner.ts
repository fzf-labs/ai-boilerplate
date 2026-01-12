/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 通用-轮播图-列表数据查询 返回值: An unexpected error response. GET /app/v1/banner/list */
export function listBanners({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.ListBannersParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.ListBannersReply>('/app/v1/banner/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
