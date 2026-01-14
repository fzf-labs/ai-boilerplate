# 常用组件使用指南

## 1. 表单组件 (useVbenForm)

### 基础用法

```typescript
import { useVbenForm } from '#/adapter/form';

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',  // horizontal | vertical | inline
  schema: [
    // 表单项配置
  ],
  showDefaultActions: false,  // 是否显示默认按钮
});
```

### 常用表单项

#### Input - 文本输入

```typescript
{
  field: 'username',
  label: '用户名',
  component: 'Input',
  rules: 'required',
  componentProps: {
    placeholder: '请输入用户名',
    maxLength: 50,
    showCount: true,
  },
}
```

#### InputNumber - 数字输入

```typescript
{
  field: 'age',
  label: '年龄',
  component: 'InputNumber',
  defaultValue: 0,
  componentProps: {
    min: 0,
    max: 150,
    placeholder: '请输入年龄',
    style: { width: '100%' },
  },
}
```

#### Select - 下拉选择

```typescript
{
  field: 'type',
  label: '类型',
  component: 'Select',
  componentProps: {
    placeholder: '请选择类型',
    allowClear: true,
    showSearch: true,
    options: [
      { label: '类型1', value: 1 },
      { label: '类型2', value: 2 },
    ],
  },
}
```

#### TreeSelect - 树形选择

```typescript
{
  field: 'deptId',
  label: '所属部门',
  component: 'TreeSelect',
  componentProps: {
    placeholder: '请选择部门',
    allowClear: true,
    treeData: [],  // 树形数据
    fieldNames: {
      label: 'name',
      value: 'id',
      children: 'children',
    },
  },
}
```

#### DatePicker - 日期选择

```typescript
{
  field: 'birthday',
  label: '生日',
  component: 'DatePicker',
  componentProps: {
    placeholder: '请选择日期',
    format: 'YYYY-MM-DD',
    valueFormat: 'YYYY-MM-DD',
    style: { width: '100%' },
  },
}
```

#### RangePicker - 日期范围

```typescript
{
  field: 'dateRange',
  label: '日期范围',
  component: 'RangePicker',
  componentProps: {
    placeholder: ['开始日期', '结束日期'],
    format: 'YYYY-MM-DD',
    valueFormat: 'YYYY-MM-DD',
    style: { width: '100%' },
  },
}
```

#### RadioGroup - 单选框组

```typescript
{
  field: 'gender',
  label: '性别',
  component: 'RadioGroup',
  defaultValue: 1,
  componentProps: {
    options: [
      { label: '男', value: 1 },
      { label: '女', value: 2 },
    ],
  },
}
```

#### CheckboxGroup - 多选框组

```typescript
{
  field: 'hobbies',
  label: '爱好',
  component: 'CheckboxGroup',
  componentProps: {
    options: [
      { label: '运动', value: 'sport' },
      { label: '阅读', value: 'read' },
      { label: '音乐', value: 'music' },
    ],
  },
}
```

#### Switch - 开关

```typescript
{
  field: 'enabled',
  label: '是否启用',
  component: 'Switch',
  defaultValue: true,
  componentProps: {
    checkedChildren: '启用',
    unCheckedChildren: '禁用',
  },
}
```

#### Textarea - 多行文本

```typescript
{
  field: 'description',
  label: '描述',
  component: 'Textarea',
  componentProps: {
    placeholder: '请输入描述',
    rows: 4,
    maxLength: 500,
    showCount: true,
  },
}
```

#### Upload - 文件上传

```typescript
{
  field: 'avatar',
  label: '头像',
  component: 'Upload',
  componentProps: {
    maxCount: 1,
    accept: 'image/*',
    beforeUpload: (file) => {
      // 上传前验证
      return true;
    },
  },
}
```

### 表单验证

#### 内置规则

```typescript
{
  field: 'email',
  label: '邮箱',
  component: 'Input',
  rules: 'required|email',  // 必填 + 邮箱格式
}

{
  field: 'phone',
  label: '手机号',
  component: 'Input',
  rules: 'required|phone',  // 必填 + 手机号格式
}

{
  field: 'url',
  label: '网址',
  component: 'Input',
  rules: 'url',  // URL 格式
}
```

#### 自定义规则

```typescript
{
  field: 'password',
  label: '密码',
  component: 'InputPassword',
  rules: [
    { required: true, message: '请输入密码' },
    { min: 6, message: '密码长度不能小于6位' },
    {
      validator: async (rule, value) => {
        if (!/[A-Z]/.test(value)) {
          return Promise.reject('密码必须包含大写字母');
        }
        return Promise.resolve();
      },
    },
  ],
}
```

### 表单联动

#### 动态显示/隐藏

