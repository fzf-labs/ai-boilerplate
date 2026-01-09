import type { AxiosProgressEvent } from '#/api/infra/file/data';

export enum UploadResultStatus {
  DONE = 'done',
  ERROR = 'error',
  SUCCESS = 'success',
  UPLOADING = 'uploading',
}

export type UploadListType = 'picture' | 'picture-card' | 'text';

export interface FileUploadProps {
  // 根据后缀，或者其他
  accept?: string[];
  api?: (file: File, onUploadProgress?: AxiosProgressEvent) => Promise<any>;
  disabled?: boolean;
  helpText?: string;
  listType?: UploadListType;
  // 最大数量的文件，Infinity不限制
  maxNumber?: number;
  // 文件最大多少MB
  maxSize?: number;
  // 是否支持多选
  multiple?: boolean;
  // support xxx.xxx.xx
  resultField?: string;
  // 场景值，用于区分不同业务场景的文件上传路径
  scene?: string;
  // 是否显示下面的描述
  showDescription?: boolean;
  value?: string | string[];
}
