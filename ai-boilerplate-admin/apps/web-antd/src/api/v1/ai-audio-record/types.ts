/* eslint-disable */
// @ts-ignore

export type AiAudioRecordInfo = {
  /** 编号 */
  id?: string;
  /** 租户编号 */
  tenantId?: string;
  /** 用户编号 */
  adminId?: string;
  /** 音乐名称 */
  title?: string;
  /** 歌词 */
  lyric?: string;
  /** 图片地址 */
  imageURL?: string;
  /** 音频地址 */
  audioURL?: string;
  /** 音频状态 */
  status?: number;
  /** 描述词 */
  description?: string;
  /** 提示词 */
  prompt?: string;
  /** 模型平台 */
  platform?: string;
  /** 模型编号 */
  modelId?: string;
  /** 模型 */
  model?: string;
  /** 生成模式 */
  generateMode?: number;
  /** 风格标签 */
  tags?: string;
  /** 时长 */
  duration?: number;
  /** 是否发布 */
  publicStatus?: boolean;
  /** 任务编号 */
  taskId?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type DeleteAiAudioRecordReply = object;

export type DeleteAiAudioRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiAudioRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiAudioRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiAudioRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiAudioRecordInfoReply = {
  info?: AiAudioRecordInfo;
};

export type GetAiAudioRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiAudioRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiAudioRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiAudioRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiAudioRecordInfo[];
};

export type GetAiAudioRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiAudioRecordListReply;
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
