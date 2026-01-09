/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type HelpCategoryInfo = {
  /** id */
  id?: string;
  /** 分类名称 */
  name?: string;
  /** 图标 */
  icon?: string;
  /** 排序 */
  order?: number;
  /** 状态(1启用 0禁用) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type ListHelpCategoriesReply = {
  /** 列表数据 */
  list?: HelpCategoryInfo[];
};

export type ListHelpCategoriesResponses = {
  /**
   * A successful response.
   */
  200: ListHelpCategoriesReply;
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
