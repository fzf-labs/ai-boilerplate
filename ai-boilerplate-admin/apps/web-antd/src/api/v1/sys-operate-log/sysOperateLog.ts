/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 系统-操作日志-列表数据查询 返回值: An unexpected error response. GET /admin/v1/sys_operate_log/list */
export function getSysOperateLogList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetSysOperateLogListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetSysOperateLogListReply>(
    '/admin/v1/sys_operate_log/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}
