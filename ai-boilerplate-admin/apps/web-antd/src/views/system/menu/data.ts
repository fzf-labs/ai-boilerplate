import type { Recordable } from '@vben/types';

import type { VbenFormSchema } from '#/adapter/form';
import type { OnActionClickFn, VxeTableGridOptions } from '#/adapter/vxe-table';
import type { SysMenuInfo } from '#/api/v1/sys-menu';

import { h } from 'vue';

import { useAccess } from '@vben/access';
import { IconifyIcon } from '@vben/icons';
import { handleTree, isHttpUrl } from '@vben/utils';

import { z } from '#/adapter/form';
import { getSysMenuList } from '#/api/v1/sys-menu';
import { $t } from '#/locales';
import { componentKeys } from '#/router/routes';
import { CommonStatusEnum, SystemMenuTypeEnum } from '#/utils/constants';

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
      fieldName: 'pid',
      label: '上级菜单',
      component: 'ApiTreeSelect',
      componentProps: {
        allowClear: true,
        api: async () => {
          const data = await getSysMenuList({ options: {} });
          data.list?.unshift({
            id: '',
            name: '顶级菜单',
          } as SysMenuInfo);
          return handleTree(data.list || []);
        },
        class: 'w-full',
        labelField: 'name',
        valueField: 'id',
        childrenField: 'children',
        placeholder: '请选择上级菜单',
        filterTreeNode(input: string, node: Recordable<any>) {
          if (!input || input.length === 0) {
            return true;
          }
          const name: string = node.label ?? '';
          if (!name) return false;
          return name.includes(input) || $t(name).includes(input);
        },
        showSearch: true,
        treeDefaultExpandedKeys: [0],
      },
      rules: 'selectRequired',
      renderComponentContent() {
        return {
          title({ label, icon }: { icon: string; label: string }) {
            const components = [];
            if (!label) return '';
            if (icon) {
              components.push(h(IconifyIcon, { class: 'size-4', icon }));
            }
            components.push(h('span', { class: '' }, $t(label || '')));
            return h('div', { class: 'flex items-center gap-1' }, components);
          },
        };
      },
    },
    {
      fieldName: 'name',
      label: '菜单名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入菜单名称',
      },
      rules: 'required',
    },
    {
      fieldName: 'type',
      label: '菜单类型',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '目录', value: SystemMenuTypeEnum.DIR },
          { label: '菜单', value: SystemMenuTypeEnum.MENU },
          { label: '按钮', value: SystemMenuTypeEnum.BUTTON },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.string().default(SystemMenuTypeEnum.DIR),
    },
    {
      fieldName: 'icon',
      label: '菜单图标',
      component: 'IconPicker',
      componentProps: {
        placeholder: '请选择菜单图标',
        prefix: 'carbon',
      },
      rules: 'required',
      dependencies: {
        triggerFields: ['type'],
        show: (values) => {
          return [SystemMenuTypeEnum.DIR, SystemMenuTypeEnum.MENU].includes(
            values.type,
          );
        },
      },
    },
    {
      fieldName: 'path',
      label: '路由地址',
      component: 'Input',
      componentProps: {
        placeholder: '请输入路由地址',
      },
      rules: z.string(),
      help: '访问的路由地址，如：`user`。如需外网地址时，则以 `http(s)://` 开头',
      dependencies: {
        triggerFields: ['type', 'parentId'],
        show: (values) => {
          return [SystemMenuTypeEnum.DIR, SystemMenuTypeEnum.MENU].includes(
            values.type,
          );
        },
        rules: (values) => {
          const schema = z.string().min(1, '路由地址不能为空');
          if (isHttpUrl(values.path)) {
            return schema;
          }
          if (values.parentId === 0) {
            return schema.refine(
              (path) => path.charAt(0) === '/',
              '路径必须以 / 开头',
            );
          }
          return schema.refine(
            (path) => path.charAt(0) !== '/',
            '路径不能以 / 开头',
          );
        },
      },
    },
    {
      fieldName: 'component',
      label: '组件地址',
      component: 'Input',
      componentProps: {
        placeholder: '请输入组件地址',
      },
      dependencies: {
        triggerFields: ['type'],
        show: (values) => {
          return [SystemMenuTypeEnum.MENU].includes(values.type);
        },
      },
    },
    {
      fieldName: 'componentName',
      label: '组件名称',
      component: 'AutoComplete',
      componentProps: {
        allowClear: true,
        class: 'w-full',
        filterOption(input: string, option: { value: string }) {
          return option.value.toLowerCase().includes(input.toLowerCase());
        },
        placeholder: '请选择组件名称',
        options: componentKeys.map((v) => ({ value: v })),
      },
      dependencies: {
        triggerFields: ['type'],
        show: (values) => {
          return [SystemMenuTypeEnum.MENU].includes(values.type);
        },
      },
    },
    {
      fieldName: 'permission',
      label: '权限标识',
      component: 'Input',
      componentProps: {
        placeholder: '请输入菜单描述',
      },
      dependencies: {
        show: (values) => {
          return [SystemMenuTypeEnum.BUTTON, SystemMenuTypeEnum.MENU].includes(
            values.type,
          );
        },
        triggerFields: ['type'],
      },
    },
    {
      fieldName: 'sort',
      label: '显示顺序',
      component: 'InputNumber',
      componentProps: {
        min: 0,
        class: 'w-full',
        controlsPosition: 'right',
        placeholder: '请输入显示顺序',
      },
      rules: 'required',
    },
    {
      fieldName: 'status',
      label: '菜单状态',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '启用', value: CommonStatusEnum.ENABLE },
          { label: '禁用', value: CommonStatusEnum.DISABLE },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number().default(CommonStatusEnum.ENABLE),
    },
  ];
}

