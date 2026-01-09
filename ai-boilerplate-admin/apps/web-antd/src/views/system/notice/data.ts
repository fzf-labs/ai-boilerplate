import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysNoticeInfo } from '#/api/v1/sys-notice';

import { useAccess } from '@vben/access';

import { z } from '#/adapter/form';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

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
      fieldName: 'type',
      label: '公告类型',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            label: '公告',
            value: 'notice',
          },
          {
            label: '通知',
            value: 'notify',
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: 'required',
    },
    {
      fieldName: 'title',
      label: '公告标题',
      component: 'Input',
      rules: 'required',
    },
    {
      fieldName: 'content',
      label: '公告内容',
      component: 'Textarea',
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '公告状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          {
            label: '启用',
            value: 1,
          },
          {
            label: '禁用',
            value: -1,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
  ];
}

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'title',
      label: '公告标题',
      component: 'Input',
      componentProps: {
        placeholder: '请输入公告标题',
        allowClear: true,
      },
    },
    {
      fieldName: 'status',
      label: '公告状态',
      component: 'Select',
      componentProps: {
        options: [
          {
            label: '启用',
            value: 1,
          },
          {
            label: '禁用',
            value: -1,
          },
        ],
        placeholder: '请选择公告状态',
        allowClear: true,
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = SysNoticeInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'type',
      title: '公告类型',
      minWidth: 100,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 'notice': {
            return '公告';
          }
          case 'notify': {
            return '通知';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'title',
      title: '公告标题',
      minWidth: 200,
    },
    {
      field: 'content',
      title: '公告内容',
      minWidth: 200,
    },
    {
      field: 'status',
      title: '公告状态',
      minWidth: 100,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case -1: {
            return '禁用';
          }
          case 1: {
            return '启用';
          }
          default: {
            return '';
          }
        }
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
      minWidth: 180,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'title',
          nameTitle: '公告',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['system:notice:update']),
          },
          {
            code: 'push',
            text: '推送',
            show: hasAccessByCodes(['system:notice:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:notice:delete']),
          },
        ],
      },
    },
  ];
}
