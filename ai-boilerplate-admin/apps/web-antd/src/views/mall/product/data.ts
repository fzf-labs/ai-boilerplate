import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MallProductInfo } from '#/api/v1/mall-product';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';

const { hasAccessByCodes } = useAccess();

/** 商品类型选项 */
export const ProductTypeOptions = [
  { label: '会员', value: 'membership' },
  { label: '服务', value: 'service' },
  { label: '商品', value: 'goods' },
];

/** 商品状态选项 */
export const ProductStatusOptions = [
  { label: '下架', value: -1 },
  { label: '待上架', value: 0 },
  { label: '在售', value: 1 },
  { label: '售罄', value: 2 },
];

/** 会员类型选项 */
export const MembershipTypeOptions = [
  { label: '普通会员', value: 'normal' },
  { label: 'VIP会员', value: 'vip' },
  { label: 'SVIP会员', value: 'svip' },
];

/** 新增/修改的表单 */
export function useFormSchema(): VbenFormSchema[] {
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
      fieldName: 'productType',
      label: '商品类型',
      component: 'Select',
      componentProps: {
        options: ProductTypeOptions,
        placeholder: '请选择商品类型',
      },
      rules: 'required',
    },
    // 会员配置（仅当商品类型为会员时显示）
    {
      fieldName: 'productConfig.membership.membershipType',
      label: '会员类型',
      component: 'Select',
      componentProps: {
        options: MembershipTypeOptions,
        placeholder: '请选择会员类型',
      },
      dependencies: {
        triggerFields: ['productType'],
        show: (values) => values?.productType === 'membership',
        required: (values) => values?.productType === 'membership',
      },
    },
    {
      fieldName: 'productConfig.membership.duration_days',
      label: '会员时长',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入会员时长天数',
        min: 1,
        style: { width: '100%' },
      },
      dependencies: {
        triggerFields: ['productType'],
        show: (values) => values?.productType === 'membership',
        required: (values) => values?.productType === 'membership',
      },
      rules: z.number().min(1, '会员时长不能小于1天'),
    },
    {
      fieldName: 'productName',
      label: '商品名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入商品名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'productDesc',
      label: '商品描述',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入商品描述（可选）',
        rows: 4,
      },
    },
    {
      fieldName: 'productImages',
      label: '商品图片',
      component: 'ImageUpload',
      componentProps: {
        maxNumber: 6,
        multiple: true,
      },
      help: '最多可上传6张商品图片',
    },
    {
      fieldName: 'productDetail',
      label: '商品详情',
      component: 'ImageUpload',
      componentProps: {
        maxNumber: 6,
        multiple: true,
      },
      help: '最多可上传6张商品详情图',
    },
    {
      fieldName: 'originalPrice',
      label: '原价',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入原价',
        min: 0,
        precision: 2,
        style: { width: '100%' },
      },
      rules: z.number().min(0, '原价不能小于0'),
    },
    {
      fieldName: 'currentPrice',
      label: '现价',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入现价',
        min: 0,
        precision: 2,
        style: { width: '100%' },
      },
      rules: z.number().min(0, '现价不能小于0'),
    },
    {
      fieldName: 'stockQuantity',
      label: '库存',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入库存（-1表示无限库存）',
        min: -1,
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'sort',
      label: '排序',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入排序值',
        min: 0,
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'RadioGroup',
      componentProps: {
        options: ProductStatusOptions,
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(0),
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'productName',
      label: '商品名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入商品名称',
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
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MallProductInfo>(
  onActionClick: OnActionClickFn<T>,
  _onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'productName',
      title: '商品名称',
      minWidth: 200,
    },
    {
      field: 'productType',
      title: '商品类型',
      minWidth: 120,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallProductInfo }) => {
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
      field: 'originalPrice',
      title: '原价',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'currentPrice',
      title: '现价',
      minWidth: 100,
      align: 'right',
      formatter: ({ cellValue }) => `¥${Number(cellValue || 0).toFixed(2)}`,
    },
    {
      field: 'stockQuantity',
      title: '库存',
      minWidth: 100,
      align: 'center',
      formatter: ({ cellValue }) =>
        cellValue === -1 ? '无限' : (cellValue || 0).toString(),
    },
    {
      field: 'soldQuantity',
      title: '已售',
      minWidth: 100,
      align: 'center',
      formatter: ({ cellValue }) => cellValue || 0,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 120,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: MallProductInfo }) => {
          const statusMap = {
            '-1': { color: 'red', text: '下架' },
            '0': { color: 'orange', text: '待上架' },
            '1': { color: 'green', text: '在售' },
            '2': { color: 'gray', text: '售罄' },
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
      field: 'productConfig.membership.membershipType',
      title: '会员类型',
      minWidth: 120,
      align: 'center',
      formatter: ({ row }) => {
        const productInfo = row as MallProductInfo;
        if (productInfo.productType !== 'membership') return '-';
        const membershipType =
          productInfo.productConfig?.membership?.membershipType;
        const typeMap = {
          normal: '普通会员',
          vip: 'VIP会员',
          svip: 'SVIP会员',
        };
        return (
          typeMap[membershipType as keyof typeof typeMap] ||
          membershipType ||
          '-'
        );
      },
    },
    {
      field: 'productConfig.membership.duration_days',
      title: '会员时长',
      minWidth: 120,
      align: 'center',
      formatter: ({ row }) => {
        const productInfo = row as MallProductInfo;
        if (productInfo.productType !== 'membership') return '-';
        const duration = productInfo.productConfig?.membership?.duration_days;
        return duration ? `${duration}天` : '-';
      },
    },
    {
      field: 'sort',
      title: '排序',
      minWidth: 80,
      align: 'center',
      formatter: ({ cellValue }) => cellValue || 0,
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
          nameField: 'productName',
          nameTitle: '商品',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['mall:product:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['mall:product:update']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['mall:product:delete']),
          },
        ],
      },
    },
  ];
}
