/* eslint-disable */
// @ts-ignore

export type AliyunConfig = {
  /** 访问密钥 */
  accessKey?: string;
  /** 密钥 */
  secretKey?: string;
  /** 桶 */
  bucket?: string;
  /** 端点 */
  endpoint?: string;
  /** 主机 */
  host?: string;
  /** 前缀 */
  prefix?: string;
  /** 盐 */
  salt?: string;
};

export type Any = {
  '@type'?: string;
};

export type CreateFileConfigReply = {
  /** 编号 */
  id?: string;
};

export type CreateFileConfigReq = {
  /** 配置名 */
  name: string;
  /** 存储器 */
  storage: string;
  /** 备注 */
  remark?: string;
  config?: StorageConfig;
};

export type CreateFileConfigResponses = {
  /**
   * A successful response.
   */
  200: CreateFileConfigReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteFileConfigReply = object;

export type DeleteFileConfigReq = {
  /** 编号 */
  id: string;
};

export type DeleteFileConfigResponses = {
  /**
   * A successful response.
   */
  200: DeleteFileConfigReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type FileConfigInfo = {
  /** 编号 */
  id?: string;
  /** 配置名 */
  name?: string;
  /** 存储器 */
  storage?: string;
  /** 备注 */
  remark?: string;
  /** 是否为主配置 */
  master?: boolean;
  config?: StorageConfig;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type FileConfigSelect = {
  /** 编号 */
  id?: string;
  /** 配置名 */
  name?: string;
};

export type FileConfigStorage = {
  /** 标签 */
  label?: string;
  /** 存储器 */
  value?: string;
};

export type GetFileConfigInfoParams = {
  /** 编号 */
  id: string;
};

export type GetFileConfigInfoReply = {
  info?: FileConfigInfo;
};

export type GetFileConfigInfoResponses = {
  /**
   * A successful response.
   */
  200: GetFileConfigInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetFileConfigListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 配置名 */
  name?: string;
  /** 存储器 */
  storage?: string;
};

export type GetFileConfigListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: FileConfigInfo[];
};

export type GetFileConfigListResponses = {
  /**
   * A successful response.
   */
  200: GetFileConfigListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetFileConfigSelectReply = {
  /** 列表数据 */
  list?: FileConfigSelect[];
};

export type GetFileConfigSelectResponses = {
  /**
   * A successful response.
   */
  200: GetFileConfigSelectReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetFileConfigStorageSelectReply = {
  /** 列表数据 */
  list?: FileConfigStorage[];
};

export type GetFileConfigStorageSelectResponses = {
  /**
   * A successful response.
   */
  200: GetFileConfigStorageSelectReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type QiniuConfig = {
  /** 访问密钥 */
  accessKey?: string;
  /** 密钥 */
  secretKey?: string;
  /** 桶 */
  bucket?: string;
  action?: string;
};

export type SetFileConfigMasterReply = object;

export type SetFileConfigMasterReq = {
  /** 编号 */
  id: string;
};

export type SetFileConfigMasterResponses = {
  /**
   * A successful response.
   */
  200: SetFileConfigMasterReply;
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

export type StorageConfig = {
  volcengine?: VolcengineConfig;
  aliyun?: AliyunConfig;
  tencent?: TencentConfig;
  qiniu?: QiniuConfig;
};

export type TencentConfig = {
  /** 访问密钥 */
  accessKey?: string;
  /** 密钥 */
  secretKey?: string;
  /** 端点 */
  endpoint?: string;
  /** 区域 */
  region?: string;
  /** 桶 */
  bucket?: string;
};

export type UpdateFileConfigReply = object;

export type UpdateFileConfigReq = {
  /** 编号 */
  id: string;
  /** 配置名 */
  name: string;
  /** 存储器 */
  storage: string;
  /** 备注 */
  remark?: string;
  config?: StorageConfig;
};

export type UpdateFileConfigResponses = {
  /**
   * A successful response.
   */
  200: UpdateFileConfigReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type VolcengineConfig = {
  /** 访问密钥 */
  accessKey?: string;
  /** 密钥 */
  secretKey?: string;
  /** 端点 */
  endpoint?: string;
  /** 区域 */
  region?: string;
  /** 桶 */
  bucket?: string;
  /** 账号ID */
  accountID?: string;
  /** 角色名称 */
  roleName?: string;
};
