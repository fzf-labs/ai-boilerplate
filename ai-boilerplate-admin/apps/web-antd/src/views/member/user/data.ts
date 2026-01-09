import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { UserInfo } from '#/api/v1/user';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 新增/修改用户的表单 */
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
      fieldName: 'phone',
      label: '手机号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入手机号',
      },
      rules: 'required',
    },
    {
      fieldName: 'nickname',
      label: '昵称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户昵称',
      },
    },
    {
      fieldName: 'gender',
      label: '性别',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '男', value: 1 },
          { label: '女', value: 2 },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(0),
    },
    {
      fieldName: 'avatar',
      label: '头像',
      component: 'ImageUpload',
      componentProps: {
        placeholder: '请选择头像',
      },
    },
    {
      fieldName: 'profile',
      label: '简介',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入用户简介',
        rows: 3,
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
      fieldName: 'phone',
      label: '手机号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入手机号',
        allowClear: true,
      },
    },
    {
      fieldName: 'nickname',
      label: '昵称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户昵称',
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
export function useGridColumns<T = UserInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'phone',
      title: '手机号',
      minWidth: 120,
    },
    {
      field: 'nickname',
      title: '昵称',
      minWidth: 120,
    },
    {
      field: 'gender',
      title: '性别',
      minWidth: 80,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 1: {
            return '男';
          }
          case 2: {
            return '女';
          }
          default: {
            return '未知';
          }
        }
      },
    },
    {
      field: 'avatar',
      title: '头像',
      minWidth: 80,
      cellRender: {
        name: 'CellImage',
      },
    },
    {
      field: 'profile',
      title: '简介',
      minWidth: 150,
      showOverflow: 'tooltip',
    },
    {
      field: 'userMembershipInfo.membershipType',
      title: '会员类型',
      minWidth: 100,
      align: 'center',
    },
    {
      field: 'userMembershipInfo.expiredAt',
      title: '到期时间',
      minWidth: 180,
      formatter: ({ row }) => {
        const membershipInfo = (row as UserInfo).userMembershipInfo;
        if (!membershipInfo || !membershipInfo.expiredAt) {
          return '';
        }
        return new Date(membershipInfo.expiredAt).toLocaleString();
      },
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
      minWidth: 200,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'nickname',
          nameTitle: '用户',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['member:user:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['member:user:update']),
          },
          {
            code: 'more',
            text: '更多操作',
            show: hasAccessByCodes(['member:user:test-token']),
            children: [
              {
                code: 'testToken',
                text: '测试Token',
                show: hasAccessByCodes(['member:user:test-token']),
              },
            ],
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['member:user:delete']),
          },
        ],
      },
    },
  ];
}
