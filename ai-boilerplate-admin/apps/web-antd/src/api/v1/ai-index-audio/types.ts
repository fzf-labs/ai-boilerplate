/* eslint-disable */
// @ts-ignore

export type AiIndexAudioRecordInfo = {
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

export type CreateAiIndexAudioRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiIndexAudioRecordReq = {
  /** 租户编号 */
  tenantId: string;
  /** 用户编号 */
  adminId: string;
  /** 音乐名称 */
  title: string;
  /** 歌词 */
  lyric?: string;
  /** 图片地址 */
  imageURL?: string;
  /** 音频地址 */
  audioURL?: string;
  /** 音频状态 */
  status: number;
  /** 描述词 */
  description?: string;
  /** 提示词 */
  prompt?: string;
  /** 模型平台 */
  platform: string;
  /** 模型编号 */
  modelId: string;
  /** 模型 */
  model: string;
  /** 生成模式 */
  generateMode: number;
  /** 风格标签 */
  tags?: string;
  /** 时长 */
  duration?: number;
  /** 是否发布 */
  publicStatus: boolean;
  /** 任务编号 */
  taskId?: string;
  /** 错误信息 */
  errorMessage?: string;
};

export type CreateAiIndexAudioRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiIndexAudioRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiIndexAudioRecordReply = object;

export type DeleteAiIndexAudioRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiIndexAudioRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiIndexAudioRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexAudioRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiIndexAudioRecordInfoReply = {
  info?: AiIndexAudioRecordInfo;
};

export type GetAiIndexAudioRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexAudioRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexAudioRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexAudioRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexAudioRecordInfo[];
};

export type GetAiIndexAudioRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexAudioRecordListReply;
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
