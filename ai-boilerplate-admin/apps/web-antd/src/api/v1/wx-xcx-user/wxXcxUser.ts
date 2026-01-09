/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 小程序用户表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_xcx_user/delete */
export function deleteWxXcxUser({
  body,
  options,
}: {
  body: API.DeleteWxXcxUserReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxXcxUserReply>('/admin/v1/wx_xcx_user/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 小程序用户表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_xcx_user/info */
export function getWxXcxUserInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxXcxUserInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxXcxUserInfoReply>('/admin/v1/wx_xcx_user/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 小程序用户表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_xcx_user/list */
export function getWxXcxUserList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxXcxUserListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxXcxUserListReply>('/admin/v1/wx_xcx_user/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
