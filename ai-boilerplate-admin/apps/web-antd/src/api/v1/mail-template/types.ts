/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateMailTemplateReply = {
  /** id */
  id?: string;
};

export type CreateMailTemplateReq = {
  /** 模板名称 */
  name: string;
  /** 模板编码 */
  code: string;
  /** 发送的邮箱账号编号 */
  accountId: string;
  /** 发送人名称 */
  nickname?: string;
  /** 模板标题 */
  title: string;
  /** 模板内容 */
  content: string;
  /** 参数数组 */
  params?: string;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status: number;
};

export type CreateMailTemplateResponses = {
  /**
   * A successful response.
   */
  200: CreateMailTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteMailTemplateReply = object;

export type DeleteMailTemplateReq = {
  /** id */
  id: string;
};

export type DeleteMailTemplateResponses = {
  /**
   * A successful response.
   */
  200: DeleteMailTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailTemplateInfoParams = {
  /** id */
  id: string;
};

export type GetMailTemplateInfoReply = {
  info?: MailTemplateInfo;
};

export type GetMailTemplateInfoResponses = {
  /**
   * A successful response.
   */
  200: GetMailTemplateInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailTemplateListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 模板编码 */
  code?: string;
  /** 模板名称 */
  name?: string;
  /** 发送的邮箱账号编号 */
  accountId?: string;
  /** 创建时间 */
  createdAt?: string[];
  /** 开启状态 */
  status?: number;
};

export type GetMailTemplateListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: MailTemplateInfo[];
};

export type GetMailTemplateListResponses = {
  /**
   * A successful response.
   */
  200: GetMailTemplateListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetMailTemplateSelectorParams = {
  /** 发送的邮箱账号编号 */
  accountId?: string;
};

export type GetMailTemplateSelectorReply = {
  /** 列表数据 */
  list?: MailTemplateSelectorItem[];
};

export type GetMailTemplateSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetMailTemplateSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type MailTemplateInfo = {
  /** id */
  id?: string;
  /** 模板名称 */
  name?: string;
  /** 模板编码 */
  code?: string;
  /** 发送的邮箱账号编号 */
  accountId?: string;
  /** 发送人名称 */
  nickname?: string;
  /** 模板标题 */
  title?: string;
  /** 模板内容 */
  content?: string;
  /** 参数数组 */
  params?: string;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type MailTemplateSelectorItem = {
  /** id */
  id?: string;
  /** 模板名称 */
  name?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateMailTemplateReply = object;

export type UpdateMailTemplateReq = {
  /** id */
  id: string;
  /** 模板名称 */
  name: string;
  /** 模板编码 */
  code: string;
  /** 发送的邮箱账号编号 */
  accountId: string;
  /** 发送人名称 */
  nickname?: string;
  /** 模板标题 */
  title: string;
  /** 模板内容 */
  content: string;
  /** 参数数组 */
  params?: string;
  /** 备注 */
  remark?: string;
  /** 开启状态 */
  status: number;
};

export type UpdateMailTemplateResponses = {
  /**
   * A successful response.
   */
  200: UpdateMailTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateMailTemplateStatusReply = object;

export type UpdateMailTemplateStatusReq = {
  /** id */
  id: string;
  /** 开启状态 */
  status: number;
};

export type UpdateMailTemplateStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateMailTemplateStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