```typescript
{
  field: 'type',
  label: '类型',
  component: 'Select',
  componentProps: {
    options: [
      { label: '文本', value: 'text' },
      { label: '图片', value: 'image' },
    ],
  },
},
{
  field: 'content',
  label: '文本内容',
  component: 'Textarea',
  // 只有类型为文本时才显示
  ifShow: ({ values }) => values.type === 'text',
},
{
  field: 'imageUrl',
  label: '图片地址',
  component: 'Upload',
  // 只有类型为图片时才显示
  ifShow: ({ values }) => values.type === 'image',
}
```

#### 动态禁用

```typescript
{
  field: 'username',
  label: '用户名',
  component: 'Input',
  componentProps: ({ values }) => ({
    disabled: !!values.id,  // 编辑时禁用
  }),
}
```

### 表单 API

```typescript
// 获取表单值
const values = await formApi.getValues();

// 设置表单值
await formApi.setValues({ username: 'admin' });

// 设置单个字段值
formApi.setFieldValue('username', 'admin');

// 验证表单
const { valid, errors } = await formApi.validate();

// 验证单个字段
await formApi.validateField('username');

// 重置表单
formApi.resetFields();

// 清空表单
formApi.clearValidate();
```

---

## 2. 表格组件 (useVbenVxeGrid)

### 基础用法

```typescript
import { useVbenVxeGrid } from '#/adapter/vxe-table';

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: [],  // 列配置
    height: 'auto',
    keepSource: true,
    pagerConfig: {
      enabled: true,
      pageSize: 20,
    },
    proxyConfig: {
      ajax: {
        query: async ({ page }) => {
          // 查询数据
        },
      },
    },
  },
});
```

### 列配置

#### 序号列

```typescript
{
  type: 'seq',
  width: 50,
  fixed: 'left',
}
```

#### 多选列

```typescript
{
  type: 'checkbox',
  width: 50,
  fixed: 'left',
}
```

#### 普通列

```typescript
{
  field: 'name',
  title: '名称',
  minWidth: 150,
  sortable: true,  // 可排序
}
```

#### 自定义渲染

```typescript
import { h } from 'vue';
import { Tag } from 'ant-design-vue';

{
  field: 'status',
  title: '状态',
  width: 100,
  cellRender: {
    name: 'VbenCellRender',
    props: ({ row }: any) => ({
      render: () => {
        const statusMap = {
          1: { color: 'green', text: '启用' },
          -1: { color: 'red', text: '禁用' },
        };
        const status = statusMap[row.status];
        return h(Tag, { color: status.color }, () => status.text);
      },
    }),
  },
}
```

#### 操作列

```typescript
import { ActionButtons } from '#/components/table-action';

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
              code: 'edit',
              label: '编辑',
              auth: 'system:user:update',
            },
            {
              code: 'delete',
              label: '删除',
              auth: 'system:user:delete',
              confirm: true,
              confirmTitle: '确定要删除吗？',
            },
          ],
          onClick: (params) => {
            console.log(params.code, params.row);
          },
        }),
    }),
  },
}
```

### 分页配置

```typescript
{
  pagerConfig: {
    enabled: true,
    pageSize: 20,
    pageSizes: [10, 20, 50, 100],
    layouts: [
      'Sizes',
      'PrevJump',
      'PrevPage',
      'Number',
      'NextPage',
      'NextJump',
      'FullJump',
      'Total',
    ],
  },
}
```

### 树形配置

```typescript
{
  treeConfig: {
    parentField: 'pid',
    rowField: 'id',
    transform: true,
    expandAll: true,
    reserve: true,
  },
  pagerConfig: {
    enabled: false,  // 树形结构通常不分页
  },
}
```

### 表格 API

```typescript
// 刷新数据
gridApi.query();

// 获取选中行
const rows = gridApi.grid.getCheckboxRecords();

// 设置选中行
gridApi.grid.setCheckboxRow(rows, true);

// 清空选中
gridApi.grid.clearCheckboxRow();

// 展开/收缩树形
gridApi.grid.setAllTreeExpand(true);

// 获取表格数据
const data = gridApi.grid.getTableData().fullData;
```

---

## 3. 弹窗组件 (useVbenModal)

### 基础用法

```typescript
import { useVbenModal } from '@vben/common-ui';

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    // 确认逻辑
    await modalApi.close();
  },
  onCancel: () => {
    // 取消逻辑
  },
  onOpenChange: (isOpen) => {
    if (isOpen) {
      // 打开时的逻辑
      const data = modalApi.getData();
    } else {
      // 关闭时的逻辑
    }
  },
});
```

### 弹窗配置

```vue
<template>
  <Modal
    :title="title"
    :width="600"
    :fullscreen="false"
    :loading="loading"
  >
    <div>弹窗内容</div>
  </Modal>
</template>
```

### 弹窗 API

