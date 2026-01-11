/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CheckUpdateReply = {
  /** 是否有更新 */
  hasUpdate?: boolean;
  updateInfo?: UpdateInfo;
};

export type CheckUpdateReq = {
  /** 包名 */
  packageName: string;
  /** 渠道 */
  channel: string;
  /** 当前build值 */
  buildNum: number;
  /** 设备序列号(用于灰度判断) */
  deviceSn?: string;
  /** 系统版本 */
  osVersion?: string;
};

export type CheckUpdateResponses = {
  /**
   * A successful response.
   */
  200: CheckUpdateReply;
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

export type UpdateInfo = {
  /** 新版本号 */
  version?: string;
  /** 新build值 */
  buildNum?: number;
  /** 更新类型(1强制 2提示 3静默) */
  updateType?: number;
  /** 更新标题 */
  title?: string;
  /** 更新日志 */
  changelog?: string;
  /** 安装包地址 */
  packageURL?: string;
  /** 安装包大小(字节) */
  packageSize?: number;
  /** 安装包MD5 */
  packageMd5?: string;
  /** 最低系统版本要求 */
  minOsVersion?: string;
};
