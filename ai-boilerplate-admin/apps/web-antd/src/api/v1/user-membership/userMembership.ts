/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 用户会员关系表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/user_membership/create */
export function createUserMembership({
  body,
  options,
}: {
  body: API.CreateUserMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateUserMembershipReply>(
    '/admin/v1/user_membership/create',
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

/** 用户会员关系表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/user_membership/delete */
export function deleteUserMembership({
  body,
  options,
}: {
  body: API.DeleteUserMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteUserMembershipReply>(
    '/admin/v1/user_membership/delete',
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

/** 用户会员关系表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/user_membership/info */
export function getUserMembershipInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMembershipInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMembershipInfoReply>(
    '/admin/v1/user_membership/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 用户会员关系表-根据用户ID查询单条数据 返回值: An unexpected error response. GET /admin/v1/user_membership/info/by_user_id */
export function getUserMembershipInfoByUserId({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMembershipInfoByUserIdParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMembershipInfoByUserIdReply>(
    '/admin/v1/user_membership/info/by_user_id',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 用户会员关系表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/user_membership/list */
export function getUserMembershipList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMembershipListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMembershipListReply>(
    '/admin/v1/user_membership/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 用户会员关系表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/user_membership/update */
export function updateUserMembership({
  body,
  options,
}: {
  body: API.UpdateUserMembershipReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateUserMembershipReply>(
    '/admin/v1/user_membership/update',
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

/** 用户会员关系表-更新状态 返回值: An unexpected error response. POST /admin/v1/user_membership/update/status */
export function updateUserMembershipStatus({
  body,
  options,
}: {
  body: API.UpdateUserMembershipStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateUserMembershipStatusReply>(
    '/admin/v1/user_membership/update/status',
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
