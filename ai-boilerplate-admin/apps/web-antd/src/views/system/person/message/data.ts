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
      fieldName: 'sender',
      label: '发送人',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSysAdminSelector({ options: {} }),
        resultField: 'list',
        labelField: 'nickname',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择发送人',
      },
    },
    {
      fieldName: 'sendTime',
      label: '发送时间',
      component: 'RangePicker',
      componentProps: {
        allowClear: true,
        ...getRangePickerDefaultProps(),
      },
    },
    {
      fieldName: 'readStatus',
      label: '已读状态',
      component: 'Select',
      componentProps: {
        options: [
          { label: '全部', value: 0 },
          { label: '已读', value: 1 },
          { label: '未读', value: -1 },
        ],
        allowClear: true,
        placeholder: '请选择是否已读',
      },
      defaultValue: 0,
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SysNotifyMessageInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      title: '',
      width: 40,
      type: 'checkbox',
    },
    {
      field: 'senderName',
      title: '发送人',
      minWidth: 180,
    },
    {
      field: 'sendTime',
      title: '发送时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'type',
      title: '消息类型',
      minWidth: 120,
    },
    {
      field: 'subject',
      title: '消息主题',
      minWidth: 200,
    },
    {
      field: 'content',
      title: '消息内容',
      minWidth: 300,
    },
    {
      field: 'receiverName',
      title: '接收人',
      minWidth: 150,
    },
    {
      field: 'readStatus',
      title: '已读状态',
      minWidth: 100,
      cellRender: {
        name: 'CellTag',
        props: (params: { row: SysNotifyMessageInfo }) => {
          const isRead = params.row.readTime && params.row.readTime !== '';
          return {
            color: isRead ? 'success' : 'warning',
            text: isRead ? '已读' : '未读',
          };
        },
      },
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
            text: '查看',
          },
          {
            code: 'read',
            text: '标记已读',
            show: (row: any) =>
              (!row.readTime || row.readTime === '') &&
              hasAccessByCodes(['system:person:message:update']),
          },
        ],
      },
    },
  ];
}
