/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** AI 配置平台表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_platform/create */
export function createAiProviderPlatform({
  body,
  options,
}: {
  body: API.CreateAiProviderPlatformReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateAiProviderPlatformReply>(
    '/admin/v1/ai_provider_platform/create',
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

/** AI 配置平台表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_platform/delete */
export function deleteAiProviderPlatform({
  body,
  options,
}: {
  body: API.DeleteAiProviderPlatformReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteAiProviderPlatformReply>(
    '/admin/v1/ai_provider_platform/delete',
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

/** AI 配置平台表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/ai_provider_platform/info */
export function getAiProviderPlatformInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiProviderPlatformInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderPlatformInfoReply>(
    '/admin/v1/ai_provider_platform/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 配置平台表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/ai_provider_platform/list */
export function getAiProviderPlatformList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetAiProviderPlatformListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderPlatformListReply>(
    '/admin/v1/ai_provider_platform/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** AI 配置平台表-获取平台选择器 返回值: An unexpected error response. GET /admin/v1/ai_provider_platform/selector */
export function getAiProviderPlatformSelector({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetAiProviderPlatformSelectorReply>(
    '/admin/v1/ai_provider_platform/selector',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** AI 配置平台表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/ai_provider_platform/update */
export function updateAiProviderPlatform({
  body,
  options,
}: {
  body: API.UpdateAiProviderPlatformReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiProviderPlatformReply>(
    '/admin/v1/ai_provider_platform/update',
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

/** AI 配置平台表-更新状态 返回值: An unexpected error response. POST /admin/v1/ai_provider_platform/update/status */
export function updateAiProviderPlatformStatus({
  body,
  options,
}: {
  body: API.UpdateAiProviderPlatformStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateAiProviderPlatformStatusReply>(
    '/admin/v1/ai_provider_platform/update/status',
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
