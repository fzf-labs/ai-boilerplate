/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 设备表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/device/delete */
export function deleteDevice({
  body,
  options,
}: {
  body: API.DeleteDeviceReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteDeviceReply>('/admin/v1/device/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 设备表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/device/info */
export function getDeviceInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetDeviceInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDeviceInfoReply>('/admin/v1/device/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 设备表-列表数据查询 返回值: An unexpected error response. POST /admin/v1/device/list */
export function getDeviceList({
  body,
  options,
}: {
  body: API.GetDeviceListReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetDeviceListReply>('/admin/v1/device/list', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 设备表-在线设备数量统计 返回值: An unexpected error response. GET /admin/v1/device/online/count */
export function getOnlineDeviceCount({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetOnlineDeviceCountReply>(
    '/admin/v1/device/online/count',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 设备表-注册设备 返回值: An unexpected error response. POST /admin/v1/device/register */
export function registerDevice({
  body,
  options,
}: {
  body: API.RegisterDeviceReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.RegisterDeviceReply>('/admin/v1/device/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 设备表-更新状态 返回值: An unexpected error response. POST /admin/v1/device/update/status */
export function updateDeviceStatus({
  body,
  options,
}: {
  body: API.UpdateDeviceStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateDeviceStatusReply>(
    '/admin/v1/device/update/status',
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
