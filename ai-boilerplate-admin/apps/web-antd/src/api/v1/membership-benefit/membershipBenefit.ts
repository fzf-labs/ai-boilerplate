/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 会员权益配置表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit/create */
export function createMembershipBenefit({
  body,
  options,
}: {
  body: API.CreateMembershipBenefitReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMembershipBenefitReply>(
    '/admin/v1/membership_benefit/create',
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

/** 会员权益配置表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit/delete */
export function deleteMembershipBenefit({
  body,
  options,
}: {
  body: API.DeleteMembershipBenefitReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMembershipBenefitReply>(
    '/admin/v1/membership_benefit/delete',
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

/** 会员权益配置表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/membership_benefit/info */
export function getMembershipBenefitInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipBenefitInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipBenefitInfoReply>(
    '/admin/v1/membership_benefit/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 会员权益配置表-获取权益标识选择器 返回值: An unexpected error response. GET /admin/v1/membership_benefit/key/select */
export function getMembershipBenefitKeySelect({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipBenefitKeySelectReply>(
    '/admin/v1/membership_benefit/key/select',
    {
      method: 'GET',
      ...(options || {}),
    },
  );
}

/** 会员权益配置表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/membership_benefit/list */
export function getMembershipBenefitList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipBenefitListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipBenefitListReply>(
    '/admin/v1/membership_benefit/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
}

/** 会员权益配置表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit/update */
export function updateMembershipBenefit({
  body,
  options,
}: {
  body: API.UpdateMembershipBenefitReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipBenefitReply>(
    '/admin/v1/membership_benefit/update',
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

/** 会员权益配置表-更新状态 返回值: An unexpected error response. POST /admin/v1/membership_benefit/update/status */
export function updateMembershipBenefitStatus({
  body,
  options,
}: {
  body: API.UpdateMembershipBenefitStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipBenefitStatusReply>(
    '/admin/v1/membership_benefit/update/status',
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
