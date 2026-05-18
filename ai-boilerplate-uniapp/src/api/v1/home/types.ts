/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type ContentDetail = {
  /** 内容ID */
  id?: string;
  /** 标题 */
  title?: string;
  /** 内容 */
  content?: string;
  /** 封面图 */
  coverImage?: string;
  /** 发布时间 */
  publishTime?: string;
  /** 标签 */
  tags?: string[];
  /** 浏览量 */
  viewCount?: number;
  /** 点赞数 */
  likeCount?: number;
};

export type ContentInfo = {
  /** 内容ID */
  id?: string;
  /** 标题 */
  title?: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverImage?: string;
  /** 发布时间 */
  publishTime?: string;
  /** 标签 */
  tags?: string[];
  /** 是否推荐 */
  isRecommend?: boolean;
  /** 是否热门 */
  isHot?: boolean;
};

export type GetContentDetailParams = {
  /** 内容ID */
  id: string;
};

export type GetContentDetailReply = {
  info?: ContentDetail;
};

export type GetContentDetailResponses = {
  /**
   * A successful response.
   */
  200: GetContentDetailReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetContentListParams = {
  /** 页码 */
  page?: number;
  /** 每页数量 */
  pageSize?: number;
};

export type GetContentListReply = {
  /** 内容列表 */
  list?: ContentInfo[];
  /** 总数 */
  total?: number;
  /** 当前页 */
  page?: number;
  /** 每页数量 */
  pageSize?: number;
};

export type GetContentListResponses = {
  /**
   * A successful response.
   */
  200: GetContentListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};
