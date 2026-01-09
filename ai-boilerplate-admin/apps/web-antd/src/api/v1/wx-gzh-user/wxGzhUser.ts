/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号粉丝表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_user/delete */
export function deleteWxGzhUser({
  body,
  options,
}: {
  body: API.DeleteWxGzhUserReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhUserReply>('/admin/v1/wx_gzh_user/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 公众号粉丝表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_user/info */
export function getWxGzhUserInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhUserInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhUserInfoReply>('/admin/v1/wx_gzh_user/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 公众号粉丝表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_user/list */
export function getWxGzhUserList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhUserListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhUserListReply>('/admin/v1/wx_gzh_user/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
