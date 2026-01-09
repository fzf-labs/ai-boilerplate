import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MembershipInfo } from '#/api/v1/membership';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 新增/修改会员类型的表单 */
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
      label: '会员类型名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入会员类型名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'type',
      label: '会员类型编码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入会员类型编码',
      },
      rules: 'required',
    },
    {
      fieldName: 'description',
      label: '会员类型描述',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入会员类型描述',
        rows: 3,
      },
    },
    {
      fieldName: 'sort',
      label: '排序',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入排序（数字越小越靠前）',
        min: 0,
        class: 'w-full',
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
      fieldName: 'name',
      label: '会员类型名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入会员类型名称',
        allowClear: true,
      },
    },
    {
      fieldName: 'type',
      label: '会员类型编码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入会员类型编码',
        allowClear: true,
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
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
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
export function useGridColumns<T = MembershipInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '会员类型名称',
      minWidth: 150,
    },
    {
      field: 'type',
      title: '会员类型编码',
      minWidth: 120,
    },
    {
      field: 'description',
      title: '会员类型描述',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'sort',
      title: '排序',
      minWidth: 80,
      align: 'center',
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
          checkedValue: CommonStatusEnum.ENABLE,
          unCheckedValue: CommonStatusEnum.DISABLE,
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
      minWidth: 160,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '会员类型',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['member:membership:update']),
          },
          {
            code: 'benefit',
            text: '权益管理',
            show: hasAccessByCodes(['member:membership-benefit:query']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['member:membership:delete']),
          },
        ],
      },
    },
  ];
}
