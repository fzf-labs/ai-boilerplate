/* eslint-disable */
// @ts-ignore

export type AliyunPolicy = {
  policy?: string;
  securityToken?: string;
  signatureVersion?: string;
  credential?: string;
  date?: string;
  signature?: string;
  host?: string;
  dir?: string;
  callback?: string;
};

export type Any = {
  '@type'?: string;
};

export type DeleteFileDatumReply = object;

export type DeleteFileDatumReq = {
  /** 文件编号 */
  id: string;
};

export type DeleteFileDatumResponses = {
  /**
   * A successful response.
   */
  200: DeleteFileDatumReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type FileDatumInfo = {
  /** 文件编号 */
  id?: string;
  /** 配置编号 */
  configId?: string;
  /** 文件名 */
  name?: string;
  /** 文件路径 */
  path?: string;
  /** 文件 URL */
  URL?: string;
  /** 文件类型 */
  ext?: string;
  /** 文件大小 */
  size?: number;
  /** 状态（-1失败,1未知,2成功） */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
};

export type GetFileDatumInfoParams = {
  /** 文件编号 */
  id: string;
};

export type GetFileDatumInfoReply = {
  info?: FileDatumInfo;
};

export type GetFileDatumInfoResponses = {
  /**
   * A successful response.
   */
  200: GetFileDatumInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetFileDatumListParams = {
  /** 页码 */
  page: number;
  /** 页数 */
  pageSize: number;
  /** 配置编号 */
  configId?: string;
  /** 文件名 */
  name?: string;
  /** 文件路径 */
  path?: string;
  /** 状态（-1失败,1未知,2成功） */
  status?: number;
};

export type GetFileDatumListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: FileDatumInfo[];
};

export type GetFileDatumListResponses = {
  /**
   * A successful response.
   */
  200: GetFileDatumListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type QiniuPolicy = {
  /** 凭证 */
  token?: string;
};

export type Status = {
  code?: number;
  message?: string;
  details?: Any[];
};

export type TencentPolicy = {
  /** 临时密钥ID */
  tmpSecretId?: string;
  /** 临时密钥Key */
  tmpSecretKey?: string;
  /** 临时 token */
  sessionToken?: string;
};

export type UploadFileOSSDefaultPolicyParams = {
  /** 文件名 */
  name: string;
  /** 文件路径 */
  path: string;
  /** 文件大小 */
  size?: number;
};

export type UploadFileOSSDefaultPolicyReply = {
  /** 文件ID */
  fileId?: string;
  /** 存储引擎 火山云:volcengine 腾讯云:tencent 阿里云:aliyun 七牛云:qiniu */
  storage?: string;
  volcengine?: VolcenginePolicy;
  tencent?: TencentPolicy;
  aliyun?: AliyunPolicy;
  qiniu?: QiniuPolicy;
};

export type UploadFileOSSDefaultPolicyResponses = {
  /**
   * A successful response.
   */
  200: UploadFileOSSDefaultPolicyReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type VolcenginePolicy = {
  /** 临时 accessKeyId */
  accessKeyId?: string;
  /** 临时 secretAccessKey */
  secretAccessKey?: string;
  /** 临时 token */
  sessionToken?: string;
  /** 端点 */
  endpoint?: string;
  /** 区域 */
  region?: string;
  /** 桶 */
  bucket?: string;
};
