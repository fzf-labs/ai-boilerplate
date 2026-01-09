import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { getRangePickerDefaultProps } from '#/utils';

// AI 音乐生成状态的枚举
export const AiMusicStatusEnum = {
  IN_PROGRESS: 10, // 进行中
  SUCCESS: 20, // 已完成
  FAIL: 30, // 已失败
};

// 音乐生成模式枚举
export const AiMusicGenerateModeEnum = {
  DESCRIPTION: 'description', // 描述模式
  LYRIC: 'lyric', // 歌词模式
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
      fieldName: 'title',
      label: '音乐名称',
      component: 'Input',
    },
    {
      fieldName: 'status',
      label: '音乐状态',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          { label: '进行中', value: AiMusicStatusEnum.IN_PROGRESS },
          { label: '已完成', value: AiMusicStatusEnum.SUCCESS },
          { label: '已失败', value: AiMusicStatusEnum.FAIL },
        ],
      },
    },
    {
      fieldName: 'generateMode',
      label: '生成模式',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          { label: '描述模式', value: AiMusicGenerateModeEnum.DESCRIPTION },
          { label: '歌词模式', value: AiMusicGenerateModeEnum.LYRIC },
        ],
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
    {
      fieldName: 'publicStatus',
      label: '是否发布',
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: [
          { label: '是', value: true },
          { label: '否', value: false },
        ],
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
      title: '音乐名称',
      minWidth: 180,
      fixed: 'left',
      field: 'title',
    },
    {
      minWidth: 150,
      title: '用户',
      slots: { default: 'adminId' },
    },
    {
      field: 'status',
      title: '音乐状态',
      minWidth: 100,
      cellRender: {
        name: 'CellDict',
      },
    },
    {
      field: 'model',
      title: '模型',
      minWidth: 150,
    },
    {
      title: '内容',
      minWidth: 180,
      slots: { default: 'content' },
    },
    {
      field: 'duration',
      title: '时长（秒）',
      minWidth: 120,
    },
    {
      field: 'prompt',
      title: '提示词',
      minWidth: 200,
    },
    {
      field: 'lyric',
      title: '歌词',
      minWidth: 200,
    },
    {
      field: 'description',
      title: '描述',
      minWidth: 200,
    },
    {
      field: 'generateMode',
      title: '生成模式',
      minWidth: 120,
      cellRender: {
        name: 'CellDict',
      },
    },
    {
      field: 'tags',
      title: '风格标签',
      minWidth: 180,
      cellRender: {
        name: 'CellTags',
      },
    },
    {
      minWidth: 100,
      title: '是否发布',
      slots: { default: 'publicStatus' },
    },
    {
      field: 'taskId',
      title: '任务编号',
      minWidth: 200,
    },
    {
      field: 'errorMessage',
      title: '错误信息',
      minWidth: 180,
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
