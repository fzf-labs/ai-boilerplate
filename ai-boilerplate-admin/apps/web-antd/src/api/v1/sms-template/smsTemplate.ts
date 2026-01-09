/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 短信模板-创建一条数据 返回值: An unexpected error response. POST /admin/v1/sms_template/create */
export function createSmsTemplate({
  body,
  options,
}: {
  body: API.CreateSmsTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateSmsTemplateReply>('/admin/v1/sms_template/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信模板-删除一条数据 返回值: An unexpected error response. POST /admin/v1/sms_template/delete */
export function deleteSmsTemplate({
  body,
  options,
}: {
  body: API.DeleteSmsTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteSmsTemplateReply>('/admin/v1/sms_template/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信模板-单条数据查询 返回值: An unexpected error response. GET /admin/v1/sms_template/info */
export function getSmsTemplateInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsTemplateInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsTemplateInfoReply>('/admin/v1/sms_template/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 短信模板-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sms_template/list */
export function getSmsTemplateList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSmsTemplateListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsTemplateListReply>('/admin/v1/sms_template/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 短信模板-模板选择器 返回值: An unexpected error response. GET /admin/v1/sms_template/selector */
export function getSmsTemplateSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSmsTemplateSelectorReply>(
    '/admin/v1/sms_template/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 短信模板-发送短信 返回值: An unexpected error response. POST /admin/v1/sms_template/send/msg */
export function sendSmsTemplateMsg({
  body,
  options,
}: {
  body: API.SendSmsTemplateMsgReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SendSmsTemplateMsgReply>(
    '/admin/v1/sms_template/send/msg',
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

/** 短信模板-更新一条数据 返回值: An unexpected error response. POST /admin/v1/sms_template/update */
export function updateSmsTemplate({
  body,
  options,
}: {
  body: API.UpdateSmsTemplateReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSmsTemplateReply>('/admin/v1/sms_template/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 短信模板-更新状态 返回值: An unexpected error response. POST /admin/v1/sms_template/update/status */
export function updateSmsTemplateStatus({
  body,
  options,
}: {
  body: API.UpdateSmsTemplateStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateSmsTemplateStatusReply>(
    '/admin/v1/sms_template/update/status',
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
