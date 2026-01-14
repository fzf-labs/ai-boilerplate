import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MembershipBenefitTypeInfo } from '#/api/v1/membership-benefit-type';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 新增/修改权益类型的表单 */
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
      fieldName: 'benefitKey',
      label: '权益标识',
      component: 'Input',
      componentProps: {
        placeholder: '请输入权益标识（如：location, screen）',
      },
      rules: 'required',
    },
    {
      fieldName: 'benefitName',
      label: '权益名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入权益名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'benefitIcon',
      label: '权益图标',
      component: 'Input',
      componentProps: {
        placeholder: '请输入权益图标URL（可选）',
      },
    },
    {
      fieldName: 'benefitDesc',
      label: '权益描述',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入权益描述',
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
      fieldName: 'benefitKey',
      label: '权益标识',
      component: 'Input',
      componentProps: {
        placeholder: '请输入权益标识',
        allowClear: true,
      },
    },
    {
      fieldName: 'benefitName',
      label: '权益名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入权益名称',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MembershipBenefitTypeInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'benefitKey',
      title: '权益标识',
      minWidth: 120,
    },
    {
      field: 'benefitName',
      title: '权益名称',
      minWidth: 150,
    },
    {
      field: 'benefitIcon',
      title: '权益图标',
      minWidth: 100,
      showOverflow: 'tooltip',
    },
    {
      field: 'benefitDesc',
      title: '权益描述',
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
      minWidth: 120,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'benefitName',
          nameTitle: '权益类型',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['member:membership-benefit-type:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['member:membership-benefit-type:delete']),
          },
        ],
      },
    },
  ];
}
