import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

// AI 图像生成状态的枚举
export const AiImageStatusEnum = {
  IN_PROGRESS: 10, // 进行中
  SUCCESS: 20, // 已完成
  FAIL: 30, // 已失败
};

// 平台枚举
export const AiImagePlatformEnum = {
  STABLE_DIFFUSION: 'stable-diffusion',
  MIDJOURNEY: 'midjourney',
  DALL_E: 'dall-e',
};

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
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
      fieldName: 'platform',
      label: '平台',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入平台名称',
      },
    },
    {
      fieldName: 'status',
      label: '绘画状态',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          { label: '进行中', value: AiImageStatusEnum.IN_PROGRESS },
          { label: '已完成', value: AiImageStatusEnum.SUCCESS },
          { label: '已失败', value: AiImageStatusEnum.FAIL },
        ],
      },
    },
    {
      fieldName: 'publicStatus',
      label: '是否发布',
      component: 'Select',
      componentProps: {
        options: [
          { label: '是', value: true },
          { label: '否', value: false },
        ],
        allowClear: true,
      },
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
export function useGridColumns(): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'id',
      title: '编号',
      minWidth: 200,
      fixed: 'left',
    },
    {
      title: '图片',
      minWidth: 110,
      fixed: 'left',
      slots: { default: 'picURL' },
    },
    {
      minWidth: 150,
      title: '用户',
      slots: { default: 'adminId' },
    },
    {
      field: 'platform',
      title: '平台',
      minWidth: 120,
    },
    {
      field: 'model',
      title: '模型',
      minWidth: 150,
    },
    {
      field: 'status',
      title: '绘画状态',
      minWidth: 100,
      cellRender: {
        name: 'CellDict',
      },
    },
    {
      minWidth: 100,
      title: '是否发布',
      slots: { default: 'publicStatus' },
    },
    {
      field: 'prompt',
      title: '提示词',
      minWidth: 200,
    },
    {
      field: 'width',
      title: '宽度',
      minWidth: 100,
    },
    {
      field: 'height',
      title: '高度',
      minWidth: 100,
    },
    {
      field: 'errorMessage',
      title: '错误信息',
      minWidth: 180,
    },
    {
      field: 'taskId',
      title: '任务编号',
      minWidth: 200,
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
