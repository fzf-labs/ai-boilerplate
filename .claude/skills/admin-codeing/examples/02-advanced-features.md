# 进阶功能实现

本示例展示如何在基础 CRUD 页面上添加搜索表单、批量操作等进阶功能。

> **前置条件**：请先阅读 [基础 CRUD 示例](./01-basic-crud.md) 了解基本的 CRUD 实现。

---

## 1. 搜索表单

在列表页面上方添加搜索表单，支持按条件筛选数据。

### 实现步骤

#### 步骤 1：定义搜索表单配置

在 `data.ts` 中添加搜索表单配置：

```typescript
/**
 * 搜索表单配置
 */
export function useSearchSchema() {
  return [
    {
      field: 'name',
      label: '部门名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入部门名称',
      },
    },
    {
      field: 'status',
      label: '状态',
      component: 'Select',
      componentProps: {
        placeholder: '请选择状态',
        allowClear: true,
        options: [
          { label: '全部', value: undefined },
          { label: '启用', value: 1 },
          { label: '禁用', value: -1 },
        ],
      },
    },
  ];
}
```

#### 步骤 2：在列表页中使用搜索表单

更新 `index.vue`：

```vue
<script lang="ts" setup>
// ... 其他导入

import { useVbenForm } from '#/adapter/form';
import { useSearchSchema } from './data';

// ... 其他代码

// 搜索表单
const [SearchForm, searchFormApi] = useVbenForm({
  layout: 'inline',
  schema: useSearchSchema(),
  showDefaultActions: false,
});

/** 搜索 */
async function onSearch() {
  const values = await searchFormApi.getValues();
  gridApi.grid.commitProxy('query', { ...values });
}

/** 重置 */
function onReset() {
  searchFormApi.resetFields();
  gridApi.query();
}

// ... 其他代码

// 更新表格配置，支持搜索参数
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    // ... 其他配置
    proxyConfig: {
      ajax: {
        query: async ({ page, form }) => {
          // form 参数包含搜索表单的值
          return await getDeptList({
            ...form,
            page: page.currentPage,
            pageSize: page.pageSize,
          });
        },
      },
    },
  },
});
</script>

<template>
  <Page auto-content-height>
    <!-- 搜索表单 -->
    <div class="mb-4 rounded-lg bg-white p-4">
      <SearchForm />
      <div class="mt-4 flex justify-end gap-2">
        <Button @click="onReset">重置</Button>
        <Button type="primary" @click="onSearch">搜索</Button>
      </div>
    </div>

    <!-- 表格 -->
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
      </template>
    </Grid>
  </Page>
</template>
```

### 高级用法

#### 日期范围搜索

```typescript
{
  field: 'dateRange',
  label: '创建时间',
  component: 'RangePicker',
  componentProps: {
    placeholder: ['开始日期', '结束日期'],
    format: 'YYYY-MM-DD',
    valueFormat: 'YYYY-MM-DD',
    style: { width: '100%' },
  },
}
```

#### 级联选择

```typescript
{
  field: 'region',
  label: '地区',
  component: 'Cascader',
  componentProps: {
    placeholder: '请选择地区',
    options: regionOptions,
    fieldNames: {
      label: 'name',
      value: 'code',
      children: 'children',
    },
  },
}
```

#### 自动搜索（输入时自动触发）

```typescript
import { useDebounceFn } from '@vueuse/core';

// 防抖搜索
const debouncedSearch = useDebounceFn(() => {
  onSearch();
}, 300);

// 监听表单值变化
watch(
  () => searchFormApi.getValues(),
  () => {
    debouncedSearch();
  },
  { deep: true }
);
```

---

## 2. 批量操作

支持选中多行数据进行批量操作（如批量删除、批量导出等）。

### 实现步骤

#### 步骤 1：启用表格多选

更新表格列配置，在 `data.ts` 中：

```typescript
export function useGridColumns(
  onActionClick: (params: any) => void,
  onStatusChange: (row: SystemDeptApi.Dept) => void,
): VxeGridProps['columns'] {
  return [
    // 添加多选列
    {
      type: 'checkbox',
      width: 50,
      fixed: 'left',
    },
    {
      type: 'seq',
      width: 50,
      fixed: 'left'
    },
    // ... 其他列配置
  ];
}
```

#### 步骤 2：实现批量操作逻辑

更新 `index.vue`：

