<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysAdminInfo } from '#/api/v1/sys-admin';
import type { SysDeptInfo } from '#/api/v1/sys-dept';

import { onMounted, ref } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { deleteSysDept, getSysDeptList } from '#/api/v1/sys-dept';
import { $t } from '#/locales';

import { useGridColumns } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const adminList = ref<SysAdminInfo[]>([]);

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
function onAppend(row: SysDeptInfo) {
  formModalApi.setData({ pid: row.id }).open();
}

/** 编辑部门 */
function onEdit(row: SysDeptInfo) {
  formModalApi.setData(row).open();
}

/** 删除部门 */
async function onDelete(row: SysDeptInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSysDept({ body: { id: row.id as string } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SysDeptInfo>) {
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

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    pagerConfig: {
      enabled: false,
    },
    proxyConfig: {
      ajax: {
        query: async (_params) => {
          return await getSysDeptList({ options: {} });
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
  } as VxeTableGridOptions<SysDeptInfo>,
});

/** 初始化 */
onMounted(async () => {
  const res = await getSysAdminSelector({ options: {} });
  adminList.value = res.list || [];
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
