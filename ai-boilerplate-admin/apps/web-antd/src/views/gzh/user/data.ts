import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { WxGzhUserInfo } from '#/api/v1/wx-gzh-user';

import { useAccess } from '@vben/access';

import { getWxGzhAccountSelector } from '#/api/v1/wx-gzh-account';

const { hasAccessByCodes } = useAccess();

/** 列表搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'appId',
      label: '公众号',
      component: 'ApiSelect',
      componentProps: () => ({
        api: async () => await getWxGzhAccountSelector({ options: {} }),
        resultField: 'list',
        labelField: 'name',
        valueField: 'appId',
        autoSelect: 'first',
      }),
      rules: 'required',
    },
    {
      fieldName: 'openid',
      label: '用户标识',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户openid',
        allowClear: true,
      },
    },
    {
      fieldName: 'nickname',
      label: '昵称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户昵称',
        allowClear: true,
      },
    },
    {
      fieldName: 'subscribeStatus',
      label: '关注状态',
      component: 'Select',
      componentProps: {
        options: [
          { label: '已关注', value: 1 },
          { label: '未关注', value: 0 },
        ],
        placeholder: '请选择关注状态',
        allowClear: true,
      },
      defaultValue: 1,
    },
  ];
}

/** 列表字段 */
export function useGridColumns<T = WxGzhUserInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'openid',
      title: '用户标识',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'nickname',
      title: '昵称',
      minWidth: 120,
      showOverflow: 'tooltip',
    },
    {
      field: 'headImageURL',
      title: '头像',
      minWidth: 80,
      align: 'center',
      cellRender: {
        name: 'CellImage',
        props: {
          height: 40,
          width: 40,
          style: 'border-radius: 50%',
        },
      },
    },
    {
      field: 'subscribeStatus',
      title: '关注状态',
      minWidth: 100,
      align: 'center',
      formatter: ({ cellValue }) => {
        return cellValue === 1 ? '已关注' : '未关注';
      },
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: WxGzhUserInfo }) => {
          return {
            color: row.subscribeStatus === 1 ? 'success' : 'default',
          };
        },
      },
    },
    {
      field: 'subscribeTime',
      title: '关注时间',
      minWidth: 160,
      formatter: 'formatDateTime',
    },
    {
      field: 'language',
      title: '语言',
      minWidth: 80,
    },
    {
      field: 'country',
      title: '国家',
      minWidth: 80,
    },
    {
      field: 'province',
      title: '省份',
      minWidth: 80,
    },
    {
      field: 'city',
      title: '城市',
      minWidth: 80,
    },
    {
      field: 'tagIds',
      title: '标签',
      minWidth: 120,
      showOverflow: 'tooltip',
    },
    {
      field: 'remark',
      title: '备注',
      minWidth: 120,
      showOverflow: 'tooltip',
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
      minWidth: 150,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'nickname',
          nameTitle: '用户',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['gzh:user:query']),
          },
        ],
      },
    },
  ];
}
