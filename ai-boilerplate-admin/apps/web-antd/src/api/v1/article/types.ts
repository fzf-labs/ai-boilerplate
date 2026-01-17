/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type ArticleInfo = {
  id?: string;
  /** 标题 */
  title?: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverImage?: string;
  /** Markdown内容 */
  contentMarkdown?: string;
  /** 状态(-1下线,0草稿,1已发布) */
  status?: number;
  /** 发布时间 */
  publishTime?: string;
  /** 标签 */
  tags?: string[];
  /** 是否推荐 */
  isRecommend?: boolean;
  /** 是否热门 */
  isHot?: boolean;
  /** 浏览量 */
  viewCount?: number;
  /** 点赞数 */
  likeCount?: number;
  createdAt?: string;
  updatedAt?: string;
};

export type CreateArticleReq = {
  title: string;
  summary?: string;
  coverImage?: string;
  contentMarkdown: string;
  status?: number;
  publishTime?: string;
  tags?: string[];
  isRecommend?: boolean;
  isHot?: boolean;
};

export type CreateArticleReply = {
  id?: string;
};

export type UpdateArticleReq = {
  id: string;
  title: string;
  summary?: string;
  coverImage?: string;
  contentMarkdown: string;
  tags?: string[];
  isRecommend?: boolean;
  isHot?: boolean;
};

export type UpdateArticleReply = object;

export type UpdateArticleStatusReq = {
  id: string;
  status: number;
  publishTime?: string;
};

export type UpdateArticleStatusReply = object;

export type DeleteArticleReq = {
  id: string;
};

export type DeleteArticleReply = object;

export type GetArticleInfoParams = {
  id: string;
};

export type GetArticleInfoReply = {
  info?: ArticleInfo;
};

export type GetArticleListParams = {
  page: number;
  pageSize: number;
  status?: number;
  keyword?: string;
};

export type GetArticleListReply = {
  total?: number;
  list?: ArticleInfo[];
};

