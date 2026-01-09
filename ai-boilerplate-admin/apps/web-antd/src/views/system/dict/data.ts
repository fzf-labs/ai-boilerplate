import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { DictDatumInfo } from '#/api/v1/dict-data';
import type { DictTypeInfo } from '#/api/v1/dict-type';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getDictTypeSelector } from '#/api/v1/dict-type';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

// ============================== 字典类型 ==============================

/** 类型新增/修改的表单 */
export function useTypeFormSchema(): VbenFormSchema[] {
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
      label: '字典名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入字典名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'type',
      label: '字典类型',
      component: 'Input',
      componentProps: {
        placeholder: '请输入字典类型',
      },
      rules: 'required',
      dependencies: {
        triggerFields: [''],
        disabled: ({ values }) => values.id,
      },
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

/** 类型列表的搜索表单 */
export function useTypeGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'name',
      label: '字典名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入字典名称',
        clearable: true,
      },
    },
    {
      fieldName: 'type',
      label: '字典类型',
      component: 'Input',
      componentProps: {
        placeholder: '请输入字典类型',
        clearable: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
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
        placeholder: '请选择状态',
        clearable: true,
      },
    },
  ];
}

/** 类型列表的字段 */
export function useTypeGridColumns<T = DictTypeInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '字典名称',
      minWidth: 180,
    },
    {
      field: 'type',
      title: '字典类型',
      minWidth: 220,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 180,
      formatter: (row) => {
        return row.cellValue === 1 ? '启用' : '禁用';
      },
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 180,
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      minWidth: 120,
      title: '操作',
      field: 'operation',
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'type',
          nameTitle: '字典类型',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['system:dict:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:dict:delete']),
          },
        ],
      },
    },
  ];
}

// ============================== 字典数据 ==============================

// TODO @芋艿：后续针对 antd，增加
/**
 * 颜色选项
 */
const colorOptions = [
  { value: '', label: '无' },
  { value: 'processing', label: '主要' },
  { value: 'success', label: '成功' },
  { value: 'default', label: '默认' },
  { value: 'warning', label: '警告' },
  { value: 'error', label: '危险' },
  { value: 'pink', label: 'pink' },
  { value: 'red', label: 'red' },
  { value: 'orange', label: 'orange' },
  { value: 'green', label: 'green' },
  { value: 'cyan', label: 'cyan' },
  { value: 'blue', label: 'blue' },
  { value: 'purple', label: 'purple' },
];

/** 数据新增/修改的表单 */
export function useDataFormSchema(): VbenFormSchema[] {
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
      fieldName: 'type',
      label: '字典类型',
      component: 'ApiSelect',
      componentProps: (values) => {
        return {
          api: getDictTypeSelector,
          placeholder: '请输入字典类型',
          class: 'w-full',
          resultField: 'list',
          labelField: 'name',
          valueField: 'type',
          disabled: !!values.id,
        };
      },
      rules: 'required',
      dependencies: {
        triggerFields: [''],
      },
    },
    {
      fieldName: 'label',
      label: '数据标签',
      component: 'Input',
      componentProps: {
        placeholder: '请输入数据标签',
      },
      rules: 'required',
    },
    {
      fieldName: 'key',
      label: '数据键',
      component: 'Input',
      componentProps: {
        placeholder: '请输入数据键',
      },
      rules: 'required',
    },
    {
      fieldName: 'value',
      label: '数据键值',
      component: 'Input',
      componentProps: {
        placeholder: '请输入数据键值',
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
            value: 1,
            label: '启用',
          },
          {
            value: -1,
            label: '禁用',
          },
        ],
        placeholder: '请选择状态',
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
    {
      fieldName: 'colorType',
      label: '颜色类型',
      component: 'Select',
      componentProps: {
        options: colorOptions,
        placeholder: '请选择颜色类型',
        class: 'w-full',
      },
    },
    {
      fieldName: 'cssClass',
      label: 'CSS Class',
      component: 'Input',
      componentProps: {
        placeholder: '请输入 CSS Class',
      },
      help: '输入 hex 模式的颜色, 例如 #108ee9',
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

/** 字典数据列表搜索表单 */
export function useDataGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'label',
      label: '字典标签',
      component: 'Input',
      componentProps: {
        placeholder: '请输入字典标签',
        clearable: true,
      },
    },
    {
      fieldName: 'key',
      label: '数据键',
      component: 'Input',
      componentProps: {
        placeholder: '请输入数据键',
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
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
        placeholder: '请选择状态',
        clearable: true,
      },
    },
  ];
}

/**
 * 字典数据表格列
 */
export function useDataGridColumns<T = DictDatumInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'label',
      title: '字典标签',
      minWidth: 180,
    },
    {
      field: 'key',
      title: '数据键',
      minWidth: 100,
    },
    {
      field: 'value',
      title: '字典键值',
      minWidth: 100,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      formatter: (row) => {
        return row.cellValue === 1 ? '启用' : '禁用';
      },
    },
    {
      field: 'colorType',
      title: '颜色类型',
      minWidth: 120,
    },
    {
      field: 'cssClass',
      title: 'CSS Class',
      minWidth: 120,
    },
    {
      title: '创建时间',
      field: 'createdAt',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      minWidth: 120,
      title: '操作',
      field: 'operation',
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'label',
          nameTitle: '字典数据',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['system:dict:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:dict:delete']),
          },
        ],
      },
    },
  ];
}
