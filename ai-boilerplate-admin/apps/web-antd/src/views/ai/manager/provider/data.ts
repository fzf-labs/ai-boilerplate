import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { AiProviderPlatformInfo } from '#/api/v1/ai-provider-platform';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getAiProviderPlatformSelector } from '#/api/v1/ai-provider-platform';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 服务商平台 - 新增/修改的表单 */
export function usePlatformFormSchema(
  onPlatformChange?: (value: string, label: string) => void,
): VbenFormSchema[] {
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
      fieldName: 'platform',
      label: '平台标识',
      component: 'ApiSelect',
      componentProps: {
        allowClear: true,
        api: getAiProviderPlatformSelector,
        resultField: 'list',
        labelField: 'label',
        valueField: 'value',
        onChange: (value: string, option: any) => {
          // 当选择平台后，触发回调，传递 value 和 label
          if (option && option.label && onPlatformChange) {
            onPlatformChange(value, option.label);
          }
        },
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'name',
      label: '平台名称',
      componentProps: {
        placeholder: '请输入平台名称（如：OpenAI、Anthropic）',
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'APIURL',
      label: 'API 地址',
      componentProps: {
        placeholder: '请输入API地址',
      },
    },
    {
      component: 'Input',
      fieldName: 'APIKey',
      label: 'API 密钥',
      componentProps: {
        placeholder: '请输入API密钥',
      },
    },
    {
      component: 'Input',
      fieldName: 'docURL',
      label: '文档地址',
      componentProps: {
        placeholder: '请输入文档地址',
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

/** 服务商平台 - 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'platform',
      label: '平台',
      component: 'ApiSelect',
      componentProps: {
        placeholder: '请输入平台标识',
        allowClear: true,
        api: getAiProviderPlatformSelector,
        resultField: 'list',
        labelField: 'label',
        valueField: 'value',
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        allowClear: true,
        placeholder: '请选择状态',
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
  ];
}

/** 服务商平台 - 列表的字段 */
export function useGridColumns<T = AiProviderPlatformInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'platform',
      title: '平台标识',
      minWidth: 120,
    },
    {
      field: 'name',
      title: '平台名称',
      minWidth: 180,
    },
    {
      field: 'APIURL',
      title: 'API 地址',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'APIKey',
      title: 'API 密钥',
      minWidth: 200,
      cellRender: {
        name: 'CellPassword',
      },
    },
    {
      field: 'docURL',
      title: '文档地址',
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
      minWidth: 280,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '服务商',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'manageModels',
            text: '模型管理',
            show: hasAccessByCodes(['ai:manager:provider:model:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['ai:manager:provider:platform:update']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['ai:manager:provider:platform:delete']),
          },
        ],
      },
    },
  ];
}
