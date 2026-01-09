/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号消息表 -创建一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_message/create */
export function createWxGzhMessage({
  body,
  options,
}: {
  body: API.CreateWxGzhMessageReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateWxGzhMessageReply>(
    '/admin/v1/wx_gzh_message/create',
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

/** 公众号消息表 -删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_message/delete */
export function deleteWxGzhMessage({
  body,
  options,
}: {
  body: API.DeleteWxGzhMessageReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhMessageReply>(
    '/admin/v1/wx_gzh_message/delete',
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

/** 公众号消息表 -单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_message/info */
export function getWxGzhMessageInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMessageInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMessageInfoReply>(
    '/admin/v1/wx_gzh_message/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号消息表 -列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_message/list */
export function getWxGzhMessageList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhMessageListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhMessageListReply>(
    '/admin/v1/wx_gzh_message/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号消息表 -更新一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_message/update */
export function updateWxGzhMessage({
  body,
  options,
}: {
  body: API.UpdateWxGzhMessageReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateWxGzhMessageReply>(
    '/admin/v1/wx_gzh_message/update',
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
