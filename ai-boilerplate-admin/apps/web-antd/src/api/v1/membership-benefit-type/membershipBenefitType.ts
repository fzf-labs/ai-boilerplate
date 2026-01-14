/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 会员权益类型表-创建一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit_type/create */
export function createMembershipBenefitType({
  body,
  options,
}: {
  body: API.CreateMembershipBenefitTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateMembershipBenefitTypeReply>(
    '/admin/v1/membership_benefit_type/create',
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

/** 会员权益类型表-删除一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit_type/delete */
export function deleteMembershipBenefitType({
  body,
  options,
}: {
  body: API.DeleteMembershipBenefitTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteMembershipBenefitTypeReply>(
    '/admin/v1/membership_benefit_type/delete',
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

/** 会员权益类型表-单条数据查询 返回值: An unexpected error response. GET /admin/v1/membership_benefit_type/info */
export function getMembershipBenefitTypeInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipBenefitTypeInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipBenefitTypeInfoReply>(
    '/admin/v1/membership_benefit_type/info',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    }
  );
}

/** 会员权益类型表-列表数据查询 返回值: An unexpected error response. GET /admin/v1/membership_benefit_type/list */
export function getMembershipBenefitTypeList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetMembershipBenefitTypeListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetMembershipBenefitTypeListReply>(
    '/admin/v1/membership_benefit_type/list',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    }
  );
}

/** 会员权益类型表-更新一条数据 返回值: An unexpected error response. POST /admin/v1/membership_benefit_type/update */
export function updateMembershipBenefitType({
  body,
  options,
}: {
  body: API.UpdateMembershipBenefitTypeReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipBenefitTypeReply>(
    '/admin/v1/membership_benefit_type/update',
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

/** 会员权益类型表-更新状态 返回值: An unexpected error response. POST /admin/v1/membership_benefit_type/update/status */
export function updateMembershipBenefitTypeStatus({
  body,
  options,
}: {
  body: API.UpdateMembershipBenefitTypeStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateMembershipBenefitTypeStatusReply>(
    '/admin/v1/membership_benefit_type/update/status',
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
