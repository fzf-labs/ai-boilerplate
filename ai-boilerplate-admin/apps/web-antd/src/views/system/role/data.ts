import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysRoleInfo } from '#/api/v1/sys-role';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum, SystemDataScopeEnum } from '#/utils/constants';

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
      label: '角色名称',
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'remark',
      label: '角色备注',
      component: 'Textarea',
    },
    {
      component: 'Select',
      fieldName: 'dataScope',
      label: '权限范围',
      componentProps: {
        class: 'w-full',
        options: [
          {
            label: '全部数据权限',
            value: SystemDataScopeEnum.ALL,
          },
          {
            label: '部门及子部门数据权限',
            value: SystemDataScopeEnum.DEPT_AND_CHILD,
          },
          {
            label: '部门数据权限',
            value: SystemDataScopeEnum.DEPT_ONLY,
          },
          {
            label: '仅本人数据权限',
            value: SystemDataScopeEnum.SELF,
          },
        ],
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
      fieldName: 'sort',
      label: '显示顺序',
      component: 'InputNumber',
      componentProps: {
        min: 0,
        class: 'w-full',
        controlsPosition: 'right',
        placeholder: '请输入显示顺序',
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '角色状态',
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
      fieldName: 'name',
      label: '角色名称',
      component: 'Input',
    },
    {
      fieldName: 'status',
      label: '角色状态',
      component: 'Select',
      componentProps: {
        allowClear: true,
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
export function useGridColumns<T = SysRoleInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '角色名称',
      minWidth: 200,
    },
    {
      field: 'remark',
      title: '角色备注',
      minWidth: 200,
    },
    {
      field: 'dataScope',
      title: '权限范围',
      minWidth: 100,
      formatter: (row) => {
        switch (row.cellValue) {
          case SystemDataScopeEnum.ALL: {
            return '全部数据权限';
          }
          case SystemDataScopeEnum.DEPT_AND_CHILD: {
            return '部门及子部门数据权限';
          }
          case SystemDataScopeEnum.DEPT_ONLY: {
            return '部门数据权限';
          }
          case SystemDataScopeEnum.SELF: {
            return '仅本人数据权限';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'sort',
      title: '显示顺序',
      minWidth: 100,
    },
    {
      field: 'status',
      title: '角色状态',
      minWidth: 100,
      formatter: (row) => {
        return row.cellValue === 1 ? '启用' : '禁用';
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
      width: 240,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '角色',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['system:role:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:role:delete']),
          },
        ],
      },
    },
  ];
}
