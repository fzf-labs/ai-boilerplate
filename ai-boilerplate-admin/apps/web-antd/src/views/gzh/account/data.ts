import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { WxGzhAccountInfo } from '#/api/v1/wx-gzh-account';

import { useAccess } from '@vben/access';

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
      fieldName: 'name',
      label: '名称',
      component: 'Input',
      rules: 'required',
      componentProps: {
        placeholder: '请输入名称',
      },
    },
    {
      fieldName: 'account',
      label: '微信号',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 账号详情] 中能找到「微信号」',
      rules: 'required',
      componentProps: {
        placeholder: '请输入微信号',
      },
    },
    {
      fieldName: 'appId',
      label: '开发者ID',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 开发接口管理] 中能找到「开发者ID(AppID)」',
      rules: 'required',
      componentProps: {
        placeholder: '请输入appId',
      },
    },
    {
      fieldName: 'appSecret',
      label: '开发者密码',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 开发接口管理] 中能找到「开发者密码(AppSecret)」',
      rules: 'required',
      componentProps: {
        placeholder: '请输入appSecret',
      },
    },
    {
      fieldName: 'URL',
      label: '服务器地址',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 开发接口管理] 中能找到「服务器地址(URL)」',
      componentProps: {
        placeholder: '请输入URL',
      },
    },
    {
      fieldName: 'token',
      label: '令牌',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 开发接口管理] 中能找到「令牌(Token)」',
      componentProps: {
        placeholder: '请输入token',
      },
    },
    {
      fieldName: 'encodingAesKey',
      label: '消息加解密密钥',
      component: 'Input',
      help: '在微信公众平台（mp.weixin.qq.com）的菜单 [设置与开发 - 公众号设置 - 开发接口管理] 中能找到「消息加解密密钥(EncodingAESKey)」',
      componentProps: {
        placeholder: '请输入encodingAesKey',
      },
    },
    {
      fieldName: 'remark',
      label: '备注',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注',
      },
    },
  ];
}

/** 搜索表单配置 */
export function useGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'account',
      label: '微信号',
      component: 'Input',
    },
    {
      fieldName: 'appId',
      label: '开发者ID',
      component: 'Input',
    },
  ];
}

/** 表格列配置 */
export function useGridColumns<T = WxGzhAccountInfo>(
  onActionClick: OnActionClickFn<T>,
): VxeTableGridOptions['columns'] {
  return [
    {
      title: '名称',
      field: 'name',
    },
    {
      title: '微信号',
      field: 'account',
    },
    {
      title: '开发者ID',
      field: 'appId',
    },
    {
      title: '二维码',
      field: 'qrCodeURL',
      width: 120,
      align: 'center',
      cellRender: {
        name: 'CellImage',
        props: ({ row }: { row: WxGzhAccountInfo }) => ({
          src: row.qrCodeURL,
          width: 60,
          height: 60,
          fallback: undefined,
        }),
      },
    },
    {
      title: '备注',
      field: 'remark',
    },
    {
      field: 'operation',
      title: '操作',
      minWidth: 180,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'name',
          nameTitle: '公众号账号',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'detail',
            text: '详情',
            show: hasAccessByCodes(['gzh:account:query']),
          },
          {
            code: 'edit',
            text: '编辑',
            show: hasAccessByCodes(['gzh:account:update']),
          },
          {
            code: 'generateQrCode',
            text: '生成二维码',
            show: hasAccessByCodes(['gzh:account:qr-code']),
          },
          {
            code: 'clearQuota',
            text: '清空 API 配额',
            show: hasAccessByCodes(['gzh:account:clear-quota']),
          },
          {
            code: 'delete',
            text: '删除',
            show: hasAccessByCodes(['gzh:account:delete']),
          },
        ],
      },
    },
  ];
}
