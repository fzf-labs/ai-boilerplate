import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SelfAppReleaseInfo } from '#/api/v1/self-app-release';

import { useAccess } from '@vben/access';

import dayjs from 'dayjs';

import { z } from '#/adapter/form';
import { parsePackageFile } from '#/utils/appInfoParser';
import { CommonStatusEnum } from '#/utils/constants';

const { hasAccessByCodes } = useAccess();

/** 版本发布新增/修改的表单 */
export function useReleaseFormSchema(
  onParsePackage?: (
    buildNum: number,
    version: string,
    packageSize?: number,
    packageMd5?: string,
    minOsVersion?: string,
  ) => void,
): VbenFormSchema[] {
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
      fieldName: 'packageName',
      label: '包名',
      component: 'Input',
      componentProps: {
        placeholder: '应用包名',
        disabled: true,
      },
      rules: 'required',
    },
    {
      fieldName: 'channel',
      label: '发布渠道',
      component: 'Select',
      componentProps: {
        placeholder: '请输入发布渠道',
        allowClear: true,
        options: [
          { label: '官方', value: 'official' },
          { label: '应用宝', value: 'yingyongbao' },
          { label: '小米', value: 'xiaomi' },
          { label: '华为', value: 'huawei' },
          { label: 'OPPO', value: 'oppo' },
          { label: 'VIVO', value: 'vivo' },
          { label: '魅族', value: 'meizu' },
          { label: '百度', value: 'baidu' },
          { label: '360', value: '360' },
        ],
      },
      rules: 'required',
    },
    {
      fieldName: 'packageURL',
      label: '安装包地址',
      component: 'FileUpload',
      componentProps: {
        placeholder: '请选择安装包',
        maxCount: 1,
        maxSize: 1024,
        scene: 'selfapp',
        onFileSelected: async (file: File) => {
          try {
            const { buildNum, version, packageSize, packageMd5, minOsVersion } =
              await parsePackageFile(file);

            if (buildNum && version && onParsePackage) {
              onParsePackage(
                buildNum,
                version,
                packageSize,
                packageMd5,
                minOsVersion,
              );
            } else {
              console.warn('未能从安装包中解析到版本信息');
            }
          } catch (error: any) {
            console.error('解析安装包失败:', error);
            // 可以在这里显示更友好的错误提示
            // message.error('解析安装包失败，请手动输入版本信息');
          }
        },
      },
      rules: 'required',
    },
    {
      fieldName: 'buildNum',
      label: 'Build号',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入Build号',
        min: 1,
        style: { width: '100%' },
      },
      rules: z.number().min(1, 'Build号必须大于0'),
    },
    {
      fieldName: 'version',
      label: '版本号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入版本号，如：1.0.0',
      },
    },
    {
      fieldName: 'packageSize',
      label: '安装包大小',
      component: 'InputNumber',
      componentProps: {
        placeholder: '请输入安装包大小(MB)',
        min: 0,
        precision: 2,
        style: { width: '100%' },
      },
    },
    {
      fieldName: 'packageMd5',
      label: '安装包MD5',
      component: 'Input',
      componentProps: {
        placeholder: '请输入安装包MD5值',
      },
    },
    {
      fieldName: 'minOsVersion',
      label: '最低系统版本',
      component: 'Input',
      componentProps: {
        placeholder: '请输入最低系统版本要求',
      },
    },
    {
      fieldName: 'publishTime',
      label: '发布时间',
      component: 'DatePicker',
      componentProps: {
        placeholder: '请选择发布时间',
        showTime: { defaultValue: dayjs('23:59:59', 'HH:mm:ss') },
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DDTHH:mm:ssZ',
        style: { width: '100%' },
      },
      rules: 'required',
    },
    {
      fieldName: 'updateType',
      label: '更新类型',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '强制更新', value: 1 },
          { label: '提示更新', value: 2 },
          { label: '静默更新', value: 3 },
        ],
        optionType: 'button',
        buttonStyle: 'solid',
      },
      rules: z.number().default(2),
    },
    {
      fieldName: 'title',
      label: '更新标题',
      component: 'Input',
      componentProps: {
        placeholder: '请输入更新标题',
      },
      rules: 'required',
    },
    {
      fieldName: 'changelog',
      label: '更新日志',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入更新日志内容',
        rows: 4,
      },
      defaultValue: '解决了用户提出的问题\n解决了提出问题的用户',
    },
    {
      fieldName: 'grayStrategy',
      label: '灰度策略',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '全量发布', value: 1 },
          { label: '自定义设备', value: 2 },
        ],
        optionType: 'button',
        buttonStyle: 'solid',
      },
      rules: z.number().default(1),
    },
    {
      fieldName: 'graySns',
      label: '灰度设备列表',
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入设备序列号，每行一个',
        rows: 4,
      },
      dependencies: {
        triggerFields: ['grayStrategy'],
        show: (values) => values.grayStrategy === 2,
      },
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
        optionType: 'button',
        buttonStyle: 'solid',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
  ];
}

