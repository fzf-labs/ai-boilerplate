import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'adminId',
      label: '用户编号',
      component: 'ApiSelect',
      componentProps: {
        api: getSysAdminSelector,
        labelField: 'nickname',
        valueField: 'id',
      },
    },
    {
      fieldName: 'type',
      label: '写作类型',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          {
            value: 1,
            label: '撰写',
          },
          {
            value: 2,
            label: '改写',
          },
        ],
      },
    },
    {
      fieldName: 'platform',
      label: '平台',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入平台名称',
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
export function useGridColumns(): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '编号',
      minWidth: 200,
      fixed: 'left',
    },
    {
      minWidth: 150,
      title: '用户',
      slots: { default: 'adminId' },
    },
    {
      field: 'type',
      title: '写作类型',
      minWidth: 100,
      cellRender: {
        name: 'CellDict',
      },
    },
    {
      field: 'platform',
      title: '平台',
      minWidth: 120,
    },
    {
      field: 'model',
      title: '模型',
      minWidth: 150,
    },
    {
      field: 'prompt',
      title: '生成内容提示',
      minWidth: 200,
    },
    {
      field: 'generatedContent',
      title: '生成的内容',
      minWidth: 250,
    },
    {
      field: 'originalContent',
      title: '原文',
      minWidth: 200,
    },
    {
      field: 'length',
      title: '长度',
      minWidth: 100,
    },
    {
      field: 'format',
      title: '格式',
      minWidth: 100,
    },
    {
      field: 'tone',
      title: '语气',
      minWidth: 100,
    },
    {
      field: 'language',
      title: '语言',
      minWidth: 100,
    },
    {
      field: 'errorMessage',
      title: '错误信息',
      minWidth: 180,
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      title: '操作',
      width: 130,
      fixed: 'right',
      slots: { default: 'actions' },
    },
  ];
}
