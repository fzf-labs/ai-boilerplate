# 详情页实现

本示例展示如何实现具有丰富视觉效果的详情页，使用 Card 组件 + 自定义样式，提供优秀的用户体验。

> **前置条件**：请先阅读 [基础 CRUD 示例](./01-basic-crud.md) 了解基本的 CRUD 实现。

## 特点

- ✅ 视觉效果丰富，用户体验好
- ✅ 使用渐变背景和卡片阴影
- ✅ 悬停动画效果
- ✅ 响应式设计
- ✅ 自定义主题色

---

## 实现代码

### 效果预览

使用 Card 组件 + 自定义样式实现视觉效果丰富的详情页。

**文件路径**：`apps/web-antd/src/views/system/dept/modules/detail.vue`

```vue
<script lang="ts" setup>
import type { SysDeptApi } from '#/api/v1/sys-dept';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Avatar, Card, Tag } from 'ant-design-vue';

import { getSysDeptInfo } from '#/api/v1/sys-dept';

const deptData = ref<SysDeptApi.SysDept>();

const getTitle = computed(() => {
  const data = deptData.value;
  return data ? `部门详情 - ${data.name}` : '部门详情';
});

// 状态映射
const statusMap: Record<number, { color: string; icon: string; text: string }> = {
  1: { color: 'success', icon: '✅', text: '启用' },
  [-1]: { color: 'error', icon: '❌', text: '禁用' },
};

// 获取状态信息
const getStatusInfo = computed(() => {
  const status = deptData.value?.status ?? -1;
  return statusMap[status] || statusMap[-1];
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      deptData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysDeptApi.SysDept>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getSysDeptInfo(data.id);
      deptData.value = res.info;
    } catch (error) {
      console.error('加载部门详情失败:', error);
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal :title="getTitle" class="dept-detail-modal w-full max-w-4xl">
    <div v-if="deptData" class="dept-detail-content">
      <!-- 部门头部信息 -->
      <div
        class="dept-header mb-6 rounded-lg bg-gradient-to-r from-blue-50 via-purple-50 to-pink-50 p-6"
      >
        <div class="flex items-start space-x-6">
          <div class="shrink-0">
            <Avatar :size="96" class="shadow-lg ring-4 ring-white">
              <template #icon>
                <span class="text-4xl">🏛️</span>
              </template>
            </Avatar>
          </div>
          <div class="flex-1">
            <div class="mb-3">
              <h2 class="mb-2 text-2xl font-bold text-gray-800">
                {{ deptData.name }}
              </h2>
              <div class="flex items-center gap-3">
                <Tag :color="getStatusInfo?.color" class="rounded-full">
                  {{ getStatusInfo?.icon }} {{ getStatusInfo?.text }}
                </Tag>
                <Tag color="blue" class="rounded-full">
                  排序: {{ deptData.sort }}
                </Tag>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <!-- 基本信息 -->
        <Card title="📋 基本信息" size="small" class="info-card">
          <div class="space-y-4">
            <div class="info-item bg-blue-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">部门ID</span>
                <span class="font-mono text-sm text-gray-800">
                  {{ deptData.id }}
                </span>
              </div>
            </div>
            <div class="info-item bg-green-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">部门名称</span>
                <span class="font-semibold text-gray-900">
                  {{ deptData.name }}
                </span>
              </div>
            </div>
            <div class="info-item bg-purple-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">负责人</span>
                <Tag color="blue" class="rounded">
                  👤 {{ deptData.adminName || '未分配' }}
                </Tag>
              </div>
            </div>
          </div>
        </Card>

        <!-- 时间记录 -->
        <Card title="⏰ 时间记录" size="small" class="time-card">
          <div class="space-y-4">
            <div class="info-item bg-cyan-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">创建时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(deptData.createdAt || '')) }}
                </span>
              </div>
            </div>
            <div class="info-item bg-emerald-50">
              <div class="flex items-center justify-between">
                <span class="font-medium text-gray-600">更新时间</span>
                <span class="font-medium text-gray-800">
                  {{ formatDateTime(String(deptData.updatedAt || '')) }}
                </span>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
/* 响应式优化 */
@media (max-width: 768px) {
  .dept-header {
    padding: 1rem;
  }

  .dept-header .flex {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }

  .dept-detail-content {
    padding: 1rem;
  }
}

.dept-detail-modal :deep(.vben-modal .ant-modal-body) {
  padding: 0;
}

.dept-detail-content {
  padding: 1.5rem;
}

/* 部门头部渐变背景增强 */
.dept-header {
  position: relative;
  overflow: hidden;
}

.dept-header::before {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: '';
  background: linear-gradient(
    135deg,
    rgb(59 130 246 / 10%),
    rgb(147 51 234 / 10%),
    rgb(236 72 153 / 10%)
  );
}

/* 信息项样式 */
.info-item {
  padding: 1rem;
  cursor: pointer;
  border-radius: 0.5rem;
  transition: all 0.2s;
}

.info-item:hover {
  box-shadow:
    0 10px 15px -3px rgb(0 0 0 / 10%),
    0 4px 6px -2px rgb(0 0 0 / 5%);
  transform: translateX(0.25rem) scale(1.02);
}

/* 卡片悬停效果 */
:deep(.ant-card) {
  background: linear-gradient(145deg, #fff, #f8fafc);
  border-color: rgb(229 231 235);
  transition: all 0.3s ease;
}

:deep(.ant-card:hover) {
  border-color: rgb(99 102 241 / 30%);
  box-shadow: 0 8px 25px rgb(0 0 0 / 10%);
  transform: translateY(-2px);
}

/* 特定卡片样式 */
.info-card {
  border-left: 4px solid #3b82f6;
}

.time-card {
  border-left: 4px solid #06b6d4;
}
</style>
```

