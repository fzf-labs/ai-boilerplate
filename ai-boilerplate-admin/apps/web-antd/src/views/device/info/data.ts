import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { DeviceInfo } from '#/api/v1/device';

import { useAccess } from '@vben/access';

import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 设备注册表单 */
export function useFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'sn',
      label: '设备序列号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入设备序列号',
      },
      rules: 'required',
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'sn',
      label: '设备序列号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入设备序列号',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '设备状态',
      component: 'Select',
      componentProps: {
        placeholder: '请选择设备状态',
        allowClear: true,
        options: [
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
      },
    },
    {
      fieldName: 'onlineSearch',
      label: '在线状态',
      component: 'Select',
      componentProps: {
        placeholder: '请选择是否在线',
        allowClear: true,
        options: [
          { label: '全部', value: 'all' },
          { label: '在线', value: 'online' },
          { label: '离线', value: 'offline' },
        ],
      },
    },
    {
      fieldName: 'registryTime',
      label: '激活时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = DeviceInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'sn',
      title: '设备序列号',
      minWidth: 220,
      fixed: 'left',
    },
    {
      field: 'name',
      title: '设备名称',
      minWidth: 120,
    },
    {
      field: 'brand',
      title: '品牌',
      minWidth: 100,
    },
    {
      field: 'model',
      title: '型号',
      minWidth: 120,
    },
    {
      field: 'mac',
      title: 'MAC地址',
      minWidth: 150,
    },
    {
      field: 'appVersion',
      title: 'APP版本',
      minWidth: 100,
    },
    {
      field: 'androidVersion',
      title: '安卓版本',
      minWidth: 100,
    },
    {
      field: 'registryTime',
      title: '激活时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'online',
      title: '是否在线',
      minWidth: 100,
      align: 'center',
      formatter: ({ cellValue }) => {
        return cellValue ? '在线' : '离线';
      },
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        attrs: { beforeChange: onStatusChange },
        name: 'CellSwitch',
        props: {
          checkedValue: 1,
          unCheckedValue: -1,
        },
      },
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 300,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '设备',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: true,
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['kid:device:delete']),
          },
        ],
      },
    },
  ];
}
