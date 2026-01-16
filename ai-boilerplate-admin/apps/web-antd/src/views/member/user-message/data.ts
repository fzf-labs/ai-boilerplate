import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { UserMessageInfo } from '#/api/v1/user-message';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';

const { hasAccessByCodes } = useAccess();

export const MessageCategoryOptions = [
  { label: '交易信息', value: 'transaction' },
  { label: '系统消息', value: 'system' },
  { label: '客服消息', value: 'service' },
];

export const AudienceTypeOptions = [
  { label: '全部用户', value: 'all' },
  { label: '条件筛选', value: 'segment' },
  { label: '指定用户', value: 'users' },
];

export const ReadStatusOptions = [
  { label: '已读', value: 1 },
  { label: '未读', value: -1 },
];

export const MembershipTypeOptions = [
  { label: '普通会员', value: 'normal' },
  { label: 'VIP会员', value: 'vip' },
  { label: 'SVIP会员', value: 'svip' },
];

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'category',
      label: '分类',
      component: 'Select',
      componentProps: {
        options: MessageCategoryOptions,
        placeholder: '请选择分类',
        allowClear: true,
      },
    },
    {
      fieldName: 'title',
      label: '标题',
      component: 'Input',
      componentProps: {
        placeholder: '请输入标题',
        allowClear: true,
      },
    },
    {
      fieldName: 'userId',
      label: '用户ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户ID',
        allowClear: true,
      },
    },
    {
      fieldName: 'messageId',
      label: '批次ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入批次ID',
        allowClear: true,
      },
    },
    {
      fieldName: 'audienceType',
      label: '投放范围',
      component: 'Select',
      componentProps: {
        options: AudienceTypeOptions,
        placeholder: '请选择投放范围',
        allowClear: true,
      },
    },
    {
      fieldName: 'readStatus',
      label: '阅读状态',
      component: 'Select',
      componentProps: {
        options: ReadStatusOptions,
        placeholder: '请选择阅读状态',
        allowClear: true,
      },
    },
  ];
}

/** 发送消息表单 */
export function useSendFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'category',
      label: '消息分类',
      component: 'Select',
      componentProps: {
        options: MessageCategoryOptions,
        placeholder: '请选择消息分类',
      },
      rules: 'required',
    },
    {
      fieldName: 'title',
      label: '标题',
      component: 'Input',
      componentProps: {
        placeholder: '请输入消息标题',
      },
      rules: 'required',
    },
    {
      fieldName: 'summary',
      label: '摘要',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入消息摘要(可选)',
        rows: 3,
      },
    },
    {
      fieldName: 'coverURL',
      label: '封面图',
      component: 'Input',
      componentProps: {
        placeholder: '请输入封面图地址(可选)',
      },
    },
    {
      fieldName: 'content',
      label: '内容',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入消息内容',
        rows: 6,
      },
      rules: 'required',
    },
    {
      fieldName: 'linkURL',
      label: '跳转链接',
      component: 'Input',
      componentProps: {
        placeholder: '请输入跳转链接(可选)',
      },
    },
    {
      fieldName: 'audienceType',
      label: '投放范围',
      component: 'RadioGroup',
      componentProps: {
        options: AudienceTypeOptions,
        optionType: 'button',
        buttonStyle: 'solid',
      },
      defaultValue: 'all',
      rules: z.string().min(1, '请选择投放范围'),
    },
    {
      fieldName: 'userIdsText',
      label: '用户ID列表',
      component: 'Textarea',
      componentProps: {
        placeholder: '多个用户ID可用逗号或换行分隔',
        rows: 4,
      },
      dependencies: {
        triggerFields: ['audienceType'],
        show: (values) => values?.audienceType === 'users',
        required: (values) => values?.audienceType === 'users',
      },
    },
    {
      fieldName: 'membershipType',
      label: '会员类型',
      component: 'Select',
      componentProps: {
        options: MembershipTypeOptions,
        placeholder: '请选择会员类型(可选)',
        allowClear: true,
      },
      dependencies: {
        triggerFields: ['audienceType'],
        show: (values) => values?.audienceType === 'segment',
      },
    },
    {
      fieldName: 'activeWithinDays',
      label: '活跃天数',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入活跃天数(可选)',
        min: 1,
        style: { width: '100%' },
      },
      dependencies: {
        triggerFields: ['audienceType'],
        show: (values) => values?.audienceType === 'segment',
      },
      rules: z.number().int().min(1).optional(),
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = UserMessageInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'category',
      title: '分类',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        const match = MessageCategoryOptions.find(
          (item) => item.value === cellValue,
        );
        return match?.label || cellValue || '-';
      },
    },
    {
      field: 'title',
      title: '标题',
      minWidth: 180,
      showOverflow: true,
    },
    {
      field: 'summary',
      title: '摘要',
      minWidth: 200,
      showOverflow: true,
    },
    {
      field: 'userId',
      title: '用户ID',
      minWidth: 160,
      showOverflow: true,
    },
    {
      field: 'audienceType',
      title: '投放范围',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        const match = AudienceTypeOptions.find(
          (item) => item.value === cellValue,
        );
        return match?.label || cellValue || '-';
      },
    },
    {
      field: 'sentAt',
      title: '发送时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'readAt',
      title: '阅读时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'messageId',
      title: '批次ID',
      minWidth: 200,
      showOverflow: true,
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 120,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'id',
          nameTitle: '用户消息',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['member:user-message:query']),
          },
        ],
      },
    },
  ];
}
