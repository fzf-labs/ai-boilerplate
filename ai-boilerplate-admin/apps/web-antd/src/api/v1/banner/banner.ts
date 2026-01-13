/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 通用-轮播图-创建一条数据 返回值: An unexpected error response. POST /admin/v1/banner/create */
export function createBanner({
  body,
  options,
}: {
  body: API.CreateBannerReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateBannerReply>('/admin/v1/banner/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 通用-轮播图-删除一条数据 返回值: An unexpected error response. POST /admin/v1/banner/delete */
export function deleteBanner({
  body,
  options,
}: {
  body: API.DeleteBannerReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteBannerReply>('/admin/v1/banner/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 通用-轮播图-单条数据查询 返回值: An unexpected error response. GET /admin/v1/banner/info */
export function getBannerInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetBannerInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetBannerInfoReply>('/admin/v1/banner/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 通用-轮播图-列表数据查询 返回值: An unexpected error response. GET /admin/v1/banner/list */
export function getBannerList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetBannerListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetBannerListReply>('/admin/v1/banner/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 通用-轮播图-更新一条数据 返回值: An unexpected error response. POST /admin/v1/banner/update */
export function updateBanner({
  body,
  options,
}: {
  body: API.UpdateBannerReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateBannerReply>('/admin/v1/banner/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 通用-轮播图-更新状态 返回值: An unexpected error response. POST /admin/v1/banner/update/status */
export function updateBannerStatus({
  body,
  options,
}: {
  body: API.UpdateBannerStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateBannerStatusReply>(
    '/admin/v1/banner/update/status',
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
