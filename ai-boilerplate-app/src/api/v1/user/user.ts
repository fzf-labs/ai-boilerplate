/* eslint-disable */
// @ts-ignore
import request from '@/http/vue-query';
import { CustomRequestOptions_ } from '@/http/types';

import * as API from './types';

/** 注销账号 返回值: An unexpected error response. POST /app/v1/user/account/delete */
export function deleteAccount({
  body,
  options,
}: {
  body: API.DeleteAccountReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.DeleteAccountReply>('/app/v1/user/account/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 发送验证码 返回值: An unexpected error response. POST /app/v1/user/code/send */
export function sendVerifyCode({
  body,
  options,
}: {
  body: API.SendVerifyCodeReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.SendVerifyCodeReply>('/app/v1/user/code/send', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 登录 返回值: An unexpected error response. POST /app/v1/user/login */
export function login({
  body,
  options,
}: {
  body: API.LoginReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.LoginReply>('/app/v1/user/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 修改密码 返回值: An unexpected error response. POST /app/v1/user/password/change */
export function changePassword({
  body,
  options,
}: {
  body: API.ChangePasswordReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.ChangePasswordReply>('/app/v1/user/password/change', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 绑定手机号 返回值: An unexpected error response. POST /app/v1/user/phone/bind */
export function bindPhone({
  body,
  options,
}: {
  body: API.BindPhoneReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.BindPhoneReply>('/app/v1/user/phone/bind', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取用户详情 返回值: An unexpected error response. GET /app/v1/user/profile */
export function getUserInfo({ options }: { options?: CustomRequestOptions_ }) {
  return request<API.GetUserInfoReply>('/app/v1/user/profile', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新用户信息 返回值: An unexpected error response. POST /app/v1/user/profile/update */
export function updateUserInfo({
  body,
  options,
}: {
  body: API.UpdateUserInfoReq;
  options?: CustomRequestOptions_;
}) {
  return request<API.UpdateUserInfoReply>('/app/v1/user/profile/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
