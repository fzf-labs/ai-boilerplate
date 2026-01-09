/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 文件表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/file_data/delete */
export function deleteFileDatum({
  body,
  options,
}: {
  body: API.DeleteFileDatumReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteFileDatumReply>('/admin/v1/file_data/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 文件表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/file_data/info */
export function getFileDatumInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetFileDatumInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileDatumInfoReply>('/admin/v1/file_data/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 文件表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/file_data/list */
export function getFileDatumList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetFileDatumListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetFileDatumListReply>('/admin/v1/file_data/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 客户端上传-默认上传到 OSS 的方式和凭证获取 返回值: An unexpected error response. GET /admin/v1/file_data/upload/oss_default_policy */
export function uploadFileOssDefaultPolicy({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.UploadFileOSSDefaultPolicyParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.UploadFileOSSDefaultPolicyReply>(
    '/admin/v1/file_data/upload/oss_default_policy',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
