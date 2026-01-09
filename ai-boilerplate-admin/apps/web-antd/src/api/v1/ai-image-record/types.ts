/* eslint-disable */
// @ts-ignore

export type AiImageRecordInfo = {
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
  /** 图片宽度 */
  width?: number;
  /** 图片高度 */
  height?: number;
  /** 绘画状态 */
  status?: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus?: boolean;
  /** 图片地址 */
  picURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
  /** mj buttons 按钮 */
  buttons?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateAiImageRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiImageRecordReq = {
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
  /** 图片宽度 */
  width: number;
  /** 图片高度 */
  height: number;
  /** 绘画状态 */
  status: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus: boolean;
  /** 图片地址 */
  picURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
  /** mj buttons 按钮 */
  buttons?: string;
};

export type CreateAiImageRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiImageRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiImageRecordReply = object;

export type DeleteAiImageRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiImageRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiImageRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiImageRecordInfoParams = {
  /** 编号 */
  id: string;
};

export type GetAiImageRecordInfoReply = {
  info?: AiImageRecordInfo;
};

export type GetAiImageRecordInfoResponses = {
  /**
   * A successful response.
   */
  200: GetAiImageRecordInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiImageRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiImageRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiImageRecordInfo[];
};

export type GetAiImageRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiImageRecordListReply;
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

export type UpdateAiImageRecordReply = object;

export type UpdateAiImageRecordReq = {
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
  /** 图片宽度 */
  width: number;
  /** 图片高度 */
  height: number;
  /** 绘画状态 */
  status: number;
  /** 完成时间 */
  finishTime?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 是否发布 */
  publicStatus: boolean;
  /** 图片地址 */
  picURL?: string;
  /** 绘制参数 */
  options?: string;
  /** 任务编号 */
  taskId?: string;
  /** mj buttons 按钮 */
  buttons?: string;
};

export type UpdateAiImageRecordResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiImageRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateAiImageRecordStatusReply = object;

export type UpdateAiImageRecordStatusReq = {
  /** 编号 */
  id: string;
  /** 绘画状态 */
  status: number;
};

export type UpdateAiImageRecordStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateAiImageRecordStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
