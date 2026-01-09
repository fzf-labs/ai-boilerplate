/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 配置模型表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_model/create */
export function createAiProviderModel({
  body,
  options,
}: {
  body: API.CreateAiProviderModelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiProviderModelReply>(
    '/admin/v1/ai_provider_model/create',
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

/** AI 配置模型表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_model/delete */
export function deleteAiProviderModel({
  body,
  options,
}: {
  body: API.DeleteAiProviderModelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiProviderModelReply>(
    '/admin/v1/ai_provider_model/delete',
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

/** AI 配置模型表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_provider_model/info */
export function getAiProviderModelInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiProviderModelInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderModelInfoReply>(
    '/admin/v1/ai_provider_model/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 配置模型表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_provider_model/list */
export function getAiProviderModelList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiProviderModelListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderModelListReply>(
    '/admin/v1/ai_provider_model/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 配置模型表-标签选择器 返回值: An unexpected error response. GET /admin/v1/ai_provider_model/selector */
export function getAiProviderModelLabelSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderModelLabelSelectorReply>(
    '/admin/v1/ai_provider_model/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** AI 配置模型表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_model/update */
export function updateAiProviderModel({
  body,
  options,
}: {
  body: API.UpdateAiProviderModelReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiProviderModelReply>(
    '/admin/v1/ai_provider_model/update',
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

/** AI 配置模型表-更新状态 返回值: An unexpected error response. POST /admin/v1/ai_provider_model/update/status */
export function updateAiProviderModelStatus({
  body,
  options,
}: {
  body: API.UpdateAiProviderModelStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiProviderModelStatusReply>(
    '/admin/v1/ai_provider_model/update/status',
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