---

## 在列表页中集成详情页

### 步骤 1：引入详情组件

更新 `index.vue`：

```vue
<script lang="ts" setup>
// ... 其他导入

import Detail from './modules/detail.vue';

// ... 其他代码

// 详情弹窗
const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,
  showConfirmButton: false,  // 详情页不需要确认按钮
});

/** 查看详情 */
function onDetail(row: SysDeptApi.SysDept) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row
}: OnActionClickParams<SysDeptApi.SysDept>) {
  switch (code) {
    case 'detail': {
      onDetail(row);
      break;
    }
    case 'append': {
      onAppend(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'delete': {
      onDelete(row);
      break;
    }
  }
}

// ... 其他代码
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <DetailModal />
    <Grid table-title="部门列表">
      <!-- ... -->
    </Grid>
  </Page>
</template>
```

### 步骤 2：在表格中添加详情按钮

更新 `data.ts`，在操作列中添加详情按钮：

```typescript
export function useGridColumns(
  onActionClick: (params: any) => void,
  onStatusChange: (row: SysDeptApi.SysDept) => void,
): VxeGridProps['columns'] {
  return [
    // ... 其他列

    {
      title: '操作',
      width: 250,  // 增加宽度以容纳详情按钮
      fixed: 'right',
      cellRender: {
        name: 'VbenCellRender',
        props: ({ row }: any) => ({
          render: () =>
            h(ActionButtons, {
              row,
              actions: [
                {
                  code: 'detail',
                  label: '详情',
                  auth: 'system:dept:query',  // 查看权限
                },
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
```

---

## 详情页扩展

### 1. 添加Tab切换

适合有多类信息需要展示的场景：

```vue
<template>
  <Modal :title="getTitle" class="w-full max-w-4xl">
    <div v-if="deptData">
      <Tabs v-model:activeKey="activeTab">
        <TabPane key="basic" tab="基本信息">
          <Descriptions bordered :column="2">
            <!-- 基本信息 -->
          </Descriptions>
        </TabPane>

        <TabPane key="members" tab="部门成员">
          <!-- 成员列表 -->
          <Table :dataSource="members" />
        </TabPane>

        <TabPane key="logs" tab="操作日志">
          <!-- 操作日志 -->
          <Timeline>
            <TimelineItem v-for="log in logs" :key="log.id">
              {{ log.content }}
            </TimelineItem>
          </Timeline>
        </TabPane>
      </Tabs>
    </div>
  </Modal>
</template>

<script setup>
import { Tabs, TabPane, Table, Timeline, TimelineItem } from 'ant-design-vue';

const activeTab = ref('basic');
const members = ref([]);
const logs = ref([]);
</script>
```

