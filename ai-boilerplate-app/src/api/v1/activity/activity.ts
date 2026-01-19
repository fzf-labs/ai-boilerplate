/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 内容-活动-列表数据查询 返回值: An unexpected error response. GET /app/v1/activity/list */
export function listActivities({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.ListActivitiesParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.ListActivitiesReply>('/app/v1/activity/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
