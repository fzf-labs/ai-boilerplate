import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { AiProviderModelInfo } from '#/api/v1/ai-provider-model';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 供应商模型 - 新增/修改的表单 */
export function useModelFormSchema(): VbenFormSchema[] {
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
      fieldName: 'platformId',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'modelType',
      label: '模型类型',
      component: 'Select',
      componentProps: {
        placeholder: '请选择模型类型',
        options: [
          { label: '文本', value: 'text' },
          { label: '图片', value: 'image' },
          { label: '音频', value: 'audio' },
          { label: '视频', value: 'video' },
        ],
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'modelId',
      label: '模型ID',
      componentProps: {
        placeholder: '请输入模型ID（如：gpt-4-turbo、claude-3-5-sonnet）',
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'modelName',
      label: '模型名称',
      componentProps: {
        placeholder: '请输入模型显示名称',
      },
      rules: 'required',
    },
    {
      component: 'Textarea',
      fieldName: 'config',
      label: '配置信息',
      componentProps: {
        placeholder: '请输入JSON格式的配置信息',
        rows: 6,
      },
    },
    {
      fieldName: 'sort',
      label: '排序',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入排序值',
        class: 'w-full',
        min: 0,
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

/** 供应商模型 - 列表的字段 */
export function useModelColumns<T = AiProviderModelInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'modelType',
      title: '模型类型',
      minWidth: 120,
    },
    {
      field: 'modelId',
      title: '模型ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'modelName',
      title: '模型名称',
      minWidth: 180,
    },
    {
      field: 'config',
      title: '配置',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'sort',
      title: '排序',
      minWidth: 80,
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
      field: 'createdAt',
      title: '创建时间',
      minWidth: 160,
      formatter: 'formatDateTime',
    },
    {
      field: 'updatedAt',
      title: '更新时间',
      minWidth: 160,
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
          nameField: 'modelName',
          nameTitle: '模型',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['ai:manager:provider:model:update']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['ai:manager:provider:model:delete']),
          },
        ],
      },
    },
  ];
}
