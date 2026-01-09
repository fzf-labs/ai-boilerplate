/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 文件配置表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/file_config/create */
export function createFileConfig({
  body,
  options,
}: {
  body: API.CreateFileConfigReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateFileConfigReply>('/admin/v1/file_config/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 文件配置表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/file_config/delete */
export function deleteFileConfig({
  body,
  options,
}: {
  body: API.DeleteFileConfigReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteFileConfigReply>('/admin/v1/file_config/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 文件配置表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/file_config/info */
export function getFileConfigInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetFileConfigInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileConfigInfoReply>('/admin/v1/file_config/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 文件配置表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/file_config/list */
export function getFileConfigList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetFileConfigListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileConfigListReply>('/admin/v1/file_config/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 文件配置表-获取所有选择器 返回值: An unexpected error response. GET /admin/v1/file_config/select */
export function getFileConfigSelect({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileConfigSelectReply>('/admin/v1/file_config/select', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 文件配置表-设置主配置 返回值: An unexpected error response. POST /admin/v1/file_config/set_master */
export function setFileConfigMaster({
  body,
  options,
}: {
  body: API.SetFileConfigMasterReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SetFileConfigMasterReply>(
    '/admin/v1/file_config/set_master',
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

/** 文件配置表-获取存储器选择器 返回值: An unexpected error response. GET /admin/v1/file_config/storage/select */
export function getFileConfigStorageSelect({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileConfigStorageSelectReply>(
    '/admin/v1/file_config/storage/select',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 文件配置表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/file_config/update */
export function updateFileConfig({
  body,
  options,
}: {
  body: API.UpdateFileConfigReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateFileConfigReply>('/admin/v1/file_config/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
