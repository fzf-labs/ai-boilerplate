/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 短信日志-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sms_log/info */
export function getSmsLogInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsLogInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsLogInfoReply>('/admin/v1/sms_log/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 短信日志-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sms_log/list */
export function getSmsLogList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsLogListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsLogListReply>('/admin/v1/sms_log/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
