/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 短信渠道-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sms_channel/create */
export function createSmsChannel({
  body,
  options,
}: {
  body: API.CreateSmsChannelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSmsChannelReply>('/admin/v1/sms_channel/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信渠道-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sms_channel/delete */
export function deleteSmsChannel({
  body,
  options,
}: {
  body: API.DeleteSmsChannelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSmsChannelReply>('/admin/v1/sms_channel/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信渠道-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sms_channel/info */
export function getSmsChannelInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsChannelInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsChannelInfoReply>('/admin/v1/sms_channel/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 短信渠道-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sms_channel/list */
export function getSmsChannelList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsChannelListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsChannelListReply>('/admin/v1/sms_channel/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 短信渠道-运营商 返回值: An unexpected error response. GET /admin/v1/sms_channel/operator */
export function getSmsChannelOperator({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsChannelOperatorReply>(
    '/admin/v1/sms_channel/operator',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 短信渠道-选择器 返回值: An unexpected error response. GET /admin/v1/sms_channel/selector */
export function getSmsChannelSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsChannelSelectorReply>(
    '/admin/v1/sms_channel/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 短信渠道-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sms_channel/update */
export function updateSmsChannel({
  body,
  options,
}: {
  body: API.UpdateSmsChannelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSmsChannelReply>('/admin/v1/sms_channel/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信渠道-更新状态 返回值: An unexpected error response. POST /admin/v1/sms_channel/update/status */
export function updateSmsChannelStatus({
  body,
  options,
}: {
  body: API.UpdateSmsChannelStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSmsChannelStatusReply>(
    '/admin/v1/sms_channel/update/status',
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
