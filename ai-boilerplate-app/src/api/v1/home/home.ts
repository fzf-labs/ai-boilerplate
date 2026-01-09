/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 获取轮播图列表 返回值: An unexpected error response. GET /app/v1/home/banner/list */
export function getBannerList({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetBannerListReply>('/app/v1/home/banner/list', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取内容详情 返回值: An unexpected error response. GET /app/v1/home/content/detail */
export function getContentDetail({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetContentDetailParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetContentDetailReply>('/app/v1/home/content/detail', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取内容列表 返回值: An unexpected error response. GET /app/v1/home/content/list */
export function getContentList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetContentListParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetContentListReply>('/app/v1/home/content/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
