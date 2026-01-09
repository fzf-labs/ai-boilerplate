import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { MailLogInfo } from '#/api/v1/mail-log';

import { useAccess } from '@vben/access';

import { getMailAccountSelector } from '#/api/v1/mail-account';
import { getMailTemplateSelector } from '#/api/v1/mail-template';
import { getRangePickerDefaultProps } from '#/utils';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 列表的搜索表单 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'accountId',
      label: '邮箱账号',
      component: 'ApiSelect',
      componentProps: {
        api: async () => await getMailAccountSelector({ options: {} }),
        resultField: 'list',
        labelField: 'mail',
        valueField: 'id',
        allowClear: true,
        placeholder: '请选择邮箱账号',
      },
    },
    {
      fieldName: 'templateId',
      label: '模板编号',
      component: 'ApiSelect',
      dependencies: {
        triggerFields: ['accountId'],
        show: (formValues) => !!formValues.accountId,
        trigger: (_, formActions) => {
          formActions.setFieldValue('templateId', ''); // 清空当前选中的模板
        },
        componentProps: (formValues) => ({
          resultField: 'list',
          labelField: 'name',
          valueField: 'id',
          allowClear: true,
          placeholder: '请选择模板编号',
          api: async (params: Record<string, any>) => {
            const accountId = params?.accountId || '';
            return await getMailTemplateSelector({ params: { accountId } });
          },
          params: { accountId: formValues.accountId },
        }),
      },
    },
    {
      fieldName: 'fromMail',
      label: '发送邮箱',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入发送邮箱',
      },
    },
    {
      fieldName: 'toMail',
      label: '收件邮箱',
      component: 'Input',
      componentProps: {
        allowClear: true,
        placeholder: '请输入收件邮箱',
      },
    },
    {
      fieldName: 'sendTime',
      label: '发送时间',
      component: 'RangePicker',
      componentProps: {
        ...getRangePickerDefaultProps(),
        allowClear: true,
      },
    },
    {
      fieldName: 'sendStatus',
      label: '发送状态',
      component: 'Select',
      componentProps: {
        options: [
          {
            label: '成功',
            value: 1,
          },
          {
            label: '失败',
            value: -1,
          },
        ],
        allowClear: true,
        placeholder: '请选择发送状态',
      },
    },
  ];
}

/** 列表的字段 */
export function useGridColumns<T = MailLogInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'fromMail',
      title: '发送邮箱',
      minWidth: 120,
    },
    {
      field: 'toMail',
      title: '收件邮箱',
      minWidth: 160,
    },
    {
      field: 'templateId',
      title: '模板编号',
      minWidth: 120,
    },
    {
      field: 'templateCode',
      title: '模板编码',
      minWidth: 120,
    },
    {
      field: 'templateTitle',
      title: '邮件标题',
      minWidth: 120,
    },
    {
      field: 'sendTime',
      title: '发送时间',
      minWidth: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'sendStatus',
      title: '发送状态',
      minWidth: 120,
      formatter: (row) => {
        switch (row.cellValue) {
          case CommonStatusEnum.DISABLE: {
            return '禁用';
          }
          case CommonStatusEnum.ENABLE: {
            return '启用';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 80,
      align: 'center',
      fixed: 'right',
      cellRender: {
        attrs: {
          nameField: 'toMail',
          nameTitle: '邮件日志',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '查看',
            show: hasAccessByCodes(['infra:mail-log:query']),
          },
        ],
      },
    },
  ];
}
