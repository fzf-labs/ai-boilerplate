/* eslint-disable */
// @ts-ignore

export type ActivityItem = {
  /** id */
  id?: string;
  /** 标题 */
  title?: string;
  /** 图片URL */
  imageURL?: string;
  /** 跳转链接 */
  linkURL?: string;
  /** 跳转类型 */
  linkType?: string;
  /** 排序 */
  sort?: number;
};

export type Any = {
  '@type'?: string;
};

export type ListActivitiesParams = {
  /** 页码 */
  page?: number;
  /** 每页数量 */
  pageSize?: number;
};

export type ListActivitiesReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: ActivityItem[];
  /** 当前页 */
  page?: number;
  /** 每页数量 */
  pageSize?: number;
};

export type ListActivitiesResponses = {
  /**
   * A successful response.
   */
  200: ListActivitiesReply;
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