/** 版本发布列表的搜索表单 */
export function useReleaseGridFormSchema(): VbenFormSchema[] {
  return [
    {
      fieldName: 'channel',
      label: '发布渠道',
      component: 'Select',
      componentProps: {
        placeholder: '请输入发布渠道',
        allowClear: true,
        options: [
          { label: '官方', value: 'official' },
          { label: '应用宝', value: 'yingyongbao' },
          { label: '小米', value: 'xiaomi' },
          { label: '华为', value: 'huawei' },
          { label: 'OPPO', value: 'oppo' },
          { label: 'VIVO', value: 'vivo' },
          { label: '魅族', value: 'meizu' },
          { label: '百度', value: 'baidu' },
          { label: '360', value: '360' },
        ],
      },
    },
    {
      fieldName: 'buildNum',
      label: 'Build号',
      component: 'Input',
      componentProps: {
        placeholder: '请输入Build号',
        allowClear: true,
      },
    },
  ];
}

/** 版本发布列表的字段 */
export function useReleaseGridColumns<T = SelfAppReleaseInfo>(
  onActionClick: OnActionClickFn<T>,
  onStatusChange?: (
    newStatus: number,
    row: T,
  ) => PromiseLike<boolean | undefined>,
): VxeTableGridOptions['columns'] {
  return [
    {
      field: 'channel',
      title: '发布渠道',
      minWidth: 120,
      formatter: ({ cellValue }) => {
        const channelMap = {
          official: '官方',
          yingyongbao: '应用宝',
          xiaomi: '小米',
          huawei: '华为',
          oppo: 'OPPO',
          vivo: 'VIVO',
          meizu: '魅族',
          baidu: '百度',
          '360': '360',
        };
        return channelMap[cellValue as keyof typeof channelMap] || '-';
      },
    },
    {
      field: 'version',
      title: '版本号',
      minWidth: 100,
      formatter: ({ cellValue }) => cellValue || '-',
    },
    {
      field: 'buildNum',
      title: 'Build号',
      minWidth: 100,
      align: 'center',
    },
    {
      field: 'updateType',
      title: '更新类型',
      minWidth: 100,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: SelfAppReleaseInfo }) => {
          const typeMap = {
            1: { color: 'red', text: '强制' },
            2: { color: 'orange', text: '提示' },
            3: { color: 'blue', text: '静默' },
          };
          const type = typeMap[row.updateType as keyof typeof typeMap];
          return {
            color: type?.color || 'default',
            text: type?.text || '未知',
          };
        },
      },
    },
    {
      field: 'title',
      title: '更新标题',
      minWidth: 200,
    },
    {
      field: 'grayStrategy',
      title: '灰度策略',
      minWidth: 120,
      align: 'center',
      cellRender: {
        name: 'CellTag',
        props: ({ row }: { row: SelfAppReleaseInfo }) => {
          const strategyMap = {
            1: { color: 'green', text: '全量' },
            2: { color: 'blue', text: '自定义' },
          };
          const strategy =
            strategyMap[row.grayStrategy as keyof typeof strategyMap];
          return {
            color: strategy?.color || 'default',
            text: strategy?.text || '未知',
          };
        },
      },
    },
    {
      field: 'status',
      title: '状态',
      minWidth: 100,
      align: 'center',
      cellRender: {
        attrs: { beforeChange: onStatusChange },
        name: 'CellSwitch',
        props: {
          checkedValue: 1,
          unCheckedValue: -1,
        },
      },
    },
    {
      field: 'publishTime',
      title: '发布时间',
      minWidth: 180,
      formatter: 'formatDateTime',
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
      minWidth: 180,
      fixed: 'right',
      align: 'center',
      cellRender: {
        attrs: {
          nameField: 'title',
          nameTitle: '版本',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'view',
            text: '详情',
            show: hasAccessByCodes(['self_app_release:query']),
          },
          {
            code: 'edit',
            show: hasAccessByCodes(['self_app_release:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['self_app_release:delete']),
          },
        ],
      },
    },
  ];
}
