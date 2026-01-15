/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 用户会员变更记录表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/user_membership_change/list */
export function getUserMembershipChangeList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetUserMembershipChangeListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetUserMembershipChangeListReply>(
    '/admin/v1/user_membership_change/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    }
  );
}
