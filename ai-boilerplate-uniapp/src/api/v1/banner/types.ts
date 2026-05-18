/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type BannerItem = {
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
  /** 展示位置 */
  position?: string;
  /** 平台 */
  platform?: string;
  /** 排序 */
  sort?: number;
};

export type ListBannersParams = {
  /** 展示位置 */
  position: string;
  /** 平台 */
  platform?: string;
};

export type ListBannersReply = {
  /** 列表数据 */
  list?: BannerItem[];
};

export type ListBannersResponses = {
  /**
   * A successful response.
   */
  200: ListBannersReply;
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
