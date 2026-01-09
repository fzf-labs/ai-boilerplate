import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysNotifyMessageInfo } from '#/api/v1/sys-notify-message';

import { useAccess } from '@vben/access';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

const { hasAccessByCodes } = useAccess();

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'receiver',
      label: '接收人',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSysAdminSelector({ options: {} }),
        resultField: 'list',
        labelField: 'nickname',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择接收人',
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
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SysNotifyMessageInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'receiverName',
      title: '接收人',
      minWidth: 120,
    },
    {
      field: 'type',
      title: '消息类型',
      minWidth: 100,
    },
    {
      field: 'subject',
      title: '消息主题',
      minWidth: 120,
    },
    {
      field: 'content',
      title: '消息内容',
      minWidth: 120,
    },
    {
      field: 'sendTime',
      title: '发送时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'readTime',
      title: '阅读时间',
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
          nameField: 'id',
          nameTitle: '站内信',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['system:notify-message:query']),
          },
        ],
      },
    },
  ];
}
