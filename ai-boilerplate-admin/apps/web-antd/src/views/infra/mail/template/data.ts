import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MailTemplateInfo } from '#/api/v1/mail-template';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getMailAccountSelector } from '#/api/v1/mail-account';
import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

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
      fieldName: 'name',
      label: '模板名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入模板名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'code',
      label: '模板编码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入模板编码',
      },
      rules: 'required',
    },
    {
      fieldName: 'accountId',
      label: '邮箱账号',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getMailAccountSelector({ options: {} }),
        class: 'w-full',
        resultField: 'list',
        labelField: 'mail',
        valueField: 'id',
        placeholder: '请选择邮箱账号',
      },
      rules: 'required',
    },
    {
      fieldName: 'nickname',
      label: '发送人名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入发送人名称',
      },
    },
    {
      fieldName: 'title',
      label: '模板标题',
      component: 'Input',
      componentProps: {
        placeholder: '请输入模板标题',
      },
      rules: 'required',
    },
    {
      fieldName: 'content',
      label: '模板内容',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入模板内容',
        height: 300,
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '启用状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            label: '启用',
            value: CommonStatusEnum.ENABLE,
          },
          {
            label: '禁用',
            value: CommonStatusEnum.DISABLE,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注',
      },
    },
  ];
}

/** 发送邮件表单 */
export function useSendMailFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'templateParams',
      label: '模板参数',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'content',
      label: '模板内容',
      component: 'Textarea',
      componentProps: {
        disabled: true,
      },
    },
    {
      fieldName: 'mail',
      label: '收件邮箱',
      component: 'Input',
      componentProps: {
        placeholder: '请输入收件邮箱',
      },
      rules: z.string().email('请输入正确的邮箱'),
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'code',
      label: '模板编码',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入模板编码',
      },
    },
    {
      fieldName: 'name',
      label: '模板名称',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入模板名称',
      },
    },
    {
      fieldName: 'accountId',
      label: '邮箱账号',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getMailAccountSelector({ options: {} }),
        resultField: 'list',
        labelField: 'mail',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择邮箱账号',
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
    {
      fieldName: 'status',
      label: '启用状态',
      component: 'Select',
      componentProps: {
        options: [
          {
            label: '启用',
            value: CommonStatusEnum.ENABLE,
          },
          {
            label: '禁用',
            value: CommonStatusEnum.DISABLE,
          },
        ],
        allowClear: true,
        placeholder: '请选择启用状态',
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MailTemplateInfo>(
  onActionClick: OnActionClickFn<T>,
  getAccountMail?: (accountId: string) => string | undefined,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'code',
      title: '模板编码',
      minWidth: 120,
    },
    {
      field: 'name',
      title: '模板名称',
      minWidth: 120,
    },
    {
      field: 'title',
      title: '模板标题',
      minWidth: 120,
    },
    {
      field: 'accountId',
      title: '邮箱账号',
      minWidth: 120,
      formatter: (row) => getAccountMail?.(row.cellValue) || '-',
    },
    {
      field: 'nickname',
      title: '发送人名称',
      minWidth: 120,
    },
    {
      field: 'status',
      title: '启用状态',
      minWidth: 100,
      formatter: (row) => {
        switch (row.cellValue) {
          case CommonStatusEnum.DISABLE: {
            return '禁用';
          }
          case CommonStatusEnum.ENABLE: {
            return '启用';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'updateAt',
      title: '更新时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 150,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '邮件模板',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:mail-template:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:mail-template:delete']),
          },
          {
            code: 'send',
            text: '测试',
            show: hasAccessByCodes(['infra:mail-template:send-mail']),
          },
        ],
      },
    },
  ];
}
