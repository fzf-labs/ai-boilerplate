/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateWxGzhMessageReply = {
  /** 主键 */
  id?: string;
};

export type CreateWxGzhMessageReq = {
  /** 公众号 appId */
  appId: string;
  /** 微信公众号的消息编号 */
  msgId?: string;
  /** 公众号粉丝的编号 */
  userId: string;
  /** 公众号粉丝标志 */
  openid: string;
  /** 消息类型 */
  messageType: string;
  /** 消息来源 */
  sendFrom: number;
  /** 消息内容 */
  content?: string;
  /** 媒体文件 id */
  mediaId?: string;
  /** 媒体文件 URL */
  mediaURL?: string;
  /** 语音识别后文本 */
  recognition?: string;
  /** 语音格式 */
  format?: string;
  /** 标题 */
  title?: string;
  /** 描述 */
  description?: string;
  /** 缩略图的媒体 id */
  thumbMediaId?: string;
  /** 缩略图的媒体 URL */
  thumbMediaURL?: string;
  /** 点击图文消息跳转链接 */
  URL?: string;
  /** 地理位置维度 */
  locationX?: number;
  /** 地理位置经度 */
  locationY?: number;
  /** 地图缩放大小 */
  scale?: number;
  /** 详细地址 */
  label?: string;
  /** 图文消息数组 */
  articles?: string;
  /** 音乐链接 */
  musicURL?: string;
  /** 高质量音乐链接 */
  hqMusicURL?: string;
  /** 事件类型 */
  event?: string;
  /** 事件 Key */
  eventKey?: string;
};

export type CreateWxGzhMessageResponses = {
  /**
   * A successful response.
   */
  200: CreateWxGzhMessageReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteWxGzhMessageReply = object;

export type DeleteWxGzhMessageReq = {
  /** 主键 */
  id: string;
};

export type DeleteWxGzhMessageResponses = {
  /**
   * A successful response.
   */
  200: DeleteWxGzhMessageReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMessageInfoParams = {
  /** 主键 */
  id: string;
};

export type GetWxGzhMessageInfoReply = {
  info?: WxGzhMessageInfo;
};

export type GetWxGzhMessageInfoResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMessageInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetWxGzhMessageListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetWxGzhMessageListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: WxGzhMessageInfo[];
};

export type GetWxGzhMessageListResponses = {
  /**
   * A successful response.
   */
  200: GetWxGzhMessageListReply;
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

export type UpdateWxGzhMessageReply = object;

export type UpdateWxGzhMessageReq = {
  /** 主键 */
  id: string;
  /** 公众号 appId */
  appId: string;
  /** 微信公众号的消息编号 */
  msgId?: string;
  /** 公众号粉丝的编号 */
  userId: string;
  /** 公众号粉丝标志 */
  openid: string;
  /** 消息类型 */
  messageType: string;
  /** 消息来源 */
  sendFrom: number;
  /** 消息内容 */
  content?: string;
  /** 媒体文件 id */
  mediaId?: string;
  /** 媒体文件 URL */
  mediaURL?: string;
  /** 语音识别后文本 */
  recognition?: string;
  /** 语音格式 */
  format?: string;
  /** 标题 */
  title?: string;
  /** 描述 */
  description?: string;
  /** 缩略图的媒体 id */
  thumbMediaId?: string;
  /** 缩略图的媒体 URL */
  thumbMediaURL?: string;
  /** 点击图文消息跳转链接 */
  URL?: string;
  /** 地理位置维度 */
  locationX?: number;
  /** 地理位置经度 */
  locationY?: number;
  /** 地图缩放大小 */
  scale?: number;
  /** 详细地址 */
  label?: string;
  /** 图文消息数组 */
  articles?: string;
  /** 音乐链接 */
  musicURL?: string;
  /** 高质量音乐链接 */
  hqMusicURL?: string;
  /** 事件类型 */
  event?: string;
  /** 事件 Key */
  eventKey?: string;
};

export type UpdateWxGzhMessageResponses = {
  /**
   * A successful response.
   */
  200: UpdateWxGzhMessageReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type WxGzhMessageInfo = {
  /** 主键 */
  id?: string;
  /** 公众号 appId */
  appId?: string;
  /** 微信公众号的消息编号 */
  msgId?: string;
  /** 公众号粉丝的编号 */
  userId?: string;
  /** 公众号粉丝标志 */
  openid?: string;
  /** 消息类型 */
  messageType?: string;
  /** 消息来源 */
  sendFrom?: number;
  /** 消息内容 */
  content?: string;
  /** 媒体文件 id */
  mediaId?: string;
  /** 媒体文件 URL */
  mediaURL?: string;
  /** 语音识别后文本 */
  recognition?: string;
  /** 语音格式 */
  format?: string;
  /** 标题 */
  title?: string;
  /** 描述 */
  description?: string;
  /** 缩略图的媒体 id */
  thumbMediaId?: string;
  /** 缩略图的媒体 URL */
  thumbMediaURL?: string;
  /** 点击图文消息跳转链接 */
  URL?: string;
  /** 地理位置维度 */
  locationX?: number;
  /** 地理位置经度 */
  locationY?: number;
  /** 地图缩放大小 */
  scale?: number;
  /** 详细地址 */
  label?: string;
  /** 图文消息数组 */
  articles?: string;
  /** 音乐链接 */
  musicURL?: string;
  /** 高质量音乐链接 */
  hqMusicURL?: string;
  /** 事件类型 */
  event?: string;
  /** 事件 Key */
  eventKey?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