```vue
<script lang="ts" setup>
// ... 其他导入

import { Modal } from 'ant-design-vue';

// ... 其他代码

// 选中的行
const selectedRows = ref<SystemDeptApi.Dept[]>([]);

// 监听表格选中状态
function onCheckboxChange() {
  selectedRows.value = gridApi.grid.getCheckboxRecords();
}

/** 批量删除 */
async function onBatchDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  if (rows.length === 0) {
    message.warning('请先选择要删除的数据');
    return;
  }

  Modal.confirm({
    title: '批量删除',
    content: `确定要删除选中的 ${rows.length} 条数据吗？`,
    onOk: async () => {
      const hideLoading = message.loading('删除中...', 0);
      try {
        // 并行删除
        await Promise.all(rows.map(row => deleteDept(row.id)));
        message.success('批量删除成功');
        gridApi.query();
        gridApi.grid.clearCheckboxRow();
      } catch (error) {
        message.error('批量删除失败');
      } finally {
        hideLoading();
      }
    },
  });
}

/** 批量导出 */
async function onBatchExport() {
  const rows = gridApi.grid.getCheckboxRecords();
  if (rows.length === 0) {
    message.warning('请先选择要导出的数据');
    return;
  }

  const hideLoading = message.loading('导出中...', 0);
  try {
    const ids = rows.map(row => row.id);
    // 调用导出接口
    await exportDept(ids);
    message.success('导出成功');
  } catch (error) {
    message.error('导出失败');
  } finally {
    hideLoading();
  }
}

// 表格配置
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    // ... 其他配置
    checkboxConfig: {
      reserve: true,  // 保留选中状态
      // 自定义哪些行可以选中
      checkMethod: ({ row }) => {
        // 例如：禁用状态的行不能选中
        return row.status !== -1;
      },
    },
    // 监听选中状态变化
    on: {
      checkboxChange: onCheckboxChange,
      checkboxAll: onCheckboxChange,
    },
  },
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

        <!-- 批量操作按钮 -->
        <Button
          type="primary"
          danger
          @click="onBatchDelete"
          :disabled="selectedRows.length === 0"
          v-access:code="['system:dept:delete']"
        >
          批量删除 ({{ selectedRows.length }})
        </Button>

        <Button
          @click="onBatchExport"
          :disabled="selectedRows.length === 0"
          v-access:code="['system:dept:export']"
        >
          批量导出 ({{ selectedRows.length }})
        </Button>
      </template>
    </Grid>
  </Page>
</template>
```

### 高级用法

#### 批量修改状态

```typescript
/** 批量启用/禁用 */
async function onBatchUpdateStatus(status: number) {
  const rows = gridApi.grid.getCheckboxRecords();
  if (rows.length === 0) {
    message.warning('请先选择要操作的数据');
    return;
  }

  const statusText = status === 1 ? '启用' : '禁用';
  Modal.confirm({
    title: `批量${statusText}`,
    content: `确定要${statusText}选中的 ${rows.length} 条数据吗？`,
    onOk: async () => {
      const hideLoading = message.loading(`${statusText}中...`, 0);
      try {
        await Promise.all(
          rows.map(row => updateDeptStatus(row.id, status))
        );
        message.success(`批量${statusText}成功`);
        gridApi.query();
      } catch (error) {
        message.error(`批量${statusText}失败`);
      } finally {
        hideLoading();
      }
    },
  });
}
```

#### 全选当前页/全选所有

```typescript
/** 全选当前页 */
function selectCurrentPage() {
  gridApi.grid.setAllCheckboxRow(true);
}

/** 全选所有（包括其他页） */
async function selectAll() {
  // 获取所有数据
  const allData = await getDeptList({ pageSize: 999999 });
  gridApi.grid.setCheckboxRow(allData.list, true);
}

/** 清空选择 */
function clearSelection() {
  gridApi.grid.clearCheckboxRow();
}
```

---

## 3. 数据导入/导出

支持从 Excel 导入数据或导出数据到 Excel。

### 导出功能

```typescript
import { exportToExcel } from '#/utils/excel-helper';

/** 导出所有数据 */
async function onExport() {
  const hideLoading = message.loading('导出中...', 0);
  try {
    // 获取所有数据
    const res = await getDeptList({ pageSize: 999999 });

    // 转换为 Excel 格式
    const data = res.list.map(item => ({
      '部门名称': item.name,
      '负责人': item.adminName,
      '排序': item.sort,
      '状态': item.status === 1 ? '启用' : '禁用',
      '创建时间': new Date(item.createdAt).toLocaleString('zh-CN'),
    }));

    // 导出
    exportToExcel(data, '部门列表');
    message.success('导出成功');
  } catch (error) {
    message.error('导出失败');
  } finally {
    hideLoading();
  }
}
```

### 导入功能

