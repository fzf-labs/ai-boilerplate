/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号账号表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_account/create */
export function createWxGzhAccount({
  body,
  options,
}: {
  body: API.CreateWxGzhAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateWxGzhAccountReply>(
    '/admin/v1/wx_gzh_account/create',
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

/** 公众号账号表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_account/delete */
export function deleteWxGzhAccount({
  body,
  options,
}: {
  body: API.DeleteWxGzhAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhAccountReply>(
    '/admin/v1/wx_gzh_account/delete',
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

/** 公众号账号表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_account/info */
export function getWxGzhAccountInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhAccountInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhAccountInfoReply>(
    '/admin/v1/wx_gzh_account/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号账号表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_account/list */
export function getWxGzhAccountList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhAccountListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhAccountListReply>(
    '/admin/v1/wx_gzh_account/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号账号表-公众号选择器 返回值: An unexpected error response. GET /admin/v1/wx_gzh_account/selector */
export function getWxGzhAccountSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhAccountSelectorReply>(
    '/admin/v1/wx_gzh_account/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 公众号账号表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_account/update */
export function updateWxGzhAccount({
  body,
  options,
}: {
  body: API.UpdateWxGzhAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateWxGzhAccountReply>(
    '/admin/v1/wx_gzh_account/update',
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
