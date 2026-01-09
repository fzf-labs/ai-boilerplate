import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysOperateLogInfo } from '#/api/v1/sys-operate-log';

import { useAccess } from '@vben/access';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

const { hasAccessByCodes } = useAccess();

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'traceId',
      label: '链路编号',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入链路编号',
      },
    },
    {
      fieldName: 'adminId',
      label: '操作人',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getSysAdminSelector({ options: {} }),
        resultField: 'list',
        labelField: 'nickname',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择操作人',
      },
    },
    {
      fieldName: 'createdAt',
      label: '操作时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SysOperateLogInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'traceId',
      title: '链路编号',
      minWidth: 100,
    },
    {
      field: 'nickname',
      title: '操作人',
      minWidth: 120,
    },
    {
      field: 'IP',
      title: '操作IP',
      minWidth: 120,
    },
    {
      field: 'useragent',
      title: '浏览器',
      minWidth: 120,
    },
    {
      field: 'header',
      title: '请求头',
      minWidth: 120,
    },
    {
      field: 'URI',
      title: 'URI',
      minWidth: 120,
    },
    {
      field: 'req',
      title: '请求参数',
      minWidth: 160,
    },
    {
      field: 'resp',
      title: '响应内容',
      minWidth: 200,
    },
    {
      field: 'createdAt',
      title: '操作时间',
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
          nameField: 'action',
          nameTitle: '操作日志',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['system:operate-log:query']),
          },
        ],
      },
    },
  ];
}
