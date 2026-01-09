/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type DeleteDeviceReply = object;

export type DeleteDeviceReq = {
  /** ID */
  sn: string;
};

export type DeleteDeviceResponses = {
  /**
   * A successful response.
   */
  200: DeleteDeviceReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeviceInfo = {
  /** ID */
  id?: string;
  /** 设备ID */
  sn?: string;
  /** 设备名称 */
  name?: string;
  /** 描述 */
  desc?: string;
  /** 设备品牌 */
  brand?: string;
  /** 设备型号 */
  model?: string;
  /** 入网型号 */
  network?: string;
  /** IMEI */
  imei?: string;
  /** cpu型号 */
  CPU?: string;
  /** mac地址 */
  mac?: string;
  /** app版本 */
  appVersion?: string;
  /** 安卓版本 */
  androidVersion?: string;
  /** RAM大小 */
  RAMSize?: number;
  /** DDR大小 */
  ddrSize?: number;
  /** 设备证书 */
  certificate?: string;
  /** 设备密钥 */
  secureKey?: string;
  /** 激活时间 */
  registryTime?: string;
  push?: DevicePush;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 在线状态 */
  online?: boolean;
};

export type DevicePush = {
  /** 推送通道ID */
  channelID?: string;
};

export type GetDeviceInfoParams = {
  /** 设备SN */
  sn: string;
};

export type GetDeviceInfoReply = {
  info?: DeviceInfo;
};

export type GetDeviceInfoResponses = {
  /**
   * A successful response.
   */
  200: GetDeviceInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetDeviceListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: DeviceInfo[];
};

export type GetDeviceListReq = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 设备SN */
  sn?: string;
  /** 状态 */
  status?: number;
  /** 注册时间 */
  registryTime?: string[];
  /** 在线状态 online:在线 offline:离线 all:全部 */
  onlineSearch?: string;
};

export type GetDeviceListResponses = {
  /**
   * A successful response.
   */
  200: GetDeviceListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetOnlineDeviceCountReply = {
  /** 在线设备数量 */
  count?: number;
};

export type GetOnlineDeviceCountResponses = {
  /**
   * A successful response.
   */
  200: GetOnlineDeviceCountReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type RegisterDeviceReply = object;

export type RegisterDeviceReq = {
  /** 设备的唯一标识序列号 */
  sn: string;
};

export type RegisterDeviceResponses = {
  /**
   * A successful response.
   */
  200: RegisterDeviceReply;
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

export type UpdateDeviceStatusReply = object;

export type UpdateDeviceStatusReq = {
  /** 设备SN */
  sn: string;
  /** 状态 */
  status?: number;
};

export type UpdateDeviceStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateDeviceStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
