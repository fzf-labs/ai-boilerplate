/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type SubmitFeedbackReply = {
  /** 反馈ID */
  id?: string;
};

export type SubmitFeedbackReq = {
  /** 问题分类 */
  category: string;
  /** 问题描述 */
  description: string;
  /** 图片列表（JSON数组） */
  images?: string;
  /** 联系方式 */
  contact?: string;
};

export type SubmitFeedbackResponses = {
  /**
   * A successful response.
   */
  200: SubmitFeedbackReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
