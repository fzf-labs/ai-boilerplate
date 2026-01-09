import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { useAccess } from '@vben/access';
import { handleTree } from '@vben/utils';

import { z } from '#/adapter/form';
import { getSysDeptList } from '#/api/v1/sys-dept';
import { getSysPostSelector } from '#/api/v1/sys-post';
import { getSysRoleSelector } from '#/api/v1/sys-role';
import { getRangePickerDefaultProps } from '#/utils';
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
      fieldName: 'nickname',
      label: '昵称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户昵称',
      },
      rules: 'required',
    },
    {
      fieldName: 'username',
      label: '用户名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户名称',
      },
      rules: 'required',
    },
    {
      label: '密码',
      fieldName: 'password',
      component: 'InputPassword',
      componentProps: {
        placeholder: '请输入用户密码',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['id'],
        show: (values) => !values.id,
      },
    },
    {
      fieldName: 'roleId',
      label: '角色',
      component: 'ApiSelect',
      componentProps: {
        api: getSysRoleSelector,
        class: 'w-full',
        resultField: 'list',
        labelField: 'name',
        valueField: 'id',
        placeholder: '请选择角色',
      },
      rules: 'required',
    },
    {
      fieldName: 'deptId',
      label: '部门',
      component: 'ApiTreeSelect',
      componentProps: {
        api: async () => {
          const data = await getSysDeptList({ options: {} });
          return handleTree(data.list || []);
        },
        class: 'w-full',
        labelField: 'name',
        valueField: 'id',
        childrenField: 'children',
        placeholder: '请选择归属部门',
        treeDefaultExpandAll: true,
      },
      rules: 'required',
    },
    {
      fieldName: 'postId',
      label: '岗位',
      component: 'ApiSelect',
      componentProps: {
        api: getSysPostSelector,
        class: 'w-full',
        resultField: 'list',
        labelField: 'name',
        valueField: 'id',
        placeholder: '请选择岗位',
      },
      rules: 'required',
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
    {
      fieldName: 'mobile',
      label: '手机号码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入手机号码',
      },
    },
    {
      fieldName: 'email',
      label: '邮箱',
      component: 'Input',
      componentProps: {
        placeholder: '请输入邮箱',
      },
    },
    {
      fieldName: 'sex',
      label: '性别',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '未知', value: 0 },
          {
            label: '男',
            value: 1,
          },
          {
            label: '女',
            value: 2,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      defaultValue: 0,
    },
  ];
}

/** 重置密码的表单 */
export function useResetPasswordFormSchema(): VbenFormSchema[] {
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
      fieldName: 'newPassword',
      label: '新密码',
      component: 'InputPassword',
      componentProps: {
        placeholder: '请输入新密码',
      },
      rules: 'required',
    },
    {
      fieldName: 'confirmPassword',
      label: '确认密码',
      component: 'InputPassword',
      componentProps: {
        placeholder: '请再次输入新密码',
      },
      dependencies: {
        rules(values: Record<string, any>) {
          const { newPassword } = values;
          return z
            .string()
            .nonempty('确认密码不能为空')
            .refine((value) => value === newPassword, '两次输入的密码不一致');
        },
        triggerFields: ['newPassword'],
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'username',
      label: '用户名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户名称',
        allowClear: true,
      },
    },
    {
      fieldName: 'mobile',
      label: '手机号码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入手机号码',
        allowClear: true,
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
export function useGridColumns<T = SysAdminInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'username',
      title: '用户名',
      minWidth: 120,
    },
    {
      field: 'nickname',
      title: '昵称',
      minWidth: 120,
    },
    {
      field: 'avatar',
      title: '头像',
      minWidth: 120,
      cellRender: {
        name: 'CellImage',
      },
    },
    {
      field: 'mobile',
      title: '手机号码',
      minWidth: 120,
    },
    {
      field: 'email',
      title: '邮箱',
      minWidth: 120,
    },
    {
      field: 'sex',
      title: '性别',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 1: {
            return '男';
          }
          case 2: {
            return '女';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'roleName',
      title: '角色',
      minWidth: 120,
    },
    {
      field: 'deptName',
      title: '部门',
      minWidth: 120,
    },
    {
      field: 'postName',
      title: '岗位',
      minWidth: 120,
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
      minWidth: 250,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'username',
          nameTitle: '用户',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['system:admin:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['system:admin:update']),
          },
          {
            code: 'reset-password',
            text: '重置密码',
            show: hasAccessByCodes(['system:admin:update-password']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['system:admin:delete']),
          },
        ],
      },
    },
  ];
}
