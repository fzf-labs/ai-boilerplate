/* eslint-disable */
// @ts-ignore

export type AiIndexVideoRecordInfo = {
  /** 编号 */
  id?: string;
  /** 用户编号 */
  adminId?: string;
  /** 提示词 */
  prompt?: string;
  /** 平台 */
  platform?: string;
  /** 模型编号 */
  modelId?: string;
  /** 模型 */
  model?: string;
  /** 状态 */
  status?: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus?: boolean;
  /** 视频地址 */
  videoURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateAiIndexVideoRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiIndexVideoRecordReq = {
  /** 用户编号 */
  adminId: string;
  /** 提示词 */
  prompt: string;
  /** 平台 */
  platform: string;
  /** 模型编号 */
  modelId?: string;
  /** 模型 */
  model: string;
  /** 状态 */
  status: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus: boolean;
  /** 视频地址 */
  videoURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
};

export type CreateAiIndexVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiIndexVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiIndexVideoRecordReply = object;

export type DeleteAiIndexVideoRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiIndexVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiIndexVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexVideoRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiIndexVideoRecordInfoReply = {
  info?: AiIndexVideoRecordInfo;
};

export type GetAiIndexVideoRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexVideoRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexVideoRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexVideoRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexVideoRecordInfo[];
};

export type GetAiIndexVideoRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexVideoRecordListReply;
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

export type UpdateAiIndexVideoRecordReply = object;

export type UpdateAiIndexVideoRecordReq = {
  /** 编号 */
  id: string;
  /** 用户编号 */
  adminId: string;
  /** 提示词 */
  prompt: string;
  /** 平台 */
  platform: string;
  /** 模型编号 */
  modelId?: string;
  /** 模型 */
  model: string;
  /** 状态 */
  status: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus: boolean;
  /** 视频地址 */
  videoURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
};

export type UpdateAiIndexVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiIndexVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateAiIndexVideoRecordStatusReply = object;

export type UpdateAiIndexVideoRecordStatusReq = {
  /** 编号 */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateAiIndexVideoRecordStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiIndexVideoRecordStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
