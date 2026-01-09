import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SelfAppInfo } from '#/api/v1/self-app';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { CommonStatusEnum } from '#/utils/constants';

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
      fieldName: 'packageName',
      label: '包名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入应用包名，如：com.example.app',
      },
      rules: 'required',
    },
    {
      fieldName: 'name',
      label: '应用名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入应用名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'description',
      label: '应用描述',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入应用描述（可选）',
        rows: 4,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'packageName',
      label: '包名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入应用包名',
        allowClear: true,
      },
    },
    {
      fieldName: 'name',
      label: '应用名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入应用名称',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        options: [
          { label: '全部', value: '' },
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
        placeholder: '请选择状态',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SelfAppInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'packageName',
      title: '包名',
      minWidth: 250,
    },
    {
      field: 'name',
      title: '应用名称',
      minWidth: 180,
    },
    {
      field: 'description',
      title: '应用描述',
      minWidth: 200,
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        attrs: { beforeChange: onStatusChange },
        name: 'CellSwitch',
        props: {
          checkedValue: 1,
          unCheckedValue: -1,
        },
      },
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
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '应用',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['self_app:query']),
          },
          {
            code: 'edit',
            show: hasAccessByCodes(['self_app:update']),
          },
          {
            code: 'versions',
            text: '版本管理',
            show: hasAccessByCodes(['self_app_release:query']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['self_app:delete']),
          },
        ],
      },
    },
  ];
}
