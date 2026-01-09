/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 邮件模版表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/mail_template/create */
export function createMailTemplate({
  body,
  options,
}: {
  body: API.CreateMailTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMailTemplateReply>(
    '/admin/v1/mail_template/create',
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

/** 邮件模版表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/mail_template/delete */
export function deleteMailTemplate({
  body,
  options,
}: {
  body: API.DeleteMailTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMailTemplateReply>(
    '/admin/v1/mail_template/delete',
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

/** 邮件模版表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/mail_template/info */
export function getMailTemplateInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailTemplateInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailTemplateInfoReply>('/admin/v1/mail_template/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 邮件模版表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/mail_template/list */
export function getMailTemplateList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailTemplateListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailTemplateListReply>('/admin/v1/mail_template/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 邮件模版表-选择器 返回值: An unexpected error response. GET /admin/v1/mail_template/selector */
export function getMailTemplateSelector({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMailTemplateSelectorParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMailTemplateSelectorReply>(
    '/admin/v1/mail_template/selector',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 邮件模版表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/mail_template/update */
export function updateMailTemplate({
  body,
  options,
}: {
  body: API.UpdateMailTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMailTemplateReply>(
    '/admin/v1/mail_template/update',
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

/** 邮件模版表-更新状态 返回值: An unexpected error response. POST /admin/v1/mail_template/update/status */
export function updateMailTemplateStatus({
  body,
  options,
}: {
  body: API.UpdateMailTemplateStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMailTemplateStatusReply>(
    '/admin/v1/mail_template/update/status',
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
