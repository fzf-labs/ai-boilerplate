import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SmsChannelInfo } from '#/api/v1/sms-channel';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getSmsChannelOperator } from '#/api/v1/sms-channel';
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
      label: '渠道名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入短信签名',
      },
      rules: 'required',
    },
    {
      fieldName: 'operator',
      label: '运营商',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSmsChannelOperator({ options: {} }),
        resultField: 'list',
        labelField: 'name',
        valueField: 'operator',
        placeholder: '请选择运营商',
      },
      rules: 'required',
    },
    {
      fieldName: 'APIKey',
      label: '短信 API 的KEY',
      component: 'Input',
      componentProps: {
        placeholder: '请输入短信 API 的KEY',
      },
      rules: 'required',
    },
    {
      fieldName: 'APISecret',
      label: '短信 API 的密钥',
      component: 'Input',
      componentProps: {
        placeholder: '请输入短信 API 的密钥',
      },
    },
    {
      fieldName: 'callbackURL',
      label: '短信发送回调 URL',
      component: 'Input',
      componentProps: {
        placeholder: '请输入短信发送回调 URL',
      },
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
      fieldName: 'channelName',
      label: '渠道名称',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入短信签名',
      },
    },
    {
      fieldName: 'operator',
      label: '运营商',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSmsChannelOperator({ options: {} }),
        resultField: 'list',
        labelField: 'name',
        valueField: 'operator',
        allowClear: true,
        placeholder: '请选择运营商',
      },
    },
    {
      fieldName: 'status',
      label: '状态',
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
export function useGridColumns<T = SmsChannelInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'name',
      title: '渠道名称',
      minWidth: 120,
    },
    {
      field: 'operatorName',
      title: '运营商',
      minWidth: 120,
    },
    {
      field: 'APIKey',
      title: '短信 API 的KEY',
      minWidth: 180,
    },
    {
      field: 'APISecret',
      title: '短信 API 的密钥',
      minWidth: 180,
    },
    {
      field: 'callbackURL',
      title: '短信发送回调 URL',
      minWidth: 180,
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 120,
    },
    {
      field: 'status',
      title: '启用状态',
      minWidth: 100,
      formatter: (row) => {
        switch (row.cellValue) {
          case CommonStatusEnum.DISABLE: {
            return '禁用';
          }
          case CommonStatusEnum.ENABLE: {
            return '启用';
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
      field: 'updatedAt',
      title: '更新时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 180,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'channelName',
          nameTitle: '短信渠道',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:sms-channel:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:sms-channel:delete']),
          },
        ],
      },
    },
  ];
}