```typescript
// 打开弹窗
modalApi.open();

// 关闭弹窗
modalApi.close();

// 设置数据
modalApi.setData({ id: '123' });

// 获取数据
const data = modalApi.getData();

// 锁定弹窗（加载中）
modalApi.lock(true);
modalApi.lock(false);
```

### 表单弹窗示例

```vue
<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';

const [Form, formApi] = useVbenForm({
  schema: [
    {
      field: 'name',
      label: '名称',
      component: 'Input',
      rules: 'required',
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) return;

    const values = await formApi.getValues();
    // 提交数据
    await submitData(values);
    await modalApi.close();
  },
  onOpenChange(isOpen) {
    if (!isOpen) {
      formApi.resetFields();
    } else {
      const data = modalApi.getData();
      if (data) {
        formApi.setValues(data);
      }
    }
  },
});
</script>

<template>
  <Modal title="编辑">
    <Form class="mx-4" />
  </Modal>
</template>
```

---

## 4. 权限控制

### 指令方式

```vue
<template>
  <!-- 按钮权限 -->
  <Button v-access:code="['system:user:create']">
    创建
  </Button>

  <!-- 多个权限（满足任一即可） -->
  <Button v-access:code="['system:user:update', 'system:user:delete']">
    操作
  </Button>

  <!-- 多个权限（需全部满足） -->
  <Button v-access:code:and="['system:user:update', 'admin']">
    高级操作
  </Button>
</template>
```

### 函数方式

```vue
<script setup>
import { useAccess } from '@vben/access';

const { hasAccessByCode } = useAccess();

if (hasAccessByCode('system:user:create')) {
  // 有创建权限
}

// 判断多个权限
if (hasAccessByCode(['system:user:update', 'system:user:delete'])) {
  // 有更新或删除权限
}
</script>
```

### 路由权限

```typescript
{
  path: '/system/user',
  name: 'SystemUser',
  component: () => import('#/views/system/user/index.vue'),
  meta: {
    title: '用户管理',
    authority: ['system:user:list'],  // 页面访问权限
  },
}
```

---

## 5. 国际化

### 使用翻译

```vue
<template>
  <!-- 模板中使用 -->
  <div>{{ $t('ui.actionTitle.create', ['用户']) }}</div>
  <div>{{ $t('ui.actionMessage.deleteSuccess') }}</div>
</template>

<script setup>
import { $t } from '#/locales';

// 脚本中使用
const message = $t('ui.actionMessage.operationSuccess');
console.log(message);
</script>
```

### 添加翻译

**文件路径**：`packages/locales/src/langs/zh-CN.json`

```json
{
  "ui": {
    "actionTitle": {
      "create": "创建{0}",
      "edit": "编辑{0}",
      "delete": "删除{0}"
    },
    "actionMessage": {
      "operationSuccess": "操作成功",
      "deleteSuccess": "删除成功"
    }
  }
}
```

---

## 6. 文件上传

### 单文件上传

```vue
<script setup>
import { ref } from 'vue';
import { uploadFile } from '#/utils/upload-helper';

const fileUrl = ref('');

async function handleUpload(file: File) {
  try {
    const url = await uploadFile(file);
    fileUrl.value = url;
    message.success('上传成功');
  } catch {
    message.error('上传失败');
  }
}
</script>

<template>
  <Upload
    :before-upload="handleUpload"
    :show-upload-list="false"
  >
    <Button>上传文件</Button>
  </Upload>
  <img v-if="fileUrl" :src="fileUrl" />
</template>
```

### 多文件上传

```vue
<script setup>
import { ref } from 'vue';

const fileList = ref<any[]>([]);

function handleChange({ fileList: newFileList }) {
  fileList.value = newFileList;
}
</script>

<template>
  <Upload
    v-model:file-list="fileList"
    :multiple="true"
    :max-count="5"
    @change="handleChange"
  >
    <Button>上传文件</Button>
  </Upload>
</template>
```

---

## 7. 消息提示

```typescript
import { message, notification } from 'ant-design-vue';

// 成功提示
message.success('操作成功');

// 错误提示
message.error('操作失败');

// 警告提示
message.warning('请注意');

// 加载提示
const hide = message.loading('加载中...', 0);
// 隐藏
hide();

// 通知
notification.success({
  message: '操作成功',
  description: '数据已保存',
});
```

---

## 8. 页面布局

```vue
<template>
  <Page auto-content-height :loading="loading">
    <!-- 搜索区域 -->
    <div class="mb-4">
      <SearchForm />
    </div>

    <!-- 表格区域 -->
    <Grid>
      <template #toolbar-tools>
        <Button type="primary" @click="onCreate">
          创建
        </Button>
      </template>
    </Grid>

    <!-- 弹窗 -->
    <FormModal @success="onRefresh" />
  </Page>
</template>
```
