/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateWxGzhAutoReplyReply = {
  /** 主键 */
  id?: string;
};

export type CreateWxGzhAutoReplyReq = {
  /** 公众号 appId */
  appId: string;
  /** 回复类型(关键词回复,收到消息回复,被关注回复) */
  type: number;
  /** 请求的关键字 */
  requestKeyword?: string;
  /** 请求的关键字匹配类型 */
  requestKeywordMatch?: number;
  /** 回复的消息类型 */
  responseMessageType: string;
  /** 回复的消息内容 */
  responseContent?: string;
  /** 回复的媒体文件 id */
  responseMediaId?: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type CreateWxGzhAutoReplyResponses = {
  /**
   * A successful response.
   */
  200: CreateWxGzhAutoReplyReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteWxGzhAutoReplyReply = object;

export type DeleteWxGzhAutoReplyReq = {
  /** 主键 */
  id: string;
};

export type DeleteWxGzhAutoReplyResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhAutoReplyReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhAutoReplyInfoParams = {
  /** 主键 */
  id: string;
};

export type GetWxGzhAutoReplyInfoReply = {
  info?: WxGzhAutoReplyInfo;
};

export type GetWxGzhAutoReplyInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhAutoReplyInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhAutoReplyListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 公众号appId */
  appId?: string;
  /** 回复类型(关键词回复,收到消息回复,被关注回复) */
  type?: number;
};

export type GetWxGzhAutoReplyListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhAutoReplyInfo[];
};

export type GetWxGzhAutoReplyListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhAutoReplyListReply;
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

export type UpdateWxGzhAutoReplyReply = object;

export type UpdateWxGzhAutoReplyReq = {
  /** 主键 */
  id: string;
  /** 公众号 appId */
  appId: string;
  /** 回复类型(关键词回复,收到消息回复,被关注回复) */
  type: number;
  /** 请求的关键字 */
  requestKeyword?: string;
  /** 请求的关键字匹配类型 */
  requestKeywordMatch?: number;
  /** 回复的消息类型 */
  responseMessageType: string;
  /** 回复的消息内容 */
  responseContent?: string;
  /** 回复的媒体文件 id */
  responseMediaId?: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateWxGzhAutoReplyResponses = {
  /**
   * A successful response.
   */
  200: UpdateWxGzhAutoReplyReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateWxGzhAutoReplyStatusReply = object;

export type UpdateWxGzhAutoReplyStatusReq = {
  /** 主键 */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateWxGzhAutoReplyStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateWxGzhAutoReplyStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type WxGzhAutoReplyInfo = {
  /** 主键 */
  id?: string;
  /** 公众号 appId */
  appId?: string;
  /** 回复类型(关键词回复,收到消息回复,被关注回复) */
  type?: number;
  /** 请求的关键字 */
  requestKeyword?: string;
  /** 请求的关键字匹配类型 */
  requestKeywordMatch?: number;
  /** 回复的消息类型 */
  responseMessageType?: string;
  /** 回复的消息内容 */
  responseContent?: string;
  /** 回复的媒体文件 id */
  responseMediaId?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
