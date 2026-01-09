/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 邮件日志表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/mail_log/delete */
export function deleteMailLog({
  body,
  options,
}: {
  body: API.DeleteMailLogReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMailLogReply>('/admin/v1/mail_log/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 邮件日志表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mail_log/info */
export function getMailLogInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailLogInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailLogInfoReply>('/admin/v1/mail_log/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 邮件日志表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mail_log/list */
export function getMailLogList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailLogListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailLogListReply>('/admin/v1/mail_log/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
