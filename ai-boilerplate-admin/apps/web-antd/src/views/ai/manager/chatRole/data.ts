import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { z } from '#/adapter/form';
import { CommonStatusEnum } from '#/utils/constants';
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
      component: 'Input',
      fieldName: 'tenantId',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      component: 'Input',
      fieldName: 'adminId',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      component: 'Input',
      fieldName: 'name',
      label: '角色名称',
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'avatar',
      label: '头像',
      rules: 'required',
      componentProps: {
        placeholder: '请输入头像URL',
      },
    },
    {
      component: 'Input',
      fieldName: 'modelId',
      label: '模型编号',
      componentProps: {
        placeholder: '请输入模型编号',
      },
    },
    {
      component: 'Input',
      fieldName: 'category',
      label: '角色类别',
      componentProps: {
        placeholder: '请输入角色类别',
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: '角色描述',
      componentProps: {
        placeholder: '请输入角色描述',
        rows: 3,
      },
      rules: 'required',
    },
    {
      fieldName: 'systemMessage',
      label: '角色上下文',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入角色上下文',
        rows: 4,
      },
    },
    {
      component: 'Input',
      fieldName: 'knowledgeIds',
      label: '关联知识库编号',
      componentProps: {
        placeholder: '请输入知识库编号，多个用逗号分隔',
      },
    },
    {
      component: 'Input',
      fieldName: 'toolIds',
      label: '关联工具编号',
      componentProps: {
        placeholder: '请输入工具编号，多个用逗号分隔',
      },
    },
    {
      fieldName: 'publicStatus',
      label: '是否公开',
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
      rules: 'required',
    },
    {
      fieldName: 'sort',
      label: '角色排序',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入角色排序',
        class: 'w-full',
        min: 0,
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '状态',
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
      componentProps: {
        placeholder: '请输入角色名称',
      },
    },
    {
      fieldName: 'category',
      label: '角色类别',
      component: 'Input',
      componentProps: {
        placeholder: '请输入角色类别',
      },
    },
    {
      fieldName: 'publicStatus',
      label: '是否公开',
      component: 'Select',
      componentProps: {
        placeholder: '请选择是否公开',
        allowClear: true,
        options: [
          { label: '是', value: true },
          { label: '否', value: false },
        ],
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns(): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '角色名称',
      minWidth: 120,
    },
    {
      title: '头像',
      slots: { default: 'avatar' },
      minWidth: 80,
    },
    {
      title: '模型编号',
      field: 'modelId',
      minWidth: 120,
    },
    {
      title: '角色类别',
      field: 'category',
      minWidth: 100,
    },
    {
      title: '角色描述',
      field: 'description',
      minWidth: 150,
      showOverflow: 'tooltip',
    },
    {
      title: '角色上下文',
      field: 'systemMessage',
      minWidth: 150,
      showOverflow: 'tooltip',
    },
    {
      title: '知识库编号',
      field: 'knowledgeIds',
      minWidth: 100,
    },
    {
      title: '工具编号',
      field: 'toolIds',
      minWidth: 100,
    },
    {
      field: 'publicStatus',
      title: '是否公开',
      minWidth: 80,
      cellRender: {
        name: 'CellDict',
        props: {
          options: [
            { label: '是', value: true },
            { label: '否', value: false },
          ],
        },
      },
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 80,
      cellRender: {
        name: 'CellDict',
        props: {
          options: [
            { label: '启用', value: CommonStatusEnum.ENABLE },
            { label: '禁用', value: CommonStatusEnum.DISABLE },
          ],
        },
      },
    },
    {
      title: '排序',
      field: 'sort',
      minWidth: 80,
    },
    {
      title: '创建时间',
      field: 'createdAt',
      minWidth: 160,
    },
    {
      title: '更新时间',
      field: 'updatedAt',
      minWidth: 160,
    },
    {
      title: '操作',
      width: 130,
      fixed: 'right',
      slots: { default: 'actions' },
    },
  ];
}
