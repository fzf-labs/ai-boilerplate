/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 自应用版本发布表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/self_app_release/create */
export function createSelfAppRelease({
  body,
  options,
}: {
  body: API.CreateSelfAppReleaseReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSelfAppReleaseReply>(
    '/admin/v1/self_app_release/create',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
}

/** 自应用版本发布表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/self_app_release/delete */
export function deleteSelfAppRelease({
  body,
  options,
}: {
  body: API.DeleteSelfAppReleaseReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSelfAppReleaseReply>(
    '/admin/v1/self_app_release/delete',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
}

/** 自应用版本发布表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/self_app_release/info */
export function getSelfAppReleaseInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSelfAppReleaseInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSelfAppReleaseInfoReply>(
    '/admin/v1/self_app_release/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 自应用版本发布表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/self_app_release/list */
export function getSelfAppReleaseList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSelfAppReleaseListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSelfAppReleaseListReply>(
    '/admin/v1/self_app_release/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 自应用版本发布表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/self_app_release/update */
export function updateSelfAppRelease({
  body,
  options,
}: {
  body: API.UpdateSelfAppReleaseReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSelfAppReleaseReply>(
    '/admin/v1/self_app_release/update',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
}

/** 自应用版本发布表-更新状态 返回值: An unexpected error response. POST /admin/v1/self_app_release/update/status */
export function updateSelfAppReleaseStatus({
  body,
  options,
}: {
  body: API.UpdateSelfAppReleaseStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSelfAppReleaseStatusReply>(
    '/admin/v1/self_app_release/update/status',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    },
  );
}