/** 列表的字段 */
export function useGridColumns(
  onActionClick: OnActionClickFn<SysMenuInfo>,
): VxeTableGridOptions<SysMenuInfo>['columns'] {
  return [
    {
      field: 'name',
      title: '菜单名称',
      minWidth: 250,
      align: 'left',
      fixed: 'left',
      slots: { default: 'name' },
      treeNode: true,
    },
    {
      field: 'type',
      title: '菜单类型',
      minWidth: 100,
      formatter: (row) => {
        switch (row.cellValue) {
          case SystemMenuTypeEnum.BUTTON: {
            return '按钮';
          }
          case SystemMenuTypeEnum.DIR: {
            return '目录';
          }
          case SystemMenuTypeEnum.MENU: {
            return '菜单';
          }
          default: {
            return '';
          }
        }
      },
    },
    {
      field: 'path',
      title: '组件路径',
      minWidth: 200,
    },
    {
      field: 'permission',
      title: '权限标识',
      minWidth: 200,
    },
    {
      field: 'icon',
      title: '图标',
      minWidth: 200,
    },
    {
      field: 'component',
      minWidth: 200,
      title: '组件地址',
    },
    {
      field: 'componentName',
      minWidth: 200,
      title: '组件名称',
    },
    {
      field: 'sort',
      title: '显示排序',
      minWidth: 100,
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
      field: 'operation',
      title: '操作',
      align: 'right',
      minWidth: 200,
      fixed: 'right',
      headerAlign: 'center',
      showOverflow: false,
      cellRender: {
        attrs: {
          nameField: 'name',
          onClick: onActionClick,
        },
        name: 'CellOperation',
        options: [
          {
            code: 'append',
            text: '新增下级',
            show: hasAccessByCodes(['system:menu:create']),
          },
          {
            code: 'edit',
            show: hasAccessByCodes(['system:menu:update']),
          },
          {
            code: 'delete',
            show: hasAccessByCodes(['system:menu:delete']),
          },
        ],
      },
    },
  ];
}
