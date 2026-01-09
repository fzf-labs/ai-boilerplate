import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { FileConfigInfo } from '#/api/v1/file-config';

import { useAccess } from '@vben/access';

import { getFileConfigStorageSelect } from '#/api/v1/file-config';
import { getRangePickerDefaultProps } from '#/utils';

const { hasAccessByCodes } = useAccess();

/** 新增/修改的表单 */
export function useFormSchema(): VbenFormSchema[] {
  return [
    {
      component: 'Input',
      fieldName: 'id',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'name',
      label: '配置名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入配置名',
      },
      rules: 'required',
    },
    {
      fieldName: 'storage',
      label: '存储器',
      component: 'ApiSelect',
      componentProps: {
        api: getFileConfigStorageSelect,
        resultField: 'list',
        labelField: 'label',
        valueField: 'value',
      },
      rules: 'required',
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注',
      },
    },
    // 阿里云 OSS 配置
    {
      fieldName: 'config.aliyun.accessKey',
      label: 'Access Key',
      component: 'Input',
      componentProps: {
        placeholder: '请输入阿里云 Access Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.secretKey',
      label: 'Secret Key',
      component: 'Input',
      componentProps: {
        type: 'password',
        placeholder: '请输入阿里云 Secret Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.bucket',
      label: '存储桶',
      component: 'Input',
      componentProps: {
        placeholder: '请输入阿里云 OSS 存储桶名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.endpoint',
      label: '端点',
      component: 'Input',
      componentProps: {
        placeholder: '请输入阿里云 OSS 端点地址',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.host',
      label: '主机地址',
      component: 'Input',
      componentProps: {
        placeholder: '请输入阿里云 OSS 主机地址',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.prefix',
      label: '前缀',
      component: 'Input',
      componentProps: {
        placeholder: '请输入文件前缀（可选）',
      },
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    {
      fieldName: 'config.aliyun.salt',
      label: '盐值',
      component: 'Input',
      componentProps: {
        placeholder: '请输入加密盐值（可选）',
      },
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'aliyun',
      },
    },
    // 腾讯云 COS 配置
    {
      fieldName: 'config.tencent.accessKey',
      label: 'Access Key',
      component: 'Input',
      componentProps: {
        placeholder: '请输入腾讯云 Access Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'tencent',
      },
    },
    {
      fieldName: 'config.tencent.secretKey',
      label: 'Secret Key',
      component: 'Input',
      componentProps: {
        type: 'password',
        placeholder: '请输入腾讯云 Secret Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'tencent',
      },
    },
    {
      fieldName: 'config.tencent.bucket',
      label: '存储桶',
      component: 'Input',
      componentProps: {
        placeholder: '请输入腾讯云 COS 存储桶名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'tencent',
      },
    },
    {
      fieldName: 'config.tencent.endpoint',
      label: '端点',
      component: 'Input',
      componentProps: {
        placeholder: '请输入腾讯云 COS 端点地址',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'tencent',
      },
    },
    {
      fieldName: 'config.tencent.region',
      label: '区域',
      component: 'Input',
      componentProps: {
        placeholder: '请输入腾讯云 COS 区域',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'tencent',
      },
    },
    // 七牛云配置
    {
      fieldName: 'config.qiniu.accessKey',
      label: 'Access Key',
      component: 'Input',
      componentProps: {
        placeholder: '请输入七牛云 Access Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'qiniu',
      },
    },
    {
      fieldName: 'config.qiniu.secretKey',
      label: 'Secret Key',
      component: 'Input',
      componentProps: {
        type: 'password',
        placeholder: '请输入七牛云 Secret Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'qiniu',
      },
    },
    {
      fieldName: 'config.qiniu.bucket',
      label: '存储空间',
      component: 'Input',
      componentProps: {
        placeholder: '请输入七牛云存储空间名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'qiniu',
      },
    },
    {
      fieldName: 'config.qiniu.action',
      label: '上传策略',
      component: 'Input',
      componentProps: {
        placeholder: '请输入上传策略（可选）',
      },
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'qiniu',
      },
    },
    // 火山云配置
    {
      fieldName: 'config.volcengine.accessKey',
      label: 'Access Key',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云 Access Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.secretKey',
      label: 'Secret Key',
      component: 'Input',
      componentProps: {
        type: 'password',
        placeholder: '请输入火山云 Secret Key',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.bucket',
      label: '存储桶',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云存储桶名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.endpoint',
      label: '端点',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云端点地址',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.region',
      label: '区域',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云区域',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.customDomain',
      label: '自定义域名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云自定义域名',
      },
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.accountID',
      label: '账号ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云账号ID',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
    {
      fieldName: 'config.volcengine.roleName',
      label: '角色名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入火山云角色名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['storage'],
        show: (formValues) => formValues.storage === 'volcengine',
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'name',
      label: '配置名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入配置名',
        clearable: true,
      },
    },
    {
      fieldName: 'storage',
      label: '存储器',
      component: 'ApiSelect',
      componentProps: {
        api: getFileConfigStorageSelect,
        resultField: 'list',
        labelField: 'label',
        valueField: 'value',
        placeholder: '请选择存储器',
        clearable: true,
      },
    },
    {
      fieldName: 'createdAt',
      label: '创建时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = FileConfigInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '配置名',
      minWidth: 120,
    },
    {
      field: 'storage',
      title: '存储器',
      width: 100,
    },
    {
      field: 'master',
      title: '主配置',
      width: 100,
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 150,
    },
    {
      field: 'createdAt',
      title: '创建时间',
      width: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      width: 280,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '文件配置',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:file-config:update']),
          },
          {
            code: 'master',
            text: '主配置',
            disabled: (row: any) => row.master,
            show: (_row: any) => hasAccessByCodes(['infra:file-config:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:file-config:delete']),
          },
        ],
      },
    },
  ];
}
