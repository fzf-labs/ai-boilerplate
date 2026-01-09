/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type GetSysOperateLogListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 链路ID */
  traceId?: string;
  /** 管理员ID */
  adminId?: string;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSysOperateLogListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysOperateLogInfo[];
};

export type GetSysOperateLogListResponses = {
  /**
   * A successful response.
   */
  200: GetSysOperateLogListReply;
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

export type SysOperateLogInfo = {
  /** 编号 */
  id?: string;
  /** 链路ID */
  traceId?: string;
  /** 管理员ID */
  adminId?: string;
  /** ip */
  IP?: string;
  /** 请求路径 */
  URI?: string;
  /** 浏览器标识 */
  useragent?: string;
  /** header */
  header?: string;
  /** 请求数据 */
  req?: string;
  /** 响应数据 */
  resp?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 管理员昵称 */
  nickname?: string;
};
