# 基础 CRUD 页面开发

本示例展示如何创建一个标准的部门管理 CRUD 页面，包含列表、创建、编辑、删除功能。

## 1. API 接口定义

**文件路径**：`apps/web-antd/src/api/system/dept/index.ts`

```typescript
import type { PageReply } from '@vben/request';
import { requestClient } from '#/api/request';

export namespace SystemDeptApi {
  /** 部门信息 */
  export interface Dept {
    id: string;
    pid: string;
    name: string;
    adminId: string;
    sort: number;
    status: number;
    createdAt: Date;
    updatedAt: Date;
    children?: Dept[];
  }

  /** 查询参数 */
  export interface DeptQuery {
    name?: string;
    status?: number;
  }
}

interface DeptInfo {
  info: SystemDeptApi.Dept;
}

/** 部门列表 */
export function getDeptList(params?: SystemDeptApi.DeptQuery) {
  return requestClient.get<PageReply<SystemDeptApi.Dept>>(
    '/admin/v1/sys_dept/list',
    { params }
  );
}

/** 查询部门详情 */
export async function getDeptInfo(id: string) {
  return requestClient.get<DeptInfo>(`/admin/v1/sys_dept/info?id=${id}`);
}

/** 新增部门 */
export async function createDept(data: SystemDeptApi.Dept) {
  return requestClient.post('/admin/v1/sys_dept/create', data);
}

/** 修改部门 */
export async function updateDept(data: SystemDeptApi.Dept) {
  return requestClient.post('/admin/v1/sys_dept/update', data);
}

/** 删除部门 */
export async function deleteDept(id: string) {
  return requestClient.post('/admin/v1/sys_dept/delete', { id });
}

/** 更新状态 */
export async function updateDeptStatus(id: string, status: number) {
  return requestClient.post('/admin/v1/sys_dept/update_status', {
    id,
    status
  });
}
```

**关键点**：
- 使用 namespace 组织类型定义
- 标准的 CRUD 函数命名：`get{Name}List`、`get{Name}Info`、`create{Name}`、`update{Name}`、`delete{Name}`
- 接口返回类型明确

---

## 2. 表格列和表单配置

**文件路径**：`apps/web-antd/src/views/system/dept/data.ts`

```typescript
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { SystemDeptApi } from '#/api/system/dept';

import { h } from 'vue';

import { Switch } from 'ant-design-vue';

import { ActionButtons } from '#/components/table-action';

/**
 * 表格列配置
 */
export function useGridColumns(
  onActionClick: (params: any) => void,
  onStatusChange: (row: SystemDeptApi.Dept) => void,
): VxeGridProps['columns'] {
  return [
    {
      type: 'seq',
      width: 50,
      fixed: 'left'
    },
    {
      field: 'name',
      title: '部门名称',
      minWidth: 200,
      treeNode: true,
    },
    {
      field: 'adminId',
      title: '负责人',
      width: 150,
      formatter: ({ row }: any) => {
        // 这里可以关联显示负责人名称
        return row.adminName || '-';
      },
    },
    {
      field: 'sort',
      title: '排序',
      width: 100,
    },
    {
      field: 'status',
      title: '状态',
      width: 100,
      cellRender: {
        name: 'VbenCellRender',
        props: ({ row }: any) => ({
          render: () =>
            h(Switch, {
              checked: row.status === 1,
              checkedChildren: '启用',
              unCheckedChildren: '禁用',
              onChange: () => {
                row.status = row.status === 1 ? -1 : 1;
                onStatusChange(row);
              },
            }),
        }),
      },
    },
    {
      field: 'createdAt',
      title: '创建时间',
      width: 180,
      formatter: ({ cellValue }) => {
        return cellValue
          ? new Date(cellValue).toLocaleString('zh-CN')
          : '';
      },
    },
    {
      title: '操作',
      width: 200,
      fixed: 'right',
      cellRender: {
        name: 'VbenCellRender',
        props: ({ row }: any) => ({
          render: () =>
            h(ActionButtons, {
              row,
              actions: [
                {
                  code: 'append',
                  label: '添加下级',
                  auth: 'system:dept:create',
                },
                {
                  code: 'edit',
                  label: '编辑',
                  auth: 'system:dept:update',
                },
                {
                  code: 'delete',
                  label: '删除',
                  auth: 'system:dept:delete',
                  confirm: true,
                  confirmTitle: `确定要删除"${row.name}"吗？`,
                },
              ],
              onClick: onActionClick,
            }),
        }),
      },
    },
  ];
}

/**
 * 表单配置
 */
export function useFormSchema() {
  return [
    {
      field: 'id',
      label: 'ID',
      component: 'Input',
      componentProps: {
        disabled: true,
      },
      ifShow: ({ values }) => !!values.id,
    },
    {
      field: 'pid',
      label: '上级部门',
      component: 'TreeSelect',
      componentProps: {
        placeholder: '请选择上级部门',
        allowClear: true,
        // 这里需要提供部门树数据
        treeData: [],
        fieldNames: {
          label: 'name',
          value: 'id',
        },
      },
    },
    {
      field: 'name',
      label: '部门名称',
      component: 'Input',
      rules: 'required',
      componentProps: {
        placeholder: '请输入部门名称',
      },
    },
    {
      field: 'adminId',
      label: '负责人',
      component: 'Select',
      componentProps: {
        placeholder: '请选择负责人',
        allowClear: true,
        showSearch: true,
        // 这里需要提供管理员选项
        options: [],
        fieldNames: {
          label: 'username',
          value: 'id',
        },
      },
    },
    {
      field: 'sort',
      label: '排序',
      component: 'InputNumber',
      defaultValue: 0,
      componentProps: {
        placeholder: '请输入排序值',
        min: 0,
        style: { width: '100%' },
      },
    },
    {
      field: 'status',
      label: '状态',
      component: 'RadioGroup',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: '启用', value: 1 },
          { label: '禁用', value: -1 },
        ],
      },
    },
  ];
}
```

