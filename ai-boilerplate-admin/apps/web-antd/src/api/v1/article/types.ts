/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type ArticleInfo = {
  /** id */
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
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type CreateArticleReply = {
  /** id */
  id?: string;
};

export type CreateArticleReq = {
  /** 标题 */
  title: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverImage?: string;
  /** Markdown内容 */
  contentMarkdown: string;
  /** 状态(-1下线,0草稿,1已发布) */
  status?: number;
  /** 发布时间(RFC3339) */
  publishTime?: string;
  /** 标签 */
  tags?: string[];
  /** 是否推荐 */
  isRecommend?: boolean;
  /** 是否热门 */
  isHot?: boolean;
};

export type CreateArticleResponses = {
  /**
   * A successful response.
   */
  200: CreateArticleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteArticleReply = object;

export type DeleteArticleReq = {
  /** id */
  id: string;
};

export type DeleteArticleResponses = {
  /**
   * A successful response.
   */
  200: DeleteArticleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetArticleInfoParams = {
  /** id */
  id: string;
};

export type GetArticleInfoReply = {
  info?: ArticleInfo;
};

export type GetArticleInfoResponses = {
  /**
   * A successful response.
   */
  200: GetArticleInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetArticleListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 状态(-1下线,0草稿,1已发布) */
  status?: number;
  /** 关键词(标题) */
  keyword?: string;
};

export type GetArticleListReply = {
  /** 总数 */
  total?: number;
  /** 列表 */
  list?: ArticleInfo[];
};

export type GetArticleListResponses = {
  /**
   * A successful response.
   */
  200: GetArticleListReply;
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

export type UpdateArticleReply = object;

export type UpdateArticleReq = {
  /** id */
  id: string;
  /** 标题 */
  title: string;
  /** 摘要 */
  summary?: string;
  /** 封面图 */
  coverImage?: string;
  /** Markdown内容 */
  contentMarkdown: string;
  /** 标签 */
  tags?: string[];
  /** 是否推荐 */
  isRecommend?: boolean;
  /** 是否热门 */
  isHot?: boolean;
};

export type UpdateArticleResponses = {
  /**
   * A successful response.
   */
  200: UpdateArticleReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateArticleStatusReply = object;

export type UpdateArticleStatusReq = {
  /** id */
  id: string;
  /** 状态(-1下线,0草稿,1已发布) */
  status: number;
  /** 发布时间(RFC3339) */
  publishTime?: string;
};

export type UpdateArticleStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateArticleStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
