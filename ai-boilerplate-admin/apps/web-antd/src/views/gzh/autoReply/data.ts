import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { WxGzhAutoReplyInfo } from '#/api/v1/wx-gzh-auto-reply';

import { useAccess } from '@vben/access';

import { getWxGzhAccountSelector } from '#/api/v1/wx-gzh-account';
import { CommonStatusEnum } from '#/utils/constants';

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
      fieldName: 'type',
      label: '回复类型',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '关键词回复', value: 1 },
          {
            label: '收到消息回复',
            value: 2,
          },
          {
            label: '被关注回复',
            value: 3,
          },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      dependencies: {
        triggerFields: ['id'],
        componentProps: (values) => ({
          disabled: !!values.id, // 有id表示编辑模式，禁用类型选择
        }),
      },
      rules: 'required',
      defaultValue: 1,
      help: '提示：被关注回复和收到消息回复每个公众号只能设置一个，关键词回复可以设置多个。编辑时不能修改回复类型。',
    },
    {
      fieldName: 'requestKeyword',
      label: '请求关键字',
      component: 'Input',
      componentProps: {
        placeholder: '请输入关键字',
      },
      help: '用户发送此关键字时触发自动回复',
      dependencies: {
        triggerFields: ['type'],
        show: (values) => values.type === 1,
      },
      rules: 'required',
    },
    {
      fieldName: 'requestKeywordMatch',
      label: '关键字匹配类型',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '全匹配', value: 1 },
          { label: '半匹配', value: 2 },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      dependencies: {
        triggerFields: ['type'],
        show: (values) => values.type === 1,
      },
      rules: 'required',
      defaultValue: 1,
    },
    {
      fieldName: 'responseMessageType',
      label: '回复消息类型',
      component: 'Select',
      componentProps: {
        placeholder: '请选择回复消息类型',
        options: [
          { label: '文本消息', value: 'text' },
          {
            label: '图片消息',
            value: 'image',
          },
          {
            label: '音频消息',
            value: 'voice',
          },
          {
            label: '视频消息',
            value: 'video',
          },
        ],
      },
      rules: 'required',
      defaultValue: 'text',
    },
    {
      fieldName: 'responseContent',
      label: '回复内容',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入回复内容',
        rows: 4,
        maxlength: 500,
        showCount: true,
      },
      dependencies: {
        triggerFields: ['responseMessageType'],
        show: (values) => values.responseMessageType === 'text',
      },
      rules: 'required',
    },
    {
      fieldName: 'responseMediaId',
      label: '媒体文件ID',
      component: 'Input',
      componentProps: {
        placeholder: '请输入媒体文件ID',
      },
      help: '图片、语音、视频、音乐、图文消息需要提供媒体文件ID',
      dependencies: {
        triggerFields: ['responseMessageType'],
        show: (values) => values.responseMessageType !== 'text',
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: 'required',
      defaultValue: CommonStatusEnum.ENABLE,
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
    {
      fieldName: 'type',
      label: '回复类型',
      component: 'Select',
      componentProps: {
        placeholder: '请选择回复类型',
        options: [
          { label: '关键词回复', value: 1 },
          {
            label: '收到消息回复',
            value: 2,
          },
          {
            label: '被关注回复',
            value: 3,
          },
        ],
      },
    },
  ];
}

/** 表格列配置 */
export function useGridColumns<T = WxGzhAutoReplyInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      title: '回复类型',
      field: 'type',
      width: 120,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 1: {
            return '关键词回复';
          }
          case 2: {
            return '收到消息回复';
          }
          case 3: {
            return '被关注回复';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      title: '请求关键字',
      field: 'requestKeyword',
      width: 150,
      formatter: ({ cellValue, row }) => {
        if (row.type === 1) {
          return cellValue || '-';
        }
        return '-';
      },
    },
    {
      title: '匹配类型',
      field: 'requestKeywordMatch',
      width: 100,
      formatter: ({ cellValue, row }) => {
        if (row.type === 1) {
          switch (cellValue) {
            case 1: {
              return '全匹配';
            }
            case 2: {
              return '半匹配';
            }
            default: {
              return '-';
            }
          }
        }
        return '-';
      },
    },
    {
      title: '消息类型',
      field: 'responseMessageType',
      width: 120,
      formatter: ({ cellValue }) => {
        switch (cellValue) {
          case 'image': {
            return '图片消息';
          }
          case 'text': {
            return '文本消息';
          }
          case 'video': {
            return '视频消息';
          }
          case 'voice': {
            return '音频消息';
          }
          default: {
            return cellValue || '-';
          }
        }
      },
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      formatter: (row) => {
        return row.cellValue === 1 ? '启用' : '禁用';
      },
    },
    {
      title: '创建时间',
      field: 'createdAt',
      width: 180,
      formatter: 'formatDateTime',
    },
    {
      title: '更新时间',
      field: 'updatedAt',
      width: 180,
      formatter: 'formatDateTime',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 180,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'id',
          nameTitle: '自动回复',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '查看',
            show: hasAccessByCodes(['gzh:auto-reply:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['gzh:auto-reply:update']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['gzh:auto-reply:delete']),
          },
        ],
      },
    },
  ];
}
