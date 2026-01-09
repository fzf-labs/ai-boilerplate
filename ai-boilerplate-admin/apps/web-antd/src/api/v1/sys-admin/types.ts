/* eslint-disable */
// @ts-ignore

export type Any = {
  '@type'?: string;
};

export type CreateSysAdminReply = {
  /** 编号 */
  id?: string;
};

export type CreateSysAdminReq = {
  /** 用户名 */
  username: string;
  /** 密码 */
  password?: string;
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 0=保密 1=女 2=男 */
  sex?: number;
  /** 邮件 */
  email?: string;
  /** 手机号 */
  mobile?: string;
  /** 角色 */
  roleId?: string;
  /** 部门 */
  deptId?: string;
  /** 岗位 */
  postId?: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type CreateSysAdminResponses = {
  /**
   * A successful response.
   */
  200: CreateSysAdminReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type DeleteSysAdminReply = object;

export type DeleteSysAdminReq = {
  /** 编号 */
  id: string;
};

export type DeleteSysAdminResponses = {
  /**
   * A successful response.
   */
  200: DeleteSysAdminReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysAdminInfoParams = {
  /** 编号 */
  id: string;
};

export type GetSysAdminInfoReply = {
  info?: SysAdminInfo;
};

export type GetSysAdminInfoResponses = {
  /**
   * A successful response.
   */
  200: GetSysAdminInfoReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysAdminListParams = {
  /** 页码 */
  page?: number;
  /** 页数 */
  pageSize?: number;
  /** 用户名 */
  username?: string;
  /** 昵称 */
  nickname?: string;
  /** 手机号 */
  mobile?: string;
  /** 角色 */
  roleId?: string;
  /** 部门 */
  deptId?: string;
  /** 岗位 */
  postId?: string;
  /** 状态 */
  status?: number;
  /** 创建时间 */
  createdAt?: string[];
};

export type GetSysAdminListReply = {
  /** 总数 */
  total?: number;
  /** 列表数据 */
  list?: SysAdminInfo[];
};

export type GetSysAdminListResponses = {
  /**
   * A successful response.
   */
  200: GetSysAdminListReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type GetSysAdminSelectorItem = {
  /** 编号 */
  id?: string;
  /** 名称 */
  nickname?: string;
};

export type GetSysAdminSelectorReply = {
  /** 列表数据 */
  list?: GetSysAdminSelectorItem[];
};

export type GetSysAdminSelectorResponses = {
  /**
   * A successful response.
   */
  200: GetSysAdminSelectorReply;
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

export type SysAdminInfo = {
  /** 编号 */
  id?: string;
  /** 用户名 */
  username?: string;
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 0=保密 1=女 2=男 */
  sex?: number;
  /** 邮件 */
  email?: string;
  /** 手机号 */
  mobile?: string;
  /** 角色 */
  roleId?: string;
  /** 部门 */
  deptId?: string;
  /** 岗位 */
  postId?: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 角色名称 */
  roleName?: string;
  /** 部门名称 */
  deptName?: string;
  /** 岗位名称 */
  postName?: string;
};

export type UpdateSysAdminPasswordReply = object;

export type UpdateSysAdminPasswordReq = {
  /** 编号 */
  id: string;
  /** 密码 */
  password: string;
};

export type UpdateSysAdminPasswordResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysAdminPasswordReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysAdminReply = object;

export type UpdateSysAdminReq = {
  /** 编号 */
  id: string;
  /** 用户名 */
  username: string;
  /** 昵称 */
  nickname?: string;
  /** 头像 */
  avatar?: string;
  /** 0=保密 1=女 2=男 */
  sex?: number;
  /** 邮件 */
  email?: string;
  /** 手机号 */
  mobile?: string;
  /** 角色 */
  roleId?: string;
  /** 部门 */
  deptId?: string;
  /** 岗位 */
  postId?: string;
  /** 状态 -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysAdminResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysAdminReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};

export type UpdateSysAdminStatusReply = object;

export type UpdateSysAdminStatusReq = {
  /** 编号 */
  id: string;
  /** -1=禁用 1=开启 */
  status?: number;
};

export type UpdateSysAdminStatusResponses = {
  /**
   * A successful response.
   */
  200: UpdateSysAdminStatusReply;
  /**
   * An unexpected error response.
   */
  default: Status;
};
