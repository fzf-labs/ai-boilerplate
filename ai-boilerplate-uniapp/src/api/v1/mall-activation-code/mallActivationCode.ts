/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 会员激活码激活 返回值: An unexpected error response. POST /app/v1/mall_activation_code/activate */
export function activateMembershipByCode({
  body,
  options,
}: {
  body: API.ActivateMembershipByCodeReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.ActivateMembershipByCodeReply>(
    '/app/v1/mall_activation_code/activate',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    }
  );
}

/** 会员激活码兑换记录 返回值: An unexpected error response. GET /app/v1/mall_activation_code/redemptions */
export function listActivationCodeRedemptions({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.ListActivationCodeRedemptionsParams;
  options?: CustomRequestOptions_;
}) {
  return request<API.ListActivationCodeRedemptionsReply>(
    '/app/v1/mall_activation_code/redemptions',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    }
  );
}
