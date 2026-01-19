/* eslint-disable */
// @ts-ignore
import { request } from '#/api/request';

import * as API from './types';

/** 内容-文章-创建一条数据 返回值: An unexpected error response. POST /admin/v1/article/create */
export function createArticle({
  body,
  options,
}: {
  body: API.CreateArticleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.CreateArticleReply>('/admin/v1/article/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-文章-删除一条数据 返回值: An unexpected error response. POST /admin/v1/article/delete */
export function deleteArticle({
  body,
  options,
}: {
  body: API.DeleteArticleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.DeleteArticleReply>('/admin/v1/article/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-文章-单条数据查询 返回值: An unexpected error response. GET /admin/v1/article/info */
export function getArticleInfo({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetArticleInfoParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetArticleInfoReply>('/admin/v1/article/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 内容-文章-列表数据查询 返回值: An unexpected error response. GET /admin/v1/article/list */
export function getArticleList({
  params,
  options,
}: {
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.GetArticleListParams;
  options?: { [key: string]: unknown };
}) {
  return request<API.GetArticleListReply>('/admin/v1/article/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 内容-文章-更新一条数据 返回值: An unexpected error response. POST /admin/v1/article/update */
export function updateArticle({
  body,
  options,
}: {
  body: API.UpdateArticleReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateArticleReply>('/admin/v1/article/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 内容-文章-更新状态(发布/下线/草稿) 返回值: An unexpected error response. POST /admin/v1/article/update/status */
export function updateArticleStatus({
  body,
  options,
}: {
  body: API.UpdateArticleStatusReq;
  options?: { [key: string]: unknown };
}) {
  return request<API.UpdateArticleStatusReply>(
    '/admin/v1/article/update/status',
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
