import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MallPaymentRecordInfo } from '#/api/v1/mall-payment-record';

import { useAccess } from '@vben/access';

const { hasAccessByCodes } = useAccess();

/** 支付渠道选项 */
export const PaymentChannelOptions = [
  { label: '微信', value: 'wechat' },
  { label: '支付宝', value: 'alipay' },
];

/** 支付方式选项 */
export const PaymentMethodOptions = [
  { label: '小程序', value: 'mini_program' },
  { label: 'H5', value: 'h5' },
  { label: '扫码', value: 'native' },
  { label: 'JS API', value: 'jsapi' },
];

/** 支付状态选项 */
export const PaymentStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '支付成功', value: 1 },
  { label: '支付失败', value: 2 },
  { label: '已退款', value: 3 },
];

/** 记录状态选项 */
export const RecordStatusOptions = [
  { label: '无效', value: -1 },
  { label: '正常', value: 1 },
];

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'orderId',
      label: '订单ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入订单ID',
        allowClear: true,
      },
    },
    {
      fieldName: 'transactionId',
      label: '交易流水号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入交易流水号',
        allowClear: true,
      },
    },
    {
      fieldName: 'paymentChannel',
      label: '支付渠道',
      component: 'Select',
      componentProps: {
        options: [{ label: '全部', value: '' }, ...PaymentChannelOptions],
        placeholder: '请选择支付渠道',
        allowClear: true,
      },
    },
    {
      fieldName: 'paymentStatus',
      label: '支付状态',
      component: 'Select',
      componentProps: {
        options: [{ label: '全部', value: '' }, ...PaymentStatusOptions],
        placeholder: '请选择支付状态',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MallPaymentRecordInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '记录ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'orderId',
      title: '订单ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'transactionId',
      title: '交易流水号',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'paymentChannel',
      title: '支付渠道',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallPaymentRecordInfo }) => {
          const channelMap = {
            wechat: { color: 'green', text: '微信' },
            alipay: { color: 'blue', text: '支付宝' },
          };
          const config = channelMap[
            row.paymentChannel as keyof typeof channelMap
          ] || {
            color: 'default',
            text: row.paymentChannel,
          };
          return {
            color: config.color,
            text: config.text,
          };
        },
      },
    },
    {
      field: 'paymentMethod',
      title: '支付方式',
      minWidth: 100,
      align: 'center',
      formatter: ({ row }: { row: MallPaymentRecordInfo }) => {
        const methodMap: Record<string, string> = {
          mini_program: '小程序',
          h5: 'H5',
          native: '扫码',
          jsapi: 'JS API',
        };
        return methodMap[row.paymentMethod ?? ''] || row.paymentMethod || '';
      },
    },
    {
      field: 'amount',
      title: '支付金额',
      minWidth: 120,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'currency',
      title: '币种',
      minWidth: 80,
      align: 'center',
    },
    {
      field: 'paymentStatus',
      title: '支付状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallPaymentRecordInfo }) => {
          const statusMap = {
            0: { color: 'orange', text: '待支付' },
            1: { color: 'green', text: '支付成功' },
            2: { color: 'red', text: '支付失败' },
            3: { color: 'purple', text: '已退款' },
          };
          const config = statusMap[
            row.paymentStatus as keyof typeof statusMap
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
      field: 'thirdPartyOrderNo',
      title: '第三方订单号',
      minWidth: 180,
      showOverflow: 'tooltip',
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'thirdPartyTransactionId',
      title: '第三方交易号',
      minWidth: 180,
      showOverflow: 'tooltip',
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'status',
      title: '记录状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallPaymentRecordInfo }) => {
          const statusMap = {
            '-1': { color: 'red', text: '无效' },
            '1': { color: 'green', text: '正常' },
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
      field: 'callbackTime',
      title: '回调时间',
      minWidth: 180,
      formatter: ({ cellValue }) => cellValue || '-',
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
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'id',
          nameTitle: '支付记录',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['payment_record:query']),
          },
        ],
      },
    },
  ];
}
