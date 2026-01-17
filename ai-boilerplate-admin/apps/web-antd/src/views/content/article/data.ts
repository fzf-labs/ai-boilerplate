import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { ArticleInfo } from '#/api/v1/article';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';

const { hasAccessByCodes } = useAccess();

export const ArticleStatusOptions = [
  { label: '下线', value: -1 },
  { label: '草稿', value: 0 },
  { label: '已发布', value: 1 },
];

/** 新增/修改的表单 */
export function useFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'id',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'title',
      label: '标题',
      component: 'Input',
      rules: 'required',
      componentProps: {
        placeholder: '请输入标题',
        allowClear: true,
      },
    },
    {
      fieldName: 'summary',
      label: '摘要',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入摘要（可选）',
        rows: 3,
      },
    },
    {
      fieldName: 'coverImage',
      label: '封面图',
      component: 'ImageUpload',
      componentProps: {
        maxNumber: 1,
      },
    },
    {
      fieldName: 'contentMarkdown',
      label: '内容(Markdown)',
      component: 'Textarea',
      rules: z.string().min(1, '请输入内容').max(200000),
      componentProps: {
        placeholder: '请输入 Markdown 内容',
        rows: 14,
      },
    },
    {
      fieldName: 'isRecommend',
      label: '推荐',
      component: 'Switch',
    },
    {
      fieldName: 'isHot',
      label: '热门',
      component: 'Switch',
    },
    {
      fieldName: 'status',
      label: '发布状态',
      component: 'RadioGroup',
      componentProps: {
        options: ArticleStatusOptions,
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(0),
      dependencies: {
        triggerFields: ['id'],
        show: (values) => !values.id,
      },
    },
    {
      fieldName: 'publishTime',
      label: '发布时间',
      component: 'DatePicker',
      componentProps: {
        placeholder: '请选择发布时间（可选）',
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DDTHH:mm:ssZ',
        style: { width: '100%' },
      },
      dependencies: {
        triggerFields: ['id', 'status'],
        show: (values) => !values.id && values.status === 1,
      },
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'keyword',
      label: '关键词',
      component: 'Input',
      componentProps: {
        placeholder: '请输入标题关键词',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        options: ArticleStatusOptions,
        placeholder: '请选择状态',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = ArticleInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'title',
      title: '标题',
      minWidth: 220,
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        const found = ArticleStatusOptions.find((item) => item.value === cellValue);
        return found?.label || '';
      },
    },
    {
      field: 'publishTime',
      title: '发布时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'isRecommend',
      title: '推荐',
      minWidth: 80,
      formatter: ({ cellValue }) => (cellValue ? '是' : '否'),
    },
    {
      field: 'isHot',
      title: '热门',
      minWidth: 80,
      formatter: ({ cellValue }) => (cellValue ? '是' : '否'),
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
      minWidth: 220,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'title',
          nameTitle: '文章',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['content:article:update']),
          },
          {
            code: 'publish',
            text: '发布',
            show: hasAccessByCodes(['content:article:publish']),
          },
          {
            code: 'unpublish',
            text: '下线',
            show: hasAccessByCodes(['content:article:publish']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['content:article:delete']),
          },
        ],
      },
    },
  ];
}

