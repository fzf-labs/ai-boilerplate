import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MallOrderInfo } from '#/api/v1/mall-order';

import { useAccess } from '@vben/access';

const { hasAccessByCodes } = useAccess();

/** 商品类型选项 */
export const ProductTypeOptions = [
  { label: '会员', value: 'membership' },
  { label: '服务', value: 'service' },
  { label: '商品', value: 'goods' },
];

/** 支付状态选项 */
export const PaymentStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '已支付', value: 1 },
  { label: '支付失败', value: 2 },
  { label: '已退款', value: 3 },
];

/** 订单状态选项 */
export const OrderStatusOptions = [
  { label: '待付款', value: 'pendingPayment' },
  { label: '待发货', value: 'pendingDelivery' },
  { label: '待收货', value: 'pendingReceipt' },
  { label: '已完成', value: 'completed' },
  { label: '已取消', value: 'canceled' },
  { label: '已退款', value: 'refunded' },
];

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
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
      fieldName: 'orderId',
      label: '订单ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入订单ID',
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
      label: '商品ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入商品ID',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MallOrderInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '订单ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'userId',
      title: '用户ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'productType',
      title: '商品类型',
      minWidth: 120,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallOrderInfo }) => {
          const typeMap = {
            membership: { color: 'purple', text: '会员' },
            service: { color: 'blue', text: '服务' },
            goods: { color: 'green', text: '商品' },
          };
          const config = typeMap[row.productType as keyof typeof typeMap] || {
            color: 'default',
            text: row.productType,
          };
          return {
            color: config.color,
            text: config.text,
          };
        },
      },
    },
    {
      field: 'productId',
      title: '商品ID',
      minWidth: 200,
      showOverflow: 'tooltip',
    },
    {
      field: 'originalAmount',
      title: '原价',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'discountAmount',
      title: '优惠金额',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'actualAmount',
      title: '实付金额',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'refundAmount',
      title: '退款金额',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'paymentMethod',
      title: '支付方式',
      minWidth: 100,
      align: 'center',
    },
    {
      field: 'paymentStatus',
      title: '支付状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallOrderInfo }) => {
          const statusMap = {
            0: { color: 'orange', text: '待支付' },
            1: { color: 'green', text: '已支付' },
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
      field: 'status',
      title: '订单状态',
      minWidth: 120,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallOrderInfo }) => {
          const statusMap = {
            pendingPayment: { color: 'orange', text: '待付款' },
            pendingDelivery: { color: 'blue', text: '待发货' },
            pendingReceipt: { color: 'cyan', text: '待收货' },
            completed: { color: 'green', text: '已完成' },
            canceled: { color: 'red', text: '已取消' },
            refunded: { color: 'purple', text: '已退款' },
          };
          const config = statusMap[row.status as keyof typeof statusMap] || {
            color: 'default',
            text: row.status,
          };
          return {
            color: config.color,
            text: config.text,
          };
        },
      },
    },
    {
      field: 'paymentTime',
      title: '支付时间',
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
      minWidth: 160,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'id',
          nameTitle: '订单',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['order:query']),
          },
          {
            code: 'payment',
            text: '支付记录',
            show: hasAccessByCodes(['payment_record:query']),
          },
        ],
      },
    },
  ];
}
