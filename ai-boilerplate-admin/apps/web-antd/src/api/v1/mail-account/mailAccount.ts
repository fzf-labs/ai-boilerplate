/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 邮箱账号表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/mail_account/create */
export function createMailAccount({
  body,
  options,
}: {
  body: API.CreateMailAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMailAccountReply>('/admin/v1/mail_account/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 邮箱账号表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/mail_account/delete */
export function deleteMailAccount({
  body,
  options,
}: {
  body: API.DeleteMailAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMailAccountReply>('/admin/v1/mail_account/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 邮箱账号表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mail_account/info */
export function getMailAccountInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailAccountInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailAccountInfoReply>('/admin/v1/mail_account/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 邮箱账号表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mail_account/list */
export function getMailAccountList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailAccountListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailAccountListReply>('/admin/v1/mail_account/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 邮箱账号表-选择器 返回值: An unexpected error response. GET /admin/v1/mail_account/selector */
export function getMailAccountSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailAccountSelectorReply>(
    '/admin/v1/mail_account/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 邮箱账号表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/mail_account/update */
export function updateMailAccount({
  body,
  options,
}: {
  body: API.UpdateMailAccountReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMailAccountReply>('/admin/v1/mail_account/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 邮箱账号表-更新状态 返回值: An unexpected error response. POST /admin/v1/mail_account/update/status */
export function updateMailAccountStatus({
  body,
  options,
}: {
  body: API.UpdateMailAccountStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMailAccountStatusReply>(
    '/admin/v1/mail_account/update/status',
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
