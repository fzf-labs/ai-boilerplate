/* eslint-disable */
// @ts-ignore

export type AiVideoRecordInfo = {
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

export type CreateAiVideoRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiVideoRecordReq = {
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

export type CreateAiVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiVideoRecordReply = object;

export type DeleteAiVideoRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiVideoRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiVideoRecordInfoReply = {
  info?: AiVideoRecordInfo;
};

export type GetAiVideoRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiVideoRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiVideoRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiVideoRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiVideoRecordInfo[];
};

export type GetAiVideoRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiVideoRecordListReply;
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

export type UpdateAiVideoRecordReply = object;

export type UpdateAiVideoRecordReq = {
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

export type UpdateAiVideoRecordResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiVideoRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateAiVideoRecordStatusReply = object;

export type UpdateAiVideoRecordStatusReq = {
  /** 编号 */
  id: string;
  /** 状态 */
  status: number;
};

export type UpdateAiVideoRecordStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiVideoRecordStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
