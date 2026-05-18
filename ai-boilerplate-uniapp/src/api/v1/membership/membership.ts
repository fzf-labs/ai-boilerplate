/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 获取会员权益列表 返回值: An unexpected error response. GET /app/v1/membership/benefits */
export function getMembershipBenefits({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipBenefitsParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.GetMembershipBenefitsReply>(
    '/app/v1/membership/benefits',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    }
  );
}

/** 获取会员权益对比（一次返回所有会员类型的权益） 返回值: An unexpected error response. GET /app/v1/membership/benefits/compare */
export function getMembershipBenefitsCompare({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetMembershipBenefitsCompareReply>(
    '/app/v1/membership/benefits/compare',
    {
      method: 'GET',
      ...(options || {}),
    }
  );
}

/** 获取用户会员基础信息 返回值: An unexpected error response. GET /app/v1/membership/info */
export function getUserMembershipInfo({
  options,
}: {
  options?: CustomRequestOptions_;
}) {
  return request<API.GetUserMembershipInfoReply>('/app/v1/membership/info', {
    method: 'GET',
    ...(options || {}),
  });
}
