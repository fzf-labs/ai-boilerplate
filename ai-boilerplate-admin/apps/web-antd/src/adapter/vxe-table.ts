import type { VxeTableGridOptions } from '@vben/plugins/vxe-table';
import type { Recordable } from '@vben/types';

import { h } from 'vue';

import { Eye, EyeOff, IconifyIcon } from '@vben/icons';
import { $te } from '@vben/locales';
import { setupVbenVxeTable, useVbenVxeGrid } from '@vben/plugins/vxe-table';
import { isFunction, isString } from '@vben/utils';

import {
  Button,
  Dropdown,
  Image,
  ImagePreviewGroup,
  Menu,
  MenuItem,
  Popconfirm,
  Switch,
  Tag,
} from 'ant-design-vue';

import { DictTag } from '#/components/dict-tag';
import { $t } from '#/locales';

import { useVbenForm } from './form';

import '#/adapter/style.css';

setupVbenVxeTable({
  configVxeTable: (vxeUI) => {
    vxeUI.setConfig({
      grid: {
        align: 'center',
        border: false,
        columnConfig: {
          resizable: true,
        },
        minHeight: 180,
        formConfig: {
          // 全局禁用vxe-table的表单配置，使用formOptions
          enabled: false,
        },
        toolbarConfig: {
          import: false, // 是否导入
          export: false, // 是否导出
          refresh: true, // 是否刷新
          print: false, // 是否打印
          zoom: true, // 是否缩放
          custom: true, // 是否自定义配置
        },
        customConfig: {
          mode: 'modal',
        },
        proxyConfig: {
          autoLoad: true,
          response: {
            list: 'list',
            result: 'list',
            total: 'total',
          },
          showActiveMsg: true,
          showResponseMsg: false,
        },
        pagerConfig: {
          enabled: true,
        },
        sortConfig: {
          multiple: true,
        },
        round: true,
        showOverflow: true,
        size: 'small',
      } as VxeTableGridOptions,
    });

    // 表格配置项可以用 cellRender: { name: 'CellImage' },
    vxeUI.renderer.add('CellImage', {
      renderTableDefault(_renderOpts, params) {
        const { column, row } = params;
        return h(Image, { src: row[column.field] });
      },
    });

    vxeUI.renderer.add('CellImages', {
      renderTableDefault(_renderOpts, params) {
        const { column, row } = params;
        if (column && column.field && row[column.field]) {
          return h(ImagePreviewGroup, {}, () => {
            return row[column.field].map((item: any) =>
              h(Image, { src: item }),
            );
          });
        }
        return '';
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellLink' },
    vxeUI.renderer.add('CellLink', {
      renderTableDefault(renderOpts, params) {
        const { props: propsFn } = renderOpts;
        const { column, row } = params;
        // 如果 props 是函数，则调用它获取实际的 props
        const props =
          typeof propsFn === 'function' ? propsFn({ row, column }) : propsFn;

        if (!props?.href && !props?.text) {
          return props?.text || '-';
        }

        const buttonProps: any = {
          size: 'small',
          type: 'link',
          onClick: props?.href
            ? () => {
                if (props.download) {
                  // 创建临时下载链接
                  const link = document.createElement('a');
                  link.href = props.href;
                  link.download = props.download;
                  document.body.append(link);
                  link.click();
                  link.remove();
                } else {
                  window.open(props.href, '_blank');
                }
              }
            : undefined,
        };

        return h(Button, buttonProps, {
          default: () => props?.text || props?.href,
        });
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellTag' },
    vxeUI.renderer.add('CellTag', {
      renderTableDefault(renderOpts, params) {
        const { props: propsFn } = renderOpts;
        const { column, row } = params;
        // 如果 props 是函数，则调用它获取实际的 props
        const props =
          typeof propsFn === 'function' ? propsFn({ row, column }) : propsFn;
        return h(Tag, { color: props?.color }, () => props?.text);
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellTags' },
    vxeUI.renderer.add('CellTags', {
      renderTableDefault(renderOpts, params) {
        const { props: propsFn } = renderOpts;
        const { column, row } = params;
        // 如果 props 是函数，则调用它获取实际的 props
        const props =
          typeof propsFn === 'function' ? propsFn({ row, column }) : propsFn;

        const tags = props?.tags || row[column.field];
        if (!tags || !Array.isArray(tags) || tags.length === 0) {
          return '';
        }

        return h(
          'div',
          {
            class: 'flex items-center gap-1',
            style: 'flex-wrap: nowrap; overflow-x: auto;',
          },
          tags.map((tag: any) =>
            h(
              Tag,
              {
                key: tag.id || tag.text || tag,
                color: tag.color || props?.color || 'default',
                size: 'small',
              },
              () => tag.text || tag.name || tag,
            ),
          ),
        );
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellDict', props:{type: 'dict_type'} } 或 props: { options: [{label: '', value: ''}] }
    vxeUI.renderer.add('CellDict', {
      renderTableDefault(renderOpts, params) {
        const { props } = renderOpts;
        const { column, row } = params;
        if (!props) {
          return '';
        }

        const cellValue = row[column.field];

        // 如果提供了 options，直接使用 options 渲染
        if (props.options && Array.isArray(props.options)) {
          const option = props.options.find(
            (opt: any) => opt.value === cellValue,
          );
          if (!option) {
            return '';
          }
          return h(
            Tag,
            {
              color: option.color || 'default',
            },
            () => option.label,
          );
        }

        // 如果提供了 type，使用 DictTag 组件
        if (props.type) {
          return h(DictTag, {
            type: props.type,
            value: cellValue?.toString(),
          });
        }

        return '';
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellPassword' },
    vxeUI.renderer.add('CellPassword', {
      renderTableDefault(_renderOpts, params) {
        const { column, row } = params;
        const value = row[column.field];

        if (!value) {
          return '';
        }

        // 为每个单元格创建独立的显示状态
        const rowId = row.id || row._X_ROW_KEY;
        const stateKey = `__show_password_${column.field}_${rowId}`;

        // 初始化状态
        if (row[stateKey] === undefined) {
          row[stateKey] = false;
        }

        const toggleShow = () => {
          row[stateKey] = !row[stateKey];
        };

        return h(
          'div',
          {
            class: 'flex items-center gap-2',
          },
          [
            h(
              'span',
              {
                class: 'flex-1 truncate',
              },
              row[stateKey] ? value : '********',
            ),
            h(
              Button,
              {
                type: 'link',
                size: 'small',
                class: 'flex items-center justify-center p-0 min-w-[24px]',
                onClick: toggleShow,
              },
              {
                icon: () =>
                  h(row[stateKey] ? Eye : EyeOff, { class: 'size-4' }),
              },
            ),
          ],
        );
      },
    });

    // 表格配置项可以用 cellRender: { name: 'CellSwitch', props: { beforeChange: () => {} } },
    vxeUI.renderer.add('CellSwitch', {
      renderTableDefault({ attrs, props }, { column, row }) {
        const loadingKey = `__loading_${column.field}`;
        const finallyProps = {
          checkedChildren: $t('common.enabled'),
          checkedValue: 1,
          unCheckedChildren: $t('common.disabled'),
          unCheckedValue: 0,
          ...props,
          checked: row[column.field],
          loading: row[loadingKey] ?? false,
          'onUpdate:checked': onChange,
        };

        async function onChange(newVal: any) {
          row[loadingKey] = true;
          try {
            const result = await attrs?.beforeChange?.(newVal, row);
            if (result !== false) {
              row[column.field] = newVal;
            }
          } finally {
            row[loadingKey] = false;
          }
        }

        return h(Switch, finallyProps);
      },
    });

    // 注册表格的操作按钮渲染器 cellRender: { name: 'CellOperation', options: ['edit', 'delete'] }
    vxeUI.renderer.add('CellOperation', {
      renderTableDefault({ attrs, options, props }, { column, row }) {
        const defaultProps = { size: 'small', type: 'link', ...props };
        let align = 'end';
        switch (column.align) {
          case 'center': {
            align = 'center';
            break;
          }
          case 'left': {
            align = 'start';
            break;
          }
          default: {
            align = 'end';
            break;
          }
        }
        const presets: Recordable<Recordable<any>> = {
          delete: {
            danger: true,
            text: $t('common.delete'),
          },
          edit: {
            text: $t('common.edit'),
          },
        };

        // 处理操作配置，支持嵌套的 children 配置
        const processOperations = (opts: any[]): Array<Recordable<any>> => {
          return opts
            .map((opt) => {
              if (isString(opt)) {
                return presets[opt]
                  ? { code: opt, ...presets[opt], ...defaultProps }
                  : {
                      code: opt,
                      text: $te(`common.${opt}`) ? $t(`common.${opt}`) : opt,
                      ...defaultProps,
                    };
              } else {
                const processedOpt = {
                  ...defaultProps,
                  ...presets[opt.code],
                  ...opt,
                };
                // 如果有 children，递归处理
                if (processedOpt.children) {
                  processedOpt.children = processOperations(
                    processedOpt.children,
                  );
                }
                return processedOpt;
              }
            })
            .map((opt) => {
              const optBtn: Recordable<any> = {};
              Object.keys(opt).forEach((key) => {
                if (key === 'children' && opt[key]) {
                  // children 需要特殊处理，递归评估每个子项
                  optBtn[key] = opt[key]
                    .map((child: any) => {
                      const childBtn: Recordable<any> = {};
                      Object.keys(child).forEach((childKey) => {
                        childBtn[childKey] = isFunction(child[childKey])
                          ? child[childKey](row)
                          : child[childKey];
                      });
                      return childBtn;
                    })
                    .filter((child: any) => child.show !== false);
                } else {
                  optBtn[key] = isFunction(opt[key]) ? opt[key](row) : opt[key];
                }
              });
              return optBtn;
            })
            .filter((opt) => opt.show !== false);
        };

        const operations = processOperations(options || ['edit', 'delete']);

        function renderBtn(opt: Recordable<any>, listen = true) {
          // 只传递Button组件需要的属性，过滤掉内部使用的属性
          const {
            code: _code,
            text: _text,
            show: _show,
            children: _children,
            icon: _icon,
            ...buttonProps
          } = opt;

          return h(
            Button,
            {
              ...props,
              ...buttonProps,
              icon: undefined, // 图标通过children渲染
              onClick: listen
                ? () =>
                    attrs?.onClick?.({
                      code: opt.code,
                      row,
                    })
                : undefined,
            },
            {
              default: () => {
                if (opt.icon) {
                  return [
                    h(IconifyIcon, { class: 'size-5', icon: opt.icon }),
                    opt.text || '',
                  ];
                }
                return opt.text || '';
              },
            },
          );
        }

        function renderConfirm(opt: Recordable<any>) {
          return h(
            Popconfirm,
            {
              getPopupContainer(el) {
                return el.closest('tbody') || document.body;
              },
              placement: 'topLeft',
              title: $t('ui.actionTitle.delete', [attrs?.nameTitle || '']),
              ...props,
              ...opt,
              icon: undefined,
              onConfirm: () => {
                attrs?.onClick?.({
                  code: opt.code,
                  row,
                });
              },
            },
            {
              default: () => renderBtn({ ...opt }, false),
              description: () =>
                h(
                  'div',
                  { class: 'truncate' },
                  $t('ui.actionMessage.deleteConfirm', [
                    row[attrs?.nameField || 'name'],
                  ]),
                ),
            },
          );
        }

        // 渲染下拉菜单
        function renderDropdown(opt: Recordable<any>) {
          const children = opt.children || [];
          if (children.length === 0) {
            return renderBtn(opt);
          }

          const menuItems = children.map((child: Recordable<any>) => {
            const handleClick = () => {
              if (child.code === 'delete') {
                // 删除操作需要确认
                const confirmTitle = $t('ui.actionTitle.delete', [
                  attrs?.nameTitle || '',
                ]);
                const confirmContent = $t('ui.actionMessage.deleteConfirm', [
                  row[attrs?.nameField || 'name'],
                ]);

                // 创建确认对话框
                import('ant-design-vue').then((AntdVue) => {
                  AntdVue.Modal.confirm({
                    title: confirmTitle,
                    content: confirmContent,
                    onOk() {
                      attrs?.onClick?.({
                        code: child.code,
                        row,
                      });
                    },
                  });
                });
              } else {
                attrs?.onClick?.({
                  code: child.code,
                  row,
                });
              }
            };

            return h(
              MenuItem,
              {
                key: child.code,
                danger: child.danger,
                onClick: handleClick,
              },
              {
                default: () => {
                  if (child.icon) {
                    return [
                      h(IconifyIcon, {
                        class: 'mr-2 size-4',
                        icon: child.icon,
                      }),
                      child.text || '',
                    ];
                  }
                  return child.text || '';
                },
              },
            );
          });

          const menu = h(
            Menu,
            {},
            {
              default: () => menuItems,
            },
          );

          return h(
            Dropdown,
            {
              placement: 'bottomLeft',
              getPopupContainer: (el: Element) =>
                el.closest('tbody') || document.body,
            },
            {
              default: () => renderBtn(opt, false),
              overlay: () => menu,
            },
          );
        }

        // 根据配置渲染按钮或下拉菜单
        const btns = operations.map((opt) => {
          if (opt.children && opt.children.length > 0) {
            return renderDropdown(opt);
          } else if (opt.code === 'delete') {
            return renderConfirm(opt);
          } else {
            return renderBtn(opt);
          }
        });

        return h(
          'div',
          {
            class: 'flex table-operations',
            style: { justifyContent: align },
          },
          {
            default: () => btns,
          },
        );
      },
    });
    // 这里可以自行扩展 vxe-table 的全局配置，比如自定义格式化
    // vxeUI.formats.add

    // 金额格式化
    vxeUI.formats.add('formatAmount', {
      cellFormatMethod({ cellValue }, digits = 2) {
        if (cellValue === null || cellValue === undefined) {
          return '';
        }
        if (isString(cellValue)) {
          cellValue = Number.parseFloat(cellValue);
        }
        // 如果非 number，则直接返回空串
        if (Number.isNaN(cellValue)) {
          return '';
        }
        return cellValue.toFixed(digits);
      },
    });
  },
  useVbenForm,
});

export { useVbenVxeGrid };

export type OnActionClickParams<T = Recordable<any>> = {
  code: string;
  row: T;
};

export type OnActionClickFn<T = Recordable<any>> = (
  params: OnActionClickParams<T>,
) => void;
export * from '#/components/table-action';
export type * from '@vben/plugins/vxe-table';
