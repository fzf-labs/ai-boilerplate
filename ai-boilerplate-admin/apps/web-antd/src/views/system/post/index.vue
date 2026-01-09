<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysPostInfo } from '#/api/v1/sys-post';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteSysPost, getSysPostList } from '#/api/v1/sys-post';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建岗位 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑岗位 */
function onEdit(row: SysPostInfo) {
  formModalApi.setData(row).open();
}

/** 删除岗位 */
async function onDelete(row: SysPostInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSysPost({ body: { id: row.id! } });
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
function onActionClick({ code, row }: OnActionClickParams<SysPostInfo>) {
  switch (code) {
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
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getSysPostList({
            page: page.currentPage,
            pageSize: page.pageSize,
            ...formValues,
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<SysPostInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="岗位列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['system:post:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['岗位']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
