/* eslint-disable */
// @ts-ignore

export type AiIndexImageRecordInfo = {
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

export type CreateAiIndexImageRecordReply = {
  /** 编号 */
  id?: string;
};

export type CreateAiIndexImageRecordReq = {
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

export type CreateAiIndexImageRecordResponses = {
  /**
   * A successful response.
   */
  200: CreateAiIndexImageRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAiIndexImageRecordReply = object;

export type DeleteAiIndexImageRecordReq = {
  /** 编号 */
  id: string;
};

export type DeleteAiIndexImageRecordResponses = {
  /**
   * A successful response.
   */
  200: DeleteAiIndexImageRecordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetAiIndexImageRecordListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
};

export type GetAiIndexImageRecordListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: AiIndexImageRecordInfo[];
};

export type GetAiIndexImageRecordListResponses = {
  /**
   * A successful response.
   */
  200: GetAiIndexImageRecordListReply;
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
