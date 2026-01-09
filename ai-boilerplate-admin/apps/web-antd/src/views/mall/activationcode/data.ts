import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MallActivationCodeInfo } from '#/api/v1/mall-activation-code';

import { useAccess } from '@vben/access';

import { getMallProductSelector } from '#/api/v1/mall-product';

const { hasAccessByCodes } = useAccess();

/** 商品类型选项 */
export const ProductTypeOptions = [
  { label: '会员', value: 'membership' },
  { label: '服务', value: 'service' },
];

/** 激活码状态选项 */
export const ActivationCodeStatusOptions = [
  { label: '已退款', value: -2 },
  { label: '禁用', value: -1 },
  { label: '库存', value: 0 },
  { label: '已售出', value: 1 },
  { label: '已激活', value: 2 },
  { label: '已过期', value: 3 },
];

/** 平台选项 */
export const PlatformOptions = [
  { label: '淘宝', value: 'taobao' },
  { label: '京东', value: 'jd' },
  { label: '拼多多', value: 'pdd' },
  { label: '官网', value: 'official' },
  { label: '其他', value: 'other' },
];

/** 新增表单 - 批量生成激活码（BatchGenerateMallActivationCodeReq）*/
export function useCreateFormSchema(): VbenFormSchema[] {
  // 接口字段：productType(必填), productId(必填), validSt(必填), validEd(必填), num(必填), platform, remark
  return [
    {
      fieldName: 'productType',
      label: '商品类型',
      component: 'Select',
      componentProps: {
        options: ProductTypeOptions,
        placeholder: '请选择商品类型',
      },
      rules: 'required',
    },
    {
      fieldName: 'productId',
      label: '商品',
      component: 'ApiSelect',
      componentProps: {
        api: async () => {
          const data = await getMallProductSelector({ params: {} });
          return data.list || [];
        },
        labelField: 'productName',
        valueField: 'id',
        placeholder: '请选择商品',
        class: 'w-full',
      },
      rules: 'required',
    },
    {
      fieldName: 'validSt',
      label: '有效期开始',
      component: 'DatePicker',
      componentProps: {
        placeholder: '请选择有效期开始时间',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        style: { width: '100%' },
      },
      rules: 'required',
    },
    {
      fieldName: 'validEd',
      label: '有效期截止',
      component: 'DatePicker',
      componentProps: {
        placeholder: '请选择有效期截止时间',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        style: { width: '100%' },
      },
      rules: 'required',
    },
    {
      fieldName: 'num',
      label: '生成数量',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入生成数量',
        min: 1,
        max: 1000,
        style: { width: '100%' },
      },
      rules: 'required',
      defaultValue: 1,
    },
    {
      fieldName: 'platform',
      label: '平台',
      component: 'Select',
      componentProps: {
        options: PlatformOptions,
        placeholder: '请选择平台（可选）',
        allowClear: true,
      },
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注（可选）',
        rows: 4,
      },
    },
  ];
}

