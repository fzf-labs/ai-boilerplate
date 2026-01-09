/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMailAccountReply = {
  /** id */
  id?: string;
};

export type CreateMailAccountReq = {
  /** 邮箱 */
  mail: string;
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
  /** SMTP 服务器域名 */
  host: string;
  /** SMTP 服务器端口 */
  port: number;
  /** 是否开启 SSL */
  sslEnable?: boolean;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status?: number;
};

export type CreateMailAccountResponses = {
  /**
   * A successful response.
   */
  200: CreateMailAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMailAccountReply = object;

export type DeleteMailAccountReq = {
  /** id */
  id: string;
};

export type DeleteMailAccountResponses = {
  /**
   * A successful response.
   */
  200: DeleteMailAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailAccountInfoParams = {
  /** id */
  id: string;
};

export type GetMailAccountInfoReply = {
  info?: MailAccountInfo;
};

export type GetMailAccountInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMailAccountInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailAccountListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 邮箱 */
  mail?: string;
  /** 用户名 */
  username?: string;
};

export type GetMailAccountListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MailAccountInfo[];
};

export type GetMailAccountListResponses = {
  /**
   * A successful response.
   */
  200: GetMailAccountListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailAccountSelectorReply = {
  /** 列表数据 */
  list?: MailAccountSelector[];
};

export type GetMailAccountSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetMailAccountSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MailAccountInfo = {
  /** id */
  id?: string;
  /** 邮箱 */
  mail?: string;
  /** 用户名 */
  username?: string;
  /** 密码 */
  password?: string;
  /** SMTP 服务器域名 */
  host?: string;
  /** SMTP 服务器端口 */
  port?: number;
  /** 是否开启 SSL */
  sslEnable?: boolean;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type MailAccountSelector = {
  /** id */
  id?: string;
  /** 邮箱 */
  mail?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMailAccountReply = object;

export type UpdateMailAccountReq = {
  /** id */
  id: string;
  /** 邮箱 */
  mail: string;
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
  /** SMTP 服务器域名 */
  host: string;
  /** SMTP 服务器端口 */
  port: number;
  /** 是否开启 SSL */
  sslEnable?: boolean;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status?: number;
};

export type UpdateMailAccountResponses = {
  /**
   * A successful response.
   */
  200: UpdateMailAccountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMailAccountStatusReply = object;

export type UpdateMailAccountStatusReq = {
  /** id */
  id: string;
  /** 开启状态 */
  status: number;
};

export type UpdateMailAccountStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMailAccountStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