```vue
<script setup>
import { Upload } from 'ant-design-vue';
import { importFromExcel } from '#/utils/excel-helper';

/** 导入 */
async function onImport(file: File) {
  const hideLoading = message.loading('导入中...', 0);
  try {
    // 解析 Excel
    const data = await importFromExcel(file);

    // 转换为 API 格式
    const list = data.map(item => ({
      name: item['部门名称'],
      adminId: item['负责人ID'],
      sort: item['排序'],
      status: item['状态'] === '启用' ? 1 : -1,
    }));

    // 批量创建
    await Promise.all(list.map(item => createDept(item)));

    message.success(`成功导入 ${list.length} 条数据`);
    gridApi.query();
  } catch (error) {
    message.error('导入失败');
  } finally {
    hideLoading();
  }
  return false;  // 阻止默认上传
}
</script>

<template>
  <Upload
    :before-upload="onImport"
    :show-upload-list="false"
    accept=".xlsx,.xls"
  >
    <Button>
      <UploadOutlined />
      导入
    </Button>
  </Upload>
</template>
```

---

## 4. 高级筛选

支持更复杂的筛选条件，如多条件组合、自定义筛选器等。

### 表格列筛选

```typescript
{
  field: 'status',
  title: '状态',
  width: 100,
  // 添加列筛选
  filters: [
    { label: '启用', value: 1 },
    { label: '禁用', value: -1 },
  ],
  filterMultiple: false,  // 单选
}
```

### 自定义筛选面板

```vue
<script setup>
import { Drawer, Form, FormItem, Input, Select, Button } from 'ant-design-vue';

const filterVisible = ref(false);
const filterForm = reactive({
  name: '',
  status: undefined,
  dateRange: [],
  adminId: undefined,
});

function openFilter() {
  filterVisible.value = true;
}

function applyFilter() {
  gridApi.grid.commitProxy('query', { ...filterForm });
  filterVisible.value = false;
}

function resetFilter() {
  Object.assign(filterForm, {
    name: '',
    status: undefined,
    dateRange: [],
    adminId: undefined,
  });
  gridApi.query();
  filterVisible.value = false;
}
</script>

<template>
  <Button @click="openFilter">
    <FilterOutlined />
    高级筛选
  </Button>

  <Drawer
    v-model:open="filterVisible"
    title="高级筛选"
    width="400"
  >
    <Form :model="filterForm" layout="vertical">
      <FormItem label="部门名称">
        <Input v-model:value="filterForm.name" placeholder="请输入" />
      </FormItem>
      <FormItem label="状态">
        <Select v-model:value="filterForm.status" placeholder="请选择">
          <SelectOption :value="1">启用</SelectOption>
          <SelectOption :value="-1">禁用</SelectOption>
        </Select>
      </FormItem>
      <FormItem label="创建时间">
        <RangePicker v-model:value="filterForm.dateRange" />
      </FormItem>
    </Form>

    <template #footer>
      <Button @click="resetFilter">重置</Button>
      <Button type="primary" @click="applyFilter">应用</Button>
    </template>
  </Drawer>
</template>
```

---

## 5. 排序功能

支持单列排序和多列排序。

### 单列排序

```typescript
{
  field: 'sort',
  title: '排序',
  width: 100,
  sortable: true,  // 启用排序
}
```

### 服务端排序

```typescript
const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    // ... 其他配置
    sortConfig: {
      remote: true,  // 服务端排序
    },
    proxyConfig: {
      ajax: {
        query: async ({ page, sort }) => {
          return await getDeptList({
            page: page.currentPage,
            pageSize: page.pageSize,
            orderBy: sort.field,
            orderType: sort.order,  // 'asc' | 'desc'
          });
        },
      },
    },
  },
});
```

---

## 6. 刷新策略

不同场景下的刷新策略。

### 自动刷新

```typescript
import { useIntervalFn } from '@vueuse/core';

// 每 30 秒自动刷新
const { pause, resume } = useIntervalFn(() => {
  gridApi.query();
}, 30000);

// 页面失焦时暂停
onDeactivated(() => pause());
// 页面激活时恢复
onActivated(() => resume());
```

### 条件刷新

```typescript
// 只在数据变化时刷新
let lastDataHash = '';

async function conditionalRefresh() {
  const data = await getDeptList();
  const currentHash = JSON.stringify(data);

  if (currentHash !== lastDataHash) {
    gridApi.query();
    lastDataHash = currentHash;
  }
}
```

---

## 相关文档

- [基础 CRUD 示例](./01-basic-crud.md) - 基础 CRUD 实现
- [详情页示例](./03-detail-page.md) - 详情页的两种实现风格
- [组件使用指南](../references/components-guide.md) - 所有可用组件的详细文档
- [最佳实践](../references/best-practices.md) - 代码规范和优化建议
