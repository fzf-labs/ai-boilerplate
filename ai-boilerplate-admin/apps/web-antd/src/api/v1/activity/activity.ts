/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 内容-活动-创建一条数据 返回值: An unexpected error response. POST /admin/v1/activity/create */
export function createActivity({
  body,
  options,
}: {
  body: API.CreateActivityReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateActivityReply>('/admin/v1/activity/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-活动-删除一条数据 返回值: An unexpected error response. POST /admin/v1/activity/delete */
export function deleteActivity({
  body,
  options,
}: {
  body: API.DeleteActivityReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteActivityReply>('/admin/v1/activity/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-活动-单条数据查询 返回值: An unexpected error response. GET /admin/v1/activity/info */
export function getActivityInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetActivityInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetActivityInfoReply>('/admin/v1/activity/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 内容-活动-列表数据查询 返回值: An unexpected error response. GET /admin/v1/activity/list */
export function getActivityList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetActivityListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetActivityListReply>('/admin/v1/activity/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 内容-活动-更新一条数据 返回值: An unexpected error response. POST /admin/v1/activity/update */
export function updateActivity({
  body,
  options,
}: {
  body: API.UpdateActivityReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateActivityReply>('/admin/v1/activity/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-活动-更新状态 返回值: An unexpected error response. POST /admin/v1/activity/update/status */
export function updateActivityStatus({
  body,
  options,
}: {
  body: API.UpdateActivityStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateActivityStatusReply>(
    '/admin/v1/activity/update/status',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    }
  );
}
