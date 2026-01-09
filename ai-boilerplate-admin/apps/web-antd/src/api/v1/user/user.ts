/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 用户表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/user/create */
export function createUser({
  body,
  options,
}: {
  body: API.CreateUserReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateUserReply>('/admin/v1/user/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 用户表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/user/delete */
export function deleteUser({
  body,
  options,
}: {
  body: API.DeleteUserReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteUserReply>('/admin/v1/user/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 用户表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/user/info */
export function getUserInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserInfoReply>('/admin/v1/user/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 用户表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/user/list */
export function getUserList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserListReply>('/admin/v1/user/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 用户表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/user/update */
export function updateUser({
  body,
  options,
}: {
  body: API.UpdateUserReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateUserReply>('/admin/v1/user/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 用户表-更新状态 返回值: An unexpected error response. POST /admin/v1/user/update/status */
export function updateUserStatus({
  body,
  options,
}: {
  body: API.UpdateUserStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateUserStatusReply>('/admin/v1/user/update/status', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
