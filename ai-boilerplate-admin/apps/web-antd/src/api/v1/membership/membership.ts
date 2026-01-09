/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 会员类型配置表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/membership/create */
export function createMembership({
  body,
  options,
}: {
  body: API.CreateMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMembershipReply>('/admin/v1/membership/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 会员类型配置表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/membership/delete */
export function deleteMembership({
  body,
  options,
}: {
  body: API.DeleteMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMembershipReply>('/admin/v1/membership/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 会员类型配置表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/membership/info */
export function getMembershipInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipInfoReply>('/admin/v1/membership/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 会员类型配置表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/membership/list */
export function getMembershipList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipListReply>('/admin/v1/membership/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 会员类型配置表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/membership/update */
export function updateMembership({
  body,
  options,
}: {
  body: API.UpdateMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipReply>('/admin/v1/membership/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 会员类型配置表-更新状态 返回值: An unexpected error response. POST /admin/v1/membership/update/status */
export function updateMembershipStatus({
  body,
  options,
}: {
  body: API.UpdateMembershipStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipStatusReply>(
    '/admin/v1/membership/update/status',
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
