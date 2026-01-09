import type { VxeTableGridOptions } from '@vben/plugins/vxe-table';

import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn } from '#/adapter/vxe-table';
import type { SysDeptInfo } from '#/api/v1/sys-dept';

import { useAccess } from '@vben/access';
import { handleTree } from '@vben/utils';

import { z } from '#/adapter/form';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getSysDeptList } from '#/api/v1/sys-dept';

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
      fieldName: 'pid',
      label: '上级部门',
      component: 'ApiTreeSelect',
      componentProps: {
        allowClear: true,
        api: async () => {
          const data = await getSysDeptList({ options: {} });
          const list = data.list || [];
          list.unshift({
            id: '',
            name: '顶级部门',
          } as SysDeptInfo);
          return handleTree(list);
        },
        class: 'w-full',
        labelField: 'name',
        valueField: 'id',
        childrenField: 'children',
        placeholder: '请选择上级部门',
        treeDefaultExpandAll: true,
      },
      rules: 'selectRequired',
    },
    {
      fieldName: 'name',
      label: '部门名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入部门名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'adminId',
      label: '负责人',
      component: 'ApiSelect',
      componentProps: {
        api: () => getSysAdminSelector({ options: {} }),
        class: 'w-full',
        resultField: 'list',
        labelField: 'nickname',
        valueField: 'id',
        placeholder: '请选择负责人',
        allowClear: true,
      },
      rules: z.string().optional(),
    },
    {
      fieldName: 'sort',
      label: '顺序',
      component: 'InputNumber',
      componentProps: {
        min: 0,
        class: 'w-full',
        controlsPosition: 'right',
        placeholder: '请输入显示顺序',
      },
      defaultValue: 0,
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            value: 1,
            label: '启用',
          },
          {
            value: -1,
            label: '禁用',
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(1),
    },
  ];
}

/** 列表的字段 */
export function useGridColumns(
  onActionClick?: OnActionClickFn<SysDeptInfo>,
): VxeTableGridOptions<SysDeptInfo>['columns'] {
  return [
    {
      field: 'name',
      title: '部门名称',
      minWidth: 150,
      align: 'left',
      fixed: 'left',
      treeNode: true,
    },
    {
      field: 'adminName',
      title: '负责人',
      minWidth: 150,
    },
    {
      field: 'sort',
      title: '顺序',
      minWidth: 100,
    },
    {
      field: 'status',
      title: '部门状态',
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
      minWidth: 200,
      align: 'right',
      fixed: 'right',
      headerAlign: 'center',
      showOverflow: false,
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '部门',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'append',
            text: '新增下级',
            show: hasAccessByCodes(['system:dept:create']),
          },
          {
            code: 'edit',
            show: hasAccessByCodes(['system:dept:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:dept:delete']),
            disabled: (row: any) => {
              return !!(row.children && row.children.length > 0);
            },
          },
        ],
      },
    },
  ];
}
