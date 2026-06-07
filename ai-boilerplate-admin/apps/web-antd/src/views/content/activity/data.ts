import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { ActivityInfo } from '#/api/v1/activity';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';

const { hasAccessByCodes } = useAccess();

export const ActivityStatusOptions = [
  { label: '禁用', value: -1 },
  { label: '启用', value: 1 },
];

export const ActivityLinkTypeOptions = [
  { label: '站内', value: 'internal' },
  { label: '外链', value: 'external' },
];

/** 新增/修改的表单 */
export function useFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'id',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'title',
      label: '标题',
      component: 'Input',
      rules: 'required',
      componentProps: {
        placeholder: '请输入标题',
        allowClear: true,
      },
    },
    {
      fieldName: 'imageURL',
      label: '活动图',
      component: 'ImageUpload',
      rules: z.string().min(1, '请上传活动图').max(500),
      componentProps: {
        maxNumber: 1,
      },
    },
    {
      fieldName: 'linkType',
      label: '跳转类型',
      component: 'RadioGroup',
      componentProps: {
        options: ActivityLinkTypeOptions,
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.string().min(1),
    },
    {
      fieldName: 'linkURL',
      label: '跳转链接',
      component: 'Input',
      rules: z.string().min(1, '请输入跳转链接').max(500),
      componentProps: {
        placeholder: '外链需以 https:// 开头，站内链接需以 /pages/ 或 app:// 开头',
        allowClear: true,
      },
    },
    {
      fieldName: 'sort',
      label: '排序',
      component: 'InputNumber',
      rules: z.number().default(0),
      componentProps: {
        min: 0,
        max: 999_999,
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'RadioGroup',
      componentProps: {
        options: ActivityStatusOptions,
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(1),
    },
    {
      fieldName: 'startTime',
      label: '开始时间',
      component: 'DatePicker',
      componentProps: {
        placeholder: '可选',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DDTHH:mm:ssZ',
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'endTime',
      label: '结束时间',
      component: 'DatePicker',
      componentProps: {
        placeholder: '可选',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DDTHH:mm:ssZ',
        style: { width: '100%' },
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'keyword',
      label: '关键词',
      component: 'Input',
      componentProps: {
        placeholder: '请输入标题关键词',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        options: ActivityStatusOptions,
        placeholder: '请选择状态',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = ActivityInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'title',
      title: '标题',
      minWidth: 220,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        const found = ActivityStatusOptions.find(
          (item) => item.value === cellValue,
        );
        return found?.label || '';
      },
    },
    {
      field: 'sort',
      title: '排序',
      minWidth: 80,
    },
    {
      field: 'startTime',
      title: '开始时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'endTime',
      title: '结束时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'updatedAt',
      title: '更新时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 240,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'title',
          nameTitle: '活动',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['content:activity:update']),
          },
          {
            code: 'enable',
            text: '启用',
            show: hasAccessByCodes(['content:activity:enable']),
          },
          {
            code: 'disable',
            text: '禁用',
            show: hasAccessByCodes(['content:activity:disable']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['content:activity:delete']),
          },
        ],
      },
    },
  ];
}
