/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 公众号消息自动回复表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_auto_reply/create */
export function createWxGzhAutoReply({
  body,
  options,
}: {
  body: API.CreateWxGzhAutoReplyReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateWxGzhAutoReplyReply>(
    '/admin/v1/wx_gzh_auto_reply/create',
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

/** 公众号消息自动回复表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_auto_reply/delete */
export function deleteWxGzhAutoReply({
  body,
  options,
}: {
  body: API.DeleteWxGzhAutoReplyReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteWxGzhAutoReplyReply>(
    '/admin/v1/wx_gzh_auto_reply/delete',
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

/** 公众号消息自动回复表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_auto_reply/info */
export function getWxGzhAutoReplyInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhAutoReplyInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhAutoReplyInfoReply>(
    '/admin/v1/wx_gzh_auto_reply/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号消息自动回复表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/wx_gzh_auto_reply/list */
export function getWxGzhAutoReplyList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetWxGzhAutoReplyListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetWxGzhAutoReplyListReply>(
    '/admin/v1/wx_gzh_auto_reply/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 公众号消息自动回复表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/wx_gzh_auto_reply/update */
export function updateWxGzhAutoReply({
  body,
  options,
}: {
  body: API.UpdateWxGzhAutoReplyReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateWxGzhAutoReplyReply>(
    '/admin/v1/wx_gzh_auto_reply/update',
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

/** 公众号消息自动回复表-更新状态 返回值: An unexpected error response. POST /admin/v1/wx_gzh_auto_reply/update/status */
export function updateWxGzhAutoReplyStatus({
  body,
  options,
}: {
  body: API.UpdateWxGzhAutoReplyStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateWxGzhAutoReplyStatusReply>(
    '/admin/v1/wx_gzh_auto_reply/update/status',
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