/** 编辑表单 - 更新激活码（UpdateMallActivationCodeReq）*/
export function useEditFormSchema(): VbenFormSchema[] {
  // 接口字段：id(必填), platform, platformSoldAt, platformOrderNo, platformBuyerId, platformBuyerName, remark
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
      fieldName: 'platform',
      label: '平台',
      component: 'Select',
      componentProps: {
        options: PlatformOptions,
        placeholder: '请选择平台',
        allowClear: true,
      },
    },
    {
      fieldName: 'platformSoldAt',
      label: '平台售出时间',
      component: 'DatePicker',
      componentProps: {
        placeholder: '请选择平台售出时间',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'platformOrderNo',
      label: '平台订单号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入平台订单号',
      },
    },
    {
      fieldName: 'platformBuyerId',
      label: '平台买家ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入平台买家ID',
      },
    },
    {
      fieldName: 'platformBuyerName',
      label: '平台买家昵称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入平台买家昵称',
      },
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注（可选）',
        rows: 4,
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'code',
      label: '激活码',
      component: 'Input',
      componentProps: {
        placeholder: '请输入激活码',
        allowClear: true,
      },
    },
    {
      fieldName: 'batchNo',
      label: '批次号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入批次号',
        allowClear: true,
      },
    },
    {
      fieldName: 'productType',
      label: '商品类型',
      component: 'Select',
      componentProps: {
        options: [{ label: '全部', value: '' }, ...ProductTypeOptions],
        placeholder: '请选择商品类型',
        allowClear: true,
      },
    },
    {
      fieldName: 'productId',
      label: '商品',
      component: 'ApiSelect',
      componentProps: {
        api: async () => {
          const data = await getMallProductSelector({ params: {} });
          return data.list || [];
        },
        labelField: 'productName',
        valueField: 'id',
        placeholder: '请选择商品',
        allowClear: true,
        showSearch: true,
        filterOption: true,
      },
    },
    {
      fieldName: 'userId',
      label: '用户ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户ID',
        allowClear: true,
      },
    },
    {
      fieldName: 'platform',
      label: '平台',
      component: 'Select',
      componentProps: {
        options: [{ label: '全部', value: '' }, ...PlatformOptions],
        placeholder: '请选择平台',
        allowClear: true,
      },
    },
    {
      fieldName: 'platformOrderNo',
      label: '平台订单号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入平台订单号',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        options: [{ label: '全部', value: '' }, ...ActivationCodeStatusOptions],
        placeholder: '请选择状态',
        allowClear: true,
      },
    },
    {
      fieldName: 'activatedAt',
      label: '激活时间',
      component: 'RangePicker',
      componentProps: {
        placeholder: ['开始时间', '结束时间'],
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'createdAt',
      label: '创建时间',
      component: 'RangePicker',
      componentProps: {
        placeholder: ['开始时间', '结束时间'],
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        style: { width: '100%' },
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MallActivationCodeInfo>(
  onActionClick: OnActionClickFn<T>,
  _onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'code',
      title: '激活码',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'batchNo',
      title: '批次号',
      minWidth: 180,
      showOverflow: 'tooltip',
    },
    {
      field: 'productType',
      title: '商品类型',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallActivationCodeInfo }) => {
          const typeMap = {
            membership: { color: 'purple', text: '会员' },
            service: { color: 'blue', text: '服务' },
            goods: { color: 'green', text: '商品' },
          };
          const config = typeMap[row.productType as keyof typeof typeMap] || {
            color: 'default',
            text: row.productType || '-',
          };
          return {
            color: config.color,
            text: config.text,
          };
        },
      },
    },
    {
      field: 'productName',
      title: '商品名称',
      minWidth: 150,
      showOverflow: 'tooltip',
    },
    {
      field: 'validSt',
      title: '有效期开始',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'validEd',
      title: '有效期截止',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'userNickname',
      title: '用户昵称',
      minWidth: 150,
      showOverflow: 'tooltip',
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'activatedAt',
      title: '激活时间',
      minWidth: 180,
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'platform',
      title: '平台',
      minWidth: 100,
      align: 'center',
      formatter: ({ cellValue }) => {
        const platformMap: Record<string, string> = {
          taobao: '淘宝',
          jd: '京东',
          pdd: '拼多多',
          official: '官网',
          other: '其他',
        };
        return cellValue ? platformMap[cellValue] || cellValue : '-';
      },
    },
    {
      field: 'platformOrderNo',
      title: '平台订单号',
      minWidth: 180,
      showOverflow: 'tooltip',
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'platformBuyerName',
      title: '买家昵称',
      minWidth: 120,
      showOverflow: 'tooltip',
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'platformSoldAt',
      title: '平台售出时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallActivationCodeInfo }) => {
          const statusMap = {
            '-2': { color: 'red', text: '已退款' },
            '-1': { color: 'red', text: '禁用' },
            '0': { color: 'orange', text: '库存' },
            '1': { color: 'blue', text: '已售出' },
            '2': { color: 'green', text: '已激活' },
            '3': { color: 'gray', text: '已过期' },
          };
          const config = statusMap[
            (row.status ?? 0).toString() as keyof typeof statusMap
          ] || {
            color: 'default',
            text: '未知',
          };
          return {
            color: config.color,
            text: config.text,
          };
        },
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
      minWidth: 200,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'code',
          nameTitle: '激活码',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['mall:activationcode:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['mall:activationcode:update']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['mall:activationcode:delete']),
          },
        ],
      },
    },
  ];
}
