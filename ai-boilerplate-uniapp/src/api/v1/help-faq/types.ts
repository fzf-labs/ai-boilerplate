/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type HelpFaqInfo = {
  /** 主键ID */
  id?: string;
  /** 分类ID */
  categoryId?: string;
  /** 问题 */
  question?: string;
  /** 答案 */
  answer?: string;
  /** 排序号 */
  orderNum?: number;
  /** 查看次数 */
  viewCount?: number;
  /** 有帮助次数 */
  helpfulCount?: number;
  /** 无帮助次数 */
  unhelpfulCount?: number;
  /** 是否热门 */
  isHot?: boolean;
  /** 状态:0-禁用,1-启用 */
  status?: number;
  createdAt?: string;
  updatedAt?: string;
};

export type ListHelpFaqsParams = {
  /** 分类ID（可选，用于筛选） */
  categoryId?: string;
  /** 搜索关键词（可选） */
  keyword?: string;
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type ListHelpFaqsReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: HelpFaqInfo[];
};

export type ListHelpFaqsResponses = {
  /**
   * A successful response.
   */
  200: ListHelpFaqsReply;
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
