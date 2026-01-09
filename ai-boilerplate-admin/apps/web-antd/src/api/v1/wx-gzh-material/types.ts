/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type DeleteWxGzhMaterialReply = object;

export type DeleteWxGzhMaterialReq = {
  /** 主键 */
  ids: string[];
};

export type DeleteWxGzhMaterialResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhMaterialReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMaterialInfoParams = {
  /** 主键 */
  id: string;
};

export type GetWxGzhMaterialInfoReply = {
  info?: WxGzhMaterialInfo;
};

export type GetWxGzhMaterialInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMaterialInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMaterialListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 公众号 appId */
  appId: string;
  /** 文件类型 */
  type: string;
  /** 文件名称 */
  name?: string;
};

export type GetWxGzhMaterialListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhMaterialInfo[];
};

export type GetWxGzhMaterialListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMaterialListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMaterialStatsParams = {
  /** 公众号 appId */
  appId: string;
};

export type GetWxGzhMaterialStatsReply = {
  /** 总数 */
  totalCount?: number;
  /** 图片数量 */
  imageCount?: number;
  /** 语音数量 */
  voiceCount?: number;
  /** 视频数量 */
  videoCount?: number;
};

export type GetWxGzhMaterialStatsResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMaterialStatsReply;
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

export type SyncWxGzhMaterialReply = object;

export type SyncWxGzhMaterialReq = {
  /** 公众号 appId */
  appId: string;
};

export type SyncWxGzhMaterialResponses = {
  /**
   * A successful response.
   */
  200: SyncWxGzhMaterialReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type WxGzhMaterialInfo = {
  /** 主键 */
  id?: string;
  /** 公众号 appId */
  appId?: string;
  /** 素材的类型，图片（image）、视频（video）、语音 （voice） */
  type?: string;
  /** 消息ID */
  mediaId?: string;
  /** 标签 */
  tags?: string[];
  /** 更新日期 */
  updateTime?: string;
  /** 图片、语音、视频素材的名字 */
  name?: string;
  /** 图片、语音、视频素材URL */
  URL?: string;
  /** 视频封面 URL */
  coverURL?: string;
  /** 视频描述 */
  description?: string;
  /** 视频分类 */
  newcat?: string;
  /** 视频子分类 */
  newsubcat?: string;
  /** 视频 ID */
  vid?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
