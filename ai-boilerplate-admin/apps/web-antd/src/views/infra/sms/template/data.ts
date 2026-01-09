import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SmsTemplateInfo } from '#/api/v1/sms-template';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { getSmsChannelSelector } from '#/api/v1/sms-channel';
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
      fieldName: 'templateType',
      label: '短信类型',
      component: 'Select',
      componentProps: {
        options: [
          {
            label: '验证码',
            value: 1,
          },
          {
            label: '通知',
            value: 2,
          },
          {
            label: '营销',
            value: 3,
          },
        ],
        class: 'w-full',
        placeholder: '请选择短信类型',
      },
      rules: 'required',
    },
    {
      fieldName: 'templateName',
      label: '模板名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入模板名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'templateCode',
      label: '模板编码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入模板编码',
      },
      rules: 'required',
    },
    {
      fieldName: 'smsChannelId',
      label: '短信渠道',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSmsChannelSelector({ options: {} }),
        class: 'w-full',
        resultField: 'list',
        labelField: 'name',
        valueField: 'id',
        placeholder: '请选择短信渠道',
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '开启状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            label: '开启',
            value: CommonStatusEnum.ENABLE,
          },
          {
            label: '关闭',
            value: CommonStatusEnum.DISABLE,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
    {
      fieldName: 'templateContent',
      label: '模板内容',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入模板内容',
      },
      rules: 'required',
    },
    {
      fieldName: 'apiTemplateId',
      label: '运营商模板编号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入运营商模板编号',
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
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'templateType',
      label: '模板类型',
      component: 'Select',
      componentProps: {
        options: [
          {
            label: '验证码',
            value: 1,
          },
          {
            label: '通知',
            value: 2,
          },
          {
            label: '营销',
            value: 3,
          },
        ],
        allowClear: true,
        placeholder: '请选择模板类型',
      },
    },
    {
      fieldName: 'templateName',
      label: '模板名称',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入模板名称',
      },
    },
    {
      fieldName: 'templateCode',
      label: '模板编码',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入模板编码',
      },
    },
    {
      fieldName: 'smsChannelId',
      label: '短信渠道',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSmsChannelSelector({ options: {} }),
        resultField: 'list',
        labelField: 'name',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择短信渠道',
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

/** 发送短信表单 */
export function useSendSmsFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'content',
      label: '模板内容',
      component: 'Textarea',
      componentProps: {
        disabled: true,
      },
    },
    {
      fieldName: 'mobile',
      label: '手机号码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入手机号码',
      },
      rules: 'required',
    },
    {
      fieldName: 'templateParams',
      label: '模板参数',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SmsTemplateInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'smsChannelName',
      title: '短信渠道',
      minWidth: 100,
    },
    {
      field: 'templateType',
      title: '模板类型',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 1: {
            return '验证码';
          }
          case 2: {
            return '通知';
          }
          case 3: {
            return '营销';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'templateName',
      title: '模板名称',
      minWidth: 120,
    },
    {
      field: 'templateCode',
      title: '模板编码',
      minWidth: 120,
    },
    {
      field: 'templateContent',
      title: '模板内容',
      minWidth: 200,
    },
    {
      field: 'apiTemplateId',
      title: '运营商模板编号',
      minWidth: 180,
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 120,
    },
    {
      field: 'status',
      title: '开启状态',
      minWidth: 100,
      formatter: ({ cellValue }) => {
        return cellValue === 1 ? '开启' : '关闭';
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
          nameField: 'name',
          nameTitle: '短信模板',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['infra:sms-template:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['infra:sms-template:delete']),
          },
          {
            code: 'sms-send',
            text: '发送短信',
            show: hasAccessByCodes(['infra:sms-template:send-sms']),
          },
        ],
      },
    },
  ];
}