**关键点**：
- 表格列支持自定义渲染（cellRender）
- 操作列使用 ActionButtons 组件统一管理
- 表单配置支持动态显示（ifShow）
- 表单验证使用 rules 属性

---

## 3. 列表页面

**文件路径**：`apps/web-antd/src/views/system/dept/index.vue`

```vue
<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SystemAdminApi } from '#/api/system/admin';
import type { SystemDeptApi } from '#/api/system/dept';

import { onMounted, ref } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getAdminSelector } from '#/api/system/admin';
import {
  deleteDept,
  getDeptList,
  updateDeptStatus
} from '#/api/system/dept';
import { $t } from '#/locales';

import { useGridColumns } from './data';
import Form from './modules/form.vue';

// 表单弹窗
const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

// 管理员列表（用于表单选择）
const adminList = ref<SystemAdminApi.Admin[]>([]);

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 切换树形展开/收缩状态 */
const isExpanded = ref(true);
function toggleExpand() {
  isExpanded.value = !isExpanded.value;
  gridApi.grid.setAllTreeExpand(isExpanded.value);
}

/** 创建部门 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 添加下级部门 */
function onAppend(row: SystemDeptApi.Dept) {
  formModalApi.setData({ pid: row.id }).open();
}

/** 编辑部门 */
function onEdit(row: SystemDeptApi.Dept) {
  formModalApi.setData(row).open();
}

/** 删除部门 */
async function onDelete(row: SystemDeptApi.Dept) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteDept(row.id as string);
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 更新状态 */
async function onStatusChange(row: SystemDeptApi.Dept) {
  try {
    await updateDeptStatus(row.id, row.status);
    message.success('状态更新成功');
  } catch {
    message.error('状态更新失败');
    // 失败后恢复原状态
    row.status = row.status === 1 ? -1 : 1;
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row
}: OnActionClickParams<SystemDeptApi.Dept>) {
  switch (code) {
    case 'append': {
      onAppend(row);
      break;
    }
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
  }
}

// 表格配置
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: useGridColumns(onActionClick, onStatusChange),
    height: 'auto',
    keepSource: true,
    pagerConfig: {
      enabled: false,  // 树形结构不分页
    },
    proxyConfig: {
      ajax: {
        query: async (_params) => {
          return await getDeptList();
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: { code: 'query' },
    },
    treeConfig: {
      parentField: 'pid',
      rowField: 'id',
      transform: true,
      expandAll: true,
      reserve: true,
    },
  } as VxeTableGridOptions<SystemDeptApi.Dept>,
});

/** 初始化 */
onMounted(async () => {
  // 加载管理员列表（用于表单选择）
  const res = await getAdminSelector();
  adminList.value = res.list;
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="部门列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['system:dept:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['部门']) }}
        </Button>
        <Button class="ml-2" @click="toggleExpand">
          {{ isExpanded ? '收缩' : '展开' }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
```

