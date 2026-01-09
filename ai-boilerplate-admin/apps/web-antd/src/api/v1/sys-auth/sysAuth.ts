/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** Auth-查询用户信息 返回值: An unexpected error response. GET /admin/v1/sys_auth/admin_info */
export function sysAuthAdminInfo({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthAdminInfoReply>('/admin/v1/sys_auth/admin_info', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Auth-登录 返回值: An unexpected error response. POST /admin/v1/sys_auth/login */
export function sysAuthLogin({
  body,
  options,
}: {
  body: API.SysAuthLoginReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthLoginReply>('/admin/v1/sys_auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Auth-退出 返回值: An unexpected error response. POST /admin/v1/sys_auth/logout */
export function sysAuthLogout({
  body,
  options,
}: {
  body: API.SysAuthLogoutReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthLogoutReply>('/admin/v1/sys_auth/logout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Auth-获取菜单 返回值: An unexpected error response. GET /admin/v1/sys_auth/menu */
export function sysAuthMenu({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthMenuReply>('/admin/v1/sys_auth/menu', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Auth-获取权限 返回值: An unexpected error response. GET /admin/v1/sys_auth/permission */
export function sysAuthPermission({
  options,
}: {
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthPermissionReply>('/admin/v1/sys_auth/permission', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Auth-更新用户信息 返回值: An unexpected error response. POST /admin/v1/sys_auth/update/admin_info */
export function sysAuthUpdateAdminInfo({
  body,
  options,
}: {
  body: API.SysAuthUpdateAdminInfoReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthUpdateAdminInfoReply>(
    '/admin/v1/sys_auth/update/admin_info',
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

/** Auth-更新密码 返回值: An unexpected error response. POST /admin/v1/sys_auth/update/admin_password */
export function sysAuthUpdateAdminPassword({
  body,
  options,
}: {
  body: API.SysAuthUpdateAdminPasswordReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.SysAuthUpdateAdminPasswordReply>(
    '/admin/v1/sys_auth/update/admin_password',
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
