/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type ButtonItem = {
  /** 按钮类型 */
  type?: string;
  /** 按钮名称 */
  name?: string;
  /** URL（可选） */
  url?: string;
  /** 小程序appid（可选） */
  appid?: string;
  /** 小程序页面路径（可选） */
  pagepath?: string;
  /** 按钮key（可选） */
  key?: string;
  /** 按钮值（可选） */
  value?: string;
  news_info?: NewsInfo;
};

export type DeleteWxGzhMenuReply = object;

export type DeleteWxGzhMenuReq = {
  /** 主键 */
  appId: string;
};

export type DeleteWxGzhMenuResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMenuInfoParams = {
  /** 主键 */
  appId: string;
};

export type GetWxGzhMenuInfoReply = {
  info?: WxGzhMenuInfo;
};

export type GetWxGzhMenuInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMenuInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMenuListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetWxGzhMenuListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhMenuInfo[];
};

export type GetWxGzhMenuListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMenuListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type NewsInfo = {
  list?: NewsItem[];
};

export type NewsItem = {
  /** 标题 */
  title?: string;
  /** 作者 */
  author?: string;
  /** 摘要 */
  digest?: string;
  /** 是否显示封面 */
  show_cover?: number;
  /** 封面URL */
  cover_url?: string;
  /** 内容URL */
  content_url?: string;
  /** 来源URL */
  source_url?: string;
};

export type SelfButton = {
  /** 按钮类型 */
  type?: string;
  /** 按钮名称 */
  name?: string;
  /** 按钮key */
  key?: string;
  /** 按钮值（可选） */
  value?: string;
  /** URL（可选） */
  url?: string;
  sub_button?: SubButton;
};

export type SelfmenuInfo = {
  /** 按钮列表 */
  button?: SelfButton[];
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type StoreWxGzhMenuReply = object;

export type StoreWxGzhMenuReq = {
  /** 微信公众号 appid */
  appId: string;
  selfmenuInfo?: SelfmenuInfo;
};

export type StoreWxGzhMenuResponses = {
  /**
   * A successful response.
   */
  200: StoreWxGzhMenuReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SubButton = {
  list?: ButtonItem[];
};

export type WxGzhMenuInfo = {
  /** 主键 */
  id?: string;
  /** 微信公众号 appid */
  appId?: string;
  /** 菜单是否开启，0代表未开启，1代表开启 */
  isMenuOpen?: number;
  selfmenuInfo?: SelfmenuInfo;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
