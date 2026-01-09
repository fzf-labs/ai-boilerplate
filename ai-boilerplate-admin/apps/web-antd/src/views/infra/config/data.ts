import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { ConfigDatumInfo } from '#/api/v1/config-data';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
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
      fieldName: 'name',
      label: '参数名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入参数名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'key',
      label: '参数键名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入参数键名',
      },
      rules: 'required',
    },
    {
      fieldName: 'value',
      label: '参数键值',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入参数键值 (支持JSON格式)',
        autoSize: { minRows: 3, maxRows: 100 },
      },
      rules: 'required',
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
      fieldName: 'name',
      label: '参数名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入参数名称',
        clearable: true,
      },
    },
    {
      fieldName: 'key',
      label: '参数键名',
      component: 'Input',
      componentProps: {
        placeholder: '请输入参数键名',
        clearable: true,
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
export function useGridColumns<T = ConfigDatumInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '参数名称',
      minWidth: 200,
    },
    {
      field: 'key',
      title: '参数键名',
      minWidth: 200,
    },
    {
      field: 'value',
      title: '参数键值',
      minWidth: 150,
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 150,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case CommonStatusEnum.DISABLE: {
            return '停用';
          }
          case CommonStatusEnum.ENABLE: {
            return '正常';
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
          nameField: 'name',
          nameTitle: '参数',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:config:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:config:delete']),
          },
        ],
      },
    },
  ];
}
