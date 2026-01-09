/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetNotificationSettingsReply = {
  settings?: NotificationSettingsInfo;
};

export type GetNotificationSettingsResponses = {
  /**
   * A successful response.
   */
  200: GetNotificationSettingsReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type NotificationSettingsInfo = {
  /** ID */
  id?: string;
  /** 用户ID */
  userId?: string;
  /** 系统通知 */
  systemNotification?: boolean;
  /** 活动通知 */
  activityNotification?: boolean;
  /** 订单通知 */
  orderNotification?: boolean;
  /** 勿扰开始时间 */
  dndStartTime?: string;
  /** 勿扰结束时间 */
  dndEndTime?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateNotificationSettingsReply = object;

export type UpdateNotificationSettingsReq = {
  /** 系统通知 */
  systemNotification?: boolean;
  /** 活动通知 */
  activityNotification?: boolean;
  /** 订单通知 */
  orderNotification?: boolean;
  /** 勿扰开始时间（格式：HH:mm） */
  dndStartTime?: string;
  /** 勿扰结束时间（格式：HH:mm） */
  dndEndTime?: string;
};

export type UpdateNotificationSettingsResponses = {
  /**
   * A successful response.
   */
  200: UpdateNotificationSettingsReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
