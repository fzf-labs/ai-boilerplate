import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SensitiveWordInfo } from '#/api/v1/sensitive-word';

import { useAccess } from '@vben/access';

import { getSensitiveWordLabsSelector } from '#/api/v1/sensitive-word';
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
      fieldName: 'word',
      label: '敏感词',
      component: 'Input',
      componentProps: {
        placeholder: '请输入敏感词',
      },
      rules: 'required',
    },
    {
      fieldName: 'lab',
      label: '标签',
      component: 'ApiSelect',
      componentProps: {
        placeholder: '请选择标签',
        api: getSensitiveWordLabsSelector,
        resultField: 'list',
        labelField: 'value',
        valueField: 'key',
        immediate: true,
      },
      rules: 'required',
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'word',
      label: '敏感词',
      component: 'Input',
      componentProps: {
        placeholder: '请输入敏感词',
        clearable: true,
      },
    },
    {
      fieldName: 'lab',
      label: '标签',
      component: 'ApiSelect',
      componentProps: {
        placeholder: '请选择标签',
        api: getSensitiveWordLabsSelector,
        resultField: 'list',
        labelField: 'value',
        valueField: 'key',
        immediate: true,
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
export function useGridColumns<T = SensitiveWordInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'word',
      title: '敏感词',
      minWidth: 150,
    },
    {
      field: 'lab',
      title: '标签',
      minWidth: 120,
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
      minWidth: 130,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'word',
          nameTitle: '敏感词',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:sensitiveword:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:sensitiveword:delete']),
          },
        ],
      },
    },
  ];
}
