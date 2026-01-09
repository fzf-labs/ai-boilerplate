/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSelfAppReleaseReply = {
  /** ID */
  id?: string;
};

export type CreateSelfAppReleaseReq = {
  /** 发布渠道 */
  channel: string;
  /** 包名 */
  packageName: string;
  /** build值 */
  buildNum: number;
  /** 版本号 */
  version?: string;
  /** 更新类型(1强制 2提示 3静默) */
  updateType: number;
  /** 更新标题 */
  title: string;
  /** 更新日志 */
  changelog?: string;
  /** 安装包地址 */
  packageURL: string;
  /** 安装包大小 */
  packageSize?: number;
  /** 安装包MD5 */
  packageMd5?: string;
  /** 最低系统版本 */
  minOsVersion?: string;
  /** 灰度策略(1全量 2按比例 3自定义) */
  grayStrategy: number;
  /** 灰度比例 */
  grayPercentage?: number;
  /** 灰度设备 */
  graySns?: string[];
  /** 发布时间 */
  publishTime?: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type CreateSelfAppReleaseResponses = {
  /**
   * A successful response.
   */
  200: CreateSelfAppReleaseReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSelfAppReleaseReply = object;

export type DeleteSelfAppReleaseReq = {
  /** ID */
  id: string;
};

export type DeleteSelfAppReleaseResponses = {
  /**
   * A successful response.
   */
  200: DeleteSelfAppReleaseReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSelfAppReleaseInfoParams = {
  /** ID */
  id: string;
};

export type GetSelfAppReleaseInfoReply = {
  info?: SelfAppReleaseInfo;
};

export type GetSelfAppReleaseInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSelfAppReleaseInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSelfAppReleaseListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 包名 */
  packageName: string;
  /** 发布渠道 */
  channel?: string;
  /** build值 */
  buildNum?: string;
};

export type GetSelfAppReleaseListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SelfAppReleaseInfo[];
};

export type GetSelfAppReleaseListResponses = {
  /**
   * A successful response.
   */
  200: GetSelfAppReleaseListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type SelfAppReleaseInfo = {
  /** ID */
  id?: string;
  /** 发布渠道 */
  channel?: string;
  /** 包名 */
  packageName?: string;
  /** build值 */
  buildNum?: number;
  /** 版本号 */
  version?: string;
  /** 更新类型(1强制 2提示 3静默) */
  updateType?: number;
  /** 更新标题 */
  title?: string;
  /** 更新日志 */
  changelog?: string;
  /** 安装包地址 */
  packageURL?: string;
  /** 安装包大小 */
  packageSize?: number;
  /** 安装包MD5 */
  packageMd5?: string;
  /** 最低系统版本 */
  minOsVersion?: string;
  /** 灰度策略(1全量 2按比例 3自定义) */
  grayStrategy?: number;
  /** 灰度比例 */
  grayPercentage?: number;
  /** 灰度设备 */
  graySns?: string[];
  /** 发布时间 */
  publishTime?: string;
  /** 状态(-1禁用 1启用) */
  status?: number;
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

export type UpdateSelfAppReleaseReply = object;

export type UpdateSelfAppReleaseReq = {
  /** ID */
  id: string;
  /** 发布渠道 */
  channel: string;
  /** 包名 */
  packageName: string;
  /** build值 */
  buildNum: number;
  /** 版本号 */
  version?: string;
  /** 更新类型(1强制 2提示 3静默) */
  updateType: number;
  /** 更新标题 */
  title: string;
  /** 更新日志 */
  changelog?: string;
  /** 安装包地址 */
  packageURL: string;
  /** 安装包大小 */
  packageSize?: number;
  /** 安装包MD5 */
  packageMd5?: string;
  /** 最低系统版本 */
  minOsVersion?: string;
  /** 灰度策略(1全量 2按比例 3自定义) */
  grayStrategy: number;
  /** 灰度比例 */
  grayPercentage?: number;
  /** 灰度设备 */
  graySns?: string[];
  /** 发布时间 */
  publishTime?: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type UpdateSelfAppReleaseResponses = {
  /**
   * A successful response.
   */
  200: UpdateSelfAppReleaseReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSelfAppReleaseStatusReply = object;

export type UpdateSelfAppReleaseStatusReq = {
  /** ID */
  id: string;
  /** 状态(-1禁用 1启用) */
  status: number;
};

export type UpdateSelfAppReleaseStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSelfAppReleaseStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
