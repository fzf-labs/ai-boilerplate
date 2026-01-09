import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

/** 列表的搜索表单 */
export function useGridFormSchemaConversation(): VbenFormSchema[] {
  return [
    {
      fieldName: 'adminId',
      label: '用户编号',
      component: 'ApiSelect',
      componentProps: {
        api: getSysAdminSelector,
        labelField: 'nickname',
        valueField: 'id',
      },
    },
    {
      fieldName: 'title',
      label: '聊天标题',
      component: 'Input',
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
export function useGridColumnsConversation(): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '对话编号',
      fixed: 'left',
      minWidth: 200,
    },
    {
      field: 'title',
      title: '对话标题',
      minWidth: 200,
      fixed: 'left',
    },
    {
      title: '用户',
      minWidth: 150,
      slots: { default: 'adminId' },
    },
    {
      field: 'roleId',
      title: '角色ID',
      minWidth: 150,
    },
    {
      field: 'model',
      title: '模型标识',
      minWidth: 150,
    },
    {
      field: 'temperature',
      title: '温度参数',
      minWidth: 100,
    },
    {
      title: '回复 Token 数',
      field: 'maxTokens',
      minWidth: 130,
    },
    {
      title: '上下文数量',
      field: 'maxContexts',
      minWidth: 130,
    },
    {
      field: 'pinned',
      title: '是否置顶',
      minWidth: 100,
      cellRender: {
        name: 'CellDict',
      },
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      title: '操作',
      width: 180,
      fixed: 'right',
      slots: { default: 'actions' },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchemaMessage(): VbenFormSchema[] {
  return [
    {
      fieldName: 'adminId',
      label: '用户编号',
      component: 'ApiSelect',
      componentProps: {
        api: getSysAdminSelector,
        labelField: 'nickname',
        valueField: 'id',
      },
    },
    {
      fieldName: 'type',
      label: '消息类型',
      component: 'Input',
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
export function useGridColumnsMessage(): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '消息编号',
      fixed: 'left',
      minWidth: 200,
    },
    {
      field: 'conversationId',
      title: '对话编号',
      minWidth: 200,
      fixed: 'left',
    },
    {
      title: '用户',
      minWidth: 150,
      slots: { default: 'adminId' },
    },
    {
      field: 'roleId',
      title: '角色ID',
      minWidth: 150,
    },
    {
      field: 'type',
      title: '消息类型',
      minWidth: 100,
    },
    {
      field: 'model',
      title: '模型标识',
      minWidth: 150,
    },
    {
      field: 'content',
      title: '消息内容',
      minWidth: 300,
    },
    {
      field: 'replyId',
      title: '回复消息编号',
      minWidth: 200,
    },
    {
      title: '携带上下文',
      field: 'useContext',
      cellRender: {
        name: 'CellDict',
      },
      minWidth: 120,
    },
    {
      field: 'createdAt',
      title: '创建时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      title: '操作',
      width: 130,
      fixed: 'right',
      slots: { default: 'actions' },
    },
  ];
}
