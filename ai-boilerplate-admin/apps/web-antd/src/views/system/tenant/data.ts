import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysTenantInfo } from '#/api/v1/sys-tenant';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
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
      label: '租户名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入租户名称',
      },
      rules: 'required',
    },
    {
      label: '过期时间',
      fieldName: 'expireTime',
      component: 'DatePicker',
      componentProps: {
        format: 'YYYY-MM-DD',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        class: 'w-full',
        placeholder: '请选择过期时间',
      },
      rules: 'required',
    },
    {
      fieldName: 'menuIds',
      label: '菜单权限',
      component: 'Input',
      formItemClass: 'items-start',
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '租户状态',
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
      label: '管理员名称',
      fieldName: 'username',
      component: 'Input',
      componentProps: {
        placeholder: '请输入管理员名称',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['id'],
        show: (values) => !values.id,
      },
    },
    {
      label: '管理员密码',
      fieldName: 'password',
      component: 'InputPassword',
      componentProps: {
        placeholder: '请输入管理员密码',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['id'],
        show: (values) => !values.id,
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'name',
      label: '租户名',
      component: 'Input',
      componentProps: {
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
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
export function useGridColumns<T = SysTenantInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '租户名',
      minWidth: 180,
    },
    {
      field: 'adminName',
      title: '管理员',
      minWidth: 180,
    },
    {
      field: 'expireTime',
      title: '过期时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 100,
    },
    {
      field: 'status',
      title: '租户状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: SysTenantInfo }) => ({
          color: row.status === 1 ? 'success' : 'default',
          text: row.status === 1 ? '启用' : '禁用',
        }),
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
          nameTitle: '租户名',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['system:tenant:query']),
          },
          {
            code: 'edit',
            show: hasAccessByCodes(['system:tenant:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:tenant:delete']),
          },
        ],
      },
    },
  ];
}
