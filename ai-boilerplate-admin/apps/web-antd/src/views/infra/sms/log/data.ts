import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SmsLogInfo } from '#/api/v1/sms-log';

import { useAccess } from '@vben/access';

import { getRangePickerDefaultProps } from '#/utils';

const { hasAccessByCodes } = useAccess();

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'mobile',
      label: '手机号',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入手机号',
      },
    },
    {
      fieldName: 'sendStatus',
      label: '发送状态',
      component: 'Select',
      componentProps: {
        options: [],
        allowClear: true,
        placeholder: '请选择发送状态',
      },
    },
    {
      fieldName: 'sendTime',
      label: '发送时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
    {
      fieldName: 'receiveStatus',
      label: '接收状态',
      component: 'Select',
      componentProps: {
        options: [],
        allowClear: true,
        placeholder: '请选择接收状态',
      },
    },
    {
      fieldName: 'receiveTime',
      label: '接收时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SmsLogInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '编号',
      minWidth: 100,
    },
    {
      field: 'mobile',
      title: '手机号',
      minWidth: 120,
    },
    {
      field: 'userId',
      title: '用户ID',
      minWidth: 120,
    },
    {
      field: 'smsChannelId',
      title: '短信渠道',
      minWidth: 120,
    },
    {
      field: 'smsTemplateId',
      title: '短信模板',
      minWidth: 120,
    },
    {
      field: 'smsParamsContent',
      title: '短信内容',
      minWidth: 300,
      showOverflow: true,
    },
    {
      field: 'sendStatus',
      title: '发送状态',
      minWidth: 120,
    },
    {
      field: 'sendTime',
      title: '发送时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'receiveStatus',
      title: '接收状态',
      minWidth: 120,
    },
    {
      field: 'receiveTime',
      title: '接收时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 120,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'mobile',
          nameTitle: '短信日志',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['infra:sms-log:query']),
          },
        ],
      },
    },
  ];
}
