import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { FileDatumInfo } from '#/api/v1/file-data';

import { useAccess } from '@vben/access';

import { getFileConfigSelect } from '#/api/v1/file-config';

const { hasAccessByCodes } = useAccess();

/** 表单的字段 */
export function useFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'file',
      label: '文件上传',
      component: 'Upload',
      componentProps: {
        placeholder: '请选择要上传的文件',
      },
      rules: 'required',
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'configId',
      label: '存储器配置',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getFileConfigSelect({ options: {} }),
        resultField: 'list',
        labelField: 'name',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择存储器配置',
      },
    },
    {
      fieldName: 'name',
      label: '文件名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入文件名',
        clearable: true,
      },
    },
    {
      fieldName: 'path',
      label: '文件路径',
      component: 'Input',
      componentProps: {
        placeholder: '请输入文件路径',
        clearable: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        placeholder: '请选择状态',
        allowClear: true,
        options: [
          { label: '失败', value: -1 },
          { label: '未知', value: 1 },
          { label: '成功', value: 2 },
        ],
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = FileDatumInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '文件名',
      minWidth: 150,
    },
    {
      field: 'path',
      title: '文件路径',
      minWidth: 200,
      showOverflow: true,
    },
    {
      field: 'URL',
      title: 'URL',
      minWidth: 200,
      showOverflow: true,
    },
    {
      field: 'size',
      title: '文件大小',
      minWidth: 80,
      formatter: ({ cellValue }) => {
        if (!cellValue) return '0 B';
        const unitArr = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
        const index = Math.floor(Math.log(cellValue) / Math.log(1024));
        const size = cellValue / 1024 ** index;
        const formattedSize = size.toFixed(2);
        return `${formattedSize} ${unitArr[index]}`;
      },
    },
    {
      field: 'ext',
      title: '文件类型',
      minWidth: 120,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 80,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case -1: {
            return '失败';
          }
          case 1: {
            return '未知';
          }
          case 2: {
            return '成功';
          }
          default: {
            return '未知';
          }
        }
      },
    },
    {
      field: 'createdAt',
      title: '上传时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      width: 160,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '文件',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'copyUrl',
            text: '复制链接',
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:file:delete']),
          },
        ],
      },
    },
  ];
}
