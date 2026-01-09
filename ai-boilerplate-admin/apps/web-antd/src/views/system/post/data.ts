import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysPostInfo } from '#/api/v1/sys-post';

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
      component: 'Input',
      fieldName: 'name',
      label: '岗位名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: '岗位编码',
      rules: 'required',
    },
    {
      fieldName: 'sort',
      label: '显示顺序',
      component: 'InputNumber',
      componentProps: {
        min: 0,
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '岗位状态',
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
    {
      fieldName: 'remark',
      label: '岗位备注',
      component: 'Textarea',
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'name',
      label: '岗位名称',
      component: 'Input',
    },
    {
      fieldName: 'code',
      label: '岗位编码',
      component: 'Input',
    },
    {
      fieldName: 'status',
      label: '岗位状态',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SysPostInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '岗位名称',
      minWidth: 200,
    },
    {
      field: 'code',
      title: '岗位编码',
      minWidth: 200,
    },
    {
      field: 'remark',
      title: '岗位备注',
      minWidth: 200,
    },
    {
      field: 'sort',
      title: '显示顺序',
      minWidth: 100,
    },
    {
      field: 'status',
      title: '岗位状态',
      minWidth: 100,
      formatter: ({ cellValue }) => {
        return cellValue === 1 ? '启用' : '禁用';
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
      minWidth: 130,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '岗位',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['system:post:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:post:delete']),
          },
        ],
      },
    },
  ];
}
