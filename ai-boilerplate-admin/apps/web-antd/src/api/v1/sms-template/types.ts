/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSmsTemplateReply = {
  /** id */
  id?: string;
};

export type CreateSmsTemplateReq = {
  /** 短信渠道编号 */
  smsChannelId: string;
  /** 模板类型 */
  templateType: number;
  /** 模板编码 */
  templateCode: string;
  /** 模板名称 */
  templateName: string;
  /** 模板内容 */
  templateContent: string;
  /** 模板参数 */
  templateParams?: string;
  /** 备注 */
  remark?: string;
  /** 短信供应商的模板编号 */
  apiTemplateId: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type CreateSmsTemplateResponses = {
  /**
   * A successful response.
   */
  200: CreateSmsTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSmsTemplateReply = object;

export type DeleteSmsTemplateReq = {
  /** id */
  id: string;
};

export type DeleteSmsTemplateResponses = {
  /**
   * A successful response.
   */
  200: DeleteSmsTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsTemplateInfoParams = {
  /** id */
  id: string;
};

export type GetSmsTemplateInfoReply = {
  info?: SmsTemplateInfo;
};

export type GetSmsTemplateInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSmsTemplateInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsTemplateListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
};

export type GetSmsTemplateListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SmsTemplateInfo[];
};

export type GetSmsTemplateListResponses = {
  /**
   * A successful response.
   */
  200: GetSmsTemplateListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSmsTemplateSelectorReply = {
  /** 列表数据 */
  list?: SmsTemplateSelector[];
};

export type GetSmsTemplateSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSmsTemplateSelectorReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SendSmsTemplateMsgReply = object;

export type SendSmsTemplateMsgReq = {
  /** id */
  id: string;
  /** 手机号 */
  phone: string;
  /** 模板参数 */
  params?: Record<string, string>;
};

export type SendSmsTemplateMsgResponses = {
  /**
   * A successful response.
   */
  200: SendSmsTemplateMsgReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SmsTemplateInfo = {
  /** id */
  id?: string;
  /** 短信渠道编号 */
  smsChannelId?: string;
  /** 模板类型 */
  templateType?: number;
  /** 模板编码 */
  templateCode?: string;
  /** 模板名称 */
  templateName?: string;
  /** 模板内容 */
  templateContent?: string;
  /** 模板参数 */
  templateParams?: string;
  /** 备注 */
  remark?: string;
  /** 短信供应商的模板编号 */
  apiTemplateId?: string;
  /** 状态(-1禁用,1开启) */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 短信渠道名称 */
  smsChannelName?: string;
};

export type SmsTemplateSelector = {
  /** id */
  id?: string;
  /** 模板名称 */
  templateName?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type UpdateSmsTemplateReply = object;

export type UpdateSmsTemplateReq = {
  /** id */
  id: string;
  /** 短信渠道编号 */
  smsChannelId: string;
  /** 模板类型 */
  templateType: number;
  /** 模板编码 */
  templateCode: string;
  /** 模板名称 */
  templateName: string;
  /** 模板内容 */
  templateContent: string;
  /** 模板参数 */
  templateParams?: string;
  /** 备注 */
  remark?: string;
  /** 短信供应商的模板编号 */
  apiTemplateId: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateSmsTemplateResponses = {
  /**
   * A successful response.
   */
  200: UpdateSmsTemplateReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSmsTemplateStatusReply = object;

export type UpdateSmsTemplateStatusReq = {
  /** id */
  id: string;
  /** 状态(-1禁用,1开启) */
  status: number;
};

export type UpdateSmsTemplateStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSmsTemplateStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