**关键点**：
- 使用 useVbenVxeGrid 创建表格
- 树形数据配置 treeConfig
- 权限控制使用 v-access:code 指令
- 国际化使用 $t() 函数

---

## 4. 表单组件

**文件路径**：`apps/web-antd/src/views/system/dept/modules/form.vue`

```vue
<script lang="ts" setup>
import type { SystemDeptApi } from '#/api/system/dept';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { createDept, getDeptInfo, updateDept } from '#/api/system/dept';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<SystemDeptApi.Dept>();

const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['部门'])
    : $t('ui.actionTitle.create', ['部门']);
});

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useFormSchema(),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    // 提交表单
    const data = (await formApi.getValues()) as SystemDeptApi.Dept;
    try {
      await (formData.value?.id ? updateDept(data) : createDept(data));
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: $t('ui.actionMessage.operationSuccess'),
        key: 'action_process_msg',
      });
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    let data = modalApi.getData<SystemDeptApi.Dept>();
    if (!data) {
      return;
    }
    if (data.id) {
      modalApi.lock();
      try {
        const res = await getDeptInfo(data.id);
        data = res.info;
      } finally {
        modalApi.lock(false);
      }
    }
    // 设置到 values
    formData.value = data;
    await formApi.setValues(formData.value);
  },
});
</script>

<template>
  <Modal :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
```

**关键点**：
- 使用 useVbenForm 创建表单
- 使用 useVbenModal 创建弹窗
- 表单验证和提交逻辑
- 创建和编辑共用一个组件

---

## 5. 路由配置

**文件路径**：`apps/web-antd/src/router/routes/modules/system.ts`

```typescript
import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: () => import('#/layouts/basic.vue'),
    meta: {
      title: '系统管理',
      icon: 'lucide:settings',
      order: 1000,
    },
    children: [
      {
        path: 'dept',
        name: 'SystemDept',
        component: () => import('#/views/system/dept/index.vue'),
        meta: {
          title: '部门管理',
          icon: 'lucide:building',
          authority: ['system:dept:list'],
        },
      },
      // ... 其他路由
    ],
  },
];

export default routes;
```

**关键点**：
- 路由懒加载
- 权限配置使用 authority 字段
- 图标使用 icon 字段

---

## 6. 目录结构

完整的目录结构如下：

```
views/system/dept/
├── index.vue              # 列表页
├── data.ts                # 表格列和表单配置
└── modules/
    └── form.vue          # 创建/编辑表单

api/system/dept/
└── index.ts              # API 接口定义

router/routes/modules/
└── system.ts             # 路由配置
```

---

## 7. 开发步骤

1. **创建 API 接口**：定义所有 CRUD 接口和类型
2. **创建表格列配置**：在 data.ts 中定义表格列
3. **创建表单配置**：在 data.ts 中定义表单项
4. **创建列表页面**：实现表格和操作逻辑
5. **创建表单组件**：实现创建/编辑表单
6. **配置路由**：添加路由配置
7. **测试验证**：测试所有 CRUD 功能

---

## 8. 常见问题

### Q: 如何处理关联数据？

A: 在表格列的 `formatter` 中处理关联显示，在表单中使用 `Select` 或 `TreeSelect` 组件。

### Q: 如何实现树形数据的拖拽排序？

A: 使用 `treeConfig.dragSort` 配置，并实现拖拽回调函数。

### Q: 如何自定义表格列的渲染？

A: 使用 `cellRender` 配置，通过 `h()` 函数渲染自定义组件。

### Q: 表单如何实现动态显示/隐藏？

A: 使用表单项的 `ifShow` 属性，根据其他字段值控制显示。

---

## 相关文档

- [进阶功能示例](./02-advanced-features.md) - 搜索表单、批量操作等
- [详情页示例](./03-detail-page.md) - 详情页的两种实现风格
- [组件使用指南](../references/components-guide.md) - 所有可用组件的详细文档
- [最佳实践](../references/best-practices.md) - 代码规范和优化建议
