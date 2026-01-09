import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MailAccountInfo } from '#/api/v1/mail-account';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
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
      fieldName: 'mail',
      label: '邮箱',
      component: 'Input',
      componentProps: {
        placeholder: '请输入邮箱',
      },
      rules: 'required',
    },
    {
      fieldName: 'username',
      label: '用户名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户名',
      },
      rules: 'required',
    },
    {
      fieldName: 'password',
      label: '密码',
      component: 'InputPassword',
      componentProps: {
        placeholder: '请输入密码',
      },
      rules: 'required',
    },
    {
      fieldName: 'host',
      label: 'SMTP 服务器域名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入 SMTP 服务器域名',
      },
      rules: 'required',
    },
    {
      fieldName: 'port',
      label: 'SMTP 服务器端口',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入 SMTP 服务器端口',
        min: 0,
        max: 65_535,
      },
      rules: 'required',
    },
    {
      fieldName: 'sslEnable',
      label: '是否开启 SSL',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            label: '是',
            value: true,
          },
          {
            label: '否',
            value: false,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.boolean().default(true),
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注',
      },
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
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'mail',
      label: '邮箱',
      component: 'Input',
      componentProps: {
        placeholder: '请输入邮箱',
        clearable: true,
      },
    },
    {
      fieldName: 'username',
      label: '用户名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户名',
        clearable: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MailAccountInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'mail',
      title: '邮箱',
      minWidth: 160,
    },
    {
      field: 'username',
      title: '用户名',
      minWidth: 160,
    },
    {
      field: 'host',
      title: 'SMTP 服务器域名',
      minWidth: 150,
    },
    {
      field: 'port',
      title: 'SMTP 服务器端口',
      minWidth: 130,
    },
    {
      field: 'sslEnable',
      title: '是否开启 SSL',
      minWidth: 120,
    },
    {
      field: 'remark',
      title: '备注',
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
      minWidth: 130,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'mail',
          nameTitle: '邮箱账号',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:mail-account:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:mail-account:delete']),
          },
        ],
      },
    },
  ];
}
