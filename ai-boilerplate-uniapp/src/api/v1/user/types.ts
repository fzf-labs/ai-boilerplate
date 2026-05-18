/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type BindPhoneReply = object;

export type BindPhoneReq = {
  /** 手机号 */
  phone: string;
  /** 验证码 */
  code: string;
};

export type BindPhoneResponses = {
  /**
   * A successful response.
   */
  200: BindPhoneReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type ChangePasswordReply = object;

export type ChangePasswordReq = {
  /** 旧密码 */
  oldPassword: string;
  /** 新密码 */
  newPassword: string;
  /** 确认新密码 */
  confirmPassword: string;
};

export type ChangePasswordResponses = {
  /**
   * A successful response.
   */
  200: ChangePasswordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteAccountReply = object;

export type DeleteAccountReq = {
  /** 密码 */
  password: string;
};

export type DeleteAccountResponses = {
  /**
   * A successful response.
   */
  200: DeleteAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetUserInfoReply = {
  info?: UserInfo;
};

export type GetUserInfoResponses = {
  /**
   * A successful response.
   */
  200: GetUserInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type LoginReply = {
  /** token */
  token?: string;
  /** 过期时间 */
  expiredAt?: number;
  /** 刷新时间 */
  refreshAt?: number;
};

export type LoginReq = {
  /** 账号(默认使用手机号) */
  username: string;
  /** 密码 */
  password: string;
};

export type LoginResponses = {
  /**
   * A successful response.
   */
  200: LoginReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SendVerifyCodeReply = object;

export type SendVerifyCodeReq = {
  /** 手机号 */
  phone: string;
};

export type SendVerifyCodeResponses = {
  /**
   * A successful response.
   */
  200: SendVerifyCodeReply;
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

export type UpdateUserInfoReply = object;

export type UpdateUserInfoReq = {
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 性别：0-未知，1-男，2-女 */
  gender?: number;
  /** 简介 */
  profile?: string;
};

export type UpdateUserInfoResponses = {
  /**
   * A successful response.
   */
  200: UpdateUserInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UserInfo = {
  /** 用户ID */
  id?: string;
  /** 手机号（脱敏） */
  phone?: string;
  /** 昵称 */
  nickname?: string;
  /** 性别：0-未知，1-男，2-女 */
  gender?: number;
  /** 头像 */
  avatar?: string;
  /** 简介 */
  profile?: string;
  /** 公众号用户Id */
  wxGzhUserId?: string;
  /** 小程序用户Id */
  wxGzhXcxId?: string;
  /** 状态：1-正常 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};