### 2. 添加关联数据

展示关联的数据，如部门下的成员：

```vue
<template>
  <Card title="👥 部门成员" size="small">
    <div class="space-y-2">
      <div
        v-for="member in members"
        :key="member.id"
        class="flex items-center justify-between rounded-lg bg-gray-50 p-3"
      >
        <div class="flex items-center gap-3">
          <Avatar :src="member.avatar" />
          <div>
            <div class="font-medium">{{ member.name }}</div>
            <div class="text-sm text-gray-500">{{ member.position }}</div>
          </div>
        </div>
        <Tag :color="member.status === 1 ? 'green' : 'red'">
          {{ member.status === 1 ? '在职' : '离职' }}
        </Tag>
      </div>
    </div>
  </Card>
</template>

<script setup>
const members = ref([]);

// 加载部门成员
async function loadMembers(deptId: string) {
  const res = await getDeptMembers(deptId);
  members.value = res.list;
}
</script>
```

### 3. 添加操作按钮

在详情页中添加快捷操作：

```vue
<template>
  <Modal :title="getTitle" class="w-full max-w-4xl">
    <template #footer>
      <Space>
        <Button @click="onEdit">编辑</Button>
        <Button type="primary" danger @click="onDelete">删除</Button>
        <Button @click="modalApi.close()">关闭</Button>
      </Space>
    </template>

    <div v-if="deptData">
      <!-- 详情内容 -->
    </div>
  </Modal>
</template>

<script setup>
import { Space } from 'ant-design-vue';

function onEdit() {
  // 跳转到编辑页或打开编辑弹窗
  modalApi.close();
  formModalApi.setData(deptData.value).open();
}

function onDelete() {
  // 删除逻辑
}
</script>
```

### 4. 添加数据图表

使用图表展示数据统计：

```vue
<template>
  <Card title="📊 数据统计" size="small">
    <div id="chart" style="width: 100%; height: 300px;"></div>
  </Card>
</template>

<script setup>
import * as echarts from 'echarts';
import { onMounted, onUnmounted } from 'vue';

let chartInstance: echarts.ECharts | null = null;

onMounted(() => {
  chartInstance = echarts.init(document.getElementById('chart'));
  chartInstance.setOption({
    // ECharts 配置
    xAxis: { type: 'category', data: ['周一', '周二', '周三'] },
    yAxis: { type: 'value' },
    series: [{ data: [120, 200, 150], type: 'bar' }],
  });
});

onUnmounted(() => {
  chartInstance?.dispose();
});
</script>
```

---

## 最佳实践

### 1. 数据加载

```typescript
// ✅ 推荐：在弹窗打开时加载数据
const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }

    const data = modalApi.getData<SysDeptApi.SysDept>();
    if (!data?.id) return;

    modalApi.lock();
    try {
      const res = await getSysDeptInfo(data.id);
      formData.value = res.info;
    } catch (error) {
      message.error('加载失败');
      modalApi.close();
    } finally {
      modalApi.lock(false);
    }
  },
});

// ❌ 不推荐：在列表页就加载所有详细数据
```

### 2. 错误处理

```typescript
// ✅ 推荐：捕获错误并提示用户
try {
  const res = await getSysDeptInfo(data.id);
  formData.value = res.info;
} catch (error) {
  message.error('加载详情失败');
  modalApi.close();  // 加载失败时关闭弹窗
}

// ❌ 不推荐：不处理错误
const res = await getSysDeptInfo(data.id);
formData.value = res.info;
```

### 3. 性能优化

```vue
<script setup>
// ✅ 推荐：使用 destroyOnClose 销毁未使用的弹窗
const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,  // 关闭时销毁组件
});

// ✅ 推荐：大量数据时使用虚拟滚动
import { VirtualList } from '@vben/common-ui';
</script>
```

---

## 相关文档

- [基础 CRUD 示例](./01-basic-crud.md) - 基础 CRUD 实现
- [进阶功能示例](./02-advanced-features.md) - 搜索、批量操作等
- [组件使用指南](../references/components-guide.md) - 所有可用组件的详细文档
- [最佳实践](../references/best-practices.md) - 代码规范和优化建议
