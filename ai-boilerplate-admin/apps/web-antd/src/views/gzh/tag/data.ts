import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { WxGzhTagInfo } from '#/api/v1/wx-gzh-tag';

import { useAccess } from '@vben/access';

import { getWxGzhAccountSelector } from '#/api/v1/wx-gzh-account';

const { hasAccessByCodes } = useAccess();

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
      fieldName: 'appId',
      label: '公众号',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'tagId',
      label: '标签编号',
      component: 'Input',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      fieldName: 'name',
      label: '标签名称',
      component: 'Input',
      rules: 'required',
      componentProps: {
        placeholder: '请输入名称',
      },
    },
  ];
}

/** 搜索表单配置 */
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
        placeholder: '请选择公众号',
        autoSelect: 'first',
      }),
      rules: 'required',
    },
  ];
}
/** 表格列配置 */
export function useGridColumns<T = WxGzhTagInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      title: '标签编号',
      field: 'tagId',
    },
    {
      title: '标签名称',
      field: 'name',
    },
    {
      title: '粉丝数',
      field: 'count',
    },
    {
      title: '创建时间',
      field: 'createdAt',
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 140,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '标签',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'edit',
            show: hasAccessByCodes(['gzh:tag:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['gzh:tag:delete']),
          },
        ],
      },
    },
  ];
}
